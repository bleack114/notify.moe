package pages

import (
	"github.com/aerogo/aero"
	"github.com/aerogo/layout"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/layout"
	"github.com/animenotifier/notify.moe/pages/admin"
	"github.com/animenotifier/notify.moe/pages/anime"
	"github.com/animenotifier/notify.moe/pages/animeimport"
	"github.com/animenotifier/notify.moe/pages/animelist"
	"github.com/animenotifier/notify.moe/pages/animelistitem"
	"github.com/animenotifier/notify.moe/pages/apiview"
	"github.com/animenotifier/notify.moe/pages/calendar"
	"github.com/animenotifier/notify.moe/pages/character"
	"github.com/animenotifier/notify.moe/pages/charge"
	"github.com/animenotifier/notify.moe/pages/companies"
	"github.com/animenotifier/notify.moe/pages/company"
	"github.com/animenotifier/notify.moe/pages/compare"
	"github.com/animenotifier/notify.moe/pages/editanime"
	"github.com/animenotifier/notify.moe/pages/editlog"
	"github.com/animenotifier/notify.moe/pages/editor"
	"github.com/animenotifier/notify.moe/pages/editor/filteranime"
	"github.com/animenotifier/notify.moe/pages/editor/filtercompanies"
	"github.com/animenotifier/notify.moe/pages/editor/filtersoundtracks"
	"github.com/animenotifier/notify.moe/pages/embed"
	"github.com/animenotifier/notify.moe/pages/episode"
	"github.com/animenotifier/notify.moe/pages/explore"
	"github.com/animenotifier/notify.moe/pages/explore/explorecolor"
	"github.com/animenotifier/notify.moe/pages/forum"
	"github.com/animenotifier/notify.moe/pages/genre"
	"github.com/animenotifier/notify.moe/pages/genres"
	"github.com/animenotifier/notify.moe/pages/group"
	"github.com/animenotifier/notify.moe/pages/groups"
	"github.com/animenotifier/notify.moe/pages/home"
	"github.com/animenotifier/notify.moe/pages/inventory"
	"github.com/animenotifier/notify.moe/pages/listimport"
	"github.com/animenotifier/notify.moe/pages/listimport/listimportanilist"
	"github.com/animenotifier/notify.moe/pages/listimport/listimportkitsu"
	"github.com/animenotifier/notify.moe/pages/listimport/listimportmyanimelist"
	"github.com/animenotifier/notify.moe/pages/login"
	"github.com/animenotifier/notify.moe/pages/me"
	"github.com/animenotifier/notify.moe/pages/newthread"
	"github.com/animenotifier/notify.moe/pages/notifications"
	"github.com/animenotifier/notify.moe/pages/paypal"
	"github.com/animenotifier/notify.moe/pages/popular"
	"github.com/animenotifier/notify.moe/pages/posts"
	"github.com/animenotifier/notify.moe/pages/profile"
	"github.com/animenotifier/notify.moe/pages/profile/profilequotes"
	"github.com/animenotifier/notify.moe/pages/profile/profiletracks"
	"github.com/animenotifier/notify.moe/pages/quote"
	"github.com/animenotifier/notify.moe/pages/quotes"
	"github.com/animenotifier/notify.moe/pages/recommended"
	"github.com/animenotifier/notify.moe/pages/search"
	"github.com/animenotifier/notify.moe/pages/settings"
	"github.com/animenotifier/notify.moe/pages/shop"
	"github.com/animenotifier/notify.moe/pages/soundtrack"
	"github.com/animenotifier/notify.moe/pages/soundtracks"
	"github.com/animenotifier/notify.moe/pages/statistics"
	"github.com/animenotifier/notify.moe/pages/support"
	"github.com/animenotifier/notify.moe/pages/terms"
	"github.com/animenotifier/notify.moe/pages/threads"
	"github.com/animenotifier/notify.moe/pages/upload"
	"github.com/animenotifier/notify.moe/pages/user"
	"github.com/animenotifier/notify.moe/pages/users"
)

