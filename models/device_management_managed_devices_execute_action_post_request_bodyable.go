package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementManagedDevicesExecuteActionPostRequestBodyable 
type DeviceManagementManagedDevicesExecuteActionPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActionName()(*ManagedDeviceRemoteAction)
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
    SetActionName(value *ManagedDeviceRemoteAction)()
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
