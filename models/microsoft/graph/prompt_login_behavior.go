package graph
import (
    "strings"
    "errors"
)
// 
type PromptLoginBehavior int

const (
    TRANSLATETOFRESHPASSWORDAUTHENTICATION_PROMPTLOGINBEHAVIOR PromptLoginBehavior = iota
    NATIVESUPPORT_PROMPTLOGINBEHAVIOR
    DISABLED_PROMPTLOGINBEHAVIOR
    UNKNOWNFUTUREVALUE_PROMPTLOGINBEHAVIOR
)

func (i PromptLoginBehavior) String() string {
    return []string{"TRANSLATETOFRESHPASSWORDAUTHENTICATION", "NATIVESUPPORT", "DISABLED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParsePromptLoginBehavior(v string) (interface{}, error) {
    result := TRANSLATETOFRESHPASSWORDAUTHENTICATION_PROMPTLOGINBEHAVIOR
    switch strings.ToUpper(v) {
        case "TRANSLATETOFRESHPASSWORDAUTHENTICATION":
            result = TRANSLATETOFRESHPASSWORDAUTHENTICATION_PROMPTLOGINBEHAVIOR
        case "NATIVESUPPORT":
            result = NATIVESUPPORT_PROMPTLOGINBEHAVIOR
        case "DISABLED":
            result = DISABLED_PROMPTLOGINBEHAVIOR
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_PROMPTLOGINBEHAVIOR
        default:
            return 0, errors.New("Unknown PromptLoginBehavior value: " + v)
    }
    return &result, nil
}
func SerializePromptLoginBehavior(values []PromptLoginBehavior) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
