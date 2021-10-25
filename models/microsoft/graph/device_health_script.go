package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeviceHealthScript struct {
    Entity
    assignments []DeviceHealthScriptAssignment;
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    description *string;
    detectionScriptContent []byte;
    detectionScriptParameters []DeviceHealthScriptParameter;
    deviceRunStates []DeviceHealthScriptDeviceState;
    displayName *string;
    enforceSignatureCheck *bool;
    highestAvailableVersion *string;
    isGlobalScript *bool;
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    publisher *string;
    remediationScriptContent []byte;
    remediationScriptParameters []DeviceHealthScriptParameter;
    roleScopeTagIds []string;
    runAs32Bit *bool;
    runAsAccount *RunAsAccountType;
    runSummary *DeviceHealthScriptRunSummary;
    version *string;
}
func NewDeviceHealthScript()(*DeviceHealthScript) {
    m := &DeviceHealthScript{
        Entity: *NewEntity(),
    }
    return m
}
func (m *DeviceHealthScript) GetAssignments()([]DeviceHealthScriptAssignment) {
    if m == nil {
        return nil
    } else {
        return m.assignments
    }
}
func (m *DeviceHealthScript) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
func (m *DeviceHealthScript) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
func (m *DeviceHealthScript) GetDetectionScriptContent()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.detectionScriptContent
    }
}
func (m *DeviceHealthScript) GetDetectionScriptParameters()([]DeviceHealthScriptParameter) {
    if m == nil {
        return nil
    } else {
        return m.detectionScriptParameters
    }
}
func (m *DeviceHealthScript) GetDeviceRunStates()([]DeviceHealthScriptDeviceState) {
    if m == nil {
        return nil
    } else {
        return m.deviceRunStates
    }
}
func (m *DeviceHealthScript) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *DeviceHealthScript) GetEnforceSignatureCheck()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enforceSignatureCheck
    }
}
func (m *DeviceHealthScript) GetHighestAvailableVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.highestAvailableVersion
    }
}
func (m *DeviceHealthScript) GetIsGlobalScript()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isGlobalScript
    }
}
func (m *DeviceHealthScript) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
func (m *DeviceHealthScript) GetPublisher()(*string) {
    if m == nil {
        return nil
    } else {
        return m.publisher
    }
}
func (m *DeviceHealthScript) GetRemediationScriptContent()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.remediationScriptContent
    }
}
func (m *DeviceHealthScript) GetRemediationScriptParameters()([]DeviceHealthScriptParameter) {
    if m == nil {
        return nil
    } else {
        return m.remediationScriptParameters
    }
}
func (m *DeviceHealthScript) GetRoleScopeTagIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.roleScopeTagIds
    }
}
func (m *DeviceHealthScript) GetRunAs32Bit()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.runAs32Bit
    }
}
func (m *DeviceHealthScript) GetRunAsAccount()(*RunAsAccountType) {
    if m == nil {
        return nil
    } else {
        return m.runAsAccount
    }
}
func (m *DeviceHealthScript) GetRunSummary()(*DeviceHealthScriptRunSummary) {
    if m == nil {
        return nil
    } else {
        return m.runSummary
    }
}
func (m *DeviceHealthScript) GetVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
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
            res[i] = v.(string)
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
func (m *DeviceHealthScript) SetAssignments(value []DeviceHealthScriptAssignment)() {
    m.assignments = value
}
func (m *DeviceHealthScript) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
func (m *DeviceHealthScript) SetDescription(value *string)() {
    m.description = value
}
func (m *DeviceHealthScript) SetDetectionScriptContent(value []byte)() {
    m.detectionScriptContent = value
}
func (m *DeviceHealthScript) SetDetectionScriptParameters(value []DeviceHealthScriptParameter)() {
    m.detectionScriptParameters = value
}
func (m *DeviceHealthScript) SetDeviceRunStates(value []DeviceHealthScriptDeviceState)() {
    m.deviceRunStates = value
}
func (m *DeviceHealthScript) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *DeviceHealthScript) SetEnforceSignatureCheck(value *bool)() {
    m.enforceSignatureCheck = value
}
func (m *DeviceHealthScript) SetHighestAvailableVersion(value *string)() {
    m.highestAvailableVersion = value
}
func (m *DeviceHealthScript) SetIsGlobalScript(value *bool)() {
    m.isGlobalScript = value
}
func (m *DeviceHealthScript) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
func (m *DeviceHealthScript) SetPublisher(value *string)() {
    m.publisher = value
}
func (m *DeviceHealthScript) SetRemediationScriptContent(value []byte)() {
    m.remediationScriptContent = value
}
func (m *DeviceHealthScript) SetRemediationScriptParameters(value []DeviceHealthScriptParameter)() {
    m.remediationScriptParameters = value
}
func (m *DeviceHealthScript) SetRoleScopeTagIds(value []string)() {
    m.roleScopeTagIds = value
}
func (m *DeviceHealthScript) SetRunAs32Bit(value *bool)() {
    m.runAs32Bit = value
}
func (m *DeviceHealthScript) SetRunAsAccount(value *RunAsAccountType)() {
    m.runAsAccount = value
}
func (m *DeviceHealthScript) SetRunSummary(value *DeviceHealthScriptRunSummary)() {
    m.runSummary = value
}
func (m *DeviceHealthScript) SetVersion(value *string)() {
    m.version = value
}
