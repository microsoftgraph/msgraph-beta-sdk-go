package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceHealthScript intune will provide customer the ability to run their Powershell Health scripts (remediation + detection) on the enrolled windows 10 Azure Active Directory joined devices.
type DeviceHealthScript struct {
    Entity
    // The OdataType property
    OdataType *string
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
    val, err := m.GetBackingStore().Get("assignments")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceHealthScriptAssignmentable)
    }
    return nil
}
// GetCreatedDateTime gets the createdDateTime property value. The timestamp of when the device health script was created. This property is read-only.
func (m *DeviceHealthScript) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetDescription gets the description property value. Description of the device health script
func (m *DeviceHealthScript) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDetectionScriptContent gets the detectionScriptContent property value. The entire content of the detection powershell script
func (m *DeviceHealthScript) GetDetectionScriptContent()([]byte) {
    val, err := m.GetBackingStore().Get("detectionScriptContent")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]byte)
    }
    return nil
}
// GetDetectionScriptParameters gets the detectionScriptParameters property value. List of ComplexType DetectionScriptParameters objects.
func (m *DeviceHealthScript) GetDetectionScriptParameters()([]DeviceHealthScriptParameterable) {
    val, err := m.GetBackingStore().Get("detectionScriptParameters")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceHealthScriptParameterable)
    }
    return nil
}
// GetDeviceHealthScriptType gets the deviceHealthScriptType property value. Indicates the type of device script.
func (m *DeviceHealthScript) GetDeviceHealthScriptType()(*DeviceHealthScriptType) {
    val, err := m.GetBackingStore().Get("deviceHealthScriptType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeviceHealthScriptType)
    }
    return nil
}
// GetDeviceRunStates gets the deviceRunStates property value. List of run states for the device health script across all devices
func (m *DeviceHealthScript) GetDeviceRunStates()([]DeviceHealthScriptDeviceStateable) {
    val, err := m.GetBackingStore().Get("deviceRunStates")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceHealthScriptDeviceStateable)
    }
    return nil
}
// GetDisplayName gets the displayName property value. Name of the device health script
func (m *DeviceHealthScript) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetEnforceSignatureCheck gets the enforceSignatureCheck property value. Indicate whether the script signature needs be checked
func (m *DeviceHealthScript) GetEnforceSignatureCheck()(*bool) {
    val, err := m.GetBackingStore().Get("enforceSignatureCheck")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceHealthScript) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceHealthScriptAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceHealthScriptAssignmentable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceHealthScriptAssignmentable)
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
    res["detectionScriptContent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetectionScriptContent(val)
        }
        return nil
    }
    res["detectionScriptParameters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceHealthScriptParameterFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceHealthScriptParameterable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceHealthScriptParameterable)
                }
            }
            m.SetDetectionScriptParameters(res)
        }
        return nil
    }
    res["deviceHealthScriptType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceHealthScriptType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceHealthScriptType(val.(*DeviceHealthScriptType))
        }
        return nil
    }
    res["deviceRunStates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceHealthScriptDeviceStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceHealthScriptDeviceStateable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceHealthScriptDeviceStateable)
                }
            }
            m.SetDeviceRunStates(res)
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
    res["enforceSignatureCheck"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnforceSignatureCheck(val)
        }
        return nil
    }
    res["highestAvailableVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHighestAvailableVersion(val)
        }
        return nil
    }
    res["isGlobalScript"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsGlobalScript(val)
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
    res["publisher"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublisher(val)
        }
        return nil
    }
    res["remediationScriptContent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemediationScriptContent(val)
        }
        return nil
    }
    res["remediationScriptParameters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceHealthScriptParameterFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceHealthScriptParameterable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceHealthScriptParameterable)
                }
            }
            m.SetRemediationScriptParameters(res)
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
    res["runAs32Bit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRunAs32Bit(val)
        }
        return nil
    }
    res["runAsAccount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRunAsAccountType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRunAsAccount(val.(*RunAsAccountType))
        }
        return nil
    }
    res["runSummary"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceHealthScriptRunSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRunSummary(val.(DeviceHealthScriptRunSummaryable))
        }
        return nil
    }
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
// GetHighestAvailableVersion gets the highestAvailableVersion property value. Highest available version for a Microsoft Proprietary script
func (m *DeviceHealthScript) GetHighestAvailableVersion()(*string) {
    val, err := m.GetBackingStore().Get("highestAvailableVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetIsGlobalScript gets the isGlobalScript property value. Determines if this is Microsoft Proprietary Script. Proprietary scripts are read-only
func (m *DeviceHealthScript) GetIsGlobalScript()(*bool) {
    val, err := m.GetBackingStore().Get("isGlobalScript")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The timestamp of when the device health script was modified. This property is read-only.
func (m *DeviceHealthScript) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastModifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetPublisher gets the publisher property value. Name of the device health script publisher
func (m *DeviceHealthScript) GetPublisher()(*string) {
    val, err := m.GetBackingStore().Get("publisher")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRemediationScriptContent gets the remediationScriptContent property value. The entire content of the remediation powershell script
func (m *DeviceHealthScript) GetRemediationScriptContent()([]byte) {
    val, err := m.GetBackingStore().Get("remediationScriptContent")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]byte)
    }
    return nil
}
// GetRemediationScriptParameters gets the remediationScriptParameters property value. List of ComplexType RemediationScriptParameters objects.
func (m *DeviceHealthScript) GetRemediationScriptParameters()([]DeviceHealthScriptParameterable) {
    val, err := m.GetBackingStore().Get("remediationScriptParameters")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceHealthScriptParameterable)
    }
    return nil
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. List of Scope Tag IDs for the device health script
func (m *DeviceHealthScript) GetRoleScopeTagIds()([]string) {
    val, err := m.GetBackingStore().Get("roleScopeTagIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetRunAs32Bit gets the runAs32Bit property value. Indicate whether PowerShell script(s) should run as 32-bit
func (m *DeviceHealthScript) GetRunAs32Bit()(*bool) {
    val, err := m.GetBackingStore().Get("runAs32Bit")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetRunAsAccount gets the runAsAccount property value. Indicates the type of execution context the app runs in.
func (m *DeviceHealthScript) GetRunAsAccount()(*RunAsAccountType) {
    val, err := m.GetBackingStore().Get("runAsAccount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*RunAsAccountType)
    }
    return nil
}
// GetRunSummary gets the runSummary property value. High level run summary for device health script.
func (m *DeviceHealthScript) GetRunSummary()(DeviceHealthScriptRunSummaryable) {
    val, err := m.GetBackingStore().Get("runSummary")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceHealthScriptRunSummaryable)
    }
    return nil
}
// GetVersion gets the version property value. Version of the device health script
func (m *DeviceHealthScript) GetVersion()(*string) {
    val, err := m.GetBackingStore().Get("version")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeviceHealthScript) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDetectionScriptParameters()))
        for i, v := range m.GetDetectionScriptParameters() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("detectionScriptParameters", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceHealthScriptType() != nil {
        cast := (*m.GetDeviceHealthScriptType()).String()
        err = writer.WriteStringValue("deviceHealthScriptType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceRunStates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceRunStates()))
        for i, v := range m.GetDeviceRunStates() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRemediationScriptParameters()))
        for i, v := range m.GetRemediationScriptParameters() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
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
    err := m.GetBackingStore().Set("assignments", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The timestamp of when the device health script was created. This property is read-only.
func (m *DeviceHealthScript) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. Description of the device health script
func (m *DeviceHealthScript) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDetectionScriptContent sets the detectionScriptContent property value. The entire content of the detection powershell script
func (m *DeviceHealthScript) SetDetectionScriptContent(value []byte)() {
    err := m.GetBackingStore().Set("detectionScriptContent", value)
    if err != nil {
        panic(err)
    }
}
// SetDetectionScriptParameters sets the detectionScriptParameters property value. List of ComplexType DetectionScriptParameters objects.
func (m *DeviceHealthScript) SetDetectionScriptParameters(value []DeviceHealthScriptParameterable)() {
    err := m.GetBackingStore().Set("detectionScriptParameters", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceHealthScriptType sets the deviceHealthScriptType property value. Indicates the type of device script.
func (m *DeviceHealthScript) SetDeviceHealthScriptType(value *DeviceHealthScriptType)() {
    err := m.GetBackingStore().Set("deviceHealthScriptType", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceRunStates sets the deviceRunStates property value. List of run states for the device health script across all devices
func (m *DeviceHealthScript) SetDeviceRunStates(value []DeviceHealthScriptDeviceStateable)() {
    err := m.GetBackingStore().Set("deviceRunStates", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. Name of the device health script
func (m *DeviceHealthScript) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetEnforceSignatureCheck sets the enforceSignatureCheck property value. Indicate whether the script signature needs be checked
func (m *DeviceHealthScript) SetEnforceSignatureCheck(value *bool)() {
    err := m.GetBackingStore().Set("enforceSignatureCheck", value)
    if err != nil {
        panic(err)
    }
}
// SetHighestAvailableVersion sets the highestAvailableVersion property value. Highest available version for a Microsoft Proprietary script
func (m *DeviceHealthScript) SetHighestAvailableVersion(value *string)() {
    err := m.GetBackingStore().Set("highestAvailableVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetIsGlobalScript sets the isGlobalScript property value. Determines if this is Microsoft Proprietary Script. Proprietary scripts are read-only
func (m *DeviceHealthScript) SetIsGlobalScript(value *bool)() {
    err := m.GetBackingStore().Set("isGlobalScript", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The timestamp of when the device health script was modified. This property is read-only.
func (m *DeviceHealthScript) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetPublisher sets the publisher property value. Name of the device health script publisher
func (m *DeviceHealthScript) SetPublisher(value *string)() {
    err := m.GetBackingStore().Set("publisher", value)
    if err != nil {
        panic(err)
    }
}
// SetRemediationScriptContent sets the remediationScriptContent property value. The entire content of the remediation powershell script
func (m *DeviceHealthScript) SetRemediationScriptContent(value []byte)() {
    err := m.GetBackingStore().Set("remediationScriptContent", value)
    if err != nil {
        panic(err)
    }
}
// SetRemediationScriptParameters sets the remediationScriptParameters property value. List of ComplexType RemediationScriptParameters objects.
func (m *DeviceHealthScript) SetRemediationScriptParameters(value []DeviceHealthScriptParameterable)() {
    err := m.GetBackingStore().Set("remediationScriptParameters", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. List of Scope Tag IDs for the device health script
func (m *DeviceHealthScript) SetRoleScopeTagIds(value []string)() {
    err := m.GetBackingStore().Set("roleScopeTagIds", value)
    if err != nil {
        panic(err)
    }
}
// SetRunAs32Bit sets the runAs32Bit property value. Indicate whether PowerShell script(s) should run as 32-bit
func (m *DeviceHealthScript) SetRunAs32Bit(value *bool)() {
    err := m.GetBackingStore().Set("runAs32Bit", value)
    if err != nil {
        panic(err)
    }
}
// SetRunAsAccount sets the runAsAccount property value. Indicates the type of execution context the app runs in.
func (m *DeviceHealthScript) SetRunAsAccount(value *RunAsAccountType)() {
    err := m.GetBackingStore().Set("runAsAccount", value)
    if err != nil {
        panic(err)
    }
}
// SetRunSummary sets the runSummary property value. High level run summary for device health script.
func (m *DeviceHealthScript) SetRunSummary(value DeviceHealthScriptRunSummaryable)() {
    err := m.GetBackingStore().Set("runSummary", value)
    if err != nil {
        panic(err)
    }
}
// SetVersion sets the version property value. Version of the device health script
func (m *DeviceHealthScript) SetVersion(value *string)() {
    err := m.GetBackingStore().Set("version", value)
    if err != nil {
        panic(err)
    }
}
// DeviceHealthScriptable 
type DeviceHealthScriptable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignments()([]DeviceHealthScriptAssignmentable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetDetectionScriptContent()([]byte)
    GetDetectionScriptParameters()([]DeviceHealthScriptParameterable)
    GetDeviceHealthScriptType()(*DeviceHealthScriptType)
    GetDeviceRunStates()([]DeviceHealthScriptDeviceStateable)
    GetDisplayName()(*string)
    GetEnforceSignatureCheck()(*bool)
    GetHighestAvailableVersion()(*string)
    GetIsGlobalScript()(*bool)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPublisher()(*string)
    GetRemediationScriptContent()([]byte)
    GetRemediationScriptParameters()([]DeviceHealthScriptParameterable)
    GetRoleScopeTagIds()([]string)
    GetRunAs32Bit()(*bool)
    GetRunAsAccount()(*RunAsAccountType)
    GetRunSummary()(DeviceHealthScriptRunSummaryable)
    GetVersion()(*string)
    SetAssignments(value []DeviceHealthScriptAssignmentable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetDetectionScriptContent(value []byte)()
    SetDetectionScriptParameters(value []DeviceHealthScriptParameterable)()
    SetDeviceHealthScriptType(value *DeviceHealthScriptType)()
    SetDeviceRunStates(value []DeviceHealthScriptDeviceStateable)()
    SetDisplayName(value *string)()
    SetEnforceSignatureCheck(value *bool)()
    SetHighestAvailableVersion(value *string)()
    SetIsGlobalScript(value *bool)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPublisher(value *string)()
    SetRemediationScriptContent(value []byte)()
    SetRemediationScriptParameters(value []DeviceHealthScriptParameterable)()
    SetRoleScopeTagIds(value []string)()
    SetRunAs32Bit(value *bool)()
    SetRunAsAccount(value *RunAsAccountType)()
    SetRunSummary(value DeviceHealthScriptRunSummaryable)()
    SetVersion(value *string)()
}
