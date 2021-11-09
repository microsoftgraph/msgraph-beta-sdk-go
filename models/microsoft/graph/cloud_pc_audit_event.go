package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type CloudPcAuditEvent struct {
    Entity
    // Friendly name of the activity. Optional.
    activity *string;
    // The date time in UTC when the activity was performed. Read-only.
    activityDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The HTTP operation type of the activity. Possible values include create, delete, patch and other. Read-only.
    activityOperationType *CloudPcAuditActivityOperationType;
    // The result of the activity. Read-only.
    activityResult *CloudPcAuditActivityResult;
    // The type of activity that was performed. Read-only.
    activityType *string;
    // 
    actor *CloudPcAuditActor;
    // Audit category. Read-only.
    category *CloudPcAuditCategory;
    // Component name. Read-only.
    componentName *string;
    // The client request identifier, used to correlate activity within the system. Read-only.
    correlationId *string;
    // Event display name. Read-only.
    displayName *string;
    // List of cloudPcAuditResource objects. Read-only.
    resources []CloudPcAuditResource;
}
// Instantiates a new cloudPcAuditEvent and sets the default values.
func NewCloudPcAuditEvent()(*CloudPcAuditEvent) {
    m := &CloudPcAuditEvent{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the activity property value. Friendly name of the activity. Optional.
func (m *CloudPcAuditEvent) GetActivity()(*string) {
    if m == nil {
        return nil
    } else {
        return m.activity
    }
}
// Gets the activityDateTime property value. The date time in UTC when the activity was performed. Read-only.
func (m *CloudPcAuditEvent) GetActivityDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.activityDateTime
    }
}
// Gets the activityOperationType property value. The HTTP operation type of the activity. Possible values include create, delete, patch and other. Read-only.
func (m *CloudPcAuditEvent) GetActivityOperationType()(*CloudPcAuditActivityOperationType) {
    if m == nil {
        return nil
    } else {
        return m.activityOperationType
    }
}
// Gets the activityResult property value. The result of the activity. Read-only.
func (m *CloudPcAuditEvent) GetActivityResult()(*CloudPcAuditActivityResult) {
    if m == nil {
        return nil
    } else {
        return m.activityResult
    }
}
// Gets the activityType property value. The type of activity that was performed. Read-only.
func (m *CloudPcAuditEvent) GetActivityType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.activityType
    }
}
// Gets the actor property value. 
func (m *CloudPcAuditEvent) GetActor()(*CloudPcAuditActor) {
    if m == nil {
        return nil
    } else {
        return m.actor
    }
}
// Gets the category property value. Audit category. Read-only.
func (m *CloudPcAuditEvent) GetCategory()(*CloudPcAuditCategory) {
    if m == nil {
        return nil
    } else {
        return m.category
    }
}
// Gets the componentName property value. Component name. Read-only.
func (m *CloudPcAuditEvent) GetComponentName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.componentName
    }
}
// Gets the correlationId property value. The client request identifier, used to correlate activity within the system. Read-only.
func (m *CloudPcAuditEvent) GetCorrelationId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.correlationId
    }
}
// Gets the displayName property value. Event display name. Read-only.
func (m *CloudPcAuditEvent) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the resources property value. List of cloudPcAuditResource objects. Read-only.
func (m *CloudPcAuditEvent) GetResources()([]CloudPcAuditResource) {
    if m == nil {
        return nil
    } else {
        return m.resources
    }
}
// The deserialization information for the current model
func (m *CloudPcAuditEvent) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activity"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivity(val)
        }
        return nil
    }
    res["activityDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivityDateTime(val)
        }
        return nil
    }
    res["activityOperationType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcAuditActivityOperationType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(CloudPcAuditActivityOperationType)
            m.SetActivityOperationType(&cast)
        }
        return nil
    }
    res["activityResult"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcAuditActivityResult)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(CloudPcAuditActivityResult)
            m.SetActivityResult(&cast)
        }
        return nil
    }
    res["activityType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivityType(val)
        }
        return nil
    }
    res["actor"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCloudPcAuditActor() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActor(val.(*CloudPcAuditActor))
        }
        return nil
    }
    res["category"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcAuditCategory)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(CloudPcAuditCategory)
            m.SetCategory(&cast)
        }
        return nil
    }
    res["componentName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComponentName(val)
        }
        return nil
    }
    res["correlationId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCorrelationId(val)
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
    res["resources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCloudPcAuditResource() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPcAuditResource, len(val))
            for i, v := range val {
                res[i] = *(v.(*CloudPcAuditResource))
            }
            m.SetResources(res)
        }
        return nil
    }
    return res
}
func (m *CloudPcAuditEvent) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the activity property value. Friendly name of the activity. Optional.
// Parameters:
//  - value : Value to set for the activity property.
func (m *CloudPcAuditEvent) SetActivity(value *string)() {
    m.activity = value
}
// Sets the activityDateTime property value. The date time in UTC when the activity was performed. Read-only.
// Parameters:
//  - value : Value to set for the activityDateTime property.
func (m *CloudPcAuditEvent) SetActivityDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.activityDateTime = value
}
// Sets the activityOperationType property value. The HTTP operation type of the activity. Possible values include create, delete, patch and other. Read-only.
// Parameters:
//  - value : Value to set for the activityOperationType property.
func (m *CloudPcAuditEvent) SetActivityOperationType(value *CloudPcAuditActivityOperationType)() {
    m.activityOperationType = value
}
// Sets the activityResult property value. The result of the activity. Read-only.
// Parameters:
//  - value : Value to set for the activityResult property.
func (m *CloudPcAuditEvent) SetActivityResult(value *CloudPcAuditActivityResult)() {
    m.activityResult = value
}
// Sets the activityType property value. The type of activity that was performed. Read-only.
// Parameters:
//  - value : Value to set for the activityType property.
func (m *CloudPcAuditEvent) SetActivityType(value *string)() {
    m.activityType = value
}
// Sets the actor property value. 
// Parameters:
//  - value : Value to set for the actor property.
func (m *CloudPcAuditEvent) SetActor(value *CloudPcAuditActor)() {
    m.actor = value
}
// Sets the category property value. Audit category. Read-only.
// Parameters:
//  - value : Value to set for the category property.
func (m *CloudPcAuditEvent) SetCategory(value *CloudPcAuditCategory)() {
    m.category = value
}
// Sets the componentName property value. Component name. Read-only.
// Parameters:
//  - value : Value to set for the componentName property.
func (m *CloudPcAuditEvent) SetComponentName(value *string)() {
    m.componentName = value
}
// Sets the correlationId property value. The client request identifier, used to correlate activity within the system. Read-only.
// Parameters:
//  - value : Value to set for the correlationId property.
func (m *CloudPcAuditEvent) SetCorrelationId(value *string)() {
    m.correlationId = value
}
// Sets the displayName property value. Event display name. Read-only.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *CloudPcAuditEvent) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the resources property value. List of cloudPcAuditResource objects. Read-only.
// Parameters:
//  - value : Value to set for the resources property.
func (m *CloudPcAuditEvent) SetResources(value []CloudPcAuditResource)() {
    m.resources = value
}
