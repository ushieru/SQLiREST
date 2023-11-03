package utils

import (
	"strings"

	"github.com/huandu/go-sqlbuilder"
)

func BuildWhereParams(cb *sqlbuilder.Cond, queries map[string]string) []string {
	delete(queries, "select")
	params := make([]string, 0)
	i := 0
	for k, v := range queries {
		sv := strings.Split(v, ".")
		if len(sv) != 2 {
			continue
		}
		condS := sv[0]
		val := sv[1]
		switch condS {
		case "eq":
			params = append(params, cb.Equal(k, val))
		case "gt":
			params = append(params, cb.GreaterThan(k, val))
		case "gte":
			params = append(params, cb.GreaterEqualThan(k, val))
		case "lt":
			params = append(params, cb.LessThan(k, val))
		case "lte":
			params = append(params, cb.LessEqualThan(k, val))
		case "neq":
			params = append(params, cb.NotEqual(k, val))
		case "like":
			params = append(params, cb.Like(k, val))
		}
		i++
	}
	return params
}
