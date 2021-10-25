package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type CloudPcAuditEvent struct {
    Entity
    activity *string;
    activityDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    activityOperationType *CloudPcAuditActivityOperationType;
    activityResult *CloudPcAuditActivityResult;
    activityType *string;
    actor *CloudPcAuditActor;
    category *CloudPcAuditCategory;
    componentName *string;
    correlationId *string;
    displayName *string;
    resources []CloudPcAuditResource;
}
func NewCloudPcAuditEvent()(*CloudPcAuditEvent) {
    m := &CloudPcAuditEvent{
        Entity: *NewEntity(),
    }
    return m
}
func (m *CloudPcAuditEvent) GetActivity()(*string) {
    if m == nil {
        return nil
    } else {
        return m.activity
    }
}
func (m *CloudPcAuditEvent) GetActivityDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.activityDateTime
    }
}
func (m *CloudPcAuditEvent) GetActivityOperationType()(*CloudPcAuditActivityOperationType) {
    if m == nil {
        return nil
    } else {
        return m.activityOperationType
    }
}
func (m *CloudPcAuditEvent) GetActivityResult()(*CloudPcAuditActivityResult) {
    if m == nil {
        return nil
    } else {
        return m.activityResult
    }
}
func (m *CloudPcAuditEvent) GetActivityType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.activityType
    }
}
func (m *CloudPcAuditEvent) GetActor()(*CloudPcAuditActor) {
    if m == nil {
        return nil
    } else {
        return m.actor
    }
}
func (m *CloudPcAuditEvent) GetCategory()(*CloudPcAuditCategory) {
    if m == nil {
        return nil
    } else {
        return m.category
    }
}
func (m *CloudPcAuditEvent) GetComponentName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.componentName
    }
}
func (m *CloudPcAuditEvent) GetCorrelationId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.correlationId
    }
}
func (m *CloudPcAuditEvent) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *CloudPcAuditEvent) GetResources()([]CloudPcAuditResource) {
    if m == nil {
        return nil
    } else {
        return m.resources
    }
}
func (m *CloudPcAuditEvent) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
        val, err := n.GetEnumValue(ParseCloudPcAuditActivityOperationType)
        if err != nil {
            return err
        }
        cast := val.(CloudPcAuditActivityOperationType)
        m.SetActivityOperationType(&cast)
        return nil
    }
    res["activityResult"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcAuditActivityResult)
        if err != nil {
            return err
        }
        cast := val.(CloudPcAuditActivityResult)
        m.SetActivityResult(&cast)
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
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCloudPcAuditActor() })
        if err != nil {
            return err
        }
        m.SetActor(val.(*CloudPcAuditActor))
        return nil
    }
    res["category"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcAuditCategory)
        if err != nil {
            return err
        }
        cast := val.(CloudPcAuditCategory)
        m.SetCategory(&cast)
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
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCloudPcAuditResource() })
        if err != nil {
            return err
        }
        res := make([]CloudPcAuditResource, len(val))
        for i, v := range val {
            res[i] = *(v.(*CloudPcAuditResource))
        }
        m.SetResources(res)
        return nil
    }
    return res
}
func (m *CloudPcAuditEvent) IsNil()(bool) {
    return m == nil
}
func (m *CloudPcAuditEvent) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
    if m.GetActivityOperationType() != nil {
        cast := m.GetActivityOperationType().String()
        err = writer.WriteStringValue("activityOperationType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetActivityResult() != nil {
        cast := m.GetActivityResult().String()
        err = writer.WriteStringValue("activityResult", &cast)
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
    if m.GetCategory() != nil {
        cast := m.GetCategory().String()
        err = writer.WriteStringValue("category", &cast)
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
func (m *CloudPcAuditEvent) SetActivity(value *string)() {
    m.activity = value
}
func (m *CloudPcAuditEvent) SetActivityDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.activityDateTime = value
}
func (m *CloudPcAuditEvent) SetActivityOperationType(value *CloudPcAuditActivityOperationType)() {
    m.activityOperationType = value
}
func (m *CloudPcAuditEvent) SetActivityResult(value *CloudPcAuditActivityResult)() {
    m.activityResult = value
}
func (m *CloudPcAuditEvent) SetActivityType(value *string)() {
    m.activityType = value
}
func (m *CloudPcAuditEvent) SetActor(value *CloudPcAuditActor)() {
    m.actor = value
}
func (m *CloudPcAuditEvent) SetCategory(value *CloudPcAuditCategory)() {
    m.category = value
}
func (m *CloudPcAuditEvent) SetComponentName(value *string)() {
    m.componentName = value
}
func (m *CloudPcAuditEvent) SetCorrelationId(value *string)() {
    m.correlationId = value
}
func (m *CloudPcAuditEvent) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *CloudPcAuditEvent) SetResources(value []CloudPcAuditResource)() {
    m.resources = value
}
