package main

import (
	"github.com/co0p/tankism/lib/collision"
	_ "github.com/lafriks/go-tiled"
)

func CheckGatorCollision(player PlayerSprite, gator GatorSprite, game *duckyGame) bool {
	playerBounds := collision.BoundingBox{
		X:      float64(player.xLoc),
		Y:      float64(player.yLoc),
		Width:  float64(DUCK_FRAME_WIDTH),
		Height: float64(DUCK_HEIGHT),
	}
	gatorBounds := collision.BoundingBox{
		X:      float64(gator.xLoc + 30),
		Y:      float64(gator.yLoc + 25),
		Width:  float64(50),
		Height: float64(50),
	}
	if collision.AABBCollision(playerBounds, gatorBounds) {
		game.running = false
		return true
	}
	return false
}
