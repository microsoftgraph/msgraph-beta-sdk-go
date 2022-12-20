package models
import (
    "errors"
)
// Provides operations to manage the collection of activityStatistics entities.
type SectionEmphasisType int

const (
    NONE_SECTIONEMPHASISTYPE SectionEmphasisType = iota
    NETURAL_SECTIONEMPHASISTYPE
    SOFT_SECTIONEMPHASISTYPE
    STRONG_SECTIONEMPHASISTYPE
    UNKNOWNFUTUREVALUE_SECTIONEMPHASISTYPE
)

func (i SectionEmphasisType) String() string {
    return []string{"none", "netural", "soft", "strong", "unknownFutureValue"}[i]
}
func ParseSectionEmphasisType(v string) (interface{}, error) {
    result := NONE_SECTIONEMPHASISTYPE
    switch v {
        case "none":
            result = NONE_SECTIONEMPHASISTYPE
        case "netural":
            result = NETURAL_SECTIONEMPHASISTYPE
        case "soft":
            result = SOFT_SECTIONEMPHASISTYPE
        case "strong":
            result = STRONG_SECTIONEMPHASISTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_SECTIONEMPHASISTYPE
        default:
            return 0, errors.New("Unknown SectionEmphasisType value: " + v)
    }
    return &result, nil
}
func SerializeSectionEmphasisType(values []SectionEmphasisType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
