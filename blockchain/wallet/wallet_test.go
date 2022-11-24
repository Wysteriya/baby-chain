package wallet

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	x, y, _ := GeneratePublicAddressAndKey()
	fmt.Println(len(x))
	fmt.Println(len(y))
}
