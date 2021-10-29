package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AuditEvent struct {
    Entity
    // Friendly name of the activity.
    activity *string;
    // The date time in UTC when the activity was performed.
    activityDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The HTTP operation type of the activity.
    activityOperationType *string;
    // The result of the activity.
    activityResult *string;
    // The type of activity that was being performed.
    activityType *string;
    // AAD user and application that are associated with the audit event.
    actor *AuditActor;
    // Audit category.
    category *string;
    // Component name.
    componentName *string;
    // The client request Id that is used to correlate activity within the system.
    correlationId *string;
    // Event display name.
    displayName *string;
    // Resources being modified.
    resources []AuditResource;
}
// Instantiates a new auditEvent and sets the default values.
func NewAuditEvent()(*AuditEvent) {
    m := &AuditEvent{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the activity property value. Friendly name of the activity.
func (m *AuditEvent) GetActivity()(*string) {
    if m == nil {
        return nil
    } else {
        return m.activity
    }
}
// Gets the activityDateTime property value. The date time in UTC when the activity was performed.
func (m *AuditEvent) GetActivityDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.activityDateTime
    }
}
// Gets the activityOperationType property value. The HTTP operation type of the activity.
func (m *AuditEvent) GetActivityOperationType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.activityOperationType
    }
}
// Gets the activityResult property value. The result of the activity.
func (m *AuditEvent) GetActivityResult()(*string) {
    if m == nil {
        return nil
    } else {
        return m.activityResult
    }
}
// Gets the activityType property value. The type of activity that was being performed.
func (m *AuditEvent) GetActivityType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.activityType
    }
}
// Gets the actor property value. AAD user and application that are associated with the audit event.
func (m *AuditEvent) GetActor()(*AuditActor) {
    if m == nil {
        return nil
    } else {
        return m.actor
    }
}
// Gets the category property value. Audit category.
func (m *AuditEvent) GetCategory()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category
    }
}
// Gets the componentName property value. Component name.
func (m *AuditEvent) GetComponentName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.componentName
    }
}
// Gets the correlationId property value. The client request Id that is used to correlate activity within the system.
func (m *AuditEvent) GetCorrelationId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.correlationId
    }
}
// Gets the displayName property value. Event display name.
func (m *AuditEvent) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the resources property value. Resources being modified.
func (m *AuditEvent) GetResources()([]AuditResource) {
    if m == nil {
        return nil
    } else {
        return m.resources
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the activity property value. Friendly name of the activity.
// Parameters:
//  - value : Value to set for the activity property.
func (m *AuditEvent) SetActivity(value *string)() {
    m.activity = value
}
// Sets the activityDateTime property value. The date time in UTC when the activity was performed.
// Parameters:
//  - value : Value to set for the activityDateTime property.
func (m *AuditEvent) SetActivityDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.activityDateTime = value
}
// Sets the activityOperationType property value. The HTTP operation type of the activity.
// Parameters:
//  - value : Value to set for the activityOperationType property.
func (m *AuditEvent) SetActivityOperationType(value *string)() {
    m.activityOperationType = value
}
// Sets the activityResult property value. The result of the activity.
// Parameters:
//  - value : Value to set for the activityResult property.
func (m *AuditEvent) SetActivityResult(value *string)() {
    m.activityResult = value
}
// Sets the activityType property value. The type of activity that was being performed.
// Parameters:
//  - value : Value to set for the activityType property.
func (m *AuditEvent) SetActivityType(value *string)() {
    m.activityType = value
}
// Sets the actor property value. AAD user and application that are associated with the audit event.
// Parameters:
//  - value : Value to set for the actor property.
func (m *AuditEvent) SetActor(value *AuditActor)() {
    m.actor = value
}
// Sets the category property value. Audit category.
// Parameters:
//  - value : Value to set for the category property.
func (m *AuditEvent) SetCategory(value *string)() {
    m.category = value
}
// Sets the componentName property value. Component name.
// Parameters:
//  - value : Value to set for the componentName property.
func (m *AuditEvent) SetComponentName(value *string)() {
    m.componentName = value
}
// Sets the correlationId property value. The client request Id that is used to correlate activity within the system.
// Parameters:
//  - value : Value to set for the correlationId property.
func (m *AuditEvent) SetCorrelationId(value *string)() {
    m.correlationId = value
}
// Sets the displayName property value. Event display name.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *AuditEvent) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the resources property value. Resources being modified.
// Parameters:
//  - value : Value to set for the resources property.
func (m *AuditEvent) SetResources(value []AuditResource)() {
    m.resources = value
}
