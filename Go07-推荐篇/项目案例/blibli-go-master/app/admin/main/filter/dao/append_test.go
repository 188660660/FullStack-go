package dao

import (
	"fmt"
	"testing"
)

func TestDao_AiScore(t *testing.T) {
	Convey("TestDao_AiScore", t, func() {
		score, err := dao.AiScore(ctx, "string")
		fmt.Printf("score:%+v \n", score)
		fmt.Printf("err:%+v \n", err)
		So(score, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})
}
