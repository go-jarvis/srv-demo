package blog

import (
	"fmt"
	"testing"
	"time"
)

func Test_Timestamp(t *testing.T) {
	ts := timestamp()
	fmt.Println(ts)

	now := time.Now()

	fmt.Println(now.Unix(), "<-now = utc-now ->", now.UTC().Unix())
	fmt.Println(now, "<-now = utc-now ->", now.UTC())

}
