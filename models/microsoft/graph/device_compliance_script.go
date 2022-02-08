package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceComplianceScript 
type DeviceComplianceScript struct {
    Entity
    // The list of group assignments for the device compliance script
    assignments []DeviceHealthScriptAssignment;
    // The timestamp of when the device compliance script was created. This property is read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Description of the device compliance script
    description *string;
    // The entire content of the detection powershell script
    detectionScriptContent []byte;
    // List of run states for the device compliance script across all devices
    deviceRunStates []DeviceComplianceScriptDeviceState;
    // Name of the device compliance script
    displayName *string;
    // Indicate whether the script signature needs be checked
    enforceSignatureCheck *bool;
    // The timestamp of when the device compliance script was modified. This property is read-only.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Name of the device compliance script publisher
    publisher *string;
    // List of Scope Tag IDs for the device compliance script
    roleScopeTagIds []string;
    // Indicate whether PowerShell script(s) should run as 32-bit
    runAs32Bit *bool;
    // Indicates the type of execution context. Possible values are: system, user.
    runAsAccount *RunAsAccountType;
    // High level run summary for device compliance script.
    runSummary *DeviceComplianceScriptRunSummary;
    // Version of the device compliance script
    version *string;
}
// NewDeviceComplianceScript instantiates a new deviceComplianceScript and sets the default values.
func NewDeviceComplianceScript()(*DeviceComplianceScript) {
    m := &DeviceComplianceScript{
        Entity: *NewEntity(),
    }
    return m
}
// GetAssignments gets the assignments property value. The list of group assignments for the device compliance script
func (m *DeviceComplianceScript) GetAssignments()([]DeviceHealthScriptAssignment) {
    if m == nil {
        return nil
    } else {
        return m.assignments
    }
}
// GetCreatedDateTime gets the createdDateTime property value. The timestamp of when the device compliance script was created. This property is read-only.
func (m *DeviceComplianceScript) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetDescription gets the description property value. Description of the device compliance script
func (m *DeviceComplianceScript) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDetectionScriptContent gets the detectionScriptContent property value. The entire content of the detection powershell script
func (m *DeviceComplianceScript) GetDetectionScriptContent()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.detectionScriptContent
    }
}
// GetDeviceRunStates gets the deviceRunStates property value. List of run states for the device compliance script across all devices
func (m *DeviceComplianceScript) GetDeviceRunStates()([]DeviceComplianceScriptDeviceState) {
    if m == nil {
        return nil
    } else {
        return m.deviceRunStates
    }
}
// GetDisplayName gets the displayName property value. Name of the device compliance script
func (m *DeviceComplianceScript) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetEnforceSignatureCheck gets the enforceSignatureCheck property value. Indicate whether the script signature needs be checked
func (m *DeviceComplianceScript) GetEnforceSignatureCheck()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enforceSignatureCheck
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The timestamp of when the device compliance script was modified. This property is read-only.
func (m *DeviceComplianceScript) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetPublisher gets the publisher property value. Name of the device compliance script publisher
func (m *DeviceComplianceScript) GetPublisher()(*string) {
    if m == nil {
        return nil
    } else {
        return m.publisher
    }
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. List of Scope Tag IDs for the device compliance script
func (m *DeviceComplianceScript) GetRoleScopeTagIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.roleScopeTagIds
    }
}
// GetRunAs32Bit gets the runAs32Bit property value. Indicate whether PowerShell script(s) should run as 32-bit
func (m *DeviceComplianceScript) GetRunAs32Bit()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.runAs32Bit
    }
}
// GetRunAsAccount gets the runAsAccount property value. Indicates the type of execution context. Possible values are: system, user.
func (m *DeviceComplianceScript) GetRunAsAccount()(*RunAsAccountType) {
    if m == nil {
        return nil
    } else {
        return m.runAsAccount
    }
}
// GetRunSummary gets the runSummary property value. High level run summary for device compliance script.
func (m *DeviceComplianceScript) GetRunSummary()(*DeviceComplianceScriptRunSummary) {
    if m == nil {
        return nil
    } else {
        return m.runSummary
    }
}
// GetVersion gets the version property value. Version of the device compliance script
func (m *DeviceComplianceScript) GetVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceComplianceScript) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceHealthScriptAssignment() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceHealthScriptAssignment, len(val))
            for i, v := range val {
                res[i] = *(v.(*DeviceHealthScriptAssignment))
            }
            m.SetAssignments(res)
        }
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["detectionScriptContent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetectionScriptContent(val)
        }
        return nil
    }
    res["deviceRunStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceComplianceScriptDeviceState() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceComplianceScriptDeviceState, len(val))
            for i, v := range val {
                res[i] = *(v.(*DeviceComplianceScriptDeviceState))
            }
            m.SetDeviceRunStates(res)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["enforceSignatureCheck"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnforceSignatureCheck(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["publisher"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublisher(val)
        }
        return nil
    }
    res["roleScopeTagIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetRoleScopeTagIds(res)
        }
        return nil
    }
    res["runAs32Bit"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRunAs32Bit(val)
        }
        return nil
    }
    res["runAsAccount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRunAsAccountType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRunAsAccount(val.(*RunAsAccountType))
        }
        return nil
    }
    res["runSummary"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceComplianceScriptRunSummary() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRunSummary(val.(*DeviceComplianceScriptRunSummary))
        }
        return nil
    }
    res["version"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
func (m *DeviceComplianceScript) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DeviceComplianceScript) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssignments() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAssignments()))
        for i, v := range m.GetAssignments() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
    if m.GetDeviceRunStates() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDeviceRunStates()))
        for i, v := range m.GetDeviceRunStates() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
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
// SetAssignments sets the assignments property value. The list of group assignments for the device compliance script
func (m *DeviceComplianceScript) SetAssignments(value []DeviceHealthScriptAssignment)() {
    if m != nil {
        m.assignments = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The timestamp of when the device compliance script was created. This property is read-only.
func (m *DeviceComplianceScript) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetDescription sets the description property value. Description of the device compliance script
func (m *DeviceComplianceScript) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDetectionScriptContent sets the detectionScriptContent property value. The entire content of the detection powershell script
func (m *DeviceComplianceScript) SetDetectionScriptContent(value []byte)() {
    if m != nil {
        m.detectionScriptContent = value
    }
}
// SetDeviceRunStates sets the deviceRunStates property value. List of run states for the device compliance script across all devices
func (m *DeviceComplianceScript) SetDeviceRunStates(value []DeviceComplianceScriptDeviceState)() {
    if m != nil {
        m.deviceRunStates = value
    }
}
// SetDisplayName sets the displayName property value. Name of the device compliance script
func (m *DeviceComplianceScript) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetEnforceSignatureCheck sets the enforceSignatureCheck property value. Indicate whether the script signature needs be checked
func (m *DeviceComplianceScript) SetEnforceSignatureCheck(value *bool)() {
    if m != nil {
        m.enforceSignatureCheck = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The timestamp of when the device compliance script was modified. This property is read-only.
func (m *DeviceComplianceScript) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetPublisher sets the publisher property value. Name of the device compliance script publisher
func (m *DeviceComplianceScript) SetPublisher(value *string)() {
    if m != nil {
        m.publisher = value
    }
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. List of Scope Tag IDs for the device compliance script
func (m *DeviceComplianceScript) SetRoleScopeTagIds(value []string)() {
    if m != nil {
        m.roleScopeTagIds = value
    }
}
// SetRunAs32Bit sets the runAs32Bit property value. Indicate whether PowerShell script(s) should run as 32-bit
func (m *DeviceComplianceScript) SetRunAs32Bit(value *bool)() {
    if m != nil {
        m.runAs32Bit = value
    }
}
// SetRunAsAccount sets the runAsAccount property value. Indicates the type of execution context. Possible values are: system, user.
func (m *DeviceComplianceScript) SetRunAsAccount(value *RunAsAccountType)() {
    if m != nil {
        m.runAsAccount = value
    }
}
// SetRunSummary sets the runSummary property value. High level run summary for device compliance script.
func (m *DeviceComplianceScript) SetRunSummary(value *DeviceComplianceScriptRunSummary)() {
    if m != nil {
        m.runSummary = value
    }
}
// SetVersion sets the version property value. Version of the device compliance script
func (m *DeviceComplianceScript) SetVersion(value *string)() {
    if m != nil {
        m.version = value
    }
}
