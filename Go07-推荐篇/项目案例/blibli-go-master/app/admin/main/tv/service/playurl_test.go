package service

import (
	"fmt"
	"testing"
)

func TestService_Playurl(t *testing.T) {
	Convey("PlayURL test", t, WithService(func(s *Service) {
		url, err := s.Playurl(2476470)
		fmt.Println(url)
		So(err, ShouldBeNil)
	}))
}
