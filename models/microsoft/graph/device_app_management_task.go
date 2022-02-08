package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceAppManagementTask 
type DeviceAppManagementTask struct {
    Entity
    // The name or email of the admin this task is assigned to.
    assignedTo *string;
    // The category. Possible values are: unknown, advancedThreatProtection.
    category *DeviceAppManagementTaskCategory;
    // The created date.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The email address of the creator.
    creator *string;
    // Notes from the creator.
    creatorNotes *string;
    // The description.
    description *string;
    // The name.
    displayName *string;
    // The due date.
    dueDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The priority. Possible values are: none, high, low.
    priority *DeviceAppManagementTaskPriority;
    // The status. Possible values are: unknown, pending, active, completed, rejected.
    status *DeviceAppManagementTaskStatus;
}
// NewDeviceAppManagementTask instantiates a new deviceAppManagementTask and sets the default values.
func NewDeviceAppManagementTask()(*DeviceAppManagementTask) {
    m := &DeviceAppManagementTask{
        Entity: *NewEntity(),
    }
    return m
}
// GetAssignedTo gets the assignedTo property value. The name or email of the admin this task is assigned to.
func (m *DeviceAppManagementTask) GetAssignedTo()(*string) {
    if m == nil {
        return nil
    } else {
        return m.assignedTo
    }
}
// GetCategory gets the category property value. The category. Possible values are: unknown, advancedThreatProtection.
func (m *DeviceAppManagementTask) GetCategory()(*DeviceAppManagementTaskCategory) {
    if m == nil {
        return nil
    } else {
        return m.category
    }
}
// GetCreatedDateTime gets the createdDateTime property value. The created date.
func (m *DeviceAppManagementTask) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetCreator gets the creator property value. The email address of the creator.
func (m *DeviceAppManagementTask) GetCreator()(*string) {
    if m == nil {
        return nil
    } else {
        return m.creator
    }
}
// GetCreatorNotes gets the creatorNotes property value. Notes from the creator.
func (m *DeviceAppManagementTask) GetCreatorNotes()(*string) {
    if m == nil {
        return nil
    } else {
        return m.creatorNotes
    }
}
// GetDescription gets the description property value. The description.
func (m *DeviceAppManagementTask) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. The name.
func (m *DeviceAppManagementTask) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetDueDateTime gets the dueDateTime property value. The due date.
func (m *DeviceAppManagementTask) GetDueDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.dueDateTime
    }
}
// GetPriority gets the priority property value. The priority. Possible values are: none, high, low.
func (m *DeviceAppManagementTask) GetPriority()(*DeviceAppManagementTaskPriority) {
    if m == nil {
        return nil
    } else {
        return m.priority
    }
}
// GetStatus gets the status property value. The status. Possible values are: unknown, pending, active, completed, rejected.
func (m *DeviceAppManagementTask) GetStatus()(*DeviceAppManagementTaskStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceAppManagementTask) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignedTo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignedTo(val)
        }
        return nil
    }
    res["category"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceAppManagementTaskCategory)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory(val.(*DeviceAppManagementTaskCategory))
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
    res["creator"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreator(val)
        }
        return nil
    }
    res["creatorNotes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatorNotes(val)
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
    res["dueDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDueDateTime(val)
        }
        return nil
    }
    res["priority"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceAppManagementTaskPriority)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPriority(val.(*DeviceAppManagementTaskPriority))
        }
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceAppManagementTaskStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*DeviceAppManagementTaskStatus))
        }
        return nil
    }
    return res
}
func (m *DeviceAppManagementTask) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DeviceAppManagementTask) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("assignedTo", m.GetAssignedTo())
        if err != nil {
            return err
        }
    }
    if m.GetCategory() != nil {
        cast := (*m.GetCategory()).String()
        err = writer.WriteStringValue("category", &cast)
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
        err = writer.WriteStringValue("creator", m.GetCreator())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("creatorNotes", m.GetCreatorNotes())
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
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("dueDateTime", m.GetDueDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetPriority() != nil {
        cast := (*m.GetPriority()).String()
        err = writer.WriteStringValue("priority", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignedTo sets the assignedTo property value. The name or email of the admin this task is assigned to.
func (m *DeviceAppManagementTask) SetAssignedTo(value *string)() {
    if m != nil {
        m.assignedTo = value
    }
}
// SetCategory sets the category property value. The category. Possible values are: unknown, advancedThreatProtection.
func (m *DeviceAppManagementTask) SetCategory(value *DeviceAppManagementTaskCategory)() {
    if m != nil {
        m.category = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The created date.
func (m *DeviceAppManagementTask) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetCreator sets the creator property value. The email address of the creator.
func (m *DeviceAppManagementTask) SetCreator(value *string)() {
    if m != nil {
        m.creator = value
    }
}
// SetCreatorNotes sets the creatorNotes property value. Notes from the creator.
func (m *DeviceAppManagementTask) SetCreatorNotes(value *string)() {
    if m != nil {
        m.creatorNotes = value
    }
}
// SetDescription sets the description property value. The description.
func (m *DeviceAppManagementTask) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. The name.
func (m *DeviceAppManagementTask) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetDueDateTime sets the dueDateTime property value. The due date.
func (m *DeviceAppManagementTask) SetDueDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.dueDateTime = value
    }
}
// SetPriority sets the priority property value. The priority. Possible values are: none, high, low.
func (m *DeviceAppManagementTask) SetPriority(value *DeviceAppManagementTaskPriority)() {
    if m != nil {
        m.priority = value
    }
}
// SetStatus sets the status property value. The status. Possible values are: unknown, pending, active, completed, rejected.
func (m *DeviceAppManagementTask) SetStatus(value *DeviceAppManagementTaskStatus)() {
    if m != nil {
        m.status = value
    }
}
