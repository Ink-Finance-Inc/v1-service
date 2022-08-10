package services

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"
	"time"

	"gorm.io/gen"
	conf "inkfinance.xyz/conf"
	"inkfinance.xyz/dal/model"
	"inkfinance.xyz/dal/query"
	request "inkfinance.xyz/request/payment"
)

type IPaymentService interface {
	GetScheduleList(ctx context.Context, req *request.DAOScheduleRequest) (*request.ScheduleList, error)
	GetTimeList(ctx context.Context, req *request.TimeQueryRequest) (*request.SchedulePayTimeList, error)
	GetPayrollList(ctx context.Context, req *request.PayrollQueryRequest) (*request.PayrollGroup, error)
	AddressCollectable(ctx context.Context, req *request.DAOCollectableRequest) (*request.AddressCollectableResponse, error) // dao | wallet
	GetDAOSignInfo(ctx context.Context, req *request.DAOScheduleRequest) (*request.DAOSignInfo, error)
	GetScheduleSignInfo(ctx context.Context, req *request.TimeQueryRequest) (*request.ScheduleSignGroupInfo, error)
	GetScheduleWithdraws(ctx context.Context, req *request.TimeQueryRequest) (request.WithdrawRecords, error)
	// 每个阶段的领取列表
	// 监听用户领取的事件
	// 每个阶段的LIST - 领取的人
	// 每个阶段没签的人
}

type PaymentService struct{}

func NewPaymentService() IPaymentService {
	var service = &PaymentService{
		// middleware
	}
	return service
}

func (PaymentService) GetScheduleWithdraws(ctx context.Context, req *request.TimeQueryRequest) (request.WithdrawRecords, error) {
	db := query.Use(conf.DB)
	tbScheduleWithdrawLog := db.ScheduleWithdrawLog
	scheduleWithdrawLogDAO := tbScheduleWithdrawLog.WithContext(ctx)
	logs, err := scheduleWithdrawLogDAO.Where(tbScheduleWithdrawLog.DaoAddress.Eq(req.DAOAddress), tbScheduleWithdrawLog.ScheduleID.Eq(req.ScheduleID), tbScheduleWithdrawLog.MemberAddress.Eq(req.WalletAddress)).Find()

	var withdrawLogs []request.WithdrawLog
	if err == nil {
		// 提醒时间 + 数量 + token
		for i := 0; i < len(logs); i++ {
			withdrawLogs = append(withdrawLogs, request.WithdrawLog{TokenAddress: logs[i].TokenAddress, ClaimAmount: logs[i].ClaimAmount, ClaimTime: logs[i].CreateTime, ScheduleID: logs[i].ScheduleID, ToTimeID: logs[i].CurrentTimeID})
		}
		// fmt.Println(withdrawLogs)
	}

	tbScheduleInfo := db.ScheduleInfo
	scheduleInfos, err := tbScheduleInfo.WithContext(ctx).Where(tbScheduleInfo.DaoAddress.Eq(req.DAOAddress), tbScheduleInfo.ScheduleID.Eq(req.ScheduleID)).First()
	sort.Slice(withdrawLogs, func(i, j int) bool {
		return withdrawLogs[i].ClaimTime > withdrawLogs[j].ClaimTime
	})
	timeIDs := calculateID(scheduleInfos.StartTime, scheduleInfos.Duration)

	timeAmount := int32(len(timeIDs))
	if scheduleInfos.ScheduleTimes == 1 {
		timeAmount = 1
	}
	return request.WithdrawRecords{TimeAmount: timeAmount, WithdrawLogs: withdrawLogs}, nil
}

func calculateID(startTime int64, duration int64) []int32 {
	currentTime := time.Now().Unix()
	var timeIDs []int32
	var startTimeID int32 = 0
	for i := startTime; ; i += duration {
		// fmt.Println("startTime: ", i, ", duration: ", duration, " currentTime:", currentTime)
		if i <= currentTime {
			startTimeID++
			timeIDs = append(timeIDs, startTimeID)
		} else {
			break
		}
	}

	return timeIDs
}

