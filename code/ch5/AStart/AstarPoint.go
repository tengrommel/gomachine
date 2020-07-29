package main

// A 星算法地图上点的结构
type AstarPoint struct {
	Point
	father *AstarPoint
	gVal   int
	hVal   int
	fVal   int
}
