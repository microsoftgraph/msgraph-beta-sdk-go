package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the deviceManagement singleton.
type ZebraFotaConnectorState int

const (
    NONE_ZEBRAFOTACONNECTORSTATE ZebraFotaConnectorState = iota
    CONNECTED_ZEBRAFOTACONNECTORSTATE
    DISCONNECTED_ZEBRAFOTACONNECTORSTATE
    UNKNOWNFUTUREVALUE_ZEBRAFOTACONNECTORSTATE
)

func (i ZebraFotaConnectorState) String() string {
    return []string{"NONE", "CONNECTED", "DISCONNECTED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseZebraFotaConnectorState(v string) (interface{}, error) {
    result := NONE_ZEBRAFOTACONNECTORSTATE
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_ZEBRAFOTACONNECTORSTATE
        case "CONNECTED":
            result = CONNECTED_ZEBRAFOTACONNECTORSTATE
        case "DISCONNECTED":
            result = DISCONNECTED_ZEBRAFOTACONNECTORSTATE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_ZEBRAFOTACONNECTORSTATE
        default:
            return 0, errors.New("Unknown ZebraFotaConnectorState value: " + v)
    }
    return &result, nil
}
func SerializeZebraFotaConnectorState(values []ZebraFotaConnectorState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
