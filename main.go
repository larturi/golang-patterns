package main

import (
	"mod_patterns/adapter"
	"mod_patterns/factory"
	"mod_patterns/observer"
	"mod_patterns/singleton"
	"mod_patterns/strategy"
)

func main() {
	factory.MainFactory(false)

	adapter.MainAdapter(false)

	singleton.MainSingleton(false)

	observer.MainObserver(false)

	strategy.MainStrategy(true)
}
