package models
import (
    "errors"
)
// Provides operations to manage the collection of accessReview entities.
type CloudPcDiskEncryptionState int

const (
    NOTAVAILABLE_CLOUDPCDISKENCRYPTIONSTATE CloudPcDiskEncryptionState = iota
    NOTENCRYPED_CLOUDPCDISKENCRYPTIONSTATE
    ENCRYPTEDUSINGPLATFORMMANAGEDKEY_CLOUDPCDISKENCRYPTIONSTATE
    ENCRYPTEDUSINGCUSTOMERMANAGEDKEY_CLOUDPCDISKENCRYPTIONSTATE
    UNKNOWNFUTUREVALUE_CLOUDPCDISKENCRYPTIONSTATE
)

func (i CloudPcDiskEncryptionState) String() string {
    return []string{"notAvailable", "notEncryped", "encryptedUsingPlatformManagedKey", "encryptedUsingCustomerManagedKey", "unknownFutureValue"}[i]
}
func ParseCloudPcDiskEncryptionState(v string) (interface{}, error) {
    result := NOTAVAILABLE_CLOUDPCDISKENCRYPTIONSTATE
    switch v {
        case "notAvailable":
            result = NOTAVAILABLE_CLOUDPCDISKENCRYPTIONSTATE
        case "notEncryped":
            result = NOTENCRYPED_CLOUDPCDISKENCRYPTIONSTATE
        case "encryptedUsingPlatformManagedKey":
            result = ENCRYPTEDUSINGPLATFORMMANAGEDKEY_CLOUDPCDISKENCRYPTIONSTATE
        case "encryptedUsingCustomerManagedKey":
            result = ENCRYPTEDUSINGCUSTOMERMANAGEDKEY_CLOUDPCDISKENCRYPTIONSTATE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CLOUDPCDISKENCRYPTIONSTATE
        default:
            return 0, errors.New("Unknown CloudPcDiskEncryptionState value: " + v)
    }
    return &result, nil
}
func SerializeCloudPcDiskEncryptionState(values []CloudPcDiskEncryptionState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
