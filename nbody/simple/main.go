package main

import (
	"image/color"
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Body struct {
	x, y, vx, vy, ax, ay, m, size float64
}

type Game struct {
	bodies []Body
}

const G = 1.0
const dt = 0.1
const numBodies = 100 // ★ここで好きな星の数を設定できます！

func (g *Game) Update() error {
	for i := range g.bodies {
		g.bodies[i].ax = 0
		g.bodies[i].ay = 0

		for j := range g.bodies {
			if i == j {
				continue
			}
			dx := g.bodies[j].x - g.bodies[i].x
			dy := g.bodies[j].y - g.bodies[i].y
			r := math.Sqrt(dx*dx + dy*dy)

			if r < 1.0 {
				r = 1.0
			}

			a := G * (g.bodies[j].m / (r * r))

			g.bodies[i].ax += a * (dx / r)
			g.bodies[i].ay += a * (dy / r)
		}
	}

	for i := range g.bodies {
		g.bodies[i].vx += g.bodies[i].ax * dt
		g.bodies[i].vy += g.bodies[i].ay * dt
		g.bodies[i].x += g.bodies[i].vx * dt
		g.bodies[i].y += g.bodies[i].vy * dt
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, b := range g.bodies {
		vector.FillCircle(screen, float32(400+b.x), float32(300+b.y), float32(b.size), color.White, true)
	}
}

func (g *Game) Layout(w, h int) (int, int) {
	return 800, 600
}

// ★星を自動で生成する関数
func createBodies(n int) []Body {
	// 毎回違う配置になるように、現在時刻で乱数を初期化
	rand.Seed(time.Now().UnixNano())

	bodies := make([]Body, 0, n)

	// 中心に「重力源」となる巨大な星を1つ置く
	centralMass := 2000.0
	bodies = append(bodies, Body{x: 0, y: 0, vx: 0, vy: 0, m: centralMass, size: 5})

	// 残りの星（n-1個）をランダムに配置
	for i := 1; i < n; i++ {
		// ランダムな角度（0 〜 360度）と距離（50 〜 250）を決める
		angle := rand.Float64() * 2 * math.Pi
		distance := 50.0 + rand.Float64()*200.0

		// 中心の星の周りを綺麗な円軌道で回るための速度を計算
		// 公式: v = sqrt(G * M / r)
		speed := math.Sqrt(G * centralMass / distance)

		// 進行方向は、中心からの向きに対して90度横に向ける
		vx := -math.Sin(angle) * speed
		vy := math.Cos(angle) * speed

		bodies = append(bodies, Body{
			x:    math.Cos(angle) * distance,
			y:    math.Sin(angle) * distance,
			vx:   vx,
			vy:   vy,
			m:    1.0 + rand.Float64()*4.0, // 質量は 1.0 〜 5.0 でランダム
			size: 2,
		})
	}
	return bodies
}

func main() {
	g := &Game{
		// ★ここで先ほど作った自動生成関数を呼び出します
		bodies: createBodies(numBodies),
	}
	ebiten.SetWindowSize(800, 600)
	ebiten.RunGame(g)
}
