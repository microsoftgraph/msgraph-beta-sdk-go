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
    switch strings.ToUpper(v) {
        case "UPGRADED":
            return UPGRADED_OPERATINGSYSTEMUPGRADEELIGIBILITY, nil
        case "UNKNOWN":
            return UNKNOWN_OPERATINGSYSTEMUPGRADEELIGIBILITY, nil
        case "NOTCAPABLE":
            return NOTCAPABLE_OPERATINGSYSTEMUPGRADEELIGIBILITY, nil
        case "CAPABLE":
            return CAPABLE_OPERATINGSYSTEMUPGRADEELIGIBILITY, nil
    }
    return 0, errors.New("Unknown OperatingSystemUpgradeEligibility value: " + v)
}
func SerializeOperatingSystemUpgradeEligibility(values []OperatingSystemUpgradeEligibility) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
