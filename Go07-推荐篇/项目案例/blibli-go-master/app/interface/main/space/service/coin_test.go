package service

import (
	"context"
	"testing"
)

func TestService_CoinVideo(t *testing.T) {
	Convey("test coin video", t, WithService(func(s *Service) {
		mid := int64(0)
		vmid := int64(88889018)
		data, err := s.CoinVideo(context.Background(), mid, vmid)
		So(err, ShouldBeNil)
		Printf("%+v", data)
	}))
}
