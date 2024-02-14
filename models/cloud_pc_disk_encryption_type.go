package models
import (
    "errors"
)
type CloudPcDiskEncryptionType int

const (
    PLATFORMMANAGEDKEY_CLOUDPCDISKENCRYPTIONTYPE CloudPcDiskEncryptionType = iota
    CUSTOMERMANAGEDKEY_CLOUDPCDISKENCRYPTIONTYPE
    UNKNOWNFUTUREVALUE_CLOUDPCDISKENCRYPTIONTYPE
)

func (i CloudPcDiskEncryptionType) String() string {
    return []string{"platformManagedKey", "customerManagedKey", "unknownFutureValue"}[i]
}
func ParseCloudPcDiskEncryptionType(v string) (any, error) {
    result := PLATFORMMANAGEDKEY_CLOUDPCDISKENCRYPTIONTYPE
    switch v {
        case "platformManagedKey":
            result = PLATFORMMANAGEDKEY_CLOUDPCDISKENCRYPTIONTYPE
        case "customerManagedKey":
            result = CUSTOMERMANAGEDKEY_CLOUDPCDISKENCRYPTIONTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CLOUDPCDISKENCRYPTIONTYPE
        default:
            return 0, errors.New("Unknown CloudPcDiskEncryptionType value: " + v)
    }
    return &result, nil
}
func SerializeCloudPcDiskEncryptionType(values []CloudPcDiskEncryptionType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CloudPcDiskEncryptionType) isMultiValue() bool {
    return false
}
