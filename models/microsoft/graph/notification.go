package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type Notification struct {
    Entity
    displayTimeToLive *int32;
    expirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    groupName *string;
    payload *PayloadTypes;
    priority *Priority;
    targetHostName *string;
    targetPolicy *TargetPolicyEndpoints;
}
func NewNotification()(*Notification) {
    m := &Notification{
        Entity: *NewEntity(),
    }
    return m
}
func (m *Notification) GetDisplayTimeToLive()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.displayTimeToLive
    }
}
func (m *Notification) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.expirationDateTime
    }
}
func (m *Notification) GetGroupName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.groupName
    }
}
func (m *Notification) GetPayload()(*PayloadTypes) {
    if m == nil {
        return nil
    } else {
        return m.payload
    }
}
func (m *Notification) GetPriority()(*Priority) {
    if m == nil {
        return nil
    } else {
        return m.priority
    }
}
func (m *Notification) GetTargetHostName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetHostName
    }
}
func (m *Notification) GetTargetPolicy()(*TargetPolicyEndpoints) {
    if m == nil {
        return nil
    } else {
        return m.targetPolicy
    }
}
func (m *Notification) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayTimeToLive"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetDisplayTimeToLive(val)
        return nil
    }
    res["expirationDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetExpirationDateTime(val)
        return nil
    }
    res["groupName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetGroupName(val)
        return nil
    }
    res["payload"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPayloadTypes() })
        if err != nil {
            return err
        }
        m.SetPayload(val.(*PayloadTypes))
        return nil
    }
    res["priority"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePriority)
        if err != nil {
            return err
        }
        cast := val.(Priority)
        m.SetPriority(&cast)
        return nil
    }
    res["targetHostName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTargetHostName(val)
        return nil
    }
    res["targetPolicy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTargetPolicyEndpoints() })
        if err != nil {
            return err
        }
        m.SetTargetPolicy(val.(*TargetPolicyEndpoints))
        return nil
    }
    return res
}
func (m *Notification) IsNil()(bool) {
    return m == nil
}
func (m *Notification) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("displayTimeToLive", m.GetDisplayTimeToLive())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("expirationDateTime", m.GetExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("groupName", m.GetGroupName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("payload", m.GetPayload())
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
    {
        err = writer.WriteStringValue("targetHostName", m.GetTargetHostName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("targetPolicy", m.GetTargetPolicy())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *Notification) SetDisplayTimeToLive(value *int32)() {
    m.displayTimeToLive = value
}
func (m *Notification) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expirationDateTime = value
}
func (m *Notification) SetGroupName(value *string)() {
    m.groupName = value
}
func (m *Notification) SetPayload(value *PayloadTypes)() {
    m.payload = value
}
func (m *Notification) SetPriority(value *Priority)() {
    m.priority = value
}
func (m *Notification) SetTargetHostName(value *string)() {
    m.targetHostName = value
}
func (m *Notification) SetTargetPolicy(value *TargetPolicyEndpoints)() {
    m.targetPolicy = value
}
