package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ChangeItemBase struct {
    Entity
}
// NewChangeItemBase instantiates a new ChangeItemBase and sets the default values.
func NewChangeItemBase()(*ChangeItemBase) {
    m := &ChangeItemBase{
        Entity: *NewEntity(),
    }
    return m
}
// CreateChangeItemBaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateChangeItemBaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.announcement":
                        return NewAnnouncement(), nil
                    case "#microsoft.graph.roadmap":
                        return NewRoadmap(), nil
                }
            }
        }
    }
    return NewChangeItemBase(), nil
}
// GetChangeItemService gets the changeItemService property value. Specifies the Microsoft Entra service name to which this item belongs. Supports $filter (eq, ne, in) and $orderby.
// returns a *string when successful
func (m *ChangeItemBase) GetChangeItemService()(*string) {
    val, err := m.GetBackingStore().Get("changeItemService")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDescription gets the description property value. Description of the new feature or change announcement. Supports $filter (eq, ne, in, startswith) and $orderby.
// returns a *string when successful
func (m *ChangeItemBase) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDocumentationUrls gets the documentationUrls property value. Link to the feature or change documentation. Supports $filter (any with eq).
// returns a []string when successful
func (m *ChangeItemBase) GetDocumentationUrls()([]string) {
    val, err := m.GetBackingStore().Get("documentationUrls")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ChangeItemBase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["changeItemService"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChangeItemService(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["documentationUrls"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetDocumentationUrls(res)
        }
        return nil
    }
    res["shortDescription"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShortDescription(val)
        }
        return nil
    }
    res["systemTags"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetSystemTags(res)
        }
        return nil
    }
    res["tags"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetTags(res)
        }
        return nil
    }
    res["title"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val)
        }
        return nil
    }
    return res
}
// GetShortDescription gets the shortDescription property value. A short description of the feature or change. Supports $filter (eq, ne, in, startswith) and $orderby.
// returns a *string when successful
func (m *ChangeItemBase) GetShortDescription()(*string) {
    val, err := m.GetBackingStore().Get("shortDescription")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSystemTags gets the systemTags property value. Microsoft Entra-specific tags. Example values: Top announcement - entraroadmaphighlightproductnews, New release highlight - entraroadmaphighlightnewfeature. Supports $filter (any with eq).
// returns a []string when successful
func (m *ChangeItemBase) GetSystemTags()([]string) {
    val, err := m.GetBackingStore().Get("systemTags")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetTags gets the tags property value. Identity and Access Management (IAM) related tags. Example values: External Identities, Reliability and Resilience. Supports $filter (any with eq).
// returns a []string when successful
func (m *ChangeItemBase) GetTags()([]string) {
    val, err := m.GetBackingStore().Get("tags")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetTitle gets the title property value. Title of the feature or change. Supports $filter (eq, ne, in, startswith) and $orderby.
// returns a *string when successful
func (m *ChangeItemBase) GetTitle()(*string) {
    val, err := m.GetBackingStore().Get("title")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ChangeItemBase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("changeItemService", m.GetChangeItemService())
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
    if m.GetDocumentationUrls() != nil {
        err = writer.WriteCollectionOfStringValues("documentationUrls", m.GetDocumentationUrls())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("shortDescription", m.GetShortDescription())
        if err != nil {
            return err
        }
    }
    if m.GetSystemTags() != nil {
        err = writer.WriteCollectionOfStringValues("systemTags", m.GetSystemTags())
        if err != nil {
            return err
        }
    }
    if m.GetTags() != nil {
        err = writer.WriteCollectionOfStringValues("tags", m.GetTags())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("title", m.GetTitle())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetChangeItemService sets the changeItemService property value. Specifies the Microsoft Entra service name to which this item belongs. Supports $filter (eq, ne, in) and $orderby.
func (m *ChangeItemBase) SetChangeItemService(value *string)() {
    err := m.GetBackingStore().Set("changeItemService", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. Description of the new feature or change announcement. Supports $filter (eq, ne, in, startswith) and $orderby.
func (m *ChangeItemBase) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDocumentationUrls sets the documentationUrls property value. Link to the feature or change documentation. Supports $filter (any with eq).
func (m *ChangeItemBase) SetDocumentationUrls(value []string)() {
    err := m.GetBackingStore().Set("documentationUrls", value)
    if err != nil {
        panic(err)
    }
}
// SetShortDescription sets the shortDescription property value. A short description of the feature or change. Supports $filter (eq, ne, in, startswith) and $orderby.
func (m *ChangeItemBase) SetShortDescription(value *string)() {
    err := m.GetBackingStore().Set("shortDescription", value)
    if err != nil {
        panic(err)
    }
}
// SetSystemTags sets the systemTags property value. Microsoft Entra-specific tags. Example values: Top announcement - entraroadmaphighlightproductnews, New release highlight - entraroadmaphighlightnewfeature. Supports $filter (any with eq).
func (m *ChangeItemBase) SetSystemTags(value []string)() {
    err := m.GetBackingStore().Set("systemTags", value)
    if err != nil {
        panic(err)
    }
}
// SetTags sets the tags property value. Identity and Access Management (IAM) related tags. Example values: External Identities, Reliability and Resilience. Supports $filter (any with eq).
func (m *ChangeItemBase) SetTags(value []string)() {
    err := m.GetBackingStore().Set("tags", value)
    if err != nil {
        panic(err)
    }
}
// SetTitle sets the title property value. Title of the feature or change. Supports $filter (eq, ne, in, startswith) and $orderby.
func (m *ChangeItemBase) SetTitle(value *string)() {
    err := m.GetBackingStore().Set("title", value)
    if err != nil {
        panic(err)
    }
}
type ChangeItemBaseable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetChangeItemService()(*string)
    GetDescription()(*string)
    GetDocumentationUrls()([]string)
    GetShortDescription()(*string)
    GetSystemTags()([]string)
    GetTags()([]string)
    GetTitle()(*string)
    SetChangeItemService(value *string)()
    SetDescription(value *string)()
    SetDocumentationUrls(value []string)()
    SetShortDescription(value *string)()
    SetSystemTags(value []string)()
    SetTags(value []string)()
    SetTitle(value *string)()
}
