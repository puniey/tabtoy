package pbsrc

// 报错行号+3
const templateText = `// Generated by github.com/davyxu/tabtoy
// DO NOT EDIT!!
// Version: {{.Version}}
syntax = "proto3";
package {{.PackageName}};

{{range $sn, $objName := $.Types.EnumNames}}
enum {{$objName}}
{	{{range $fi,$field := $.Types.AllFieldByName $objName}}
	{{$field.FieldName}} = {{$field.Value}}; // {{$field.Name}} {{end}}
}
{{end}}
{{range $sn, $objName := $.Types.StructNames}}
message {{$objName}}
{ {{range $fi,$field := $.Types.AllFieldByName $objName}}
	{{PbType $field}} {{$field.FieldName}} {{PbTag $fi $field}}; // {{$field.Name}} {{end}}
}
{{end}}

// Combine struct
message {{.CombineStructName}}
{	{{range $ti, $tab := $.Datas.AllTables}}
	repeated {{$tab.HeaderType}} {{$tab.HeaderType}} {{PbCombineField $ti}}; // table: {{$tab.HeaderType}} {{end}}
}

`
