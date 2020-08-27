package Graph

type WeightGraph interface {
	Graph
	Weight(edge Edge) int                  //返回边长的权重
	AddWeightEdge(edge Edge, weight int)   // 增加有向图的权重
	AddWeightEdgeBi(edge Edge, weight int) // 增加无向图的权重
	TotalWeight() int                      // 统计权重的总量
}

type AdjacencyMatrixWithWeight struct {
	AdjacencyMatrix
	weights     map[Edge]int
	totalWeight int
}

func NewAdjacencyMatrixWithWeight() WeightGraph {
	return new(AdjacencyMatrixWithWeight).Init()
}

func (g *AdjacencyMatrixWithWeight) Init() WeightGraph {
	g.AdjacencyMatrix.init()       // 初始化图
	g.weights = make(map[Edge]int) // 权重
	g.totalWeight = 0              // 总重量
	return g
}

func (g *AdjacencyMatrixWithWeight) Weight(edge Edge) int {
	if value, ok := g.weights[edge]; ok {
		return value
	}
	return -1
}
func (g *AdjacencyMatrixWithWeight) AddWeightEdge(edge Edge, weight int) {
	g.AdjacencyMatrix.AddEdge(edge)
	if _, ok := g.weights[edge]; ok {
		g.totalWeight = g.totalWeight - g.weights[edge] + weight
	} else {
		g.totalWeight = g.totalWeight + weight
	}
	g.weights[edge] = weight
}
func (g *AdjacencyMatrixWithWeight) AddWeightEdgeBi(edge Edge, weight int) {
	g.AddWeightEdge(edge, weight)
	g.AddWeightEdge(Edge{edge.End, edge.Start}, weight)
}

func (g *AdjacencyMatrixWithWeight) TotalWeight() int {
	return g.totalWeight
}

// 继承的覆盖
func (g *AdjacencyMatrixWithWeight) DeleteEdge(edge Edge) {
	g.AdjacencyMatrix.DeleteEdge(edge)
	if w, ok := g.weights[edge]; ok {
		g.totalWeight -= w
		delete(g.weights, edge) // 删除权重
	}
}

// 删除
func (g *AdjacencyMatrixWithWeight) DeleteEdgeBi(edge Edge) {
	g.DeleteEdge(edge)
	g.DeleteEdge(Edge{edge.End, edge.Start})
}
