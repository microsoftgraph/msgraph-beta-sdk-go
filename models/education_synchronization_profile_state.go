package models
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
    return []string{"deleting", "deletionFailed", "provisioningFailed", "provisioned", "provisioning", "unknownFutureValue"}[i]
}
func ParseEducationSynchronizationProfileState(v string) (any, error) {
    result := DELETING_EDUCATIONSYNCHRONIZATIONPROFILESTATE
    switch v {
        case "deleting":
            result = DELETING_EDUCATIONSYNCHRONIZATIONPROFILESTATE
        case "deletionFailed":
            result = DELETIONFAILED_EDUCATIONSYNCHRONIZATIONPROFILESTATE
        case "provisioningFailed":
            result = PROVISIONINGFAILED_EDUCATIONSYNCHRONIZATIONPROFILESTATE
        case "provisioned":
            result = PROVISIONED_EDUCATIONSYNCHRONIZATIONPROFILESTATE
        case "provisioning":
            result = PROVISIONING_EDUCATIONSYNCHRONIZATIONPROFILESTATE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_EDUCATIONSYNCHRONIZATIONPROFILESTATE
        default:
            return nil, nil
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
func (i EducationSynchronizationProfileState) isMultiValue() bool {
    return false
}
