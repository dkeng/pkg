package time_test

import (
	"fmt"
	"testing"

	"github.com/dkeng/pkg/time"
)

func TestTodayStartEndTime(t *testing.T) {
	t1, t2 := time.TodayStartEndTime()
	fmt.Println(t1, t2)
}
