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
    switch strings.ToUpper(v) {
        case "TRANSLATETOFRESHPASSWORDAUTHENTICATION":
            return TRANSLATETOFRESHPASSWORDAUTHENTICATION_PROMPTLOGINBEHAVIOR, nil
        case "NATIVESUPPORT":
            return NATIVESUPPORT_PROMPTLOGINBEHAVIOR, nil
        case "DISABLED":
            return DISABLED_PROMPTLOGINBEHAVIOR, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_PROMPTLOGINBEHAVIOR, nil
    }
    return 0, errors.New("Unknown PromptLoginBehavior value: " + v)
}
func SerializePromptLoginBehavior(values []PromptLoginBehavior) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
