package crypto

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
	"unsafe"

	"github.com/cbergoon/merkletree"
	"github.com/ethereum/go-ethereum/common"
)

//
//func Test_buildMerkleTree(t *testing.T) {
//	empty := trie.NewEmpty(trie.NewDatabase(rawdb.NewMemoryDatabase()))
//	var strings []string
//	for i := 0; i < 1000000; i++ {
//		// 生成随机字节数组
//		var buf [20]byte
//		_, err := rand.Read(buf[:])
//		if err != nil {
//			panic(err)
//		}
//
//		// 将字节数组转换为 0x 地址
//		addr := common.BytesToAddress(buf[:])
//		strings = append(strings, addr.Hex())
//		empty.Update(addr.Bytes(), addr.Bytes())
//	}
//
//	empty.Hash()
//	size := unsafe.Sizeof(*empty)
//	fmt.Printf("Person object size: %d bytes\n", size)
//	return
//}

type AddressContent struct {
	address common.Address
}

func NewAddressContent(address string) AddressContent {
	return AddressContent{
		address: common.HexToAddress(address),
	}
}

func (a AddressContent) CalculateHash() ([]byte, error) {
	hash := a.address.Hash()
	return hash[:], nil
}

func (a AddressContent) Equals(other merkletree.Content) (bool, error) {
	return bytes.Equal(a.address.Bytes(), other.(AddressContent).address.Bytes()), nil
}

func Test_buildMerkleTree(t *testing.T) {
	var addrs = []string{"0x817016163775AaF0B25DF274fB4b18edB67E1F26", "0x19b2A18D2850acCdf90418Bd68DAC8b07123350B"}

	content := make([]merkletree.Content, 0, len(addrs))
	for i := 0; i < len(addrs); i++ {
		content = append(content, AddressContent{
			address: common.HexToAddress(addrs[i]),
		})
	}
	tree, err := merkletree.NewTree(content)
	if err != nil {
		fmt.Print(err)
	}
	//path, _, _ := tree.GetMerklePath(NewAddressContent("0x817016163775AaF0B25DF274fB4b18edB67E1F26"))
	verifyContent, err := tree.VerifyContent(NewAddressContent("0x19b2A18D2850acCdf90418Bd68DAC8b07123350B"))
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(verifyContent)
}

func sizeOfStruct(v interface{}) uintptr {
	var size uintptr

	val := reflect.ValueOf(v)
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldSize := unsafe.Sizeof(field.Interface())
		size += fieldSize
	}

	return size
}