func (PaymentService) GetScheduleSignInfo(ctx context.Context, req *request.TimeQueryRequest) (*request.ScheduleSignGroupInfo, error) {
	db := query.Use(conf.DB)
	tbScheduleInfo := db.ScheduleInfo
	scheduleInfos, err := tbScheduleInfo.WithContext(ctx).Where(tbScheduleInfo.DaoAddress.Eq(req.DAOAddress), tbScheduleInfo.ScheduleID.Eq(req.ScheduleID)).First()

	if err == nil {
		scheduleSignInfo, err := buildScheduleSignInfo(ctx, scheduleInfos)

		sort.Slice(scheduleSignInfo.TimeSignList, func(i, j int) bool {
			return scheduleSignInfo.TimeSignList[i].TimeID > scheduleSignInfo.TimeSignList[j].TimeID
		})
		return &scheduleSignInfo, err
	}

	return &request.ScheduleSignGroupInfo{}, nil
}

func (PaymentService) AddressCollectable(ctx context.Context, req *request.DAOCollectableRequest) (*request.AddressCollectableResponse, error) {
	db := query.Use(conf.DB)
	tbScheduleMemberInfo := db.ScheduleMemberInfo

	scheduleMemberInfoDAO := tbScheduleMemberInfo.WithContext(ctx)

	memberExit, err := scheduleMemberInfoDAO.Where(tbScheduleMemberInfo.DaoAddress.Eq(req.DAOAddress), tbScheduleMemberInfo.MemberAddr.Eq(req.WalletAddress)).Count()
	if err == nil && memberExit > 0 {
		return &request.AddressCollectableResponse{Collectable: true}, nil
	}
	return &request.AddressCollectableResponse{Collectable: false}, nil
}

func (PaymentService) GetScheduleList(ctx context.Context, req *request.DAOScheduleRequest) (*request.ScheduleList, error) {
	db := query.Use(conf.DB)
	tbScheduleInfo := db.ScheduleInfo
	var scheduleIDs []int32

	fmt.Println("GetScheduleList start: ", req)

	var conds []gen.Condition
	if len(req.MemberAddress) > 0 {
		tbScheduleMemberInfo := db.ScheduleMemberInfo
		err := tbScheduleMemberInfo.WithContext(ctx).Select(tbScheduleMemberInfo.ScheduleID).Where(tbScheduleMemberInfo.DaoAddress.Eq(req.DAOAddress), tbScheduleMemberInfo.MemberAddr.Eq(req.MemberAddress)).Scan(&scheduleIDs)
		if err == nil {

		}
		if len(scheduleIDs) > 0 {
			conds = append(conds, tbScheduleInfo.ScheduleID.In(scheduleIDs...))
		}
	}

	fmt.Println("GetScheduleList query is :", conds)

	if req.ScheduleType != 0 {
		conds = append(conds, tbScheduleInfo.ScheduleTimes.Eq(int64(req.ScheduleType)))
	}

	if req.DAOAddress != "" {
		conds = append(conds, tbScheduleInfo.DaoAddress.Eq(req.DAOAddress))

	}

	scheduleList, err := db.ScheduleInfo.WithContext(ctx).Where(conds...).Find()

	fmt.Println("GetScheduleList DAO:", conds)

	req = &request.DAOScheduleRequest{}

	if err == nil {
		var schedules []request.ScheduleInfo
		for i := 0; i < len(scheduleList); i++ {
			schedules = append(schedules, request.ScheduleInfo{ScheduleID: scheduleList[i].ScheduleID, StartTime: scheduleList[i].StartTime, Duration: scheduleList[i].Duration})
		}

		sort.Slice(schedules, func(i, j int) bool {
			return schedules[i].ScheduleID > schedules[j].ScheduleID
		})

		return &request.ScheduleList{ScheduleList: schedules}, err
	}

	return &request.ScheduleList{}, nil
}
func (PaymentService) GetTimeList(ctx context.Context, req *request.TimeQueryRequest) (*request.SchedulePayTimeList, error) {
	db := query.Use(conf.DB)
	tbScheduleInfo := db.ScheduleInfo
	scheduleInfo, err := tbScheduleInfo.WithContext(ctx).Where(tbScheduleInfo.DaoAddress.Eq(req.DAOAddress), tbScheduleInfo.ScheduleID.Eq(req.ScheduleID)).First()

	if err == nil {

		startTime := scheduleInfo.StartTime
		duration := scheduleInfo.Duration
		var timeIDs []int32
		var startTimeID int32 = 0
		currentTime := time.Now().Unix() + scheduleInfo.Duration
		var schedulePayTimes []request.SchedulePayTime

		if scheduleInfo.ScheduleTimes == 1 {
			// pay once
			schedulePayTimes = append(schedulePayTimes, request.SchedulePayTime{ScheduleID: req.ScheduleID, TimeID: 1, PayTime: startTime})
		} else {
			for i := startTime; ; i += duration {
				fmt.Println("startTime: ", i, ", duration: ", duration, " currentTime:", currentTime)
				if i <= currentTime {
					startTimeID++
					schedulePayTimes = append(schedulePayTimes, request.SchedulePayTime{ScheduleID: req.ScheduleID, TimeID: int64(startTimeID), PayTime: i})
					timeIDs = append(timeIDs, startTimeID)
				} else {

					break
				}
			}
		}

		sort.Slice(schedulePayTimes, func(i, j int) bool {
			return schedulePayTimes[i].TimeID > schedulePayTimes[j].TimeID
		})

		tbSchedulesignInfo := db.ScheduleSignedInfo
		tbSocialInfo := db.SocialInfo
		var signMembers []request.SignInfo
		err = tbSchedulesignInfo.WithContext(ctx).
			Select(tbSchedulesignInfo.TimeID.As("TimeID"), tbSocialInfo.SocialAccount.As("MediaAccount"), tbSocialInfo.MediaType.As("MediaType"), tbSchedulesignInfo.SignAddress.As("Signer"), tbSchedulesignInfo.CreateTime.As("SignTime"), tbSchedulesignInfo.BlockTime.As("BlockSignTime")).
			Join(tbSocialInfo, tbSocialInfo.MainAccounts.EqCol(tbSchedulesignInfo.SignAddress)).
			Where(tbSchedulesignInfo.DaoAddress.Eq(scheduleInfo.DaoAddress), tbSchedulesignInfo.ScheduleID.Eq(scheduleInfo.ScheduleID)).Scan(&signMembers)
		fmt.Println(signMembers)
		fmt.Println(schedulePayTimes)
		schedulePayTimes = bindSignMember(schedulePayTimes, signMembers)

		return &request.SchedulePayTimeList{SchedulePayTimeList: schedulePayTimes}, err
	}

	return &request.SchedulePayTimeList{}, nil
}

