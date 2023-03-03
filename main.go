package main

import (
	"flag"
	"fmt"
)

func main() {
	add := flag.Bool("add", false, "Create node")
	reset := flag.Bool("reset", false, "Reset db")
	name := flag.String("name", "", "Nombre nodo")
	parent := flag.Int("parent", 0, "Nodo padre")

	flag.Parse()

	repo := NewSqliteRepo()
	defer repo.Close()

	if *add {
		id, err := repo.SaveNode(*parent, *name)
		if err != nil {
			fmt.Println(err)
			return
		}
	  fmt.Printf("Record id: %d", id)
    return
	}

  if *reset {
    fmt.Println("Limpiando db ...")
    repo.Clean()
    return
  }

  flag.PrintDefaults()
}
