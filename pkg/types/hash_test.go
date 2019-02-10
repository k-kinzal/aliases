package types_test

import (
	"fmt"

	"github.com/k-kinzal/aliases/pkg/types"
)

func ExampleMD5() {
	var hasher types.Hasher = types.MD5
	fmt.Println(hasher("abcde"))
	fmt.Println(hasher(struct{}{}))
	// Output:
	// ab56b4d92b40713acc5af89985d4b786
	// 99914b932bd37a50b983c5e7c90ae93b
}

func ExampleSHA1() {
	var hasher types.Hasher = types.SHA1
	fmt.Println(hasher("abcde"))
	fmt.Println(hasher(struct{}{}))
	// Output:
	// 03de6c570bfe24bfc328ccd7ca46b76eadaf4334
	// bf21a9e8fbc5a3846fb05b4fa0859e0917b2202f
}

func ExampleSHA256() {
	var hasher types.Hasher = types.SHA256
	fmt.Println(hasher("abcde"))
	fmt.Println(hasher(struct{}{}))
	// Output:
	// 36bbe50ed96841d10443bcb670d6554f0a34b761be67ec9c4a8ad2c0c44ca42c
	// 44136fa355b3678a1146ad16f7e8649e94fb4fc21fe77e8310c060f61caaff8a
}
