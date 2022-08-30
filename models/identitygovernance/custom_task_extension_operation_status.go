package identitygovernance
import (
    "errors"
)
// Provides operations to call the resume method.
type CustomTaskExtensionOperationStatus int

const (
    COMPLETED_CUSTOMTASKEXTENSIONOPERATIONSTATUS CustomTaskExtensionOperationStatus = iota
    FAILED_CUSTOMTASKEXTENSIONOPERATIONSTATUS
    UNKNOWNFUTUREVALUE_CUSTOMTASKEXTENSIONOPERATIONSTATUS
)

func (i CustomTaskExtensionOperationStatus) String() string {
    return []string{"completed", "failed", "unknownFutureValue"}[i]
}
func ParseCustomTaskExtensionOperationStatus(v string) (interface{}, error) {
    result := COMPLETED_CUSTOMTASKEXTENSIONOPERATIONSTATUS
    switch v {
        case "completed":
            result = COMPLETED_CUSTOMTASKEXTENSIONOPERATIONSTATUS
        case "failed":
            result = FAILED_CUSTOMTASKEXTENSIONOPERATIONSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CUSTOMTASKEXTENSIONOPERATIONSTATUS
        default:
            return 0, errors.New("Unknown CustomTaskExtensionOperationStatus value: " + v)
    }
    return &result, nil
}
func SerializeCustomTaskExtensionOperationStatus(values []CustomTaskExtensionOperationStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
