package models
import (
    "errors"
    "math"
    "strings"
)
// FileVault State
type FileVaultState int

const (
    // FileVault State Success
    SUCCESS_FILEVAULTSTATE = 1
    // FileVault has been enabled by user and is not being managed by policy
    DRIVEENCRYPTEDBYUSER_FILEVAULTSTATE = 2
    // FileVault policy is successfully installed but user has not started encryption
    USERDEFERREDENCRYPTION_FILEVAULTSTATE = 4
    // FileVault recovery key escrow is not enabled
    ESCROWNOTENABLED_FILEVAULTSTATE = 8
)

func (i FileVaultState) String() string {
    var values []string
    options := []string{"success", "driveEncryptedByUser", "userDeferredEncryption", "escrowNotEnabled"}
    for p := 0; p < 4; p++ {
        mantis := FileVaultState(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
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
