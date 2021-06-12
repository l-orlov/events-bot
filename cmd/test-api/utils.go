package main

import (
	"strconv"
	"strings"
)

func buildCategoriesParam(categories []string) string {
	ret := ""

	for i, category := range categories {
		if i != 0 {
			ret += "," + category
		} else {
			ret += category
		}
	}

	return ret
}

func buildIsFreeParam(isFree bool) string {
	if isFree {
		return "1"
	}

	return "0"
}

func buildUrl(params Params) string {
	builder := strings.Builder{}

	builder.WriteString(kudaGoUrl)

	builder.WriteString(languageParam)
	builder.WriteString("ru")

	builder.WriteString(locationParam)
	builder.WriteString(params.Location)

	builder.WriteString(dataStartParam)
	builder.WriteString(strconv.FormatInt(params.StartDate, 10))

	builder.WriteString(dataEndParam)
	builder.WriteString(strconv.FormatInt(params.EndDate, 10))

	builder.WriteString(categoriesParam)
	builder.WriteString(buildCategoriesParam(params.Categories))

	builder.WriteString(isFreeParam)
	builder.WriteString(buildIsFreeParam(params.IsFree))

	return builder.String()
}
