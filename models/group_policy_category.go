package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupPolicyCategory 
type GroupPolicyCategory struct {
    Entity
}
// NewGroupPolicyCategory instantiates a new groupPolicyCategory and sets the default values.
func NewGroupPolicyCategory()(*GroupPolicyCategory) {
    m := &GroupPolicyCategory{
        Entity: *NewEntity(),
    }
    return m
}
// CreateGroupPolicyCategoryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupPolicyCategoryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupPolicyCategory(), nil
}
// GetChildren gets the children property value. The children categories
func (m *GroupPolicyCategory) GetChildren()([]GroupPolicyCategoryable) {
    val, err := m.GetBackingStore().Get("children")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]GroupPolicyCategoryable)
    }
    return nil
}
// GetDefinitionFile gets the definitionFile property value. The id of the definition file the category came from
func (m *GroupPolicyCategory) GetDefinitionFile()(GroupPolicyDefinitionFileable) {
    val, err := m.GetBackingStore().Get("definitionFile")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(GroupPolicyDefinitionFileable)
    }
    return nil
}
// GetDefinitions gets the definitions property value. The immediate GroupPolicyDefinition children of the category
func (m *GroupPolicyCategory) GetDefinitions()([]GroupPolicyDefinitionable) {
    val, err := m.GetBackingStore().Get("definitions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]GroupPolicyDefinitionable)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The string id of the category's display name
func (m *GroupPolicyCategory) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupPolicyCategory) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["children"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGroupPolicyCategoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GroupPolicyCategoryable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(GroupPolicyCategoryable)
                }
            }
            m.SetChildren(res)
        }
        return nil
    }
    res["definitionFile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupPolicyDefinitionFileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefinitionFile(val.(GroupPolicyDefinitionFileable))
        }
        return nil
    }
    res["definitions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGroupPolicyDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GroupPolicyDefinitionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(GroupPolicyDefinitionable)
                }
            }
            m.SetDefinitions(res)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["ingestionSource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseIngestionSource)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIngestionSource(val.(*IngestionSource))
        }
        return nil
    }
    res["isRoot"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsRoot(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["parent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupPolicyCategoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParent(val.(GroupPolicyCategoryable))
        }
        return nil
    }
    return res
}
// GetIngestionSource gets the ingestionSource property value. Category Ingestion source
func (m *GroupPolicyCategory) GetIngestionSource()(*IngestionSource) {
    val, err := m.GetBackingStore().Get("ingestionSource")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*IngestionSource)
    }
    return nil
}
// GetIsRoot gets the isRoot property value. Defines if the category is a root category
func (m *GroupPolicyCategory) GetIsRoot()(*bool) {
    val, err := m.GetBackingStore().Get("isRoot")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time the entity was last modified.
func (m *GroupPolicyCategory) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastModifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetParent gets the parent property value. The parent category
func (m *GroupPolicyCategory) GetParent()(GroupPolicyCategoryable) {
    val, err := m.GetBackingStore().Get("parent")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(GroupPolicyCategoryable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *GroupPolicyCategory) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetChildren() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetChildren()))
        for i, v := range m.GetChildren() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDefinitions()))
        for i, v := range m.GetDefinitions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
    err := m.GetBackingStore().Set("children", value)
    if err != nil {
        panic(err)
    }
}
// SetDefinitionFile sets the definitionFile property value. The id of the definition file the category came from
func (m *GroupPolicyCategory) SetDefinitionFile(value GroupPolicyDefinitionFileable)() {
    err := m.GetBackingStore().Set("definitionFile", value)
    if err != nil {
        panic(err)
    }
}
// SetDefinitions sets the definitions property value. The immediate GroupPolicyDefinition children of the category
func (m *GroupPolicyCategory) SetDefinitions(value []GroupPolicyDefinitionable)() {
    err := m.GetBackingStore().Set("definitions", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The string id of the category's display name
func (m *GroupPolicyCategory) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetIngestionSource sets the ingestionSource property value. Category Ingestion source
func (m *GroupPolicyCategory) SetIngestionSource(value *IngestionSource)() {
    err := m.GetBackingStore().Set("ingestionSource", value)
    if err != nil {
        panic(err)
    }
}
// SetIsRoot sets the isRoot property value. Defines if the category is a root category
func (m *GroupPolicyCategory) SetIsRoot(value *bool)() {
    err := m.GetBackingStore().Set("isRoot", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time the entity was last modified.
func (m *GroupPolicyCategory) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetParent sets the parent property value. The parent category
func (m *GroupPolicyCategory) SetParent(value GroupPolicyCategoryable)() {
    err := m.GetBackingStore().Set("parent", value)
    if err != nil {
        panic(err)
    }
}
// GroupPolicyCategoryable 
type GroupPolicyCategoryable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetChildren()([]GroupPolicyCategoryable)
    GetDefinitionFile()(GroupPolicyDefinitionFileable)
    GetDefinitions()([]GroupPolicyDefinitionable)
    GetDisplayName()(*string)
    GetIngestionSource()(*IngestionSource)
    GetIsRoot()(*bool)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetParent()(GroupPolicyCategoryable)
    SetChildren(value []GroupPolicyCategoryable)()
    SetDefinitionFile(value GroupPolicyDefinitionFileable)()
    SetDefinitions(value []GroupPolicyDefinitionable)()
    SetDisplayName(value *string)()
    SetIngestionSource(value *IngestionSource)()
    SetIsRoot(value *bool)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetParent(value GroupPolicyCategoryable)()
}
