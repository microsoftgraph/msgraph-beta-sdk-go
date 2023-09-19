package models
import (
    "errors"
    "strings"
)
// Exchange Active Sync services.
type EasServices int

const (
    NONE_EASSERVICES EasServices = iota
    // Enables synchronization of calendars.
    CALENDARS_EASSERVICES
    // Enables synchronization of contacts.
    CONTACTS_EASSERVICES
    // Enables synchronization of email.
    EMAIL_EASSERVICES
    // Enables synchronization of notes.
    NOTES_EASSERVICES
    // Enables synchronization of reminders.
    REMINDERS_EASSERVICES
)

func (i EasServices) String() string {
    var values []string
    for p := EasServices(1); p <= REMINDERS_EASSERVICES; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"none", "calendars", "contacts", "email", "notes", "reminders"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseEasServices(v string) (any, error) {
    var result EasServices
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "none":
                result |= NONE_EASSERVICES
            case "calendars":
                result |= CALENDARS_EASSERVICES
            case "contacts":
                result |= CONTACTS_EASSERVICES
            case "email":
                result |= EMAIL_EASSERVICES
            case "notes":
                result |= NOTES_EASSERVICES
            case "reminders":
                result |= REMINDERS_EASSERVICES
            default:
                return 0, errors.New("Unknown EasServices value: " + v)
        }
    }
    return &result, nil
}
func SerializeEasServices(values []EasServices) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i EasServices) isMultiValue() bool {
    return true
}
