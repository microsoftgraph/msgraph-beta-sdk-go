package externalconnectors
import (
    "errors"
    "strings"
)
// 
type ContentExperienceType int

const (
    SEARCH_CONTENTEXPERIENCETYPE ContentExperienceType = iota
    COMPLIANCE_CONTENTEXPERIENCETYPE
    UNKNOWNFUTUREVALUE_CONTENTEXPERIENCETYPE
)

func (i ContentExperienceType) String() string {
    var values []string
    for p := ContentExperienceType(1); p <= UNKNOWNFUTUREVALUE_CONTENTEXPERIENCETYPE; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"search", "compliance", "unknownFutureValue"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseContentExperienceType(v string) (any, error) {
    var result ContentExperienceType
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "search":
                result |= SEARCH_CONTENTEXPERIENCETYPE
            case "compliance":
                result |= COMPLIANCE_CONTENTEXPERIENCETYPE
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_CONTENTEXPERIENCETYPE
            default:
                return 0, errors.New("Unknown ContentExperienceType value: " + v)
        }
    }
    return &result, nil
}
func SerializeContentExperienceType(values []ContentExperienceType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ContentExperienceType) isMultiValue() bool {
    return true
}
