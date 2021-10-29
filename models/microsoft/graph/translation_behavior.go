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
    switch strings.ToUpper(v) {
        case "ASK":
            return ASK_TRANSLATIONBEHAVIOR, nil
        case "YES":
            return YES_TRANSLATIONBEHAVIOR, nil
        case "NO":
            return NO_TRANSLATIONBEHAVIOR, nil
    }
    return 0, errors.New("Unknown TranslationBehavior value: " + v)
}
func SerializeTranslationBehavior(values []TranslationBehavior) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
