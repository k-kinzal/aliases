package config_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/k-kinzal/aliases/pkg/aliases/config"
)

func ExampleUnmarshal() {
	content := `
/path/to/command1:
  image: alpine
  tag: latest
  env:
    ONE: 1
/path/to/command2:
  image: alpine
  tag: latest
  env:
    TWO: 2
  dependencies:
  - /path/to/command1
/path/to/command3:
  image: alpine
  tag: latest
  env:
    THREE: 3
  dependencies:
  - /path/to/command2
  - /path/to/command4:
      image: alpine
      tag: latest
      env:
        ONE: 4
`
	conf, err := config.Unmarshal([]byte(content))
	if err != nil {
		panic(err)
	}
	{
		opt, err := conf.Get("/path/to/command1")
		if err != nil {
			panic(err)
		}
		out, _ := json.Marshal(opt.Env)
		fmt.Println("/path/to/command1", string(out))
	}
	{
		opt, err := conf.Get("/path/to/command2")
		if err != nil {
			panic(err)
		}
		out, _ := json.Marshal(opt.Env)
		fmt.Println("/path/to/command2", string(out))
	}
	{
		opt, err := conf.Get("/path/to/command3")
		if err != nil {
			panic(err)
		}
		out, _ := json.Marshal(opt.Env)
		fmt.Println("/path/to/command3", string(out))
	}
	// Output:
	// /path/to/command1 {"ONE":"1"}
	// /path/to/command2 {"ONE":"1","TWO":"2"}
	// /path/to/command3 {"ONE":"4","THREE":"3","TWO":"2"}
}

func ExampleLoadConfig() {
	content := `
/path/to/command1:
  image: alpine
  tag: latest
  name: alpine1
`
	file, err := ioutil.TempFile("/tmp", "")
	if err != nil {
		panic(err)
	}
	if _, err := file.Write([]byte(content)); err != nil {
		panic(err)
	}

	conf, err := config.LoadConfig(file.Name())
	if err != nil {
		panic(err)
	}

	opt, err := conf.Get("/path/to/command1")
	if err != nil {
		panic(err)
	}
	fmt.Println(*opt.Name)
	// Output: alpine1
}

func ExampleConfig_Get() {
	content := `
/path/to/command1:
  image: alpine
  tag: latest
  name: alpine1
`
	conf, err := config.Unmarshal([]byte(content))
	if err != nil {
		panic(err)
	}
	opt, err := conf.Get("/path/to/command1")
	if err != nil {
		panic(err)
	}
	fmt.Println(*opt.Name)
	// Output: alpine1
}

func ExampleConfig_Slice() {
	content := `
/path/to/command1:
  image: alpine
  tag: latest
  name: alpine1
/path/to/command2:
  image: alpine
  tag: latest
  name: alpine2
`
	conf, err := config.Unmarshal([]byte(content))
	if err != nil {
		panic(err)
	}
	slice := conf.Slice()
	fmt.Println(len(slice))
	// Output: 2
}

func ExampleConfig_Binaries() {
	content := `
/path/to/command1:
  docker:
    image: docker
    tag: 18.09.1
  image: alpine
  tag: latest
  name: alpine1
  dependencies:
  - /path/to/command2:
      docker:
        image: docker
        tag: 18.06.1
      image: alpine
      tag: latest
      name: alpine2
  - /path/to/command3:
      docker:
        image: docker
        tag: 18.09.0
      image: alpine
      tag: latest
      name: alpine3
      dependencies:
      - /path/to/command4:
          docker:
            image: docker
            tag: 18.06.1
          image: alpine
          tag: latest
          name: alpine4
`
	conf, err := config.Unmarshal([]byte(content))
	if err != nil {
		panic(err)
	}
	for _, binary := range conf.Binaries("/tmp") {
		out, err := json.Marshal(binary)
		if err != nil {
			panic(nil)
		}
		fmt.Println(string(out))
	}
	// Output:
	// {"Image":"docker","Tag":"18.06.1","Path":"/tmp/docker-18-06-1"}
	// {"Image":"docker","Tag":"18.09.1","Path":"/tmp/docker-18-09-1"}
	// {"Image":"docker","Tag":"18.09.0","Path":"/tmp/docker-18-09-0"}
}
