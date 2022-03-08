package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ConditionalAccessConditionSetable 
type ConditionalAccessConditionSetable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetApplications()(ConditionalAccessApplicationsable)
    GetClientApplications()(ConditionalAccessClientApplicationsable)
    GetClientAppTypes()([]ConditionalAccessClientApp)
    GetDevices()(ConditionalAccessDevicesable)
    GetDeviceStates()(ConditionalAccessDeviceStatesable)
    GetLocations()(ConditionalAccessLocationsable)
    GetPlatforms()(ConditionalAccessPlatformsable)
    GetServicePrincipalRiskLevels()([]RiskLevel)
    GetSignInRiskLevels()([]RiskLevel)
    GetUserRiskLevels()([]RiskLevel)
    GetUsers()(ConditionalAccessUsersable)
    SetApplications(value ConditionalAccessApplicationsable)()
    SetClientApplications(value ConditionalAccessClientApplicationsable)()
    SetClientAppTypes(value []ConditionalAccessClientApp)()
    SetDevices(value ConditionalAccessDevicesable)()
    SetDeviceStates(value ConditionalAccessDeviceStatesable)()
    SetLocations(value ConditionalAccessLocationsable)()
    SetPlatforms(value ConditionalAccessPlatformsable)()
    SetServicePrincipalRiskLevels(value []RiskLevel)()
    SetSignInRiskLevels(value []RiskLevel)()
    SetUserRiskLevels(value []RiskLevel)()
    SetUsers(value ConditionalAccessUsersable)()
}
