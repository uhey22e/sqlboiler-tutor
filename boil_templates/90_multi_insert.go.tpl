{{- $alias := .Aliases.Table .Table.Name}}

// Insert inserts multiple records
func (o {{$alias.UpSingular}}Slice) Insert({{if .NoContext}}exec boil.Executor{{else}}ctx context.Context, exec boil.ContextExecutor{{end}}, columns boil.Columns) error {
	for _, v := range o {
		if err := v.Insert({{if not .NoContext}}ctx, {{end -}} exec, columns); err != nil {
			return err
		}
	}
	return nil
}
