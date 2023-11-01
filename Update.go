package main

import "fmt"

func (game *duckyGame) Update() error {
	getPlayerInput(game)

	for i, _ := range game.barrierTiles {
		if CheckCollision(game.player, game.barrierTiles[i], game) {
			game.collisionDetected = true
		}
	}

	fmt.Printf("%b  %d\n", game.collisionDetected, game.player.direction)

	game.player.frameDelay += 1
	if game.player.frameDelay%FRAME_COUNT == 0 {
		game.player.frame += 1
		if game.player.frame >= FRAMES_PER_SHEET {
			game.player.frame = 0
		}
		if game.player.direction == LEFT && !game.collisionDetected {
			game.player.xLoc -= 5
		} else if game.player.direction == RIGHT && !game.collisionDetected {
			game.player.xLoc += 5
		} else if game.player.direction == DOWN && !game.collisionDetected {
			game.player.yLoc += 5
		} else if game.player.direction == UP && !game.collisionDetected {
			game.player.yLoc -= 5
		}
	}
	return nil
}
