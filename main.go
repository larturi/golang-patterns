package main

import (
	"mod_patterns/adapter"
	"mod_patterns/factory"
	"mod_patterns/observer"
	"mod_patterns/singleton"
)

func main() {
	factory.MainFactory(false)

	adapter.MainAdapter(false)

	singleton.MainSingleton(false)

	observer.MainObserver(true)
}
