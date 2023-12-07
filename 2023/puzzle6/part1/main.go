package main

func main() {
	times := []int{56, 97, 77, 93}
	distances := []int{499, 2210, 1097, 1440}

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
				println("Race ", j+1, "MS pressed:", i, "Distance traveled:", distance)
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
