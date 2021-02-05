package wen

import "strings"

type Map map[string]interface{}

func (m Map) M(s string) Map {
	return m[s].(map[string]interface{})
}

func (m Map) S(s string) string {
	return m[s].(string)
}

func (m Map) Int(s string) int {
	return m[s].(int)
}
func (m Map) Bool(s string) bool {
	return m[s].(bool)
}
func (m Map) I(s string) interface{} {
	return m[s]
}
func (m Map) Have(s string) bool {
	_, ok := m[s]
	return ok
}

func (m Map) Get(s string) interface{} {
	splits := strings.Split(s, ".")
	if len(splits) == 1 {
		return m.I(s)
	}
	var r = m
	for index, split := range splits {
		if index == len(splits)-1 {
			return r.I(split)
		} else {
			r = r.M(split)
		}
	}
	return nil
}
