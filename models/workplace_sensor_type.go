package models
type WorkplaceSensorType int

const (
    // The occupancy sensor type.
    OCCUPANCY_WORKPLACESENSORTYPE WorkplaceSensorType = iota
    // The people count sensor type.
    PEOPLECOUNT_WORKPLACESENSORTYPE
    // The inferred Occupancy sensor type.
    INFERREDOCCUPANCY_WORKPLACESENSORTYPE
    // The heartbeat sensor type.
    HEARTBEAT_WORKPLACESENSORTYPE
    // The unknown feature value.
    UNKNOWNFUTUREVALUE_WORKPLACESENSORTYPE
)

func (i WorkplaceSensorType) String() string {
    return []string{"occupancy", "peopleCount", "inferredOccupancy", "heartbeat", "unknownFutureValue"}[i]
}
func ParseWorkplaceSensorType(v string) (any, error) {
    result := OCCUPANCY_WORKPLACESENSORTYPE
    switch v {
        case "occupancy":
            result = OCCUPANCY_WORKPLACESENSORTYPE
        case "peopleCount":
            result = PEOPLECOUNT_WORKPLACESENSORTYPE
        case "inferredOccupancy":
            result = INFERREDOCCUPANCY_WORKPLACESENSORTYPE
        case "heartbeat":
            result = HEARTBEAT_WORKPLACESENSORTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_WORKPLACESENSORTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWorkplaceSensorType(values []WorkplaceSensorType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WorkplaceSensorType) isMultiValue() bool {
    return false
}
