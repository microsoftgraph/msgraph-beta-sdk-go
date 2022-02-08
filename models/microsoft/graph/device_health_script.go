package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceHealthScript 
type DeviceHealthScript struct {
    Entity
    // The list of group assignments for the device health script
    assignments []DeviceHealthScriptAssignment;
    // The timestamp of when the device health script was created. This property is read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Description of the device health script
    description *string;
    // The entire content of the detection powershell script
    detectionScriptContent []byte;
    // List of ComplexType DetectionScriptParameters objects.
    detectionScriptParameters []DeviceHealthScriptParameter;
    // List of run states for the device health script across all devices
    deviceRunStates []DeviceHealthScriptDeviceState;
    // Name of the device health script
    displayName *string;
    // Indicate whether the script signature needs be checked
    enforceSignatureCheck *bool;
    // Highest available version for a Microsoft Proprietary script
    highestAvailableVersion *string;
    // Determines if this is Microsoft Proprietary Script. Proprietary scripts are read-only
    isGlobalScript *bool;
    // The timestamp of when the device health script was modified. This property is read-only.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Name of the device health script publisher
    publisher *string;
    // The entire content of the remediation powershell script
    remediationScriptContent []byte;
    // List of ComplexType RemediationScriptParameters objects.
    remediationScriptParameters []DeviceHealthScriptParameter;
    // List of Scope Tag IDs for the device health script
    roleScopeTagIds []string;
    // Indicate whether PowerShell script(s) should run as 32-bit
    runAs32Bit *bool;
    // Indicates the type of execution context. Possible values are: system, user.
    runAsAccount *RunAsAccountType;
    // High level run summary for device health script.
    runSummary *DeviceHealthScriptRunSummary;
    // Version of the device health script
    version *string;
}
// NewDeviceHealthScript instantiates a new deviceHealthScript and sets the default values.
func NewDeviceHealthScript()(*DeviceHealthScript) {
    m := &DeviceHealthScript{
        Entity: *NewEntity(),
    }
    return m
}
// GetAssignments gets the assignments property value. The list of group assignments for the device health script
func (m *DeviceHealthScript) GetAssignments()([]DeviceHealthScriptAssignment) {
    if m == nil {
        return nil
    } else {
        return m.assignments
    }
}
// GetCreatedDateTime gets the createdDateTime property value. The timestamp of when the device health script was created. This property is read-only.
func (m *DeviceHealthScript) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetDescription gets the description property value. Description of the device health script
func (m *DeviceHealthScript) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDetectionScriptContent gets the detectionScriptContent property value. The entire content of the detection powershell script
func (m *DeviceHealthScript) GetDetectionScriptContent()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.detectionScriptContent
    }
}
// GetDetectionScriptParameters gets the detectionScriptParameters property value. List of ComplexType DetectionScriptParameters objects.
func (m *DeviceHealthScript) GetDetectionScriptParameters()([]DeviceHealthScriptParameter) {
    if m == nil {
        return nil
    } else {
        return m.detectionScriptParameters
    }
}
// GetDeviceRunStates gets the deviceRunStates property value. List of run states for the device health script across all devices
func (m *DeviceHealthScript) GetDeviceRunStates()([]DeviceHealthScriptDeviceState) {
    if m == nil {
        return nil
    } else {
        return m.deviceRunStates
    }
}
// GetDisplayName gets the displayName property value. Name of the device health script
func (m *DeviceHealthScript) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetEnforceSignatureCheck gets the enforceSignatureCheck property value. Indicate whether the script signature needs be checked
func (m *DeviceHealthScript) GetEnforceSignatureCheck()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enforceSignatureCheck
    }
}
// GetHighestAvailableVersion gets the highestAvailableVersion property value. Highest available version for a Microsoft Proprietary script
func (m *DeviceHealthScript) GetHighestAvailableVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.highestAvailableVersion
    }
}
// GetIsGlobalScript gets the isGlobalScript property value. Determines if this is Microsoft Proprietary Script. Proprietary scripts are read-only
func (m *DeviceHealthScript) GetIsGlobalScript()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isGlobalScript
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The timestamp of when the device health script was modified. This property is read-only.
func (m *DeviceHealthScript) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetPublisher gets the publisher property value. Name of the device health script publisher
func (m *DeviceHealthScript) GetPublisher()(*string) {
    if m == nil {
        return nil
    } else {
        return m.publisher
    }
}
// GetRemediationScriptContent gets the remediationScriptContent property value. The entire content of the remediation powershell script
func (m *DeviceHealthScript) GetRemediationScriptContent()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.remediationScriptContent
    }
}
// GetRemediationScriptParameters gets the remediationScriptParameters property value. List of ComplexType RemediationScriptParameters objects.
func (m *DeviceHealthScript) GetRemediationScriptParameters()([]DeviceHealthScriptParameter) {
    if m == nil {
        return nil
    } else {
        return m.remediationScriptParameters
    }
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. List of Scope Tag IDs for the device health script
func (m *DeviceHealthScript) GetRoleScopeTagIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.roleScopeTagIds
    }
}
// GetRunAs32Bit gets the runAs32Bit property value. Indicate whether PowerShell script(s) should run as 32-bit
func (m *DeviceHealthScript) GetRunAs32Bit()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.runAs32Bit
    }
}
// GetRunAsAccount gets the runAsAccount property value. Indicates the type of execution context. Possible values are: system, user.
func (m *DeviceHealthScript) GetRunAsAccount()(*RunAsAccountType) {
    if m == nil {
        return nil
    } else {
        return m.runAsAccount
    }
}
// GetRunSummary gets the runSummary property value. High level run summary for device health script.
func (m *DeviceHealthScript) GetRunSummary()(*DeviceHealthScriptRunSummary) {
    if m == nil {
        return nil
    } else {
        return m.runSummary
    }
}
// GetVersion gets the version property value. Version of the device health script
func (m *DeviceHealthScript) GetVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceHealthScript) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    res["detectionScriptParameters"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceHealthScriptParameter() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceHealthScriptParameter, len(val))
            for i, v := range val {
                res[i] = *(v.(*DeviceHealthScriptParameter))
            }
            m.SetDetectionScriptParameters(res)
        }
        return nil
    }
    res["deviceRunStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceHealthScriptDeviceState() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceHealthScriptDeviceState, len(val))
            for i, v := range val {
                res[i] = *(v.(*DeviceHealthScriptDeviceState))
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
    res["highestAvailableVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHighestAvailableVersion(val)
        }
        return nil
    }
    res["isGlobalScript"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsGlobalScript(val)
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
    res["remediationScriptContent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemediationScriptContent(val)
        }
        return nil
    }
    res["remediationScriptParameters"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceHealthScriptParameter() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceHealthScriptParameter, len(val))
            for i, v := range val {
                res[i] = *(v.(*DeviceHealthScriptParameter))
            }
            m.SetRemediationScriptParameters(res)
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
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceHealthScriptRunSummary() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRunSummary(val.(*DeviceHealthScriptRunSummary))
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
func (m *DeviceHealthScript) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DeviceHealthScript) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
    if m.GetDetectionScriptParameters() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDetectionScriptParameters()))
        for i, v := range m.GetDetectionScriptParameters() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("detectionScriptParameters", cast)
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
    {
        err = writer.WriteByteArrayValue("remediationScriptContent", m.GetRemediationScriptContent())
        if err != nil {
            return err
        }
    }
    if m.GetRemediationScriptParameters() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRemediationScriptParameters()))
        for i, v := range m.GetRemediationScriptParameters() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
