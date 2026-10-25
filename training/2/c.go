package main

import (
	"fmt"
	"math"
)

type Titik struct {
	x, y float64
}

func hitung(i int, j int, cx [3]float64, cy [3]float64, r [3]float64) (float64, float64, float64, float64) {
	var a, b, c, d float64
	a = r[j]*r[j] - r[i]*r[i]
	b = -2 * (r[j]*r[j]*cx[i] - r[i]*r[i]*cx[j])
	c = -2 * (r[j]*r[j]*cy[i] - r[i]*r[i]*cy[j])
	d = r[j]*r[j]*(cx[i]*cx[i]+cy[i]*cy[i]) - r[i]*r[i]*(cx[j]*cx[j]+cy[j]*cy[j])
	return a, b, c, d
}

func potongGarisLingkaran(la float64, lb float64, lc float64, A float64, B float64, C float64, D float64) []Titik {
	var res []Titik
	res = make([]Titik, 0)
	var h, k, R float64
	h = -B / (2 * A)
	k = -C / (2 * A)
	var temp float64
	temp = (B*B+C*C)/(4*A*A) - D/A
	if temp < 0 {
		temp = 0
	}
	R = math.Sqrt(temp)

	var val, den float64
	val = la*h + lb*k + lc
	den = la*la + lb*lb
	if den < 1e-12 {
		return res
	}
	var dist, h2 float64
	dist = val * val / den
	h2 = R*R - dist
	if h2 < -1e-9 {
		return res
	}
	if h2 < 0 {
		h2 = 0
	}
	var px, py, hh float64
	px = h - la*val/den
	py = k - lb*val/den
	hh = math.Sqrt(h2 / den)
	res = append(res, Titik{px - lb*hh, py + la*hh})
	if hh > 1e-12 {
		res = append(res, Titik{px + lb*hh, py - la*hh})
	}
	return res
}

func main() {
	var cx, cy, r [3]float64
	var i, j int
	for i = 0; i < 3; i++ {
		fmt.Scan(&cx[i], &cy[i], &r[i])
	}

	var A1, B1, C1, D1 float64
	var A2, B2, C2, D2 float64
	A1, B1, C1, D1 = hitung(0, 1, cx, cy, r)
	A2, B2, C2, D2 = hitung(0, 2, cx, cy, r)

	var kandidat []Titik
	kandidat = make([]Titik, 0)

	var eps float64
	eps = 1e-12

	if math.Abs(A1) < eps && math.Abs(A2) < eps {
		var det float64
		det = B1*C2 - B2*C1
		if math.Abs(det) > eps {
			var px, py float64
			px = (C1*D2 - C2*D1) / det
			py = (D1*B2 - B1*D2) / det
			kandidat = append(kandidat, Titik{px, py})
		}
	} else if math.Abs(A1) < eps {
		kandidat = append(kandidat, potongGarisLingkaran(B1, C1, D1, A2, B2, C2, D2)...)
	} else if math.Abs(A2) < eps {
		kandidat = append(kandidat, potongGarisLingkaran(B2, C2, D2, A1, B1, C1, D1)...)
	} else {
		var la, lb, lc float64
		la = A2*B1 - A1*B2
		lb = A2*C1 - A1*C2
		lc = A2*D1 - A1*D2
		kandidat = append(kandidat, potongGarisLingkaran(la, lb, lc, A1, B1, C1, D1)...)
	}

	var terbaik float64
	terbaik = 1e18
	var hasilX, hasilY float64
	var ada bool
	ada = false

	for j = 0; j < len(kandidat); j++ {
		var p Titik
		p = kandidat[j]
		var d0, d1, d2 float64
		d0 = math.Hypot(p.x-cx[0], p.y-cy[0])
		d1 = math.Hypot(p.x-cx[1], p.y-cy[1])
		d2 = math.Hypot(p.x-cx[2], p.y-cy[2])
		if d0 < r[0]-1e-6 || d1 < r[1]-1e-6 || d2 < r[2]-1e-6 {
			continue
		}
		var rasio0, rasio1, rasio2 float64
		rasio0 = d0 / r[0]
		rasio1 = d1 / r[1]
		rasio2 = d2 / r[2]
		if math.Abs(rasio0-rasio1) > 1e-5 || math.Abs(rasio0-rasio2) > 1e-5 {
			continue
		}
		if rasio0 < terbaik {
			terbaik = rasio0
			hasilX = p.x
			hasilY = p.y
			ada = true
		}
	}

	if ada {
		fmt.Printf("%.5f %.5f\n", hasilX, hasilY)
	}
}
