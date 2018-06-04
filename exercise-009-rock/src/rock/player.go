package main

import (
	"math/rand"
)

type Player interface {
  Type() string
  Play() int
}

type RandoRex struct {
}

// Type returns the type of the player
func (r *RandoRex) Type() string {
	return "RandoRex"
}

// Play returns a move
func (r *RandoRex) Play() int {
	choice := rand.Int() % 3
	return choice
}

type Obsessed struct {
  move int
}

func (o *Obsessed) Type() string {
  return "Obsessed"
}

func (o *Obsessed) Play() int {
  return o.move
}

type Flipper struct {
  move1 int
  move2 int
}

func (f *Flipper) Type() string {
  return "Flipper"
}

func (f *Flipper) Play() int {
  choice := rand.Int() % 2
  if choice == 0 {
    return f.move1
  } else {
    return f.move2
  }
}

var prev_choice = 0
type Cyclone struct {
}

func (c *Cyclone) Type() string {
  return "Cyclone"
}

func (c *Cyclone) Play() int {
  if prev_choice == 3 {
    prev_choice = 0
  }
  prev_choice += 1
  return prev_choice
}