package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

func (game duckyGame) Draw(screen *ebiten.Image) {
	drawOptions := ebiten.DrawImageOptions{}
	for tileY := 0; tileY < game.waterMap.Level.Height; tileY += 1 {
		for tileX := 0; tileX < game.waterMap.Level.Width; tileX += 1 {
			drawOptions.GeoM.Reset()
			TileXpos := float64(game.waterMap.Level.TileWidth * tileX)
			TileYpos := float64(game.waterMap.Level.TileHeight * tileY)
			drawOptions.GeoM.Translate(TileXpos, TileYpos)
			tileToDraw :=
				game.waterMap.Level.Layers[0].Tiles[tileY*game.waterMap.Level.Width+tileX]
			ebitenTileToDraw := game.waterMap.tileHash[tileToDraw.ID]
			screen.DrawImage(ebitenTileToDraw,
				&drawOptions)
		}
	}
	//yLoc := 300.0 //when you add up and down move this to the player sprite and update it in update
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Reset()
	op.GeoM.Translate(float64(game.player.xLoc), float64(game.player.yLoc))
	screen.DrawImage(game.player.spriteSheet.SubImage(image.Rect(game.player.frame*DUCK_FRAME_WIDTH,
		game.player.direction*DUCK_HEIGHT,
		game.player.frame*DUCK_FRAME_WIDTH+DUCK_FRAME_WIDTH,
		game.player.direction*DUCK_HEIGHT+DUCK_HEIGHT)).(*ebiten.Image), op)
	op.GeoM.Reset()
	op.GeoM.Translate(float64(game.gator.xLoc), float64(game.gator.yLoc))

}

func (game duckyGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
