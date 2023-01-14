package utils

// redis key configuration
const (
	LoveBookPrefix = "ls_"

	// jwt token
	AuthClientPurpose          = LoveBookPrefix + "auth_client_"
	AuthorizeDataPurpose       = LoveBookPrefix + "authorize_data_"
	AccessDataPurpose          = LoveBookPrefix + "access_data_"
	AccessTokenPurpose         = LoveBookPrefix + "access_token_"
	RefreshTokenPurpose        = LoveBookPrefix + "refresh_token_"
	AccessTokenForUserPurpose  = LoveBookPrefix + "access_token_user_"
	RefreshTokenForUserPurpose = LoveBookPrefix + "refresh_token_user_"
)
