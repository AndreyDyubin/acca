package tests

import (
	"testing"
)

func Test01Basic_01SimpleTrasnferWithoutHold(t *testing.T) {
	cur := "curr"

	tests := []cmdBatch{
		{
			"InternalTransfer",
			[]command{
				CmdInitAccounts(cur, []accountInfo{
					{
						AccKey:  "1",
						Balance: 10,
					},
					{
						AccKey:  "2",
						Balance: 20,
					},
					{
						AccKey:  "3",
						Balance: 30,
					},
				}),
				CmdTransfers([]transfers{
					transfers{
						{
							SrcAcc: "1",
							DstAcc: "2",
							Type:   Internal,
							Amount: 9,
							Reason: "fortesting",
							Meta: MetaData{
								"foo": "bar",
							},
							Hold: false,
						},
						{
							SrcAcc: "2",
							DstAcc: "3",
							Type:   Internal,
							Amount: 29,
							Reason: "fortesting",
							Meta: MetaData{
								"foo": "bar",
							},
							Hold: false,
						},
					},
				}),
				CmdCheckStatuses("draft"),
				CmdExecute(1),
				CmdCheckStatuses("accepted"),
				CmdCheckBalances(map[string]int64{
					"1": 10 - 9,
					"2": 20 + 9 - 29,
					"3": 30 + 29,
				}),
			},
		},
		{
			"InternalTransfer_Error_NotEnoughMoney",
			[]command{
				CmdInitAccounts(cur, []accountInfo{
					{
						AccKey:  "1",
						Balance: 10,
					},
					{
						AccKey:  "2",
						Balance: 20,
					},
					{
						AccKey:  "3",
						Balance: 30,
					},
				}),
				CmdTransfers([]transfers{
					{
						{
							SrcAcc: "1",
							DstAcc: "2",
							Type:   Internal,
							Amount: 1,
							Reason: "fortesting",
							Meta: MetaData{
								"foo": "bar",
							},
							Hold: false,
						},
						{
							SrcAcc: "2",
							DstAcc: "3",
							Type:   Internal,
							Amount: 1000,
							Reason: "fortesting",
							Meta: MetaData{
								"foo": "bar",
							},
							Hold: false,
						},
					},
				}),
				CmdCheckStatuses("draft"),
				CmdExecute(1),
				CmdCheckStatuses("accepted"),
				CmdCheckBalances(map[string]int64{
					"1": 9,
					"2": -979,
					"3": 1030,
				}),
			},
		},
		{
			"InternalTransfer_EmptyListTransfers",
			[]command{
				CmdInitAccounts(cur, []accountInfo{
					{
						AccKey:  "1",
						Balance: 10,
					},
					{
						AccKey:  "2",
						Balance: 20,
					},
					{
						AccKey:  "3",
						Balance: 30,
					},
				}),
				CmdTransfers([]transfers{
					transfers{}, // empty list transaction
				}),
				CmdCheckStatuses("draft"),
				CmdExecute(1),
				CmdCheckStatuses("accepted"),
				CmdCheckBalances(map[string]int64{
					"1": 10,
					"2": 20,
					"3": 30,
				}),
			},
		},

		{
			"InternalTransfer_Recharge",
			[]command{
				CmdInitAccounts(cur, []accountInfo{
					{
						AccKey:  "payment_gateway",
						Balance: 0,
					},
					{
						AccKey:  "client",
						Balance: 0,
					},
				}),
				CmdTransfers([]transfers{
					transfers{
						{
							SrcAcc: "payment_gateway",
							DstAcc: "client",
							Type:   Recharge,
							Amount: 102,
							Reason: "fortesting",
							Meta: MetaData{
								"foo": "bar",
							},
							Hold: false,
						},
					},
				}),
				CmdCheckStatuses("draft"),
				CmdExecute(1),
				CmdCheckStatuses("accepted"),
				CmdCheckBalances(map[string]int64{
					"payment_gateway": 102,
					"client":          102,
				}),
			},
		},

		{
			"InternalTransfer_Withdraw",
			[]command{
				CmdInitAccounts(cur, []accountInfo{
					{
						AccKey:  "payment_gateway",
						Balance: 100,
					},
					{
						AccKey:  "client",
						Balance: 10,
					},
				}),
				CmdTransfers([]transfers{
					transfers{
						{
							SrcAcc: "client",
							DstAcc: "payment_gateway",
							Type:   Withdraw,
							Amount: 10,
							Reason: "fortesting",
							Meta: MetaData{
								"foo": "bar",
							},
							Hold: false,
						},
					},
				}),
				CmdCheckStatuses("draft"),
				CmdExecute(1),
				CmdCheckStatuses("accepted"),
				CmdCheckBalances(map[string]int64{
					"payment_gateway": 90,
					"client":          0,
				}),
			},
		},

		{
			"InternalTransfer_WithdrawError1_NotEnoughMoneyFromSysAccount",
			[]command{
				CmdInitAccounts(cur, []accountInfo{
					{
						AccKey:  "payment_gateway",
						Balance: 0,
					},
					{
						AccKey:  "client",
						Balance: 10,
					},
				}),
				CmdTransfers([]transfers{
					transfers{
						{
							SrcAcc: "client",
							DstAcc: "payment_gateway",
							Type:   Withdraw,
							Amount: 10,
							Reason: "fortesting",
							Meta: MetaData{
								"foo": "bar",
							},
							Hold: false,
						},
					},
				}),
				CmdCheckStatuses("draft"),
				CmdExecute(1),
				CmdCheckStatuses("accepted"),
				CmdCheckBalances(map[string]int64{
					"payment_gateway": -10,
					"client":          0,
				}),
			},
		},
	}

	runTests(t, tests)
}
func Test01Basic_02SimpleTransferWithHold(t *testing.T) {
	cur := "curr"

	tests := []cmdBatch{
		{
			"InternalTransferWithHold",
			[]command{
				CmdInitAccounts(cur, []accountInfo{
					{
						AccKey:  "1",
						Balance: 10,
					},
					{
						AccKey:  "2",
						Balance: 20,
					},
					{
						AccKey:  "3",
						Balance: 30,
					},
					{
						AccKey:  "hold1",
						Balance: 0,
					},
				}),
				CmdTransfers([]transfers{
					transfers{
						{
							SrcAcc: "1",
							DstAcc: "2",
							Type:   Internal,
							Amount: 9,
							Reason: "fortesting",
							Meta: MetaData{
								"foo": "bar",
							},
							Hold:    true,
							HoldAcc: "hold1",
						},
						{
							SrcAcc: "2",
							DstAcc: "3",
							Type:   Internal,
							Amount: 19,
							Reason: "fortesting",
							Meta: MetaData{
								"foo": "bar",
							},
							Hold:    true,
							HoldAcc: "hold1",
						},
						{
							SrcAcc: "3",
							DstAcc: "1",
							Type:   Internal,
							Amount: 29,
							Reason: "fortesting",
							Meta: MetaData{
								"foo": "bar",
							},
							Hold:    true,
							HoldAcc: "hold1",
						},
					},
				}),
				CmdCheckStatuses("draft"),
				CmdExecute(1),
				CmdCheckStatuses("auth"),
				CmdCheckBalances(map[string]int64{
					"1":     10 - 9,
					"2":     20 - 19,
					"3":     30 - 29,
					"hold1": 9 + 19 + 29,
				}),
				CmdApprove(0),
				CmdExecute(1),
				CmdCheckStatuses("accepted"),
				CmdCheckBalances(map[string]int64{
					"1":     10 - 9 + 29,
					"2":     20 - 19 + 9,
					"3":     30 - 29 + 19,
					"hold1": 0,
				}),
			},
		},
		{
			"InternalTransferWithHold_NotEnoughMoney",
			[]command{
				CmdInitAccounts(cur, []accountInfo{
					{
						AccKey:  "1",
						Balance: 10,
					},
					{
						AccKey:  "2",
						Balance: 20,
					},
					{
						AccKey:  "3",
						Balance: 30,
					},
					{
						AccKey:  "hold1",
						Balance: 0,
					},
					{
						AccKey:  "other1",
						Balance: 0,
					},
				}),
				CmdTransfers([]transfers{
					transfers{
						{
							SrcAcc: "1",
							DstAcc: "2",
							Type:   Internal,
							Amount: 9,
							Reason: "fortesting",
							Meta: MetaData{
								"foo": "bar",
							},
							Hold:    true,
							HoldAcc: "hold1",
						},
						{
							SrcAcc: "2",
							DstAcc: "3",
							Type:   Internal,
							Amount: 19,
							Reason: "fortesting",
							Meta: MetaData{
								"foo": "bar",
							},
							Hold:    true,
							HoldAcc: "hold1",
						},
						{
							SrcAcc: "3",
							DstAcc: "1",
							Type:   Internal,
							Amount: 29,
							Reason: "fortesting",
							Meta: MetaData{
								"foo": "bar",
							},
							Hold:    true,
							HoldAcc: "hold1",
						},
					},
					transfers{
						{
							SrcAcc: "hold1",
							DstAcc: "other1",
							Type:   Internal,
							Amount: 9 + 19 + 29,
							Reason: "fortesting",
							Meta: MetaData{
								"foo": "bar",
							},
							Hold: false,
						},
					},
				}),
				CmdCheckStatuses("draft", "draft"),
				CmdExecute(2),
				CmdCheckStatuses("auth", "accepted"),
				CmdCheckBalances(map[string]int64{
					"1":      10 - 9,
					"2":      20 - 19,
					"3":      30 - 29,
					"hold1":  0,
					"other1": 9 + 19 + 29,
				}),
				CmdApprove(0),
				CmdExecute(1),
				CmdCheckStatuses("accepted", "accepted"),
				CmdCheckBalances(map[string]int64{
					"1":      10 - 9 + 29,
					"2":      20 + 9 - 19,
					"3":      30 + 19 - 29,
					"hold1":  0 - 9 - 19 - 29,
					"other1": 9 + 19 + 29,
				}),
				//CmdRollback(0), // not enough money
				//CmdExecute(1),
				//CmdCheckStatuses("accepted", "accepted"),
				//CmdCheckBalances(map[string]int64{
				//	"1":      10 - 9,
				//	"2":      20 - 19,
				//	"3":      30 - 29,
				//	"hold1":  0,
				//	"other1": 9 + 19 + 29,
				//}),
				//
				CmdTransfers([]transfers{
					transfers{
						{
							SrcAcc: "other1",
							DstAcc: "hold1",
							Type:   Internal,
							Amount: 9 + 19 + 29,
							Reason: "fortesting",
							Meta: MetaData{
								"foo": "bar",
							},
							Hold: false,
						},
					},
				}),
				CmdExecute(1),
				CmdCheckStatuses("accepted", "accepted", "accepted"),
				CmdCheckBalances(map[string]int64{
					"1":      30,
					"2":      10,
					"3":      20,
					"hold1":  0,
					"other1": 0,
				}),

				//CmdRollback(0), // not enough money
				CmdExecute(1),

				CmdCheckStatuses("accepted", "accepted", "accepted"),
				CmdCheckBalances(map[string]int64{
					"1":      30,
					"2":      10,
					"3":      20,
					"hold1":  0,
					"other1": 0,
				}),
			},
		},
	}

	runTests(t, tests)
}

