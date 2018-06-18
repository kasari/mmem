package mmem

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets7399258b52a66e58e2c3974b0f4ec3404326567b = "package repository\n\n{{ $Table := .Name | toPascal }}\n{{ $table := .Name | toCamel }}\n\ntype {{ $Table }}Repository struct {\n\trecords []*data.{{ $Table }}\n\n\t{{- range .UniqueColumns }}\n\t\trecordBy{{ .Name | toPascal }} map[{{ .Type }}]*data.{{ $Table }}\n\t{{- end}}\n}\n\nfunc New{{ $Table }}Repository() *{{ $Table }}Repository {\n\treturn &{{ $Table }}Repository{}\n}\n\n\nfunc (repo *{{ $Table }}Repository) All() []*data.{{ $Table }} {\n\treturn repo.records\n}\n\nfunc (repo *{{ $Table }}Repository) Where(predicate func({{ $table }}*data.{{ $Table }}) bool) []*data.{{ $Table }} {\n\tresult := []*data.{{ $Table }}{}\n\tfor _, record := range repo.records {\n\t\tif predicate(record) {\n\t\t\tresult = append(result, record)\n\t\t}\n\t}\n\treturn result\n}\n\n{{ range .UniqueColumns }}{{ $Column := .Name | toPascal }}{{ $column := .Name | toPascal }}\n\tfunc (repo *{{ $Table }}Repository) FindBy{{ $Column }}({{ $column }} {{ .Type }}) *data.{{ $Table }} {\n\t\treturn repo.recordBy{{ $Column }}[{{ $column }}]\n\t}\n{{ end }}\n\nfunc (repo *{{ $Table }}Repository) Add({{ $table }}*data.{{ $Table }}) error {\n\trepo.records = append(repo.records, {{ $table }})\n\n\t{{ range .UniqueColumns }}{{ $Column := .Name | toPascal }}\n\t\tif _, ok := repo.recordBy{{ $Column }}[{{ $table }}.{{ $Column }}]; ok {\n\t\t\treturn fmt.Errorf(\"can't add record because duplicate {{ $Column }}(%v)\", {{ $table }}.{{ $Column }})\n\t\t}\n\t\trepo.recordBy{{ $Column }}[{{ $table }}.{{ $Column }}] = {{ $table }}\n\t{{ end }}\n\n\treturn nil\n}\n"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/tmpl": []string{"repository.tmpl"}, "/": []string{"tmpl"}}, map[string]*assets.File{
	"/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1529290802, 1529290802000000000),
		Data:     nil,
	}, "/tmpl": &assets.File{
		Path:     "/tmpl",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1525464675, 1525464675000000000),
		Data:     nil,
	}, "/tmpl/repository.tmpl": &assets.File{
		Path:     "/tmpl/repository.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1529291893, 1529291893000000000),
		Data:     []byte(_Assets7399258b52a66e58e2c3974b0f4ec3404326567b),
	}}, "")
