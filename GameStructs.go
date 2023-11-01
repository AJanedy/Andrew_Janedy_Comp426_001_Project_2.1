package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lafriks/go-tiled"
)

type GameMap struct {
	Level    *tiled.Map
	tileHash map[uint32]*ebiten.Image
}
type BarrierTile struct {
	barrierTile tiled.LayerTile
	xLoc        int
	yLoc        int
	height      int
	width       int
}

type PlayerSprite struct {
	spriteSheet *ebiten.Image
	xLoc        int
	yLoc        int
	direction   int
	frame       int
	frameDelay  int
}

type duckyGame struct {
	player            PlayerSprite
	gator             GatorSprite
	waterMap          GameMap
	barrierTiles      []BarrierTile
	collisionDetected bool
}

type GatorSprite struct {
	spriteSheet *ebiten.Image
	xLoc        int
	yLoc        int
	direction   int
	frame       int
	frameDelay  int
}
