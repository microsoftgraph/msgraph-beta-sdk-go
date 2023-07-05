package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsDriverUpdateProfile windows Driver Update Profile
type WindowsDriverUpdateProfile struct {
    Entity
}
// NewWindowsDriverUpdateProfile instantiates a new windowsDriverUpdateProfile and sets the default values.
func NewWindowsDriverUpdateProfile()(*WindowsDriverUpdateProfile) {
    m := &WindowsDriverUpdateProfile{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWindowsDriverUpdateProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsDriverUpdateProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsDriverUpdateProfile(), nil
}
// GetApprovalType gets the approvalType property value. An enum type to represent approval type of a driver update profile.
func (m *WindowsDriverUpdateProfile) GetApprovalType()(*DriverUpdateProfileApprovalType) {
    val, err := m.GetBackingStore().Get("approvalType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DriverUpdateProfileApprovalType)
    }
    return nil
}
// GetAssignments gets the assignments property value. The list of group assignments of the profile.
func (m *WindowsDriverUpdateProfile) GetAssignments()([]WindowsDriverUpdateProfileAssignmentable) {
    val, err := m.GetBackingStore().Get("assignments")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]WindowsDriverUpdateProfileAssignmentable)
    }
    return nil
}
// GetCreatedDateTime gets the createdDateTime property value. The date time that the profile was created.
func (m *WindowsDriverUpdateProfile) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetDeploymentDeferralInDays gets the deploymentDeferralInDays property value. Deployment deferral settings in days, only applicable when ApprovalType is set to automatic approval.
func (m *WindowsDriverUpdateProfile) GetDeploymentDeferralInDays()(*int32) {
    val, err := m.GetBackingStore().Get("deploymentDeferralInDays")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetDescription gets the description property value. The description of the profile which is specified by the user.
func (m *WindowsDriverUpdateProfile) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceReporting gets the deviceReporting property value. Number of devices reporting for this profile
func (m *WindowsDriverUpdateProfile) GetDeviceReporting()(*int32) {
    val, err := m.GetBackingStore().Get("deviceReporting")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The display name for the profile.
func (m *WindowsDriverUpdateProfile) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDriverInventories gets the driverInventories property value. Driver inventories for this profile.
func (m *WindowsDriverUpdateProfile) GetDriverInventories()([]WindowsDriverUpdateInventoryable) {
    val, err := m.GetBackingStore().Get("driverInventories")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]WindowsDriverUpdateInventoryable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsDriverUpdateProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["approvalType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDriverUpdateProfileApprovalType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApprovalType(val.(*DriverUpdateProfileApprovalType))
        }
        return nil
    }
    res["assignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsDriverUpdateProfileAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsDriverUpdateProfileAssignmentable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(WindowsDriverUpdateProfileAssignmentable)
                }
            }
            m.SetAssignments(res)
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["deploymentDeferralInDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeploymentDeferralInDays(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["deviceReporting"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceReporting(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["driverInventories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsDriverUpdateInventoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsDriverUpdateInventoryable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(WindowsDriverUpdateInventoryable)
                }
            }
            m.SetDriverInventories(res)
        }
        return nil
    }
    res["inventorySyncStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindowsDriverUpdateProfileInventorySyncStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInventorySyncStatus(val.(WindowsDriverUpdateProfileInventorySyncStatusable))
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["newUpdates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNewUpdates(val)
        }
        return nil
    }
    res["roleScopeTagIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetRoleScopeTagIds(res)
        }
        return nil
    }
    return res
}
// GetInventorySyncStatus gets the inventorySyncStatus property value. Driver inventory sync status for this profile.
func (m *WindowsDriverUpdateProfile) GetInventorySyncStatus()(WindowsDriverUpdateProfileInventorySyncStatusable) {
    val, err := m.GetBackingStore().Get("inventorySyncStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(WindowsDriverUpdateProfileInventorySyncStatusable)
    }
    return nil
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date time that the profile was last modified.
func (m *WindowsDriverUpdateProfile) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastModifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetNewUpdates gets the newUpdates property value. Number of new driver updates available for this profile.
func (m *WindowsDriverUpdateProfile) GetNewUpdates()(*int32) {
    val, err := m.GetBackingStore().Get("newUpdates")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. List of Scope Tags for this Driver Update entity.
func (m *WindowsDriverUpdateProfile) GetRoleScopeTagIds()([]string) {
    val, err := m.GetBackingStore().Get("roleScopeTagIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *WindowsDriverUpdateProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetApprovalType() != nil {
        cast := (*m.GetApprovalType()).String()
        err = writer.WriteStringValue("approvalType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignments()))
        for i, v := range m.GetAssignments() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("assignments", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("deploymentDeferralInDays", m.GetDeploymentDeferralInDays())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("deviceReporting", m.GetDeviceReporting())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetDriverInventories() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDriverInventories()))
        for i, v := range m.GetDriverInventories() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("driverInventories", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("inventorySyncStatus", m.GetInventorySyncStatus())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("newUpdates", m.GetNewUpdates())
        if err != nil {
            return err
        }
    }
    if m.GetRoleScopeTagIds() != nil {
        err = writer.WriteCollectionOfStringValues("roleScopeTagIds", m.GetRoleScopeTagIds())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApprovalType sets the approvalType property value. An enum type to represent approval type of a driver update profile.
func (m *WindowsDriverUpdateProfile) SetApprovalType(value *DriverUpdateProfileApprovalType)() {
    err := m.GetBackingStore().Set("approvalType", value)
    if err != nil {
        panic(err)
    }
}
// SetAssignments sets the assignments property value. The list of group assignments of the profile.
func (m *WindowsDriverUpdateProfile) SetAssignments(value []WindowsDriverUpdateProfileAssignmentable)() {
    err := m.GetBackingStore().Set("assignments", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The date time that the profile was created.
func (m *WindowsDriverUpdateProfile) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetDeploymentDeferralInDays sets the deploymentDeferralInDays property value. Deployment deferral settings in days, only applicable when ApprovalType is set to automatic approval.
func (m *WindowsDriverUpdateProfile) SetDeploymentDeferralInDays(value *int32)() {
    err := m.GetBackingStore().Set("deploymentDeferralInDays", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. The description of the profile which is specified by the user.
func (m *WindowsDriverUpdateProfile) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceReporting sets the deviceReporting property value. Number of devices reporting for this profile
func (m *WindowsDriverUpdateProfile) SetDeviceReporting(value *int32)() {
    err := m.GetBackingStore().Set("deviceReporting", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The display name for the profile.
func (m *WindowsDriverUpdateProfile) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetDriverInventories sets the driverInventories property value. Driver inventories for this profile.
func (m *WindowsDriverUpdateProfile) SetDriverInventories(value []WindowsDriverUpdateInventoryable)() {
    err := m.GetBackingStore().Set("driverInventories", value)
    if err != nil {
        panic(err)
    }
}
// SetInventorySyncStatus sets the inventorySyncStatus property value. Driver inventory sync status for this profile.
func (m *WindowsDriverUpdateProfile) SetInventorySyncStatus(value WindowsDriverUpdateProfileInventorySyncStatusable)() {
    err := m.GetBackingStore().Set("inventorySyncStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date time that the profile was last modified.
func (m *WindowsDriverUpdateProfile) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetNewUpdates sets the newUpdates property value. Number of new driver updates available for this profile.
func (m *WindowsDriverUpdateProfile) SetNewUpdates(value *int32)() {
    err := m.GetBackingStore().Set("newUpdates", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. List of Scope Tags for this Driver Update entity.
func (m *WindowsDriverUpdateProfile) SetRoleScopeTagIds(value []string)() {
    err := m.GetBackingStore().Set("roleScopeTagIds", value)
    if err != nil {
        panic(err)
    }
}
// WindowsDriverUpdateProfileable 
type WindowsDriverUpdateProfileable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApprovalType()(*DriverUpdateProfileApprovalType)
    GetAssignments()([]WindowsDriverUpdateProfileAssignmentable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDeploymentDeferralInDays()(*int32)
    GetDescription()(*string)
    GetDeviceReporting()(*int32)
    GetDisplayName()(*string)
    GetDriverInventories()([]WindowsDriverUpdateInventoryable)
    GetInventorySyncStatus()(WindowsDriverUpdateProfileInventorySyncStatusable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetNewUpdates()(*int32)
    GetRoleScopeTagIds()([]string)
    SetApprovalType(value *DriverUpdateProfileApprovalType)()
    SetAssignments(value []WindowsDriverUpdateProfileAssignmentable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDeploymentDeferralInDays(value *int32)()
    SetDescription(value *string)()
    SetDeviceReporting(value *int32)()
    SetDisplayName(value *string)()
    SetDriverInventories(value []WindowsDriverUpdateInventoryable)()
    SetInventorySyncStatus(value WindowsDriverUpdateProfileInventorySyncStatusable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetNewUpdates(value *int32)()
    SetRoleScopeTagIds(value []string)()
}
