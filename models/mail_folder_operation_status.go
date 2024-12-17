package models
type MailFolderOperationStatus int

const (
    NOTSTARTED_MAILFOLDEROPERATIONSTATUS MailFolderOperationStatus = iota
    RUNNING_MAILFOLDEROPERATIONSTATUS
    SUCCEEDED_MAILFOLDEROPERATIONSTATUS
    FAILED_MAILFOLDEROPERATIONSTATUS
    UNKNOWNFUTUREVALUE_MAILFOLDEROPERATIONSTATUS
)

func (i MailFolderOperationStatus) String() string {
    return []string{"notStarted", "running", "succeeded", "failed", "unknownFutureValue"}[i]
}
func ParseMailFolderOperationStatus(v string) (any, error) {
    result := NOTSTARTED_MAILFOLDEROPERATIONSTATUS
    switch v {
        case "notStarted":
            result = NOTSTARTED_MAILFOLDEROPERATIONSTATUS
        case "running":
            result = RUNNING_MAILFOLDEROPERATIONSTATUS
        case "succeeded":
            result = SUCCEEDED_MAILFOLDEROPERATIONSTATUS
        case "failed":
            result = FAILED_MAILFOLDEROPERATIONSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_MAILFOLDEROPERATIONSTATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeMailFolderOperationStatus(values []MailFolderOperationStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i MailFolderOperationStatus) isMultiValue() bool {
    return false
}
