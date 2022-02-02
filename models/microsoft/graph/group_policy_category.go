package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// GroupPolicyCategory 
type GroupPolicyCategory struct {
    Entity
    // The children categories
    children []GroupPolicyCategory;
    // The id of the definition file the category came from
    definitionFile *GroupPolicyDefinitionFile;
    // The immediate GroupPolicyDefinition children of the category
    definitions []GroupPolicyDefinition;
    // The string id of the category's display name
    displayName *string;
    // Defines if the category is a root category
    isRoot *bool;
    // The date and time the entity was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The parent category
    parent *GroupPolicyCategory;
}
// NewGroupPolicyCategory instantiates a new groupPolicyCategory and sets the default values.
func NewGroupPolicyCategory()(*GroupPolicyCategory) {
    m := &GroupPolicyCategory{
        Entity: *NewEntity(),
    }
    return m
}
// GetChildren gets the children property value. The children categories
func (m *GroupPolicyCategory) GetChildren()([]GroupPolicyCategory) {
    if m == nil {
        return nil
    } else {
        return m.children
    }
}
// GetDefinitionFile gets the definitionFile property value. The id of the definition file the category came from
func (m *GroupPolicyCategory) GetDefinitionFile()(*GroupPolicyDefinitionFile) {
    if m == nil {
        return nil
    } else {
        return m.definitionFile
    }
}
// GetDefinitions gets the definitions property value. The immediate GroupPolicyDefinition children of the category
func (m *GroupPolicyCategory) GetDefinitions()([]GroupPolicyDefinition) {
    if m == nil {
        return nil
    } else {
        return m.definitions
    }
}
// GetDisplayName gets the displayName property value. The string id of the category's display name
func (m *GroupPolicyCategory) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetIsRoot gets the isRoot property value. Defines if the category is a root category
func (m *GroupPolicyCategory) GetIsRoot()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isRoot
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time the entity was last modified.
func (m *GroupPolicyCategory) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetParent gets the parent property value. The parent category
func (m *GroupPolicyCategory) GetParent()(*GroupPolicyCategory) {
    if m == nil {
        return nil
    } else {
        return m.parent
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupPolicyCategory) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["children"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGroupPolicyCategory() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GroupPolicyCategory, len(val))
            for i, v := range val {
                res[i] = *(v.(*GroupPolicyCategory))
            }
            m.SetChildren(res)
        }
        return nil
    }
    res["definitionFile"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGroupPolicyDefinitionFile() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefinitionFile(val.(*GroupPolicyDefinitionFile))
        }
        return nil
    }
    res["definitions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGroupPolicyDefinition() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GroupPolicyDefinition, len(val))
            for i, v := range val {
                res[i] = *(v.(*GroupPolicyDefinition))
            }
            m.SetDefinitions(res)
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
    res["isRoot"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsRoot(val)
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
    res["parent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGroupPolicyCategory() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParent(val.(*GroupPolicyCategory))
        }
        return nil
    }
    return res
}
func (m *GroupPolicyCategory) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *GroupPolicyCategory) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetChildren() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetChildren()))
        for i, v := range m.GetChildren() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("children", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("definitionFile", m.GetDefinitionFile())
        if err != nil {
            return err
        }
    }
    if m.GetDefinitions() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDefinitions()))
        for i, v := range m.GetDefinitions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("definitions", cast)
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
        err = writer.WriteBoolValue("isRoot", m.GetIsRoot())
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
        err = writer.WriteObjectValue("parent", m.GetParent())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetChildren sets the children property value. The children categories
func (m *GroupPolicyCategory) SetChildren(value []GroupPolicyCategory)() {
    if m != nil {
        m.children = value
    }
}
// SetDefinitionFile sets the definitionFile property value. The id of the definition file the category came from
func (m *GroupPolicyCategory) SetDefinitionFile(value *GroupPolicyDefinitionFile)() {
    if m != nil {
        m.definitionFile = value
    }
}
// SetDefinitions sets the definitions property value. The immediate GroupPolicyDefinition children of the category
func (m *GroupPolicyCategory) SetDefinitions(value []GroupPolicyDefinition)() {
    if m != nil {
        m.definitions = value
    }
}
// SetDisplayName sets the displayName property value. The string id of the category's display name
func (m *GroupPolicyCategory) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetIsRoot sets the isRoot property value. Defines if the category is a root category
func (m *GroupPolicyCategory) SetIsRoot(value *bool)() {
    if m != nil {
        m.isRoot = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time the entity was last modified.
func (m *GroupPolicyCategory) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetParent sets the parent property value. The parent category
func (m *GroupPolicyCategory) SetParent(value *GroupPolicyCategory)() {
    if m != nil {
        m.parent = value
    }
}
