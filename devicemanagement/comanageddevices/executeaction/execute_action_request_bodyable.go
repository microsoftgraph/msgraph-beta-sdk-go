package executeaction

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// ExecuteActionRequestBodyable 
type ExecuteActionRequestBodyable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetActionName()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagedDeviceRemoteAction)
    GetCarrierUrl()(*string)
    GetDeprovisionReason()(*string)
    GetDeviceIds()([]string)
    GetDeviceName()(*string)
    GetKeepEnrollmentData()(*bool)
    GetKeepUserData()(*bool)
    GetNotificationBody()(*string)
    GetNotificationTitle()(*string)
    GetOrganizationalUnitPath()(*string)
    GetPersistEsimDataPlan()(*bool)
    SetActionName(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagedDeviceRemoteAction)()
    SetCarrierUrl(value *string)()
    SetDeprovisionReason(value *string)()
    SetDeviceIds(value []string)()
    SetDeviceName(value *string)()
    SetKeepEnrollmentData(value *bool)()
    SetKeepUserData(value *bool)()
    SetNotificationBody(value *string)()
    SetNotificationTitle(value *string)()
    SetOrganizationalUnitPath(value *string)()
    SetPersistEsimDataPlan(value *bool)()
}