func Test01Basic_03ErrorInMiddle(t *testing.T) {
	cur := "curr"

	tests := []cmdBatch{
		{
			"TestQueue",
			[]command{
				CmdInitAccounts(cur, []accountInfo{
					{
						AccKey:  "1",
						Balance: 10,
					},
					{
						AccKey:  "2",
						Balance: 20,
					},
					{
						AccKey:  "3",
						Balance: 30,
					},
					{
						AccKey:  "hold1",
						Balance: 0,
					},
				}),
				CmdTransfers([]transfers{
					transfers{
						{
							SrcAcc: "1",
							DstAcc: "2",
							Type:   Internal,
							Amount: 3,
							Reason: "fortesting",
							Meta: MetaData{
								"foo": "bar",
							},
							Hold:    true,
							HoldAcc: "hold1",
						},
					},
				}),
				CmdTransfers([]transfers{
					transfers{
						{
							SrcAcc: "2",
							DstAcc: "3",
							Type:   Internal,
							Amount: 1000000,
							Reason: "fortesting",
							Meta: MetaData{
								"foo": "bar",
							},
							Hold:    true,
							HoldAcc: "hold1",
						},
					},
				}),
				CmdTransfers([]transfers{
					transfers{
						{
							SrcAcc: "3",
							DstAcc: "1",
							Type:   Internal,
							Amount: 3,
							Reason: "fortesting",
							Meta: MetaData{
								"foo": "bar",
							},
							Hold:    true,
							HoldAcc: "hold1",
						},
					},
				}),
				CmdTransfers([]transfers{
					transfers{
						{
							SrcAcc: "1",
							DstAcc: "2",
							Type:   Internal,
							Amount: 1,
							Reason: "fortesting",
							Meta: MetaData{
								"foo": "bar",
							},
							Hold:    true,
							HoldAcc: "hold1",
						},
					},
				}),

				CmdCheckStatuses("draft", "draft", "draft", "draft"),
				CmdExecute(4),
				CmdCheckStatuses("auth", "auth", "auth", "auth"),
				CmdApprove(0, 2, 3),
				// CmdApprove(0, 2, 3), // idempotency?
				CmdExecute(3),
				CmdCheckStatuses("accepted", "auth", "accepted", "accepted"),
			},
		},
	}

	runTests(t, tests)
}

