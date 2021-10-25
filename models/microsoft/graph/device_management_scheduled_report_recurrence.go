package graph
import (
    "strings"
    "errors"
)
type DeviceManagementScheduledReportRecurrence int

const (
    NONE_DEVICEMANAGEMENTSCHEDULEDREPORTRECURRENCE DeviceManagementScheduledReportRecurrence = iota
    DAILY_DEVICEMANAGEMENTSCHEDULEDREPORTRECURRENCE
    WEEKLY_DEVICEMANAGEMENTSCHEDULEDREPORTRECURRENCE
    MONTHLY_DEVICEMANAGEMENTSCHEDULEDREPORTRECURRENCE
)

func (i DeviceManagementScheduledReportRecurrence) String() string {
    return []string{"NONE", "DAILY", "WEEKLY", "MONTHLY"}[i]
}
func ParseDeviceManagementScheduledReportRecurrence(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_DEVICEMANAGEMENTSCHEDULEDREPORTRECURRENCE, nil
        case "DAILY":
            return DAILY_DEVICEMANAGEMENTSCHEDULEDREPORTRECURRENCE, nil
        case "WEEKLY":
            return WEEKLY_DEVICEMANAGEMENTSCHEDULEDREPORTRECURRENCE, nil
        case "MONTHLY":
            return MONTHLY_DEVICEMANAGEMENTSCHEDULEDREPORTRECURRENCE, nil
    }
    return 0, errors.New("Unknown DeviceManagementScheduledReportRecurrence value: " + v)
}
func SerializeDeviceManagementScheduledReportRecurrence(values []DeviceManagementScheduledReportRecurrence) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
