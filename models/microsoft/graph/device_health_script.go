package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new deviceHealthScript and sets the default values.
func NewDeviceHealthScript()(*DeviceHealthScript) {
    m := &DeviceHealthScript{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the assignments property value. The list of group assignments for the device health script
func (m *DeviceHealthScript) GetAssignments()([]DeviceHealthScriptAssignment) {
    if m == nil {
        return nil
    } else {
        return m.assignments
    }
}
// Gets the createdDateTime property value. The timestamp of when the device health script was created. This property is read-only.
func (m *DeviceHealthScript) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the description property value. Description of the device health script
func (m *DeviceHealthScript) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the detectionScriptContent property value. The entire content of the detection powershell script
func (m *DeviceHealthScript) GetDetectionScriptContent()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.detectionScriptContent
    }
}
// Gets the detectionScriptParameters property value. List of ComplexType DetectionScriptParameters objects.
func (m *DeviceHealthScript) GetDetectionScriptParameters()([]DeviceHealthScriptParameter) {
    if m == nil {
        return nil
    } else {
        return m.detectionScriptParameters
    }
}
// Gets the deviceRunStates property value. List of run states for the device health script across all devices
func (m *DeviceHealthScript) GetDeviceRunStates()([]DeviceHealthScriptDeviceState) {
    if m == nil {
        return nil
    } else {
        return m.deviceRunStates
    }
}
// Gets the displayName property value. Name of the device health script
func (m *DeviceHealthScript) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the enforceSignatureCheck property value. Indicate whether the script signature needs be checked
func (m *DeviceHealthScript) GetEnforceSignatureCheck()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enforceSignatureCheck
    }
}
// Gets the highestAvailableVersion property value. Highest available version for a Microsoft Proprietary script
func (m *DeviceHealthScript) GetHighestAvailableVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.highestAvailableVersion
    }
}
// Gets the isGlobalScript property value. Determines if this is Microsoft Proprietary Script. Proprietary scripts are read-only
func (m *DeviceHealthScript) GetIsGlobalScript()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isGlobalScript
    }
}
// Gets the lastModifiedDateTime property value. The timestamp of when the device health script was modified. This property is read-only.
func (m *DeviceHealthScript) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the publisher property value. Name of the device health script publisher
func (m *DeviceHealthScript) GetPublisher()(*string) {
    if m == nil {
        return nil
    } else {
        return m.publisher
    }
}
// Gets the remediationScriptContent property value. The entire content of the remediation powershell script
func (m *DeviceHealthScript) GetRemediationScriptContent()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.remediationScriptContent
    }
}
// Gets the remediationScriptParameters property value. List of ComplexType RemediationScriptParameters objects.
func (m *DeviceHealthScript) GetRemediationScriptParameters()([]DeviceHealthScriptParameter) {
    if m == nil {
        return nil
    } else {
        return m.remediationScriptParameters
    }
}
// Gets the roleScopeTagIds property value. List of Scope Tag IDs for the device health script
func (m *DeviceHealthScript) GetRoleScopeTagIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.roleScopeTagIds
    }
}
// Gets the runAs32Bit property value. Indicate whether PowerShell script(s) should run as 32-bit
func (m *DeviceHealthScript) GetRunAs32Bit()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.runAs32Bit
    }
}
// Gets the runAsAccount property value. Indicates the type of execution context. Possible values are: system, user.
func (m *DeviceHealthScript) GetRunAsAccount()(*RunAsAccountType) {
    if m == nil {
        return nil
    } else {
        return m.runAsAccount
    }
}
// Gets the runSummary property value. High level run summary for device health script.
func (m *DeviceHealthScript) GetRunSummary()(*DeviceHealthScriptRunSummary) {
    if m == nil {
        return nil
    } else {
        return m.runSummary
    }
}
// Gets the version property value. Version of the device health script
func (m *DeviceHealthScript) GetVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
// The deserialization information for the current model
func (m *DeviceHealthScript) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceHealthScriptAssignment() })
        if err != nil {
            return err
        }
        res := make([]DeviceHealthScriptAssignment, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceHealthScriptAssignment))
        }
        m.SetAssignments(res)
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCreatedDateTime(val)
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
        return nil
    }
    res["detectionScriptContent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        m.SetDetectionScriptContent(val)
        return nil
    }
    res["detectionScriptParameters"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceHealthScriptParameter() })
        if err != nil {
            return err
        }
        res := make([]DeviceHealthScriptParameter, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceHealthScriptParameter))
        }
        m.SetDetectionScriptParameters(res)
        return nil
    }
    res["deviceRunStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceHealthScriptDeviceState() })
        if err != nil {
            return err
        }
        res := make([]DeviceHealthScriptDeviceState, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceHealthScriptDeviceState))
        }
        m.SetDeviceRunStates(res)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["enforceSignatureCheck"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetEnforceSignatureCheck(val)
        return nil
    }
    res["highestAvailableVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetHighestAvailableVersion(val)
        return nil
    }
    res["isGlobalScript"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsGlobalScript(val)
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastModifiedDateTime(val)
        return nil
    }
    res["publisher"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPublisher(val)
        return nil
    }
    res["remediationScriptContent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        m.SetRemediationScriptContent(val)
        return nil
    }
    res["remediationScriptParameters"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceHealthScriptParameter() })
        if err != nil {
            return err
        }
        res := make([]DeviceHealthScriptParameter, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceHealthScriptParameter))
        }
        m.SetRemediationScriptParameters(res)
        return nil
    }
    res["roleScopeTagIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetRoleScopeTagIds(res)
        return nil
    }
    res["runAs32Bit"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetRunAs32Bit(val)
        return nil
    }
    res["runAsAccount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRunAsAccountType)
        if err != nil {
            return err
        }
        cast := val.(RunAsAccountType)
        m.SetRunAsAccount(&cast)
        return nil
    }
    res["runSummary"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceHealthScriptRunSummary() })
        if err != nil {
            return err
        }
        m.SetRunSummary(val.(*DeviceHealthScriptRunSummary))
        return nil
    }
    res["version"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetVersion(val)
        return nil
    }
    return res
}
func (m *DeviceHealthScript) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DeviceHealthScript) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
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
    {
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
    {
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
    {
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
    {
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
        cast := m.GetRunAsAccount().String()
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
// Sets the assignments property value. The list of group assignments for the device health script
// Parameters:
//  - value : Value to set for the assignments property.
func (m *DeviceHealthScript) SetAssignments(value []DeviceHealthScriptAssignment)() {
    m.assignments = value
}
// Sets the createdDateTime property value. The timestamp of when the device health script was created. This property is read-only.
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *DeviceHealthScript) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the description property value. Description of the device health script
// Parameters:
//  - value : Value to set for the description property.
func (m *DeviceHealthScript) SetDescription(value *string)() {
    m.description = value
}
// Sets the detectionScriptContent property value. The entire content of the detection powershell script
// Parameters:
//  - value : Value to set for the detectionScriptContent property.
func (m *DeviceHealthScript) SetDetectionScriptContent(value []byte)() {
    m.detectionScriptContent = value
}
// Sets the detectionScriptParameters property value. List of ComplexType DetectionScriptParameters objects.
// Parameters:
//  - value : Value to set for the detectionScriptParameters property.
func (m *DeviceHealthScript) SetDetectionScriptParameters(value []DeviceHealthScriptParameter)() {
    m.detectionScriptParameters = value
}
// Sets the deviceRunStates property value. List of run states for the device health script across all devices
// Parameters:
//  - value : Value to set for the deviceRunStates property.
func (m *DeviceHealthScript) SetDeviceRunStates(value []DeviceHealthScriptDeviceState)() {
    m.deviceRunStates = value
}
// Sets the displayName property value. Name of the device health script
// Parameters:
//  - value : Value to set for the displayName property.
func (m *DeviceHealthScript) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the enforceSignatureCheck property value. Indicate whether the script signature needs be checked
// Parameters:
//  - value : Value to set for the enforceSignatureCheck property.
func (m *DeviceHealthScript) SetEnforceSignatureCheck(value *bool)() {
    m.enforceSignatureCheck = value
}
// Sets the highestAvailableVersion property value. Highest available version for a Microsoft Proprietary script
// Parameters:
//  - value : Value to set for the highestAvailableVersion property.
func (m *DeviceHealthScript) SetHighestAvailableVersion(value *string)() {
    m.highestAvailableVersion = value
}
// Sets the isGlobalScript property value. Determines if this is Microsoft Proprietary Script. Proprietary scripts are read-only
// Parameters:
//  - value : Value to set for the isGlobalScript property.
func (m *DeviceHealthScript) SetIsGlobalScript(value *bool)() {
    m.isGlobalScript = value
}
// Sets the lastModifiedDateTime property value. The timestamp of when the device health script was modified. This property is read-only.
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *DeviceHealthScript) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the publisher property value. Name of the device health script publisher
// Parameters:
//  - value : Value to set for the publisher property.
func (m *DeviceHealthScript) SetPublisher(value *string)() {
    m.publisher = value
}
// Sets the remediationScriptContent property value. The entire content of the remediation powershell script
// Parameters:
//  - value : Value to set for the remediationScriptContent property.
func (m *DeviceHealthScript) SetRemediationScriptContent(value []byte)() {
    m.remediationScriptContent = value
}
// Sets the remediationScriptParameters property value. List of ComplexType RemediationScriptParameters objects.
// Parameters:
//  - value : Value to set for the remediationScriptParameters property.
func (m *DeviceHealthScript) SetRemediationScriptParameters(value []DeviceHealthScriptParameter)() {
    m.remediationScriptParameters = value
}
// Sets the roleScopeTagIds property value. List of Scope Tag IDs for the device health script
// Parameters:
//  - value : Value to set for the roleScopeTagIds property.
func (m *DeviceHealthScript) SetRoleScopeTagIds(value []string)() {
    m.roleScopeTagIds = value
}
// Sets the runAs32Bit property value. Indicate whether PowerShell script(s) should run as 32-bit
// Parameters:
//  - value : Value to set for the runAs32Bit property.
func (m *DeviceHealthScript) SetRunAs32Bit(value *bool)() {
    m.runAs32Bit = value
}
// Sets the runAsAccount property value. Indicates the type of execution context. Possible values are: system, user.
// Parameters:
//  - value : Value to set for the runAsAccount property.
func (m *DeviceHealthScript) SetRunAsAccount(value *RunAsAccountType)() {
    m.runAsAccount = value
}
// Sets the runSummary property value. High level run summary for device health script.
// Parameters:
//  - value : Value to set for the runSummary property.
func (m *DeviceHealthScript) SetRunSummary(value *DeviceHealthScriptRunSummary)() {
    m.runSummary = value
}
// Sets the version property value. Version of the device health script
// Parameters:
//  - value : Value to set for the version property.
func (m *DeviceHealthScript) SetVersion(value *string)() {
    m.version = value
}
