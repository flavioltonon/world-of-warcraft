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
	return newNamespace("static", region.String())
}

func NewDynamicNamespace(region Region) Namespace {
	return newNamespace("dynamic", region.String())
}

func newNamespace(prefix string, region string) Namespace {
	return Namespace(fmt.Sprintf("%s_%s", prefix, strings.ToLower(region)))
}
