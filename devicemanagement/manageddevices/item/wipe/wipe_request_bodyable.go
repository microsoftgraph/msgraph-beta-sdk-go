package wipe

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WipeRequestBodyable 
type WipeRequestBodyable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetKeepEnrollmentData()(*bool)
    GetKeepUserData()(*bool)
    GetMacOsUnlockCode()(*string)
    GetPersistEsimDataPlan()(*bool)
    GetUseProtectedWipe()(*bool)
    SetKeepEnrollmentData(value *bool)()
    SetKeepUserData(value *bool)()
    SetMacOsUnlockCode(value *string)()
    SetPersistEsimDataPlan(value *bool)()
    SetUseProtectedWipe(value *bool)()
}
