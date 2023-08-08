package models
import (
    "errors"
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
    return []string{"none", "newPartners", "existingPartners", "unknownFutureValue"}[i]
}
func ParseTemplateApplicationLevel(v string) (any, error) {
    result := NONE_TEMPLATEAPPLICATIONLEVEL
    switch v {
        case "none":
            result = NONE_TEMPLATEAPPLICATIONLEVEL
        case "newPartners":
            result = NEWPARTNERS_TEMPLATEAPPLICATIONLEVEL
        case "existingPartners":
            result = EXISTINGPARTNERS_TEMPLATEAPPLICATIONLEVEL
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_TEMPLATEAPPLICATIONLEVEL
        default:
            return 0, errors.New("Unknown TemplateApplicationLevel value: " + v)
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
