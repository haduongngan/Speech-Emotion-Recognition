package middleware

const (
	TypeRefresh = "TypeRefresh"
	TypeAccess  = "TypeAccess"
)

// func VerifyToken(tokenString string, typeToken string) (jwt.Token, error) {
// 	// Verify the token
// 	token, err := infrastructure.GetDecodeAuth().Decode(tokenString)
// 	if err != nil {
// 		return nil, err
// 	}
// 	// Check token valid
// 	if _, ok := token.AsMap(context.Background()); ok != nil {
// 		return nil, errors.New("Token invalid")
// 	}

// 	var tokenUUID string
// 	var ok bool
// 	claims, _ := token.AsMap(context.Background())
// 	if typeToken == TypeRefresh {
// 		tokenUUID, ok = claims["refresh_uuid"].(string)
// 		if !ok {
// 			return nil, errors.New("Claims of token is invalid")
// 		}
// 	} else {
// 		tokenUUID, ok = claims["access_uuid"].(string)
// 		if !ok {
// 			return nil, errors.New("Claims of token is invalid")
// 		}
// 	}
// 	if userID, err := infrastructure.FetchAuth(tokenUUID); err != nil || userID == 0 {
// 		return nil, errors.New("Token is expired")
// 	}

// 	return token, nil
// }

// func CreateToken(userID uint, role string) (*model.TokenDetail, error) {
// 	var err error
// 	// Create token details
// 	tokenDetail := &model.TokenDetail{}

// 	tokenDetail.AtExpires = time.Now().Add(time.Minute * time.Duration(infrastructure.GetExtendAccessMinute())).Unix()
// 	tokenDetail.AccessUUID = utils.GetPattern(userID) + uuid.NewV4().String()
// 	tokenDetail.RtExpires = time.Now().Add(time.Hour * time.Duration(infrastructure.GetExtendRefreshHour())).Unix()
// 	tokenDetail.RefreshUUID = utils.GetPattern(userID) + uuid.NewV4().String()

// 	// Create Access Token
// 	atClaims := make(map[string]interface{})
// 	atClaims["access_uuid"] = tokenDetail.AccessUUID
// 	atClaims["user_id"] = userID
// 	atClaims["role"] = role
// 	atClaims["exp"] = tokenDetail.AtExpires

// 	_, tokenDetail.AccessToken, err = infrastructure.GetEncodeAuth().Encode(atClaims)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Create Resfresh Token
// 	rtClaims := make(map[string]interface{})
// 	rtClaims["refresh_uuid"] = tokenDetail.RefreshUUID
// 	rtClaims["user_id"] = userID
// 	rtClaims["role"] = role
// 	rtClaims["exp"] = tokenDetail.RtExpires
// 	_, tokenDetail.RefreshToken, err = infrastructure.GetEncodeAuth().Encode(rtClaims)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return tokenDetail, nil
// }

// func CreateAuth(userID uint, tokenDetail *model.TokenDetail) error {
// 	// converting Unix to UTC(to Time Object)
// 	accessToken := time.Unix(tokenDetail.AtExpires, 0)
// 	refreshToken := time.Unix(tokenDetail.RtExpires, 0)
// 	now := time.Now()

// 	if errAccess := infrastructure.
// 		GetRedisClient().
// 		Set(tokenDetail.AccessUUID, strconv.Itoa(int(userID)), accessToken.Sub(now)).
// 		Err(); errAccess != nil {
// 		return errAccess
// 	}

// 	if errRefresh := infrastructure.
// 		GetRedisClient().
// 		Set(tokenDetail.RefreshUUID, strconv.Itoa(int(userID)), refreshToken.Sub(now)).
// 		Err(); errRefresh != nil {
// 		return errRefresh
// 	}

// 	return nil
// }

// // RefreshToken receive refresh token and return pair of token
// func RefreshToken(token jwt.Token) (*model.TokenDetail, error) {
// 	// Get values from claims and check it
// 	claims, _ := token.AsMap(context.Background())
// 	refreshUUID, ok := claims["refresh_uuid"].(string)
// 	if !ok {
// 		return nil, errors.New("Claims of token is invalid")
// 	}
// 	userID, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
// 	if err != nil {
// 		return nil, err
// 	}
// 	role, ok := claims["role"].(string)
// 	if !ok {
// 		return nil, errors.New("Claims of token is invalid")
// 	}

// 	// Delete the previous RefreshToken
// 	deleted, err := infrastructure.DeleteAuth(refreshUUID)
// 	if err != nil || deleted == 0 {
// 		return nil, errors.New("Redis Server has not refreshUUID")
// 	}

// 	// Recreate new pairs of refresh and access token
// 	tokenDetail, err := CreateToken(uint(userID), role)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Save the tokens metadata to redis
// 	if err := CreateAuth(uint(userID), tokenDetail); err != nil {
// 		return nil, err
// 	}

// 	return tokenDetail, nil
// }

// func GetClaimsAfterAuthen(r *http.Request) map[string]interface{} {
// 	accessCookie, _ := r.Cookie(infrastructure.NameAccessTokenInCookie)

// 	token, _ := infrastructure.GetDecodeAuth().Decode(accessCookie.Value)

// 	claims, _ := token.AsMap(context.Background())
// 	return claims
// }

// func Authenticator(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		accessCookie, err := r.Cookie(infrastructure.NameAccessTokenInCookie)
// 		if err != nil {
// 			refreshCookie, err := r.Cookie(infrastructure.NameRefreshTokenInCookie)
// 			if err != nil {
// 				w.WriteHeader(http.StatusUnauthorized)
// 				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
// 				return
// 			}

// 			// Verify the refresh token
// 			refreshToken, err := VerifyToken(refreshCookie.Value, TypeRefresh)
// 			if err != nil {
// 				w.WriteHeader(http.StatusUnauthorized)
// 				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
// 				return
// 			}

// 			tokenDetail, err := RefreshToken(refreshToken)
// 			if err != nil {
// 				w.WriteHeader(http.StatusUnauthorized)
// 				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
// 				return
// 			}

// 			http.SetCookie(w, &http.Cookie{
// 				Name:  infrastructure.NameRefreshTokenInCookie,
// 				Value: tokenDetail.RefreshToken,
// 				// HttpOnly: true,
// 				MaxAge:   infrastructure.GetExtendAccessMinute() * 60,
// 				SameSite: http.SameSiteLaxMode,
// 				Path:     "/",
// 			})

// 			http.SetCookie(w, &http.Cookie{
// 				Name:  infrastructure.NameAccessTokenInCookie,
// 				Value: tokenDetail.AccessToken,
// 				// HttpOnly: true,
// 				MaxAge:   infrastructure.GetExtendRefreshHour() * 3600,
// 				SameSite: http.SameSiteLaxMode,
// 				Path:     "/",
// 			})

// 			next.ServeHTTP(w, r)
// 			return
// 		}

// 		// Verify the access token
// 		if _, err := VerifyToken(accessCookie.Value, TypeAccess); err != nil {
// 			w.WriteHeader(http.StatusUnauthorized)
// 			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
// 			return
// 		}

// 		next.ServeHTTP(w, r)
// 	})
// }
