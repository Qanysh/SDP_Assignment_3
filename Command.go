package main

import "fmt"

type Command interface {
	Execute()
}

type Light struct {
	On bool
}

func (l *Light) TurnOn() {
	l.On = true
	fmt.Println("Свет включен")
}

func (l *Light) TurnOff() {
	l.On = false
	fmt.Println("Свет выключен")
}

type TurnOnLightCommand struct {
	Light *Light
}

func (c *TurnOnLightCommand) Execute() {
	c.Light.TurnOn()
}

type TurnOffLightCommand struct {
	Light *Light
}

func (c *TurnOffLightCommand) Execute() {
	c.Light.TurnOff()
}

type RemoteControl struct {
	Command Command
}

func (r *RemoteControl) PressButton() {
	r.Command.Execute()
}

func main() {
	light := &Light{}
	turnOnCommand := &TurnOnLightCommand{Light: light}
	turnOffCommand := &TurnOffLightCommand{Light: light}

	remote := &RemoteControl{}

	remote.Command = turnOnCommand
	remote.PressButton()

	remote.Command = turnOffCommand
	remote.PressButton()
}
