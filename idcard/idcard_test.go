package idcard

import (
	"fmt"
	"testing"
)

func TestIsValidCitizenNo(t *testing.T) {
	id := []byte("your id card")
	fmt.Println(IsValidCitizenNo(&id))
}
