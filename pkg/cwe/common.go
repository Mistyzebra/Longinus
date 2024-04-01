package cwe

type Finding struct {
	Line        int
	Description string
}

type Auditor interface {
	Audit(code string) ([]Finding, error)
}