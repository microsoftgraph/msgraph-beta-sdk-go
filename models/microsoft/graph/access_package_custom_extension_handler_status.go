package graph
import (
    "strings"
    "errors"
)
// 
type AccessPackageCustomExtensionHandlerStatus int

const (
    REQUESTSENT_ACCESSPACKAGECUSTOMEXTENSIONHANDLERSTATUS AccessPackageCustomExtensionHandlerStatus = iota
    REQUESTRECEIVED_ACCESSPACKAGECUSTOMEXTENSIONHANDLERSTATUS
    UNKNOWNFUTUREVALUE_ACCESSPACKAGECUSTOMEXTENSIONHANDLERSTATUS
)

func (i AccessPackageCustomExtensionHandlerStatus) String() string {
    return []string{"REQUESTSENT", "REQUESTRECEIVED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseAccessPackageCustomExtensionHandlerStatus(v string) (interface{}, error) {
    result := REQUESTSENT_ACCESSPACKAGECUSTOMEXTENSIONHANDLERSTATUS
    switch strings.ToUpper(v) {
        case "REQUESTSENT":
            result = REQUESTSENT_ACCESSPACKAGECUSTOMEXTENSIONHANDLERSTATUS
        case "REQUESTRECEIVED":
            result = REQUESTRECEIVED_ACCESSPACKAGECUSTOMEXTENSIONHANDLERSTATUS
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_ACCESSPACKAGECUSTOMEXTENSIONHANDLERSTATUS
        default:
            return 0, errors.New("Unknown AccessPackageCustomExtensionHandlerStatus value: " + v)
    }
    return &result, nil
}
func SerializeAccessPackageCustomExtensionHandlerStatus(values []AccessPackageCustomExtensionHandlerStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
