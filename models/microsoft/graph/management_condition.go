package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ManagementCondition struct {
    Entity
    // The applicable platforms for this management condition.
    applicablePlatforms []DevicePlatformType;
    // The time the management condition was created. Generated service side.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The admin defined description of the management condition.
    description *string;
    // The admin defined name of the management condition.
    displayName *string;
    // ETag of the management condition. Updated service side.
    eTag *string;
    // The management condition statements associated to the management condition.
    managementConditionStatements []ManagementConditionStatement;
    // The time the management condition was last modified. Updated service side.
    modifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Unique name for the management condition. Used in management condition expressions.
    uniqueName *string;
}
// Instantiates a new managementCondition and sets the default values.
func NewManagementCondition()(*ManagementCondition) {
    m := &ManagementCondition{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the applicablePlatforms property value. The applicable platforms for this management condition.
func (m *ManagementCondition) GetApplicablePlatforms()([]DevicePlatformType) {
    if m == nil {
        return nil
    } else {
        return m.applicablePlatforms
    }
}
// Gets the createdDateTime property value. The time the management condition was created. Generated service side.
func (m *ManagementCondition) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the description property value. The admin defined description of the management condition.
func (m *ManagementCondition) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. The admin defined name of the management condition.
func (m *ManagementCondition) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the eTag property value. ETag of the management condition. Updated service side.
func (m *ManagementCondition) GetETag()(*string) {
    if m == nil {
        return nil
    } else {
        return m.eTag
    }
}
// Gets the managementConditionStatements property value. The management condition statements associated to the management condition.
func (m *ManagementCondition) GetManagementConditionStatements()([]ManagementConditionStatement) {
    if m == nil {
        return nil
    } else {
        return m.managementConditionStatements
    }
}
// Gets the modifiedDateTime property value. The time the management condition was last modified. Updated service side.
func (m *ManagementCondition) GetModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.modifiedDateTime
    }
}
// Gets the uniqueName property value. Unique name for the management condition. Used in management condition expressions.
func (m *ManagementCondition) GetUniqueName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.uniqueName
    }
}
// The deserialization information for the current model
func (m *ManagementCondition) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    res["managementConditionStatements"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagementConditionStatement() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagementConditionStatement, len(val))
            for i, v := range val {
                res[i] = *(v.(*ManagementConditionStatement))
            }
            m.SetManagementConditionStatements(res)
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
    res["uniqueName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUniqueName(val)
        }
        return nil
    }
    return res
}
func (m *ManagementCondition) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ManagementCondition) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetManagementConditionStatements()))
        for i, v := range m.GetManagementConditionStatements() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("managementConditionStatements", cast)
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
    {
        err = writer.WriteStringValue("uniqueName", m.GetUniqueName())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the applicablePlatforms property value. The applicable platforms for this management condition.
// Parameters:
//  - value : Value to set for the applicablePlatforms property.
func (m *ManagementCondition) SetApplicablePlatforms(value []DevicePlatformType)() {
    m.applicablePlatforms = value
}
// Sets the createdDateTime property value. The time the management condition was created. Generated service side.
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *ManagementCondition) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the description property value. The admin defined description of the management condition.
// Parameters:
//  - value : Value to set for the description property.
func (m *ManagementCondition) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. The admin defined name of the management condition.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *ManagementCondition) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the eTag property value. ETag of the management condition. Updated service side.
// Parameters:
//  - value : Value to set for the eTag property.
func (m *ManagementCondition) SetETag(value *string)() {
    m.eTag = value
}
// Sets the managementConditionStatements property value. The management condition statements associated to the management condition.
// Parameters:
//  - value : Value to set for the managementConditionStatements property.
func (m *ManagementCondition) SetManagementConditionStatements(value []ManagementConditionStatement)() {
    m.managementConditionStatements = value
}
// Sets the modifiedDateTime property value. The time the management condition was last modified. Updated service side.
// Parameters:
//  - value : Value to set for the modifiedDateTime property.
func (m *ManagementCondition) SetModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.modifiedDateTime = value
}
// Sets the uniqueName property value. Unique name for the management condition. Used in management condition expressions.
// Parameters:
//  - value : Value to set for the uniqueName property.
func (m *ManagementCondition) SetUniqueName(value *string)() {
    m.uniqueName = value
}
