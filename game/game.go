package game

import (
	"log"
	"time"
)

const frameRate = 3.0

type Game struct {
	Name       string
	Sessions   []*PlayerSession
	JoinChanel chan *PlayerSession
}

func NewGame(name string) *Game {
	game := Game{
		Name:       name,
		JoinChanel: make(chan *PlayerSession),
	}
	go game.runListen()
	go game.loop()
	return &game
}

func (g *Game) runListen() {
	for {
		p_sess := <-g.JoinChanel

		p_sess.Balls = append(p_sess.Balls, NewBall(50, 100.0, 100.0))
		p_sess.Direction = Cordinate{X: 10, Y: 10}
		g.Sessions = append(g.Sessions, p_sess)
		log.Println(g.Sessions)
		log.Printf("player %s has joined\n", p_sess.Player.Name)
	}
}

func (g *Game) loop() {
	var stepDelay int32 = 20
	for {
		time.Sleep(time.Duration(stepDelay) * time.Millisecond)
		for _, session := range g.Sessions {
			log.Printf("player %s session", session.Player.Name)
			balls_pos := session.moveBalls()
			log.Println(balls_pos)
			session.sendClientCommand(PlayerSessionCommand{Method: "player_state", Params: CommandParams{"Balls": balls_pos}})

		}
	}
}

func (g Game) JoinPlayer(session *PlayerSession) {
	g.JoinChanel <- session
}
