package app

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"

	"github.com/GoAuth/models"
	"github.com/GoAuth/utils"
)

var JwtAuthentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		noAuth := []string{"/api/user/new", "/api/user/login"} //endpoint yang tidak memerlukan auth
		requestPath := r.URL.Path                              //mengambil endpoint yang akan di hit

		//pengecekan apabila endpoint yang akan di hit adalah endpoint yang tidak memerlukan auth,
		//maka langsung dilakukan return
		for _, value := range noAuth {
			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		response := make(map[string]interface{})
		tokenHeader := r.Header.Get("Authorization") //mengambil isi token dari request

		//Token missing, 403 Forbidden. Unauthorized
		if tokenHeader == "" {
			response = utils.Message(false, "Missing auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			utils.Response(w, response)
			return
		}

		// //Token biasanya datang dalam format `Bearer {token-body}`, kami memeriksa apakah token yang diambil cocok dengan persyaratan ini
		// splitted := strings.Split(tokenHeader, " ")
		// log.Println(splitted[1])

		// if len(splitted) != 2 {
		// 	response = utils.Message(false, "Invalid/Malformated auth token")
		// 	w.WriteHeader(http.StatusForbidden)
		// 	w.Header().Add("Content-Type", "application/json")
		// 	utils.Response(w, response)
		// 	return
		// }

		// tokenPart := splitted[1] //mengambil bagian token yang kita pakai
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenHeader, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})

		if err != nil {
			response := utils.Message(false, "Malformated auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			utils.Response(w, response)
			return
		}

		//token is invalid, mungkin tidak sing di server ini
		if !token.Valid {
			response := utils.Message(false, "Token is not valid")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			utils.Response(w, response)
			return
		}

		fmt.Sprintf("user ", tk.UserId)
		ctx := context.WithValue(r.Context(), "user", tk.UserId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
