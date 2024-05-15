package models
type AwsPolicyType int

const (
    SYSTEM_AWSPOLICYTYPE AwsPolicyType = iota
    CUSTOM_AWSPOLICYTYPE
    UNKNOWNFUTUREVALUE_AWSPOLICYTYPE
)

func (i AwsPolicyType) String() string {
    return []string{"system", "custom", "unknownFutureValue"}[i]
}
func ParseAwsPolicyType(v string) (any, error) {
    result := SYSTEM_AWSPOLICYTYPE
    switch v {
        case "system":
            result = SYSTEM_AWSPOLICYTYPE
        case "custom":
            result = CUSTOM_AWSPOLICYTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_AWSPOLICYTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAwsPolicyType(values []AwsPolicyType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AwsPolicyType) isMultiValue() bool {
    return false
}
