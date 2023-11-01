package main

import (
	"github.com/co0p/tankism/lib/collision"
)

func CheckBreadCollision(player PlayerSprite, bread BreadSprite) bool {
	playerBounds := collision.BoundingBox{
		X:      float64(player.xLoc),
		Y:      float64(player.yLoc),
		Width:  float64(DUCK_FRAME_WIDTH),
		Height: float64(DUCK_HEIGHT),
	}
	barrierBounds := collision.BoundingBox{
		X:      float64(bread.xLoc),
		Y:      float64(bread.yLoc),
		Width:  float64(10),
		Height: float64(10),
	}
	if collision.AABBCollision(playerBounds, barrierBounds) {
		return true
	}
	return false
}
