package main

import (
	"fmt"
	"github.com/kuroko1t/GoMNIST"
	"github.com/kuroko1t/gdeep"
	"github.com/kuroko1t/gmat"
)

// 损失函数

func main() {
	train, test, err := GoMNIST.Load("./minst") // 读取训练数据,测试数据
	if err != nil {
		fmt.Println(err)
	}
	trainDataSize := len(train.ImagesFloatNorm) // 训练数据的大小
	testDataSize := len(test.ImagesFloatNorm)   // 测试数据大小
	fmt.Println(testDataSize)

	batchSize := 128                                     // 批量大小
	inputSize := 784                                     // 输入大小
	hiddenSize := 20                                     // 隐藏大小
	outputSize := 10                                     // 输出结果
	learningRate := 0.01                                 // 学习的差异
	epochNum := 1                                        // 训练次数
	iterationNum := trainDataSize * epochNum / batchSize // 迭代次数

	// 构造神经网络
	// 神经网络输出
	dropout1 := &gdeep.Dropout{}
	dropout2 := &gdeep.Dropout{}

	// 中间层神经网络
	layer := []gdeep.LayerInterface{}                                    // Dense稠密
	gdeep.LayerAdd(&layer, &gdeep.Dense{}, []int{inputSize, hiddenSize}) // 第一层输入
	gdeep.LayerAdd(&layer, &gdeep.Relu{})
	gdeep.LayerAdd(&layer, dropout1, 0.2)

	gdeep.LayerAdd(&layer, &gdeep.Dense{}, []int{hiddenSize, hiddenSize}) // 第二层输入
	gdeep.LayerAdd(&layer, &gdeep.Relu{})
	gdeep.LayerAdd(&layer, dropout2, 0.2)

	gdeep.LayerAdd(&layer, &gdeep.Dense{}, []int{hiddenSize, outputSize})
	gdeep.LayerAdd(&layer, &gdeep.SoftmaxWithLoss{})                           // 损失最低
	momentum := &gdeep.Momentum{LearningRate: learningRate, MomentumRate: 0.9} // 训练步骤

	iter := 0
	for i := 0; i < iterationNum; i++ {
		if (i+2)*batchSize > trainDataSize {
			iter = 0 // 归零
		}
		// 图片数组
		imageBatch := train.ImagesFloatNorm[:][iter*batchSize : (iter+1)*batchSize] // 图片数组
		// 标签数组
		labelBatch := train.LabelsOneHot[:][iter*batchSize : (iter+1)*batchSize]
		x := gmat.Make2DInitArray(imageBatch)    // 训练用的二维数组
		t := gmat.Make2DInitArray(labelBatch)    // 结果的数组
		loss := gdeep.Run(layer, momentum, x, t) // 运行神经网络
		gdeep.AvePrint(loss, "loss")
		iter++
	}
}
