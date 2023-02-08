package client

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	sum, err := md5Checksum("../bin/launcheq.exe")
	if err != nil {
		t.Fatalf("checksum: %s", err)
	}
	fmt.Println(sum)
}
