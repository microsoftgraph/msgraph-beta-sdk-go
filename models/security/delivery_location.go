package security
import (
    "errors"
)
type DeliveryLocation int

const (
    UNKNOWN_DELIVERYLOCATION DeliveryLocation = iota
    INBOX_FOLDER_DELIVERYLOCATION
    JUNKFOLDER_DELIVERYLOCATION
    DELETEDFOLDER_DELIVERYLOCATION
    QUARANTINE_DELIVERYLOCATION
    ONPREM_EXTERNAL_DELIVERYLOCATION
    FAILED_DELIVERYLOCATION
    DROPPED_DELIVERYLOCATION
    OTHERS_DELIVERYLOCATION
    UNKNOWNFUTUREVALUE_DELIVERYLOCATION
)

func (i DeliveryLocation) String() string {
    return []string{"unknown", "inbox_folder", "junkFolder", "deletedFolder", "quarantine", "onprem_external", "failed", "dropped", "others", "unknownFutureValue"}[i]
}
func ParseDeliveryLocation(v string) (any, error) {
    result := UNKNOWN_DELIVERYLOCATION
    switch v {
        case "unknown":
            result = UNKNOWN_DELIVERYLOCATION
        case "inbox_folder":
            result = INBOX_FOLDER_DELIVERYLOCATION
        case "junkFolder":
            result = JUNKFOLDER_DELIVERYLOCATION
        case "deletedFolder":
            result = DELETEDFOLDER_DELIVERYLOCATION
        case "quarantine":
            result = QUARANTINE_DELIVERYLOCATION
        case "onprem_external":
            result = ONPREM_EXTERNAL_DELIVERYLOCATION
        case "failed":
            result = FAILED_DELIVERYLOCATION
        case "dropped":
            result = DROPPED_DELIVERYLOCATION
        case "others":
            result = OTHERS_DELIVERYLOCATION
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_DELIVERYLOCATION
        default:
            return 0, errors.New("Unknown DeliveryLocation value: " + v)
    }
    return &result, nil
}
func SerializeDeliveryLocation(values []DeliveryLocation) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DeliveryLocation) isMultiValue() bool {
    return false
}
