package main

func main() {
	time := 56977793
	distances := 499221010971440

	waysToWin := 0
	for i := 0; i < time; i++ {
		if i == 0 {
			continue
		}

		distance := (time - i) * i

		if distance > distances {
			waysToWin++
		}
	}

	println(waysToWin)
}
