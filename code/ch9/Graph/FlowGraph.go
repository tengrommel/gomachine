package Graph

type FlowGraph interface {
	Graph
	Cap(edge Edge) int                   // 最大宽度
	Flow(edge Edge) int                  // 当前流量宽度
	AddEdgeWithCap(edge Edge, data int)  // 根据宽度增加边长
	AddEdgeWithFlow(edge Edge, data int) // 根据流量增加边长
}

func newAdjacencyMatrixWithFlow() FlowGraph {
	return new(AdjacencyMatrixWithFlow).init()
}

func (g *AdjacencyMatrixWithFlow) init() FlowGraph {
	g.AdjacencyMatrix.init()
	g.cap = make(map[Edge]int)
	g.flow = make(map[Edge]int)
	return g
}

type AdjacencyMatrixWithFlow struct {
	AdjacencyMatrix              // 矩阵
	cap             map[Edge]int // 边长的最大宽度
	flow            map[Edge]int // 边长的当前流量
}

func (g *AdjacencyMatrixWithFlow) Cap(edge Edge) int {
	if _, ok := g.cap[edge]; !ok {
		return 0
	}
	return g.cap[edge]
}

func (g *AdjacencyMatrixWithFlow) Flow(edge Edge) int {
	if _, ok := g.flow[edge]; !ok {
		return 0
	}
	return g.flow[edge]
}

func (g *AdjacencyMatrixWithFlow) AddEdgeWithCap(edge Edge, data int) {
	g.AdjacencyMatrix.AddEdge(edge)
	g.cap[edge] = data
}

func (g *AdjacencyMatrixWithFlow) AddEdgeWithFlow(edge Edge, data int) {
	g.AdjacencyMatrix.AddEdge(edge)
	g.cap[edge] = data
}

// 覆盖重写
func (g *AdjacencyMatrixWithFlow) DeleteEdge(edge Edge) {
	g.AdjacencyMatrix.DeleteEdge(edge)
	delete(g.cap, edge)
	delete(g.flow, edge)
}

func (g *AdjacencyMatrixWithFlow) AddEdgeBi(edge Edge) {
	panic("WithFlow没有无向图操作")
}

func (g *AdjacencyMatrixWithFlow) DeleteEdgeBi(edge Edge) {
	panic("WithFlow没有无向图操作")
}
