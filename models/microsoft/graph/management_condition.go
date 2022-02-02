package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ManagementCondition 
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
// NewManagementCondition instantiates a new managementCondition and sets the default values.
func NewManagementCondition()(*ManagementCondition) {
    m := &ManagementCondition{
        Entity: *NewEntity(),
    }
    return m
}
// GetApplicablePlatforms gets the applicablePlatforms property value. The applicable platforms for this management condition.
func (m *ManagementCondition) GetApplicablePlatforms()([]DevicePlatformType) {
    if m == nil {
        return nil
    } else {
        return m.applicablePlatforms
    }
}
// GetCreatedDateTime gets the createdDateTime property value. The time the management condition was created. Generated service side.
func (m *ManagementCondition) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetDescription gets the description property value. The admin defined description of the management condition.
func (m *ManagementCondition) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. The admin defined name of the management condition.
func (m *ManagementCondition) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetETag gets the eTag property value. ETag of the management condition. Updated service side.
func (m *ManagementCondition) GetETag()(*string) {
    if m == nil {
        return nil
    } else {
        return m.eTag
    }
}
// GetManagementConditionStatements gets the managementConditionStatements property value. The management condition statements associated to the management condition.
func (m *ManagementCondition) GetManagementConditionStatements()([]ManagementConditionStatement) {
    if m == nil {
        return nil
    } else {
        return m.managementConditionStatements
    }
}
// GetModifiedDateTime gets the modifiedDateTime property value. The time the management condition was last modified. Updated service side.
func (m *ManagementCondition) GetModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.modifiedDateTime
    }
}
// GetUniqueName gets the uniqueName property value. Unique name for the management condition. Used in management condition expressions.
func (m *ManagementCondition) GetUniqueName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.uniqueName
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
// Serialize serializes information the current object
func (m *ManagementCondition) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
    if m.GetManagementConditionStatements() != nil {
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
// SetApplicablePlatforms sets the applicablePlatforms property value. The applicable platforms for this management condition.
func (m *ManagementCondition) SetApplicablePlatforms(value []DevicePlatformType)() {
    if m != nil {
        m.applicablePlatforms = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The time the management condition was created. Generated service side.
func (m *ManagementCondition) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetDescription sets the description property value. The admin defined description of the management condition.
func (m *ManagementCondition) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. The admin defined name of the management condition.
func (m *ManagementCondition) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetETag sets the eTag property value. ETag of the management condition. Updated service side.
func (m *ManagementCondition) SetETag(value *string)() {
    if m != nil {
        m.eTag = value
    }
}
// SetManagementConditionStatements sets the managementConditionStatements property value. The management condition statements associated to the management condition.
func (m *ManagementCondition) SetManagementConditionStatements(value []ManagementConditionStatement)() {
    if m != nil {
        m.managementConditionStatements = value
    }
}
// SetModifiedDateTime sets the modifiedDateTime property value. The time the management condition was last modified. Updated service side.
func (m *ManagementCondition) SetModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.modifiedDateTime = value
    }
}
// SetUniqueName sets the uniqueName property value. Unique name for the management condition. Used in management condition expressions.
func (m *ManagementCondition) SetUniqueName(value *string)() {
    if m != nil {
        m.uniqueName = value
    }
}
