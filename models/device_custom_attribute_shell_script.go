package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceCustomAttributeShellScript represents a custom attribute script for macOS.
type DeviceCustomAttributeShellScript struct {
    Entity
    // The list of group assignments for the device management script.
    assignments []DeviceManagementScriptAssignmentable
    // The date and time the device management script was created. This property is read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The name of the custom attribute.
    customAttributeName *string
    // Represents the expected type for a macOS custom attribute script value.
    customAttributeType *DeviceCustomAttributeValueType
    // Optional description for the device management script.
    description *string
    // List of run states for this script across all devices.
    deviceRunStates []DeviceManagementScriptDeviceStateable
    // Name of the device management script.
    displayName *string
    // Script file name.
    fileName *string
    // The list of group assignments for the device management script.
    groupAssignments []DeviceManagementScriptGroupAssignmentable
    // The date and time the device management script was last modified. This property is read-only.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // List of Scope Tag IDs for this PowerShellScript instance.
    roleScopeTagIds []string
    // Indicates the type of execution context the app runs in.
    runAsAccount *RunAsAccountType
    // Run summary for device management script.
    runSummary DeviceManagementScriptRunSummaryable
    // The script content.
    scriptContent []byte
    // List of run states for this script across all users.
    userRunStates []DeviceManagementScriptUserStateable
}
// NewDeviceCustomAttributeShellScript instantiates a new deviceCustomAttributeShellScript and sets the default values.
func NewDeviceCustomAttributeShellScript()(*DeviceCustomAttributeShellScript) {
    m := &DeviceCustomAttributeShellScript{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.deviceCustomAttributeShellScript";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceCustomAttributeShellScriptFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceCustomAttributeShellScriptFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceCustomAttributeShellScript(), nil
}
// GetAssignments gets the assignments property value. The list of group assignments for the device management script.
func (m *DeviceCustomAttributeShellScript) GetAssignments()([]DeviceManagementScriptAssignmentable) {
    return m.assignments
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time the device management script was created. This property is read-only.
func (m *DeviceCustomAttributeShellScript) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetCustomAttributeName gets the customAttributeName property value. The name of the custom attribute.
func (m *DeviceCustomAttributeShellScript) GetCustomAttributeName()(*string) {
    return m.customAttributeName
}
// GetCustomAttributeType gets the customAttributeType property value. Represents the expected type for a macOS custom attribute script value.
func (m *DeviceCustomAttributeShellScript) GetCustomAttributeType()(*DeviceCustomAttributeValueType) {
    return m.customAttributeType
}
// GetDescription gets the description property value. Optional description for the device management script.
func (m *DeviceCustomAttributeShellScript) GetDescription()(*string) {
    return m.description
}
// GetDeviceRunStates gets the deviceRunStates property value. List of run states for this script across all devices.
func (m *DeviceCustomAttributeShellScript) GetDeviceRunStates()([]DeviceManagementScriptDeviceStateable) {
    return m.deviceRunStates
}
// GetDisplayName gets the displayName property value. Name of the device management script.
func (m *DeviceCustomAttributeShellScript) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceCustomAttributeShellScript) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceManagementScriptAssignmentFromDiscriminatorValue , m.SetAssignments)
    res["createdDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetCreatedDateTime)
    res["customAttributeName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCustomAttributeName)
    res["customAttributeType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseDeviceCustomAttributeValueType , m.SetCustomAttributeType)
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["deviceRunStates"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceManagementScriptDeviceStateFromDiscriminatorValue , m.SetDeviceRunStates)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["fileName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetFileName)
    res["groupAssignments"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceManagementScriptGroupAssignmentFromDiscriminatorValue , m.SetGroupAssignments)
    res["lastModifiedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastModifiedDateTime)
    res["roleScopeTagIds"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetRoleScopeTagIds)
    res["runAsAccount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseRunAsAccountType , m.SetRunAsAccount)
    res["runSummary"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateDeviceManagementScriptRunSummaryFromDiscriminatorValue , m.SetRunSummary)
    res["scriptContent"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetByteArrayValue(m.SetScriptContent)
    res["userRunStates"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceManagementScriptUserStateFromDiscriminatorValue , m.SetUserRunStates)
    return res
}
// GetFileName gets the fileName property value. Script file name.
func (m *DeviceCustomAttributeShellScript) GetFileName()(*string) {
    return m.fileName
}
// GetGroupAssignments gets the groupAssignments property value. The list of group assignments for the device management script.
func (m *DeviceCustomAttributeShellScript) GetGroupAssignments()([]DeviceManagementScriptGroupAssignmentable) {
    return m.groupAssignments
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time the device management script was last modified. This property is read-only.
func (m *DeviceCustomAttributeShellScript) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. List of Scope Tag IDs for this PowerShellScript instance.
func (m *DeviceCustomAttributeShellScript) GetRoleScopeTagIds()([]string) {
    return m.roleScopeTagIds
}
// GetRunAsAccount gets the runAsAccount property value. Indicates the type of execution context the app runs in.
func (m *DeviceCustomAttributeShellScript) GetRunAsAccount()(*RunAsAccountType) {
    return m.runAsAccount
}
// GetRunSummary gets the runSummary property value. Run summary for device management script.
func (m *DeviceCustomAttributeShellScript) GetRunSummary()(DeviceManagementScriptRunSummaryable) {
    return m.runSummary
}
// GetScriptContent gets the scriptContent property value. The script content.
func (m *DeviceCustomAttributeShellScript) GetScriptContent()([]byte) {
    return m.scriptContent
}
// GetUserRunStates gets the userRunStates property value. List of run states for this script across all users.
func (m *DeviceCustomAttributeShellScript) GetUserRunStates()([]DeviceManagementScriptUserStateable) {
    return m.userRunStates
}
// Serialize serializes information the current object
func (m *DeviceCustomAttributeShellScript) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssignments() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAssignments())
        err = writer.WriteCollectionOfObjectValues("assignments", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customAttributeName", m.GetCustomAttributeName())
        if err != nil {
            return err
        }
    }
    if m.GetCustomAttributeType() != nil {
        cast := (*m.GetCustomAttributeType()).String()
        err = writer.WriteStringValue("customAttributeType", &cast)
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
    if m.GetDeviceRunStates() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDeviceRunStates())
        err = writer.WriteCollectionOfObjectValues("deviceRunStates", cast)
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
    {
        err = writer.WriteStringValue("fileName", m.GetFileName())
        if err != nil {
            return err
        }
    }
    if m.GetGroupAssignments() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetGroupAssignments())
        err = writer.WriteCollectionOfObjectValues("groupAssignments", cast)
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
    if m.GetRunAsAccount() != nil {
        cast := (*m.GetRunAsAccount()).String()
        err = writer.WriteStringValue("runAsAccount", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("runSummary", m.GetRunSummary())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("scriptContent", m.GetScriptContent())
        if err != nil {
            return err
        }
    }
    if m.GetUserRunStates() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetUserRunStates())
        err = writer.WriteCollectionOfObjectValues("userRunStates", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignments sets the assignments property value. The list of group assignments for the device management script.
func (m *DeviceCustomAttributeShellScript) SetAssignments(value []DeviceManagementScriptAssignmentable)() {
    m.assignments = value
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time the device management script was created. This property is read-only.
func (m *DeviceCustomAttributeShellScript) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetCustomAttributeName sets the customAttributeName property value. The name of the custom attribute.
func (m *DeviceCustomAttributeShellScript) SetCustomAttributeName(value *string)() {
    m.customAttributeName = value
}
// SetCustomAttributeType sets the customAttributeType property value. Represents the expected type for a macOS custom attribute script value.
func (m *DeviceCustomAttributeShellScript) SetCustomAttributeType(value *DeviceCustomAttributeValueType)() {
    m.customAttributeType = value
}
// SetDescription sets the description property value. Optional description for the device management script.
func (m *DeviceCustomAttributeShellScript) SetDescription(value *string)() {
    m.description = value
}
// SetDeviceRunStates sets the deviceRunStates property value. List of run states for this script across all devices.
func (m *DeviceCustomAttributeShellScript) SetDeviceRunStates(value []DeviceManagementScriptDeviceStateable)() {
    m.deviceRunStates = value
}
// SetDisplayName sets the displayName property value. Name of the device management script.
func (m *DeviceCustomAttributeShellScript) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetFileName sets the fileName property value. Script file name.
func (m *DeviceCustomAttributeShellScript) SetFileName(value *string)() {
    m.fileName = value
}
// SetGroupAssignments sets the groupAssignments property value. The list of group assignments for the device management script.
func (m *DeviceCustomAttributeShellScript) SetGroupAssignments(value []DeviceManagementScriptGroupAssignmentable)() {
    m.groupAssignments = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time the device management script was last modified. This property is read-only.
func (m *DeviceCustomAttributeShellScript) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. List of Scope Tag IDs for this PowerShellScript instance.
func (m *DeviceCustomAttributeShellScript) SetRoleScopeTagIds(value []string)() {
    m.roleScopeTagIds = value
}
// SetRunAsAccount sets the runAsAccount property value. Indicates the type of execution context the app runs in.
func (m *DeviceCustomAttributeShellScript) SetRunAsAccount(value *RunAsAccountType)() {
    m.runAsAccount = value
}
// SetRunSummary sets the runSummary property value. Run summary for device management script.
func (m *DeviceCustomAttributeShellScript) SetRunSummary(value DeviceManagementScriptRunSummaryable)() {
    m.runSummary = value
}
// SetScriptContent sets the scriptContent property value. The script content.
func (m *DeviceCustomAttributeShellScript) SetScriptContent(value []byte)() {
    m.scriptContent = value
}
// SetUserRunStates sets the userRunStates property value. List of run states for this script across all users.
func (m *DeviceCustomAttributeShellScript) SetUserRunStates(value []DeviceManagementScriptUserStateable)() {
    m.userRunStates = value
}
