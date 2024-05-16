package models
type TransformationTrimType int

const (
    LEADING_TRANSFORMATIONTRIMTYPE TransformationTrimType = iota
    TRAILING_TRANSFORMATIONTRIMTYPE
    LEADINGANDTRAILING_TRANSFORMATIONTRIMTYPE
    UNKNOWNFUTUREVALUE_TRANSFORMATIONTRIMTYPE
)

func (i TransformationTrimType) String() string {
    return []string{"leading", "trailing", "leadingAndTrailing", "unknownFutureValue"}[i]
}
func ParseTransformationTrimType(v string) (any, error) {
    result := LEADING_TRANSFORMATIONTRIMTYPE
    switch v {
        case "leading":
            result = LEADING_TRANSFORMATIONTRIMTYPE
        case "trailing":
            result = TRAILING_TRANSFORMATIONTRIMTYPE
        case "leadingAndTrailing":
            result = LEADINGANDTRAILING_TRANSFORMATIONTRIMTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_TRANSFORMATIONTRIMTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeTransformationTrimType(values []TransformationTrimType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i TransformationTrimType) isMultiValue() bool {
    return false
}
