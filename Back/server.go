package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/[username]/todoapp-graphql-go-react/app/graph"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"xorm.io/xorm"
)

const defaultPort = "8080"

func main() {
	// DBへの接続情報を設定
	connectionString := "user=postgres password=postgres dbname=testdb host=db port=5432 sslmode=disable"
	engine, err := xorm.NewEngine("postgres", connectionString)
	if err != nil {
		log.Fatalln("error - create engine: ", err)
	}

	// DBへ接続
	err = engine.Ping()
	if err != nil {
		log.Fatalln("error - connect DB: ", err)
	}
	log.Println("success - connect DB")

	// 環境変数からポート番号を取得、設定されていない場合はデフォルト値を使用
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// gqlgenのサーバを新規作成し、リゾルバとしてDB接続を渡す
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{DB: engine}}))

	// フロントエンドから接続可能にするためCORSを設定
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})
	handler := c.Handler(http.DefaultServeMux)

	// ルートURLでGraphQLのPlaygroundを起動
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	// /queryのパスでGraphQLのエンドポイントを設定
	http.Handle("/query", srv)

	// 注意: Docker環境のportとローカル環境のportが違うため、実際にローカル環境から接続するportは異なり8081である
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	// サーバを起動、エラーが発生した場合はログに出力して終了
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
