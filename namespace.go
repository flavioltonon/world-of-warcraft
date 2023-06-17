package worldofwarcraft

import (
	"fmt"
	"strings"
)

type Namespace string

func (n Namespace) String() string {
	return string(n)
}

func NewStaticNamespace(region Region) Namespace {
	return newNamespace("static", region)
}

func NewDynamicNamespace(region Region) Namespace {
	return newNamespace("dynamic", region)
}

func newNamespace(prefix string, region Region) Namespace {
	return Namespace(fmt.Sprintf("%s-%s", prefix, strings.ToLower(region.String())))
}
