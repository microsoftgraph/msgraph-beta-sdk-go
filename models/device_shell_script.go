package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceShellScript intune will provide customer the ability to run their Shell scripts on the enrolled Mac OS devices. The script can be run once or periodically.
type DeviceShellScript struct {
    Entity
    // The list of group assignments for the device management script.
    assignments []DeviceManagementScriptAssignmentable
    // Does not notify the user a script is being executed
    blockExecutionNotifications *bool
    // The date and time the device management script was created. This property is read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Optional description for the device management script.
    description *string
    // List of run states for this script across all devices.
    deviceRunStates []DeviceManagementScriptDeviceStateable
    // Name of the device management script.
    displayName *string
    // The interval for script to run. If not defined the script will run once
    executionFrequency *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // Script file name.
    fileName *string
    // The list of group assignments for the device management script.
    groupAssignments []DeviceManagementScriptGroupAssignmentable
    // The date and time the device management script was last modified. This property is read-only.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Number of times for the script to be retried if it fails
    retryCount *int32
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
// NewDeviceShellScript instantiates a new deviceShellScript and sets the default values.
func NewDeviceShellScript()(*DeviceShellScript) {
    m := &DeviceShellScript{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.deviceShellScript";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceShellScriptFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceShellScriptFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceShellScript(), nil
}
// GetAssignments gets the assignments property value. The list of group assignments for the device management script.
func (m *DeviceShellScript) GetAssignments()([]DeviceManagementScriptAssignmentable) {
    return m.assignments
}
// GetBlockExecutionNotifications gets the blockExecutionNotifications property value. Does not notify the user a script is being executed
func (m *DeviceShellScript) GetBlockExecutionNotifications()(*bool) {
    return m.blockExecutionNotifications
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time the device management script was created. This property is read-only.
func (m *DeviceShellScript) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDescription gets the description property value. Optional description for the device management script.
func (m *DeviceShellScript) GetDescription()(*string) {
    return m.description
}
// GetDeviceRunStates gets the deviceRunStates property value. List of run states for this script across all devices.
func (m *DeviceShellScript) GetDeviceRunStates()([]DeviceManagementScriptDeviceStateable) {
    return m.deviceRunStates
}
// GetDisplayName gets the displayName property value. Name of the device management script.
func (m *DeviceShellScript) GetDisplayName()(*string) {
    return m.displayName
}
// GetExecutionFrequency gets the executionFrequency property value. The interval for script to run. If not defined the script will run once
func (m *DeviceShellScript) GetExecutionFrequency()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.executionFrequency
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceShellScript) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceManagementScriptAssignmentFromDiscriminatorValue , m.SetAssignments)
    res["blockExecutionNotifications"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetBlockExecutionNotifications)
    res["createdDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetCreatedDateTime)
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["deviceRunStates"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceManagementScriptDeviceStateFromDiscriminatorValue , m.SetDeviceRunStates)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["executionFrequency"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetISODurationValue(m.SetExecutionFrequency)
    res["fileName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetFileName)
    res["groupAssignments"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceManagementScriptGroupAssignmentFromDiscriminatorValue , m.SetGroupAssignments)
    res["lastModifiedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastModifiedDateTime)
    res["retryCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetRetryCount)
    res["roleScopeTagIds"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetRoleScopeTagIds)
    res["runAsAccount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseRunAsAccountType , m.SetRunAsAccount)
    res["runSummary"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateDeviceManagementScriptRunSummaryFromDiscriminatorValue , m.SetRunSummary)
    res["scriptContent"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetByteArrayValue(m.SetScriptContent)
    res["userRunStates"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceManagementScriptUserStateFromDiscriminatorValue , m.SetUserRunStates)
    return res
}
// GetFileName gets the fileName property value. Script file name.
func (m *DeviceShellScript) GetFileName()(*string) {
    return m.fileName
}
// GetGroupAssignments gets the groupAssignments property value. The list of group assignments for the device management script.
func (m *DeviceShellScript) GetGroupAssignments()([]DeviceManagementScriptGroupAssignmentable) {
    return m.groupAssignments
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time the device management script was last modified. This property is read-only.
func (m *DeviceShellScript) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetRetryCount gets the retryCount property value. Number of times for the script to be retried if it fails
func (m *DeviceShellScript) GetRetryCount()(*int32) {
    return m.retryCount
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. List of Scope Tag IDs for this PowerShellScript instance.
func (m *DeviceShellScript) GetRoleScopeTagIds()([]string) {
    return m.roleScopeTagIds
}
// GetRunAsAccount gets the runAsAccount property value. Indicates the type of execution context the app runs in.
func (m *DeviceShellScript) GetRunAsAccount()(*RunAsAccountType) {
    return m.runAsAccount
}
// GetRunSummary gets the runSummary property value. Run summary for device management script.
func (m *DeviceShellScript) GetRunSummary()(DeviceManagementScriptRunSummaryable) {
    return m.runSummary
}
// GetScriptContent gets the scriptContent property value. The script content.
func (m *DeviceShellScript) GetScriptContent()([]byte) {
    return m.scriptContent
}
// GetUserRunStates gets the userRunStates property value. List of run states for this script across all users.
func (m *DeviceShellScript) GetUserRunStates()([]DeviceManagementScriptUserStateable) {
    return m.userRunStates
}
// Serialize serializes information the current object
func (m *DeviceShellScript) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteBoolValue("blockExecutionNotifications", m.GetBlockExecutionNotifications())
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
        err = writer.WriteISODurationValue("executionFrequency", m.GetExecutionFrequency())
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
    {
        err = writer.WriteInt32Value("retryCount", m.GetRetryCount())
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
func (m *DeviceShellScript) SetAssignments(value []DeviceManagementScriptAssignmentable)() {
    m.assignments = value
}
// SetBlockExecutionNotifications sets the blockExecutionNotifications property value. Does not notify the user a script is being executed
func (m *DeviceShellScript) SetBlockExecutionNotifications(value *bool)() {
    m.blockExecutionNotifications = value
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time the device management script was created. This property is read-only.
func (m *DeviceShellScript) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDescription sets the description property value. Optional description for the device management script.
func (m *DeviceShellScript) SetDescription(value *string)() {
    m.description = value
}
// SetDeviceRunStates sets the deviceRunStates property value. List of run states for this script across all devices.
func (m *DeviceShellScript) SetDeviceRunStates(value []DeviceManagementScriptDeviceStateable)() {
    m.deviceRunStates = value
}
// SetDisplayName sets the displayName property value. Name of the device management script.
func (m *DeviceShellScript) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetExecutionFrequency sets the executionFrequency property value. The interval for script to run. If not defined the script will run once
func (m *DeviceShellScript) SetExecutionFrequency(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.executionFrequency = value
}
// SetFileName sets the fileName property value. Script file name.
func (m *DeviceShellScript) SetFileName(value *string)() {
    m.fileName = value
}
// SetGroupAssignments sets the groupAssignments property value. The list of group assignments for the device management script.
func (m *DeviceShellScript) SetGroupAssignments(value []DeviceManagementScriptGroupAssignmentable)() {
    m.groupAssignments = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time the device management script was last modified. This property is read-only.
func (m *DeviceShellScript) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetRetryCount sets the retryCount property value. Number of times for the script to be retried if it fails
func (m *DeviceShellScript) SetRetryCount(value *int32)() {
    m.retryCount = value
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. List of Scope Tag IDs for this PowerShellScript instance.
func (m *DeviceShellScript) SetRoleScopeTagIds(value []string)() {
    m.roleScopeTagIds = value
}
// SetRunAsAccount sets the runAsAccount property value. Indicates the type of execution context the app runs in.
func (m *DeviceShellScript) SetRunAsAccount(value *RunAsAccountType)() {
    m.runAsAccount = value
}
// SetRunSummary sets the runSummary property value. Run summary for device management script.
func (m *DeviceShellScript) SetRunSummary(value DeviceManagementScriptRunSummaryable)() {
    m.runSummary = value
}
// SetScriptContent sets the scriptContent property value. The script content.
func (m *DeviceShellScript) SetScriptContent(value []byte)() {
    m.scriptContent = value
}
// SetUserRunStates sets the userRunStates property value. List of run states for this script across all users.
func (m *DeviceShellScript) SetUserRunStates(value []DeviceManagementScriptUserStateable)() {
    m.userRunStates = value
}
