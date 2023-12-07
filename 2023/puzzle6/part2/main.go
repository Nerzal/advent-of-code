package main

func main() {
	times := []int{56977793}
	distances := []int{499221010971440}

	result := 0
	for j, time := range times {
		waysToWin := 0
		for i := 0; i < time; i++ {
			if i == 0 {
				continue
			}

			distance := (time - i) * i

			if distance > distances[j] {
				waysToWin++
			}
		}

		if result == 0 {
			result = waysToWin
		} else {
			result *= waysToWin
		}
	}

	println(result)
}
