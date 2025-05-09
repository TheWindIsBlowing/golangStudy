package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

func printAll(vals []interface{}) {
	for _, val := range vals {
		fmt.Println(val)
	}
}

func main() {
	// 报错
	// str := []string{"stanley", "david", "oscar"}
	// printAll(str)

	func1()
}

func func1() {
	input := `
	{
		"created_at": "Thu May 31 00:00:01 +0000 2012"
	}
	`
	var val map[string]Timestamp
	if err := json.Unmarshal([]byte(input), &val); err != nil {
		fmt.Println("err: ", err)
	}

	fmt.Println(val)
	for k, v := range val {
		fmt.Println(k, reflect.TypeOf(v))
	}
	fmt.Println(time.Time(val["created_at"]))
}

type Unmarshaler interface {
	UnmarshalJSON([]byte) error
}

type Timestamp time.Time

func (t *Timestamp) UnmarshalJSON(b []byte) error {
	v, err := time.Parse(time.RubyDate, string(b[1:len(b)-1]))
	if err != nil {
		return err
	}

	*t = Timestamp(v)
	return nil
}
