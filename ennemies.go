package main

import (
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type EnnemyState struct {
	PosX   float64
	PosY   float64
	SpeedX int
	SpeedY int
	Angle  float64
}

type Ennemy struct {
	CurrentState  EnnemyState
	PreviousState EnnemyState
}

func addEnnemy(baseSpeed int) {
	State := EnnemyState{}
	speed := rand.Intn(15+baseSpeed) + baseSpeed
	rng := rand.Float64()
	if rng < 0.25 {
		State = EnnemyState{
			PosX:   screenWidth / 2,
			PosY:   0,
			SpeedX: 0,
			SpeedY: speed,
			Angle:  3 * math.Pi / 2,
		}
	} else if rng < 0.5 {
		State = EnnemyState{
			PosX:   screenWidth / 2,
			PosY:   screenHeight,
			SpeedX: 0,
			SpeedY: -speed,
			Angle:  math.Pi / 2,
		}
	} else if rng < 0.75 {
		State = EnnemyState{
			PosX:   0,
			PosY:   screenHeight / 2,
			SpeedX: speed,
			SpeedY: 0,
			Angle:  7,
		}
	} else {
		State = EnnemyState{
			PosX:   screenWidth,
			PosY:   screenHeight / 2,
			SpeedX: -speed,
			SpeedY: 0,
			Angle:  0,
		}
	}

	newEnnemy := &Ennemy{
		CurrentState: State,
	}
	newEnnemy.PreviousState = newEnnemy.CurrentState
	spawnedEnnemies = append(spawnedEnnemies, newEnnemy)
}

func (g *Game) pickEnnemy() {
	var ecart time.Duration
	var baseSpeed int
	ecart = time.Duration(1000 - g.count/10)
	baseSpeed = 1 + g.count/1000
	if g.count > 6000 {
		ecart = 400
	}
	if rand.Float64() < float64(g.count)/6000 && time.Since(lastEnnemy) > ecart*time.Millisecond {
		lastEnnemy = time.Now()
		addEnnemy(baseSpeed)
	}
}

func (g *Game) moveEnnemies() {
	var ennemiesAlive []*Ennemy
	for _, e := range spawnedEnnemies {
		// Calculates new pos
		e.CurrentState.PosX = e.PreviousState.PosX + float64(e.PreviousState.SpeedX)/4.0
		e.CurrentState.PosY = e.PreviousState.PosY + float64(e.PreviousState.SpeedY)/4.0
		e.PreviousState = e.CurrentState

		// Check hit box
		X, Y := e.CurrentState.PosX, e.CurrentState.PosY
		if !(X > (screenWidth/2)-70 && X < (screenWidth/2)+70 && Y > (screenHeight/2)-70 && Y < (screenHeight/2)+70) {
			ennemiesAlive = append(ennemiesAlive, e)
		} else {
			// Check for success or loss
			if checkAngles(user.Position, e.CurrentState.Angle) {
				user.Score -= 1
				p := g.audioContext.NewPlayerFromBytes(hitSound)
				p.SetVolume(0.1)
				p.Play()
			} else {
				user.Score += 2
			}
			if -user.Score <= 0 {
				user.GameOver = true
			}
		}
	}
	spawnedEnnemies = ennemiesAlive
}

func (g *Game) drawAllEnnemies(screen *ebiten.Image) {
	w, h := ennemies.Size()
	op := &ebiten.DrawImageOptions{}

	for _, e := range spawnedEnnemies {
		// Handle ennemies images
		op.GeoM.Reset()
		op.GeoM.Translate(-float64(w)/2, -float64(h)/2)

		if e.CurrentState.Angle == 7 {
			op.GeoM.Scale(-1, 1)
		} else {
			op.GeoM.Rotate(e.CurrentState.Angle)
		}
		op.GeoM.Translate(e.CurrentState.PosX, e.CurrentState.PosY)
		op.ColorM.RotateHue(e.CurrentState.Angle)
		screen.DrawImage(ennemies, op)
	}
}

func checkAngles(pAng, eAng float64) bool {
	switch eAng {
	case 3 * math.Pi / 2:
		if pAng == 0 {
			return true
		}
		return false
	case math.Pi / 2:
		if pAng == math.Pi {
			return true
		}
		return false
	case 0:
		if pAng == math.Pi/2 {
			return true
		}
		return false
	case 7:
		if pAng == 3*math.Pi/2 {
			return true
		}
		return false
	}
	return false
}
