package Graph

const (
	while = 0 // 图的类型，无向图，有向图，混合图
	gray  = 1
	block = 2
)

type Edge struct {
	Start, End interface{} // 开始点 结束点
}

type Graph interface {
	AddVertex(interface{})        // 增加顶点
	CheckVertex(interface{}) bool // 判断顶点是否存在
	DeleteVertex(interface{})     // 删除顶点

	AddEdge(edge Edge)        // 增加边长
	CheckEdge(edge Edge) bool // 判断边长是否流通
	DeleteEdge(edge Edge)     // 删除边长

	AddEdgeBi(edge Edge)    // 增加边长，处理无向
	DeleteEdgeBi(edge Edge) // 删除无向的边长

	AllVertices() []interface{} // 返回所有的顶点
	AllEdges() []Edge           // 返回所有的边长

	AllConnectedVertices() []interface{} // 返回所有的联通的顶点
	AllConnectedEdges() []Edge           // 返回所有的联通的边长
	Transpose() Graph                    // 返回旋转图
}

// 邻接矩阵 图的基本结构
type AdjacencyMatrix struct {
	matrix LinkedMap
	Graph
}
