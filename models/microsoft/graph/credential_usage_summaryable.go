package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CredentialUsageSummaryable 
type CredentialUsageSummaryable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAuthMethod()(*UsageAuthMethod)
    GetFailureActivityCount()(*int32)
    GetFeature()(*FeatureType)
    GetSuccessfulActivityCount()(*int32)
    SetAuthMethod(value *UsageAuthMethod)()
    SetFailureActivityCount(value *int32)()
    SetFeature(value *FeatureType)()
    SetSuccessfulActivityCount(value *int32)()
}
