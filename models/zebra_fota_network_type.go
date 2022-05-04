package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the deviceManagement singleton.
type ZebraFotaNetworkType int

const (
    ANY_ZEBRAFOTANETWORKTYPE ZebraFotaNetworkType = iota
    WIFI_ZEBRAFOTANETWORKTYPE
    CELLULAR_ZEBRAFOTANETWORKTYPE
    WIFIANDCELLULAR_ZEBRAFOTANETWORKTYPE
    UNKNOWNFUTUREVALUE_ZEBRAFOTANETWORKTYPE
)

func (i ZebraFotaNetworkType) String() string {
    return []string{"ANY", "WIFI", "CELLULAR", "WIFIANDCELLULAR", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseZebraFotaNetworkType(v string) (interface{}, error) {
    result := ANY_ZEBRAFOTANETWORKTYPE
    switch strings.ToUpper(v) {
        case "ANY":
            result = ANY_ZEBRAFOTANETWORKTYPE
        case "WIFI":
            result = WIFI_ZEBRAFOTANETWORKTYPE
        case "CELLULAR":
            result = CELLULAR_ZEBRAFOTANETWORKTYPE
        case "WIFIANDCELLULAR":
            result = WIFIANDCELLULAR_ZEBRAFOTANETWORKTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_ZEBRAFOTANETWORKTYPE
        default:
            return 0, errors.New("Unknown ZebraFotaNetworkType value: " + v)
    }
    return &result, nil
}
func SerializeZebraFotaNetworkType(values []ZebraFotaNetworkType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
