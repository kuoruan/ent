// +build ignore

package main

import (
	"log"

	"github.com/facebook/ent/entc"
	"github.com/facebook/ent/entc/gen"
)

func main() {
	opts := []entc.Option{
		entc.TemplateFiles("./ent/template/debug.tmpl"),
	}
	err := entc.Generate("./ent/schema", &gen.Config{
		Templates: []*gen.Template{
			gen.MustParse(gen.NewTemplate("test_import").
				ParseFiles("./ent/template/test_import.tmpl")),
		},
		Target: "./ent",
	}, opts...)
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}