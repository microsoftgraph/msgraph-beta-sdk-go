package managedtenants
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
    return []string{"string", "integer", "boolean", "guid", "stringCollection", "integerCollection", "booleanCollection", "guidCollection", "unknownFutureValue"}[i]
}
func ParseManagementParameterValueType(v string) (any, error) {
    result := STRING_MANAGEMENTPARAMETERVALUETYPE
    switch v {
        case "string":
            result = STRING_MANAGEMENTPARAMETERVALUETYPE
        case "integer":
            result = INTEGER_MANAGEMENTPARAMETERVALUETYPE
        case "boolean":
            result = BOOLEAN_MANAGEMENTPARAMETERVALUETYPE
        case "guid":
            result = GUID_MANAGEMENTPARAMETERVALUETYPE
        case "stringCollection":
            result = STRINGCOLLECTION_MANAGEMENTPARAMETERVALUETYPE
        case "integerCollection":
            result = INTEGERCOLLECTION_MANAGEMENTPARAMETERVALUETYPE
        case "booleanCollection":
            result = BOOLEANCOLLECTION_MANAGEMENTPARAMETERVALUETYPE
        case "guidCollection":
            result = GUIDCOLLECTION_MANAGEMENTPARAMETERVALUETYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_MANAGEMENTPARAMETERVALUETYPE
        default:
            return nil, nil
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
func (i ManagementParameterValueType) isMultiValue() bool {
    return false
}
