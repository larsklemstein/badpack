package badpack

import "github.com/larsklemstein/badpack/internal/actor"

func Say(msg string) {
	actor.PrintToConsole("v2 v2 v2==>" + msg)
}
