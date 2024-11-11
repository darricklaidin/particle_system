package particles

import (
	"math"
	"slices"
	"strings"
	"time"
)

type Particle struct {
	Lifetime int64
	Speed    float64

	X float64
	Y float64
}

type ParticleParams struct {
	MaxLife  int64
	MaxSpeed float64

	ParticleCount int

	MaxRows    int
	MaxColumns int

	nextPosition NextPosition
	ascii        Ascii
	reset        Reset
}

type NextPosition func(particle *Particle, deltaMS int64)
type Ascii func(row int, col int, counts [][]int) string
type Reset func(particle *Particle, params *ParticleParams)

type ParticleSystem struct {
	ParticleParams
	particles []*Particle

	lastTime int64
}

func NewParticleSystem(params ParticleParams) ParticleSystem {
	particles := make([]*Particle, 0)
	for i := 0; i < params.ParticleCount; i++ {
		particles = append(particles, &Particle{})
	}
	return ParticleSystem{
		ParticleParams: params,
		particles:      particles,

		lastTime: time.Now().UnixMilli(),
	}
}

func (ps *ParticleSystem) Start() {
	for _, p := range ps.particles {
		ps.reset(p, &ps.ParticleParams)
	}
}

func (ps *ParticleSystem) Update() {
	var now int64 = time.Now().UnixMilli()
	var deltaMS int64 = now - ps.lastTime
	ps.lastTime = now

	for _, p := range ps.particles {
		ps.nextPosition(p, deltaMS)

		if p.X >= float64(ps.MaxRows) || p.Y >= float64(ps.MaxColumns) || p.Lifetime <= 0 {
			ps.reset(p, &ps.ParticleParams)
		}
	}
}

func (ps *ParticleSystem) Display() string {
	var counts [][]int = make([][]int, 0)

	// Initializes the counts array
	for row := 0; row < ps.MaxColumns; row++ {
		var count []int = make([]int, 0)
		for col := 0; col < ps.MaxRows; col++ {
			count = append(count, 0)
		}
		counts = append(counts, count)
	}

	// Updates the counts array
	for _, p := range ps.particles {
		var row int = int(math.Floor(p.Y))
		// Use math.Round here to include maxX
		var col int = int(math.Round(p.X))
		counts[row][col]++
	}

	// Update the grid with ASCII
	var out [][]string = make([][]string, 0)
	for r, row := range counts {
		var outRow []string = make([]string, 0)
		for c := range row {
			outRow = append(outRow, ps.ascii(r, c, counts))
		}
		out = append(out, outRow)
	}

	slices.Reverse(out)
	var outStr []string = make([]string, 0)
	for _, row := range out {
		outStr = append(outStr, strings.Join(row, ""))
	}

	return strings.Join(outStr, "\n")
}
