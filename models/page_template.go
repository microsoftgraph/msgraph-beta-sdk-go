package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PageTemplate struct {
    BaseSitePage
}
// NewPageTemplate instantiates a new PageTemplate and sets the default values.
func NewPageTemplate()(*PageTemplate) {
    m := &PageTemplate{
        BaseSitePage: *NewBaseSitePage(),
    }
    return m
}
// CreatePageTemplateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePageTemplateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPageTemplate(), nil
}
// GetCanvasLayout gets the canvasLayout property value. The layout of the content in a given SharePoint page template, including horizontal sections and vertical sections.
// returns a CanvasLayoutable when successful
func (m *PageTemplate) GetCanvasLayout()(CanvasLayoutable) {
    val, err := m.GetBackingStore().Get("canvasLayout")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(CanvasLayoutable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PageTemplate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseSitePage.GetFieldDeserializers()
    res["canvasLayout"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCanvasLayoutFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCanvasLayout(val.(CanvasLayoutable))
        }
        return nil
    }
    res["titleArea"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTitleAreaFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitleArea(val.(TitleAreaable))
        }
        return nil
    }
    res["webParts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWebPartFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WebPartable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(WebPartable)
                }
            }
            m.SetWebParts(res)
        }
        return nil
    }
    return res
}
// GetTitleArea gets the titleArea property value. The title area on the SharePoint page template.
// returns a TitleAreaable when successful
func (m *PageTemplate) GetTitleArea()(TitleAreaable) {
    val, err := m.GetBackingStore().Get("titleArea")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TitleAreaable)
    }
    return nil
}
// GetWebParts gets the webParts property value. The collection of web parts on the SharePoint page.
// returns a []WebPartable when successful
func (m *PageTemplate) GetWebParts()([]WebPartable) {
    val, err := m.GetBackingStore().Get("webParts")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]WebPartable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *PageTemplate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseSitePage.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("canvasLayout", m.GetCanvasLayout())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("titleArea", m.GetTitleArea())
        if err != nil {
            return err
        }
    }
    if m.GetWebParts() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWebParts()))
        for i, v := range m.GetWebParts() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("webParts", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCanvasLayout sets the canvasLayout property value. The layout of the content in a given SharePoint page template, including horizontal sections and vertical sections.
func (m *PageTemplate) SetCanvasLayout(value CanvasLayoutable)() {
    err := m.GetBackingStore().Set("canvasLayout", value)
    if err != nil {
        panic(err)
    }
}
// SetTitleArea sets the titleArea property value. The title area on the SharePoint page template.
func (m *PageTemplate) SetTitleArea(value TitleAreaable)() {
    err := m.GetBackingStore().Set("titleArea", value)
    if err != nil {
        panic(err)
    }
}
// SetWebParts sets the webParts property value. The collection of web parts on the SharePoint page.
func (m *PageTemplate) SetWebParts(value []WebPartable)() {
    err := m.GetBackingStore().Set("webParts", value)
    if err != nil {
        panic(err)
    }
}
type PageTemplateable interface {
    BaseSitePageable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCanvasLayout()(CanvasLayoutable)
    GetTitleArea()(TitleAreaable)
    GetWebParts()([]WebPartable)
    SetCanvasLayout(value CanvasLayoutable)()
    SetTitleArea(value TitleAreaable)()
    SetWebParts(value []WebPartable)()
}
