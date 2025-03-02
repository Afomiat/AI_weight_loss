package usecase

import "github.com/Afomiat/AI_weight_loss/backend/internal/ai"

type ExerciseUsecase struct{}

func NewExerciseUsecase() *ExerciseUsecase {
	return &ExerciseUsecase{}
}

func (uc *ExerciseUsecase) GetRecommendation(goal string) (string, error) {
	return ai.GetExerciseRecommendation(goal)
}
