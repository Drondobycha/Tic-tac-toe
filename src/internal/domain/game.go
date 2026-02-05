package domain

import "github.com/google/uuid"

type GameField [3][3]int

const (
	Empty = 0
	X     = 1
	O     = 2
)

type GameStatus int

const (
	StatusInProgress GameStatus = iota
	StatusXWon
	StatusOWon
	StatusDraw
)

type Game struct {
	ID     uuid.UUID
	Field  GameField
	Status GameStatus
}
