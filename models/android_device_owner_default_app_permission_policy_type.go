package models
// Android Device Owner default app permission policy type.
type AndroidDeviceOwnerDefaultAppPermissionPolicyType int

const (
    // Device default value, no intent.
    DEVICEDEFAULT_ANDROIDDEVICEOWNERDEFAULTAPPPERMISSIONPOLICYTYPE AndroidDeviceOwnerDefaultAppPermissionPolicyType = iota
    // Prompt.
    PROMPT_ANDROIDDEVICEOWNERDEFAULTAPPPERMISSIONPOLICYTYPE
    // Auto grant.
    AUTOGRANT_ANDROIDDEVICEOWNERDEFAULTAPPPERMISSIONPOLICYTYPE
    // Auto deny.
    AUTODENY_ANDROIDDEVICEOWNERDEFAULTAPPPERMISSIONPOLICYTYPE
)

func (i AndroidDeviceOwnerDefaultAppPermissionPolicyType) String() string {
    return []string{"deviceDefault", "prompt", "autoGrant", "autoDeny"}[i]
}
func ParseAndroidDeviceOwnerDefaultAppPermissionPolicyType(v string) (any, error) {
    result := DEVICEDEFAULT_ANDROIDDEVICEOWNERDEFAULTAPPPERMISSIONPOLICYTYPE
    switch v {
        case "deviceDefault":
            result = DEVICEDEFAULT_ANDROIDDEVICEOWNERDEFAULTAPPPERMISSIONPOLICYTYPE
        case "prompt":
            result = PROMPT_ANDROIDDEVICEOWNERDEFAULTAPPPERMISSIONPOLICYTYPE
        case "autoGrant":
            result = AUTOGRANT_ANDROIDDEVICEOWNERDEFAULTAPPPERMISSIONPOLICYTYPE
        case "autoDeny":
            result = AUTODENY_ANDROIDDEVICEOWNERDEFAULTAPPPERMISSIONPOLICYTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAndroidDeviceOwnerDefaultAppPermissionPolicyType(values []AndroidDeviceOwnerDefaultAppPermissionPolicyType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AndroidDeviceOwnerDefaultAppPermissionPolicyType) isMultiValue() bool {
    return false
}
