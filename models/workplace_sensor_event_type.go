package models
type WorkplaceSensorEventType int

const (
    BADGEIN_WORKPLACESENSOREVENTTYPE WorkplaceSensorEventType = iota
    BADGEOUT_WORKPLACESENSOREVENTTYPE
    UNKNOWNFUTUREVALUE_WORKPLACESENSOREVENTTYPE
)

func (i WorkplaceSensorEventType) String() string {
    return []string{"badgeIn", "badgeOut", "unknownFutureValue"}[i]
}
func ParseWorkplaceSensorEventType(v string) (any, error) {
    result := BADGEIN_WORKPLACESENSOREVENTTYPE
    switch v {
        case "badgeIn":
            result = BADGEIN_WORKPLACESENSOREVENTTYPE
        case "badgeOut":
            result = BADGEOUT_WORKPLACESENSOREVENTTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_WORKPLACESENSOREVENTTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWorkplaceSensorEventType(values []WorkplaceSensorEventType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WorkplaceSensorEventType) isMultiValue() bool {
    return false
}
