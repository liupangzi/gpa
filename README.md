## GPA: yet another Golang Persistence API

### Installation

```bash
go get github.com/liupangzi/gpa
```

### Document

#### 1. .sql to Go struct

`gpa mysql model -h` can translate mysql ddl file to Go struct respecting the data alignment, e.g.:

- .sql ddl file
![sql](./docs/images/sql.png "ddl")

- Generated Golang struct
![model](./docs/images/model.png "Go struct")

#### 2. Interface to Implementation

`gpa mysql impl -h` will generate Go implementations from pre-defined interface based on [sqlx](https://github.com/jmoiron/sqlx), e.g.:

- Golang interface with annotations
![interface](./docs/images/interface.png "Golang interface with annotations")

- Generated implementations
![implementation](./docs/images/impl.png "Generated implementations")
