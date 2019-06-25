package dao

import (
	"context"
	"flag"
	"path/filepath"

	"go-common/app/admin/ep/melloi/conf"

	_ "github.com/go-sql-driver/mysql"
)

var (
	d *Dao
	c context.Context
)

func init() {
	dir, _ := filepath.Abs("../cmd/convey-test.toml")
	flag.Set("conf", dir)
	conf.Init()
	d = New(conf.Conf)
	d.httpClient.SetTransport(gock.DefaultTransport)
	c = ctx()
}

func ctx() context.Context {
	return context.Background()
}
