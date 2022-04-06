package ports

type ArithmeticPort interface {
	Addition(number1 int32, number2 int32) (int32, error)
	Subtraction(number1 int32, number2 int32) (int32, error)
	Multiplication(number1 int32, number2 int32) (int32, error)
	Division(number1 int32, number2 int32) (int32, error)
}
