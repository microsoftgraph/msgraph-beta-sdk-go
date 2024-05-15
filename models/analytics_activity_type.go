package models
type AnalyticsActivityType int

const (
    EMAIL_ANALYTICSACTIVITYTYPE AnalyticsActivityType = iota
    MEETING_ANALYTICSACTIVITYTYPE
    FOCUS_ANALYTICSACTIVITYTYPE
    CHAT_ANALYTICSACTIVITYTYPE
    CALL_ANALYTICSACTIVITYTYPE
)

func (i AnalyticsActivityType) String() string {
    return []string{"Email", "Meeting", "Focus", "Chat", "Call"}[i]
}
func ParseAnalyticsActivityType(v string) (any, error) {
    result := EMAIL_ANALYTICSACTIVITYTYPE
    switch v {
        case "Email":
            result = EMAIL_ANALYTICSACTIVITYTYPE
        case "Meeting":
            result = MEETING_ANALYTICSACTIVITYTYPE
        case "Focus":
            result = FOCUS_ANALYTICSACTIVITYTYPE
        case "Chat":
            result = CHAT_ANALYTICSACTIVITYTYPE
        case "Call":
            result = CALL_ANALYTICSACTIVITYTYPE
        default:
            return nil, nil
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
func (i AnalyticsActivityType) isMultiValue() bool {
    return false
}
