package baseCase

import (
	"fmt"
	"reflect"
	"testing"
)

type Video struct {
	Name string `json:"name" doc:"你的名字"`
	Sex  string `json:"sex" doc:"哈哈"`
}

func findDoc(str any) map[any]any {
	t := reflect.TypeOf(str).Elem()
	doc := map[any]any{}

	for i := 0; i < t.NumField(); i++ {
		doc[t.Field(i).Tag.Get("json")] = t.Field(i).Tag.Get("doc")
	}
	return doc
}

func TestReflect(t *testing.T) {
	var str Video
	doc := findDoc(&str)
	fmt.Println(doc)
}
