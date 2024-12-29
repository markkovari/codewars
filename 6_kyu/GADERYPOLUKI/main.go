package main

import (
	"bytes"
	"strings"
)

func makeSwitchMap(keys string) map[string]string {
	asLowerCase := strings.ToLower(keys)
	asUpperCase := strings.ToUpper(keys)
	switchMap := make(map[string]string, 0)
	for i := 0; i < len(keys); i++ {
		if i%2 == 0 {
			switchMap[string(asUpperCase[i])] = string(asUpperCase[i+1])
			switchMap[string(asLowerCase[i])] = string(asLowerCase[i+1])
			switchMap[string(asUpperCase[i+1])] = string(asUpperCase[i])
			switchMap[string(asLowerCase[i+1])] = string(asLowerCase[i])
		} else {
			switchMap[string(asUpperCase[i-1])] = string(asUpperCase[i])
			switchMap[string(asLowerCase[i-1])] = string(asLowerCase[i])
			switchMap[string(asUpperCase[i])] = string(asUpperCase[i-1])
			switchMap[string(asLowerCase[i])] = string(asLowerCase[i-1])
		}
	}
	return switchMap
}

func replace(of, cypher string) string {
	switchMap := makeSwitchMap(cypher)
	var buf bytes.Buffer
	for _, letter := range strings.Split(of, "") {
		if val, ok := switchMap[letter]; ok {
			buf.WriteString(val)
		} else {
			buf.WriteString(letter)
		}
	}
	return buf.String()
}

func Encode(str, key string) string {
	return replace(str, key)
}

func Decode(str, key string) string {
	return replace(str, key)
}
