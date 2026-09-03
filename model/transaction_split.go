package model

import (
	"context"

	"github.com/uptrace/bun"
)

type TransactionSplit struct {
	bun.BaseModel `bun:"alias:ts"`

	TransactionID int32   `bun:"transaction_id,pk"`
	MemberID      int32   `bun:"member_id,pk"`
	Amount        float64 `bun:"amount,notnull"`
}

func (t TransactionSplit) Member(db bun.IDB, ctx context.Context) (m *GroupMember, err error) {
	m = &GroupMember{ID: t.MemberID}
	err = db.NewSelect().Model(m).WherePK().Scan(ctx)
	return
}
