package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new managementConditionStatement and sets the default values.
func NewManagementConditionStatement()(*ManagementConditionStatement) {
    m := &ManagementConditionStatement{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the applicablePlatforms property value. The applicable platforms for this management condition statement.
func (m *ManagementConditionStatement) GetApplicablePlatforms()([]DevicePlatformType) {
    if m == nil {
        return nil
    } else {
        return m.applicablePlatforms
    }
}
// Gets the createdDateTime property value. The time the management condition statement was created. Generated service side.
func (m *ManagementConditionStatement) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the description property value. The admin defined description of the management condition statement.
func (m *ManagementConditionStatement) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. The admin defined name of the management condition statement.
func (m *ManagementConditionStatement) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the eTag property value. ETag of the management condition statement. Updated service side.
func (m *ManagementConditionStatement) GetETag()(*string) {
    if m == nil {
        return nil
    } else {
        return m.eTag
    }
}
// Gets the expression property value. The management condition statement expression used to evaluate if a management condition statement was activated/deactivated.
func (m *ManagementConditionStatement) GetExpression()(*ManagementConditionExpression) {
    if m == nil {
        return nil
    } else {
        return m.expression
    }
}
// Gets the managementConditions property value. The management conditions associated to the management condition statement.
func (m *ManagementConditionStatement) GetManagementConditions()([]ManagementCondition) {
    if m == nil {
        return nil
    } else {
        return m.managementConditions
    }
}
// Gets the modifiedDateTime property value. The time the management condition statement was last modified. Updated service side.
func (m *ManagementConditionStatement) GetModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.modifiedDateTime
    }
}
// The deserialization information for the current model
func (m *ManagementConditionStatement) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["applicablePlatforms"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseDevicePlatformType)
        if err != nil {
            return err
        }
        res := make([]DevicePlatformType, len(val))
        for i, v := range val {
            res[i] = *(v.(*DevicePlatformType))
        }
        m.SetApplicablePlatforms(res)
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
    res["eTag"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetETag(val)
        return nil
    }
    res["expression"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagementConditionExpression() })
        if err != nil {
            return err
        }
        m.SetExpression(val.(*ManagementConditionExpression))
        return nil
    }
    res["managementConditions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagementCondition() })
        if err != nil {
            return err
        }
        res := make([]ManagementCondition, len(val))
        for i, v := range val {
            res[i] = *(v.(*ManagementCondition))
        }
        m.SetManagementConditions(res)
        return nil
    }
    res["modifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetModifiedDateTime(val)
        return nil
    }
    return res
}
func (m *ManagementConditionStatement) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ManagementConditionStatement) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
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
    {
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
// Sets the applicablePlatforms property value. The applicable platforms for this management condition statement.
// Parameters:
//  - value : Value to set for the applicablePlatforms property.
func (m *ManagementConditionStatement) SetApplicablePlatforms(value []DevicePlatformType)() {
    m.applicablePlatforms = value
}
// Sets the createdDateTime property value. The time the management condition statement was created. Generated service side.
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *ManagementConditionStatement) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the description property value. The admin defined description of the management condition statement.
// Parameters:
//  - value : Value to set for the description property.
func (m *ManagementConditionStatement) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. The admin defined name of the management condition statement.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *ManagementConditionStatement) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the eTag property value. ETag of the management condition statement. Updated service side.
// Parameters:
//  - value : Value to set for the eTag property.
func (m *ManagementConditionStatement) SetETag(value *string)() {
    m.eTag = value
}
// Sets the expression property value. The management condition statement expression used to evaluate if a management condition statement was activated/deactivated.
// Parameters:
//  - value : Value to set for the expression property.
func (m *ManagementConditionStatement) SetExpression(value *ManagementConditionExpression)() {
    m.expression = value
}
// Sets the managementConditions property value. The management conditions associated to the management condition statement.
// Parameters:
//  - value : Value to set for the managementConditions property.
func (m *ManagementConditionStatement) SetManagementConditions(value []ManagementCondition)() {
    m.managementConditions = value
}
// Sets the modifiedDateTime property value. The time the management condition statement was last modified. Updated service side.
// Parameters:
//  - value : Value to set for the modifiedDateTime property.
func (m *ManagementConditionStatement) SetModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.modifiedDateTime = value
}
