package main

import (
	"embed"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lafriks/go-tiled"
)

//go:embed assets/*
var EmbeddedAssets embed.FS

const mapPath = "gameMap.tmx"

const (
	DUCK_FRAME_WIDTH = 100
	DUCK_HEIGHT      = 92
	FRAME_COUNT      = 4
	FRAMES_PER_SHEET = 2
)

const (
	RIGHT = iota
	LEFT
	UP
	DOWN
)

func main() {
	gameMap, err := tiled.LoadFile(mapPath)
	windowWidth := gameMap.Width * gameMap.TileWidth
	windowHeight := gameMap.Height * gameMap.TileHeight

	ebiten.SetWindowSize(windowWidth, windowHeight)
	if err != nil {
		fmt.Printf("error parsing map: %s", err.Error())
	}
	ebitenImageMap := makeEbitenImagesFromMap(*gameMap)

	barrierTilesArray := make([]BarrierTile, 0, 56)

	for tileY := 0; tileY < gameMap.Height; tileY += 1 {
		for tileX := 0; tileX < gameMap.Width; tileX += 1 {
			TileXpos := float64(gameMap.TileWidth * tileX)
			TileYpos := float64(gameMap.TileHeight * tileY)
			TileHeight := float64(gameMap.TileHeight)
			TileWidth := float64(gameMap.TileWidth)

			if gameMap.Layers[0].Tiles[tileY*gameMap.Width+tileX].ID == 0 {
				barrierTilesArray = append(barrierTilesArray, BarrierTile{
					barrierTile: *gameMap.Layers[0].Tiles[tileY*gameMap.Width+tileX],
					xLoc:        int(TileXpos),
					yLoc:        int(TileYpos),
					height:      int(TileHeight),
					width:       int(TileWidth),
				})
			}
		}
	}

	for i, _ := range barrierTilesArray {
		fmt.Printf("xLoc: %d  yLoc: %d  tile#: %d\n",
			barrierTilesArray[i].xLoc,
			barrierTilesArray[i].yLoc,
			i)
	}

	duckAnimation := LoadEmbeddedImage("", "jankyDuck6.png")
	myPlayer := PlayerSprite{spriteSheet: duckAnimation,
		xLoc: windowWidth / 2,
		yLoc: windowHeight / 2,
	}
	gatorAnimation := LoadEmbeddedImage("", "jankyGator.png")
	enemyGator := GatorSprite{spriteSheet: gatorAnimation,
		xLoc: windowWidth + 100,
		yLoc: windowHeight - 100}

	myMap := GameMap{
		Level:    gameMap,
		tileHash: ebitenImageMap,
	}
	thisGame := duckyGame{
		player:            myPlayer,
		gator:             enemyGator,
		waterMap:          myMap,
		barrierTiles:      barrierTilesArray,
		collisionDetected: false,
	}
	ebiten.SetWindowTitle("Animated Sprite")
	ebiten.RunGame(&thisGame)
}
