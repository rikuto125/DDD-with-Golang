package main

import (
	"example.com/m/config"
	"example.com/m/infrastructure/persistence"
	"example.com/m/interface/handler"
	"example.com/m/usecase"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	// 依存関係を定義
	userPersistence := persistence.NewUserPersistence(config.Connect())
	userUseCase := usecase.NewUserUseCase(userPersistence)
	userHandler := handler.NewUserHandler(userUseCase)

	// ルーティングの設定
	router := httprouter.New()
	router.GET("/api/users", userHandler.Index)

	// サーバ起動
	http.ListenAndServe(":8080", &Server{router})
	log.Fatal(http.ListenAndServe(":8080", router))
}

type Server struct {
	r *httprouter.Router
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "https://example.com")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Add("Access-Control-Allow-Headers", "Origin")
	w.Header().Add("Access-Control-Allow-Headers", "X-Requested-With")
	w.Header().Add("Access-Control-Allow-Headers", "Accept")
	w.Header().Add("Access-Control-Allow-Headers", "Accept-Language")
	w.Header().Set("Content-Type", "application/json")
	s.r.ServeHTTP(w, r)
}
