package graph
import (
    "strings"
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
    return []string{"EVERYONE", "EVERYONEEXCEPTGUESTS", "MODERATORS", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseUserNewMessageRestriction(v string) (interface{}, error) {
    result := EVERYONE_USERNEWMESSAGERESTRICTION
    switch strings.ToUpper(v) {
        case "EVERYONE":
            result = EVERYONE_USERNEWMESSAGERESTRICTION
        case "EVERYONEEXCEPTGUESTS":
            result = EVERYONEEXCEPTGUESTS_USERNEWMESSAGERESTRICTION
        case "MODERATORS":
            result = MODERATORS_USERNEWMESSAGERESTRICTION
        case "UNKNOWNFUTUREVALUE":
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