func (m *DeviceHealthScript) SetAssignments(value []DeviceHealthScriptAssignment)() {
    if m != nil {
        m.assignments = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The timestamp of when the device health script was created. This property is read-only.
func (m *DeviceHealthScript) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetDescription sets the description property value. Description of the device health script
func (m *DeviceHealthScript) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDetectionScriptContent sets the detectionScriptContent property value. The entire content of the detection powershell script
func (m *DeviceHealthScript) SetDetectionScriptContent(value []byte)() {
    if m != nil {
        m.detectionScriptContent = value
    }
}
// SetDetectionScriptParameters sets the detectionScriptParameters property value. List of ComplexType DetectionScriptParameters objects.
func (m *DeviceHealthScript) SetDetectionScriptParameters(value []DeviceHealthScriptParameter)() {
    if m != nil {
        m.detectionScriptParameters = value
    }
}
// SetDeviceRunStates sets the deviceRunStates property value. List of run states for the device health script across all devices
func (m *DeviceHealthScript) SetDeviceRunStates(value []DeviceHealthScriptDeviceState)() {
    if m != nil {
        m.deviceRunStates = value
    }
}
// SetDisplayName sets the displayName property value. Name of the device health script
func (m *DeviceHealthScript) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetEnforceSignatureCheck sets the enforceSignatureCheck property value. Indicate whether the script signature needs be checked
func (m *DeviceHealthScript) SetEnforceSignatureCheck(value *bool)() {
    if m != nil {
        m.enforceSignatureCheck = value
    }
}
// SetHighestAvailableVersion sets the highestAvailableVersion property value. Highest available version for a Microsoft Proprietary script
func (m *DeviceHealthScript) SetHighestAvailableVersion(value *string)() {
    if m != nil {
        m.highestAvailableVersion = value
    }
}
// SetIsGlobalScript sets the isGlobalScript property value. Determines if this is Microsoft Proprietary Script. Proprietary scripts are read-only
func (m *DeviceHealthScript) SetIsGlobalScript(value *bool)() {
    if m != nil {
        m.isGlobalScript = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The timestamp of when the device health script was modified. This property is read-only.
func (m *DeviceHealthScript) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetPublisher sets the publisher property value. Name of the device health script publisher
func (m *DeviceHealthScript) SetPublisher(value *string)() {
    if m != nil {
        m.publisher = value
    }
}
// SetRemediationScriptContent sets the remediationScriptContent property value. The entire content of the remediation powershell script
func (m *DeviceHealthScript) SetRemediationScriptContent(value []byte)() {
    if m != nil {
        m.remediationScriptContent = value
    }
}
// SetRemediationScriptParameters sets the remediationScriptParameters property value. List of ComplexType RemediationScriptParameters objects.
func (m *DeviceHealthScript) SetRemediationScriptParameters(value []DeviceHealthScriptParameter)() {
    if m != nil {
        m.remediationScriptParameters = value
    }
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. List of Scope Tag IDs for the device health script
func (m *DeviceHealthScript) SetRoleScopeTagIds(value []string)() {
    if m != nil {
        m.roleScopeTagIds = value
    }
}
// SetRunAs32Bit sets the runAs32Bit property value. Indicate whether PowerShell script(s) should run as 32-bit
func (m *DeviceHealthScript) SetRunAs32Bit(value *bool)() {
    if m != nil {
        m.runAs32Bit = value
    }
}
// SetRunAsAccount sets the runAsAccount property value. Indicates the type of execution context. Possible values are: system, user.
func (m *DeviceHealthScript) SetRunAsAccount(value *RunAsAccountType)() {
    if m != nil {
        m.runAsAccount = value
    }
}
// SetRunSummary sets the runSummary property value. High level run summary for device health script.
func (m *DeviceHealthScript) SetRunSummary(value *DeviceHealthScriptRunSummary)() {
    if m != nil {
        m.runSummary = value
    }
}
// SetVersion sets the version property value. Version of the device health script
func (m *DeviceHealthScript) SetVersion(value *string)() {
    if m != nil {
        m.version = value
    }
}
