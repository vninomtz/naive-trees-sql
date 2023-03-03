package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)



type SqliteRepo struct {
  db *sql.DB
}

type NodeRepository interface {
  Save(parent *Node, user *Node) (int, error)
  SaveNode(parent int, value string) (int, error)
  Close()
  Clean()
}

func NewSqliteRepo() NodeRepository{
	//os.Remove("./foo.db")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		log.Fatal(err)
	}
	stmt := `
  CREATE TABLE IF NOT EXISTS Nodes (
    node_id INTEGER PRIMARY KEY,
    parent_id INTEGER,
    value INTEGER,
    FOREIGN KEY(parent_id) REFERENCES Nodes(node_id)
  );
  `
  _, err = db.Exec(stmt)
	if err != nil {
		log.Printf("%q: %s\n", err, stmt)
    log.Fatalf("Error creando DB")
	}

  return &SqliteRepo{
    db: db,
  }
}

func (r *SqliteRepo) Close()  {
  r.db.Close()
}

func (r *SqliteRepo) Clean()  {
  stmt := "delete from Nodes;"
  _, err := r.db.Exec(stmt)
	if err != nil {
		log.Printf("%q: %s\n", err, stmt)
    log.Fatalf("Error limpiando db")
	}
}

func (r *SqliteRepo) Save(parent *Node, user *Node) (int, error)  {
  // TODO: implementar
  fmt.Println(parent)
  fmt.Println(user)
  return 0, nil
}

func (r *SqliteRepo) SaveNode(parent int, value string) (int, error)  {
  // TODO: implementar
  query := "INSERT INTO Nodes (parent_id, value) VALUES (?, ?)"

  res, err := r.db.Exec(query, parent, value)
	if err != nil {
		log.Printf("%q: %s\n", err, query)
    return 0, err
	}

  id, err := res.LastInsertId()
	if err != nil {
		log.Printf("%q:\n", err)
    return 0, err
	}
  return int(id), nil
}
