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
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_WINDOWSDEFENDERAPPLICATIONCONTROLSUPPLEMENTALPOLICYSTATUSES, nil
        case "SUCCESS":
            return SUCCESS_WINDOWSDEFENDERAPPLICATIONCONTROLSUPPLEMENTALPOLICYSTATUSES, nil
        case "TOKENERROR":
            return TOKENERROR_WINDOWSDEFENDERAPPLICATIONCONTROLSUPPLEMENTALPOLICYSTATUSES, nil
        case "NOTAUTHORIZEDBYTOKEN":
            return NOTAUTHORIZEDBYTOKEN_WINDOWSDEFENDERAPPLICATIONCONTROLSUPPLEMENTALPOLICYSTATUSES, nil
        case "POLICYNOTFOUND":
            return POLICYNOTFOUND_WINDOWSDEFENDERAPPLICATIONCONTROLSUPPLEMENTALPOLICYSTATUSES, nil
    }
    return 0, errors.New("Unknown WindowsDefenderApplicationControlSupplementalPolicyStatuses value: " + v)
}
func SerializeWindowsDefenderApplicationControlSupplementalPolicyStatuses(values []WindowsDefenderApplicationControlSupplementalPolicyStatuses) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
