package gadget

import "fmt"

type TapPlayer struct {
	Batteries string
}

func (t TapPlayer) Play(song string) {
	fmt.Println("Playing", song)
}

func (t TapPlayer) Stop() {
	fmt.Println("Stopped!")
}

type TapRecorder struct {
	Microphones string
}

func (t TapRecorder) Play(song string) {
	fmt.Println("Playing", song)
}

func (t TapRecorder) Record() {
	fmt.Println("Recording")
}

func (t TapRecorder) Stop() {
	fmt.Println("Stopped!")
}
