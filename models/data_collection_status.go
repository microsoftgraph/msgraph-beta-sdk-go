package models
import (
    "errors"
)
type DataCollectionStatus int

const (
    ONLINE_DATACOLLECTIONSTATUS DataCollectionStatus = iota
    OFFLINE_DATACOLLECTIONSTATUS
    UNKNOWNFUTUREVALUE_DATACOLLECTIONSTATUS
)

func (i DataCollectionStatus) String() string {
    return []string{"online", "offline", "unknownFutureValue"}[i]
}
func ParseDataCollectionStatus(v string) (any, error) {
    result := ONLINE_DATACOLLECTIONSTATUS
    switch v {
        case "online":
            result = ONLINE_DATACOLLECTIONSTATUS
        case "offline":
            result = OFFLINE_DATACOLLECTIONSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_DATACOLLECTIONSTATUS
        default:
            return 0, errors.New("Unknown DataCollectionStatus value: " + v)
    }
    return &result, nil
}
func SerializeDataCollectionStatus(values []DataCollectionStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DataCollectionStatus) isMultiValue() bool {
    return false
}
