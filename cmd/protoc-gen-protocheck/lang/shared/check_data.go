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
	ContainerFieldNames  []string
	HasMethodsAvailable  bool // false for proto2,proto3, true for edition2023
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
	IsSetFuncRequired    bool
	IsSetConditionSource string
}
