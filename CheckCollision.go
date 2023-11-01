package main

import (
	"github.com/co0p/tankism/lib/collision"
	_ "github.com/lafriks/go-tiled"
)

func CheckCollision(player PlayerSprite, barrier BarrierTile, game *duckyGame) bool {
	playerBounds := collision.BoundingBox{
		X:      float64(player.xLoc),
		Y:      float64(player.yLoc),
		Width:  float64(DUCK_FRAME_WIDTH),
		Height: float64(DUCK_HEIGHT),
	}
	barrierBounds := collision.BoundingBox{
		X:      float64(barrier.xLoc),
		Y:      float64(barrier.yLoc),
		Width:  float64(barrier.width),
		Height: float64(barrier.height),
	}
	if collision.AABBCollision(playerBounds, barrierBounds) && game.player.direction == RIGHT {
		return true
	}
	return false
}
