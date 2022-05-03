package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the deviceManagement singleton.
type ZebraFotaScheduleMode int

const (
    INSTALLNOW_ZEBRAFOTASCHEDULEMODE ZebraFotaScheduleMode = iota
    SCHEDULED_ZEBRAFOTASCHEDULEMODE
    UNKNOWNFUTUREVALUE_ZEBRAFOTASCHEDULEMODE
)

func (i ZebraFotaScheduleMode) String() string {
    return []string{"INSTALLNOW", "SCHEDULED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseZebraFotaScheduleMode(v string) (interface{}, error) {
    result := INSTALLNOW_ZEBRAFOTASCHEDULEMODE
    switch strings.ToUpper(v) {
        case "INSTALLNOW":
            result = INSTALLNOW_ZEBRAFOTASCHEDULEMODE
        case "SCHEDULED":
            result = SCHEDULED_ZEBRAFOTASCHEDULEMODE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_ZEBRAFOTASCHEDULEMODE
        default:
            return 0, errors.New("Unknown ZebraFotaScheduleMode value: " + v)
    }
    return &result, nil
}
func SerializeZebraFotaScheduleMode(values []ZebraFotaScheduleMode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
