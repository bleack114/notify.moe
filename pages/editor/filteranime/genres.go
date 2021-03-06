package filteranime

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
)

const maxGenreEntries = 70

// Genres ...
func Genres(ctx *aero.Context) string {
	return editorList(
		ctx,
		"Anime without genres",
		func(anime *arn.Anime) bool {
			return len(anime.Genres) == 0
		},
		nil,
	)
}