func bindSignMember(schedulePayTimes []request.SchedulePayTime, signMembers []request.SignInfo) []request.SchedulePayTime {
	for i := 0; i < len(schedulePayTimes); i++ {

		for j := 0; j < len(signMembers); j++ {
			fmt.Println("compare: ", schedulePayTimes[i].TimeID, int64(signMembers[j].TimeID))
			if schedulePayTimes[i].TimeID == int64(signMembers[j].TimeID) {

				signer := request.SignerInfo{}
				signer.FullAccount = signMembers[j].MediaAccount + "@" + signMembers[j].MediaType + "<" + signMembers[j].Signer + ">"
				signer.SignerAddress = signMembers[j].Signer
				signer.SignDate = signMembers[j].SignTime
				signer.BlockSignTime = signMembers[j].BlockSignTime

				schedulePayTimes[i].TimeSigners = append(schedulePayTimes[i].TimeSigners, signer)
			}
		}
	}

	return schedulePayTimes

}

func (PaymentService) GetPayrollList(ctx context.Context, req *request.PayrollQueryRequest) (*request.PayrollGroup, error) {
	db := query.Use(conf.DB)
	tbScheduleMemberInfo := db.ScheduleMemberInfo
	tbSocialInfo := db.SocialInfo

	var members []request.Members
	err := tbScheduleMemberInfo.WithContext(ctx).
		Select(tbScheduleMemberInfo.ListTitle.As("Title"), tbSocialInfo.SocialAccount.As("Payee"), tbSocialInfo.MediaType.As("MediaType"), tbScheduleMemberInfo.MemberAddr.As("PayeeAddress"), tbScheduleMemberInfo.OncePay.As("Amount"), tbScheduleMemberInfo.Description.As("Description"), tbScheduleMemberInfo.TokenAddr.As("TokenAddress")).
		LeftJoin(tbSocialInfo, tbSocialInfo.MainAccounts.EqCol(tbScheduleMemberInfo.MemberAddr)).
		Where(tbScheduleMemberInfo.DaoAddress.Eq(req.DAOAddress), tbScheduleMemberInfo.ScheduleID.Eq(req.ScheduleID)).Scan(&members)

	if err != nil {
		return &request.PayrollGroup{}, err
	}

	// group
	groupMap := map[string][]request.Members{}
	for i := 0; i < len(members); i++ {
		if members[i].MediaType == "" {
			members[i].FullAccount = members[i].PayeeAddress
		} else {
			members[i].FullAccount = members[i].Payee + "@" + members[i].MediaType + "<" + members[i].PayeeAddress + ">"
		}
		key := members[i].Title
		theGroup, exist := groupMap[key]
		if exist {
			theGroup = append(theGroup, members[i])
			groupMap[key] = theGroup
		} else {
			var membersArray []request.Members
			membersArray = append(membersArray, members[i])
			groupMap[key] = membersArray
		}

	}

	var payrollGroupList []request.PayrollList
	for groupKey := range groupMap {

		payrollList := request.PayrollList{}
		payrollList.Title = groupKey
		payrollList.MemberList = groupMap[groupKey]
		payrollGroupList = append(payrollGroupList, payrollList)

	}

	return &request.PayrollGroup{GroupList: payrollGroupList}, nil
}

