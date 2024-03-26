package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	ebitVector "github.com/hajimehoshi/ebiten/v2/vector"
	physicsVector "github.com/quartercastle/vector"
)

func setMag(currentVector physicsVector.MutableVector, desiredMag float64) physicsVector.MutableVector {

	// The normalized magnitude times desired
	return currentVector.Unit().Scale(desiredMag) 
}

func limitMag(vector physicsVector.MutableVector, maxMag float64) physicsVector.MutableVector {

	// If the maxMag is over shot, scale it down to be within bounds
	if vector.Magnitude() > maxMag {
		return setMag(vector, maxMag)
	}
	return vector
}

type Boid struct {
	position     physicsVector.MutableVector
	velocity     physicsVector.MutableVector
	acceleration physicsVector.MutableVector
	maxSpeed float64
	maxMag float64
}

func (b *Boid) distance(other *Boid) float64 {

	// Distance formula
	dx := b.position.X() - other.position.X()
	dy := b.position.Y() - other.position.Y()
	return math.Sqrt(dx*dx + dy*dy)
}

func (b *Boid) SetBounds() {

	// If gonna over on either X-side, set to opposite side
	if b.position.X() > (screenWidth/2)+5 {
		b.position[0] = 0
	} else if (b.position.X()) < -5 {
		b.position[0] = (screenWidth / 2) + 5
	}

	// If gonna over on either Y-side, set to opposite side
	if b.position.Y() > (screenHeight/2)+5 {
		b.position[1] = 0
	} else if (b.position.Y()) < -5 {
		b.position[1] = (screenHeight / 2) + 5
	}
}



func (b*Boid) addToAverageAlign (boid *Boid, steeringVector physicsVector.MutableVector, boidDistance float64) physicsVector.MutableVector {
	return steeringVector.Add(physicsVector.Vector(boid.velocity))
}
func (b*Boid) addToAverageCohesion (boid *Boid, steeringVector physicsVector.MutableVector, boidDistance float64) physicsVector.MutableVector {
	return steeringVector.Add(physicsVector.Vector(boid.position))
}
func (b*Boid) addToAverageSeparation (boid *Boid, steeringVector physicsVector.MutableVector, boidDistance float64) physicsVector.MutableVector {
		// Create a difference vector from a cloned position: myBoidPosition - nearByBoid
		difference := b.position.Clone().Sub(physicsVector.Vector(boid.position))
		
		// Divide my the distance between them, as item gets closer magnitude of difference is increased
		difference = difference.Scale(1.0 / float64(boidDistance))
		
		// Add to be averaged
		return steeringVector.Add(physicsVector.Vector(difference))
}

func (b *Boid) averageFinalAlignVector(steeringVector physicsVector.MutableVector, numberSurroundingBoids int) physicsVector.MutableVector {

	steeringVector.Scale(1.0 / float64(numberSurroundingBoids))
	steeringVector = setMag(steeringVector, b.maxSpeed)
	steeringVector.Sub(physicsVector.Vector(b.velocity))
	return limitMag(steeringVector, b.maxMag)
}

func (b *Boid) averageFinalCohesionVector(steeringVector physicsVector.MutableVector, numberSurroundingBoids int) physicsVector.MutableVector {

	steeringVector.Scale(1.0 / float64(numberSurroundingBoids))
	steeringVector.Sub(physicsVector.Vector(b.position))
	steeringVector = setMag(steeringVector, b.maxSpeed)
	steeringVector.Sub(physicsVector.Vector(b.velocity))
	return limitMag(steeringVector, b.maxMag)
}
func (b *Boid) averageFinalSeparationVector(steeringVector physicsVector.MutableVector, numberSurroundingBoids int) physicsVector.MutableVector {
	steeringVector.Scale(1.0 / float64(numberSurroundingBoids))
	steeringVector = setMag(steeringVector, b.maxSpeed)
	steeringVector.Sub(physicsVector.Vector(b.velocity))
	return limitMag(steeringVector, b.maxMag)
}
func (b *Boid) ModifyVector(
	boids []*Boid, 
	perception float64, 
	addToAverageFunc func(*Boid, physicsVector.MutableVector, float64) physicsVector.MutableVector,
	averageFinalFunc func(physicsVector.MutableVector, int) physicsVector.MutableVector) (physicsVector.MutableVector) {

	// Create an empty vector as a steering force, will be modified later to an average
	steering := physicsVector.MutableVector{0, 0}
	numberSurroundingBoids := 0

	for _, boid  := range boids {

		// Get the distance from the boid to the current boid object
		boidDistance := b.distance(boid)

		// If not the same boid and within perception radius
		if  boid != b && boidDistance <= perception {

			// UNIQUE: Return steering vector
			steering = addToAverageFunc(boid, steering, boidDistance)

			// Increase count
			numberSurroundingBoids++
		}
	}

	// If we found a nearby node
	if numberSurroundingBoids > 0 {

		// UNIQUE: Ajusting the final steering vector
		steering = averageFinalFunc(steering, numberSurroundingBoids)
		
	}

	return steering
}

func (b *Boid) Flock(boids []*Boid) {

	alignment := b.ModifyVector(boids, 50, b.addToAverageAlign, b.averageFinalAlignVector)
	cohesion := b.ModifyVector(boids, 50, b.addToAverageCohesion, b.averageFinalCohesionVector)
	separation := b.ModifyVector(boids, 50, b.addToAverageSeparation, b.averageFinalSeparationVector)

	b.acceleration.Add(physicsVector.Vector(alignment).Scale(0.1))
	b.acceleration.Add(physicsVector.Vector(separation).Scale(1.15))
	b.acceleration.Add(physicsVector.Vector(cohesion).Scale(1.3))
}

func (b *Boid) Update() {

	b.position.Add(physicsVector.Vector(b.velocity))
	b.velocity.Add(physicsVector.Vector(b.acceleration))
	b.velocity = limitMag(b.velocity, b.maxSpeed)
	b.acceleration.Scale(float64(0))
}

func (b *Boid) Show(screen *ebiten.Image) {

	rValue := uint8(math.Abs(math.Sin(b.position.X()/100)) * 255)	
	gValue := uint8(math.Abs(math.Sin(b.position.Y()/100)) * 255)
	bValue := uint8((math.Abs(math.Cos(b.position.X()/100)) + math.Abs(math.Cos(b.position.Y()/100))) / 2 * 255)

	ebitVector.DrawFilledCircle(screen, float32(b.position.X()), float32(b.position.Y()), 1.5, color.RGBA{
		R: rValue,
		G: gValue,
		B: bValue,
		A: uint8(255),
	}, false)
}
