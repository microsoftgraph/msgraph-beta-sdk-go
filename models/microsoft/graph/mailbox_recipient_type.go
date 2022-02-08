package graph
import (
    "strings"
    "errors"
)
// 
type MailboxRecipientType int

const (
    UNKNOWN_MAILBOXRECIPIENTTYPE MailboxRecipientType = iota
    USER_MAILBOXRECIPIENTTYPE
    LINKED_MAILBOXRECIPIENTTYPE
    SHARED_MAILBOXRECIPIENTTYPE
    ROOM_MAILBOXRECIPIENTTYPE
    EQUIPMENT_MAILBOXRECIPIENTTYPE
    OTHERS_MAILBOXRECIPIENTTYPE
)

func (i MailboxRecipientType) String() string {
    return []string{"UNKNOWN", "USER", "LINKED", "SHARED", "ROOM", "EQUIPMENT", "OTHERS"}[i]
}
func ParseMailboxRecipientType(v string) (interface{}, error) {
    result := UNKNOWN_MAILBOXRECIPIENTTYPE
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_MAILBOXRECIPIENTTYPE
        case "USER":
            result = USER_MAILBOXRECIPIENTTYPE
        case "LINKED":
            result = LINKED_MAILBOXRECIPIENTTYPE
        case "SHARED":
            result = SHARED_MAILBOXRECIPIENTTYPE
        case "ROOM":
            result = ROOM_MAILBOXRECIPIENTTYPE
        case "EQUIPMENT":
            result = EQUIPMENT_MAILBOXRECIPIENTTYPE
        case "OTHERS":
            result = OTHERS_MAILBOXRECIPIENTTYPE
        default:
            return 0, errors.New("Unknown MailboxRecipientType value: " + v)
    }
    return &result, nil
}
func SerializeMailboxRecipientType(values []MailboxRecipientType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
