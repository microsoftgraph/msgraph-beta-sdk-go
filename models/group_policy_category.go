package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupPolicyCategory 
type GroupPolicyCategory struct {
    Entity
    // The children categories
    children []GroupPolicyCategoryable
    // The id of the definition file the category came from
    definitionFile GroupPolicyDefinitionFileable
    // The immediate GroupPolicyDefinition children of the category
    definitions []GroupPolicyDefinitionable
    // The string id of the category's display name
    displayName *string
    // Category Ingestion source
    ingestionSource *IngestionSource
    // Defines if the category is a root category
    isRoot *bool
    // The date and time the entity was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The parent category
    parent GroupPolicyCategoryable
}
// NewGroupPolicyCategory instantiates a new groupPolicyCategory and sets the default values.
func NewGroupPolicyCategory()(*GroupPolicyCategory) {
    m := &GroupPolicyCategory{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.groupPolicyCategory";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateGroupPolicyCategoryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupPolicyCategoryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupPolicyCategory(), nil
}
// GetChildren gets the children property value. The children categories
func (m *GroupPolicyCategory) GetChildren()([]GroupPolicyCategoryable) {
    return m.children
}
// GetDefinitionFile gets the definitionFile property value. The id of the definition file the category came from
func (m *GroupPolicyCategory) GetDefinitionFile()(GroupPolicyDefinitionFileable) {
    return m.definitionFile
}
// GetDefinitions gets the definitions property value. The immediate GroupPolicyDefinition children of the category
func (m *GroupPolicyCategory) GetDefinitions()([]GroupPolicyDefinitionable) {
    return m.definitions
}
// GetDisplayName gets the displayName property value. The string id of the category's display name
func (m *GroupPolicyCategory) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupPolicyCategory) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["children"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateGroupPolicyCategoryFromDiscriminatorValue , m.SetChildren)
    res["definitionFile"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateGroupPolicyDefinitionFileFromDiscriminatorValue , m.SetDefinitionFile)
    res["definitions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateGroupPolicyDefinitionFromDiscriminatorValue , m.SetDefinitions)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["ingestionSource"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseIngestionSource , m.SetIngestionSource)
    res["isRoot"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsRoot)
    res["lastModifiedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastModifiedDateTime)
    res["parent"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateGroupPolicyCategoryFromDiscriminatorValue , m.SetParent)
    return res
}
// GetIngestionSource gets the ingestionSource property value. Category Ingestion source
func (m *GroupPolicyCategory) GetIngestionSource()(*IngestionSource) {
    return m.ingestionSource
}
// GetIsRoot gets the isRoot property value. Defines if the category is a root category
func (m *GroupPolicyCategory) GetIsRoot()(*bool) {
    return m.isRoot
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time the entity was last modified.
func (m *GroupPolicyCategory) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetParent gets the parent property value. The parent category
func (m *GroupPolicyCategory) GetParent()(GroupPolicyCategoryable) {
    return m.parent
}
// Serialize serializes information the current object
func (m *GroupPolicyCategory) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetChildren() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetChildren())
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
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDefinitions())
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
    if m.GetIngestionSource() != nil {
        cast := (*m.GetIngestionSource()).String()
        err = writer.WriteStringValue("ingestionSource", &cast)
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
func (m *GroupPolicyCategory) SetChildren(value []GroupPolicyCategoryable)() {
    m.children = value
}
// SetDefinitionFile sets the definitionFile property value. The id of the definition file the category came from
func (m *GroupPolicyCategory) SetDefinitionFile(value GroupPolicyDefinitionFileable)() {
    m.definitionFile = value
}
// SetDefinitions sets the definitions property value. The immediate GroupPolicyDefinition children of the category
func (m *GroupPolicyCategory) SetDefinitions(value []GroupPolicyDefinitionable)() {
    m.definitions = value
}
// SetDisplayName sets the displayName property value. The string id of the category's display name
func (m *GroupPolicyCategory) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetIngestionSource sets the ingestionSource property value. Category Ingestion source
func (m *GroupPolicyCategory) SetIngestionSource(value *IngestionSource)() {
    m.ingestionSource = value
}
// SetIsRoot sets the isRoot property value. Defines if the category is a root category
func (m *GroupPolicyCategory) SetIsRoot(value *bool)() {
    m.isRoot = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time the entity was last modified.
func (m *GroupPolicyCategory) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetParent sets the parent property value. The parent category
func (m *GroupPolicyCategory) SetParent(value GroupPolicyCategoryable)() {
    m.parent = value
}
