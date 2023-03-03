package main

import (
	"database/sql"
	"log"
	"os"
	_ "github.com/mattn/go-sqlite3"
)



type SqliteRepo struct {
  db *sql.DB
}

type NodeRepository interface {
  Save(parent *Node, user *Node) (int, error)
  Close()
}

func NewSqliteRepo() NodeRepository{
	os.Remove("./foo.db")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		log.Fatal(err)
	}

	stmt := `
  CREATE TABLE Nodes (
    node_id INTEGER PRIMARY KEY,
    parent_id INTEGER,
    value INTEGER,
    FOREIGN KEY(parent_id) REFERENCES Nodes(node_id)
  );
  delete from Nodes;
  `
	_, err = db.Exec(stmt)
	if err != nil {
		log.Printf("%q: %s\n", err, stmt)
    log.Fatalf("Error creando db")
	}
  return &SqliteRepo{
    db: db,
  }
}

func (r *SqliteRepo) Close()  {
  r.db.Close()
}

func (r *SqliteRepo) Save(parent *Node, user *Node) (int, error)  {
  // TODO: implementar
  return 0, nil
}
