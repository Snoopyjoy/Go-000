package dao

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

type User struct {
	ID       uint64 `json:"id" db:"id"`
	Nickname string `json:"nickname" db:"nickname"`
}

func FindUserByID(ctx context.Context, ID uint64) (*User, error) {
	u, err := mockGetUser(ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// sql.ErrNoRows改为dao内部的sentinel err, 屏蔽底层数据库实现
			return nil, fmt.Errorf("[%d] %w", ID, ErrNotFound)
		}
		return nil, errors.Wrapf(err, "find user [%d] err", ID)
	}
	return u, nil
}

// mockGetUser 虚假的db
// 假装这是一个user库，从库里根据id查找user
// uid 为 1001 时有值，否则返回 sql.ErrNoRows
func mockGetUser(uid uint64) (*User, error) {
	if uid == 1001 {
		return &User{
			ID:       1001,
			Nickname: "周星星",
		}, nil
	}

	if uid == 1003 {
		return nil, errors.New("test internal error")
	}

	return nil, sql.ErrNoRows
}
