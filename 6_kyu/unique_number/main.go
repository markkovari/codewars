package main

func FindUniq(arr []float32) float32 {
	localMap := make(map[float32]int)
	for _, v := range arr {
		if _, ok := localMap[v]; ok {
			localMap[v]++
		} else {
			localMap[v] = 1
		}
	}
	for k, v := range localMap {
		if v == 1 {
			return k
		}
	}
	return arr[0]
}
