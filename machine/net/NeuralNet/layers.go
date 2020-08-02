package NeuralNet

// 神经网络的层

type Layer interface {
	Forward(Tensor) Tensor         // 前进
	Backward(Tensor) Tensor        // 后退
	GetWeights() map[string]Tensor // 抓取权重
	GetGrads() map[string]Tensor   // 抓取损失
}

type Linear struct {
	inputs  map[string]Tensor
	grads   map[string]Tensor
	weights map[string]Tensor
}

func NewLinear(inputsSize TensorShape, outSize TensorShape) Linear {
	l := Linear{
		inputs:  make(map[string]Tensor),
		grads:   make(map[string]Tensor),
		weights: make(map[string]Tensor),
	}
	return l
}

type Relu struct {
	inputs  map[string]Tensor
	grads   map[string]Tensor
	weights map[string]Tensor
}

func NewRelu(inputSize TensorShape, outSize TensorShape) Relu {
	l := Relu{
		inputs:  make(map[string]Tensor),
		grads:   make(map[string]Tensor),
		weights: make(map[string]Tensor),
	}
	return l
}

func (l Relu) GetWeights() map[string]Tensor {
	return l.weights
}
