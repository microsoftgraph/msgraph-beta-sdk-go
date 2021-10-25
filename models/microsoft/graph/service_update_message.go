package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ServiceUpdateMessage struct {
    ServiceAnnouncementBase
    actionRequiredByDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    body *ItemBody;
    category *ServiceUpdateCategory;
    isMajorChange *bool;
    services []string;
    severity *ServiceUpdateSeverity;
    tags []string;
    viewPoint *ServiceUpdateMessageViewpoint;
}
func NewServiceUpdateMessage()(*ServiceUpdateMessage) {
    m := &ServiceUpdateMessage{
        ServiceAnnouncementBase: *NewServiceAnnouncementBase(),
    }
    return m
}
func (m *ServiceUpdateMessage) GetActionRequiredByDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.actionRequiredByDateTime
    }
}
func (m *ServiceUpdateMessage) GetBody()(*ItemBody) {
    if m == nil {
        return nil
    } else {
        return m.body
    }
}
func (m *ServiceUpdateMessage) GetCategory()(*ServiceUpdateCategory) {
    if m == nil {
        return nil
    } else {
        return m.category
    }
}
func (m *ServiceUpdateMessage) GetIsMajorChange()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isMajorChange
    }
}
func (m *ServiceUpdateMessage) GetServices()([]string) {
    if m == nil {
        return nil
    } else {
        return m.services
    }
}
func (m *ServiceUpdateMessage) GetSeverity()(*ServiceUpdateSeverity) {
    if m == nil {
        return nil
    } else {
        return m.severity
    }
}
func (m *ServiceUpdateMessage) GetTags()([]string) {
    if m == nil {
        return nil
    } else {
        return m.tags
    }
}
func (m *ServiceUpdateMessage) GetViewPoint()(*ServiceUpdateMessageViewpoint) {
    if m == nil {
        return nil
    } else {
        return m.viewPoint
    }
}
func (m *ServiceUpdateMessage) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ServiceAnnouncementBase.GetFieldDeserializers()
    res["actionRequiredByDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetActionRequiredByDateTime(val)
        return nil
    }
    res["body"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItemBody() })
        if err != nil {
            return err
        }
        m.SetBody(val.(*ItemBody))
        return nil
    }
    res["category"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseServiceUpdateCategory)
        if err != nil {
            return err
        }
        cast := val.(ServiceUpdateCategory)
        m.SetCategory(&cast)
        return nil
    }
    res["isMajorChange"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsMajorChange(val)
        return nil
    }
    res["services"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetServices(res)
        return nil
    }
    res["severity"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseServiceUpdateSeverity)
        if err != nil {
            return err
        }
        cast := val.(ServiceUpdateSeverity)
        m.SetSeverity(&cast)
        return nil
    }
    res["tags"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetTags(res)
        return nil
    }
    res["viewPoint"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewServiceUpdateMessageViewpoint() })
        if err != nil {
            return err
        }
        m.SetViewPoint(val.(*ServiceUpdateMessageViewpoint))
        return nil
    }
    return res
}
func (m *ServiceUpdateMessage) IsNil()(bool) {
    return m == nil
}
func (m *ServiceUpdateMessage) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.ServiceAnnouncementBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("actionRequiredByDateTime", m.GetActionRequiredByDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("body", m.GetBody())
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
        err = writer.WriteBoolValue("isMajorChange", m.GetIsMajorChange())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("services", m.GetServices())
        if err != nil {
            return err
        }
    }
    if m.GetSeverity() != nil {
        cast := m.GetSeverity().String()
        err = writer.WriteStringValue("severity", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("tags", m.GetTags())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("viewPoint", m.GetViewPoint())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *ServiceUpdateMessage) SetActionRequiredByDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.actionRequiredByDateTime = value
}
func (m *ServiceUpdateMessage) SetBody(value *ItemBody)() {
    m.body = value
}
func (m *ServiceUpdateMessage) SetCategory(value *ServiceUpdateCategory)() {
    m.category = value
}
func (m *ServiceUpdateMessage) SetIsMajorChange(value *bool)() {
    m.isMajorChange = value
}
func (m *ServiceUpdateMessage) SetServices(value []string)() {
    m.services = value
}
func (m *ServiceUpdateMessage) SetSeverity(value *ServiceUpdateSeverity)() {
    m.severity = value
}
func (m *ServiceUpdateMessage) SetTags(value []string)() {
    m.tags = value
}
func (m *ServiceUpdateMessage) SetViewPoint(value *ServiceUpdateMessageViewpoint)() {
    m.viewPoint = value
}
