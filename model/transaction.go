package model

import (
	"context"
	"errors"
	"time"

	"github.com/uptrace/bun"
)

var (
	ErroTotalAmountsNotMatch = errors.New("total amounts not match")
)

type Transaction struct {
	bun.BaseModel `bun:"alias:t"`

	ID          int32     `bun:"id,pk,autoincrement"`
	Amount      float64   `bun:"amount,notnull"`
	Description *string   `bun:"description"`
	MemberID    int32     `bun:"member_id,notnull"`
	CreatedAt   time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp"`

	NewSplits []*TransactionSplit `bun:"-"`
}

func (t Transaction) Splits(db bun.IDB, ctx context.Context) (tss []TransactionSplit, err error) {
	err = db.NewSelect().Model(&tss).Where("transaction_id = ?", t.ID).Scan(ctx)
	return
}

func (t Transaction) Member(db bun.IDB, ctx context.Context) (m *GroupMember, err error) {
	m = &GroupMember{ID: t.MemberID}
	err = db.NewSelect().Model(m).WherePK().Scan(ctx)
	return
}

var _ bun.BeforeAppendModelHook = (*Transaction)(nil)

func (t *Transaction) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		var total float64
		for _, s := range t.NewSplits {
			total += s.Amount
		}
		if total != t.Amount {
			return ErroTotalAmountsNotMatch
		}
	}
	return nil
}
