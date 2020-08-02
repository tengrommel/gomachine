package NeuralNet

//神经网络的核心
type NeuralNet interface {
	Forward(inputs Tensor) Tensor // 正向
	Backward(grad Tensor)         // 后退
	GetLayers() []Layer           // 获取多层神经网络
}

type Sequential struct {
	layers []Layer
}

// 循环前进
func (seq Sequential) Forward(inputs Tensor) Tensor {
	for _, layer := range seq.layers {
		inputs = layer.Forward(inputs)
	}
	return inputs
}

// 反向神经网络
func (seq Sequential) Backward(grad Tensor) {
	revLayers := make([]Layer, len(seq.layers))
	for i := 0; i < len(revLayers); i++ {
		revLayers[i] = seq.layers[len(seq.layers)-i-1]
	}
	for _, layer := range revLayers {
		grad = layer.Backward(grad)
	}
}

func (seq Sequential) GetLayers() []Layer {
	return seq.layers
}

func NewSequential(layers []Layer) Sequential {
	seq := Sequential{layers: layers}
	return seq
}
