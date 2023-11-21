package models
import (
    "errors"
)
// 
type VirtualAppointmentSmsType int

const (
    CONFIRMATION_VIRTUALAPPOINTMENTSMSTYPE VirtualAppointmentSmsType = iota
    RESCHEDULE_VIRTUALAPPOINTMENTSMSTYPE
    CANCELLATION_VIRTUALAPPOINTMENTSMSTYPE
    UNKNOWNFUTUREVALUE_VIRTUALAPPOINTMENTSMSTYPE
)

func (i VirtualAppointmentSmsType) String() string {
    return []string{"confirmation", "reschedule", "cancellation", "unknownFutureValue"}[i]
}
func ParseVirtualAppointmentSmsType(v string) (any, error) {
    result := CONFIRMATION_VIRTUALAPPOINTMENTSMSTYPE
    switch v {
        case "confirmation":
            result = CONFIRMATION_VIRTUALAPPOINTMENTSMSTYPE
        case "reschedule":
            result = RESCHEDULE_VIRTUALAPPOINTMENTSMSTYPE
        case "cancellation":
            result = CANCELLATION_VIRTUALAPPOINTMENTSMSTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_VIRTUALAPPOINTMENTSMSTYPE
        default:
            return 0, errors.New("Unknown VirtualAppointmentSmsType value: " + v)
    }
    return &result, nil
}
func SerializeVirtualAppointmentSmsType(values []VirtualAppointmentSmsType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i VirtualAppointmentSmsType) isMultiValue() bool {
    return false
}
