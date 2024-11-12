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

var dirs = [][]int{
	{-1, -1},
	{-1, 0},
	{-1, 1},

	{0, -1},
	{0, 1},

	{1, -1},
	{1, 0},
	{1, 1},
}

func countParticlesAround(row int, col int, counts [][]int) int {
	count := 0

	for _, dir := range dirs {
		r := row + dir[0]
		c := col + dir[1]

		if r < 0 || r == len(counts) || c < 0 || c == len(counts[0]) {
			continue
		}

		count += counts[row+dir[0]][col+dir[1]]
	}

	return count
}

// removes particles which have a lot of adjacent particles
func normalize(row int, col int, counts [][]int) {
	if countParticlesAround(row, col, counts) > 4 {
		counts[row][col] = 0
	}
}

func NewCoffee(width int, height int, scale float64) Coffee {
	startTime := time.Now().UnixMilli()
	ascii := func(row int, col int, counts [][]int) string {
		// normalize(row, col, counts)

		var count int = counts[row][col]
		if count < 1 {
			return " "
		}

		direction := row + int(((time.Now().UnixMilli()-startTime)/2000)%2)

		if countParticlesAround(row, col, counts) > 3 {
			if direction%2 == 0 {
				return "{"
			}
			return "}"
		}

		return "."
	}

	_ = ascii

	asciiFire := func(row int, col int, counts [][]int) string {
		count := counts[row][col]
		if count == 0 {
			return " "
		}
		if count < 3 {
			return "░"
		}
		if count < 5 {
			return "▒"
		}
		if count < 7 {
			return "▓"
		}
		return "█"
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
				ascii:        asciiFire,
				reset:        reset,
			},
		),
	}
}
