package main

import (
)



func main() {
  repo :=  NewSqliteRepo()
  defer repo.Close()

  service := NewOrg(repo)

  service.AddUser(nil, "Victor", "left")
}
