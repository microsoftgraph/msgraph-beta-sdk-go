package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_DATATYPE, nil
        case "BOOLEAN":
            return BOOLEAN_DATATYPE, nil
        case "INT64":
            return INT64_DATATYPE, nil
        case "DOUBLE":
            return DOUBLE_DATATYPE, nil
        case "STRING":
            return STRING_DATATYPE, nil
        case "DATETIME":
            return DATETIME_DATATYPE, nil
        case "VERSION":
            return VERSION_DATATYPE, nil
        case "BASE64":
            return BASE64_DATATYPE, nil
        case "XML":
            return XML_DATATYPE, nil
        case "BOOLEANARRAY":
            return BOOLEANARRAY_DATATYPE, nil
        case "INT64ARRAY":
            return INT64ARRAY_DATATYPE, nil
        case "DOUBLEARRAY":
            return DOUBLEARRAY_DATATYPE, nil
        case "STRINGARRAY":
            return STRINGARRAY_DATATYPE, nil
        case "DATETIMEARRAY":
            return DATETIMEARRAY_DATATYPE, nil
        case "VERSIONARRAY":
            return VERSIONARRAY_DATATYPE, nil
    }
    return 0, errors.New("Unknown DataType value: " + v)
}
func SerializeDataType(values []DataType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
