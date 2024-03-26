package main

import (
	// "fmt"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	physicsVector "github.com/quartercastle/vector"
)

// Set screen height and widths constants, as well as some performance constants
const (
	screenHeight = 480
	screenWidth  = 640
	maxSpeed = 5
	maxMag = 0.05
	initialVelocityMultiplier = 3
	numberOfBoids = 200
)

// Create a flock, list of boids
var (
	flock = []*Boid{}
)

func init() {

	// Populate the flock
	for i := 0; i < numberOfBoids; i++ {

		// Create a new boid
		newBoid := Boid{
			position: physicsVector.MutableVector{
				float64(rand.Intn(int(screenWidth / 2))),
				float64(rand.Intn(int(screenHeight / 2))),
			},
			velocity: physicsVector.MutableVector{
				float64((rand.Float64() * 2 - 1) * initialVelocityMultiplier),
				float64((rand.Float64() * 2 - 1) * initialVelocityMultiplier),
			},
			acceleration: physicsVector.MutableVector{
				0,
				0,
			},
			maxSpeed: maxSpeed,
			maxMag: maxMag,
		}

		// Add it to the flock
		flock = append(flock, &newBoid)
	}
	
}

type Game struct{}

func (g *Game) Update() error {

	for _, boid := range flock {
		
		// Set the bounds, flock, and update the boid
		boid.SetBounds()
		boid.Flock(flock)
		boid.Update()

	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	
	// For each boid, show on screen
	for _, boid := range flock {
		boid.Show(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	// ebiten.SetWindowResizable(true)

	ebiten.SetWindowTitle("Boids Algorithm")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}