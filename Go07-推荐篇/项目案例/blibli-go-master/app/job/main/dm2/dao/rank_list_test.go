package dao

import (
	"context"
	"testing"
)

func TestRankList(t *testing.T) {
	Convey("", t, func() {
		res, err := testDao.RankList(context.Background(), 185, 3)
		So(err, ShouldBeNil)
		So(res, ShouldNotBeNil)
		t.Log(res)
		for _, v := range res.List {
			t.Log(v)
		}
	})
}
