package ports

type ArithmeticPort interface {
	Addition(numberOne int32, numberTwo int32) (int32, error)
	Subtraction(numberOne int32, numberTwo int32) (int32, error)
	Multiplication(numberOne int32, numberTwo int32) (int32, error)
	Division(numberOne int32, numberTwo int32) (int32, error)
}
