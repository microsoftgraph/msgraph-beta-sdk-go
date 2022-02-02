package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WindowsDefenderApplicationControlSupplementalPolicy 
type WindowsDefenderApplicationControlSupplementalPolicy struct {
    Entity
    // The associated group assignments for this WindowsDefenderApplicationControl supplemental policy.
    assignments []WindowsDefenderApplicationControlSupplementalPolicyAssignment;
    // The WindowsDefenderApplicationControl supplemental policy content in byte array format.
    content []byte;
    // The WindowsDefenderApplicationControl supplemental policy content's file name.
    contentFileName *string;
    // The date and time when the WindowsDefenderApplicationControl supplemental policy was uploaded.
    creationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // WindowsDefenderApplicationControl supplemental policy deployment summary.
    deploySummary *WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary;
    // The description of WindowsDefenderApplicationControl supplemental policy.
    description *string;
    // The list of device deployment states for this WindowsDefenderApplicationControl supplemental policy.
    deviceStatuses []WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus;
    // The display name of WindowsDefenderApplicationControl supplemental policy.
    displayName *string;
    // The date and time when the WindowsDefenderApplicationControl supplemental policy was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // List of Scope Tags for this WindowsDefenderApplicationControl supplemental policy entity.
    roleScopeTagIds []string;
    // The WindowsDefenderApplicationControl supplemental policy's version.
    version *string;
}
// NewWindowsDefenderApplicationControlSupplementalPolicy instantiates a new windowsDefenderApplicationControlSupplementalPolicy and sets the default values.
func NewWindowsDefenderApplicationControlSupplementalPolicy()(*WindowsDefenderApplicationControlSupplementalPolicy) {
    m := &WindowsDefenderApplicationControlSupplementalPolicy{
        Entity: *NewEntity(),
    }
    return m
}
// GetAssignments gets the assignments property value. The associated group assignments for this WindowsDefenderApplicationControl supplemental policy.
func (m *WindowsDefenderApplicationControlSupplementalPolicy) GetAssignments()([]WindowsDefenderApplicationControlSupplementalPolicyAssignment) {
    if m == nil {
        return nil
    } else {
        return m.assignments
    }
}
// GetContent gets the content property value. The WindowsDefenderApplicationControl supplemental policy content in byte array format.
func (m *WindowsDefenderApplicationControlSupplementalPolicy) GetContent()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.content
    }
}
// GetContentFileName gets the contentFileName property value. The WindowsDefenderApplicationControl supplemental policy content's file name.
func (m *WindowsDefenderApplicationControlSupplementalPolicy) GetContentFileName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contentFileName
    }
}
// GetCreationDateTime gets the creationDateTime property value. The date and time when the WindowsDefenderApplicationControl supplemental policy was uploaded.
func (m *WindowsDefenderApplicationControlSupplementalPolicy) GetCreationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.creationDateTime
    }
}
// GetDeploySummary gets the deploySummary property value. WindowsDefenderApplicationControl supplemental policy deployment summary.
func (m *WindowsDefenderApplicationControlSupplementalPolicy) GetDeploySummary()(*WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary) {
    if m == nil {
        return nil
    } else {
        return m.deploySummary
    }
}
// GetDescription gets the description property value. The description of WindowsDefenderApplicationControl supplemental policy.
func (m *WindowsDefenderApplicationControlSupplementalPolicy) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDeviceStatuses gets the deviceStatuses property value. The list of device deployment states for this WindowsDefenderApplicationControl supplemental policy.
func (m *WindowsDefenderApplicationControlSupplementalPolicy) GetDeviceStatuses()([]WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) {
    if m == nil {
        return nil
    } else {
        return m.deviceStatuses
    }
}
// GetDisplayName gets the displayName property value. The display name of WindowsDefenderApplicationControl supplemental policy.
func (m *WindowsDefenderApplicationControlSupplementalPolicy) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time when the WindowsDefenderApplicationControl supplemental policy was last modified.
func (m *WindowsDefenderApplicationControlSupplementalPolicy) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. List of Scope Tags for this WindowsDefenderApplicationControl supplemental policy entity.
func (m *WindowsDefenderApplicationControlSupplementalPolicy) GetRoleScopeTagIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.roleScopeTagIds
    }
}
// GetVersion gets the version property value. The WindowsDefenderApplicationControl supplemental policy's version.
func (m *WindowsDefenderApplicationControlSupplementalPolicy) GetVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsDefenderApplicationControlSupplementalPolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWindowsDefenderApplicationControlSupplementalPolicyAssignment() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsDefenderApplicationControlSupplementalPolicyAssignment, len(val))
            for i, v := range val {
                res[i] = *(v.(*WindowsDefenderApplicationControlSupplementalPolicyAssignment))
            }
            m.SetAssignments(res)
        }
        return nil
    }
    res["content"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContent(val)
        }
        return nil
    }
    res["contentFileName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentFileName(val)
        }
        return nil
    }
    res["creationDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreationDateTime(val)
        }
        return nil
    }
    res["deploySummary"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeploySummary(val.(*WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary))
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
    res["deviceStatuses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus, len(val))
            for i, v := range val {
                res[i] = *(v.(*WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus))
            }
            m.SetDeviceStatuses(res)
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
func (m *WindowsDefenderApplicationControlSupplementalPolicy) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *WindowsDefenderApplicationControlSupplementalPolicy) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err = writer.WriteByteArrayValue("content", m.GetContent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("contentFileName", m.GetContentFileName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("creationDateTime", m.GetCreationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deploySummary", m.GetDeploySummary())
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
    if m.GetDeviceStatuses() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDeviceStatuses()))
        for i, v := range m.GetDeviceStatuses() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("deviceStatuses", cast)
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
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
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
        err = writer.WriteStringValue("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignments sets the assignments property value. The associated group assignments for this WindowsDefenderApplicationControl supplemental policy.
func (m *WindowsDefenderApplicationControlSupplementalPolicy) SetAssignments(value []WindowsDefenderApplicationControlSupplementalPolicyAssignment)() {
    if m != nil {
        m.assignments = value
    }
}
// SetContent sets the content property value. The WindowsDefenderApplicationControl supplemental policy content in byte array format.
func (m *WindowsDefenderApplicationControlSupplementalPolicy) SetContent(value []byte)() {
    if m != nil {
        m.content = value
    }
}
// SetContentFileName sets the contentFileName property value. The WindowsDefenderApplicationControl supplemental policy content's file name.
func (m *WindowsDefenderApplicationControlSupplementalPolicy) SetContentFileName(value *string)() {
    if m != nil {
        m.contentFileName = value
    }
}
// SetCreationDateTime sets the creationDateTime property value. The date and time when the WindowsDefenderApplicationControl supplemental policy was uploaded.
func (m *WindowsDefenderApplicationControlSupplementalPolicy) SetCreationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.creationDateTime = value
    }
}
// SetDeploySummary sets the deploySummary property value. WindowsDefenderApplicationControl supplemental policy deployment summary.
func (m *WindowsDefenderApplicationControlSupplementalPolicy) SetDeploySummary(value *WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary)() {
    if m != nil {
        m.deploySummary = value
    }
}
// SetDescription sets the description property value. The description of WindowsDefenderApplicationControl supplemental policy.
func (m *WindowsDefenderApplicationControlSupplementalPolicy) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDeviceStatuses sets the deviceStatuses property value. The list of device deployment states for this WindowsDefenderApplicationControl supplemental policy.
func (m *WindowsDefenderApplicationControlSupplementalPolicy) SetDeviceStatuses(value []WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus)() {
    if m != nil {
        m.deviceStatuses = value
    }
}
// SetDisplayName sets the displayName property value. The display name of WindowsDefenderApplicationControl supplemental policy.
func (m *WindowsDefenderApplicationControlSupplementalPolicy) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time when the WindowsDefenderApplicationControl supplemental policy was last modified.
func (m *WindowsDefenderApplicationControlSupplementalPolicy) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. List of Scope Tags for this WindowsDefenderApplicationControl supplemental policy entity.
func (m *WindowsDefenderApplicationControlSupplementalPolicy) SetRoleScopeTagIds(value []string)() {
    if m != nil {
        m.roleScopeTagIds = value
    }
}
// SetVersion sets the version property value. The WindowsDefenderApplicationControl supplemental policy's version.
func (m *WindowsDefenderApplicationControlSupplementalPolicy) SetVersion(value *string)() {
    if m != nil {
        m.version = value
    }
}
