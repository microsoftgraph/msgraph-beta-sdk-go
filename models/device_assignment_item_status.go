package models
// A list of possible assignment item action status values for the application or configuration regarding their executed action on the managed device. For example, a configuration included in the deviceAssignmentItems list has just been executed the action. Its status starts with inProgress until it's successfully removed to reflect as removed status or failed to be removed to reflect as error status on the managed device. Similar status change happens for restoration process
type DeviceAssignmentItemStatus int

const (
    // Default. Indicates that the device assignment action to remove or restore an application or a configuration is 'initiated' on the managed device
    INITIATED_DEVICEASSIGNMENTITEMSTATUS DeviceAssignmentItemStatus = iota
    // Indicates that the device assignment action to remove or restore an application or a configuration is 'in progress' on the managed device
    INPROGRESS_DEVICEASSIGNMENTITEMSTATUS
    // Indicates that the application or configuration has been successfully removed on the managed device
    REMOVED_DEVICEASSIGNMENTITEMSTATUS
    // Indicates that the application or configuration has failed to be removed or restored on the managed device. The error may be retriable depending on the intent action message and error code
    ERROR_DEVICEASSIGNMENTITEMSTATUS
    // Indicates that the application or configuration has been successfully restored on the managed device
    SUCCEEDED_DEVICEASSIGNMENTITEMSTATUS
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_DEVICEASSIGNMENTITEMSTATUS
)

func (i DeviceAssignmentItemStatus) String() string {
    return []string{"initiated", "inProgress", "removed", "error", "succeeded", "unknownFutureValue"}[i]
}
func ParseDeviceAssignmentItemStatus(v string) (any, error) {
    result := INITIATED_DEVICEASSIGNMENTITEMSTATUS
    switch v {
        case "initiated":
            result = INITIATED_DEVICEASSIGNMENTITEMSTATUS
        case "inProgress":
            result = INPROGRESS_DEVICEASSIGNMENTITEMSTATUS
        case "removed":
            result = REMOVED_DEVICEASSIGNMENTITEMSTATUS
        case "error":
            result = ERROR_DEVICEASSIGNMENTITEMSTATUS
        case "succeeded":
            result = SUCCEEDED_DEVICEASSIGNMENTITEMSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_DEVICEASSIGNMENTITEMSTATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDeviceAssignmentItemStatus(values []DeviceAssignmentItemStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DeviceAssignmentItemStatus) isMultiValue() bool {
    return false
}
