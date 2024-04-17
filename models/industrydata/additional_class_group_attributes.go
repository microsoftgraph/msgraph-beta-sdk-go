package industrydata
import (
    "errors"
)
type AdditionalClassGroupAttributes int

const (
    COURSETITLE_ADDITIONALCLASSGROUPATTRIBUTES AdditionalClassGroupAttributes = iota
    COURSECODE_ADDITIONALCLASSGROUPATTRIBUTES
    COURSESUBJECT_ADDITIONALCLASSGROUPATTRIBUTES
    COURSEGRADELEVEL_ADDITIONALCLASSGROUPATTRIBUTES
    COURSEEXTERNALID_ADDITIONALCLASSGROUPATTRIBUTES
    ACADEMICSESSIONTITLE_ADDITIONALCLASSGROUPATTRIBUTES
    ACADEMICSESSIONEXTERNALID_ADDITIONALCLASSGROUPATTRIBUTES
    CLASSCODE_ADDITIONALCLASSGROUPATTRIBUTES
    UNKNOWNFUTUREVALUE_ADDITIONALCLASSGROUPATTRIBUTES
)

func (i AdditionalClassGroupAttributes) String() string {
    return []string{"courseTitle", "courseCode", "courseSubject", "courseGradeLevel", "courseExternalId", "academicSessionTitle", "academicSessionExternalId", "classCode", "unknownFutureValue"}[i]
}
func ParseAdditionalClassGroupAttributes(v string) (any, error) {
    result := COURSETITLE_ADDITIONALCLASSGROUPATTRIBUTES
    switch v {
        case "courseTitle":
            result = COURSETITLE_ADDITIONALCLASSGROUPATTRIBUTES
        case "courseCode":
            result = COURSECODE_ADDITIONALCLASSGROUPATTRIBUTES
        case "courseSubject":
            result = COURSESUBJECT_ADDITIONALCLASSGROUPATTRIBUTES
        case "courseGradeLevel":
            result = COURSEGRADELEVEL_ADDITIONALCLASSGROUPATTRIBUTES
        case "courseExternalId":
            result = COURSEEXTERNALID_ADDITIONALCLASSGROUPATTRIBUTES
        case "academicSessionTitle":
            result = ACADEMICSESSIONTITLE_ADDITIONALCLASSGROUPATTRIBUTES
        case "academicSessionExternalId":
            result = ACADEMICSESSIONEXTERNALID_ADDITIONALCLASSGROUPATTRIBUTES
        case "classCode":
            result = CLASSCODE_ADDITIONALCLASSGROUPATTRIBUTES
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ADDITIONALCLASSGROUPATTRIBUTES
        default:
            return 0, errors.New("Unknown AdditionalClassGroupAttributes value: " + v)
    }
    return &result, nil
}
func SerializeAdditionalClassGroupAttributes(values []AdditionalClassGroupAttributes) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AdditionalClassGroupAttributes) isMultiValue() bool {
    return false
}