func (PaymentService) GetDAOSignInfo(ctx context.Context, req *request.DAOScheduleRequest) (*request.DAOSignInfo, error) {
	db := query.Use(conf.DB)
	tbScheduleInfo := db.ScheduleInfo
	scheduleInfos, err := tbScheduleInfo.WithContext(ctx).Where(tbScheduleInfo.DaoAddress.Eq(req.DAOAddress), tbScheduleInfo.ScheduleTimes.Eq(int64(req.ScheduleType))).Find()
	var scheduleSignInfoList []request.ScheduleSignGroupInfo
	if err == nil {
		for i := 0; i < len(scheduleInfos); i++ {
			fmt.Println(scheduleInfos[i])
			scheduleSignInfo, err := buildScheduleSignInfo(ctx, scheduleInfos[i])
			fmt.Println(err, scheduleSignInfo)
			if err == nil {
				scheduleSignInfoList = append(scheduleSignInfoList, scheduleSignInfo)
			}
		}
		return &request.DAOSignInfo{ScheduleSignList: scheduleSignInfoList}, err
	}
	return &request.DAOSignInfo{}, nil
}

func buildScheduleSignInfo(ctx context.Context, scheduleInfo *model.ScheduleInfo) (request.ScheduleSignGroupInfo, error) {

	startTime := scheduleInfo.StartTime
	duration := scheduleInfo.Duration
	var timeIDs []int32
	var startTimeID int32 = 0
	currentTime := time.Now().Unix() + scheduleInfo.Duration
	var schedulePayTimes []request.SchedulePayTime
	// 组装
	if scheduleInfo.ScheduleTimes == 1 {
		// pay once
		schedulePayTimes = append(schedulePayTimes, request.SchedulePayTime{ScheduleID: scheduleInfo.ScheduleID, TimeID: 1, PayTime: startTime})
	} else {
		for i := startTime; ; i += duration {
			// fmt.Println("startTime: ", i, ", duration: ", duration, " currentTime:", currentTime)
			if i <= currentTime {
				startTimeID++
				schedulePayTimes = append(schedulePayTimes, request.SchedulePayTime{ScheduleID: scheduleInfo.ScheduleID, TimeID: int64(startTimeID), PayTime: i})
				timeIDs = append(timeIDs, startTimeID)
			} else {
				break
			}
		}
	}

	fmt.Println(schedulePayTimes)

	// 针对于Times，组装每个TimeID签名人的数据
	// var timeSignList []request.TimeSignInfo
	// 查询出来每个
	db := query.Use(conf.DB)
	tbSchedulesignInfo := db.ScheduleSignedInfo
	tbSocialInfo := db.SocialInfo

	var scheduleSignGroup request.ScheduleSignGroupInfo

	scheduleSignGroup.ScheduleID = scheduleInfo.ScheduleID

	var signMembers []request.SignInfo
	err := tbSchedulesignInfo.WithContext(ctx).
		Select(tbSchedulesignInfo.TimeID.As("TimeID"), tbSocialInfo.SocialAccount.As("MediaAccount"), tbSocialInfo.MediaType.As("MediaType"), tbSchedulesignInfo.SignAddress.As("Signer"), tbSchedulesignInfo.CreateTime.As("SignTime"), tbSchedulesignInfo.BlockTime.As("BlockSignTime")).
		Join(tbSocialInfo, tbSocialInfo.MainAccounts.EqCol(tbSchedulesignInfo.SignAddress)).
		Where(tbSchedulesignInfo.DaoAddress.Eq(scheduleInfo.DaoAddress), tbSchedulesignInfo.ScheduleID.Eq(scheduleInfo.ScheduleID)).Scan(&signMembers)

	if err == nil {
		for i := 0; i < len(signMembers); i++ {
			signMembers[i].FullAccount = signMembers[i].MediaAccount + "@" + signMembers[i].MediaType + "<" + signMembers[i].Signer + ">"
			signMembers[i].PayTime = time.Unix(getPayTime(signMembers[i].TimeID, schedulePayTimes), 0).Format("2006-01-02 15:04:05")
			signMembers[i].BlockPayTime = getPayTime(signMembers[i].TimeID, schedulePayTimes)
			signMembers[i].SignTime = signMembers[i].SignTime[0:10] + " " + signMembers[i].SignTime[11:19]

			exist := false
			for j := 0; j < len(scheduleSignGroup.TimeSignList); j++ {
				if scheduleSignGroup.TimeSignList[j].TimeID == signMembers[i].TimeID {
					scheduleSignGroup.TimeSignList[j].SignList = append(scheduleSignGroup.TimeSignList[j].SignList, signMembers[i])
					exist = true
				}
			}

			if !exist {
				g := request.TimeSignList{TimeID: signMembers[i].TimeID}
				g.SignList = append(g.SignList, signMembers[i])
				scheduleSignGroup.TimeSignList = append(scheduleSignGroup.TimeSignList, g)
			}
		}
	}

	return scheduleSignGroup, nil
}

