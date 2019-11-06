package account

import (
	"fmt"
	"github.com/Zilliqa/gozilliqa-sdk/util"
	"testing"
)

var f = "{\"address\":\"9bfec715a6bd658fcb62b0f8cc9bfa2ade71434a\",\"id\":\"1497eb45-3a52-4c5a-97eb-88d5e790fcd0\",\"version\":3,\"crypto\":{\"cipher\":\"aes-128-ctr\",\"ciphertext\":\"3ddd39cb13c95ccdc150c962fadaebfa7a2fca3221c81e276491d70a5d621dd5\",\"kdf\":\"pbkdf2\",\"mac\":\"980f95923582693dad2038ea4e1119a934332c53d620ebe38b7e3b7928e57d05\",\"cipherparams\":{\"iv\":\"39a7beef25795f912572718363dba9f4\"},\"kdfparams\":{\"n\":8192,\"c\":262144,\"r\":8,\"p\":1,\"dklen\":32,\"salt\":\"4f3ddae640ebe3cb45a133c583d03e5da25c36baf4472343fb5f6a0c899b78f1\"}}}"

func TestToFile(t *testing.T) {
	file, err := ToFile("e19d05c5452598e24caad4a0d85a49146f7be089515c905ae6a19e8a578a6930", "xiaohuo", 0)
	if err != nil {
		t.Error(err.Error())
	}
	if f != file {
		t.Failed()
	}
}

func TestFromFile(t *testing.T) {
	a, err := FromFile(f, "xiaohuo")
	if err != nil {
		t.Error(err.Error())
	}

	fmt.Println(util.EncodeHex(a.PrivateKey))
	if util.EncodeHex(a.PrivateKey) != "e19d05c5452598e24caad4a0d85a49146f7be089515c905ae6a19e8a578a6930" {
		t.Failed()
	}
}

func TestNewHDAccount(t *testing.T) {
	expectAddresses := []string{
		"bea456fb58094be1c7f99bb6d1584dcec642b0b0",
		"aacdf9c84bba51878c8681c72f035b62135d6d7e",
		"852f52532c3c928269bdd3b83ac88e25a04d6b3b",
		"0237f40d30d3c37c9b77577acbb11c972cc58664",
		"cd6cb5bc8f3ee8ff7a91b060ce341feb6fc40e21",
		"ecd9d875c7366432a7ce403a7702dfa3e7f09602",
		"9165ae9ceeb155fb75d9c1fee2041f12c6e1f5ea",
		"0723dd96243491ee84a925edb657f24582aec899",
		"68275607e8bdf7cfa248b5f5a07b576f9ef39cd1",
		"4878d8eb9a63493a6de066eb1458cab672dc8cfd",
	}

	var mnemonic = "cart hat drip lava jelly keep device journey bean mango rocket festival"
	for i := 0; i < 10; i++ {
		account, err := NewHDAccount(mnemonic, uint32(i))
		if err != nil {
			t.Failed()
		} else {
			if !contains(expectAddresses, account.Address) {
				t.Failed()
			}
		}
	}
}

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}
