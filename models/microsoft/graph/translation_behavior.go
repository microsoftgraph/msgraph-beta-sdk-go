package graph
import (
    "strings"
    "errors"
)
// 
type TranslationBehavior int

const (
    ASK_TRANSLATIONBEHAVIOR TranslationBehavior = iota
    YES_TRANSLATIONBEHAVIOR
    NO_TRANSLATIONBEHAVIOR
)

func (i TranslationBehavior) String() string {
    return []string{"ASK", "YES", "NO"}[i]
}
func ParseTranslationBehavior(v string) (interface{}, error) {
    result := ASK_TRANSLATIONBEHAVIOR
    switch strings.ToUpper(v) {
        case "ASK":
            result = ASK_TRANSLATIONBEHAVIOR
        case "YES":
            result = YES_TRANSLATIONBEHAVIOR
        case "NO":
            result = NO_TRANSLATIONBEHAVIOR
        default:
            return 0, errors.New("Unknown TranslationBehavior value: " + v)
    }
    return &result, nil
}
func SerializeTranslationBehavior(values []TranslationBehavior) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