func getPayTime(timeID int32, schedulePayTimes []request.SchedulePayTime) int64 {
	for i := 0; i < len(schedulePayTimes); i++ {
		if timeID == int32(schedulePayTimes[i].TimeID) {
			return schedulePayTimes[i].PayTime
		}
	}
	return 0
}

func (PaymentService) GetAuditReportContent(ctx context.Context, req *request.AuditReportRequest) (*request.AuditReportResp, error) {

	db := query.Use(conf.DB)
	tbAuditReport := db.AuditReport
	tbSocialInfo := db.SocialInfo
	tbProposal := db.ProposalDecision

	proposal, err := tbProposal.WithContext(ctx).Where(tbProposal.ProposalID.Eq(req.ProposalID)).First()
	if err != nil {
		return nil, err
	}

	if proposal.AuditPeriod <= 0 {
		return nil, errors.New("no duration")
	}

	var reportDetails []request.AuditReportDetail
	err = tbAuditReport.WithContext(ctx).
		Select(tbAuditReport.ReportContent.As("Data"), tbAuditReport.OperatorAddress.As("AuditerAddr"), tbSocialInfo.SocialAccount.As("AuditerMedia"), tbSocialInfo.MediaType.As("AuditerMediaType"), tbAuditReport.CreateTime.As("SignDate")).
		Join(tbSocialInfo, tbSocialInfo.MainAccounts.EqCol(tbAuditReport.OperatorAddress)).Order(tbAuditReport.CreateTime.Desc()).Where(tbAuditReport.DaoAddress.Eq(req.DAOAddress)).Scan(&reportDetails)

	//.ScanByPage(&reportDetails, (req.Page-1)*req.Size, req.Size)
	// fmt.Println(count)
	// startTime := proposal.AgreeTime
	// duration := int64(proposal.AuditPeriod)
	// currentTime := time.Now().Unix()
	// currentTimeID := 0
	// var auditReportTab []request.AuditReportDetail
	startTime := proposal.AgreeTime
	duration := int64(proposal.AuditPeriod)
	for i := 0; i < len(reportDetails); i++ {

		reportDetails[i].AuditerFullInfo = reportDetails[i].AuditerMedia + "@" + reportDetails[i].AuditerMediaType + "<" + reportDetails[i].AuditerAddr + ">"
		reportDetails[i].IssueID = int32((int64(reportDetails[i].SignDate)-startTime)/duration + 1)
		reportDetails[i].IsAudited = 1

		reportDetails[i].StringData = string(reportDetails[i].Data)[64:]
		reportDetails[i].StringData = reportDetails[i].StringData[:strings.LastIndex(reportDetails[i].StringData, "}")+1]
		reportDetails[i].Data = nil
	}

	// sort.Slice(reportDetails, func(i, j int) bool {
	// 	return reportDetails[i].CreateTime > reportDetails[j].CreateTime
	// })

	auditTimes := calculateID(proposal.AgreeTime, int64(proposal.AuditPeriod))

	sort.Slice(auditTimes, func(i, j int) bool {
		return i > j
	})

	var auditScheduleGroups []request.AuditReportGroup

	for i := 0; i < len(auditTimes); i++ {

		scheduleGroup := request.AuditReportGroup{AuditSchedule: uint32(auditTimes[i])}

		scheduleReports := findReports(auditTimes[i], reportDetails)

		scheduleGroup.AuditReportList = scheduleReports

		sort.Slice(scheduleGroup.AuditReportList, func(i2, j int) bool {
			return scheduleGroup.AuditReportList[i2].SignDate > scheduleGroup.AuditReportList[j].SignDate
		})

		auditScheduleGroups = append(auditScheduleGroups, scheduleGroup)
	}

	sort.Slice(auditScheduleGroups, func(i, j int) bool {
		return auditScheduleGroups[i].AuditSchedule > auditScheduleGroups[j].AuditSchedule
	})

	return &request.AuditReportResp{AuditReportGroup: auditScheduleGroups}, nil
}

