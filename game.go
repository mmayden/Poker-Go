package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 600
)

// Game implements ebiten.Game interface
type Game struct {
	sceneManager *SceneManager
	input        Input
}

// Layout takes the outside size (ex: window size) and returns the (logical) screen size
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

// Update proceeds game state
// Update is called every tick (1/60 [s] by default)
func (g *Game) Update() error {
	if g.sceneManager == nil {
		g.SceneManager = &SceneManager{}
		g.sceneManager.GoTo(&TitleScene{})
	}
	g.input.Update()
	if err := g.sceneManager.Update(&g.input); err != nil {
		return err
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.sceneManager.Draw(screen)
}
