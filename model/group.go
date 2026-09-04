package model

import (
	"context"

	"github.com/uptrace/bun"
)

type Group struct {
	bun.BaseModel `bun:"alias:g"`

	ID   int32  `bun:"id,pk,autoincrement"`
	Name string `bun:"name,notnull"`
}

func (g Group) Members(db bun.IDB, ctx context.Context) (mems []GroupMember, err error) {
	err = db.NewSelect().Model(&mems).Where("group_id = ?", g.ID).Scan(ctx)
	return
}

func (g Group) Transactions(db bun.IDB, ctx context.Context) (ts []Transaction, err error) {
	err = db.NewSelect().Model(&ts).
		Join("JOIN group_members AS gm ON gm.id = t.member_id").
		Where("gm.group_id = ?", g.ID).
		OrderBy("created_at", bun.OrderDesc).
		Scan(ctx)
	return
}

func (g Group) Settlements(db bun.IDB, ctx context.Context) (rs []Settlement, err error) {
	var ss []Settlement
	if err := db.NewSelect().Model((*TransactionSplit)(nil)).
		ColumnExpr("SUM(ts.amount * -1) AS amount, ts.member_id AS from_id, t.member_id AS to_id").
		Join("JOIN transactions AS t ON t.id = ts.transaction_id").
		Join("JOIN group_members AS gm ON gm.id = t.member_id").
		Where("gm.group_id = ? AND t.member_id != ts.member_id", g.ID).
		Group("from_id", "to_id").
		Scan(ctx, &ss); err != nil {
		return nil, err
	}

	maps := map[int32]map[int32]float64{}
	for _, s := range ss {
		if maps[s.ToID] != nil && maps[s.ToID][s.FromID] != 0 {
			maps[s.ToID][s.FromID] -= s.Amount
		} else {
			if maps[s.FromID] == nil {
				maps[s.FromID] = map[int32]float64{}
			}
			maps[s.FromID][s.ToID] = s.Amount
		}
	}

	for fromId, m := range maps {
		for toId, v := range m {
			if v > 0 {
				rs = append(rs, Settlement{fromId, toId, v})
			} else if v < 0 {
				rs = append(rs, Settlement{toId, fromId, v * -1})
			}
		}
	}

	return
}
