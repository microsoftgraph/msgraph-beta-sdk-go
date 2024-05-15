package models
// State of lost mode, indicating if lost mode is enabled or disabled
type LostModeState int

const (
    // Lost mode is disabled.
    DISABLED_LOSTMODESTATE LostModeState = iota
    // Lost mode is enabled.
    ENABLED_LOSTMODESTATE
)

func (i LostModeState) String() string {
    return []string{"disabled", "enabled"}[i]
}
func ParseLostModeState(v string) (any, error) {
    result := DISABLED_LOSTMODESTATE
    switch v {
        case "disabled":
            result = DISABLED_LOSTMODESTATE
        case "enabled":
            result = ENABLED_LOSTMODESTATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeLostModeState(values []LostModeState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i LostModeState) isMultiValue() bool {
    return false
}
