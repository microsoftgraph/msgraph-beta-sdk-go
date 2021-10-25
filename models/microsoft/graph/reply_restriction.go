package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "EVERYONE":
            return EVERYONE_REPLYRESTRICTION, nil
        case "AUTHORANDMODERATORS":
            return AUTHORANDMODERATORS_REPLYRESTRICTION, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_REPLYRESTRICTION, nil
    }
    return 0, errors.New("Unknown ReplyRestriction value: " + v)
}
func SerializeReplyRestriction(values []ReplyRestriction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
