package shared

import (
	_ "embed"
)

type FileData struct {
	PkgName            string
	Messages           []MessageData
	JavaOuterClassname string
}
type MessageData struct {
	InitFuncName         string
	LowercaseMessageName string
	MessageName          string
	ObjectTypeName       string
	MessageCheckers      []CheckerData
	FieldCheckers        []CheckerData
	MessageFieldNames    []string
	MapFields            map[string]FieldData
	RepeatedFields       map[string]FieldData
}
type FieldData struct {
	Name            string
	KeyJavaType     string
	ElementJavaType string
}

func (md MessageData) HasChecker() bool {
	return len(md.FieldCheckers) > 0 || len(md.MessageCheckers) > 0
}

type CheckerData struct {
	Comment              string
	FieldName            string // empty for message level
	IsOptional           bool
	OneOfType            string // set for oneof
	OneOfFieldName       string // set for oneof
	ID                   string
	Expr                 string
	Fail                 string
	IsSetConditionSource string
	IsHasMethodAvailable bool
}
