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
