package archive

import (
	"context"
	"testing"
)

func Test_AuditTypesConf(t *testing.T) {
	Convey("AuditTypesConf", t, func() {
		configs, err := d.AuditTypesConf(context.TODO())
		So(err, ShouldBeNil)
		Println(configs)
	})
}
