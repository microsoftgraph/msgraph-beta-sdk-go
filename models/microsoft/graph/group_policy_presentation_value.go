package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// GroupPolicyPresentationValue 
type GroupPolicyPresentationValue struct {
    Entity
    // The date and time the object was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The group policy definition value associated with the presentation value.
    definitionValue *GroupPolicyDefinitionValue;
    // The date and time the object was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The group policy presentation associated with the presentation value.
    presentation *GroupPolicyPresentation;
}
// NewGroupPolicyPresentationValue instantiates a new groupPolicyPresentationValue and sets the default values.
func NewGroupPolicyPresentationValue()(*GroupPolicyPresentationValue) {
    m := &GroupPolicyPresentationValue{
        Entity: *NewEntity(),
    }
    return m
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time the object was created.
func (m *GroupPolicyPresentationValue) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetDefinitionValue gets the definitionValue property value. The group policy definition value associated with the presentation value.
func (m *GroupPolicyPresentationValue) GetDefinitionValue()(*GroupPolicyDefinitionValue) {
    if m == nil {
        return nil
    } else {
        return m.definitionValue
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time the object was last modified.
func (m *GroupPolicyPresentationValue) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetPresentation gets the presentation property value. The group policy presentation associated with the presentation value.
func (m *GroupPolicyPresentationValue) GetPresentation()(*GroupPolicyPresentation) {
    if m == nil {
        return nil
    } else {
        return m.presentation
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupPolicyPresentationValue) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["definitionValue"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGroupPolicyDefinitionValue() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefinitionValue(val.(*GroupPolicyDefinitionValue))
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
    res["presentation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGroupPolicyPresentation() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPresentation(val.(*GroupPolicyPresentation))
        }
        return nil
    }
    return res
}
func (m *GroupPolicyPresentationValue) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *GroupPolicyPresentationValue) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("definitionValue", m.GetDefinitionValue())
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
        err = writer.WriteObjectValue("presentation", m.GetPresentation())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time the object was created.
func (m *GroupPolicyPresentationValue) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetDefinitionValue sets the definitionValue property value. The group policy definition value associated with the presentation value.
func (m *GroupPolicyPresentationValue) SetDefinitionValue(value *GroupPolicyDefinitionValue)() {
    if m != nil {
        m.definitionValue = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time the object was last modified.
func (m *GroupPolicyPresentationValue) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetPresentation sets the presentation property value. The group policy presentation associated with the presentation value.
func (m *GroupPolicyPresentationValue) SetPresentation(value *GroupPolicyPresentation)() {
    if m != nil {
        m.presentation = value
    }
}
