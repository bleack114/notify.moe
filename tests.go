package main

var routeTests = map[string][]string{
	// User
	"/user/:nick": []string{
		"/+Akyoto",
	},

	"/user/:nick/forum/threads": []string{
		"/+Akyoto/forum/threads",
	},

	"/user/:nick/forum/posts": []string{
		"/+Akyoto/forum/posts",
	},

	"/user/:nick/soundtracks/added": []string{
		"/+Akyoto/soundtracks/added",
	},

	"/user/:nick/soundtracks/added/from/:index": []string{
		"/+Akyoto/soundtracks/added/from/3",
	},

	"/user/:nick/soundtracks/liked": []string{
		"/+Akyoto/soundtracks/liked",
	},

	"/user/:nick/soundtracks/liked/from/:index": []string{
		"/+Akyoto/soundtracks/liked/from/3",
	},

	"/user/:nick/quotes/added": []string{
		"/+Scott/quotes/added",
	},

	"/user/:nick/quotes/added/from/:index": []string{
		"/+Scott/quotes/added/from/3",
	},

	"/user/:nick/quotes/liked": []string{
		"/+Scott/quotes/liked",
	},

	"/user/:nick/quotes/liked/from/:index": []string{
		"/+Scott/quotes/liked/from/3",
	},

	"/user/:nick/followers": []string{
		"/+Akyoto/followers",
	},

	"/user/:nick/stats": []string{
		"/+Akyoto/stats",
	},

	"/user/:nick/animelist/anime/:id": []string{
		"/+Akyoto/animelist/anime/7929",
	},

	"/user/:nick/animelist/watching": []string{
		"/+Akyoto/animelist/watching",
	},

	"/user/:nick/animelist/watching/from/:index": []string{
		"/+Akyoto/animelist/watching/from/3",
	},

	"/user/:nick/animelist/completed": []string{
		"/+Akyoto/animelist/completed",
	},

	"/user/:nick/animelist/completed/from/:index": []string{
		"/+Akyoto/animelist/completed/from/3",
	},

	"/user/:nick/animelist/planned": []string{
		"/+Akyoto/animelist/planned",
	},

	"/user/:nick/animelist/planned/from/:index": []string{
		"/+Akyoto/animelist/planned/from/3",
	},

	"/user/:nick/animelist/hold": []string{
		"/+Akyoto/animelist/hold",
	},

	"/user/:nick/animelist/hold/from/:index": []string{
		"/+Akyoto/animelist/hold/from/3",
	},

	"/user/:nick/animelist/dropped": []string{
		"/+Akyoto/animelist/dropped",
	},

	"/user/:nick/animelist/dropped/from/:index": []string{
		"/+Akyoto/animelist/dropped/from/3",
	},

	"/user/:nick/recommended/anime": []string{
		"/+Akyoto/recommended/anime",
	},

	// Pages
	"/anime/:id": []string{
		"/anime/1",
	},

	"/anime/:id/characters": []string{
		"/anime/1/characters",
	},

	"/anime/:id/episodes": []string{
		"/anime/1/episodes",
	},

	"/anime/:id/tracks": []string{
		"/anime/1/tracks",
	},

	"/thread/:id": []string{
		"/thread/HJgS7c2K",
	},

	"/post/:id": []string{
		"/post/B1RzshnK",
	},

	"/forum/:tag": []string{
		"/forum/general",
	},

	"/search/:term": []string{
		"/search/Dragon Ball",
	},

	"/quote/:id": []string{
		"/quote/gUZugd6zR",
	},

	"/quote/:id/edit": []string{
		"/quote/gUZugd6zR/edit",
	},

	"/quotes/from/:index": []string{
		"/quotes/from/2",
	},

	"/quotes/best/from/:index": []string{
		"/quotes/best/from/2",
	},

	"/soundtrack/:id": []string{
		"/soundtrack/h0ac8sKkg",
	},

	"/soundtrack/:id/edit": []string{
		"/soundtrack/h0ac8sKkg/edit",
	},

	"/soundtrack/:id/history": []string{
		"/soundtrack/h0ac8sKkg/history",
	},

	"/soundtracks": []string{
		"/soundtracks",
	},

	"/soundtracks/from/:index": []string{
		"/soundtracks/from/12",
	},

	"/soundtracks/best": []string{
		"/soundtracks/best",
	},

	"/soundtracks/best/from/:index": []string{
		"/soundtracks/best/from/12",
	},

	"/soundtracks/tag/:tag": []string{
		"/soundtracks/tag/moe",
	},

	"/soundtracks/tag/:tag/from/:index": []string{
		"/soundtracks/tag/moe/from/3",
	},

	"/character/:id": []string{
		"/character/6556",
	},

	"/compare/animelist/:nick-1/:nick-2": []string{
		"/compare/animelist/Akyoto/Scott",
	},

	"/explore/anime/:year/:status/:type": []string{
		"/explore/anime/2011/finished/tv",
	},

	// API
	"/api/anime/:id": []string{
		"/api/anime/1",
	},

	"/api/thread/:id": []string{
		"/api/thread/HJgS7c2K",
	},

	"/api/post/:id": []string{
		"/api/post/B1RzshnK",
	},

	"/api/animelist/:id": []string{
		"/api/animelist/4J6qpK1ve",
	},

	"/api/settings/:id": []string{
		"/api/settings/4J6qpK1ve",
	},

	"/api/user/:id": []string{
		"/api/user/4J6qpK1ve",
	},

	"/api/emailtouser/:id": []string{
		"/api/emailtouser/e.urbach@gmail.com",
	},

	"/api/googletouser/:id": []string{
		"/api/googletouser/106530160120373282283",
	},

	"/api/facebooktouser/:id": []string{
		"/api/facebooktouser/10207576239700188",
	},

	"/api/nicktouser/:id": []string{
		"/api/nicktouser/Akyoto",
	},

	"/api/analytics/:id": []string{
		"/api/analytics/4J6qpK1ve",
	},

	"/api/soundtrack/:id": []string{
		"/api/soundtrack/h0ac8sKkg",
	},

	"/api/userfollows/:id": []string{
		"/api/userfollows/4J6qpK1ve",
	},

	"/api/anilisttoanime/:id": []string{
		"/api/anilisttoanime/527",
	},

	"/api/animecharacters/:id": []string{
		"/api/animecharacters/323",
	},

	"/api/animeepisodes/:id": []string{
		"/api/animeepisodes/323",
	},

	"/anime/:id/episode/:episode-number": []string{
		"/anime/12230/episode/5",
	},

	"/api/character/:id": []string{
		"/api/character/6556",
	},

	"/api/pushsubscriptions/:id": []string{
		"/api/pushsubscriptions/4J6qpK1ve",
	},

	"/api/myanimelisttoanime/:id": []string{
		"/api/myanimelisttoanime/527",
	},

	// Images
	"/images/avatars/large/:file": []string{
		"/images/avatars/large/4J6qpK1ve.webp",
	},

	"/images/avatars/small/:file": []string{
		"/images/avatars/small/4J6qpK1ve.webp",
	},

	"/images/brand/:file": []string{
		"/images/brand/64.webp",
	},

	"/images/login/:file": []string{
		"/images/login/google",
	},

	"/images/elements/:file": []string{
		"/images/elements/no-avatar.svg",
	},

	// Extra tests for higher coverage
	"/_/+Akyoto": []string{
		"/_/+Akyoto",
	},

	"/_/search/dragon": []string{
		"/_/search/dragon",
	},

	// Disable these tests because they require authorization
	"/auth/google":                                   nil,
	"/auth/google/callback":                          nil,
	"/auth/facebook":                                 nil,
	"/auth/facebook/callback":                        nil,
	"/dashboard":                                     nil,
	"/import":                                        nil,
	"/import/anilist/animelist":                      nil,
	"/import/anilist/animelist/finish":               nil,
	"/import/myanimelist/animelist":                  nil,
	"/import/myanimelist/animelist/finish":           nil,
	"/import/kitsu/animelist":                        nil,
	"/import/kitsu/animelist/finish":                 nil,
	"/animelist/watching":                            nil,
	"/animelist/completed":                           nil,
	"/animelist/planned":                             nil,
	"/animelist/hold":                                nil,
	"/animelist/dropped":                             nil,
	"/notifications":                                 nil,
	"/user/:nick/notifications":                      nil,
	"/api/test/notification":                         nil,
	"/api/paypal/payment/create":                     nil,
	"/api/userfollows/:id/get/:item":                 nil,
	"/api/userfollows/:id/get/:item/:property":       nil,
	"/api/pushsubscriptions/:id/get/:item":           nil,
	"/api/pushsubscriptions/:id/get/:item/:property": nil,
	"/api/count/notifications/unseen":                nil,
	"/api/mark/notifications/seen":                   nil,
	"/editor/kitsu/new/anime":                        nil,
	"/paypal/success":                                nil,
	"/paypal/cancel":                                 nil,
	"/anime/:id/edit":                                nil,
	"/new/thread":                                    nil,
	"/admin/purchases":                               nil,
	"/editor/anilist":                                nil,
	"/editor/shoboi":                                 nil,
	"/dark-flame-master":                             nil,
	"/user":                                          nil,
	"/settings":                                      nil,
	"/settings/accounts":                             nil,
	"/settings/notifications":                        nil,
	"/settings/apps":                                 nil,
	"/settings/avatar":                               nil,
	"/settings/formatting":                           nil,
	"/settings/pro":                                  nil,
	"/shop":                                          nil,
	"/shop/history":                                  nil,
	"/support":                                       nil,
	"/charge":                                        nil,
	"/log":                                           nil,
	"/inventory":                                     nil,
	"/extension/embed":                               nil,
}
