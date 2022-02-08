package graph
import (
    "strings"
    "errors"
)
// 
type AndroidManagedStoreAccountEnrollmentTarget int

const (
    NONE_ANDROIDMANAGEDSTOREACCOUNTENROLLMENTTARGET AndroidManagedStoreAccountEnrollmentTarget = iota
    ALL_ANDROIDMANAGEDSTOREACCOUNTENROLLMENTTARGET
    TARGETED_ANDROIDMANAGEDSTOREACCOUNTENROLLMENTTARGET
    TARGETEDASENROLLMENTRESTRICTIONS_ANDROIDMANAGEDSTOREACCOUNTENROLLMENTTARGET
)

func (i AndroidManagedStoreAccountEnrollmentTarget) String() string {
    return []string{"NONE", "ALL", "TARGETED", "TARGETEDASENROLLMENTRESTRICTIONS"}[i]
}
func ParseAndroidManagedStoreAccountEnrollmentTarget(v string) (interface{}, error) {
    result := NONE_ANDROIDMANAGEDSTOREACCOUNTENROLLMENTTARGET
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_ANDROIDMANAGEDSTOREACCOUNTENROLLMENTTARGET
        case "ALL":
            result = ALL_ANDROIDMANAGEDSTOREACCOUNTENROLLMENTTARGET
        case "TARGETED":
            result = TARGETED_ANDROIDMANAGEDSTOREACCOUNTENROLLMENTTARGET
        case "TARGETEDASENROLLMENTRESTRICTIONS":
            result = TARGETEDASENROLLMENTRESTRICTIONS_ANDROIDMANAGEDSTOREACCOUNTENROLLMENTTARGET
        default:
            return 0, errors.New("Unknown AndroidManagedStoreAccountEnrollmentTarget value: " + v)
    }
    return &result, nil
}
func SerializeAndroidManagedStoreAccountEnrollmentTarget(values []AndroidManagedStoreAccountEnrollmentTarget) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