func findReports(auditTime int32, reportDetails []request.AuditReportDetail) []request.AuditReportDetail {
	var findedReport []request.AuditReportDetail

	for i := 0; i < len(reportDetails); i++ {

		if reportDetails[i].IssueID == auditTime {
			findedReport = append(findedReport, reportDetails[i])
		}

	}

	return findedReport

}

func (PaymentService) GetAuditReportTab(ctx context.Context, req *request.AuditReportRequest) (*request.AuditReportResp, error) {

	db := query.Use(conf.DB)
	tbAuditReport := db.AuditReport

	tbProposal := db.ProposalDecision

	proposal, err := tbProposal.WithContext(ctx).Where(tbProposal.ProposalID.Eq(req.ProposalID)).First()
	if err != nil {
		return nil, err
	}
	if proposal.AuditPeriod <= 0 {
		return nil, errors.New("no duration")
	}
	var conds []gen.Condition
	conds = append(conds, tbAuditReport.DaoAddress.Eq(req.DAOAddress))
	conds = append(conds, tbAuditReport.OperatorAddress.Eq(req.UserAddress))

	auditReports, _ := tbAuditReport.WithContext(ctx).Where(conds...).Find()

	startTime := proposal.AgreeTime
	duration := int64(proposal.AuditPeriod)
	currentTime := time.Now().Unix()
	currentTimeID := 0
	var auditReportTab []request.AuditReportDetail
	// fmt.Println(auditReports)
	for i := startTime; ; i += duration {
		fmt.Println("startTime: ", i, ", duration: ", duration, " currentTime:", currentTime)
		if i <= currentTime {
			currentTimeID++

			r, exist := getAuditReport(auditReports, startTime, duration, currentTimeID)

			if exist {

				detail := request.AuditReportDetail{IssueID: int32(currentTimeID), IsAudited: 1}
				detail.StringData = string(r.ReportContent)[64:]
				detail.Data = nil
				detail.Duration = duration
				detail.AuditPeriodTime = startTime + (int64(currentTimeID-1))*duration
				detail.StringData = detail.StringData[:strings.LastIndex(detail.StringData, "}")+1]
				detail.SignDate = int32(r.CreateTime)

				auditReportTab = append(auditReportTab, detail)
			} else {
				auditReportTab = append(auditReportTab, request.AuditReportDetail{IssueID: int32(currentTimeID), IsAudited: 0, Duration: duration, AuditPeriodTime: startTime + (int64(currentTimeID-1))*duration})
			}

		} else {
			fmt.Println("over --- ")
			break
		}
	}

	sort.Slice(auditReportTab, func(i, j int) bool {
		return auditReportTab[i].IssueID > auditReportTab[j].IssueID
	})
	return &request.AuditReportResp{AuditReportTabs: auditReportTab}, nil
}

func getAuditReport(auditReports []*model.AuditReport, startTime, duration int64, currentTimeID int) (model.AuditReport, bool) {
	for i := 0; i < len(auditReports); i++ {

		fmt.Println("compare: ", currentTimeID, " with: ", ((auditReports[i].CreateTime-startTime)/duration)+1)
		if (auditReports[i].CreateTime-startTime)/duration+1 == int64(currentTimeID) {
			return *auditReports[i], true
		}
	}

	return model.AuditReport{}, false
}
