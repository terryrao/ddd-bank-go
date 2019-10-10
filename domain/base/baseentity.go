package base

import (
	"fmt"
	"reflect"
)

type Entity struct {
	Id int64
}

func (b *Entity) equals(e *Entity) bool {
	bType := reflect.TypeOf(b)
	fmt.Printf("%v", bType)
	eType := reflect.TypeOf(b)
	if bType != eType {
		return false
	}
	return true

}
