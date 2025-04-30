package main

import (
	goamp "github.com/alterminal/go-amp"
)

func main() {
	str := "hello"
	sbuf := goamp.Marshal(str)
	goamp.Unmarshal(sbuf)
	s := make([]string, 0)
	s = append(s, "wefwe")
	sbuf = goamp.Marshal(s)
	goamp.Unmarshal(sbuf)
	m := make(map[string]string, 0)
	m["hello"] = "world"
	m["world"] = "false"
	sbuf = goamp.Marshal(m)
	// um, _ := goamp.Unmarshal(sbuf).(map[any]any)
	// v := um["world"]
	// fmt.Println(v)
}
