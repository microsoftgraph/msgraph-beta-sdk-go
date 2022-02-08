package graph
import (
    "strings"
    "errors"
)
// 
type AndroidDeviceOwnerEnrollmentTokenType int

const (
    DEFAULT_ANDROIDDEVICEOWNERENROLLMENTTOKENTYPE AndroidDeviceOwnerEnrollmentTokenType = iota
    CORPORATEOWNEDDEDICATEDDEVICEWITHAZUREADSHAREDMODE_ANDROIDDEVICEOWNERENROLLMENTTOKENTYPE
)

func (i AndroidDeviceOwnerEnrollmentTokenType) String() string {
    return []string{"DEFAULT", "CORPORATEOWNEDDEDICATEDDEVICEWITHAZUREADSHAREDMODE"}[i]
}
func ParseAndroidDeviceOwnerEnrollmentTokenType(v string) (interface{}, error) {
    result := DEFAULT_ANDROIDDEVICEOWNERENROLLMENTTOKENTYPE
    switch strings.ToUpper(v) {
        case "DEFAULT":
            result = DEFAULT_ANDROIDDEVICEOWNERENROLLMENTTOKENTYPE
        case "CORPORATEOWNEDDEDICATEDDEVICEWITHAZUREADSHAREDMODE":
            result = CORPORATEOWNEDDEDICATEDDEVICEWITHAZUREADSHAREDMODE_ANDROIDDEVICEOWNERENROLLMENTTOKENTYPE
        default:
            return 0, errors.New("Unknown AndroidDeviceOwnerEnrollmentTokenType value: " + v)
    }
    return &result, nil
}
func SerializeAndroidDeviceOwnerEnrollmentTokenType(values []AndroidDeviceOwnerEnrollmentTokenType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
