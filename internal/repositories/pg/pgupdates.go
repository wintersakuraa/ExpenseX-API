package pg

import (
	"fmt"
	"strings"
)

type pgUpdates struct {
	assigments []string
	args       []interface{}
	argId      int
}

func newPGUpdates() *pgUpdates {
	return &pgUpdates{
		assigments: make([]string, 0),
		args:       make([]interface{}, 0),
		argId:      1,
	}
}

func (u *pgUpdates) add(key string, value interface{}) {
	u.args = append(u.args, value)
	u.assigments = append(u.assigments, fmt.Sprintf("%s = $%d", key, u.argId))
	u.argId++
}

func (u *pgUpdates) getSetQuery() string {
	return strings.Join(u.assigments, ", ")
}

func (u *pgUpdates) appendArg(value interface{}) {
	u.args = append(u.args, value)
	u.argId++
}
