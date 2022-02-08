package graph
import (
    "strings"
    "errors"
)
// 
type WindowsDefenderApplicationControlSupplementalPolicyStatuses int

const (
    UNKNOWN_WINDOWSDEFENDERAPPLICATIONCONTROLSUPPLEMENTALPOLICYSTATUSES WindowsDefenderApplicationControlSupplementalPolicyStatuses = iota
    SUCCESS_WINDOWSDEFENDERAPPLICATIONCONTROLSUPPLEMENTALPOLICYSTATUSES
    TOKENERROR_WINDOWSDEFENDERAPPLICATIONCONTROLSUPPLEMENTALPOLICYSTATUSES
    NOTAUTHORIZEDBYTOKEN_WINDOWSDEFENDERAPPLICATIONCONTROLSUPPLEMENTALPOLICYSTATUSES
    POLICYNOTFOUND_WINDOWSDEFENDERAPPLICATIONCONTROLSUPPLEMENTALPOLICYSTATUSES
)

func (i WindowsDefenderApplicationControlSupplementalPolicyStatuses) String() string {
    return []string{"UNKNOWN", "SUCCESS", "TOKENERROR", "NOTAUTHORIZEDBYTOKEN", "POLICYNOTFOUND"}[i]
}
func ParseWindowsDefenderApplicationControlSupplementalPolicyStatuses(v string) (interface{}, error) {
    result := UNKNOWN_WINDOWSDEFENDERAPPLICATIONCONTROLSUPPLEMENTALPOLICYSTATUSES
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_WINDOWSDEFENDERAPPLICATIONCONTROLSUPPLEMENTALPOLICYSTATUSES
        case "SUCCESS":
            result = SUCCESS_WINDOWSDEFENDERAPPLICATIONCONTROLSUPPLEMENTALPOLICYSTATUSES
        case "TOKENERROR":
            result = TOKENERROR_WINDOWSDEFENDERAPPLICATIONCONTROLSUPPLEMENTALPOLICYSTATUSES
        case "NOTAUTHORIZEDBYTOKEN":
            result = NOTAUTHORIZEDBYTOKEN_WINDOWSDEFENDERAPPLICATIONCONTROLSUPPLEMENTALPOLICYSTATUSES
        case "POLICYNOTFOUND":
            result = POLICYNOTFOUND_WINDOWSDEFENDERAPPLICATIONCONTROLSUPPLEMENTALPOLICYSTATUSES
        default:
            return 0, errors.New("Unknown WindowsDefenderApplicationControlSupplementalPolicyStatuses value: " + v)
    }
    return &result, nil
}
func SerializeWindowsDefenderApplicationControlSupplementalPolicyStatuses(values []WindowsDefenderApplicationControlSupplementalPolicyStatuses) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
