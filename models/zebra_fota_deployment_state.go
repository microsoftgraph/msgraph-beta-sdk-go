package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the deviceManagement singleton.
type ZebraFotaDeploymentState int

const (
    PENDINGCREATION_ZEBRAFOTADEPLOYMENTSTATE ZebraFotaDeploymentState = iota
    CREATEFAILED_ZEBRAFOTADEPLOYMENTSTATE
    CREATED_ZEBRAFOTADEPLOYMENTSTATE
    INPROGRESS_ZEBRAFOTADEPLOYMENTSTATE
    COMPLETED_ZEBRAFOTADEPLOYMENTSTATE
    PENDINGCANCEL_ZEBRAFOTADEPLOYMENTSTATE
    CANCELED_ZEBRAFOTADEPLOYMENTSTATE
    UNKNOWNFUTUREVALUE_ZEBRAFOTADEPLOYMENTSTATE
)

func (i ZebraFotaDeploymentState) String() string {
    return []string{"PENDINGCREATION", "CREATEFAILED", "CREATED", "INPROGRESS", "COMPLETED", "PENDINGCANCEL", "CANCELED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseZebraFotaDeploymentState(v string) (interface{}, error) {
    result := PENDINGCREATION_ZEBRAFOTADEPLOYMENTSTATE
    switch strings.ToUpper(v) {
        case "PENDINGCREATION":
            result = PENDINGCREATION_ZEBRAFOTADEPLOYMENTSTATE
        case "CREATEFAILED":
            result = CREATEFAILED_ZEBRAFOTADEPLOYMENTSTATE
        case "CREATED":
            result = CREATED_ZEBRAFOTADEPLOYMENTSTATE
        case "INPROGRESS":
            result = INPROGRESS_ZEBRAFOTADEPLOYMENTSTATE
        case "COMPLETED":
            result = COMPLETED_ZEBRAFOTADEPLOYMENTSTATE
        case "PENDINGCANCEL":
            result = PENDINGCANCEL_ZEBRAFOTADEPLOYMENTSTATE
        case "CANCELED":
            result = CANCELED_ZEBRAFOTADEPLOYMENTSTATE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_ZEBRAFOTADEPLOYMENTSTATE
        default:
            return 0, errors.New("Unknown ZebraFotaDeploymentState value: " + v)
    }
    return &result, nil
}
func SerializeZebraFotaDeploymentState(values []ZebraFotaDeploymentState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
