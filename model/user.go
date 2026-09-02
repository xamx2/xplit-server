package model

import (
	"context"

	"github.com/uptrace/bun"
)

type User struct {
	ID int32 `bun:"id,pk,autoincrement"`
}

func (u User) CreateGroup(db bun.IDB, ctx context.Context, g *Group) error {
	return db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		if _, err := tx.NewInsert().Model(g).Exec(ctx); err != nil {
			return err
		}
		if _, err := tx.NewInsert().Model(&GroupMember{
			UserID:  &u.ID,
			GroupID: g.ID,
			Role:    GroupMemberRoleOwner,
		}).Exec(ctx); err != nil {
			return err
		}
		return nil
	})
}

func (u User) Groups(db bun.IDB, ctx context.Context) (gs []Group, err error) {
	err = db.NewSelect().Model(&gs).
		Join("JOIN group_members AS gm ON gm.group_id = g.id").
		Where("gm.user_id = ?", u.ID).
		Scan(ctx)
	return
}

func (u User) Group(db bun.IDB, ctx context.Context, id any) (g *Group, err error) {
	g = new(Group)
	err = db.NewSelect().Model(g).
		Join("JOIN group_members AS gm ON gm.group_id = g.id").
		Where("gm.group_id = ? AND gm.user_id = ?", id, u.ID).
		Scan(ctx)
	return
}
