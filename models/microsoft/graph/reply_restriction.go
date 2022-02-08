package graph
import (
    "strings"
    "errors"
)
// 
type ReplyRestriction int

const (
    EVERYONE_REPLYRESTRICTION ReplyRestriction = iota
    AUTHORANDMODERATORS_REPLYRESTRICTION
    UNKNOWNFUTUREVALUE_REPLYRESTRICTION
)

func (i ReplyRestriction) String() string {
    return []string{"EVERYONE", "AUTHORANDMODERATORS", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseReplyRestriction(v string) (interface{}, error) {
    result := EVERYONE_REPLYRESTRICTION
    switch strings.ToUpper(v) {
        case "EVERYONE":
            result = EVERYONE_REPLYRESTRICTION
        case "AUTHORANDMODERATORS":
            result = AUTHORANDMODERATORS_REPLYRESTRICTION
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_REPLYRESTRICTION
        default:
            return 0, errors.New("Unknown ReplyRestriction value: " + v)
    }
    return &result, nil
}
func SerializeReplyRestriction(values []ReplyRestriction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
