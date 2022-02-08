package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// GroupPolicyDefinitionValue 
type GroupPolicyDefinitionValue struct {
    Entity
    // Specifies how the value should be configured. This can be either as a Policy or as a Preference. Possible values are: policy, preference.
    configurationType *GroupPolicyConfigurationType;
    // The date and time the object was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The associated group policy definition with the value.
    definition *GroupPolicyDefinition;
    // Enables or disables the associated group policy definition.
    enabled *bool;
    // The date and time the entity was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The associated group policy presentation values with the definition value.
    presentationValues []GroupPolicyPresentationValue;
}
// NewGroupPolicyDefinitionValue instantiates a new groupPolicyDefinitionValue and sets the default values.
func NewGroupPolicyDefinitionValue()(*GroupPolicyDefinitionValue) {
    m := &GroupPolicyDefinitionValue{
        Entity: *NewEntity(),
    }
    return m
}
// GetConfigurationType gets the configurationType property value. Specifies how the value should be configured. This can be either as a Policy or as a Preference. Possible values are: policy, preference.
func (m *GroupPolicyDefinitionValue) GetConfigurationType()(*GroupPolicyConfigurationType) {
    if m == nil {
        return nil
    } else {
        return m.configurationType
    }
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time the object was created.
func (m *GroupPolicyDefinitionValue) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetDefinition gets the definition property value. The associated group policy definition with the value.
func (m *GroupPolicyDefinitionValue) GetDefinition()(*GroupPolicyDefinition) {
    if m == nil {
        return nil
    } else {
        return m.definition
    }
}
// GetEnabled gets the enabled property value. Enables or disables the associated group policy definition.
func (m *GroupPolicyDefinitionValue) GetEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enabled
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time the entity was last modified.
func (m *GroupPolicyDefinitionValue) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetPresentationValues gets the presentationValues property value. The associated group policy presentation values with the definition value.
func (m *GroupPolicyDefinitionValue) GetPresentationValues()([]GroupPolicyPresentationValue) {
    if m == nil {
        return nil
    } else {
        return m.presentationValues
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupPolicyDefinitionValue) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["configurationType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseGroupPolicyConfigurationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigurationType(val.(*GroupPolicyConfigurationType))
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
    res["definition"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGroupPolicyDefinition() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefinition(val.(*GroupPolicyDefinition))
        }
        return nil
    }
    res["enabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnabled(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["presentationValues"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGroupPolicyPresentationValue() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GroupPolicyPresentationValue, len(val))
            for i, v := range val {
                res[i] = *(v.(*GroupPolicyPresentationValue))
            }
            m.SetPresentationValues(res)
        }
        return nil
    }
    return res
}
func (m *GroupPolicyDefinitionValue) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *GroupPolicyDefinitionValue) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetConfigurationType() != nil {
        cast := (*m.GetConfigurationType()).String()
        err = writer.WriteStringValue("configurationType", &cast)
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
        err = writer.WriteObjectValue("definition", m.GetDefinition())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("enabled", m.GetEnabled())
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
    if m.GetPresentationValues() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPresentationValues()))
        for i, v := range m.GetPresentationValues() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("presentationValues", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetConfigurationType sets the configurationType property value. Specifies how the value should be configured. This can be either as a Policy or as a Preference. Possible values are: policy, preference.
func (m *GroupPolicyDefinitionValue) SetConfigurationType(value *GroupPolicyConfigurationType)() {
    if m != nil {
        m.configurationType = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time the object was created.
func (m *GroupPolicyDefinitionValue) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetDefinition sets the definition property value. The associated group policy definition with the value.
func (m *GroupPolicyDefinitionValue) SetDefinition(value *GroupPolicyDefinition)() {
    if m != nil {
        m.definition = value
    }
}
// SetEnabled sets the enabled property value. Enables or disables the associated group policy definition.
func (m *GroupPolicyDefinitionValue) SetEnabled(value *bool)() {
    if m != nil {
        m.enabled = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time the entity was last modified.
func (m *GroupPolicyDefinitionValue) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetPresentationValues sets the presentationValues property value. The associated group policy presentation values with the definition value.
func (m *GroupPolicyDefinitionValue) SetPresentationValues(value []GroupPolicyPresentationValue)() {
    if m != nil {
        m.presentationValues = value
    }
}
