package gpaparser

type BindvarsSelect struct {
	StructName string
	Statement  string
	Method     string
	Param      string
	Arg        string
}

func RenderSelectTpl(goImplGenerator *GoImplGenerator, method *GoMethod) []byte {
	return nil
}
