package dao

import (
	"context"
	"testing"
)

func TestDao_PayCenterWallet(t *testing.T) {
	Convey("normal", t, func() {
		mid := int64(1)
		w, e := testDao.PayCenterWallet(context.Background(), mid, "android")
		So(e, ShouldBeNil)
		t.Logf("%+v", w)

	})
}
