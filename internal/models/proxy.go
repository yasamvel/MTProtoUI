package models

type Proxy struct {
    ID       uint
    Server   string
    Port     int
    Type     string
    Secret   string
    Domain   string
    Enabled  bool
}
