package models
import (
    "errors"
    "strings"
)
// 
type TemplateApplicationLevel int

const (
    NONE_TEMPLATEAPPLICATIONLEVEL TemplateApplicationLevel = iota
    NEWPARTNERS_TEMPLATEAPPLICATIONLEVEL
    EXISTINGPARTNERS_TEMPLATEAPPLICATIONLEVEL
    UNKNOWNFUTUREVALUE_TEMPLATEAPPLICATIONLEVEL
)

func (i TemplateApplicationLevel) String() string {
    var values []string
    for p := TemplateApplicationLevel(1); p <= UNKNOWNFUTUREVALUE_TEMPLATEAPPLICATIONLEVEL; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"none", "newPartners", "existingPartners", "unknownFutureValue"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseTemplateApplicationLevel(v string) (any, error) {
    var result TemplateApplicationLevel
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "none":
                result |= NONE_TEMPLATEAPPLICATIONLEVEL
            case "newPartners":
                result |= NEWPARTNERS_TEMPLATEAPPLICATIONLEVEL
            case "existingPartners":
                result |= EXISTINGPARTNERS_TEMPLATEAPPLICATIONLEVEL
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_TEMPLATEAPPLICATIONLEVEL
            default:
                return 0, errors.New("Unknown TemplateApplicationLevel value: " + v)
        }
    }
    return &result, nil
}
func SerializeTemplateApplicationLevel(values []TemplateApplicationLevel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i TemplateApplicationLevel) isMultiValue() bool {
    return true
}
