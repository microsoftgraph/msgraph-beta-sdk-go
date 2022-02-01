package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ManagementConditionStatement 
type ManagementConditionStatement struct {
    Entity
    // The applicable platforms for this management condition statement.
    applicablePlatforms []DevicePlatformType;
    // The time the management condition statement was created. Generated service side.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The admin defined description of the management condition statement.
    description *string;
    // The admin defined name of the management condition statement.
    displayName *string;
    // ETag of the management condition statement. Updated service side.
    eTag *string;
    // The management condition statement expression used to evaluate if a management condition statement was activated/deactivated.
    expression *ManagementConditionExpression;
    // The management conditions associated to the management condition statement.
    managementConditions []ManagementCondition;
    // The time the management condition statement was last modified. Updated service side.
    modifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
}
// NewManagementConditionStatement instantiates a new managementConditionStatement and sets the default values.
func NewManagementConditionStatement()(*ManagementConditionStatement) {
    m := &ManagementConditionStatement{
        Entity: *NewEntity(),
    }
    return m
}
// GetApplicablePlatforms gets the applicablePlatforms property value. The applicable platforms for this management condition statement.
func (m *ManagementConditionStatement) GetApplicablePlatforms()([]DevicePlatformType) {
    if m == nil {
        return nil
    } else {
        return m.applicablePlatforms
    }
}
// GetCreatedDateTime gets the createdDateTime property value. The time the management condition statement was created. Generated service side.
func (m *ManagementConditionStatement) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetDescription gets the description property value. The admin defined description of the management condition statement.
func (m *ManagementConditionStatement) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. The admin defined name of the management condition statement.
func (m *ManagementConditionStatement) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetETag gets the eTag property value. ETag of the management condition statement. Updated service side.
func (m *ManagementConditionStatement) GetETag()(*string) {
    if m == nil {
        return nil
    } else {
        return m.eTag
    }
}
// GetExpression gets the expression property value. The management condition statement expression used to evaluate if a management condition statement was activated/deactivated.
func (m *ManagementConditionStatement) GetExpression()(*ManagementConditionExpression) {
    if m == nil {
        return nil
    } else {
        return m.expression
    }
}
// GetManagementConditions gets the managementConditions property value. The management conditions associated to the management condition statement.
func (m *ManagementConditionStatement) GetManagementConditions()([]ManagementCondition) {
    if m == nil {
        return nil
    } else {
        return m.managementConditions
    }
}
// GetModifiedDateTime gets the modifiedDateTime property value. The time the management condition statement was last modified. Updated service side.
func (m *ManagementConditionStatement) GetModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.modifiedDateTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagementConditionStatement) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["applicablePlatforms"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseDevicePlatformType)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DevicePlatformType, len(val))
            for i, v := range val {
                res[i] = *(v.(*DevicePlatformType))
            }
            m.SetApplicablePlatforms(res)
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
    res["eTag"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetETag(val)
        }
        return nil
    }
    res["expression"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagementConditionExpression() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpression(val.(*ManagementConditionExpression))
        }
        return nil
    }
    res["managementConditions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagementCondition() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagementCondition, len(val))
            for i, v := range val {
                res[i] = *(v.(*ManagementCondition))
            }
            m.SetManagementConditions(res)
        }
        return nil
    }
    res["modifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModifiedDateTime(val)
        }
        return nil
    }
    return res
}
func (m *ManagementConditionStatement) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ManagementConditionStatement) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetApplicablePlatforms() != nil {
        err = writer.WriteCollectionOfStringValues("applicablePlatforms", SerializeDevicePlatformType(m.GetApplicablePlatforms()))
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
        err = writer.WriteStringValue("eTag", m.GetETag())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("expression", m.GetExpression())
        if err != nil {
            return err
        }
    }
    if m.GetManagementConditions() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetManagementConditions()))
        for i, v := range m.GetManagementConditions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("managementConditions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("modifiedDateTime", m.GetModifiedDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApplicablePlatforms sets the applicablePlatforms property value. The applicable platforms for this management condition statement.
func (m *ManagementConditionStatement) SetApplicablePlatforms(value []DevicePlatformType)() {
    if m != nil {
        m.applicablePlatforms = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The time the management condition statement was created. Generated service side.
func (m *ManagementConditionStatement) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetDescription sets the description property value. The admin defined description of the management condition statement.
func (m *ManagementConditionStatement) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. The admin defined name of the management condition statement.
func (m *ManagementConditionStatement) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetETag sets the eTag property value. ETag of the management condition statement. Updated service side.
func (m *ManagementConditionStatement) SetETag(value *string)() {
    if m != nil {
        m.eTag = value
    }
}
// SetExpression sets the expression property value. The management condition statement expression used to evaluate if a management condition statement was activated/deactivated.
func (m *ManagementConditionStatement) SetExpression(value *ManagementConditionExpression)() {
    if m != nil {
        m.expression = value
    }
}
// SetManagementConditions sets the managementConditions property value. The management conditions associated to the management condition statement.
func (m *ManagementConditionStatement) SetManagementConditions(value []ManagementCondition)() {
    if m != nil {
        m.managementConditions = value
    }
}
// SetModifiedDateTime sets the modifiedDateTime property value. The time the management condition statement was last modified. Updated service side.
func (m *ManagementConditionStatement) SetModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.modifiedDateTime = value
    }
}
