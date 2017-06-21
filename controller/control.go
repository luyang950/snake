package controller

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"snake/snake"
)

func StartControl(stopGameSignal chan bool, startControlSignal chan bool, player *snake.Snake) {
	select {
	case <-startControlSignal:
		for {
			event_queue := make(chan termbox.Event)
			go func() {
				for {
					event_queue <- termbox.PollEvent()
				}
			}()
			ev := <-event_queue

			switch ev.Type {
			case termbox.EventKey:
				switch ev.Key {
				case termbox.KeyArrowUp:
					if player.Head.Next != nil {
						if player.OppoDirection() == snake.Up {
							continue
						}
					}
					player.HeadUp()
				case termbox.KeyArrowDown:
					if player.Head.Next != nil {
						if player.OppoDirection() == snake.Down {
							continue
						}
					}
					player.HeadDown()
				case termbox.KeyArrowLeft:
					if player.Head.Next != nil {
						if player.OppoDirection() == snake.Left {
							continue
						}
					}
					player.HeadLeft()
				case termbox.KeyArrowRight:
					if player.Head.Next != nil {
						if player.OppoDirection() == snake.Right {
							continue
						}
					}
					player.HeadRight()
				case termbox.KeyCtrlC:
					stopGameSignal <- true
				default:
					fmt.Println("other")
				}
			}
		}
	}
}
