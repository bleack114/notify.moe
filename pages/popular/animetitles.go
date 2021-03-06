package popular

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
)

// AnimeTitles returns a list of the 500 most popular anime titles.
func AnimeTitles(ctx *aero.Context) string {
	maxLength, err := ctx.GetInt("count")

	if err != nil {
		return ctx.Error(http.StatusBadRequest, "Invalid value for count parameter", err)
	}

	popularAnimeTitles := []string{}
	popularAnime := arn.AllAnime()
	arn.SortAnimeByPopularity(popularAnime)

	if len(popularAnime) > maxLength {
		popularAnime = popularAnime[:maxLength]
	}

	for _, anime := range popularAnime {
		popularAnimeTitles = append(popularAnimeTitles, anime.Title.Canonical)

		if arn.ContainsUnicodeLetters(anime.Title.Japanese) {
			popularAnimeTitles = append(popularAnimeTitles, anime.Title.Japanese)
		}
	}

	return ctx.JSON(popularAnimeTitles)
}
