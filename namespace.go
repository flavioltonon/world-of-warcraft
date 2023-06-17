package worldofwarcraft

import "fmt"

type Namespace string

func (n Namespace) String() string {
	return string(n)
}

func NewStaticNamespace(region Region) Namespace {
	return Namespace(fmt.Sprintf("static_%s", region))
}

func NewDynamicNamespace(region Region) Namespace {
	return Namespace(fmt.Sprintf("dynamic_%s", region))
}
