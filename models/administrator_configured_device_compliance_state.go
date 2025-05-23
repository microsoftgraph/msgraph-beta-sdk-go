// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package models
// Administrator configured device compliance state Enum
type AdministratorConfiguredDeviceComplianceState int

const (
    // Set compliance state based on other compliance polices
    BASEDONDEVICECOMPLIANCEPOLICY_ADMINISTRATORCONFIGUREDDEVICECOMPLIANCESTATE AdministratorConfiguredDeviceComplianceState = iota
    // Set compliance to nonCompliant
    NONCOMPLIANT_ADMINISTRATORCONFIGUREDDEVICECOMPLIANCESTATE
)

func (i AdministratorConfiguredDeviceComplianceState) String() string {
    return []string{"basedOnDeviceCompliancePolicy", "nonCompliant"}[i]
}
func ParseAdministratorConfiguredDeviceComplianceState(v string) (any, error) {
    result := BASEDONDEVICECOMPLIANCEPOLICY_ADMINISTRATORCONFIGUREDDEVICECOMPLIANCESTATE
    switch v {
        case "basedOnDeviceCompliancePolicy":
            result = BASEDONDEVICECOMPLIANCEPOLICY_ADMINISTRATORCONFIGUREDDEVICECOMPLIANCESTATE
        case "nonCompliant":
            result = NONCOMPLIANT_ADMINISTRATORCONFIGUREDDEVICECOMPLIANCESTATE
        default:
            return nil, nil
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
func (i AdministratorConfiguredDeviceComplianceState) isMultiValue() bool {
    return false
}
