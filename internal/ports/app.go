package ports

type APIPort interface {
	GetAddition(numberOne, numberTwo int32) (int32, error)
	GetSubtraction(numberOne, numberTwo int32) (int32, error)
	GetMultiplication(numberOne, numberTwo int32) (int32, error)
	GetDivision(numberOne, numberTwo int32) (int32, error)
}
