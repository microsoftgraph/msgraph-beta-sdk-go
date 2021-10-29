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
    switch strings.ToUpper(v) {
        case "EMAIL":
            return EMAIL_ANALYTICSACTIVITYTYPE, nil
        case "MEETING":
            return MEETING_ANALYTICSACTIVITYTYPE, nil
        case "FOCUS":
            return FOCUS_ANALYTICSACTIVITYTYPE, nil
        case "CHAT":
            return CHAT_ANALYTICSACTIVITYTYPE, nil
        case "CALL":
            return CALL_ANALYTICSACTIVITYTYPE, nil
    }
    return 0, errors.New("Unknown AnalyticsActivityType value: " + v)
}
func SerializeAnalyticsActivityType(values []AnalyticsActivityType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
