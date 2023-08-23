package models
import (
    "errors"
)
// 
type OriginalTransferMethods int

const (
    NONE_ORIGINALTRANSFERMETHODS OriginalTransferMethods = iota
    DEVICECODEFLOW_ORIGINALTRANSFERMETHODS
    AUTHENTICATIONTRANSFER_ORIGINALTRANSFERMETHODS
    UNKNOWNFUTUREVALUE_ORIGINALTRANSFERMETHODS
)

func (i OriginalTransferMethods) String() string {
    return []string{"none", "deviceCodeFlow", "authenticationTransfer", "unknownFutureValue"}[i]
}
func ParseOriginalTransferMethods(v string) (any, error) {
    result := NONE_ORIGINALTRANSFERMETHODS
    switch v {
        case "none":
            result = NONE_ORIGINALTRANSFERMETHODS
        case "deviceCodeFlow":
            result = DEVICECODEFLOW_ORIGINALTRANSFERMETHODS
        case "authenticationTransfer":
            result = AUTHENTICATIONTRANSFER_ORIGINALTRANSFERMETHODS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ORIGINALTRANSFERMETHODS
        default:
            return 0, errors.New("Unknown OriginalTransferMethods value: " + v)
    }
    return &result, nil
}
func SerializeOriginalTransferMethods(values []OriginalTransferMethods) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
