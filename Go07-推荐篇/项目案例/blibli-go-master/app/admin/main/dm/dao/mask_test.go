package dao

import (
	"context"
	"testing"
)

func TestGenerateMask(t *testing.T) {
	Convey("Test generate mask", t, func() {
		testDao.GenerateMask(context.TODO(), 32, 11, 1, 0, 0, 0, 0)
	})
}
