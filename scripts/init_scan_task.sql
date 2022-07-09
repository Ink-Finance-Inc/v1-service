use db_inkfinance;
truncate contract_events_progress;
INSERT INTO `contract_events_progress` (`contract_address`, `events`, `start_height`, `current_height`, `valid`, `update_time`, `comments`)
VALUES
	('0xDBCF5a529EE7C4190fb0a27f5c04AcC8914C4800', 'DAO_CREATE_MONITOR', 0, 0, 1, now(), ''),
	('0x3B727529F7FB07E93aa6b6b20b8A0516701272C3', 'INCOME_REPORT', 0, 0, 1, now(), ''),
	('0x63A73397d8749203708f8fE4bbc248BA5254cC24', 'ProposalDecision', 0, 0, 1, now(), '');
