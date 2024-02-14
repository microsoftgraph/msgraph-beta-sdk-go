package models
import (
    "errors"
    "math"
    "strings"
)
// Advanced BitLocker State
type AdvancedBitLockerState int

const (
    // Advanced BitLocker State Success
    SUCCESS_ADVANCEDBITLOCKERSTATE = 1
    // User never gave consent for Encryption
    NOUSERCONSENT_ADVANCEDBITLOCKERSTATE = 2
    // Un-protected OS Volume was detected
    OSVOLUMEUNPROTECTED_ADVANCEDBITLOCKERSTATE = 4
    // TPM not used for protection of OS volume, but is required by policy
    OSVOLUMETPMREQUIRED_ADVANCEDBITLOCKERSTATE = 8
    // TPM only protection not used for OS volume, but is required by policy
    OSVOLUMETPMONLYREQUIRED_ADVANCEDBITLOCKERSTATE = 16
    // TPM+PIN protection not used for OS volume, but is required by policy
    OSVOLUMETPMPINREQUIRED_ADVANCEDBITLOCKERSTATE = 32
    // TPM+Startup Key protection not used for OS volume, but is required by policy
    OSVOLUMETPMSTARTUPKEYREQUIRED_ADVANCEDBITLOCKERSTATE = 64
    // TPM+PIN+Startup Key not used for OS volume, but is required by policy
    OSVOLUMETPMPINSTARTUPKEYREQUIRED_ADVANCEDBITLOCKERSTATE = 128
    // Encryption method of OS Volume is different than that set by policy
    OSVOLUMEENCRYPTIONMETHODMISMATCH_ADVANCEDBITLOCKERSTATE = 256
    // Recovery key backup failed
    RECOVERYKEYBACKUPFAILED_ADVANCEDBITLOCKERSTATE = 512
    // Fixed Drive not encrypted
    FIXEDDRIVENOTENCRYPTED_ADVANCEDBITLOCKERSTATE = 1024
    // Encryption method of Fixed Drive is different than that set by policy
    FIXEDDRIVEENCRYPTIONMETHODMISMATCH_ADVANCEDBITLOCKERSTATE = 2048
    // Logged on user is non-admin. This requires “AllowStandardUserEncryption” policy set to 1
    LOGGEDONUSERNONADMIN_ADVANCEDBITLOCKERSTATE = 4096
    // WinRE is not configured
    WINDOWSRECOVERYENVIRONMENTNOTCONFIGURED_ADVANCEDBITLOCKERSTATE = 8192
    // TPM is not available for BitLocker. This means TPM is not present, or TPM unavailable registry override is set or host OS is on portable/rome-able drive
    TPMNOTAVAILABLE_ADVANCEDBITLOCKERSTATE = 16384
    // TPM is not ready for BitLocker
    TPMNOTREADY_ADVANCEDBITLOCKERSTATE = 32768
    // Network not available. This is required for recovery key backup. This is reported for Drive Encryption capable devices
    NETWORKERROR_ADVANCEDBITLOCKERSTATE = 65536
)

func (i AdvancedBitLockerState) String() string {
    var values []string
    options := []string{"success", "noUserConsent", "osVolumeUnprotected", "osVolumeTpmRequired", "osVolumeTpmOnlyRequired", "osVolumeTpmPinRequired", "osVolumeTpmStartupKeyRequired", "osVolumeTpmPinStartupKeyRequired", "osVolumeEncryptionMethodMismatch", "recoveryKeyBackupFailed", "fixedDriveNotEncrypted", "fixedDriveEncryptionMethodMismatch", "loggedOnUserNonAdmin", "windowsRecoveryEnvironmentNotConfigured", "tpmNotAvailable", "tpmNotReady", "networkError"}
    for p := 0; p < 17; p++ {
        mantis := AdvancedBitLockerState(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseAdvancedBitLockerState(v string) (any, error) {
    var result AdvancedBitLockerState
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "success":
                result |= SUCCESS_ADVANCEDBITLOCKERSTATE
            case "noUserConsent":
                result |= NOUSERCONSENT_ADVANCEDBITLOCKERSTATE
            case "osVolumeUnprotected":
                result |= OSVOLUMEUNPROTECTED_ADVANCEDBITLOCKERSTATE
            case "osVolumeTpmRequired":
                result |= OSVOLUMETPMREQUIRED_ADVANCEDBITLOCKERSTATE
            case "osVolumeTpmOnlyRequired":
                result |= OSVOLUMETPMONLYREQUIRED_ADVANCEDBITLOCKERSTATE
            case "osVolumeTpmPinRequired":
                result |= OSVOLUMETPMPINREQUIRED_ADVANCEDBITLOCKERSTATE
            case "osVolumeTpmStartupKeyRequired":
                result |= OSVOLUMETPMSTARTUPKEYREQUIRED_ADVANCEDBITLOCKERSTATE
            case "osVolumeTpmPinStartupKeyRequired":
                result |= OSVOLUMETPMPINSTARTUPKEYREQUIRED_ADVANCEDBITLOCKERSTATE
            case "osVolumeEncryptionMethodMismatch":
                result |= OSVOLUMEENCRYPTIONMETHODMISMATCH_ADVANCEDBITLOCKERSTATE
            case "recoveryKeyBackupFailed":
                result |= RECOVERYKEYBACKUPFAILED_ADVANCEDBITLOCKERSTATE
            case "fixedDriveNotEncrypted":
                result |= FIXEDDRIVENOTENCRYPTED_ADVANCEDBITLOCKERSTATE
            case "fixedDriveEncryptionMethodMismatch":
                result |= FIXEDDRIVEENCRYPTIONMETHODMISMATCH_ADVANCEDBITLOCKERSTATE
            case "loggedOnUserNonAdmin":
                result |= LOGGEDONUSERNONADMIN_ADVANCEDBITLOCKERSTATE
            case "windowsRecoveryEnvironmentNotConfigured":
                result |= WINDOWSRECOVERYENVIRONMENTNOTCONFIGURED_ADVANCEDBITLOCKERSTATE
            case "tpmNotAvailable":
                result |= TPMNOTAVAILABLE_ADVANCEDBITLOCKERSTATE
            case "tpmNotReady":
                result |= TPMNOTREADY_ADVANCEDBITLOCKERSTATE
            case "networkError":
                result |= NETWORKERROR_ADVANCEDBITLOCKERSTATE
            default:
                return 0, errors.New("Unknown AdvancedBitLockerState value: " + v)
        }
    }
    return &result, nil
}
func SerializeAdvancedBitLockerState(values []AdvancedBitLockerState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AdvancedBitLockerState) isMultiValue() bool {
    return true
}
