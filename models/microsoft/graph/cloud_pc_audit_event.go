package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CloudPcAuditEvent 
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
// NewCloudPcAuditEvent instantiates a new cloudPcAuditEvent and sets the default values.
func NewCloudPcAuditEvent()(*CloudPcAuditEvent) {
    m := &CloudPcAuditEvent{
        Entity: *NewEntity(),
    }
    return m
}
// GetActivity gets the activity property value. Friendly name of the activity. Optional.
func (m *CloudPcAuditEvent) GetActivity()(*string) {
    if m == nil {
        return nil
    } else {
        return m.activity
    }
}
// GetActivityDateTime gets the activityDateTime property value. The date time in UTC when the activity was performed. Read-only.
func (m *CloudPcAuditEvent) GetActivityDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.activityDateTime
    }
}
// GetActivityOperationType gets the activityOperationType property value. The HTTP operation type of the activity. Possible values include create, delete, patch and other. Read-only.
func (m *CloudPcAuditEvent) GetActivityOperationType()(*CloudPcAuditActivityOperationType) {
    if m == nil {
        return nil
    } else {
        return m.activityOperationType
    }
}
// GetActivityResult gets the activityResult property value. The result of the activity. Read-only.
func (m *CloudPcAuditEvent) GetActivityResult()(*CloudPcAuditActivityResult) {
    if m == nil {
        return nil
    } else {
        return m.activityResult
    }
}
// GetActivityType gets the activityType property value. The type of activity that was performed. Read-only.
func (m *CloudPcAuditEvent) GetActivityType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.activityType
    }
}
// GetActor gets the actor property value. 
func (m *CloudPcAuditEvent) GetActor()(*CloudPcAuditActor) {
    if m == nil {
        return nil
    } else {
        return m.actor
    }
}
// GetCategory gets the category property value. Audit category. Read-only.
func (m *CloudPcAuditEvent) GetCategory()(*CloudPcAuditCategory) {
    if m == nil {
        return nil
    } else {
        return m.category
    }
}
// GetComponentName gets the componentName property value. Component name. Read-only.
func (m *CloudPcAuditEvent) GetComponentName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.componentName
    }
}
// GetCorrelationId gets the correlationId property value. The client request identifier, used to correlate activity within the system. Read-only.
func (m *CloudPcAuditEvent) GetCorrelationId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.correlationId
    }
}
// GetDisplayName gets the displayName property value. Event display name. Read-only.
func (m *CloudPcAuditEvent) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetResources gets the resources property value. List of cloudPcAuditResource objects. Read-only.
func (m *CloudPcAuditEvent) GetResources()([]CloudPcAuditResource) {
    if m == nil {
        return nil
    } else {
        return m.resources
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
            m.SetActivityOperationType(val.(*CloudPcAuditActivityOperationType))
        }
        return nil
    }
    res["activityResult"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcAuditActivityResult)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivityResult(val.(*CloudPcAuditActivityResult))
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
            m.SetCategory(val.(*CloudPcAuditCategory))
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
// Serialize serializes information the current object
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
        cast := (*m.GetActivityOperationType()).String()
        err = writer.WriteStringValue("activityOperationType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetActivityResult() != nil {
        cast := (*m.GetActivityResult()).String()
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
        cast := (*m.GetCategory()).String()
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
    if m.GetResources() != nil {
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
// SetActivity sets the activity property value. Friendly name of the activity. Optional.
func (m *CloudPcAuditEvent) SetActivity(value *string)() {
    if m != nil {
        m.activity = value
    }
}
// SetActivityDateTime sets the activityDateTime property value. The date time in UTC when the activity was performed. Read-only.
func (m *CloudPcAuditEvent) SetActivityDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.activityDateTime = value
    }
}
// SetActivityOperationType sets the activityOperationType property value. The HTTP operation type of the activity. Possible values include create, delete, patch and other. Read-only.
func (m *CloudPcAuditEvent) SetActivityOperationType(value *CloudPcAuditActivityOperationType)() {
    if m != nil {
        m.activityOperationType = value
    }
}
// SetActivityResult sets the activityResult property value. The result of the activity. Read-only.
func (m *CloudPcAuditEvent) SetActivityResult(value *CloudPcAuditActivityResult)() {
    if m != nil {
        m.activityResult = value
    }
}
// SetActivityType sets the activityType property value. The type of activity that was performed. Read-only.
func (m *CloudPcAuditEvent) SetActivityType(value *string)() {
    if m != nil {
        m.activityType = value
    }
}
// SetActor sets the actor property value. 
func (m *CloudPcAuditEvent) SetActor(value *CloudPcAuditActor)() {
    if m != nil {
        m.actor = value
    }
}
// SetCategory sets the category property value. Audit category. Read-only.
func (m *CloudPcAuditEvent) SetCategory(value *CloudPcAuditCategory)() {
    if m != nil {
        m.category = value
    }
}
// SetComponentName sets the componentName property value. Component name. Read-only.
func (m *CloudPcAuditEvent) SetComponentName(value *string)() {
    if m != nil {
        m.componentName = value
    }
}
// SetCorrelationId sets the correlationId property value. The client request identifier, used to correlate activity within the system. Read-only.
func (m *CloudPcAuditEvent) SetCorrelationId(value *string)() {
    if m != nil {
        m.correlationId = value
    }
}
// SetDisplayName sets the displayName property value. Event display name. Read-only.
func (m *CloudPcAuditEvent) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetResources sets the resources property value. List of cloudPcAuditResource objects. Read-only.
func (m *CloudPcAuditEvent) SetResources(value []CloudPcAuditResource)() {
    if m != nil {
        m.resources = value
    }
}
