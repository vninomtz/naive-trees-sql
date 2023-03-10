package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/vninomtz/naive-trees-sql/internal/repo"
)

func main() {
	host := flag.String("host", "localhost", "Server host")
	port := flag.String("port", "8000", "Port for http server")

	flag.Parse()

	http.HandleFunc("/api/nodes", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			nodes, err := repo.NewSqliteRepo().GetNodes()
      if err != nil {
        log.Println(err)
        responseError(w, err.Error())
        return
      }
      w.Header().Set("Content-Type", "application/json")
      json.NewEncoder(w).Encode(nodes)
			return
		}

		w.WriteHeader(http.StatusMethodNotAllowed)
	})

	addr := fmt.Sprintf("%s:%s", *host, *port)

	fmt.Printf("Server runing at: http://%s:%s", *host, *port)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func responseError(w http.ResponseWriter, msg string)  {
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(msg)
}

