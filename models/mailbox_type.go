package models
type MailboxType int

const (
    UNKNOWN_MAILBOXTYPE MailboxType = iota
    USER_MAILBOXTYPE
    SHARED_MAILBOXTYPE
    UNKNOWNFUTUREVALUE_MAILBOXTYPE
)

func (i MailboxType) String() string {
    return []string{"unknown", "user", "shared", "unknownFutureValue"}[i]
}
func ParseMailboxType(v string) (any, error) {
    result := UNKNOWN_MAILBOXTYPE
    switch v {
        case "unknown":
            result = UNKNOWN_MAILBOXTYPE
        case "user":
            result = USER_MAILBOXTYPE
        case "shared":
            result = SHARED_MAILBOXTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_MAILBOXTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeMailboxType(values []MailboxType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i MailboxType) isMultiValue() bool {
    return false
}
