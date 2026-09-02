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
