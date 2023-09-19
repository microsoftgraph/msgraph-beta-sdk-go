package models
import (
    "errors"
)
// Indicates the type of elevation occured
type PrivilegeManagementElevationType int

const (
    // Default. If the type was unknown on the client for some reasons.
    UNDETERMINED_PRIVILEGEMANAGEMENTELEVATIONTYPE PrivilegeManagementElevationType = iota
    // The elevation was done without any use of endpoint privilege management. For example: the administrator on a client machine elevated an application with their admin right.
    UNMANAGEDELEVATION_PRIVILEGEMANAGEMENTELEVATIONTYPE
    // The elevation was done using the endpoint privilege management zero touch elevation policy.
    ZEROTOUCHELEVATION_PRIVILEGEMANAGEMENTELEVATIONTYPE
    // The elevation was done using the endpoint privilege management user confirmed elevation policy.
    USERCONFIRMEDELEVATION_PRIVILEGEMANAGEMENTELEVATIONTYPE
    // The elevation was done using the endpoint privilege management support approved elevation policy.
    SUPPORTAPPROVEDELEVATION_PRIVILEGEMANAGEMENTELEVATIONTYPE
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_PRIVILEGEMANAGEMENTELEVATIONTYPE
)

func (i PrivilegeManagementElevationType) String() string {
    return []string{"undetermined", "unmanagedElevation", "zeroTouchElevation", "userConfirmedElevation", "supportApprovedElevation", "unknownFutureValue"}[i]
}
func ParsePrivilegeManagementElevationType(v string) (any, error) {
    result := UNDETERMINED_PRIVILEGEMANAGEMENTELEVATIONTYPE
    switch v {
        case "undetermined":
            result = UNDETERMINED_PRIVILEGEMANAGEMENTELEVATIONTYPE
        case "unmanagedElevation":
            result = UNMANAGEDELEVATION_PRIVILEGEMANAGEMENTELEVATIONTYPE
        case "zeroTouchElevation":
            result = ZEROTOUCHELEVATION_PRIVILEGEMANAGEMENTELEVATIONTYPE
        case "userConfirmedElevation":
            result = USERCONFIRMEDELEVATION_PRIVILEGEMANAGEMENTELEVATIONTYPE
        case "supportApprovedElevation":
            result = SUPPORTAPPROVEDELEVATION_PRIVILEGEMANAGEMENTELEVATIONTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_PRIVILEGEMANAGEMENTELEVATIONTYPE
        default:
            return 0, errors.New("Unknown PrivilegeManagementElevationType value: " + v)
    }
    return &result, nil
}
func SerializePrivilegeManagementElevationType(values []PrivilegeManagementElevationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PrivilegeManagementElevationType) isMultiValue() bool {
    return false
}
