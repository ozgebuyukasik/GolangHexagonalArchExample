package ports

type DbPort interface {
	CloseDbConnection()
	AddToHistory(result int32, operation string) error
}
