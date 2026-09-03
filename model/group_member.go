package model

import (
	"context"
	"errors"
	"time"

	"github.com/uptrace/bun"
)

var (
	ErrNotOwner = errors.New("not owner")
)

type GroupMemberRole string

const (
	GroupMemberRoleMember GroupMemberRole = "member"
	GroupMemberRoleOwner  GroupMemberRole = "owner"
)

type GroupMember struct {
	bun.BaseModel `bun:"alias:gm"`

	ID        int32           `bun:"id,pk,autoincrement"`
	GroupID   int32           `bun:"group_id,notnull,unique:group_user"`
	UserID    *int32          `bun:"user_id,unique:group_user"`
	Name      *string         `bun:"name"`
	Role      GroupMemberRole `bun:"role,nullzero,notnull,default:'member'"`
	CreatedAt time.Time       `bun:"created_at,nullzero,notnull,default:current_timestamp"`
}

func (gm GroupMember) CreateGroupMember(db bun.IDB, ctx context.Context, m *GroupMember) error {
	if gm.Role != GroupMemberRoleOwner {
		return ErrNotOwner
	}
	m.GroupID = gm.GroupID
	_, err := db.NewInsert().Model(m).Exec(ctx)
	return err
}

func (gm GroupMember) CreateTransaction(db bun.IDB, ctx context.Context, t *Transaction) error {
	return db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		if _, err := tx.NewInsert().Model(t).Exec(ctx); err != nil {
			return err
		}
		for _, s := range t.NewSplits {
			s.TransactionID = t.ID
		}
		if _, err := tx.NewInsert().Model(&t.NewSplits).Exec(ctx); err != nil {
			return err
		}
		return nil
	})
}
