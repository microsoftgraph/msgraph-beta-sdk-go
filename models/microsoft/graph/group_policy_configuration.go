package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type GroupPolicyConfiguration struct {
    Entity
    // The list of group assignments for the configuration.
    assignments []GroupPolicyConfigurationAssignment;
    // The date and time the object was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The list of enabled or disabled group policy definition values for the configuration.
    definitionValues []GroupPolicyDefinitionValue;
    // User provided description for the resource object.
    description *string;
    // User provided name for the resource object.
    displayName *string;
    // The date and time the entity was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The list of scope tags for the configuration.
    roleScopeTagIds []string;
}
// Instantiates a new groupPolicyConfiguration and sets the default values.
func NewGroupPolicyConfiguration()(*GroupPolicyConfiguration) {
    m := &GroupPolicyConfiguration{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the assignments property value. The list of group assignments for the configuration.
func (m *GroupPolicyConfiguration) GetAssignments()([]GroupPolicyConfigurationAssignment) {
    if m == nil {
        return nil
    } else {
        return m.assignments
    }
}
// Gets the createdDateTime property value. The date and time the object was created.
func (m *GroupPolicyConfiguration) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the definitionValues property value. The list of enabled or disabled group policy definition values for the configuration.
func (m *GroupPolicyConfiguration) GetDefinitionValues()([]GroupPolicyDefinitionValue) {
    if m == nil {
        return nil
    } else {
        return m.definitionValues
    }
}
// Gets the description property value. User provided description for the resource object.
func (m *GroupPolicyConfiguration) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. User provided name for the resource object.
func (m *GroupPolicyConfiguration) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the lastModifiedDateTime property value. The date and time the entity was last modified.
func (m *GroupPolicyConfiguration) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the roleScopeTagIds property value. The list of scope tags for the configuration.
func (m *GroupPolicyConfiguration) GetRoleScopeTagIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.roleScopeTagIds
    }
}
// The deserialization information for the current model
func (m *GroupPolicyConfiguration) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGroupPolicyConfigurationAssignment() })
        if err != nil {
            return err
        }
        res := make([]GroupPolicyConfigurationAssignment, len(val))
        for i, v := range val {
            res[i] = *(v.(*GroupPolicyConfigurationAssignment))
        }
        m.SetAssignments(res)
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
    res["definitionValues"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGroupPolicyDefinitionValue() })
        if err != nil {
            return err
        }
        res := make([]GroupPolicyDefinitionValue, len(val))
        for i, v := range val {
            res[i] = *(v.(*GroupPolicyDefinitionValue))
        }
        m.SetDefinitionValues(res)
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
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastModifiedDateTime(val)
        return nil
    }
    res["roleScopeTagIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetRoleScopeTagIds(res)
        return nil
    }
    return res
}
func (m *GroupPolicyConfiguration) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *GroupPolicyConfiguration) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAssignments()))
        for i, v := range m.GetAssignments() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("assignments", cast)
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDefinitionValues()))
        for i, v := range m.GetDefinitionValues() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("definitionValues", cast)
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
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("roleScopeTagIds", m.GetRoleScopeTagIds())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the assignments property value. The list of group assignments for the configuration.
// Parameters:
//  - value : Value to set for the assignments property.
func (m *GroupPolicyConfiguration) SetAssignments(value []GroupPolicyConfigurationAssignment)() {
    m.assignments = value
}
// Sets the createdDateTime property value. The date and time the object was created.
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *GroupPolicyConfiguration) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the definitionValues property value. The list of enabled or disabled group policy definition values for the configuration.
// Parameters:
//  - value : Value to set for the definitionValues property.
func (m *GroupPolicyConfiguration) SetDefinitionValues(value []GroupPolicyDefinitionValue)() {
    m.definitionValues = value
}
// Sets the description property value. User provided description for the resource object.
// Parameters:
//  - value : Value to set for the description property.
func (m *GroupPolicyConfiguration) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. User provided name for the resource object.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *GroupPolicyConfiguration) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the lastModifiedDateTime property value. The date and time the entity was last modified.
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *GroupPolicyConfiguration) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the roleScopeTagIds property value. The list of scope tags for the configuration.
// Parameters:
//  - value : Value to set for the roleScopeTagIds property.
func (m *GroupPolicyConfiguration) SetRoleScopeTagIds(value []string)() {
    m.roleScopeTagIds = value
}
