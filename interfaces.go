package main

import (
	"fmt"
	"go-demo/interfaces"
)

func main() {
	doSpeak(interfaces.HardCodedSpeaker{})
	doSpeak(interfaces.ConfigurableSpeaker{Speech: "Configured speech"})
}

func doSpeak(speaker interfaces.Speaker) {
	fmt.Println(speaker.Speak())
}
