package graph
import (
    "strings"
    "errors"
)
// 
type OperatingSystemUpgradeEligibility int

const (
    UPGRADED_OPERATINGSYSTEMUPGRADEELIGIBILITY OperatingSystemUpgradeEligibility = iota
    UNKNOWN_OPERATINGSYSTEMUPGRADEELIGIBILITY
    NOTCAPABLE_OPERATINGSYSTEMUPGRADEELIGIBILITY
    CAPABLE_OPERATINGSYSTEMUPGRADEELIGIBILITY
)

func (i OperatingSystemUpgradeEligibility) String() string {
    return []string{"UPGRADED", "UNKNOWN", "NOTCAPABLE", "CAPABLE"}[i]
}
func ParseOperatingSystemUpgradeEligibility(v string) (interface{}, error) {
    result := UPGRADED_OPERATINGSYSTEMUPGRADEELIGIBILITY
    switch strings.ToUpper(v) {
        case "UPGRADED":
            result = UPGRADED_OPERATINGSYSTEMUPGRADEELIGIBILITY
        case "UNKNOWN":
            result = UNKNOWN_OPERATINGSYSTEMUPGRADEELIGIBILITY
        case "NOTCAPABLE":
            result = NOTCAPABLE_OPERATINGSYSTEMUPGRADEELIGIBILITY
        case "CAPABLE":
            result = CAPABLE_OPERATINGSYSTEMUPGRADEELIGIBILITY
        default:
            return 0, errors.New("Unknown OperatingSystemUpgradeEligibility value: " + v)
    }
    return &result, nil
}
func SerializeOperatingSystemUpgradeEligibility(values []OperatingSystemUpgradeEligibility) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
