package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"strconv"
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
	op := &ebiten.DrawImageOptions{}

	if !game.gator1.gatorFed && !game.gator2.gatorFed {
		op.GeoM.Reset()
		op.GeoM.Translate(float64(game.gator1.xLoc), float64(game.gator1.yLoc))
		screen.DrawImage(game.gator1.hungryGator, op)
		op.GeoM.Reset()
		op.GeoM.Translate(float64(game.gator2.xLoc), float64(game.gator2.yLoc))
		screen.DrawImage(game.gator2.hungryGator, op)
		op.GeoM.Reset()
		op.GeoM.Translate(float64(game.player.xLoc), float64(game.player.yLoc))
		screen.DrawImage(game.player.spriteSheet.SubImage(image.Rect(game.player.frame*DUCK_FRAME_WIDTH,
			game.player.direction*DUCK_HEIGHT,
			game.player.frame*DUCK_FRAME_WIDTH+DUCK_FRAME_WIDTH,
			game.player.direction*DUCK_HEIGHT+DUCK_HEIGHT)).(*ebiten.Image), op)
	} else if game.gator1.gatorFed {
		op.GeoM.Reset()
		op.GeoM.Translate(float64(game.gator1.xLoc), float64(game.gator1.yLoc))
		screen.DrawImage(game.gator1.happyGator, op)
		op.GeoM.Reset()
		op.GeoM.Translate(float64(game.gator2.xLoc), float64(game.gator2.yLoc))
		screen.DrawImage(game.gator2.hungryGator, op)
	} else if game.gator2.gatorFed {
		op.GeoM.Reset()
		op.GeoM.Translate(float64(game.gator1.xLoc), float64(game.gator1.yLoc))
		screen.DrawImage(game.gator1.hungryGator, op)
		op.GeoM.Reset()
		op.GeoM.Translate(float64(game.gator2.xLoc), float64(game.gator2.yLoc))
		screen.DrawImage(game.gator2.happyGator, op)
	}
	for _, breadCrumb := range game.breadCrumbs {
		op.GeoM.Reset()
		op.GeoM.Translate(float64(breadCrumb.xLoc), float64(breadCrumb.yLoc))
		screen.DrawImage(breadCrumb.bread, op)
	}
	DrawCenteredText(screen, game.typeface, "High Score: "+strconv.Itoa(game.highscore), 200, 100)
	DrawCenteredText(screen, game.typeface, "Score: "+strconv.Itoa(game.score), 800, 100)
}

func (game duckyGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
