package main

import (
	"fmt"

	"github.com/edfan0930/wav/playwav"
)

func main() {

	fmt.Println("init buffers", playwav.InitMandarinBuffers([]string{"0.wav", "1.wav", "2.wav", "3.wav", "4.wav", "5.wav", "6.wav", "7.wav", "8.wav", "9.wav", "號.wav"}))
	playwav.InitSpeaker()

	for {

		fmt.Println("play", playwav.PlayWithMandarinSingle(3001))
	}

}
