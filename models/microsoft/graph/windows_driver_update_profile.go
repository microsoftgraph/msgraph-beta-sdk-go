package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WindowsDriverUpdateProfile 
type WindowsDriverUpdateProfile struct {
    Entity
    // Driver update profile approval type. For example, manual or automatic approval. Possible values are: manual, automatic.
    approvalType *DriverUpdateProfileApprovalType;
    // The list of group assignments of the profile.
    assignments []WindowsDriverUpdateProfileAssignment;
    // The date time that the profile was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Deployment deferral settings in days, only applicable when ApprovalType is set to automatic approval.
    deploymentDeferralInDays *int32;
    // The description of the profile which is specified by the user.
    description *string;
    // Number of devices reporting for this profile
    deviceReporting *int32;
    // The display name for the profile.
    displayName *string;
    // Driver inventories for this profile.
    driverInventories []WindowsDriverUpdateInventory;
    // The date time that the profile was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Number of new driver updates available for this profile.
    newUpdates *int32;
    // List of Scope Tags for this Driver Update entity.
    roleScopeTagIds []string;
}
// NewWindowsDriverUpdateProfile instantiates a new windowsDriverUpdateProfile and sets the default values.
func NewWindowsDriverUpdateProfile()(*WindowsDriverUpdateProfile) {
    m := &WindowsDriverUpdateProfile{
        Entity: *NewEntity(),
    }
    return m
}
// GetApprovalType gets the approvalType property value. Driver update profile approval type. For example, manual or automatic approval. Possible values are: manual, automatic.
func (m *WindowsDriverUpdateProfile) GetApprovalType()(*DriverUpdateProfileApprovalType) {
    if m == nil {
        return nil
    } else {
        return m.approvalType
    }
}
// GetAssignments gets the assignments property value. The list of group assignments of the profile.
func (m *WindowsDriverUpdateProfile) GetAssignments()([]WindowsDriverUpdateProfileAssignment) {
    if m == nil {
        return nil
    } else {
        return m.assignments
    }
}
// GetCreatedDateTime gets the createdDateTime property value. The date time that the profile was created.
func (m *WindowsDriverUpdateProfile) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetDeploymentDeferralInDays gets the deploymentDeferralInDays property value. Deployment deferral settings in days, only applicable when ApprovalType is set to automatic approval.
func (m *WindowsDriverUpdateProfile) GetDeploymentDeferralInDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.deploymentDeferralInDays
    }
}
// GetDescription gets the description property value. The description of the profile which is specified by the user.
func (m *WindowsDriverUpdateProfile) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDeviceReporting gets the deviceReporting property value. Number of devices reporting for this profile
func (m *WindowsDriverUpdateProfile) GetDeviceReporting()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.deviceReporting
    }
}
// GetDisplayName gets the displayName property value. The display name for the profile.
func (m *WindowsDriverUpdateProfile) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetDriverInventories gets the driverInventories property value. Driver inventories for this profile.
func (m *WindowsDriverUpdateProfile) GetDriverInventories()([]WindowsDriverUpdateInventory) {
    if m == nil {
        return nil
    } else {
        return m.driverInventories
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date time that the profile was last modified.
func (m *WindowsDriverUpdateProfile) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetNewUpdates gets the newUpdates property value. Number of new driver updates available for this profile.
func (m *WindowsDriverUpdateProfile) GetNewUpdates()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.newUpdates
    }
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. List of Scope Tags for this Driver Update entity.
func (m *WindowsDriverUpdateProfile) GetRoleScopeTagIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.roleScopeTagIds
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsDriverUpdateProfile) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["approvalType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDriverUpdateProfileApprovalType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApprovalType(val.(*DriverUpdateProfileApprovalType))
        }
        return nil
    }
    res["assignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWindowsDriverUpdateProfileAssignment() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsDriverUpdateProfileAssignment, len(val))
            for i, v := range val {
                res[i] = *(v.(*WindowsDriverUpdateProfileAssignment))
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
    res["deploymentDeferralInDays"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeploymentDeferralInDays(val)
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
    res["deviceReporting"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceReporting(val)
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
    res["driverInventories"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWindowsDriverUpdateInventory() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsDriverUpdateInventory, len(val))
            for i, v := range val {
                res[i] = *(v.(*WindowsDriverUpdateInventory))
            }
            m.SetDriverInventories(res)
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
    res["newUpdates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNewUpdates(val)
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
    return res
}
func (m *WindowsDriverUpdateProfile) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *WindowsDriverUpdateProfile) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDriverInventories()))
        for i, v := range m.GetDriverInventories() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("driverInventories", cast)
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
// SetApprovalType sets the approvalType property value. Driver update profile approval type. For example, manual or automatic approval. Possible values are: manual, automatic.
func (m *WindowsDriverUpdateProfile) SetApprovalType(value *DriverUpdateProfileApprovalType)() {
    if m != nil {
        m.approvalType = value
    }
}
// SetAssignments sets the assignments property value. The list of group assignments of the profile.
func (m *WindowsDriverUpdateProfile) SetAssignments(value []WindowsDriverUpdateProfileAssignment)() {
    if m != nil {
        m.assignments = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The date time that the profile was created.
func (m *WindowsDriverUpdateProfile) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetDeploymentDeferralInDays sets the deploymentDeferralInDays property value. Deployment deferral settings in days, only applicable when ApprovalType is set to automatic approval.
func (m *WindowsDriverUpdateProfile) SetDeploymentDeferralInDays(value *int32)() {
    if m != nil {
        m.deploymentDeferralInDays = value
    }
}
// SetDescription sets the description property value. The description of the profile which is specified by the user.
func (m *WindowsDriverUpdateProfile) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDeviceReporting sets the deviceReporting property value. Number of devices reporting for this profile
func (m *WindowsDriverUpdateProfile) SetDeviceReporting(value *int32)() {
    if m != nil {
        m.deviceReporting = value
    }
}
// SetDisplayName sets the displayName property value. The display name for the profile.
func (m *WindowsDriverUpdateProfile) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetDriverInventories sets the driverInventories property value. Driver inventories for this profile.
func (m *WindowsDriverUpdateProfile) SetDriverInventories(value []WindowsDriverUpdateInventory)() {
    if m != nil {
        m.driverInventories = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date time that the profile was last modified.
func (m *WindowsDriverUpdateProfile) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetNewUpdates sets the newUpdates property value. Number of new driver updates available for this profile.
func (m *WindowsDriverUpdateProfile) SetNewUpdates(value *int32)() {
    if m != nil {
        m.newUpdates = value
    }
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. List of Scope Tags for this Driver Update entity.
func (m *WindowsDriverUpdateProfile) SetRoleScopeTagIds(value []string)() {
    if m != nil {
        m.roleScopeTagIds = value
    }
}
