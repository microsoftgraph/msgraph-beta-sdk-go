package graph
import (
    "strings"
    "errors"
)
// 
type AnalyticsActivityType int

const (
    EMAIL_ANALYTICSACTIVITYTYPE AnalyticsActivityType = iota
    MEETING_ANALYTICSACTIVITYTYPE
    FOCUS_ANALYTICSACTIVITYTYPE
    CHAT_ANALYTICSACTIVITYTYPE
    CALL_ANALYTICSACTIVITYTYPE
)

func (i AnalyticsActivityType) String() string {
    return []string{"EMAIL", "MEETING", "FOCUS", "CHAT", "CALL"}[i]
}
func ParseAnalyticsActivityType(v string) (interface{}, error) {
    result := EMAIL_ANALYTICSACTIVITYTYPE
    switch strings.ToUpper(v) {
        case "EMAIL":
            result = EMAIL_ANALYTICSACTIVITYTYPE
        case "MEETING":
            result = MEETING_ANALYTICSACTIVITYTYPE
        case "FOCUS":
            result = FOCUS_ANALYTICSACTIVITYTYPE
        case "CHAT":
            result = CHAT_ANALYTICSACTIVITYTYPE
        case "CALL":
            result = CALL_ANALYTICSACTIVITYTYPE
        default:
            return 0, errors.New("Unknown AnalyticsActivityType value: " + v)
    }
    return &result, nil
}
func SerializeAnalyticsActivityType(values []AnalyticsActivityType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
