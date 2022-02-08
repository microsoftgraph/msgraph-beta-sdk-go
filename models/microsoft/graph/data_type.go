package graph
import (
    "strings"
    "errors"
)
// 
type DataType int

const (
    NONE_DATATYPE DataType = iota
    BOOLEAN_DATATYPE
    INT64_DATATYPE
    DOUBLE_DATATYPE
    STRING_DATATYPE
    DATETIME_DATATYPE
    VERSION_DATATYPE
    BASE64_DATATYPE
    XML_DATATYPE
    BOOLEANARRAY_DATATYPE
    INT64ARRAY_DATATYPE
    DOUBLEARRAY_DATATYPE
    STRINGARRAY_DATATYPE
    DATETIMEARRAY_DATATYPE
    VERSIONARRAY_DATATYPE
)

func (i DataType) String() string {
    return []string{"NONE", "BOOLEAN", "INT64", "DOUBLE", "STRING", "DATETIME", "VERSION", "BASE64", "XML", "BOOLEANARRAY", "INT64ARRAY", "DOUBLEARRAY", "STRINGARRAY", "DATETIMEARRAY", "VERSIONARRAY"}[i]
}
func ParseDataType(v string) (interface{}, error) {
    result := NONE_DATATYPE
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_DATATYPE
        case "BOOLEAN":
            result = BOOLEAN_DATATYPE
        case "INT64":
            result = INT64_DATATYPE
        case "DOUBLE":
            result = DOUBLE_DATATYPE
        case "STRING":
            result = STRING_DATATYPE
        case "DATETIME":
            result = DATETIME_DATATYPE
        case "VERSION":
            result = VERSION_DATATYPE
        case "BASE64":
            result = BASE64_DATATYPE
        case "XML":
            result = XML_DATATYPE
        case "BOOLEANARRAY":
            result = BOOLEANARRAY_DATATYPE
        case "INT64ARRAY":
            result = INT64ARRAY_DATATYPE
        case "DOUBLEARRAY":
            result = DOUBLEARRAY_DATATYPE
        case "STRINGARRAY":
            result = STRINGARRAY_DATATYPE
        case "DATETIMEARRAY":
            result = DATETIMEARRAY_DATATYPE
        case "VERSIONARRAY":
            result = VERSIONARRAY_DATATYPE
        default:
            return 0, errors.New("Unknown DataType value: " + v)
    }
    return &result, nil
}
func SerializeDataType(values []DataType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
