package models
import (
    "errors"
)
// 
type UserNewMessageRestriction int

const (
    EVERYONE_USERNEWMESSAGERESTRICTION UserNewMessageRestriction = iota
    EVERYONEEXCEPTGUESTS_USERNEWMESSAGERESTRICTION
    MODERATORS_USERNEWMESSAGERESTRICTION
    UNKNOWNFUTUREVALUE_USERNEWMESSAGERESTRICTION
)

func (i UserNewMessageRestriction) String() string {
    return []string{"everyone", "everyoneExceptGuests", "moderators", "unknownFutureValue"}[i]
}
func ParseUserNewMessageRestriction(v string) (any, error) {
    result := EVERYONE_USERNEWMESSAGERESTRICTION
    switch v {
        case "everyone":
            result = EVERYONE_USERNEWMESSAGERESTRICTION
        case "everyoneExceptGuests":
            result = EVERYONEEXCEPTGUESTS_USERNEWMESSAGERESTRICTION
        case "moderators":
            result = MODERATORS_USERNEWMESSAGERESTRICTION
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_USERNEWMESSAGERESTRICTION
        default:
            return 0, errors.New("Unknown UserNewMessageRestriction value: " + v)
    }
    return &result, nil
}
func SerializeUserNewMessageRestriction(values []UserNewMessageRestriction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i UserNewMessageRestriction) isMultiValue() bool {
    return false
}
