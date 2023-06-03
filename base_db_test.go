package postgresql

import (
	"errors"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestRetry(t *testing.T) {
	Convey("重试可以成功", t, func() {
		var err error
		Retry("retry", func() error {
			err = errors.New("retry error")
			return err
		})
		So(err, ShouldNotBeNil)

		start := time.Now()
		count := 0
		Retry("retry", func() error {
			count++
			if count <= 10 {
				err = errors.New("invalid connection")
				return err
			}
			err = nil
			return nil
		})
		sub := time.Since(start)
		So(err, ShouldBeNil)
		So(sub, ShouldBeGreaterThan, time.Millisecond)
	})
	Convey("只捕获固定的错误", t, func() {
		for _, msg := range errMsg {
			var err error
			start := time.Now()
			Retry(fmt.Sprintf("retry-%s", msg), func() error {
				err = fmt.Errorf("my-error: %s", msg)
				return err
			})
			So(err, ShouldNotBeNil)
			So(time.Since(start), ShouldBeGreaterThan, time.Millisecond)
		}
	})
}
