package types_test

import (
	"fmt"

	"github.com/k-kinzal/aliases/pkg/types"
)

func ExampleMD5() {
	var hasher types.Hasher = types.MD5
	fmt.Println(hasher(struct{}{}))
	// Output: 99914b932bd37a50b983c5e7c90ae93b
}

func ExampleSHA1() {
	var hasher types.Hasher = types.SHA1
	fmt.Println(hasher(struct{}{}))
	// Output: bf21a9e8fbc5a3846fb05b4fa0859e0917b2202f
}

func ExampleSHA256() {
	var hasher types.Hasher = types.SHA256
	fmt.Println(hasher(struct{}{}))
	// Output: 44136fa355b3678a1146ad16f7e8649e94fb4fc21fe77e8310c060f61caaff8a
}
