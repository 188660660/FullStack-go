package service

import (
	"testing"
)

func Test_QueryUserInfo(t *testing.T) {
	Convey("query user or create user", t, func() {
		_, err := s.QueryUser("zhanglu")
		So(err, ShouldBeNil)
	})
}
