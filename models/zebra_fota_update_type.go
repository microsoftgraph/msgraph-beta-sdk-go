package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the deviceManagement singleton.
type ZebraFotaUpdateType int

const (
    CUSTOM_ZEBRAFOTAUPDATETYPE ZebraFotaUpdateType = iota
    LATEST_ZEBRAFOTAUPDATETYPE
    AUTO_ZEBRAFOTAUPDATETYPE
    UNKNOWNFUTUREVALUE_ZEBRAFOTAUPDATETYPE
)

func (i ZebraFotaUpdateType) String() string {
    return []string{"CUSTOM", "LATEST", "AUTO", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseZebraFotaUpdateType(v string) (interface{}, error) {
    result := CUSTOM_ZEBRAFOTAUPDATETYPE
    switch strings.ToUpper(v) {
        case "CUSTOM":
            result = CUSTOM_ZEBRAFOTAUPDATETYPE
        case "LATEST":
            result = LATEST_ZEBRAFOTAUPDATETYPE
        case "AUTO":
            result = AUTO_ZEBRAFOTAUPDATETYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_ZEBRAFOTAUPDATETYPE
        default:
            return 0, errors.New("Unknown ZebraFotaUpdateType value: " + v)
    }
    return &result, nil
}
func SerializeZebraFotaUpdateType(values []ZebraFotaUpdateType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
