package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type EducationAssignmentDefaults struct {
    Entity
    addedStudentAction *EducationAddedStudentAction;
    addToCalendarAction *EducationAddToCalendarOptions;
    dueTime *string;
    notificationChannelUrl *string;
}
func NewEducationAssignmentDefaults()(*EducationAssignmentDefaults) {
    m := &EducationAssignmentDefaults{
        Entity: *NewEntity(),
    }
    return m
}
func (m *EducationAssignmentDefaults) GetAddedStudentAction()(*EducationAddedStudentAction) {
    if m == nil {
        return nil
    } else {
        return m.addedStudentAction
    }
}
func (m *EducationAssignmentDefaults) GetAddToCalendarAction()(*EducationAddToCalendarOptions) {
    if m == nil {
        return nil
    } else {
        return m.addToCalendarAction
    }
}
func (m *EducationAssignmentDefaults) GetDueTime()(*string) {
    if m == nil {
        return nil
    } else {
        return m.dueTime
    }
}
func (m *EducationAssignmentDefaults) GetNotificationChannelUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.notificationChannelUrl
    }
}
func (m *EducationAssignmentDefaults) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["addedStudentAction"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseEducationAddedStudentAction)
        if err != nil {
            return err
        }
        cast := val.(EducationAddedStudentAction)
        m.SetAddedStudentAction(&cast)
        return nil
    }
    res["addToCalendarAction"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseEducationAddToCalendarOptions)
        if err != nil {
            return err
        }
        cast := val.(EducationAddToCalendarOptions)
        m.SetAddToCalendarAction(&cast)
        return nil
    }
    res["dueTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDueTime(val)
        return nil
    }
    res["notificationChannelUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetNotificationChannelUrl(val)
        return nil
    }
    return res
}
func (m *EducationAssignmentDefaults) IsNil()(bool) {
    return m == nil
}
func (m *EducationAssignmentDefaults) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAddedStudentAction() != nil {
        cast := m.GetAddedStudentAction().String()
        err = writer.WriteStringValue("addedStudentAction", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAddToCalendarAction() != nil {
        cast := m.GetAddToCalendarAction().String()
        err = writer.WriteStringValue("addToCalendarAction", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("dueTime", m.GetDueTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("notificationChannelUrl", m.GetNotificationChannelUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *EducationAssignmentDefaults) SetAddedStudentAction(value *EducationAddedStudentAction)() {
    m.addedStudentAction = value
}
func (m *EducationAssignmentDefaults) SetAddToCalendarAction(value *EducationAddToCalendarOptions)() {
    m.addToCalendarAction = value
}
func (m *EducationAssignmentDefaults) SetDueTime(value *string)() {
    m.dueTime = value
}
func (m *EducationAssignmentDefaults) SetNotificationChannelUrl(value *string)() {
    m.notificationChannelUrl = value
}
