package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsDriverUpdateProfile windows Driver Update Profile
type WindowsDriverUpdateProfile struct {
    Entity
    // An enum type to represent approval type of a driver update profile.
    approvalType *DriverUpdateProfileApprovalType
    // The list of group assignments of the profile.
    assignments []WindowsDriverUpdateProfileAssignmentable
    // The date time that the profile was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Deployment deferral settings in days, only applicable when ApprovalType is set to automatic approval.
    deploymentDeferralInDays *int32
    // The description of the profile which is specified by the user.
    description *string
    // Number of devices reporting for this profile
    deviceReporting *int32
    // The display name for the profile.
    displayName *string
    // Driver inventories for this profile.
    driverInventories []WindowsDriverUpdateInventoryable
    // Driver inventory sync status for this profile.
    inventorySyncStatus WindowsDriverUpdateProfileInventorySyncStatusable
    // The date time that the profile was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Number of new driver updates available for this profile.
    newUpdates *int32
    // List of Scope Tags for this Driver Update entity.
    roleScopeTagIds []string
}
// NewWindowsDriverUpdateProfile instantiates a new windowsDriverUpdateProfile and sets the default values.
func NewWindowsDriverUpdateProfile()(*WindowsDriverUpdateProfile) {
    m := &WindowsDriverUpdateProfile{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.windowsDriverUpdateProfile";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWindowsDriverUpdateProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsDriverUpdateProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsDriverUpdateProfile(), nil
}
// GetApprovalType gets the approvalType property value. An enum type to represent approval type of a driver update profile.
func (m *WindowsDriverUpdateProfile) GetApprovalType()(*DriverUpdateProfileApprovalType) {
    return m.approvalType
}
// GetAssignments gets the assignments property value. The list of group assignments of the profile.
func (m *WindowsDriverUpdateProfile) GetAssignments()([]WindowsDriverUpdateProfileAssignmentable) {
    return m.assignments
}
// GetCreatedDateTime gets the createdDateTime property value. The date time that the profile was created.
func (m *WindowsDriverUpdateProfile) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDeploymentDeferralInDays gets the deploymentDeferralInDays property value. Deployment deferral settings in days, only applicable when ApprovalType is set to automatic approval.
func (m *WindowsDriverUpdateProfile) GetDeploymentDeferralInDays()(*int32) {
    return m.deploymentDeferralInDays
}
// GetDescription gets the description property value. The description of the profile which is specified by the user.
func (m *WindowsDriverUpdateProfile) GetDescription()(*string) {
    return m.description
}
// GetDeviceReporting gets the deviceReporting property value. Number of devices reporting for this profile
func (m *WindowsDriverUpdateProfile) GetDeviceReporting()(*int32) {
    return m.deviceReporting
}
// GetDisplayName gets the displayName property value. The display name for the profile.
func (m *WindowsDriverUpdateProfile) GetDisplayName()(*string) {
    return m.displayName
}
// GetDriverInventories gets the driverInventories property value. Driver inventories for this profile.
func (m *WindowsDriverUpdateProfile) GetDriverInventories()([]WindowsDriverUpdateInventoryable) {
    return m.driverInventories
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsDriverUpdateProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["approvalType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseDriverUpdateProfileApprovalType , m.SetApprovalType)
    res["assignments"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateWindowsDriverUpdateProfileAssignmentFromDiscriminatorValue , m.SetAssignments)
    res["createdDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetCreatedDateTime)
    res["deploymentDeferralInDays"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetDeploymentDeferralInDays)
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["deviceReporting"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetDeviceReporting)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["driverInventories"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateWindowsDriverUpdateInventoryFromDiscriminatorValue , m.SetDriverInventories)
    res["inventorySyncStatus"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateWindowsDriverUpdateProfileInventorySyncStatusFromDiscriminatorValue , m.SetInventorySyncStatus)
    res["lastModifiedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastModifiedDateTime)
    res["newUpdates"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetNewUpdates)
    res["roleScopeTagIds"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetRoleScopeTagIds)
    return res
}
// GetInventorySyncStatus gets the inventorySyncStatus property value. Driver inventory sync status for this profile.
func (m *WindowsDriverUpdateProfile) GetInventorySyncStatus()(WindowsDriverUpdateProfileInventorySyncStatusable) {
    return m.inventorySyncStatus
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date time that the profile was last modified.
func (m *WindowsDriverUpdateProfile) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetNewUpdates gets the newUpdates property value. Number of new driver updates available for this profile.
func (m *WindowsDriverUpdateProfile) GetNewUpdates()(*int32) {
    return m.newUpdates
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. List of Scope Tags for this Driver Update entity.
func (m *WindowsDriverUpdateProfile) GetRoleScopeTagIds()([]string) {
    return m.roleScopeTagIds
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
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAssignments())
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
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDriverInventories())
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
    m.approvalType = value
}
// SetAssignments sets the assignments property value. The list of group assignments of the profile.
func (m *WindowsDriverUpdateProfile) SetAssignments(value []WindowsDriverUpdateProfileAssignmentable)() {
    m.assignments = value
}
// SetCreatedDateTime sets the createdDateTime property value. The date time that the profile was created.
func (m *WindowsDriverUpdateProfile) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDeploymentDeferralInDays sets the deploymentDeferralInDays property value. Deployment deferral settings in days, only applicable when ApprovalType is set to automatic approval.
func (m *WindowsDriverUpdateProfile) SetDeploymentDeferralInDays(value *int32)() {
    m.deploymentDeferralInDays = value
}
// SetDescription sets the description property value. The description of the profile which is specified by the user.
func (m *WindowsDriverUpdateProfile) SetDescription(value *string)() {
    m.description = value
}
// SetDeviceReporting sets the deviceReporting property value. Number of devices reporting for this profile
func (m *WindowsDriverUpdateProfile) SetDeviceReporting(value *int32)() {
    m.deviceReporting = value
}
// SetDisplayName sets the displayName property value. The display name for the profile.
func (m *WindowsDriverUpdateProfile) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetDriverInventories sets the driverInventories property value. Driver inventories for this profile.
func (m *WindowsDriverUpdateProfile) SetDriverInventories(value []WindowsDriverUpdateInventoryable)() {
    m.driverInventories = value
}
// SetInventorySyncStatus sets the inventorySyncStatus property value. Driver inventory sync status for this profile.
func (m *WindowsDriverUpdateProfile) SetInventorySyncStatus(value WindowsDriverUpdateProfileInventorySyncStatusable)() {
    m.inventorySyncStatus = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date time that the profile was last modified.
func (m *WindowsDriverUpdateProfile) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetNewUpdates sets the newUpdates property value. Number of new driver updates available for this profile.
func (m *WindowsDriverUpdateProfile) SetNewUpdates(value *int32)() {
    m.newUpdates = value
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. List of Scope Tags for this Driver Update entity.
func (m *WindowsDriverUpdateProfile) SetRoleScopeTagIds(value []string)() {
    m.roleScopeTagIds = value
}
