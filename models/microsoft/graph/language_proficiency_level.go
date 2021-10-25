package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "ELEMENTARY":
            return ELEMENTARY_LANGUAGEPROFICIENCYLEVEL, nil
        case "CONVERSATIONAL":
            return CONVERSATIONAL_LANGUAGEPROFICIENCYLEVEL, nil
        case "LIMITEDWORKING":
            return LIMITEDWORKING_LANGUAGEPROFICIENCYLEVEL, nil
        case "PROFESSIONALWORKING":
            return PROFESSIONALWORKING_LANGUAGEPROFICIENCYLEVEL, nil
        case "FULLPROFESSIONAL":
            return FULLPROFESSIONAL_LANGUAGEPROFICIENCYLEVEL, nil
        case "NATIVEORBILINGUAL":
            return NATIVEORBILINGUAL_LANGUAGEPROFICIENCYLEVEL, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_LANGUAGEPROFICIENCYLEVEL, nil
    }
    return 0, errors.New("Unknown LanguageProficiencyLevel value: " + v)
}
func SerializeLanguageProficiencyLevel(values []LanguageProficiencyLevel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
