package ctl

import (
	"context"
	"errors"
)

type key int

var userKey key

type UserInfo struct {
	ID uint `json:"id"`
}

func GetUserInfo(ctx context.Context) (*UserInfo, error) {
	user, ok := FromContext(ctx)
	if !ok {
		return nil, errors.New("fail to access user information")
	}
	return user, nil
}

func NewContext(ctx context.Context, user *UserInfo) context.Context {
	return context.WithValue(ctx, userKey, user)
}

func FromContext(ctx context.Context) (*UserInfo, bool) {
	user, ok := ctx.Value(userKey).(*UserInfo)
	return user, ok
}
