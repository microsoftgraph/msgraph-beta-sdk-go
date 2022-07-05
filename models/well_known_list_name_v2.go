package models
import (
    "errors"
)
// Provides operations to manage the collection of accessReview entities.
type WellKnownListName_v2 int

const (
    NONE_WELLKNOWNLISTNAME_V2 WellKnownListName_v2 = iota
    DEFAULTLIST_WELLKNOWNLISTNAME_V2
    FLAGGEDEMAILS_WELLKNOWNLISTNAME_V2
    UNKNOWNFUTUREVALUE_WELLKNOWNLISTNAME_V2
)

func (i WellKnownListName_v2) String() string {
    return []string{"none", "defaultList", "flaggedEmails", "unknownFutureValue"}[i]
}
func ParseWellKnownListName_v2(v string) (interface{}, error) {
    result := NONE_WELLKNOWNLISTNAME_V2
    switch v {
        case "none":
            result = NONE_WELLKNOWNLISTNAME_V2
        case "defaultList":
            result = DEFAULTLIST_WELLKNOWNLISTNAME_V2
        case "flaggedEmails":
            result = FLAGGEDEMAILS_WELLKNOWNLISTNAME_V2
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_WELLKNOWNLISTNAME_V2
        default:
            return 0, errors.New("Unknown WellKnownListName_v2 value: " + v)
    }
    return &result, nil
}
func SerializeWellKnownListName_v2(values []WellKnownListName_v2) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
