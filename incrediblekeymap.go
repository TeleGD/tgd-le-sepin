package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

/*
	Redefine needed input for compatibility
*/

func inputOr(inputs []bool) bool {
	for _, b := range inputs {
		if b {
			return true
		}
	}
	return false
}

func (g *Game) pressedUp() bool {
	inputs := []bool{
		inpututil.IsKeyJustPressed(ebiten.KeyArrowUp),
	}
	for _, id := range g.gamepadids {
		inputs = append(inputs, inpututil.IsStandardGamepadButtonJustPressed(id, ebiten.StandardGamepadButtonLeftTop))
		inputs = append(inputs, inpututil.IsStandardGamepadButtonJustPressed(id, ebiten.StandardGamepadButtonRightTop))
	}
	return inputOr(inputs)
}

func (g *Game) pressedDown() bool {
	inputs := []bool{
		inpututil.IsKeyJustPressed(ebiten.KeyArrowDown),
	}
	for _, id := range g.gamepadids {
		inputs = append(inputs, inpututil.IsStandardGamepadButtonJustPressed(id, ebiten.StandardGamepadButtonLeftBottom))
		inputs = append(inputs, inpututil.IsStandardGamepadButtonJustPressed(id, ebiten.StandardGamepadButtonRightBottom))
	}
	return inputOr(inputs)
}

func (g *Game) pressedRight() bool {
	inputs := []bool{
		inpututil.IsKeyJustPressed(ebiten.KeyArrowRight),
	}
	for _, id := range g.gamepadids {
		inputs = append(inputs, inpututil.IsStandardGamepadButtonJustPressed(id, ebiten.StandardGamepadButtonLeftRight))
		inputs = append(inputs, inpututil.IsStandardGamepadButtonJustPressed(id, ebiten.StandardGamepadButtonRightRight))
	}
	return inputOr(inputs)
}

func (g *Game) pressedLeft() bool {
	inputs := []bool{
		inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft),
	}
	for _, id := range g.gamepadids {
		inputs = append(inputs, inpututil.IsStandardGamepadButtonJustPressed(id, ebiten.StandardGamepadButtonLeftLeft))
		inputs = append(inputs, inpututil.IsStandardGamepadButtonJustPressed(id, ebiten.StandardGamepadButtonRightLeft))
	}
	return inputOr(inputs)
}

func (g *Game) pressedEsc() bool {
	inputs := []bool{
		inpututil.IsKeyJustPressed(ebiten.KeyEscape),
	}
	for _, id := range g.gamepadids {
		inputs = append(inputs, inpututil.IsStandardGamepadButtonJustPressed(id, ebiten.StandardGamepadButtonCenterCenter))
		inputs = append(inputs, inpututil.IsStandardGamepadButtonJustPressed(id, ebiten.StandardGamepadButtonCenterRight))
		inputs = append(inputs, inpututil.IsStandardGamepadButtonJustPressed(id, ebiten.StandardGamepadButtonCenterLeft))
	}
	return inputOr(inputs)
}

func (g *Game) pressedEnter() bool {
	inputs := []bool{
		inpututil.IsKeyJustPressed(ebiten.KeyEnter),
		inpututil.IsKeyJustPressed(ebiten.KeyKPEnter),
	}
	for _, id := range g.gamepadids {
		inputs = append(inputs, inpututil.IsStandardGamepadButtonJustPressed(id, ebiten.StandardGamepadButtonFrontBottomLeft))
		inputs = append(inputs, inpututil.IsStandardGamepadButtonJustPressed(id, ebiten.StandardGamepadButtonFrontBottomRight))
		inputs = append(inputs, inpututil.IsStandardGamepadButtonJustPressed(id, ebiten.StandardGamepadButtonFrontTopLeft))
		inputs = append(inputs, inpututil.IsStandardGamepadButtonJustPressed(id, ebiten.StandardGamepadButtonFrontTopRight))
	}
	return inputOr(inputs)
}
