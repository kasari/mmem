package mmem

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"unicode"

	"github.com/pkg/errors"
	"github.com/serenize/snaker"
	"golang.org/x/tools/imports"
)

type Table struct {
	Name          string
	UniqueColumns []Column
}

type Column struct {
	Name string
	Type string
}

func Generate(outDir string, structs []interface{}) error {
	tables := []*Table{}

	for _, s := range structs {
		if s == nil {
			return fmt.Errorf("nil is not supported")
		}

		val := reflect.Indirect(reflect.ValueOf(s))
		rt := val.Type()

		uniqueColumns := []Column{}

		for i := 0; i < rt.NumField(); i++ {
			f := rt.Field(i)
			value, ok := f.Tag.Lookup("mmem")
			if ok {
				if value == "unique" {
					uniqueColumns = append(uniqueColumns, Column{
						Name: f.Name,
						Type: f.Type.String(),
					})
				}
			}
		}

		tables = append(tables, &Table{
			Name:          rt.Name(),
			UniqueColumns: uniqueColumns,
		})
	}

	f, err := Assets.Open("/tmpl/repository.tmpl")
	if err != nil {
		return err
	}
	text, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	repoTmpl, err := template.New("repository.tmpl").Funcs(FuncMap).Parse(string(text))
	if err != nil {
		return errors.Wrap(err, "can't parse template")
	}

	for _, table := range tables {
		outPath := filepath.Join(outDir, snaker.CamelToSnake(table.Name)+".auto.go")

		buf := &bytes.Buffer{}
		if err := repoTmpl.Execute(buf, table); err != nil {
			return err
		}

		data, err := imports.Process(outDir, buf.Bytes(), nil)
		if err != nil {
			return err
		}

		ioutil.WriteFile(outPath, data, 0644)
	}

	return nil
}

var FuncMap = template.FuncMap{
	"toPascal": snaker.SnakeToCamel,
	"toCamel":  toCamel,
}

func toCamel(s string) string {
	s = snaker.SnakeToCamel(s)
	if len(s) > 0 {
		w := []rune(s)
		w[0] = unicode.ToLower(w[0])
		return string(w)
	}
	return s
}
