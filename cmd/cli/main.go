package main

import (
	"flag"
	"fmt"

	"github.com/vninomtz/naive-trees-sql/internal/repo"
)

func main() {
	add := flag.Bool("add", false, "Create node")
	reset := flag.Bool("reset", false, "Reset db")
	list := flag.Bool("ls", false, "List nodes db")
	name := flag.String("name", "", "Nombre nodo")
	parent := flag.Int("parent", 0, "Nodo padre")

	flag.Parse()

	repo := repo.NewSqliteRepo()
	defer repo.Close()

	if *list {
		nodes, err := repo.GetNodes()
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, n := range nodes {
			fmt.Printf("Node{Key:%d,Value: %s, Parent: %d}\n",n.Key, n.Value, n.Parent)
		}
		return
	}

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
