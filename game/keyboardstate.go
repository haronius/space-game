package game

import (
	"honnef.co/go/js/dom"
)

//KeyboardState holds the curent key state and handles key events
type KeyboardState struct {
	up    bool
	down  bool
	left  bool
	right bool
}

//SetLeftTrue will set left property to true
func (s *KeyboardState) SetLeftTrue(event dom.Event) {
	s.left = true
}

//SetLeftFalse will set left property to false
func (s *KeyboardState) SetLeftFalse(event dom.Event) {
	s.left = false
}

//SetRightTrue will set right property to true
func (s *KeyboardState) SetRightTrue(event dom.Event) {
	s.right = true
}

//SetRightFalse will set right property to false
func (s *KeyboardState) SetRightFalse(event dom.Event) {
	s.right = false
}

//SetUpTrue will set up property to false
func (s *KeyboardState) SetUpTrue(event dom.Event) {
	s.up = true
}

//SetUpFalse will set up property to true
func (s *KeyboardState) SetUpFalse(event dom.Event) {
	s.up = false
}

//HandleKeyDown responds to event listener by settings the appropriate key state to true
func (s *KeyboardState) HandleKeyDown(event dom.Event) {
	keyCode := event.(*dom.KeyboardEvent).KeyCode
	if keyCode == 37 {
		s.left = true
	}
	if keyCode == 38 {
		s.up = true
	}
	if keyCode == 39 {
		s.right = true
	}
}

//HandleKeyUp responds to event listener by settings the appropriate key state to false
func (s *KeyboardState) HandleKeyUp(event dom.Event) {
	keyCode := event.(*dom.KeyboardEvent).KeyCode
	if keyCode == 37 {
		s.left = false
	}
	if keyCode == 38 {
		s.up = false
	}
	if keyCode == 39 {
		s.right = false
	}
}
