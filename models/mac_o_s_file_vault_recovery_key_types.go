package models
import (
    "errors"
    "math"
    "strings"
)
// Recovery key types for macOS FileVault
type MacOSFileVaultRecoveryKeyTypes int

const (
    // Device default value, no intent.
    NOTCONFIGURED_MACOSFILEVAULTRECOVERYKEYTYPES = 1
    // An institutional recovery key is like a “master” recovery key that can be used to unlock any device whose password has been lost.
    INSTITUTIONALRECOVERYKEY_MACOSFILEVAULTRECOVERYKEYTYPES = 2
    // A personal recovery key is a unique code that can be used to unlock the user’s device, even if the password to the device is lost.
    PERSONALRECOVERYKEY_MACOSFILEVAULTRECOVERYKEYTYPES = 4
)

func (i MacOSFileVaultRecoveryKeyTypes) String() string {
    var values []string
    options := []string{"notConfigured", "institutionalRecoveryKey", "personalRecoveryKey"}
    for p := 0; p < 3; p++ {
        mantis := MacOSFileVaultRecoveryKeyTypes(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseMacOSFileVaultRecoveryKeyTypes(v string) (any, error) {
    var result MacOSFileVaultRecoveryKeyTypes
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "notConfigured":
                result |= NOTCONFIGURED_MACOSFILEVAULTRECOVERYKEYTYPES
            case "institutionalRecoveryKey":
                result |= INSTITUTIONALRECOVERYKEY_MACOSFILEVAULTRECOVERYKEYTYPES
            case "personalRecoveryKey":
                result |= PERSONALRECOVERYKEY_MACOSFILEVAULTRECOVERYKEYTYPES
            default:
                return 0, errors.New("Unknown MacOSFileVaultRecoveryKeyTypes value: " + v)
        }
    }
    return &result, nil
}
func SerializeMacOSFileVaultRecoveryKeyTypes(values []MacOSFileVaultRecoveryKeyTypes) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i MacOSFileVaultRecoveryKeyTypes) isMultiValue() bool {
    return true
}
