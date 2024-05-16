package models
type SamlAttributeNameFormat int

const (
    UNSPECIFIED_SAMLATTRIBUTENAMEFORMAT SamlAttributeNameFormat = iota
    URI_SAMLATTRIBUTENAMEFORMAT
    BASIC_SAMLATTRIBUTENAMEFORMAT
    UNKNOWNFUTUREVALUE_SAMLATTRIBUTENAMEFORMAT
)

func (i SamlAttributeNameFormat) String() string {
    return []string{"unspecified", "uri", "basic", "unknownFutureValue"}[i]
}
func ParseSamlAttributeNameFormat(v string) (any, error) {
    result := UNSPECIFIED_SAMLATTRIBUTENAMEFORMAT
    switch v {
        case "unspecified":
            result = UNSPECIFIED_SAMLATTRIBUTENAMEFORMAT
        case "uri":
            result = URI_SAMLATTRIBUTENAMEFORMAT
        case "basic":
            result = BASIC_SAMLATTRIBUTENAMEFORMAT
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_SAMLATTRIBUTENAMEFORMAT
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeSamlAttributeNameFormat(values []SamlAttributeNameFormat) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SamlAttributeNameFormat) isMultiValue() bool {
    return false
}
