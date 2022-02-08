package graph
import (
    "strings"
    "errors"
)
// 
type AdministratorConfiguredDeviceComplianceState int

const (
    BASEDONDEVICECOMPLIANCEPOLICY_ADMINISTRATORCONFIGUREDDEVICECOMPLIANCESTATE AdministratorConfiguredDeviceComplianceState = iota
    NONCOMPLIANT_ADMINISTRATORCONFIGUREDDEVICECOMPLIANCESTATE
)

func (i AdministratorConfiguredDeviceComplianceState) String() string {
    return []string{"BASEDONDEVICECOMPLIANCEPOLICY", "NONCOMPLIANT"}[i]
}
func ParseAdministratorConfiguredDeviceComplianceState(v string) (interface{}, error) {
    result := BASEDONDEVICECOMPLIANCEPOLICY_ADMINISTRATORCONFIGUREDDEVICECOMPLIANCESTATE
    switch strings.ToUpper(v) {
        case "BASEDONDEVICECOMPLIANCEPOLICY":
            result = BASEDONDEVICECOMPLIANCEPOLICY_ADMINISTRATORCONFIGUREDDEVICECOMPLIANCESTATE
        case "NONCOMPLIANT":
            result = NONCOMPLIANT_ADMINISTRATORCONFIGUREDDEVICECOMPLIANCESTATE
        default:
            return 0, errors.New("Unknown AdministratorConfiguredDeviceComplianceState value: " + v)
    }
    return &result, nil
}
func SerializeAdministratorConfiguredDeviceComplianceState(values []AdministratorConfiguredDeviceComplianceState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
