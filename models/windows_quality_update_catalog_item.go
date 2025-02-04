package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsQualityUpdateCatalogItem windows update catalog item entity
type WindowsQualityUpdateCatalogItem struct {
    WindowsUpdateCatalogItem
}
// NewWindowsQualityUpdateCatalogItem instantiates a new WindowsQualityUpdateCatalogItem and sets the default values.
func NewWindowsQualityUpdateCatalogItem()(*WindowsQualityUpdateCatalogItem) {
    m := &WindowsQualityUpdateCatalogItem{
        WindowsUpdateCatalogItem: *NewWindowsUpdateCatalogItem(),
    }
    odataTypeValue := "#microsoft.graph.windowsQualityUpdateCatalogItem"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWindowsQualityUpdateCatalogItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWindowsQualityUpdateCatalogItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsQualityUpdateCatalogItem(), nil
}
// GetClassification gets the classification property value. Windows quality update category
// returns a *WindowsQualityUpdateCategory when successful
func (m *WindowsQualityUpdateCatalogItem) GetClassification()(*WindowsQualityUpdateCategory) {
    val, err := m.GetBackingStore().Get("classification")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*WindowsQualityUpdateCategory)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WindowsQualityUpdateCatalogItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.WindowsUpdateCatalogItem.GetFieldDeserializers()
    res["classification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsQualityUpdateCategory)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClassification(val.(*WindowsQualityUpdateCategory))
        }
        return nil
    }
    res["isExpeditable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsExpeditable(val)
        }
        return nil
    }
    res["kbArticleId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKbArticleId(val)
        }
        return nil
    }
    res["productRevisions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsQualityUpdateCatalogProductRevisionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsQualityUpdateCatalogProductRevisionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(WindowsQualityUpdateCatalogProductRevisionable)
                }
            }
            m.SetProductRevisions(res)
        }
        return nil
    }
    res["qualityUpdateCadence"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsQualityUpdateCadence)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQualityUpdateCadence(val.(*WindowsQualityUpdateCadence))
        }
        return nil
    }
    return res
}
// GetIsExpeditable gets the isExpeditable property value. Flag indicating if update qualifies for expedite
// returns a *bool when successful
func (m *WindowsQualityUpdateCatalogItem) GetIsExpeditable()(*bool) {
    val, err := m.GetBackingStore().Get("isExpeditable")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetKbArticleId gets the kbArticleId property value. Knowledge base article id
// returns a *string when successful
func (m *WindowsQualityUpdateCatalogItem) GetKbArticleId()(*string) {
    val, err := m.GetBackingStore().Get("kbArticleId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetProductRevisions gets the productRevisions property value. The operating system product revisions that are released as part of this quality update. Read-only.
// returns a []WindowsQualityUpdateCatalogProductRevisionable when successful
func (m *WindowsQualityUpdateCatalogItem) GetProductRevisions()([]WindowsQualityUpdateCatalogProductRevisionable) {
    val, err := m.GetBackingStore().Get("productRevisions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]WindowsQualityUpdateCatalogProductRevisionable)
    }
    return nil
}
// GetQualityUpdateCadence gets the qualityUpdateCadence property value. The publishing cadence of the quality update. Possible values are: monthly, outOfBand. This property cannot be modified and is automatically populated when the catalog is created.
// returns a *WindowsQualityUpdateCadence when successful
func (m *WindowsQualityUpdateCatalogItem) GetQualityUpdateCadence()(*WindowsQualityUpdateCadence) {
    val, err := m.GetBackingStore().Get("qualityUpdateCadence")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*WindowsQualityUpdateCadence)
    }
    return nil
}
// Serialize serializes information the current object
func (m *WindowsQualityUpdateCatalogItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.WindowsUpdateCatalogItem.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetClassification() != nil {
        cast := (*m.GetClassification()).String()
        err = writer.WriteStringValue("classification", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isExpeditable", m.GetIsExpeditable())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("kbArticleId", m.GetKbArticleId())
        if err != nil {
            return err
        }
    }
    if m.GetProductRevisions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetProductRevisions()))
        for i, v := range m.GetProductRevisions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("productRevisions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetQualityUpdateCadence() != nil {
        cast := (*m.GetQualityUpdateCadence()).String()
        err = writer.WriteStringValue("qualityUpdateCadence", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetClassification sets the classification property value. Windows quality update category
func (m *WindowsQualityUpdateCatalogItem) SetClassification(value *WindowsQualityUpdateCategory)() {
    err := m.GetBackingStore().Set("classification", value)
    if err != nil {
        panic(err)
    }
}
// SetIsExpeditable sets the isExpeditable property value. Flag indicating if update qualifies for expedite
func (m *WindowsQualityUpdateCatalogItem) SetIsExpeditable(value *bool)() {
    err := m.GetBackingStore().Set("isExpeditable", value)
    if err != nil {
        panic(err)
    }
}
// SetKbArticleId sets the kbArticleId property value. Knowledge base article id
func (m *WindowsQualityUpdateCatalogItem) SetKbArticleId(value *string)() {
    err := m.GetBackingStore().Set("kbArticleId", value)
    if err != nil {
        panic(err)
    }
}
// SetProductRevisions sets the productRevisions property value. The operating system product revisions that are released as part of this quality update. Read-only.
func (m *WindowsQualityUpdateCatalogItem) SetProductRevisions(value []WindowsQualityUpdateCatalogProductRevisionable)() {
    err := m.GetBackingStore().Set("productRevisions", value)
    if err != nil {
        panic(err)
    }
}
// SetQualityUpdateCadence sets the qualityUpdateCadence property value. The publishing cadence of the quality update. Possible values are: monthly, outOfBand. This property cannot be modified and is automatically populated when the catalog is created.
func (m *WindowsQualityUpdateCatalogItem) SetQualityUpdateCadence(value *WindowsQualityUpdateCadence)() {
    err := m.GetBackingStore().Set("qualityUpdateCadence", value)
    if err != nil {
        panic(err)
    }
}
type WindowsQualityUpdateCatalogItemable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    WindowsUpdateCatalogItemable
    GetClassification()(*WindowsQualityUpdateCategory)
    GetIsExpeditable()(*bool)
    GetKbArticleId()(*string)
    GetProductRevisions()([]WindowsQualityUpdateCatalogProductRevisionable)
    GetQualityUpdateCadence()(*WindowsQualityUpdateCadence)
    SetClassification(value *WindowsQualityUpdateCategory)()
    SetIsExpeditable(value *bool)()
    SetKbArticleId(value *string)()
    SetProductRevisions(value []WindowsQualityUpdateCatalogProductRevisionable)()
    SetQualityUpdateCadence(value *WindowsQualityUpdateCadence)()
}
