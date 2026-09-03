package input

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/xamx2/xplit-server/model"
)

type TransactionInput struct {
	Amount      float64
	Description graphql.NullString
	MemberID    int32
	Splits      []TransactionSplitInput
}

func (i *TransactionInput) Decode(t *model.Transaction) error {
	t.Amount = i.Amount
	t.Description = i.Description.Value
	t.MemberID = i.MemberID

	for _, s := range i.Splits {
		ts := new(model.TransactionSplit)
		if err := s.Decode(ts); err != nil {
			return err
		}
		t.NewSplits = append(t.NewSplits, ts)
	}

	return nil
}

type TransactionSplitInput struct {
	MemberID int32
	Amount   float64
}

func (i TransactionSplitInput) Decode(ts *model.TransactionSplit) error {
	ts.Amount = i.Amount
	ts.MemberID = i.MemberID
	return nil
}
