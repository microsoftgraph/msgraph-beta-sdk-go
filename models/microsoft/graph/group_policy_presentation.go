package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// GroupPolicyPresentation 
type GroupPolicyPresentation struct {
    Entity
    // The group policy definition associated with the presentation.
    definition *GroupPolicyDefinition;
    // Localized text label for any presentation entity. The default value is empty.
    label *string;
    // The date and time the entity was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
}
// NewGroupPolicyPresentation instantiates a new groupPolicyPresentation and sets the default values.
func NewGroupPolicyPresentation()(*GroupPolicyPresentation) {
    m := &GroupPolicyPresentation{
        Entity: *NewEntity(),
    }
    return m
}
// GetDefinition gets the definition property value. The group policy definition associated with the presentation.
func (m *GroupPolicyPresentation) GetDefinition()(*GroupPolicyDefinition) {
    if m == nil {
        return nil
    } else {
        return m.definition
    }
}
// GetLabel gets the label property value. Localized text label for any presentation entity. The default value is empty.
func (m *GroupPolicyPresentation) GetLabel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.label
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time the entity was last modified.
func (m *GroupPolicyPresentation) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupPolicyPresentation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["label"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLabel(val)
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
    return res
}
func (m *GroupPolicyPresentation) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *GroupPolicyPresentation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("definition", m.GetDefinition())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("label", m.GetLabel())
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
    return nil
}
// SetDefinition sets the definition property value. The group policy definition associated with the presentation.
func (m *GroupPolicyPresentation) SetDefinition(value *GroupPolicyDefinition)() {
    if m != nil {
        m.definition = value
    }
}
// SetLabel sets the label property value. Localized text label for any presentation entity. The default value is empty.
func (m *GroupPolicyPresentation) SetLabel(value *string)() {
    if m != nil {
        m.label = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time the entity was last modified.
func (m *GroupPolicyPresentation) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
