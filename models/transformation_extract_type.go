package models
type TransformationExtractType int

const (
    PREFIX_TRANSFORMATIONEXTRACTTYPE TransformationExtractType = iota
    SUFFIX_TRANSFORMATIONEXTRACTTYPE
    UNKNOWNFUTUREVALUE_TRANSFORMATIONEXTRACTTYPE
)

func (i TransformationExtractType) String() string {
    return []string{"prefix", "suffix", "unknownFutureValue"}[i]
}
func ParseTransformationExtractType(v string) (any, error) {
    result := PREFIX_TRANSFORMATIONEXTRACTTYPE
    switch v {
        case "prefix":
            result = PREFIX_TRANSFORMATIONEXTRACTTYPE
        case "suffix":
            result = SUFFIX_TRANSFORMATIONEXTRACTTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_TRANSFORMATIONEXTRACTTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeTransformationExtractType(values []TransformationExtractType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i TransformationExtractType) isMultiValue() bool {
    return false
}
