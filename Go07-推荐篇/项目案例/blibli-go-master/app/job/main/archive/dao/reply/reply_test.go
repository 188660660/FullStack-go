package reply

import (
	"testing"
)

func Test_ChangeSubjectMid(t *testing.T) {
	Convey("ChangeSubjectMid", t, func() {
		d.ChangeSubjectMid(0, 1684013)
	})
}
