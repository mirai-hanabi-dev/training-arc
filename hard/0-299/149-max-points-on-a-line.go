package main

type Point struct {
	x, y int
}

func maxPoints(points [][]int) int {
	maxLine := 0
	for i := 0; i < len(points)-1; i++ {
		vectors := make(map[Point]int)
		p1 := Point{
			x: points[i][0],
			y: points[i][1],
		}
		for j := i + 1; j < len(points); j++ {
			p2 := Point{
				x: points[j][0],
				y: points[j][1],
			}
			vector := Point{
				x: p2.x - p1.x,
				y: p2.y - p1.y,
			}
			g := gcd(vector.x, vector.y)
			if g != 0 {
				vector.x /= g
				vector.y /= g
			}
			if vector.x < 0 {
				vector.x = -vector.x
				vector.y = -vector.y
			}
			vectors[vector] += 1
			if vectors[vector] > maxLine {
				maxLine = vectors[vector]
			}
		}
	}

	return maxLine + 1
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
