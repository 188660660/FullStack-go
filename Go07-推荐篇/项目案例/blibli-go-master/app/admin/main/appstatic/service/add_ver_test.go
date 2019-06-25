package service

import (
	"testing"
)

func TestService_DefaultPkg(t *testing.T) {
	Convey("AddFile should return without err", t, WithService(func(svf *Service) {
		err := svf.DefaultPkg(74, 3, 54)
		So(err, ShouldBeNil)
	}))
}
