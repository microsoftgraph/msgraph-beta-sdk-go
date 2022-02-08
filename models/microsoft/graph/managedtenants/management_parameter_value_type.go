package managedtenants
import (
    "strings"
    "errors"
)
// 
type ManagementParameterValueType int

const (
    STRING_MANAGEMENTPARAMETERVALUETYPE ManagementParameterValueType = iota
    INTEGER_MANAGEMENTPARAMETERVALUETYPE
    BOOLEAN_MANAGEMENTPARAMETERVALUETYPE
    GUID_MANAGEMENTPARAMETERVALUETYPE
    STRINGCOLLECTION_MANAGEMENTPARAMETERVALUETYPE
    INTEGERCOLLECTION_MANAGEMENTPARAMETERVALUETYPE
    BOOLEANCOLLECTION_MANAGEMENTPARAMETERVALUETYPE
    GUIDCOLLECTION_MANAGEMENTPARAMETERVALUETYPE
    UNKNOWNFUTUREVALUE_MANAGEMENTPARAMETERVALUETYPE
)

func (i ManagementParameterValueType) String() string {
    return []string{"STRING", "INTEGER", "BOOLEAN", "GUID", "STRINGCOLLECTION", "INTEGERCOLLECTION", "BOOLEANCOLLECTION", "GUIDCOLLECTION", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseManagementParameterValueType(v string) (interface{}, error) {
    result := STRING_MANAGEMENTPARAMETERVALUETYPE
    switch strings.ToUpper(v) {
        case "STRING":
            result = STRING_MANAGEMENTPARAMETERVALUETYPE
        case "INTEGER":
            result = INTEGER_MANAGEMENTPARAMETERVALUETYPE
        case "BOOLEAN":
            result = BOOLEAN_MANAGEMENTPARAMETERVALUETYPE
        case "GUID":
            result = GUID_MANAGEMENTPARAMETERVALUETYPE
        case "STRINGCOLLECTION":
            result = STRINGCOLLECTION_MANAGEMENTPARAMETERVALUETYPE
        case "INTEGERCOLLECTION":
            result = INTEGERCOLLECTION_MANAGEMENTPARAMETERVALUETYPE
        case "BOOLEANCOLLECTION":
            result = BOOLEANCOLLECTION_MANAGEMENTPARAMETERVALUETYPE
        case "GUIDCOLLECTION":
            result = GUIDCOLLECTION_MANAGEMENTPARAMETERVALUETYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_MANAGEMENTPARAMETERVALUETYPE
        default:
            return 0, errors.New("Unknown ManagementParameterValueType value: " + v)
    }
    return &result, nil
}
func SerializeManagementParameterValueType(values []ManagementParameterValueType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
