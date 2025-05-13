package main

import "github.com/zyz880615/HeadFirstByGo/ch11/gadget"

type Player interface {
	Play(string)
	Stop()
}

func TryOut(player Player) {
	player.Play("Test Track")
	player.Stop()
	recorder, ok := player.(gadget.TapRecorder)
	if ok {
		recorder.Record()
	}
}

func main() {
	TryOut(gadget.TapPlayer{})
	TryOut(gadget.TapRecorder{})
}
