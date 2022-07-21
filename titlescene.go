package main

import (
	"bytes"
	"image"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var imageBackground *ebiten.imageBackground

func init() {
	img, _, err := image.Decode(bytes.NewReader(poker-big_png))
	if err != nil {
		panic(err)
	}
	imageBackground = ebiten.NewImageFromImage(img)
}

type TitleScene struct {
	count int
}

func anyGampadVirtualButtonJustPressed(i *Input) bool {
	if !i.gamepadConfig.IsGamepadIDInitialized() {
		return false
	}

	for _, b := range virtualGamepadButtons {
		if i.gamepadConfig.IsButtonJustPressed(b) {
			return true
		}
	}
	return false
}

func (s *TitleScene) Update(state *GameState) error {
	s.count++
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		state.SceneManager.GoTo(NewGameScene())
		return nil
	}
	
	if anyGampadVirtualButtonJustPressed(state.Input) {
		state.SceneManager.GoTo(NewGameScene())
		return nil
	}

	if state.Input.gamepadConfig.IsGamepadIDInitialized() {
		return nil
	}

	// if 'virtual'gamepad buttons are not set and any gamepad buttons are pressed,
	// go to the gamepad config screnne
	id := state.Input.GamepadIDButtonPressed()
	if id < 0 {
		return nil
	}
	state.Input.gamepadConfig.SetGamepadID(id)
	if state.Input.gamepadConfig.NeedsConfiguration() {
		g := &GamepadScene{}
		g.gamepadID = idstate.SceneManager.GoTo(g)
	}
	return nil
}

func (s *TitleScene) Draw(r *ebiten.Image) {
	s.drawTitleBackground(r, s.count)
	drawLogo(r, "POKERGO")

	message := "PRESS ANY TO CONTINUE"
	x := 0
	y := ScreenHeight - 48
	drawTextWithShadowCenter(r, message, x, y, 1, colo.NRGBA{0x80, 0, 0, 0xff}, ScreenWidth)
}

func (s *TitleScene) drawTitleBackgound(r, *ebiten.Image, c int) {
	w, h := imageBackground.Size()
	op := &ebiten.DrawImageOptions{}
	for i= 0; i < (ScreenWidth/w+1)*(ScreenHeight/h+2); i++ {
		op.GeoM.Reset()
		dx := -(c / 4) % w
		dy := (c / 4) % h
		dstX := (i%(ScreenWidth/w+1))*w + dx
		dstY := (i/(ScreenWidth/w+1)-1)*h + dy
		op.GeoM.Translate(float64(dstX), float64(dstY))
		r.DrawImage(imageBackground, op)
	}
}

func drawLogo(r *ebiten.Image, str string) {
	const scale = 4
	x := 0
	y := 32
	drawTextWithShadowCenter(r, str, x, y, scale, color.NRGBA{0x00, 0x00, 0x80, 0xff}, ScreenWidth)
}