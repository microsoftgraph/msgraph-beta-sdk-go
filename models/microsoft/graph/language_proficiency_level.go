package graph
import (
    "strings"
    "errors"
)
// 
type LanguageProficiencyLevel int

const (
    ELEMENTARY_LANGUAGEPROFICIENCYLEVEL LanguageProficiencyLevel = iota
    CONVERSATIONAL_LANGUAGEPROFICIENCYLEVEL
    LIMITEDWORKING_LANGUAGEPROFICIENCYLEVEL
    PROFESSIONALWORKING_LANGUAGEPROFICIENCYLEVEL
    FULLPROFESSIONAL_LANGUAGEPROFICIENCYLEVEL
    NATIVEORBILINGUAL_LANGUAGEPROFICIENCYLEVEL
    UNKNOWNFUTUREVALUE_LANGUAGEPROFICIENCYLEVEL
)

func (i LanguageProficiencyLevel) String() string {
    return []string{"ELEMENTARY", "CONVERSATIONAL", "LIMITEDWORKING", "PROFESSIONALWORKING", "FULLPROFESSIONAL", "NATIVEORBILINGUAL", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseLanguageProficiencyLevel(v string) (interface{}, error) {
    result := ELEMENTARY_LANGUAGEPROFICIENCYLEVEL
    switch strings.ToUpper(v) {
        case "ELEMENTARY":
            result = ELEMENTARY_LANGUAGEPROFICIENCYLEVEL
        case "CONVERSATIONAL":
            result = CONVERSATIONAL_LANGUAGEPROFICIENCYLEVEL
        case "LIMITEDWORKING":
            result = LIMITEDWORKING_LANGUAGEPROFICIENCYLEVEL
        case "PROFESSIONALWORKING":
            result = PROFESSIONALWORKING_LANGUAGEPROFICIENCYLEVEL
        case "FULLPROFESSIONAL":
            result = FULLPROFESSIONAL_LANGUAGEPROFICIENCYLEVEL
        case "NATIVEORBILINGUAL":
            result = NATIVEORBILINGUAL_LANGUAGEPROFICIENCYLEVEL
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_LANGUAGEPROFICIENCYLEVEL
        default:
            return 0, errors.New("Unknown LanguageProficiencyLevel value: " + v)
    }
    return &result, nil
}
func SerializeLanguageProficiencyLevel(values []LanguageProficiencyLevel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
