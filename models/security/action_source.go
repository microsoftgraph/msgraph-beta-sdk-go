package security
type ActionSource int

const (
    MANUAL_ACTIONSOURCE ActionSource = iota
    AUTOMATIC_ACTIONSOURCE
    RECOMMENDED_ACTIONSOURCE
    DEFAULT_ACTIONSOURCE
)

func (i ActionSource) String() string {
    return []string{"manual", "automatic", "recommended", "default"}[i]
}
func ParseActionSource(v string) (any, error) {
    result := MANUAL_ACTIONSOURCE
    switch v {
        case "manual":
            result = MANUAL_ACTIONSOURCE
        case "automatic":
            result = AUTOMATIC_ACTIONSOURCE
        case "recommended":
            result = RECOMMENDED_ACTIONSOURCE
        case "default":
            result = DEFAULT_ACTIONSOURCE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeActionSource(values []ActionSource) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ActionSource) isMultiValue() bool {
    return false
}
