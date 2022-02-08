package graph
import (
    "strings"
    "errors"
)
// 
type FileVaultState int

const (
    SUCCESS_FILEVAULTSTATE FileVaultState = iota
    DRIVEENCRYPTEDBYUSER_FILEVAULTSTATE
    USERDEFERREDENCRYPTION_FILEVAULTSTATE
    ESCROWNOTENABLED_FILEVAULTSTATE
)

func (i FileVaultState) String() string {
    return []string{"SUCCESS", "DRIVEENCRYPTEDBYUSER", "USERDEFERREDENCRYPTION", "ESCROWNOTENABLED"}[i]
}
func ParseFileVaultState(v string) (interface{}, error) {
    result := SUCCESS_FILEVAULTSTATE
    switch strings.ToUpper(v) {
        case "SUCCESS":
            result = SUCCESS_FILEVAULTSTATE
        case "DRIVEENCRYPTEDBYUSER":
            result = DRIVEENCRYPTEDBYUSER_FILEVAULTSTATE
        case "USERDEFERREDENCRYPTION":
            result = USERDEFERREDENCRYPTION_FILEVAULTSTATE
        case "ESCROWNOTENABLED":
            result = ESCROWNOTENABLED_FILEVAULTSTATE
        default:
            return 0, errors.New("Unknown FileVaultState value: " + v)
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
