package models
import (
    "errors"
    "strings"
)
// FileVault State
type FileVaultState int

const (
    // FileVault State Success
    SUCCESS_FILEVAULTSTATE FileVaultState = iota
    // FileVault has been enabled by user and is not being managed by policy
    DRIVEENCRYPTEDBYUSER_FILEVAULTSTATE
    // FileVault policy is successfully installed but user has not started encryption
    USERDEFERREDENCRYPTION_FILEVAULTSTATE
    // FileVault recovery key escrow is not enabled
    ESCROWNOTENABLED_FILEVAULTSTATE
)

func (i FileVaultState) String() string {
    var values []string
    for p := FileVaultState(1); p <= ESCROWNOTENABLED_FILEVAULTSTATE; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"success", "driveEncryptedByUser", "userDeferredEncryption", "escrowNotEnabled"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseFileVaultState(v string) (any, error) {
    var result FileVaultState
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "success":
                result |= SUCCESS_FILEVAULTSTATE
            case "driveEncryptedByUser":
                result |= DRIVEENCRYPTEDBYUSER_FILEVAULTSTATE
            case "userDeferredEncryption":
                result |= USERDEFERREDENCRYPTION_FILEVAULTSTATE
            case "escrowNotEnabled":
                result |= ESCROWNOTENABLED_FILEVAULTSTATE
            default:
                return 0, errors.New("Unknown FileVaultState value: " + v)
        }
    }
    return &result, nil
}
func SerializeFileVaultState(values []FileVaultState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i FileVaultState) isMultiValue() bool {
    return true
}
