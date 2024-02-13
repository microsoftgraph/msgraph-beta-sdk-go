package models
import (
    "errors"
    "math"
    "strings"
)
// Exchange Active Sync services.
type EasServices int

const (
    NONE_EASSERVICES = 1
    // Enables synchronization of calendars.
    CALENDARS_EASSERVICES = 2
    // Enables synchronization of contacts.
    CONTACTS_EASSERVICES = 4
    // Enables synchronization of email.
    EMAIL_EASSERVICES = 8
    // Enables synchronization of notes.
    NOTES_EASSERVICES = 16
    // Enables synchronization of reminders.
    REMINDERS_EASSERVICES = 32
)

func (i EasServices) String() string {
    var values []string
    options := []string{"none", "calendars", "contacts", "email", "notes", "reminders"}
    for p := 0; p < 6; p++ {
        mantis := EasServices(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
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
