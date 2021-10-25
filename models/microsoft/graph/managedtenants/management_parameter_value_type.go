package managedtenants
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "STRING":
            return STRING_MANAGEMENTPARAMETERVALUETYPE, nil
        case "INTEGER":
            return INTEGER_MANAGEMENTPARAMETERVALUETYPE, nil
        case "BOOLEAN":
            return BOOLEAN_MANAGEMENTPARAMETERVALUETYPE, nil
        case "GUID":
            return GUID_MANAGEMENTPARAMETERVALUETYPE, nil
        case "STRINGCOLLECTION":
            return STRINGCOLLECTION_MANAGEMENTPARAMETERVALUETYPE, nil
        case "INTEGERCOLLECTION":
            return INTEGERCOLLECTION_MANAGEMENTPARAMETERVALUETYPE, nil
        case "BOOLEANCOLLECTION":
            return BOOLEANCOLLECTION_MANAGEMENTPARAMETERVALUETYPE, nil
        case "GUIDCOLLECTION":
            return GUIDCOLLECTION_MANAGEMENTPARAMETERVALUETYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_MANAGEMENTPARAMETERVALUETYPE, nil
    }
    return 0, errors.New("Unknown ManagementParameterValueType value: " + v)
}
func SerializeManagementParameterValueType(values []ManagementParameterValueType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
