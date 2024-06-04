package models
// Represents the Validation error of the device membership target.The API will validate the device membership targets specified by the admin to ensure that they exist, that they are of the proper type, and any other target requirements are met such as that the Intune Device Provisioning First Party App is an owner of the target.
type EnrollmentTimeDeviceMembershipTargetValidationErrorCode int

const (
    // Default. Indicates the status of device membership target is not specified. Do not use.
    UNKNOWN_ENROLLMENTTIMEDEVICEMEMBERSHIPTARGETVALIDATIONERRORCODE EnrollmentTimeDeviceMembershipTargetValidationErrorCode = iota
    // Indicates device membership target cannot be found.
    SECURITYGROUPNOTFOUND_ENROLLMENTTIMEDEVICEMEMBERSHIPTARGETVALIDATIONERRORCODE
    // Indicates device membership target is not a security group.
    NOTSECURITYGROUP_ENROLLMENTTIMEDEVICEMEMBERSHIPTARGETVALIDATIONERRORCODE
    // Indicates device membership target which is security group but not a static one.
    NOTSTATICSECURITYGROUP_ENROLLMENTTIMEDEVICEMEMBERSHIPTARGETVALIDATIONERRORCODE
    // Indicates required first party app not the owner of that device membership target.
    FIRSTPARTYAPPNOTANOWNER_ENROLLMENTTIMEDEVICEMEMBERSHIPTARGETVALIDATIONERRORCODE
    // Indicates that device membership target of type security group is not in the RBAC scope of the caller.
    SECURITYGROUPNOTINCALLERSCOPE_ENROLLMENTTIMEDEVICEMEMBERSHIPTARGETVALIDATIONERRORCODE
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_ENROLLMENTTIMEDEVICEMEMBERSHIPTARGETVALIDATIONERRORCODE
)

func (i EnrollmentTimeDeviceMembershipTargetValidationErrorCode) String() string {
    return []string{"unknown", "securityGroupNotFound", "notSecurityGroup", "notStaticSecurityGroup", "firstPartyAppNotAnOwner", "securityGroupNotInCallerScope", "unknownFutureValue"}[i]
}
func ParseEnrollmentTimeDeviceMembershipTargetValidationErrorCode(v string) (any, error) {
    result := UNKNOWN_ENROLLMENTTIMEDEVICEMEMBERSHIPTARGETVALIDATIONERRORCODE
    switch v {
        case "unknown":
            result = UNKNOWN_ENROLLMENTTIMEDEVICEMEMBERSHIPTARGETVALIDATIONERRORCODE
        case "securityGroupNotFound":
            result = SECURITYGROUPNOTFOUND_ENROLLMENTTIMEDEVICEMEMBERSHIPTARGETVALIDATIONERRORCODE
        case "notSecurityGroup":
            result = NOTSECURITYGROUP_ENROLLMENTTIMEDEVICEMEMBERSHIPTARGETVALIDATIONERRORCODE
        case "notStaticSecurityGroup":
            result = NOTSTATICSECURITYGROUP_ENROLLMENTTIMEDEVICEMEMBERSHIPTARGETVALIDATIONERRORCODE
        case "firstPartyAppNotAnOwner":
            result = FIRSTPARTYAPPNOTANOWNER_ENROLLMENTTIMEDEVICEMEMBERSHIPTARGETVALIDATIONERRORCODE
        case "securityGroupNotInCallerScope":
            result = SECURITYGROUPNOTINCALLERSCOPE_ENROLLMENTTIMEDEVICEMEMBERSHIPTARGETVALIDATIONERRORCODE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ENROLLMENTTIMEDEVICEMEMBERSHIPTARGETVALIDATIONERRORCODE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeEnrollmentTimeDeviceMembershipTargetValidationErrorCode(values []EnrollmentTimeDeviceMembershipTargetValidationErrorCode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i EnrollmentTimeDeviceMembershipTargetValidationErrorCode) isMultiValue() bool {
    return false
}
