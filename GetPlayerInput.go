package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func getPlayerInput(game *duckyGame) {
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) && game.player.xLoc > 0 {
		game.player.direction = LEFT
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowRight) &&
		game.player.xLoc < (game.waterMap.Level.Width*game.waterMap.Level.TileWidth)-DUCK_FRAME_WIDTH {
		game.player.direction = RIGHT
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowDown) && game.player.yLoc > 0 {
		game.player.direction = DOWN
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowUp) &&
		game.player.yLoc < (game.waterMap.Level.Height*game.waterMap.Level.TileHeight)-DUCK_HEIGHT {
		game.player.direction = UP
	}

}
