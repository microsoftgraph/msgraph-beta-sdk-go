package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// EducationSynchronizationProfileable 
type EducationSynchronizationProfileable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDataProvider()(EducationSynchronizationDataProviderable)
    GetDisplayName()(*string)
    GetErrors()([]EducationSynchronizationErrorable)
    GetExpirationDate()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)
    GetHandleSpecialCharacterConstraint()(*bool)
    GetIdentitySynchronizationConfiguration()(EducationIdentitySynchronizationConfigurationable)
    GetLicensesToAssign()([]EducationSynchronizationLicenseAssignmentable)
    GetProfileStatus()(EducationSynchronizationProfileStatusable)
    GetState()(*EducationSynchronizationProfileState)
    SetDataProvider(value EducationSynchronizationDataProviderable)()
    SetDisplayName(value *string)()
    SetErrors(value []EducationSynchronizationErrorable)()
    SetExpirationDate(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)()
    SetHandleSpecialCharacterConstraint(value *bool)()
    SetIdentitySynchronizationConfiguration(value EducationIdentitySynchronizationConfigurationable)()
    SetLicensesToAssign(value []EducationSynchronizationLicenseAssignmentable)()
    SetProfileStatus(value EducationSynchronizationProfileStatusable)()
    SetState(value *EducationSynchronizationProfileState)()
}
