package main

type GameModel struct {
	Level                        int
	NumLines                     int
	Score                        int
	targetNumberOfLinesToLevelUp int
	IsGameOver                   bool
}

const linesPerLevel = 10

var scoreCoefficients = [4]int{40, 100, 300, 1200}

func (model *GameModel) updateGameModel(clearedLines int) {
	if clearedLines == 0 {
		return
	}
	model.Score += scoreCoefficients[clearedLines-1] * (model.Level + 1)
	model.NumLines += clearedLines
	if model.NumLines >= model.targetNumberOfLinesToLevelUp {
		model.targetNumberOfLinesToLevelUp += linesPerLevel
		model.Level++
	}
}
