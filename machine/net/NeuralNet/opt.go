package NeuralNet

// 优化
type Optimezer interface {
	Step()
}

// 优化神经网络
type SGD struct {
	LR  float64
	Net NeuralNet
}

func (sgd SGD) Step() {
	// 循环每层神经网络
	for _, layer := range sgd.Net.GetLayers() {
		weights := layer.GetWeights()
		grads := layer.GetGrads()
		for name, weight := range weights {
			grad := grads[name]
			gradScaled := MulFloat64(NewTensorFromFloat64(grad.Shape, sgd.LR), grad)
			weights[name] = SubFloat64(weight, gradScaled) // 权重等于差距
		}
	}
}
