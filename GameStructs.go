package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lafriks/go-tiled"
	"golang.org/x/image/font"
	"time"
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
	score             int
	gator1            GatorSprite
	gator2            GatorSprite
	waterMap          GameMap
	breadSprite       *ebiten.Image
	barrierTiles      []BarrierTile
	collisionDetected bool
	gatorDetected     bool
	timer             time.Time
	timerActive       bool
	running           bool
	breadCrumbs       []BreadSprite
	typeface          font.Face
}

type GatorSprite struct {
	hungryGator *ebiten.Image
	happyGator  *ebiten.Image
	xLoc        int
	yLoc        int
	gatorFed    bool
	direction   int
	frame       int
	frameDelay  int
}
type BreadSprite struct {
	bread *ebiten.Image
	xLoc  int
	yLoc  int
}