// Configure registers the page routes in the application.
func Configure(app *aero.Application) {
	l := layout.New(app)

	// Set render function for the layout
	l.Render = fullpage.Render

	// Main menu
	l.Page("/", home.Get)
	l.Page("/explore", explore.Get)
	l.Page("/explore/anime/:year/:status/:type", explore.Filter)
	l.Page("/explore/color/:color/anime", explorecolor.AnimeByAverageColor)
	l.Page("/explore/color/:color/anime/from/:index", explorecolor.AnimeByAverageColor)
	l.Page("/login", login.Get)
	l.Page("/api", apiview.Get)
	// l.Ajax("/dashboard", dashboard.Get)
	// l.Ajax("/best/anime", best.Get)
	// l.Ajax("/artworks", artworks.Get)
	// l.Ajax("/amvs", amvs.Get)

	// Forum
	l.Page("/forum", forum.Get)
	l.Page("/forum/:tag", forum.Get)
	l.Page("/thread/:id", threads.Get)
	l.Page("/post/:id", posts.Get)
	l.Page("/new/thread", newthread.Get)

	// User lists
	l.Page("/users", users.Active)
	l.Page("/users/noavatar", users.ActiveNoAvatar)
	l.Page("/users/games/osu", users.Osu)
	l.Page("/users/games/overwatch", users.Overwatch)
	l.Page("/users/staff", users.Staff)
	l.Page("/users/pro", users.Pro)
	l.Page("/users/editors", users.Editors)

	// Statistics
	l.Page("/statistics", statistics.Get)
	l.Page("/statistics/anime", statistics.Anime)

	// Anime
	l.Page("/anime/:id", anime.Get)
	l.Page("/anime/:id/episodes", anime.Episodes)
	l.Page("/anime/:id/characters", anime.Characters)
	l.Page("/anime/:id/tracks", anime.Tracks)
	l.Page("/anime/:id/episode/:episode-number", episode.Get)

	// Edit anime
	l.Page("/anime/:id/edit", editanime.Main)
	l.Page("/anime/:id/edit/images", editanime.Images)
	l.Page("/anime/:id/edit/characters", editanime.Characters)
	l.Page("/anime/:id/edit/relations", editanime.Relations)
	l.Page("/anime/:id/edit/episodes", editanime.Episodes)
	l.Page("/anime/:id/edit/history", editanime.History)

	// Characters
	l.Page("/character/:id", character.Get)

	// Quotes
	l.Page("/quote/:id", quote.Get)
	l.Page("/quote/:id/edit", quote.Edit)
	l.Page("/quote/:id/history", quote.History)
	l.Page("/quotes", quotes.Latest)
	l.Page("/quotes/from/:index", quotes.Latest)
	l.Page("/quotes/best", quotes.Best)
	l.Page("/quotes/best/from/:index", quotes.Best)

	// Calendar
	l.Page("/calendar", calendar.Get)

	// Companies
	l.Page("/company/:id", company.Get)
	l.Page("/company/:id/edit", company.Edit)
	l.Page("/company/:id/history", company.History)
	l.Page("/companies", companies.Popular)
	l.Page("/companies/from/:index", companies.Popular)
	l.Page("/companies/all", companies.All)

	// Settings
	l.Page("/settings", settings.Get(components.SettingsPersonal))
	l.Page("/settings/accounts", settings.Get(components.SettingsAccounts))
	l.Page("/settings/notifications", settings.Get(components.SettingsNotifications))
	l.Page("/settings/apps", settings.Get(components.SettingsApps))
	l.Page("/settings/formatting", settings.Get(components.SettingsFormatting))
	l.Page("/settings/pro", settings.Get(components.SettingsPro))

	// Soundtracks
	l.Page("/soundtracks", soundtracks.Latest)
	l.Page("/soundtracks/from/:index", soundtracks.Latest)
	l.Page("/soundtracks/best", soundtracks.Best)
	l.Page("/soundtracks/best/from/:index", soundtracks.Best)
	l.Page("/soundtracks/tag/:tag", soundtracks.FilterByTag)
	l.Page("/soundtracks/tag/:tag/from/:index", soundtracks.FilterByTag)
	l.Page("/soundtrack/:id", soundtrack.Get)
	l.Page("/soundtrack/:id/edit", soundtrack.Edit)
	l.Page("/soundtrack/:id/history", soundtrack.History)

	// Groups
	l.Page("/groups", groups.Get)
	l.Page("/group/:id", group.Get)
	l.Page("/group/:id/edit", group.Edit)
	l.Page("/group/:id/forum", group.Forum)

	// Notifications
	l.Page("/notifications", notifications.ByUser)
	l.Page("/notifications/all", notifications.All)

	// User profiles
	l.Page("/user", user.Get)
	l.Page("/user/:nick", profile.Get)
	l.Page("/user/:nick/forum/threads", profile.GetThreadsByUser)
	l.Page("/user/:nick/forum/posts", profile.GetPostsByUser)
	l.Page("/user/:nick/soundtracks/added", profiletracks.Added)
	l.Page("/user/:nick/soundtracks/added/from/:index", profiletracks.Added)
	l.Page("/user/:nick/soundtracks/liked", profiletracks.Liked)
	l.Page("/user/:nick/soundtracks/liked/from/:index", profiletracks.Liked)
	l.Page("/user/:nick/quotes/added", profilequotes.Added)
	l.Page("/user/:nick/quotes/added/from/:index", profilequotes.Added)
	l.Page("/user/:nick/quotes/liked", profilequotes.Liked)
	l.Page("/user/:nick/quotes/liked/from/:index", profilequotes.Liked)
	l.Page("/user/:nick/stats", profile.GetStatsByUser)
	l.Page("/user/:nick/followers", profile.GetFollowers)
	l.Page("/user/:nick/animelist/anime/:id", animelistitem.Get)
	l.Page("/user/:nick/recommended/anime", recommended.Anime)
	l.Page("/user/:nick/notifications", notifications.ByUser)

	// Anime list
	l.Page("/user/:nick/animelist/watching", animelist.FilterByStatus(arn.AnimeListStatusWatching))
	l.Page("/user/:nick/animelist/completed", animelist.FilterByStatus(arn.AnimeListStatusCompleted))
	l.Page("/user/:nick/animelist/planned", animelist.FilterByStatus(arn.AnimeListStatusPlanned))
	l.Page("/user/:nick/animelist/hold", animelist.FilterByStatus(arn.AnimeListStatusHold))
	l.Page("/user/:nick/animelist/dropped", animelist.FilterByStatus(arn.AnimeListStatusDropped))

	l.Page("/user/:nick/animelist/watching/from/:index", animelist.FilterByStatus(arn.AnimeListStatusWatching))
	l.Page("/user/:nick/animelist/completed/from/:index", animelist.FilterByStatus(arn.AnimeListStatusCompleted))
	l.Page("/user/:nick/animelist/planned/from/:index", animelist.FilterByStatus(arn.AnimeListStatusPlanned))
	l.Page("/user/:nick/animelist/hold/from/:index", animelist.FilterByStatus(arn.AnimeListStatusHold))
	l.Page("/user/:nick/animelist/dropped/from/:index", animelist.FilterByStatus(arn.AnimeListStatusDropped))

	l.Page("/animelist/watching", animelist.Redirect)
	l.Page("/animelist/completed", animelist.Redirect)
	l.Page("/animelist/planned", animelist.Redirect)
	l.Page("/animelist/hold", animelist.Redirect)
	l.Page("/animelist/dropped", animelist.Redirect)

	// Compare
	l.Page("/compare/animelist/:nick-1/:nick-2", compare.AnimeList)

	// Search
	l.Page("/search/*term", search.Get)
	l.Page("/empty-search", search.GetEmptySearch)
	l.Page("/anime-search/*term", search.Anime)
	l.Page("/character-search/*term", search.Characters)
	l.Page("/forum-search/*term", search.Forum)
	l.Page("/soundtrack-search/*term", search.SoundTracks)
	l.Page("/user-search/*term", search.Users)
	l.Page("/company-search/*term", search.Companies)

	// Shop
	l.Page("/support", support.Get)
	l.Page("/shop", shop.Get)
	l.Page("/inventory", inventory.Get)
	l.Page("/charge", charge.Get)
	l.Page("/shop/history", shop.PurchaseHistory)
	app.Post("/api/shop/buy/:item/:quantity", shop.BuyItem)

	// Import anime
	app.Post("/api/import/kitsu/anime/:id", animeimport.Kitsu)

	// Upload
	app.Post("/api/upload/avatar", upload.Avatar)
	app.Post("/api/upload/cover", upload.Cover)
	app.Post("/api/upload/anime/:id/image", upload.AnimeImage)

	// Admin
	l.Page("/admin", admin.Get)
	l.Page("/admin/webdev", admin.WebDev)
	l.Page("/admin/purchases", admin.PurchaseHistory)

	// Editor
	l.Page("/editor", editor.Get)

	// Editor links can be filtered by year, status and type
	editorFilterable := func(route string, handler func(ctx *aero.Context) string) {
		l.Page(route+"/:year/:status/:type", handler)
	}

	// Editor - Anime
	editorFilterable("/editor/anime/synopsis", filteranime.Synopsis)
	editorFilterable("/editor/anime/genres", filteranime.Genres)
	editorFilterable("/editor/anime/startdate", filteranime.StartDate)
	editorFilterable("/editor/anime/mapping/shoboi", filteranime.Shoboi)
	editorFilterable("/editor/anime/mapping/anilist", filteranime.AniList)
	editorFilterable("/editor/anime/mapping/mal", filteranime.MAL)
	editorFilterable("/editor/anime/mapping/duplicate", filteranime.DuplicateMappings)
	editorFilterable("/editor/anime/image/lowres", filteranime.LowResolutionAnimeImages)
	editorFilterable("/editor/anime/image/ultralowres", filteranime.UltraLowResolutionAnimeImages)

	// Editor - MALdiff
	editorFilterable("/editor/mal/diff/anime", editor.CompareMAL)

	// Editor - Kitsu
	l.Page("/editor/kitsu/new/anime", editor.NewKitsuAnime)

	// Editor - Companies
	l.Page("/editor/companies/description", filtercompanies.NoDescription)

	// Editor - Soundtracks
	l.Page("/editor/soundtracks/links", filtersoundtracks.NoLinks)

	// Log
	l.Page("/log", editlog.Get)

	// Mixed
	// l.Page("/database", database.Get)
	// app.Get("/api/select/:data-type/where/:field/is/:field-value", database.Select)

	// Import
	l.Page("/import", listimport.Get)
	l.Page("/import/anilist/animelist", listimportanilist.Preview)
	l.Page("/import/anilist/animelist/finish", listimportanilist.Finish)
	l.Page("/import/myanimelist/animelist", listimportmyanimelist.Preview)
	l.Page("/import/myanimelist/animelist/finish", listimportmyanimelist.Finish)
	l.Page("/import/kitsu/animelist", listimportkitsu.Preview)
	l.Page("/import/kitsu/animelist/finish", listimportkitsu.Finish)

	// Browser extension
	l.Page("/extension/embed", embed.Get)

	// API
	app.Get("/api/me", me.Get)
	app.Get("/api/popular/anime/titles/:count", popular.AnimeTitles)
	app.Get("/api/test/notification", notifications.Test)
	app.Get("/api/count/notifications/unseen", notifications.CountUnseen)
	app.Get("/api/mark/notifications/seen", notifications.MarkNotificationsAsSeen)
	app.Get("/api/random/soundtrack", soundtrack.Random)
	app.Get("/api/next/soundtrack", soundtrack.Next)

	// Legal stuff
	l.Page("/terms", terms.Get)

	// PayPal
	l.Page("/paypal/success", paypal.Success)
	l.Page("/paypal/cancel", paypal.Cancel)
	app.Post("/api/paypal/payment/create", paypal.CreatePayment)

	// Genres
	l.Page("/genres", genres.Get)
	l.Page("/genre/:name", genre.Get)
}
