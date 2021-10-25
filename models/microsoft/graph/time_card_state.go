package graph
import (
    "strings"
    "errors"
)
type TimeCardState int

const (
    CLOCKEDIN_TIMECARDSTATE TimeCardState = iota
    ONBREAK_TIMECARDSTATE
    CLOCKEDOUT_TIMECARDSTATE
    UNKNOWNFUTUREVALUE_TIMECARDSTATE
)

func (i TimeCardState) String() string {
    return []string{"CLOCKEDIN", "ONBREAK", "CLOCKEDOUT", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseTimeCardState(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "CLOCKEDIN":
            return CLOCKEDIN_TIMECARDSTATE, nil
        case "ONBREAK":
            return ONBREAK_TIMECARDSTATE, nil
        case "CLOCKEDOUT":
            return CLOCKEDOUT_TIMECARDSTATE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_TIMECARDSTATE, nil
    }
    return 0, errors.New("Unknown TimeCardState value: " + v)
}
func SerializeTimeCardState(values []TimeCardState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
