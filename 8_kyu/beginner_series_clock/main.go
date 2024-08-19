package main

func main() {

}

func Past(h, m, s int) int {
	second := 1000
	minute := 60 * second
	hour := 60 * minute
	return h * hour + m * minute + s * second
}

