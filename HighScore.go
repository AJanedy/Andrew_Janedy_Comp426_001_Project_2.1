package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func LoadHighScore(game *duckyGame) {

	file, err := os.Open("highScore.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		file.Close()
		game.highscore, _ = strconv.Atoi(line)
	}
}

func SaveHighScore(newHighScore int) {
	filePath := "highScore.txt"

	err := ioutil.WriteFile(filePath, []byte(strconv.Itoa(newHighScore)), 0644)
	if err != nil {
		panic(err)
	}
}
