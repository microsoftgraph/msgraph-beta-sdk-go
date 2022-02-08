package graph
import (
    "strings"
    "errors"
)
// 
type AospWifiSecurityType int

const (
    NONE_AOSPWIFISECURITYTYPE AospWifiSecurityType = iota
    WPA_AOSPWIFISECURITYTYPE
    WEP_AOSPWIFISECURITYTYPE
)

func (i AospWifiSecurityType) String() string {
    return []string{"NONE", "WPA", "WEP"}[i]
}
func ParseAospWifiSecurityType(v string) (interface{}, error) {
    result := NONE_AOSPWIFISECURITYTYPE
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_AOSPWIFISECURITYTYPE
        case "WPA":
            result = WPA_AOSPWIFISECURITYTYPE
        case "WEP":
            result = WEP_AOSPWIFISECURITYTYPE
        default:
            return 0, errors.New("Unknown AospWifiSecurityType value: " + v)
    }
    return &result, nil
}
func SerializeAospWifiSecurityType(values []AospWifiSecurityType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
