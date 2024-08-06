package models
// Represents the type of the targets that devices will become members of when enrolled with the associated profile. Possible values are staticSecurityGroup.
type EnrollmentTimeDeviceMembershipTargetType int

const (
    // Default value. Do not use.
    UNKNOWN_ENROLLMENTTIMEDEVICEMEMBERSHIPTARGETTYPE EnrollmentTimeDeviceMembershipTargetType = iota
    // Indicates the device membership target specified refer to static Entra Security Groups.
    STATICSECURITYGROUP_ENROLLMENTTIMEDEVICEMEMBERSHIPTARGETTYPE
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_ENROLLMENTTIMEDEVICEMEMBERSHIPTARGETTYPE
)

func (i EnrollmentTimeDeviceMembershipTargetType) String() string {
    return []string{"unknown", "staticSecurityGroup", "unknownFutureValue"}[i]
}
func ParseEnrollmentTimeDeviceMembershipTargetType(v string) (any, error) {
    result := UNKNOWN_ENROLLMENTTIMEDEVICEMEMBERSHIPTARGETTYPE
    switch v {
        case "unknown":
            result = UNKNOWN_ENROLLMENTTIMEDEVICEMEMBERSHIPTARGETTYPE
        case "staticSecurityGroup":
            result = STATICSECURITYGROUP_ENROLLMENTTIMEDEVICEMEMBERSHIPTARGETTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ENROLLMENTTIMEDEVICEMEMBERSHIPTARGETTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeEnrollmentTimeDeviceMembershipTargetType(values []EnrollmentTimeDeviceMembershipTargetType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i EnrollmentTimeDeviceMembershipTargetType) isMultiValue() bool {
    return false
}
