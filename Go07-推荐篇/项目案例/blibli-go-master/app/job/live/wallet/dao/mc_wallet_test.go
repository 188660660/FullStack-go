package dao

import (
	"testing"
)

func TestDelWalletCache(t *testing.T) {
	Convey("Test Del Memcache", t, func() {
		once.Do(startService)
		err := d.DelWalletCache(ctx, 10000)
		So(err, ShouldBeNil)
	})
}
