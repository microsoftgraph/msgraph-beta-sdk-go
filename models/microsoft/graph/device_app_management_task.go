package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeviceAppManagementTask struct {
    Entity
    assignedTo *string;
    category *DeviceAppManagementTaskCategory;
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    creator *string;
    creatorNotes *string;
    description *string;
    displayName *string;
    dueDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    priority *DeviceAppManagementTaskPriority;
    status *DeviceAppManagementTaskStatus;
}
func NewDeviceAppManagementTask()(*DeviceAppManagementTask) {
    m := &DeviceAppManagementTask{
        Entity: *NewEntity(),
    }
    return m
}
func (m *DeviceAppManagementTask) GetAssignedTo()(*string) {
    if m == nil {
        return nil
    } else {
        return m.assignedTo
    }
}
func (m *DeviceAppManagementTask) GetCategory()(*DeviceAppManagementTaskCategory) {
    if m == nil {
        return nil
    } else {
        return m.category
    }
}
func (m *DeviceAppManagementTask) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
func (m *DeviceAppManagementTask) GetCreator()(*string) {
    if m == nil {
        return nil
    } else {
        return m.creator
    }
}
func (m *DeviceAppManagementTask) GetCreatorNotes()(*string) {
    if m == nil {
        return nil
    } else {
        return m.creatorNotes
    }
}
func (m *DeviceAppManagementTask) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
func (m *DeviceAppManagementTask) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *DeviceAppManagementTask) GetDueDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.dueDateTime
    }
}
func (m *DeviceAppManagementTask) GetPriority()(*DeviceAppManagementTaskPriority) {
    if m == nil {
        return nil
    } else {
        return m.priority
    }
}
func (m *DeviceAppManagementTask) GetStatus()(*DeviceAppManagementTaskStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
func (m *DeviceAppManagementTask) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignedTo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAssignedTo(val)
        return nil
    }
    res["category"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceAppManagementTaskCategory)
        if err != nil {
            return err
        }
        cast := val.(DeviceAppManagementTaskCategory)
        m.SetCategory(&cast)
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
    res["creator"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCreator(val)
        return nil
    }
    res["creatorNotes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCreatorNotes(val)
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
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["dueDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetDueDateTime(val)
        return nil
    }
    res["priority"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceAppManagementTaskPriority)
        if err != nil {
            return err
        }
        cast := val.(DeviceAppManagementTaskPriority)
        m.SetPriority(&cast)
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceAppManagementTaskStatus)
        if err != nil {
            return err
        }
        cast := val.(DeviceAppManagementTaskStatus)
        m.SetStatus(&cast)
        return nil
    }
    return res
}
func (m *DeviceAppManagementTask) IsNil()(bool) {
    return m == nil
}
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
        cast := m.GetCategory().String()
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
        cast := m.GetPriority().String()
        err = writer.WriteStringValue("priority", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := m.GetStatus().String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *DeviceAppManagementTask) SetAssignedTo(value *string)() {
    m.assignedTo = value
}
func (m *DeviceAppManagementTask) SetCategory(value *DeviceAppManagementTaskCategory)() {
    m.category = value
}
func (m *DeviceAppManagementTask) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
func (m *DeviceAppManagementTask) SetCreator(value *string)() {
    m.creator = value
}
func (m *DeviceAppManagementTask) SetCreatorNotes(value *string)() {
    m.creatorNotes = value
}
func (m *DeviceAppManagementTask) SetDescription(value *string)() {
    m.description = value
}
func (m *DeviceAppManagementTask) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *DeviceAppManagementTask) SetDueDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.dueDateTime = value
}
func (m *DeviceAppManagementTask) SetPriority(value *DeviceAppManagementTaskPriority)() {
    m.priority = value
}
func (m *DeviceAppManagementTask) SetStatus(value *DeviceAppManagementTaskStatus)() {
    m.status = value
}
