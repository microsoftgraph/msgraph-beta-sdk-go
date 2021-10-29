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
    switch strings.ToUpper(v) {
        case "BIRTHDAY":
            return BIRTHDAY_PERSONANNUALEVENTTYPE, nil
        case "WEDDING":
            return WEDDING_PERSONANNUALEVENTTYPE, nil
        case "WORK":
            return WORK_PERSONANNUALEVENTTYPE, nil
        case "OTHER":
            return OTHER_PERSONANNUALEVENTTYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_PERSONANNUALEVENTTYPE, nil
    }
    return 0, errors.New("Unknown PersonAnnualEventType value: " + v)
}
func SerializePersonAnnualEventType(values []PersonAnnualEventType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
