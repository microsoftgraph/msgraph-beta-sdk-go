package graph
import (
    "strings"
    "errors"
)
// 
type EducationSynchronizationProfileState int

const (
    DELETING_EDUCATIONSYNCHRONIZATIONPROFILESTATE EducationSynchronizationProfileState = iota
    DELETIONFAILED_EDUCATIONSYNCHRONIZATIONPROFILESTATE
    PROVISIONINGFAILED_EDUCATIONSYNCHRONIZATIONPROFILESTATE
    PROVISIONED_EDUCATIONSYNCHRONIZATIONPROFILESTATE
    PROVISIONING_EDUCATIONSYNCHRONIZATIONPROFILESTATE
    UNKNOWNFUTUREVALUE_EDUCATIONSYNCHRONIZATIONPROFILESTATE
)

func (i EducationSynchronizationProfileState) String() string {
    return []string{"DELETING", "DELETIONFAILED", "PROVISIONINGFAILED", "PROVISIONED", "PROVISIONING", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseEducationSynchronizationProfileState(v string) (interface{}, error) {
    result := DELETING_EDUCATIONSYNCHRONIZATIONPROFILESTATE
    switch strings.ToUpper(v) {
        case "DELETING":
            result = DELETING_EDUCATIONSYNCHRONIZATIONPROFILESTATE
        case "DELETIONFAILED":
            result = DELETIONFAILED_EDUCATIONSYNCHRONIZATIONPROFILESTATE
        case "PROVISIONINGFAILED":
            result = PROVISIONINGFAILED_EDUCATIONSYNCHRONIZATIONPROFILESTATE
        case "PROVISIONED":
            result = PROVISIONED_EDUCATIONSYNCHRONIZATIONPROFILESTATE
        case "PROVISIONING":
            result = PROVISIONING_EDUCATIONSYNCHRONIZATIONPROFILESTATE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_EDUCATIONSYNCHRONIZATIONPROFILESTATE
        default:
            return 0, errors.New("Unknown EducationSynchronizationProfileState value: " + v)
    }
    return &result, nil
}
func SerializeEducationSynchronizationProfileState(values []EducationSynchronizationProfileState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
