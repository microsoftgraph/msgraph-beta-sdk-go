package graph
import (
    "strings"
    "errors"
)
// 
type PersonAnnualEventType int

const (
    BIRTHDAY_PERSONANNUALEVENTTYPE PersonAnnualEventType = iota
    WEDDING_PERSONANNUALEVENTTYPE
    WORK_PERSONANNUALEVENTTYPE
    OTHER_PERSONANNUALEVENTTYPE
    UNKNOWNFUTUREVALUE_PERSONANNUALEVENTTYPE
)

func (i PersonAnnualEventType) String() string {
    return []string{"BIRTHDAY", "WEDDING", "WORK", "OTHER", "UNKNOWNFUTUREVALUE"}[i]
}
func ParsePersonAnnualEventType(v string) (interface{}, error) {
    result := BIRTHDAY_PERSONANNUALEVENTTYPE
    switch strings.ToUpper(v) {
        case "BIRTHDAY":
            result = BIRTHDAY_PERSONANNUALEVENTTYPE
        case "WEDDING":
            result = WEDDING_PERSONANNUALEVENTTYPE
        case "WORK":
            result = WORK_PERSONANNUALEVENTTYPE
        case "OTHER":
            result = OTHER_PERSONANNUALEVENTTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_PERSONANNUALEVENTTYPE
        default:
            return 0, errors.New("Unknown PersonAnnualEventType value: " + v)
    }
    return &result, nil
}
func SerializePersonAnnualEventType(values []PersonAnnualEventType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
