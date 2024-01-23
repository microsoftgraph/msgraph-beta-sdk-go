package models
import (
    "errors"
)
// The set of available policy types that can be configured for approval. There is no default value for this enum, indicating that the policy type must always be chosen.
type OperationApprovalPolicyType int

const (
    // Indicates that the configured policy type is for a Device Action.
    DEVICEACTION_OPERATIONAPPROVALPOLICYTYPE OperationApprovalPolicyType = iota
    // Indicates that the configured policy type is for a Device Wipe Action.
    DEVICEWIPE_OPERATIONAPPROVALPOLICYTYPE
    // Indicates that the configured policy type is for a Device Retire Action.
    DEVICERETIRE_OPERATIONAPPROVALPOLICYTYPE
    // Indicates that the configured policy type is for a Retire Non-Compliant Devices Action.
    DEVICERETIRENONCOMPLIANT_OPERATIONAPPROVALPOLICYTYPE
    // Indicates that the configured policy type is for a Device Delete Action.
    DEVICEDELETE_OPERATIONAPPROVALPOLICYTYPE
    // Indicates that the configured policy type is for a Device Lock Action.
    DEVICELOCK_OPERATIONAPPROVALPOLICYTYPE
    // Indicates that the configured policy type is for a Device Erase Action.
    DEVICEERASE_OPERATIONAPPROVALPOLICYTYPE
    // Indicates that the configured policy type is for a Device Disable Activation Lock Action.
    DEVICEDISABLEACTIVATIONLOCK_OPERATIONAPPROVALPOLICYTYPE
    // Indicates that the configured policy type is for a Windows Enrollment.
    WINDOWSENROLLMENT_OPERATIONAPPROVALPOLICYTYPE
    // Indicates that the configured policy type is for a Compliance Policy.
    COMPLIANCEPOLICY_OPERATIONAPPROVALPOLICYTYPE
    // Indicates that the configured policy type is for a Configuration Policy.
    CONFIGURATIONPOLICY_OPERATIONAPPROVALPOLICYTYPE
    // Indicates that the configured policy type is for an App Protection Policy.
    APPPROTECTIONPOLICY_OPERATIONAPPROVALPOLICYTYPE
    // Indicates that the configured policy type is for a Policy Set.
    POLICYSET_OPERATIONAPPROVALPOLICYTYPE
    // Indicates that the configured policy type is for a Filter.
    FILTER_OPERATIONAPPROVALPOLICYTYPE
    // Indicates that the configured policy type is for an Endpoint Security Policy.
    ENDPOINTSECURITYPOLICY_OPERATIONAPPROVALPOLICYTYPE
    // Indicates that the configured policy type is an application type, such as mobile apps or built-in apps.
    APP_OPERATIONAPPROVALPOLICYTYPE
    // Indicates that the configured policy type is a script type, such as Powershell scripts or remediation scripts.
    SCRIPTS_OPERATIONAPPROVALPOLICYTYPE
    // Indicates that the configured policy type is for a Role.
    ROLE_OPERATIONAPPROVALPOLICYTYPE
    // Indicates that the configured policy type is for a Device Reset Passcode Action.
    DEVICERESETPASSCODE_OPERATIONAPPROVALPOLICYTYPE
    // Indicates that the configured policy type is for a Custom Organizational Message.
    CUSTOMORGANIZATIONALMESSAGE_OPERATIONAPPROVALPOLICYTYPE
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_OPERATIONAPPROVALPOLICYTYPE
)

func (i OperationApprovalPolicyType) String() string {
    return []string{"deviceAction", "deviceWipe", "deviceRetire", "deviceRetireNonCompliant", "deviceDelete", "deviceLock", "deviceErase", "deviceDisableActivationLock", "windowsEnrollment", "compliancePolicy", "configurationPolicy", "appProtectionPolicy", "policySet", "filter", "endpointSecurityPolicy", "app", "scripts", "role", "deviceResetPasscode", "customOrganizationalMessage", "unknownFutureValue"}[i]
}
func ParseOperationApprovalPolicyType(v string) (any, error) {
    result := DEVICEACTION_OPERATIONAPPROVALPOLICYTYPE
    switch v {
        case "deviceAction":
            result = DEVICEACTION_OPERATIONAPPROVALPOLICYTYPE
        case "deviceWipe":
            result = DEVICEWIPE_OPERATIONAPPROVALPOLICYTYPE
        case "deviceRetire":
            result = DEVICERETIRE_OPERATIONAPPROVALPOLICYTYPE
        case "deviceRetireNonCompliant":
            result = DEVICERETIRENONCOMPLIANT_OPERATIONAPPROVALPOLICYTYPE
        case "deviceDelete":
            result = DEVICEDELETE_OPERATIONAPPROVALPOLICYTYPE
        case "deviceLock":
            result = DEVICELOCK_OPERATIONAPPROVALPOLICYTYPE
        case "deviceErase":
            result = DEVICEERASE_OPERATIONAPPROVALPOLICYTYPE
        case "deviceDisableActivationLock":
            result = DEVICEDISABLEACTIVATIONLOCK_OPERATIONAPPROVALPOLICYTYPE
        case "windowsEnrollment":
            result = WINDOWSENROLLMENT_OPERATIONAPPROVALPOLICYTYPE
        case "compliancePolicy":
            result = COMPLIANCEPOLICY_OPERATIONAPPROVALPOLICYTYPE
        case "configurationPolicy":
            result = CONFIGURATIONPOLICY_OPERATIONAPPROVALPOLICYTYPE
        case "appProtectionPolicy":
            result = APPPROTECTIONPOLICY_OPERATIONAPPROVALPOLICYTYPE
        case "policySet":
            result = POLICYSET_OPERATIONAPPROVALPOLICYTYPE
        case "filter":
            result = FILTER_OPERATIONAPPROVALPOLICYTYPE
        case "endpointSecurityPolicy":
            result = ENDPOINTSECURITYPOLICY_OPERATIONAPPROVALPOLICYTYPE
        case "app":
            result = APP_OPERATIONAPPROVALPOLICYTYPE
        case "scripts":
            result = SCRIPTS_OPERATIONAPPROVALPOLICYTYPE
        case "role":
            result = ROLE_OPERATIONAPPROVALPOLICYTYPE
        case "deviceResetPasscode":
            result = DEVICERESETPASSCODE_OPERATIONAPPROVALPOLICYTYPE
        case "customOrganizationalMessage":
            result = CUSTOMORGANIZATIONALMESSAGE_OPERATIONAPPROVALPOLICYTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_OPERATIONAPPROVALPOLICYTYPE
        default:
            return 0, errors.New("Unknown OperationApprovalPolicyType value: " + v)
    }
    return &result, nil
}
func SerializeOperationApprovalPolicyType(values []OperationApprovalPolicyType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i OperationApprovalPolicyType) isMultiValue() bool {
    return false
}
