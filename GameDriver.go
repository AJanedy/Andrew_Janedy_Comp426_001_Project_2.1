package main

import (
	"embed"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/lafriks/go-tiled"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"log"
	"math/rand"
	"time"
)

//go:embed assets/*
var EmbeddedAssets embed.FS

const mapPath = "gameMap.tmx"

const (
	DUCK_FRAME_WIDTH  = 100
	DUCK_HEIGHT       = 92
	FRAME_COUNT       = 4
	FRAMES_PER_SHEET  = 2
	UPDATE_INTERVAL   = time.Second
	SOUND_SAMPLE_RATE = 20000
)

const (
	RIGHT = iota
	LEFT
	UP
	DOWN
)

var highScore int

func main() {
	gameMap, err := tiled.LoadFile(mapPath)
	windowWidth := gameMap.Width * gameMap.TileWidth
	windowHeight := gameMap.Height * gameMap.TileHeight
	ebiten.SetWindowSize(windowWidth, windowHeight)

	if err != nil {
		fmt.Printf("error parsing map: %s", err.Error())
	}
	ebitenImageMap := makeEbitenImagesFromMap(*gameMap)

	soundContext := audio.NewContext(SOUND_SAMPLE_RATE)

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
	breadPict := LoadEmbeddedImage("", "bread.png")
	allBread := make([]BreadSprite, 0, 5)
	for i := 0; i < cap(allBread); i++ {
		allBread = append(allBread, BreadSprite{
			bread: breadPict,
			xLoc:  75 + rand.Intn(800),
			yLoc:  75 + rand.Intn(800),
		})
	}

	duckAnimation := LoadEmbeddedImage("", "jankyDuck6.png")
	myPlayer := PlayerSprite{spriteSheet: duckAnimation,
		xLoc: windowWidth / 2,
		yLoc: windowHeight / 2,
	}

	gatorAnimation1 := LoadEmbeddedImage("", "jankyGator1.png")
	gatorAnimation2 := LoadEmbeddedImage("", "jankyGator2.png")

	enemyGator1 := GatorSprite{
		hungryGator: gatorAnimation1,
		happyGator:  gatorAnimation2,
		xLoc:        windowWidth - 200,
		yLoc:        windowHeight - 200,
		gatorSpeed:  3,
		gatorFed:    false,
	}

	enemyGator2 := GatorSprite{
		hungryGator: gatorAnimation1,
		happyGator:  gatorAnimation2,
		xLoc:        windowWidth - 700,
		yLoc:        windowHeight - 800,
		gatorSpeed:  3,
		gatorFed:    false,
	}
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}
	scoreFont, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}

	myMap := GameMap{
		Level:    gameMap,
		tileHash: ebitenImageMap,
	}
	thisGame := duckyGame{
		player:            myPlayer,
		score:             0,
		gator1:            enemyGator1,
		gator2:            enemyGator2,
		waterMap:          myMap,
		barrierTiles:      barrierTilesArray,
		collisionDetected: false,
		timerActive:       true,
		running:           true,
		breadCrumbs:       allBread,
		breadSprite:       breadPict,
		typeface:          scoreFont,
		duckSound:         LoadWav("duckSound.wav", soundContext),
		gameOverSound:     LoadWav("agh.wav", soundContext),
		boinkSound:        LoadWav("boink.wav", soundContext),
	}
	ebiten.SetWindowTitle("Janky Duck Eat Bread but Don't Be Gator Snack...the Game!")
	LoadHighScore(&thisGame)
	ebiten.RunGame(&thisGame)
}
