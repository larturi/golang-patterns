package main

import (
	"mod_patterns/adapter"
	"mod_patterns/factory"
	"mod_patterns/singleton"
)

func main() {
	factory.MainFactory()

	adapter.MainAdapter()

	singleton.MainSingleton()
}
