package service

import (
	"context"
	"testing"
)

func Test_QueryUserInfo(t *testing.T) {
	Convey("query user when user exists", t, func() {
		userInfo, err := s.QueryUserInfo(context.Background(), "fengyifeng")
		So(err, ShouldBeNil)
		So(userInfo.EMail, ShouldEqual, "fengyifeng@bilibili.com")
	})
}
