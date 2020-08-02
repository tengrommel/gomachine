package NeuralNet

import "math/rand"

// 张量

type TensorShape [2]int

// 存储行于列
func NewTensorShape(row int, cols int) TensorShape {
	return TensorShape{
		row,
		cols,
	}
}

// 实现2维张量
type Tensor struct {
	Data  [][]float64 // 数组
	Shape TensorShape // 数组的维度
}

func NewTensor(Shape TensorShape) Tensor {
	data := make([][]float64, Shape[0], Shape[0]) // 二维数组
	for i := range data {
		data[i] = make([]float64, Shape[1], Shape[1])
	}
	t := Tensor{
		Data:  data,
		Shape: Shape,
	}
	return t
}

// 新建张量,按照某个数据初始化
func NewTensorFromFloat64(Shape TensorShape, initValue float64) Tensor {
	t := NewTensor(Shape)
	for i := 0; i < Shape[0]; i++ {
		for j := 0; j < Shape[1]; j++ {
			t.Data[i][j] = initValue
		}
	}
	return t
}

func NewTensorRandomNormal(Shape TensorShape) Tensor {
	t := NewTensor(Shape)
	for i := 0; i < Shape[0]; i++ {
		for j := 0; j < Shape[1]; j++ {
			t.Data[i][j] = rand.NormFloat64()
		}
	}
	return t
}

// 张量的乘法
func DotFloat64(x Tensor, y Tensor) Tensor {
	out := NewTensor(NewTensorShape(x.Shape[0], y.Shape[1]))
	for i := 0; i < x.Shape[0]; i++ {
		for j := 0; j < y.Shape[1]; j++ {
			o := 0.0
			for k := 0; k < x.Shape[1]; k++ {
				o += x.Data[j][k] * y.Data[k][j]
			}
			out.Data[i][j] = o
		}
	}
	return out
}

// 张量的加法
func AddFloat64(x Tensor, y Tensor) Tensor {
	out := x
	for i := 0; i < x.Shape[0]; i++ {
		for j := 0; j < x.Shape[1]; j++ {
			out.Data[i][j] += y.Data[i][j]
		}
	}
	return out
}

// 张量的转置
func TransposeFloat64(x Tensor) Tensor {
	out := NewTensor(NewTensorShape(x.Shape[1], x.Shape[0]))
	for i := 0; i < x.Shape[0]; i++ {
		for j := 0; j < x.Shape[1]; j++ {
			out.Data[j][i] = x.Data[i][j]
		}
	}
	return out
}

// 每一列统计到第一行
func SumAcross0Float64(x Tensor) Tensor {
	out := NewTensor(NewTensorShape(1, x.Shape[1]))
	for i := 0; i < x.Shape[0]; i++ {
		for j := 0; j < x.Shape[1]; j++ {
			out.Data[0][j] += x.Data[i][j]
		}
	}
	return out
}

// 取出两个张量最大值 合成一个张量
func MaxFloat64(x, y Tensor) Tensor {
	out := NewTensor(x.Shape)
	for i := 0; i < x.Shape[0]; i++ {
		for j := 0; j < x.Shape[1]; j++ {
			out.Data[i][j] = y.Data[i][j]
			if x.Data[i][j] > y.Data[i][j] {
				out.Data[i][j] = x.Data[i][j]
			}
		}
	}
	return out
}

// 根据x保存y
func WhereFloat64(x, y Tensor) Tensor {
	out := NewTensor(x.Shape)
	for i := 0; i < x.Shape[0]; i++ {
		for j := 0; j < x.Shape[1]; j++ {
			if x.Data[i][j] > 0 {
				out.Data[i][j] = y.Data[i][j]
			}
		}
	}
	return out
}

func MulFloat64(x, y Tensor) Tensor {
	out := NewTensor(x.Shape)
	for i := 0; i < x.Shape[0]; i++ {
		for j := 0; j < x.Shape[1]; j++ {
			out.Data[i][j] = x.Data[i][j] * y.Data[i][j]
		}
	}
	return out
}
