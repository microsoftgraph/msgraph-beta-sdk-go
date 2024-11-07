package models
type CustomSecurityAttributeComparisonOperator int

const (
    EQUALS_CUSTOMSECURITYATTRIBUTECOMPARISONOPERATOR CustomSecurityAttributeComparisonOperator = iota
    UNKNOWNFUTUREVALUE_CUSTOMSECURITYATTRIBUTECOMPARISONOPERATOR
)

func (i CustomSecurityAttributeComparisonOperator) String() string {
    return []string{"equals", "unknownFutureValue"}[i]
}
func ParseCustomSecurityAttributeComparisonOperator(v string) (any, error) {
    result := EQUALS_CUSTOMSECURITYATTRIBUTECOMPARISONOPERATOR
    switch v {
        case "equals":
            result = EQUALS_CUSTOMSECURITYATTRIBUTECOMPARISONOPERATOR
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CUSTOMSECURITYATTRIBUTECOMPARISONOPERATOR
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCustomSecurityAttributeComparisonOperator(values []CustomSecurityAttributeComparisonOperator) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CustomSecurityAttributeComparisonOperator) isMultiValue() bool {
    return false
}
