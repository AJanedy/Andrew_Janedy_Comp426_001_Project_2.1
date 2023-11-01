package main

import (
	"math/rand"
	"time"
)

func (game *duckyGame) Update() error {

	if game.running {

		getPlayerInput(game)
		gatorSpeed := 3

		for i, _ := range game.barrierTiles {
			if CheckBarrierCollision(game.player, game.barrierTiles[i]) {
				game.collisionDetected = true
			}
		}
		for i, _ := range game.breadCrumbs {
			if CheckBreadCollision(game.player, game.breadCrumbs[i]) {
				game.breadCrumbs[i] = BreadSprite{
					bread: game.breadSprite,
					xLoc:  75 + rand.Intn(800),
					yLoc:  75 + rand.Intn(800),
				}
				game.score += 1
			}
		}
		if CheckGatorCollision(game.player, game.gator1, game) {
			game.gatorDetected = true
			game.gator1.gatorFed = true
		}
		if CheckGatorCollision(game.player, game.gator2, game) {
			game.gatorDetected = true
			game.gator2.gatorFed = true
		}

		if game.timerActive && time.Since(game.timer) >= UPDATE_INTERVAL {
			game.gator1.direction = rand.Intn(4)
			game.gator2.direction = rand.Intn(4)
			game.timer = time.Now()
		}

		if game.gator1.direction == RIGHT && game.gator1.xLoc < 805 {
			game.gator1.xLoc += gatorSpeed
		} else if game.gator1.direction == LEFT && game.gator1.xLoc > 60 {
			game.gator1.xLoc -= gatorSpeed
		} else if game.gator1.direction == DOWN && game.gator1.yLoc < 770 {
			game.gator1.yLoc += gatorSpeed
		} else if game.gator1.direction == UP && game.gator1.yLoc > 25 {
			game.gator1.yLoc -= gatorSpeed
		}

		if game.gator2.direction == RIGHT && game.gator2.xLoc < 805 {
			game.gator2.xLoc += gatorSpeed
		} else if game.gator2.direction == LEFT && game.gator2.xLoc > 60 {
			game.gator2.xLoc -= gatorSpeed
		} else if game.gator2.direction == DOWN && game.gator2.yLoc < 770 {
			game.gator2.yLoc += gatorSpeed
		} else if game.gator2.direction == UP && game.gator2.yLoc > 25 {
			game.gator2.yLoc -= gatorSpeed
		}

		game.player.frameDelay += 1
		if game.player.frameDelay%FRAME_COUNT == 0 {
			game.player.frame += 1
			if game.player.frame >= FRAMES_PER_SHEET {
				game.player.frame = 0
			}
			if game.player.direction == LEFT && !game.collisionDetected {
				game.player.xLoc -= 8
			} else if game.player.direction == RIGHT && !game.collisionDetected {
				game.player.xLoc += 8
			} else if game.player.direction == DOWN && !game.collisionDetected {
				game.player.yLoc += 8
			} else if game.player.direction == UP && !game.collisionDetected {
				game.player.yLoc -= 8
			}
		}
	}
	return nil
}
