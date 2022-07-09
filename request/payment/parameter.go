package account

type PayrollQueryRequest struct {
	DAOAddress string `json:"daoAddress"`
	ScheduleID int32  `json:"scheduleID"`
}

type DAOCollectableRequest struct {
	DAOAddress    string `json:"daoAddress"`
	WalletAddress string `json:"walletAddress"`
}

type AuditReportRequest struct {
	DAOAddress  string `json:"daoAddress"`
	ProposalID  string `json:"proposalID"`
	UserAddress string `json:"userAddress"`
	Page        int    `json:"page"`
	Size        int    `json:"size"`
}
type DAOScheduleRequest struct {
	DAOAddress    string `json:"daoAddress"`
	ScheduleType  int32  `json:"scheduleType"`
	MemberAddress string `json:"memberAddress"`
}
type AddressCollectableResponse struct {
	Collectable bool `json:"collectable"`
}

type TimeQueryRequest struct {
	DAOAddress    string `json:"daoAddress"`
	ScheduleID    int32  `json:"scheduleID"`
	WalletAddress string `json:"walletAddress"`
}
type ScheduleInfo struct {
	ScheduleID int32 `json:"scheduleID"`
	Duration   int64 `json:"duration"`
	StartTime  int64 `json:"startTime"`
}

type DAOPayMember struct {
	ScheduleID int32 `json:"scheduleID"`
	Duration   int64 `json:"duration"`
	StartTime  int64 `json:"startTime"`
}

type ScheduleList struct {
	ScheduleList []ScheduleInfo `json:"scheduleList"`
}

type SchedulePayTimeList struct {
	SchedulePayTimeList []SchedulePayTime `json:"schedulePayTimeList"`
}

type SchedulePayTime struct {
	ScheduleID  int32        `json:"scheduleID"`
	TimeID      int64        `json:"timeID"`
	PayTime     int64        `json:"payTime"`
	TimeSigners []SignerInfo `json:"timeSigners"`
}

type PayrollGroup struct {
	GroupList []PayrollList
}

type PayrollList struct {
	Title      string    `json:"title"`
	MemberList []Members `json:"list"`
}

type Members struct {
	FullAccount  string `json:"fullAccount"`
	MediaType    string `json:"mediaType"`
	Payee        string `json:"payee"`
	PayeeAddress string `json:"payeeAddres"`
	Amount       string `json:"amount"`
	TokenAddress string `json:"tokenAddress"`
	Description  string `json:"desc"`
	Title        string `json:"title"`
}

type SignerInfo struct {
	FullAccount   string `json:"fullAccount"`
	SignerAddress string `json:"signerAddress"`
	SignDate      string `json:"signDate"`
	BlockSignTime uint64 `json:"blockSignTime"`
}
type Result struct {
	Success int64 `json:"success"`
}

type DAOSignInfo struct {
	ScheduleSignList []ScheduleSignGroupInfo `json:"scheduleSignList"`
}

type ScheduleSignGroupInfo struct {
	ScheduleID   int32          `json:"scheduleID"`
	TimeSignList []TimeSignList `json:"scheduleSignList"`
}

type TimeSignList struct {
	TimeID   int32      `json:"timeID"`
	SignList []SignInfo `json:"scheduleSignList"`
}

type WithdrawRecords struct {
	TimeAmount   int32         `json:"timeAmount"`
	WithdrawLogs []WithdrawLog `json:"withdrawLogs"`
}

type WithdrawLog struct {
	TokenAddress string `json:"tokenAddress"`
	ClaimAmount  string `json:"claimAmount"`
	ClaimTime    int64  `json:"claimTime"`
	ToTimeID     int32  `json:"timeID"`
	ScheduleID   int32  `json:"scheduleID"`
}

type SignInfo struct {
	FullAccount   string `json:"fullAccount"`
	MediaType     string `json:"mediaType"`
	MediaAccount  string `json:"mediaAccount"`
	TimeID        int32  `json:"timeID"`
	PayTime       string `json:"payTime"`
	BlockPayTime  int64  `json:"blockPayTime"`
	SignTime      string `json:"signTime"`
	Signer        string `json:"signer"`
	BlockSignTime uint64 `json:"blockSignTime"`
}

type AuditReportDetail struct {
	AuditerMedia     string `json:"auditerMedia"`
	AuditerMediaType string `json:"auditerMediaType"`
	AuditerFullInfo  string `json:"auditerFullInfo"`
	AuditerAddr      string `json:"auditerAddr"`
	SignDate         int32  `json:"signDate"`
	IssueID          int32  `json:"issueID"`
	IsAudited        int32  `json:"isAudited"`
	Data             []byte `json:"data"`
	StringData       string `json:"stringData"`
	AuditPeriodTime  int64  `json:"auditPeriodTime"`
	Duration         int64  `json:"duration"`
}

type AuditReportGroup struct {
	AuditSchedule   uint32              `json:"auditSchedule"`
	AuditReportList []AuditReportDetail `json:"auditReportList"`
}
type AuditReportResp struct {
	AuditReportTabs  []AuditReportDetail `json:"auditReportTabs"`
	AuditReportGroup []AuditReportGroup  `json:"auditReporGroup"`
}
