package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_MAILBOXRECIPIENTTYPE, nil
        case "USER":
            return USER_MAILBOXRECIPIENTTYPE, nil
        case "LINKED":
            return LINKED_MAILBOXRECIPIENTTYPE, nil
        case "SHARED":
            return SHARED_MAILBOXRECIPIENTTYPE, nil
        case "ROOM":
            return ROOM_MAILBOXRECIPIENTTYPE, nil
        case "EQUIPMENT":
            return EQUIPMENT_MAILBOXRECIPIENTTYPE, nil
        case "OTHERS":
            return OTHERS_MAILBOXRECIPIENTTYPE, nil
    }
    return 0, errors.New("Unknown MailboxRecipientType value: " + v)
}
func SerializeMailboxRecipientType(values []MailboxRecipientType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
