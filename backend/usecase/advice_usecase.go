package usecase

import (
    "github.com/Afomiat/AI_weight_loss/backend/internal/ai"
)

func GetAdvice() (string, error) {
    return ai.GetFitnessAdvice()
}
