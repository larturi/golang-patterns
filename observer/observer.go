package observer

import "fmt"

type Topic interface {
	register(observer Observer)
	broadcast()
}

type Observer interface {
	getId() string
	updateValue(string)
}

type Item struct {
	observers []Observer
	name      string
	avaiable  bool
}

func NewItem(name string) *Item {
	return &Item{
		name: name,
	}
}

func (i *Item) UpdateAvaiable() {
	fmt.Printf("Item %s is avaiable\n", i.name)
	i.avaiable = true
	i.broadcast()
}

func (i *Item) broadcast() {
	for _, observer := range i.observers {
		observer.updateValue(i.name)
	}
}

func (i *Item) register(observer Observer) {
	i.observers = append(i.observers, observer)
}

type EmailClient struct {
	id string
}

func (ec *EmailClient) getId() string {
	return ec.id
}

func (ec *EmailClient) updateValue(value string) {
	fmt.Printf("Sending Email - %s avaiable from client %s\n", value, ec.id)
}

func MainObserver(run bool) {
	if !run {
		return
	}

	nvidiaItem := NewItem("RTX 3080")

	firstObserver := &EmailClient{
		id: "client1",
	}

	secondObserver := &EmailClient{
		id: "client2",
	}

	nvidiaItem.register(firstObserver)
	nvidiaItem.register(secondObserver)

	nvidiaItem.UpdateAvaiable()
}
