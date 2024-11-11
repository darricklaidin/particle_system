package particles

import (
	"math"
	"math/rand"
)

type Coffee struct {
	ParticleSystem
}

func ascii(row int, col int, counts [][]int) string {
	var count int = counts[row][col]
	if count < 3 {
		return " "
	}
	if count < 6 {
		return "."
	}
	if count < 9 {
		return ":"
	}
	if count < 12 {
		return "{"
	}
	return "}"
}

func reset(particle *Particle, params *ParticleParams) {
	particle.Lifetime = int64(math.Floor(float64(params.MaxLife) * rand.Float64()))
	particle.Speed = params.MaxSpeed * rand.Float64()

	var maxX float64 = math.Floor(float64(params.MaxRows) / 2)
	// FIXME: this is not normal distribution probability
	particle.X = maxX + math.Max(-maxX, math.Min(rand.NormFloat64(), maxX))
	particle.Y = 0
}

func nextPosition(particle *Particle, deltaMS int64) {
	particle.Lifetime -= deltaMS

	if particle.Lifetime <= 0 {
		return
	}

	particle.Y += particle.Speed * (float64(deltaMS) / 1000.0)
}

func NewCoffee(width int, height int) Coffee {
	return Coffee{
		ParticleSystem: NewParticleSystem(
			ParticleParams{
				MaxLife:  7000,
				MaxSpeed: 1,

				ParticleCount: 100,

				MaxRows:    width,
				MaxColumns: height,

				nextPosition: nextPosition,
				ascii:        ascii,
				reset:        reset,
			},
		),
	}
}
