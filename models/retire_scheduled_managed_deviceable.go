package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RetireScheduledManagedDeviceable 
type RetireScheduledManagedDeviceable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetComplianceState()(*ComplianceStatus)
    GetDeviceCompliancePolicyId()(*string)
    GetDeviceCompliancePolicyName()(*string)
    GetDeviceType()(*DeviceType)
    GetId()(*string)
    GetManagedDeviceId()(*string)
    GetManagedDeviceName()(*string)
    GetManagementAgent()(*ManagementAgentType)
    GetOdataType()(*string)
    GetOwnerType()(*ManagedDeviceOwnerType)
    GetRetireAfterDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRoleScopeTagIds()([]string)
    SetComplianceState(value *ComplianceStatus)()
    SetDeviceCompliancePolicyId(value *string)()
    SetDeviceCompliancePolicyName(value *string)()
    SetDeviceType(value *DeviceType)()
    SetId(value *string)()
    SetManagedDeviceId(value *string)()
    SetManagedDeviceName(value *string)()
    SetManagementAgent(value *ManagementAgentType)()
    SetOdataType(value *string)()
    SetOwnerType(value *ManagedDeviceOwnerType)()
    SetRetireAfterDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRoleScopeTagIds(value []string)()
}
