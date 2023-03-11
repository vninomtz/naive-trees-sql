package main

import (
	"flag"
	"fmt"

	"github.com/vninomtz/naive-trees-sql/internal/repo"
	"github.com/vninomtz/naive-trees-sql/internal/tree"
)

func main() {
	add := flag.Bool("add", false, "Create node")
	reset := flag.Bool("reset", false, "Reset db")
	list := flag.Bool("ls", false, "List nodes db")
	name := flag.String("name", "", "Nombre nodo")
	parent := flag.Int("parent", 0, "Nodo padre")
	fill := flag.Bool("fill", false, "Llenar db con datos")
	num := flag.Int("no", 0, "Generar nodos")

	flag.Parse()

	repo := repo.NewSqliteRepo()
	defer repo.Close()

	if *list {
		nodes, err := repo.GetNodes()
		if err != nil {
			fmt.Println(err)
			return
		}
		t := tree.NewLevelOrderTree(nodes)
		t.PrintLevelOrder()
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

	if *fill {
		nodes := []*tree.Node{}

		for i := 0; i < *num; i++ {
			n := &tree.Node{
				Key:   i + 1,
				Value: fmt.Sprintf("N-%d", i+1),
			}
			nodes = append(nodes, n)
		}
    tree.NewLevelOrderTree(nodes)

    for _, n := range nodes {
      repo.Save(n)
    }
    return
	}

	flag.PrintDefaults()
}
