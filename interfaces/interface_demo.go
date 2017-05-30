package interfaces

type Speaker interface {
	Speak() string
}

type HardCodedSpeaker struct {
}

func (p HardCodedSpeaker) Speak() string {
	return "I speak!"
}

type ConfigurableSpeaker struct {
	Speech string
}

func (speaker ConfigurableSpeaker) Speak() string {
	return speaker.Speech
}
