package main

import (
	"github.com/Constructor-io/swag/v2"
	"github.com/Constructor-io/swag/v2/testdata/delims/api"
	_ "github.com/Constructor-io/swag/v2/testdata/delims/docs"
)

func ReadDoc() string {
	doc, _ := swag.ReadDoc("CustomDelims")
	return doc
}

// @title Swagger Example API
// @version 1.0
// @description Testing custom template delimeters
// @termsOfService http://swagger.io/terms/

func main() {
	api.MyFunc()
}
