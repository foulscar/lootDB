package main

func main() {
	err := initGame()
	if err != nil {
		panic(err)
	}
	initCMDHandlers()

	play()
}
