package gamebox

import (
	"fmt"
	"math/rand"
	"snake/food"
	"snake/snake"
	"time"
)

// an 2-dim array to represent the game ground
// 0: empty
// 1: brick
// 2: snake body(player)
// 3: food
const (
	groundWidth  = 30
	groundHeight = 10

	brick     = 1
	body      = 2
	sweetFood = 3
)

type GameBox struct {
	fps   int
	speed int

	width  int
	height int

	ground [groundHeight][groundWidth]int

	brickTexture string
	snakeTexture string
	foodTexture  string

	player *snake.Snake
	food   *food.Food
}

func (g *GameBox) Init(startSignal chan bool, genFoodSignal chan bool, fps, speed int, brickTexture, snakeTexture, foodTexture string) {

	g.fps = fps
	g.speed = speed
	g.width = groundWidth
	g.height = groundHeight
	g.brickTexture = brickTexture
	g.snakeTexture = snakeTexture
	g.foodTexture = foodTexture

	// top wall
	for i := 0; i < g.width; i++ {
		g.ground[0][i] = brick
	}

	// left&right wall
	for i := 1; i < g.height-1; i++ {
		g.ground[i][0] = brick

		for j := 1; j < g.width-2; j++ {
			g.ground[i][j] = 0
		}
		g.ground[i][g.width-1] = brick
	}

	// bottom wall
	for j := 0; j < g.width; j++ {
		g.ground[g.height-1][j] = brick
	}

	g.player = &snake.Snake{}

	startSignal <- true
	genFoodSignal <- true
}

func (g *GameBox) Player() *snake.Snake {
	return g.player
}

func (g *GameBox) FPS() int {
	return g.fps
}

func (g *GameBox) Speed() int {
	return g.speed
}

func (g *GameBox) Draw() {
	for i := 0; i < g.height; i++ {
		fmt.Println()
		for j := 0; j < g.width; j++ {
			switch g.ground[i][j] {
			case 0:
				fmt.Print(" ")
			case 1:
				fmt.Print(g.brickTexture)
			case 2:
				fmt.Print(g.snakeTexture)
			case 3:
				fmt.Print(g.foodTexture)
			}
		}
	}
}

func (g *GameBox) Update(score *int, initPlayerSignal, genFoodSignal, growSignal, dieSignal chan bool) {
	select {
	case <-initPlayerSignal:
		for {
			select {
			case <-growSignal:
			default:
				g.player.AutoMove()
			}

			var food = food.Food{}

			// sweep the ground
			for i := 1; i < g.height-1; i++ {
				for j := 1; j < g.width-2; j++ {
					if g.ground[i][j] == sweetFood {
						food.X = i
						food.Y = j
					} else {
						g.ground[i][j] = 0
					}
				}
			}

			// put player into the ground
			cur := g.player.Head
			for {
				// todo
				// there is a small bug: player can turn head to the back if operate fast enough
				if g.ground[cur.X][cur.Y] == brick /*|| g.ground[cur.X][cur.Y] == body*/ {
					g.player.Die(dieSignal)
				}
				if cur.X == food.X && cur.Y == food.Y {
					g.player.Grow()
					*score++
					growSignal <- true
					genFoodSignal <- true

				}
				g.ground[cur.X][cur.Y] = body
				if cur.Next != nil {
					cur = cur.Next
				} else {
					break
				}
			}

			time.Sleep(1000 / time.Duration(g.speed) * time.Millisecond)
		}
	}
}

func (g *GameBox) GenFood(genFoodSignal chan bool) food.Food {
	fmt.Println("gen food")
	for {
		select {
		case <-genFoodSignal:
			var emptySpaces = make([]food.Food, 0)

			for i := 1; i < groundHeight-1; i++ {
				for j := 1; j < groundWidth-1; j++ {
					if g.ground[i][j] == 0 {
						emptySpace := food.Food{
							X: i,
							Y: j,
						}
						emptySpaces = append(emptySpaces, emptySpace)
					}
				}
			}

			randIndex := rand.Intn(len(emptySpaces))
			fmt.Print(randIndex)
			g.ground[emptySpaces[randIndex].X][emptySpaces[randIndex].Y] = sweetFood
		}
	}
}
