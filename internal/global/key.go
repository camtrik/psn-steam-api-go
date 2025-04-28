package global

import "time"

// Redis Settings
const (
	// PSN
	USER_TITLES_KEY = "user_titles:%s"

	// STEAM
	OWNED_GAMES_KEY         = "owned_games:%s"
	PLAYER_ACHIEVEMENTS_KEY = "player_achievements:%s:%s"

	DEFAULT_EXPIRATION = 2 * time.Hour

	// XBOX
	XBOX_PLAYER_ACHIEVEMENTS_KEY = "xbox_player_achievements"
	XBOX_PLAYER_GAME_STATS_KEY   = "xbox_player_game_stats:%s"
)

// PSN
const (
	PSN_AUTH_BASE_URL   = "https://ca.account.sony.com/api/authz/v3/oauth"
	PSN_TROPHY_BASE_URL = "https://m.np.playstation.com/api/trophy"
	PSN_STORE_BASE_URL  = "https://www.playstation.com/en-hk/games"
)

// steam
const (
	STEAM_API_BASE_URL = "http://api.steampowered.com"

	STEAM_ICON_BASE_URL = "https://media.steampowered.com/steamcommunity/public/images/apps"

	STEAM_STORE_BASE_URL = "https://store.steampowered.com/app"

	STEAM_CAPSULE_BASE_URL   = "https://cdn.cloudflare.steamstatic.com/steam/apps"
	STEAM_CAPSULE_ART_SMALL  = "capsule_231x87.jpg"
	STEAM_CAPSULE_ART_MEDIUM = "capsule_467x181.jpg"
	STEAM_CAPSULE_ART_LARGE  = "capsule_616x353.jpg"
)

// xbox
const (
	XBOX_API_BASE_URL = "https://xbl.io/api/v2"
)
