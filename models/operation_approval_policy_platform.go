package models
import (
    "errors"
    "strings"
)
// The set of available platforms for the OperationApprovalPolicy. Allows configuration of a policy to specific platform(s) for approval. If no specific platform is required or applicable, the platform is `notApplicable`.
type OperationApprovalPolicyPlatform int

const (
    // Default. Indicates that a policy platform is not required for a chosen policy type.
    NOTAPPLICABLE_OPERATIONAPPROVALPOLICYPLATFORM OperationApprovalPolicyPlatform = iota
    // Indicates that the configured policy platform is for Android Device Administrator.
    ANDROIDDEVICEADMINISTRATOR_OPERATIONAPPROVALPOLICYPLATFORM
    // Indicates that the configured policy platform is for Android Enterprise.
    ANDROIDENTERPRISE_OPERATIONAPPROVALPOLICYPLATFORM
    // Indicates that the configured policy platform is for iOS/iPadOS.
    IOSIPADOS_OPERATIONAPPROVALPOLICYPLATFORM
    // Indicates that the configured policy platform is for macOS.
    MACOS_OPERATIONAPPROVALPOLICYPLATFORM
    // Indicates that the configured policy platform is for Windows 10 and later.
    WINDOWS10ANDLATER_OPERATIONAPPROVALPOLICYPLATFORM
    // Indicates that the configured policy platform is for Windows 8.1 and later.
    WINDOWS81ANDLATER_OPERATIONAPPROVALPOLICYPLATFORM
    // Indicates that the configured policy platform is for Windows 10X.
    WINDOWS10X_OPERATIONAPPROVALPOLICYPLATFORM
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_OPERATIONAPPROVALPOLICYPLATFORM
)

func (i OperationApprovalPolicyPlatform) String() string {
    var values []string
    for p := OperationApprovalPolicyPlatform(1); p <= UNKNOWNFUTUREVALUE_OPERATIONAPPROVALPOLICYPLATFORM; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"notApplicable", "androidDeviceAdministrator", "androidEnterprise", "iOSiPadOS", "macOS", "windows10AndLater", "windows81AndLater", "windows10X", "unknownFutureValue"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseOperationApprovalPolicyPlatform(v string) (any, error) {
    var result OperationApprovalPolicyPlatform
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "notApplicable":
                result |= NOTAPPLICABLE_OPERATIONAPPROVALPOLICYPLATFORM
            case "androidDeviceAdministrator":
                result |= ANDROIDDEVICEADMINISTRATOR_OPERATIONAPPROVALPOLICYPLATFORM
            case "androidEnterprise":
                result |= ANDROIDENTERPRISE_OPERATIONAPPROVALPOLICYPLATFORM
            case "iOSiPadOS":
                result |= IOSIPADOS_OPERATIONAPPROVALPOLICYPLATFORM
            case "macOS":
                result |= MACOS_OPERATIONAPPROVALPOLICYPLATFORM
            case "windows10AndLater":
                result |= WINDOWS10ANDLATER_OPERATIONAPPROVALPOLICYPLATFORM
            case "windows81AndLater":
                result |= WINDOWS81ANDLATER_OPERATIONAPPROVALPOLICYPLATFORM
            case "windows10X":
                result |= WINDOWS10X_OPERATIONAPPROVALPOLICYPLATFORM
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_OPERATIONAPPROVALPOLICYPLATFORM
            default:
                return 0, errors.New("Unknown OperationApprovalPolicyPlatform value: " + v)
        }
    }
    return &result, nil
}
func SerializeOperationApprovalPolicyPlatform(values []OperationApprovalPolicyPlatform) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i OperationApprovalPolicyPlatform) isMultiValue() bool {
    return true
}
