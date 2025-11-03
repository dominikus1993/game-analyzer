package services

import (
	"context"

	"github.com/dominikus1993/game-analyzer/pkg/model"
)

type GameDataSourceService interface {
	ProvideGames(ctx context.Context) (model.GamesStream, error)
}
