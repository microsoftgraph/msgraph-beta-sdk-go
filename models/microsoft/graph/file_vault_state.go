package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "SUCCESS":
            return SUCCESS_FILEVAULTSTATE, nil
        case "DRIVEENCRYPTEDBYUSER":
            return DRIVEENCRYPTEDBYUSER_FILEVAULTSTATE, nil
        case "USERDEFERREDENCRYPTION":
            return USERDEFERREDENCRYPTION_FILEVAULTSTATE, nil
        case "ESCROWNOTENABLED":
            return ESCROWNOTENABLED_FILEVAULTSTATE, nil
    }
    return 0, errors.New("Unknown FileVaultState value: " + v)
}
func SerializeFileVaultState(values []FileVaultState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
