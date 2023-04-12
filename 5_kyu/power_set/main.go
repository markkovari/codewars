package main

var changeSet map[string]string = map[string]string{}

const (
	from string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	to   string = "NOPQRSTUVWXYZABCDEFGHIJKLMnopqrstuvwxyzabcdefghijklm"
)

func init() {
	for i := 0; i < len(from); i++ {
		changeSet[string(from[i])] = string(to[i])
	}
}

func Rot13(msg string) string {
	result := ""
	for _, c := range msg {
		val, exists := changeSet[string(c)]
		if exists {
			result += val
		} else {
			result += string(c)
		}
	}
	return result
}

func main() {
	println("Hello, World!")
}
