package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "EVERYONE":
            return EVERYONE_USERNEWMESSAGERESTRICTION, nil
        case "EVERYONEEXCEPTGUESTS":
            return EVERYONEEXCEPTGUESTS_USERNEWMESSAGERESTRICTION, nil
        case "MODERATORS":
            return MODERATORS_USERNEWMESSAGERESTRICTION, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_USERNEWMESSAGERESTRICTION, nil
    }
    return 0, errors.New("Unknown UserNewMessageRestriction value: " + v)
}
func SerializeUserNewMessageRestriction(values []UserNewMessageRestriction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
