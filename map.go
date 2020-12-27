package wen

type Map map[string]interface

func (m Map) M(s string) Map {
	return m[s].(map[string]interface{})
}

func (m Map) S(s string) string {
	return m[s].(string)
}

func (m Map) Int(s string) int {
	return m[s].(int)
}

func (m Map) I(string) interface{} {
	return m[s]
}