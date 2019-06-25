package dao

import (
	"context"
	"testing"
)

func TestDMAdvert(t *testing.T) {
	Convey("dm advert", t, func() {
		var (
			c             = context.TODO()
			aid     int64 = 10100572
			cid     int64 = 10115256
			mid     int64 = 12345881
			build   int64 = 8111
			adExtra       = ""
			buvid         = "5400000"
			mobiApp       = "android"
		)
		testDao.DMAdvert(c, aid, cid, mid, build, buvid, mobiApp, adExtra)
	})
}
