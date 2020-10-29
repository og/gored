package red

import (
	"strings"
)

type Key struct {
	value string
}
func (k Key) String() string {
	return k.value
}
func NewKey(k ...string) Key {
	return Key{value: strings.Join(k, ":")}
}


func keysToStrings(keys []Key) (stringKeys []string) {
	for _, key := range keys {
		stringKeys = append(stringKeys, key.String())
	}
	return
}

