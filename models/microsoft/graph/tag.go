package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/ediscovery"
)

// Tag 
type Tag struct {
    Entity
    // Indicates whether a single or multiple child tags can be associated with a document. Possible values are: One, Many.  This value controls whether the UX presents the tags as checkboxes or a radio button group.
    childSelectability *i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3.ChildSelectability;
    // Returns the tags that are a child of a tag.
    childTags []Tag;
    // The user who created the tag.
    createdBy *IdentitySet;
    // The description for the tag.
    description *string;
    // Display name of the tag.
    displayName *string;
    // The date and time the tag was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Returns the parent tag of the specified tag.
    parent *Tag;
}
// NewTag instantiates a new tag and sets the default values.
func NewTag()(*Tag) {
    m := &Tag{
        Entity: *NewEntity(),
    }
    return m
}
// GetChildSelectability gets the childSelectability property value. Indicates whether a single or multiple child tags can be associated with a document. Possible values are: One, Many.  This value controls whether the UX presents the tags as checkboxes or a radio button group.
func (m *Tag) GetChildSelectability()(*i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3.ChildSelectability) {
    if m == nil {
        return nil
    } else {
        return m.childSelectability
    }
}
// GetChildTags gets the childTags property value. Returns the tags that are a child of a tag.
func (m *Tag) GetChildTags()([]Tag) {
    if m == nil {
        return nil
    } else {
        return m.childTags
    }
}
// GetCreatedBy gets the createdBy property value. The user who created the tag.
func (m *Tag) GetCreatedBy()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.createdBy
    }
}
// GetDescription gets the description property value. The description for the tag.
func (m *Tag) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. Display name of the tag.
func (m *Tag) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time the tag was last modified.
func (m *Tag) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetParent gets the parent property value. Returns the parent tag of the specified tag.
func (m *Tag) GetParent()(*Tag) {
    if m == nil {
        return nil
    } else {
        return m.parent
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Tag) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["childSelectability"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3.ParseChildSelectability)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChildSelectability(val.(*i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3.ChildSelectability))
        }
        return nil
    }
    res["childTags"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTag() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Tag, len(val))
            for i, v := range val {
                res[i] = *(v.(*Tag))
            }
            m.SetChildTags(res)
        }
        return nil
    }
    res["createdBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(*IdentitySet))
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
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTag() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParent(val.(*Tag))
        }
        return nil
    }
    return res
}
func (m *Tag) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Tag) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetChildSelectability() != nil {
        cast := (*m.GetChildSelectability()).String()
        err = writer.WriteStringValue("childSelectability", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetChildTags() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetChildTags()))
        for i, v := range m.GetChildTags() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("childTags", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
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
        err = writer.WriteObjectValue("parent", m.GetParent())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetChildSelectability sets the childSelectability property value. Indicates whether a single or multiple child tags can be associated with a document. Possible values are: One, Many.  This value controls whether the UX presents the tags as checkboxes or a radio button group.
func (m *Tag) SetChildSelectability(value *i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3.ChildSelectability)() {
    if m != nil {
        m.childSelectability = value
    }
}
// SetChildTags sets the childTags property value. Returns the tags that are a child of a tag.
func (m *Tag) SetChildTags(value []Tag)() {
    if m != nil {
        m.childTags = value
    }
}
// SetCreatedBy sets the createdBy property value. The user who created the tag.
func (m *Tag) SetCreatedBy(value *IdentitySet)() {
    if m != nil {
        m.createdBy = value
    }
}
// SetDescription sets the description property value. The description for the tag.
func (m *Tag) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. Display name of the tag.
func (m *Tag) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time the tag was last modified.
func (m *Tag) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetParent sets the parent property value. Returns the parent tag of the specified tag.
func (m *Tag) SetParent(value *Tag)() {
    if m != nil {
        m.parent = value
    }
}
