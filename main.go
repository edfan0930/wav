package main

import (
	"fmt"

	"github.com/edfan0930/wav/playwav"
)

func main() {

	fmt.Println("init buffers", playwav.InitMandarinBuffers([]string{"0.wav", "1.wav", "2.wav", "3.wav", "4.wav", "5.wav", "6.wav", "7.wav", "8.wav", "9.wav", "號.wav", "來賓.wav", "請到.wav", "櫃台.wav"}))
	fmt.Println("init buffers", playwav.InitEnglishBuffers([]string{"0.wav", "1.wav", "2.wav", "3.wav", "4.wav", "5.wav", "6.wav", "7.wav", "8.wav", "9.wav", "號.wav", "來賓.wav", "請到.wav", "櫃台.wav"}))
	fmt.Println("init buffers", playwav.InitHakkaBuffers([]string{"0.wav", "1.wav", "2.wav", "3.wav", "4.wav", "5.wav", "6.wav", "7.wav", "8.wav", "9.wav", "號.wav", "來賓.wav", "請到.wav", "櫃台.wav", "拾.wav", "佰.wav", "仟.wav"}))
	fmt.Println("init buffers", playwav.InitHokkienBuffers([]string{"0.wav", "1.wav", "2.wav", "3.wav", "4.wav", "5.wav", "6.wav", "7.wav", "8.wav", "9.wav", "號.wav", "來賓.wav", "請到.wav", "櫃台.wav", "拾.wav", "佰.wav", "仟.wav"}))
	playwav.InitSpeaker()
	for {

		//		fmt.Println("hakka", playwav.PlayWithHakka(3021, 1))
		fmt.Println("hokkien", playwav.PlayWithHokkien(3001, 1))
	}

}
