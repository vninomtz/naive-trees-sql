DROP TABLE IF EXISTS Nodes;

CREATE TABLE Nodes (
  node_id INTEGER PRIMARY KEY,
  parent_id INTEGER,
  value INTEGER,
  FOREIGN KEY(parent_id) REFERENCES Nodes(node_id)
);

