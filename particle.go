package particlesystem

import (
	"math"
	"time"
)

type Particle struct {
	lifetime int64
	speed    float64

	x float64
	y float64
}

type ParticleParams struct {
	MaxLife  int64
	MaxSpeed float64

	ParticleCount int

	X int // max X
	Y int // max Y

	nextPosition NextPosition
	ascii        Ascii
	reset        Reset
}

type NextPosition func(particle *Particle, deltaMS int64)
type Ascii func(row int, col int, counts [][]int) rune
type Reset func(particle *Particle, params *ParticleParams)

type ParticleSystem struct {
	ParticleParams
	particles []*Particle

	lastTime int64
}

func NewParticleSystem(params ParticleParams) ParticleSystem {
	return ParticleSystem{
		ParticleParams: params,
		lastTime:       time.Now().UnixMilli(),
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

		if p.x >= float64(ps.X) || p.y >= float64(ps.Y) {
			ps.reset(p, &ps.ParticleParams)
		}
	}
}

func (ps *ParticleSystem) Display() [][]rune {
	var counts [][]int = make([][]int, 0)

	// Initializes the counts array
	for row := 0; row < ps.Y; row++ {
		var count []int = make([]int, 0)
		for col := 0; col < ps.X; col++ {
			count = append(count, 0)
		}
		counts = append(counts, count)
	}

	// Updates the counts array
	for _, p := range ps.particles {
		var row int = int(math.Floor(p.y))
		var col int = int(math.Floor(p.x))
		counts[row][col]++
	}

	// Update the grid with ASCII
	out := make([][]rune, 0)
	for r, row := range counts {
		outRow := make([]rune, 0)
		for c := range row {
			outRow = append(outRow, ps.ascii(r, c, counts))
		}
		out = append(out, outRow)
	}

	return out
}
