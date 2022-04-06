package arithmetic

type Adapter struct {
}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (arithmeticAdapter Adapter) Addition(operandOne int32, operandTwo int32) (int32, error) {
	return operandOne + operandTwo, nil
}

func (artihmeticAdapter Adapter) Subtraction(operandOne int32, operandTwo int32) (int32, error) {
	return operandOne - operandTwo, nil
}
func (arithmeticAdapter Adapter) Multiplication(operandOne int32, operandTwo int32) (int32, error) {
	return operandOne * operandTwo, nil
}

func (artihmeticAdapter Adapter) Division(operandOne int32, operandTwo int32) (int32, error) {
	return operandOne / operandTwo, nil
}
