package particles

import (
	"math"
	"math/rand"
	"time"
)

type Coffee struct {
	ParticleSystem
}

func reset(particle *Particle, params *ParticleParams) {
	particle.Lifetime = int64(math.Floor(float64(params.MaxLife) * rand.Float64()))
	particle.Speed = params.MaxSpeed * rand.Float64()

	var maxX float64 = math.Floor(float64(params.MaxRows) / 2)
	particle.X = maxX + math.Max(-maxX, math.Min(rand.NormFloat64()*params.XStd, maxX))
	particle.Y = 0
}

func nextPosition(particle *Particle, deltaMS int64) {
	particle.Lifetime -= deltaMS

	if particle.Lifetime <= 0 {
		return
	}

	particle.Y += particle.Speed * (float64(deltaMS) / 1000.0)
}

func NewCoffee(width int, height int, scale float64) Coffee {
	startTime := time.Now().UnixMilli()
	ascii := func(row int, col int, counts [][]int) string {
		var count int = counts[row][col]
		if count < 2 {
			return " "
		}
		// if count == 1 {
		// 	return "."
		// }
		direction := row + int(((time.Now().UnixMilli()-startTime)/2000)%2)
		if direction%2 == 0 {
			return "}"
		}
		return "{"
	}

	return Coffee{
		ParticleSystem: NewParticleSystem(
			ParticleParams{
				MaxLife:  6000,
				MaxSpeed: 1.5,

				ParticleCount: 700,

				XStd:       scale,
				MaxRows:    width,
				MaxColumns: height,

				nextPosition: nextPosition,
				ascii:        ascii,
				reset:        reset,
			},
		),
	}
}
