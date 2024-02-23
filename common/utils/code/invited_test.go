package invite_code

import (
	"fmt"
	"testing"
)

func TestDecode(t *testing.T) {
	var arr = []string{}
	for i := 0; i < 1000; i++ {
		code := Encode(uint64(i))
		// fmt.Printf("user: %d, code: %s\n", i, code)
		arr = append(arr, code)
		userId := Decode(code)

		fmt.Printf("oldId: %d,userId: %d, code: %s\n", i, userId, code)

	}
	fmt.Println(len(arr))
	fmt.Println(len(removeDuplicates(arr)))
}

func removeDuplicates(arr []string) []string {
	uniqueMap := make(map[string]bool)
	uniqueSlice := []string{}

	for _, str := range arr {
		if !uniqueMap[str] {
			uniqueMap[str] = true
			uniqueSlice = append(uniqueSlice, str)
		}
	}
	return uniqueSlice
}

func TestDecod3e(t *testing.T) {

	var result []string
	for i := 0; i < 5; i++ {
		encode := Encode(uint64(i))
		result = append(result, encode)
	}

	fmt.Println(len(result))
	fmt.Println(result)
}
