package models
import (
    "errors"
)
// 
type AuthenticationStrengthResult int

const (
    NOTSET_AUTHENTICATIONSTRENGTHRESULT AuthenticationStrengthResult = iota
    SKIPPEDFORPROOFUP_AUTHENTICATIONSTRENGTHRESULT
    SATISFIED_AUTHENTICATIONSTRENGTHRESULT
    SINGLECHALLENGEREQUIRED_AUTHENTICATIONSTRENGTHRESULT
    MULTIPLECHALLENGESREQUIRED_AUTHENTICATIONSTRENGTHRESULT
    SINGLEREGISTRATIONREQUIRED_AUTHENTICATIONSTRENGTHRESULT
    MULTIPLEREGISTRATIONSREQUIRED_AUTHENTICATIONSTRENGTHRESULT
    CANNOTSATISFYDUETOCOMBINATIONCONFIGURATION_AUTHENTICATIONSTRENGTHRESULT
    CANNOTSATISFY_AUTHENTICATIONSTRENGTHRESULT
    UNKNOWNFUTUREVALUE_AUTHENTICATIONSTRENGTHRESULT
)

func (i AuthenticationStrengthResult) String() string {
    return []string{"notSet", "skippedForProofUp", "satisfied", "singleChallengeRequired", "multipleChallengesRequired", "singleRegistrationRequired", "multipleRegistrationsRequired", "cannotSatisfyDueToCombinationConfiguration", "cannotSatisfy", "unknownFutureValue"}[i]
}
func ParseAuthenticationStrengthResult(v string) (any, error) {
    result := NOTSET_AUTHENTICATIONSTRENGTHRESULT
    switch v {
        case "notSet":
            result = NOTSET_AUTHENTICATIONSTRENGTHRESULT
        case "skippedForProofUp":
            result = SKIPPEDFORPROOFUP_AUTHENTICATIONSTRENGTHRESULT
        case "satisfied":
            result = SATISFIED_AUTHENTICATIONSTRENGTHRESULT
        case "singleChallengeRequired":
            result = SINGLECHALLENGEREQUIRED_AUTHENTICATIONSTRENGTHRESULT
        case "multipleChallengesRequired":
            result = MULTIPLECHALLENGESREQUIRED_AUTHENTICATIONSTRENGTHRESULT
        case "singleRegistrationRequired":
            result = SINGLEREGISTRATIONREQUIRED_AUTHENTICATIONSTRENGTHRESULT
        case "multipleRegistrationsRequired":
            result = MULTIPLEREGISTRATIONSREQUIRED_AUTHENTICATIONSTRENGTHRESULT
        case "cannotSatisfyDueToCombinationConfiguration":
            result = CANNOTSATISFYDUETOCOMBINATIONCONFIGURATION_AUTHENTICATIONSTRENGTHRESULT
        case "cannotSatisfy":
            result = CANNOTSATISFY_AUTHENTICATIONSTRENGTHRESULT
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_AUTHENTICATIONSTRENGTHRESULT
        default:
            return 0, errors.New("Unknown AuthenticationStrengthResult value: " + v)
    }
    return &result, nil
}
func SerializeAuthenticationStrengthResult(values []AuthenticationStrengthResult) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
