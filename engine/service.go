package engine

import (
	"strings"

	"github.com/pkg/errors"
	"gopkg.in/reform.v1"
)

func NewService(db *reform.DB) *Service {
	return &Service{
		db: db,
	}
}

type Service struct {
	db *reform.DB
	tp *transactionProcessor
}

// NewInvoice новый счет на оплату
func (s *Service) NewInvoice(key, strategy string, meta, payload *[]byte) (int64, error) {
	key = strings.TrimSpace(strings.ToLower(key))
	strategy = strings.TrimSpace(strings.ToLower(strategy))

	newInvoice := &Invoice{
		Key:      key,
		Strategy: strategy,
		Meta:     meta,
		Payload:  payload,
	}
	if err := s.db.Insert(newInvoice); err != nil {
		return 0, errors.Wrap(err, "failed insert new invoice")
	}
	return newInvoice.InvoiceID, nil
}

func (s *Service) AddTransaction(invoiceID int64, key, provider string, meta *[]byte) (int64, error) {
	invoice := &Invoice{InvoiceID: invoiceID}
	if err := s.db.Reload(invoice); err != nil {
		return 0, errors.Wrap(err, "failed find invocie by ID")
	}
	if !invoice.Status.Match(DRAFT_I) {
		return 0, errors.New("not allowed add transaction (invoice is not draft)")
	}

	newTransaction := &Transaction{
		InvoiceID: invoiceID,

		// TODO
		// - внутренний перевод
		// - перевод с карты
		Provider: Provider(provider), // TODO: валидация провайдера

		Meta: meta,
		// TODO: add key field
	}
	if err := s.db.Insert(newTransaction); err != nil {
		return 0, errors.Wrap(err, "failed insert new transaction")
	}
	return newTransaction.TransactionID, nil
}

func (s *Service) AuthInvoice(invoiceID int64) error {
	panic("not implemented")
}

func (s *Service) RejectInvoice(invoiceID int64) error {
	panic("not implemented")
}

func (s *Service) AcceptInvoice(invoiceID int64) error {
	panic("not implemented")
}

func (s *Service) RejectTx(txID int64) error {
	panic("not implemented")
}

func (s *Service) AcceptTx(txID int64) error {
	panic("not implemented")
}

// func (s *Service) InternalTransfer(srcAccID, dstAccID, amount int64) (int64, error) {
// 	var newInvoiceID, newTransactionID int64

// 	err := s.db.InTransaction(func(tx *reform.TX) error {
// 		newInvoice := &Invoice{
// 			Key:         "simple1",
// 			Strategy:    "simple",
// 			TotalAmount: amount,
// 		}
// 		if err := s.db.Insert(newInvoice); err != nil {
// 			return errors.Wrap(err, "failed insert new invoice")
// 		}

// 		newInvoiceID = newInvoice.InvoiceID
// 		newTransaction := &Transaction{
// 			Provider:  "internal",
// 			InvoiceID: newInvoiceID,
// 		}
// 		if err := s.db.Insert(newTransaction); err != nil {
// 			return errors.Wrap(err, "failed insert new transaction")
// 		}
// 		newTransactionID = newTransaction.TransactionID

// 		opers := []*Operation{
// 			{
// 				SrcAccID: srcAccID,
// 				DstAccID: dstAccID,
// 				Amount:   amount,
// 				Strategy: SIMPLE_OPS,
// 			},
// 		}
// 		for _, oper := range opers {
// 			oper.TransactionID = newTransaction.TransactionID
// 			oper.InvoiceID = newTransaction.InvoiceID
// 			if err := s.db.Insert(oper); err != nil {
// 				return errors.Wrap(err, "failed insert new operation")
// 			}
// 		}
// 		return nil
// 	})

// 	if err != nil {
// 		return 0, err
// 	}

// 	return newInvoiceID, nil
// }

// type TransactionProcessor interface {
// 	Process(txID int64, updatedAt time.Time, currentStatus, nextStatus TransactionStatus)
// }