func Test01Basic_03SimpleTransferReject(t *testing.T) {
	cur := "curr"

	tests := []cmdBatch{
		{
			"InternalTransferWithHoldAndReject",
			[]command{
				CmdInitAccounts(cur, []accountInfo{
					{
						AccKey:  "11",
						Balance: 10,
					},
					{
						AccKey:  "22",
						Balance: 20,
					},
					{
						AccKey:  "33",
						Balance: 30,
					},
					{
						AccKey:  "hold11",
						Balance: 0,
					},
				}),
				CmdTransfers([]transfers{
					transfers{
						{
							SrcAcc: "11",
							DstAcc: "22",
							Type:   Internal,
							Amount: 9,
							Reason: "fortesting",
							Meta: MetaData{
								"foo": "bar",
							},
							Hold:    true,
							HoldAcc: "hold11",
						},
						{
							SrcAcc: "22",
							DstAcc: "33",
							Type:   Internal,
							Amount: 19,
							Reason: "fortesting",
							Meta: MetaData{
								"foo": "bar",
							},
							Hold:    true,
							HoldAcc: "hold11",
						},
						{
							SrcAcc: "33",
							DstAcc: "11",
							Type:   Internal,
							Amount: 29,
							Reason: "fortesting",
							Meta: MetaData{
								"foo": "bar",
							},
							Hold:    true,
							HoldAcc: "hold11",
						},
					},
				}),
				CmdCheckStatuses("draft"),
				CmdExecute(1),
				CmdCheckStatuses("auth"),
				CmdCheckBalances(map[string]int64{
					"11":     10 - 9,
					"22":     20 - 19,
					"33":     30 - 29,
					"hold11": 9 + 19 + 29,
				}),
				CmdReject(0),
				CmdExecute(1),
				CmdCheckStatuses("rejected"),
				CmdCheckBalances(map[string]int64{
					"11":     10,
					"22":     20,
					"33":     30,
					"hold11": 0,
				}),
			},
		},
		{
			"InternalTransferWithHoldAndRejectInDraft",
			[]command{
				CmdInitAccounts(cur, []accountInfo{
					{
						AccKey:  "11",
						Balance: 10,
					},
					{
						AccKey:  "22",
						Balance: 20,
					},
					{
						AccKey:  "33",
						Balance: 30,
					},
					{
						AccKey:  "hold11",
						Balance: 0,
					},
				}),
				CmdTransfers([]transfers{
					transfers{
						{
							SrcAcc: "11",
							DstAcc: "22",
							Type:   Internal,
							Amount: 9,
							Reason: "fortesting",
							Meta: MetaData{
								"foo": "bar",
							},
							Hold:    true,
							HoldAcc: "hold11",
						},
						{
							SrcAcc: "22",
							DstAcc: "33",
							Type:   Internal,
							Amount: 19,
							Reason: "fortesting",
							Meta: MetaData{
								"foo": "bar",
							},
							Hold:    true,
							HoldAcc: "hold11",
						},
						{
							SrcAcc: "33",
							DstAcc: "11",
							Type:   Internal,
							Amount: 29,
							Reason: "fortesting",
							Meta: MetaData{
								"foo": "bar",
							},
							Hold:    true,
							HoldAcc: "hold11",
						},
					},
				}),
				CmdCheckStatuses("draft"),
				CmdReject(0),
				// CmdReject(0), // idempotency?
				CmdExecute(1),
				CmdCheckStatuses("draft"),
				CmdCheckBalances(map[string]int64{
					"11":      10,
					"22":      20,
					"33":      30,
					"hold111": 0,
				}),
				CmdReject(0),
				// CmdReject(0), // idempotency?
				CmdExecute(1),
				CmdCheckStatuses("draft"),
				CmdCheckBalances(map[string]int64{
					"11":      10,
					"22":      20,
					"33":      30,
					"hold111": 0,
				}),
			},
		},
	}

	runTests(t, tests)
}
