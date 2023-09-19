package networkaccess
import (
    "errors"
)
// 
type DhGroup int

const (
    DHGROUP14_DHGROUP DhGroup = iota
    DHGROUP24_DHGROUP
    DHGROUP2048_DHGROUP
    ECP256_DHGROUP
    ECP384_DHGROUP
    UNKNOWNFUTUREVALUE_DHGROUP
)

func (i DhGroup) String() string {
    return []string{"dhGroup14", "dhGroup24", "dhGroup2048", "ecp256", "ecp384", "unknownFutureValue"}[i]
}
func ParseDhGroup(v string) (any, error) {
    result := DHGROUP14_DHGROUP
    switch v {
        case "dhGroup14":
            result = DHGROUP14_DHGROUP
        case "dhGroup24":
            result = DHGROUP24_DHGROUP
        case "dhGroup2048":
            result = DHGROUP2048_DHGROUP
        case "ecp256":
            result = ECP256_DHGROUP
        case "ecp384":
            result = ECP384_DHGROUP
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_DHGROUP
        default:
            return 0, errors.New("Unknown DhGroup value: " + v)
    }
    return &result, nil
}
func SerializeDhGroup(values []DhGroup) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DhGroup) isMultiValue() bool {
    return false
}
