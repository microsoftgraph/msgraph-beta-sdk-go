package models
// A list of possible assignment item action intent values on the application or configuration when executing this action on the managed device. For example, if the application or configuration is intended to be removed on the managed device, then the intent value is remove, and if the application or configuration already under removal through previous actions and is now intended to be restored on the managed device, then the intent value is restore
type DeviceAssignmentItemIntent int

const (
    // Default. Indicates that the deployed application or configuration is intended to be removed on the managed device
    REMOVE_DEVICEASSIGNMENTITEMINTENT DeviceAssignmentItemIntent = iota
    // Indicates that the application or configuration already under removal through previous actions and is now intended to be restored on the managed device
    RESTORE_DEVICEASSIGNMENTITEMINTENT
    // Evolvable enumeration sentinel value. Do not use
    UNKNOWNFUTUREVALUE_DEVICEASSIGNMENTITEMINTENT
)

func (i DeviceAssignmentItemIntent) String() string {
    return []string{"remove", "restore", "unknownFutureValue"}[i]
}
func ParseDeviceAssignmentItemIntent(v string) (any, error) {
    result := REMOVE_DEVICEASSIGNMENTITEMINTENT
    switch v {
        case "remove":
            result = REMOVE_DEVICEASSIGNMENTITEMINTENT
        case "restore":
            result = RESTORE_DEVICEASSIGNMENTITEMINTENT
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_DEVICEASSIGNMENTITEMINTENT
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDeviceAssignmentItemIntent(values []DeviceAssignmentItemIntent) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DeviceAssignmentItemIntent) isMultiValue() bool {
    return false
}
