package models
import (
    "errors"
)
// 
type CloudPcConnectivityEventType int

const (
    UNKNOWN_CLOUDPCCONNECTIVITYEVENTTYPE CloudPcConnectivityEventType = iota
    USERCONNECTION_CLOUDPCCONNECTIVITYEVENTTYPE
    USERTROUBLESHOOTING_CLOUDPCCONNECTIVITYEVENTTYPE
    DEVICEHEALTHCHECK_CLOUDPCCONNECTIVITYEVENTTYPE
    UNKNOWNFUTUREVALUE_CLOUDPCCONNECTIVITYEVENTTYPE
)

func (i CloudPcConnectivityEventType) String() string {
    return []string{"unknown", "userConnection", "userTroubleshooting", "deviceHealthCheck", "unknownFutureValue"}[i]
}
func ParseCloudPcConnectivityEventType(v string) (any, error) {
    result := UNKNOWN_CLOUDPCCONNECTIVITYEVENTTYPE
    switch v {
        case "unknown":
            result = UNKNOWN_CLOUDPCCONNECTIVITYEVENTTYPE
        case "userConnection":
            result = USERCONNECTION_CLOUDPCCONNECTIVITYEVENTTYPE
        case "userTroubleshooting":
            result = USERTROUBLESHOOTING_CLOUDPCCONNECTIVITYEVENTTYPE
        case "deviceHealthCheck":
            result = DEVICEHEALTHCHECK_CLOUDPCCONNECTIVITYEVENTTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CLOUDPCCONNECTIVITYEVENTTYPE
        default:
            return 0, errors.New("Unknown CloudPcConnectivityEventType value: " + v)
    }
    return &result, nil
}
func SerializeCloudPcConnectivityEventType(values []CloudPcConnectivityEventType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CloudPcConnectivityEventType) isMultiValue() bool {
    return false
}
