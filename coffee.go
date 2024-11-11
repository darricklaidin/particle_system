package particlesystem

import (
	"math"
	"math/rand"
)

type Coffee struct {
	ParticleSystem
}

func ascii(row int, col int, counts [][]int) rune {
	return '}'
}

func reset(particle *Particle, params *ParticleParams) {
	particle.Lifetime = int64(math.Floor(float64(params.MaxLife+1) * rand.Float64()))
	particle.Speed = math.Floor((params.MaxSpeed + 1) * rand.Float64())

	var maxX float64 = math.Floor(float64(params.X) / 2)
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
				MaxLife:       7,
				MaxSpeed:      0.5,
				ParticleCount: 100,

				reset:        reset,
				ascii:        ascii,
				nextPosition: nextPosition,
			},
		),
	}
}
