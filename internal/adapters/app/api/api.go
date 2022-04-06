package api

import "hexArch/internal/ports"

type Adapter struct {
	dbPort         ports.DbPort
	arithmeticPort ports.ArithmeticPort
}

func NewAdapter(dbPort ports.DbPort, arithmeticPort ports.ArithmeticPort) *Adapter {
	return &Adapter{dbPort: dbPort, arithmeticPort: arithmeticPort}
}
func (apiAdapter Adapter) GetAddition(numberOne, numberTwo int32) (int32, error) {
	result, err := apiAdapter.arithmeticPort.Addition(numberOne, numberTwo)

	if err != nil {
		return 0, err
	}
	err = apiAdapter.dbPort.AddToHistory(result, "addition")
	if err != nil {
		return 0, nil
	}
	return result, nil
}
func (apiAdapter Adapter) GetSubtraction(numberOne, numberTwo int32) (int32, error) {
	result, err := apiAdapter.arithmeticPort.Subtraction(numberOne, numberTwo)

	if err != nil {
		return 0, err
	}
	err = apiAdapter.dbPort.AddToHistory(result, "subtraction")
	if err != nil {
		return 0, nil
	}
	return result, nil
}
func (apiAdapter Adapter) GetMultiplication(numberOne, numberTwo int32) (int32, error) {
	result, err := apiAdapter.arithmeticPort.Multiplication(numberOne, numberTwo)

	if err != nil {
		return 0, err
	}
	err = apiAdapter.dbPort.AddToHistory(result, "multiplication")
	if err != nil {
		return 0, nil
	}
	return result, nil
}
func (apiAdapter Adapter) GetDivision(numberOne, numberTwo int32) (int32, error) {
	result, err := apiAdapter.arithmeticPort.Division(numberOne, numberTwo)

	if err != nil {
		return 0, err
	}
	err = apiAdapter.dbPort.AddToHistory(result, "division")
	if err != nil {
		return 0, nil
	}
	return result, nil
}
