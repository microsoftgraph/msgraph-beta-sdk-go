package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AuditEvent struct {
    Entity
    activity *string;
    activityDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    activityOperationType *string;
    activityResult *string;
    activityType *string;
    actor *AuditActor;
    category *string;
    componentName *string;
    correlationId *string;
    displayName *string;
    resources []AuditResource;
}
func NewAuditEvent()(*AuditEvent) {
    m := &AuditEvent{
        Entity: *NewEntity(),
    }
    return m
}
func (m *AuditEvent) GetActivity()(*string) {
    if m == nil {
        return nil
    } else {
        return m.activity
    }
}
func (m *AuditEvent) GetActivityDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.activityDateTime
    }
}
func (m *AuditEvent) GetActivityOperationType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.activityOperationType
    }
}
func (m *AuditEvent) GetActivityResult()(*string) {
    if m == nil {
        return nil
    } else {
        return m.activityResult
    }
}
func (m *AuditEvent) GetActivityType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.activityType
    }
}
func (m *AuditEvent) GetActor()(*AuditActor) {
    if m == nil {
        return nil
    } else {
        return m.actor
    }
}
func (m *AuditEvent) GetCategory()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category
    }
}
func (m *AuditEvent) GetComponentName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.componentName
    }
}
func (m *AuditEvent) GetCorrelationId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.correlationId
    }
}
func (m *AuditEvent) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *AuditEvent) GetResources()([]AuditResource) {
    if m == nil {
        return nil
    } else {
        return m.resources
    }
}
func (m *AuditEvent) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activity"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetActivity(val)
        return nil
    }
    res["activityDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetActivityDateTime(val)
        return nil
    }
    res["activityOperationType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetActivityOperationType(val)
        return nil
    }
    res["activityResult"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetActivityResult(val)
        return nil
    }
    res["activityType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetActivityType(val)
        return nil
    }
    res["actor"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAuditActor() })
        if err != nil {
            return err
        }
        m.SetActor(val.(*AuditActor))
        return nil
    }
    res["category"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCategory(val)
        return nil
    }
    res["componentName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetComponentName(val)
        return nil
    }
    res["correlationId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCorrelationId(val)
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
    res["resources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAuditResource() })
        if err != nil {
            return err
        }
        res := make([]AuditResource, len(val))
        for i, v := range val {
            res[i] = *(v.(*AuditResource))
        }
        m.SetResources(res)
        return nil
    }
    return res
}
func (m *AuditEvent) IsNil()(bool) {
    return m == nil
}
func (m *AuditEvent) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("activity", m.GetActivity())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("activityDateTime", m.GetActivityDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("activityOperationType", m.GetActivityOperationType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("activityResult", m.GetActivityResult())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("activityType", m.GetActivityType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("actor", m.GetActor())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("category", m.GetCategory())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("componentName", m.GetComponentName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("correlationId", m.GetCorrelationId())
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetResources()))
        for i, v := range m.GetResources() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("resources", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *AuditEvent) SetActivity(value *string)() {
    m.activity = value
}
func (m *AuditEvent) SetActivityDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.activityDateTime = value
}
func (m *AuditEvent) SetActivityOperationType(value *string)() {
    m.activityOperationType = value
}
func (m *AuditEvent) SetActivityResult(value *string)() {
    m.activityResult = value
}
func (m *AuditEvent) SetActivityType(value *string)() {
    m.activityType = value
}
func (m *AuditEvent) SetActor(value *AuditActor)() {
    m.actor = value
}
func (m *AuditEvent) SetCategory(value *string)() {
    m.category = value
}
func (m *AuditEvent) SetComponentName(value *string)() {
    m.componentName = value
}
func (m *AuditEvent) SetCorrelationId(value *string)() {
    m.correlationId = value
}
func (m *AuditEvent) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *AuditEvent) SetResources(value []AuditResource)() {
    m.resources = value
}
