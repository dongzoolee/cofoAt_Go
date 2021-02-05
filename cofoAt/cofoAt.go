package main

import (
	"cf/getData"
	"cf/updateUser"
	"fmt"
)

func main() {
	getData.GetContest()
	fmt.Println(getData.GetTier("djs100201"))
	// functional.AddUser("nant0313")
	getData.GetUserCng("djs100201")
	getData.GetContCng("698")
	getData.GetContCng("1919")
	if !functional.UserExists("nant0313") {
		functional.AddUser("nant0313", "init")
	} else {
		var channelIdx int = functional.ChannelExists("nant0313", "do")
		if channelIdx != -1 {
			functional.DelChannel("nant0313", channelIdx)
		} else {
			functional.AddChannel("nant0313", "do")
		}
	}
}
