package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"os"
	"os/exec"
	"snake/controller"
	"snake/gamebox"
	"time"
)

const (
	fps   = 60
	speed = 5

	brickTexture = "#"
	snakeTexture = "*"
	foodTexture  = "o"
)

var game = gamebox.GameBox{}
var score = 0

func initGame(startSignal, genFoodSignal chan bool) {
	game.Init(startSignal, genFoodSignal, fps, speed, brickTexture, snakeTexture, foodTexture)
}

func initPlayer(initPlayerSignal chan bool, startControlSignal chan bool) {
	game.Player().Init()

	initPlayerSignal <- true
	startControlSignal <- true
}

func frame(startSignal chan bool) {
	select {
	case <-startSignal:
		for {
			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			cmd.Run()

			fmt.Println(time.Now())
			fmt.Println("fps:", game.FPS(), "speed:", game.Speed())
			fmt.Println("Score:", score)
			game.Draw()
			time.Sleep(1000 * time.Millisecond / time.Duration(game.FPS()))
		}
	}

}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	var startSignal = make(chan bool, 1)
	var initPlayerSignal = make(chan bool, 1)
	var startControlSignal = make(chan bool, 1)
	var stopGameSignal = make(chan bool, 1)
	var genFoodSignal = make(chan bool, 1)
	var growSignal = make(chan bool, 1)
	var dieSignal = make(chan bool, 1)

	initGame(startSignal, genFoodSignal)
	initPlayer(initPlayerSignal, startControlSignal)

	go frame(startSignal)
	go game.Update(&score, initPlayerSignal, genFoodSignal, growSignal, dieSignal)
	go controller.StartControl(stopGameSignal, startControlSignal, game.Player())
	go game.GenFood(genFoodSignal)

	// todo ending screen
	for {
		select {
		case <-dieSignal:
			return
		case <-stopGameSignal:
			return
		}
	}
}
