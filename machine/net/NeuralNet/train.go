package NeuralNet

import "fmt"

// 训练

func Train(net NeuralNet, inputs Tensor, targets Tensor, numEpochs int, lr float64) {
	loss := MSE{}
	opt := SGD{lr, net}
	for epoch := 0; epoch < numEpochs; epoch++ {
		epochLoss := 0.0
		predicted := net.Forward(inputs) // 输入
		epochLoss += loss.Loss(predicted, targets)
		grad := loss.Grad(predicted, targets)
		net.Backward(grad) // 反向
		opt.Step()         // 逐步执行
		if epoch%10 == 0 {
			fmt.Printf("Epoch/loss:%v\t%f \n", epoch, epochLoss)
		}
	}
}
