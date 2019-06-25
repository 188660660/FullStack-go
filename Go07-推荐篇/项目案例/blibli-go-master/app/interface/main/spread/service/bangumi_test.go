package service

import (
	"testing"
)

func Test_Bangumi(t *testing.T) {
	Convey("bangumi", t, func() {
		res, err := s.BangumiContent(c, 1, 10, 1, "douban")
		So(err, ShouldBeNil)
		So(res, ShouldNotBeEmpty)
	})
}

func Test_BangumiOff(t *testing.T) {
	Convey("bangumi", t, func() {
		res, err := s.BangumiOff(c, 1, 10, 1, 1, "douban")
		So(err, ShouldBeNil)
		So(res, ShouldNotBeEmpty)
	})
}
