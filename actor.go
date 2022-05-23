package badpack

import "github.com/larsklemstein/badpack/internal/actor"

func Say(msg string) {
	actor.PrintToConsole("==>" + msg)
}
