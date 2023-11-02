package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"os"
)

func LoadWav(soundFile string, context *audio.Context) *audio.Player {
	pewFile, err := os.Open(soundFile)
	if err != nil {
		fmt.Println("Error Loading sound: ", err)
	}
	pewSound, err := wav.DecodeWithoutResampling(pewFile)
	if err != nil {
		fmt.Println("Error interpreting sound file: ", err)
	}
	soundPlayer, err := context.NewPlayer(pewSound)
	if err != nil {
		fmt.Println("Couldn't create sound player: ", err)
	}
	return soundPlayer
}
