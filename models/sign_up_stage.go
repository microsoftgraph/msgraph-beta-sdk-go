package models
type SignUpStage int

const (
    CREDENTIALCOLLECTION_SIGNUPSTAGE SignUpStage = iota
    CREDENTIALVALIDATION_SIGNUPSTAGE
    CREDENTIALFEDERATION_SIGNUPSTAGE
    CONSENT_SIGNUPSTAGE
    ATTRIBUTECOLLECTIONANDVALIDATION_SIGNUPSTAGE
    USERCREATION_SIGNUPSTAGE
    TENANTCONSENT_SIGNUPSTAGE
    UNKNOWNFUTUREVALUE_SIGNUPSTAGE
)

func (i SignUpStage) String() string {
    return []string{"credentialCollection", "credentialValidation", "credentialFederation", "consent", "attributeCollectionAndValidation", "userCreation", "tenantConsent", "unknownFutureValue"}[i]
}
func ParseSignUpStage(v string) (any, error) {
    result := CREDENTIALCOLLECTION_SIGNUPSTAGE
    switch v {
        case "credentialCollection":
            result = CREDENTIALCOLLECTION_SIGNUPSTAGE
        case "credentialValidation":
            result = CREDENTIALVALIDATION_SIGNUPSTAGE
        case "credentialFederation":
            result = CREDENTIALFEDERATION_SIGNUPSTAGE
        case "consent":
            result = CONSENT_SIGNUPSTAGE
        case "attributeCollectionAndValidation":
            result = ATTRIBUTECOLLECTIONANDVALIDATION_SIGNUPSTAGE
        case "userCreation":
            result = USERCREATION_SIGNUPSTAGE
        case "tenantConsent":
            result = TENANTCONSENT_SIGNUPSTAGE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_SIGNUPSTAGE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeSignUpStage(values []SignUpStage) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SignUpStage) isMultiValue() bool {
    return false
}
