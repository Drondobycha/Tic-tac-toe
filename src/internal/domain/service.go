package domain

type GameService interface {
	GetNextMove(field GameField) (row, col int, err error)
	ValidateMove(savedField, receivedField GameField) error
	CheckGameStatus(field GameField) GameStatus
}
