package usecase

import (
	"github.com/Afomiat/AI_weight_loss/backend/internal/ai"
	"github.com/Afomiat/AI_weight_loss/backend/domain"
	"github.com/Afomiat/AI_weight_loss/backend/repository"
)

type ExerciseUsecase struct{}

func NewExerciseUsecase() *ExerciseUsecase {
	return &ExerciseUsecase{}
}

func (uc *ExerciseUsecase) GetRecommendation(goal string) (string, error) {
	// Get AI recommendation
	advice, err := ai.GetExerciseRecommendation(goal)
	if err != nil {
		return "", err
	}

	// Store in repository before returning
	exercise := domain.Exercise{
		Goal:   goal,
		Advice: advice,
	}
	if err := repository.AddExercise(exercise); err != nil {
		return "", err
	}

	return advice, nil
}