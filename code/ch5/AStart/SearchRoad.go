package main

type SearchRoad struct {
	theMap  *Map                   // 地图
	start   AstarPoint             // 开始
	end     AstarPoint             // 结束
	closeLi map[string]*AstarPoint // 关闭
	openLi  map[string]*AstarPoint // 不通的路
	openSet map[string]*AstarPoint // 通路
	TheRoad []*AstarPoint          // 通道
}
