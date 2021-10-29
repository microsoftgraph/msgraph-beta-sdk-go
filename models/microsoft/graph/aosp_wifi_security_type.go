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
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_AOSPWIFISECURITYTYPE, nil
        case "WPA":
            return WPA_AOSPWIFISECURITYTYPE, nil
        case "WEP":
            return WEP_AOSPWIFISECURITYTYPE, nil
    }
    return 0, errors.New("Unknown AospWifiSecurityType value: " + v)
}
func SerializeAospWifiSecurityType(values []AospWifiSecurityType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
