package service

import (
	"context"
	"http_demo/dao"
)

type UserDTO struct {
	ID       uint64
	Nickname string
}

type UserService struct {
}

func (s *UserService) GetUserByID(ctx context.Context, id uint64) (*UserDTO, error) {
	user, err := dao.FindUserByID(ctx, id)
	if err != nil {
		// 内部中间调用，err直接传递
		// 如果针对 err not found 有特殊逻辑处理
		// 则使用 errors.Is(err, dao.ErrNotFound) 判断
		// 处理了这个错误后，就不应该在往上层传递这个错误了
		return nil, err
	}
	ret := &UserDTO{
		ID:       user.ID,
		Nickname: user.Nickname,
	}
	return ret, nil
}
