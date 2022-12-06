package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceHealthScript intune will provide customer the ability to run their Powershell Health scripts (remediation + detection) on the enrolled windows 10 Azure Active Directory joined devices.
type DeviceHealthScript struct {
    Entity
    // The list of group assignments for the device health script
    assignments []DeviceHealthScriptAssignmentable
    // The timestamp of when the device health script was created. This property is read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Description of the device health script
    description *string
    // The entire content of the detection powershell script
    detectionScriptContent []byte
    // List of ComplexType DetectionScriptParameters objects.
    detectionScriptParameters []DeviceHealthScriptParameterable
    // List of run states for the device health script across all devices
    deviceRunStates []DeviceHealthScriptDeviceStateable
    // Name of the device health script
    displayName *string
    // Indicate whether the script signature needs be checked
    enforceSignatureCheck *bool
    // Highest available version for a Microsoft Proprietary script
    highestAvailableVersion *string
    // Determines if this is Microsoft Proprietary Script. Proprietary scripts are read-only
    isGlobalScript *bool
    // The timestamp of when the device health script was modified. This property is read-only.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Name of the device health script publisher
    publisher *string
    // The entire content of the remediation powershell script
    remediationScriptContent []byte
    // List of ComplexType RemediationScriptParameters objects.
    remediationScriptParameters []DeviceHealthScriptParameterable
    // List of Scope Tag IDs for the device health script
    roleScopeTagIds []string
    // Indicate whether PowerShell script(s) should run as 32-bit
    runAs32Bit *bool
    // Indicates the type of execution context the app runs in.
    runAsAccount *RunAsAccountType
    // High level run summary for device health script.
    runSummary DeviceHealthScriptRunSummaryable
    // Version of the device health script
    version *string
}
// NewDeviceHealthScript instantiates a new deviceHealthScript and sets the default values.
func NewDeviceHealthScript()(*DeviceHealthScript) {
    m := &DeviceHealthScript{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceHealthScriptFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceHealthScriptFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceHealthScript(), nil
}
// GetAssignments gets the assignments property value. The list of group assignments for the device health script
func (m *DeviceHealthScript) GetAssignments()([]DeviceHealthScriptAssignmentable) {
    return m.assignments
}
// GetCreatedDateTime gets the createdDateTime property value. The timestamp of when the device health script was created. This property is read-only.
func (m *DeviceHealthScript) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDescription gets the description property value. Description of the device health script
func (m *DeviceHealthScript) GetDescription()(*string) {
    return m.description
}
// GetDetectionScriptContent gets the detectionScriptContent property value. The entire content of the detection powershell script
func (m *DeviceHealthScript) GetDetectionScriptContent()([]byte) {
    return m.detectionScriptContent
}
// GetDetectionScriptParameters gets the detectionScriptParameters property value. List of ComplexType DetectionScriptParameters objects.
func (m *DeviceHealthScript) GetDetectionScriptParameters()([]DeviceHealthScriptParameterable) {
    return m.detectionScriptParameters
}
// GetDeviceRunStates gets the deviceRunStates property value. List of run states for the device health script across all devices
func (m *DeviceHealthScript) GetDeviceRunStates()([]DeviceHealthScriptDeviceStateable) {
    return m.deviceRunStates
}
// GetDisplayName gets the displayName property value. Name of the device health script
func (m *DeviceHealthScript) GetDisplayName()(*string) {
    return m.displayName
}
// GetEnforceSignatureCheck gets the enforceSignatureCheck property value. Indicate whether the script signature needs be checked
func (m *DeviceHealthScript) GetEnforceSignatureCheck()(*bool) {
    return m.enforceSignatureCheck
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceHealthScript) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceHealthScriptAssignmentFromDiscriminatorValue , m.SetAssignments)
    res["createdDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetCreatedDateTime)
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["detectionScriptContent"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetByteArrayValue(m.SetDetectionScriptContent)
    res["detectionScriptParameters"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceHealthScriptParameterFromDiscriminatorValue , m.SetDetectionScriptParameters)
    res["deviceRunStates"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceHealthScriptDeviceStateFromDiscriminatorValue , m.SetDeviceRunStates)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["enforceSignatureCheck"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetEnforceSignatureCheck)
    res["highestAvailableVersion"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetHighestAvailableVersion)
    res["isGlobalScript"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsGlobalScript)
    res["lastModifiedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastModifiedDateTime)
    res["publisher"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPublisher)
    res["remediationScriptContent"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetByteArrayValue(m.SetRemediationScriptContent)
    res["remediationScriptParameters"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceHealthScriptParameterFromDiscriminatorValue , m.SetRemediationScriptParameters)
    res["roleScopeTagIds"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetRoleScopeTagIds)
    res["runAs32Bit"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetRunAs32Bit)
    res["runAsAccount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseRunAsAccountType , m.SetRunAsAccount)
    res["runSummary"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateDeviceHealthScriptRunSummaryFromDiscriminatorValue , m.SetRunSummary)
    res["version"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetVersion)
    return res
}
// GetHighestAvailableVersion gets the highestAvailableVersion property value. Highest available version for a Microsoft Proprietary script
func (m *DeviceHealthScript) GetHighestAvailableVersion()(*string) {
    return m.highestAvailableVersion
}
// GetIsGlobalScript gets the isGlobalScript property value. Determines if this is Microsoft Proprietary Script. Proprietary scripts are read-only
func (m *DeviceHealthScript) GetIsGlobalScript()(*bool) {
    return m.isGlobalScript
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The timestamp of when the device health script was modified. This property is read-only.
func (m *DeviceHealthScript) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetPublisher gets the publisher property value. Name of the device health script publisher
func (m *DeviceHealthScript) GetPublisher()(*string) {
    return m.publisher
}
// GetRemediationScriptContent gets the remediationScriptContent property value. The entire content of the remediation powershell script
func (m *DeviceHealthScript) GetRemediationScriptContent()([]byte) {
    return m.remediationScriptContent
}
// GetRemediationScriptParameters gets the remediationScriptParameters property value. List of ComplexType RemediationScriptParameters objects.
func (m *DeviceHealthScript) GetRemediationScriptParameters()([]DeviceHealthScriptParameterable) {
    return m.remediationScriptParameters
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. List of Scope Tag IDs for the device health script
func (m *DeviceHealthScript) GetRoleScopeTagIds()([]string) {
    return m.roleScopeTagIds
}
// GetRunAs32Bit gets the runAs32Bit property value. Indicate whether PowerShell script(s) should run as 32-bit
func (m *DeviceHealthScript) GetRunAs32Bit()(*bool) {
    return m.runAs32Bit
}
// GetRunAsAccount gets the runAsAccount property value. Indicates the type of execution context the app runs in.
func (m *DeviceHealthScript) GetRunAsAccount()(*RunAsAccountType) {
    return m.runAsAccount
}
// GetRunSummary gets the runSummary property value. High level run summary for device health script.
func (m *DeviceHealthScript) GetRunSummary()(DeviceHealthScriptRunSummaryable) {
    return m.runSummary
}
// GetVersion gets the version property value. Version of the device health script
func (m *DeviceHealthScript) GetVersion()(*string) {
    return m.version
}
// Serialize serializes information the current object
func (m *DeviceHealthScript) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("detectionScriptContent", m.GetDetectionScriptContent())
        if err != nil {
            return err
        }
    }
    if m.GetDetectionScriptParameters() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDetectionScriptParameters())
        err = writer.WriteCollectionOfObjectValues("detectionScriptParameters", cast)
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
        err = writer.WriteBoolValue("enforceSignatureCheck", m.GetEnforceSignatureCheck())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("highestAvailableVersion", m.GetHighestAvailableVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isGlobalScript", m.GetIsGlobalScript())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("publisher", m.GetPublisher())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("remediationScriptContent", m.GetRemediationScriptContent())
        if err != nil {
            return err
        }
    }
    if m.GetRemediationScriptParameters() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetRemediationScriptParameters())
        err = writer.WriteCollectionOfObjectValues("remediationScriptParameters", cast)
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
    {
        err = writer.WriteBoolValue("runAs32Bit", m.GetRunAs32Bit())
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
        err = writer.WriteStringValue("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignments sets the assignments property value. The list of group assignments for the device health script
func (m *DeviceHealthScript) SetAssignments(value []DeviceHealthScriptAssignmentable)() {
    m.assignments = value
}
// SetCreatedDateTime sets the createdDateTime property value. The timestamp of when the device health script was created. This property is read-only.
func (m *DeviceHealthScript) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDescription sets the description property value. Description of the device health script
func (m *DeviceHealthScript) SetDescription(value *string)() {
    m.description = value
}
// SetDetectionScriptContent sets the detectionScriptContent property value. The entire content of the detection powershell script
func (m *DeviceHealthScript) SetDetectionScriptContent(value []byte)() {
    m.detectionScriptContent = value
}
// SetDetectionScriptParameters sets the detectionScriptParameters property value. List of ComplexType DetectionScriptParameters objects.
func (m *DeviceHealthScript) SetDetectionScriptParameters(value []DeviceHealthScriptParameterable)() {
    m.detectionScriptParameters = value
}
// SetDeviceRunStates sets the deviceRunStates property value. List of run states for the device health script across all devices
func (m *DeviceHealthScript) SetDeviceRunStates(value []DeviceHealthScriptDeviceStateable)() {
    m.deviceRunStates = value
}
// SetDisplayName sets the displayName property value. Name of the device health script
func (m *DeviceHealthScript) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetEnforceSignatureCheck sets the enforceSignatureCheck property value. Indicate whether the script signature needs be checked
func (m *DeviceHealthScript) SetEnforceSignatureCheck(value *bool)() {
    m.enforceSignatureCheck = value
}
// SetHighestAvailableVersion sets the highestAvailableVersion property value. Highest available version for a Microsoft Proprietary script
func (m *DeviceHealthScript) SetHighestAvailableVersion(value *string)() {
    m.highestAvailableVersion = value
}
// SetIsGlobalScript sets the isGlobalScript property value. Determines if this is Microsoft Proprietary Script. Proprietary scripts are read-only
func (m *DeviceHealthScript) SetIsGlobalScript(value *bool)() {
    m.isGlobalScript = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The timestamp of when the device health script was modified. This property is read-only.
func (m *DeviceHealthScript) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetPublisher sets the publisher property value. Name of the device health script publisher
func (m *DeviceHealthScript) SetPublisher(value *string)() {
    m.publisher = value
}
// SetRemediationScriptContent sets the remediationScriptContent property value. The entire content of the remediation powershell script
func (m *DeviceHealthScript) SetRemediationScriptContent(value []byte)() {
    m.remediationScriptContent = value
}
// SetRemediationScriptParameters sets the remediationScriptParameters property value. List of ComplexType RemediationScriptParameters objects.
func (m *DeviceHealthScript) SetRemediationScriptParameters(value []DeviceHealthScriptParameterable)() {
    m.remediationScriptParameters = value
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. List of Scope Tag IDs for the device health script
func (m *DeviceHealthScript) SetRoleScopeTagIds(value []string)() {
    m.roleScopeTagIds = value
}
// SetRunAs32Bit sets the runAs32Bit property value. Indicate whether PowerShell script(s) should run as 32-bit
func (m *DeviceHealthScript) SetRunAs32Bit(value *bool)() {
    m.runAs32Bit = value
}
// SetRunAsAccount sets the runAsAccount property value. Indicates the type of execution context the app runs in.
func (m *DeviceHealthScript) SetRunAsAccount(value *RunAsAccountType)() {
    m.runAsAccount = value
}
// SetRunSummary sets the runSummary property value. High level run summary for device health script.
func (m *DeviceHealthScript) SetRunSummary(value DeviceHealthScriptRunSummaryable)() {
    m.runSummary = value
}
// SetVersion sets the version property value. Version of the device health script
func (m *DeviceHealthScript) SetVersion(value *string)() {
    m.version = value
}
