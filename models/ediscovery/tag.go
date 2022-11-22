package ediscovery

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// Tag provides operations to manage the collection of accessReviewDecision entities.
type Tag struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // Indicates whether a single or multiple child tags can be associated with a document. Possible values are: One, Many.  This value controls whether the UX presents the tags as checkboxes or a radio button group.
    childSelectability *ChildSelectability
    // Returns the tags that are a child of a tag.
    childTags []Tagable
    // The user who created the tag.
    createdBy ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable
    // The description for the tag.
    description *string
    // Display name of the tag.
    displayName *string
    // The date and time the tag was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Returns the parent tag of the specified tag.
    parent Tagable
}
// NewTag instantiates a new tag and sets the default values.
func NewTag()(*Tag) {
    m := &Tag{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateTagFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTagFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTag(), nil
}
// GetChildSelectability gets the childSelectability property value. Indicates whether a single or multiple child tags can be associated with a document. Possible values are: One, Many.  This value controls whether the UX presents the tags as checkboxes or a radio button group.
func (m *Tag) GetChildSelectability()(*ChildSelectability) {
    return m.childSelectability
}
// GetChildTags gets the childTags property value. Returns the tags that are a child of a tag.
func (m *Tag) GetChildTags()([]Tagable) {
    return m.childTags
}
// GetCreatedBy gets the createdBy property value. The user who created the tag.
func (m *Tag) GetCreatedBy()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable) {
    return m.createdBy
}
// GetDescription gets the description property value. The description for the tag.
func (m *Tag) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. Display name of the tag.
func (m *Tag) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Tag) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["childSelectability"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseChildSelectability , m.SetChildSelectability)
    res["childTags"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateTagFromDiscriminatorValue , m.SetChildTags)
    res["createdBy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateIdentitySetFromDiscriminatorValue , m.SetCreatedBy)
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["lastModifiedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastModifiedDateTime)
    res["parent"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateTagFromDiscriminatorValue , m.SetParent)
    return res
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time the tag was last modified.
func (m *Tag) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetParent gets the parent property value. Returns the parent tag of the specified tag.
func (m *Tag) GetParent()(Tagable) {
    return m.parent
}
// Serialize serializes information the current object
func (m *Tag) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetChildTags())
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
func (m *Tag) SetChildSelectability(value *ChildSelectability)() {
    m.childSelectability = value
}
// SetChildTags sets the childTags property value. Returns the tags that are a child of a tag.
func (m *Tag) SetChildTags(value []Tagable)() {
    m.childTags = value
}
// SetCreatedBy sets the createdBy property value. The user who created the tag.
func (m *Tag) SetCreatedBy(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable)() {
    m.createdBy = value
}
// SetDescription sets the description property value. The description for the tag.
func (m *Tag) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. Display name of the tag.
func (m *Tag) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time the tag was last modified.
func (m *Tag) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetParent sets the parent property value. Returns the parent tag of the specified tag.
func (m *Tag) SetParent(value Tagable)() {
    m.parent = value
}
