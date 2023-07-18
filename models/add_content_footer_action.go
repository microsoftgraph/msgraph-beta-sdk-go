package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AddContentFooterAction 
type AddContentFooterAction struct {
    InformationProtectionAction
    // The OdataType property
    OdataType *string
}
// NewAddContentFooterAction instantiates a new addContentFooterAction and sets the default values.
func NewAddContentFooterAction()(*AddContentFooterAction) {
    m := &AddContentFooterAction{
        InformationProtectionAction: *NewInformationProtectionAction(),
    }
    odataTypeValue := "#microsoft.graph.addContentFooterAction"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAddContentFooterActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAddContentFooterActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAddContentFooterAction(), nil
}
// GetAlignment gets the alignment property value. The alignment property
func (m *AddContentFooterAction) GetAlignment()(*ContentAlignment) {
    val, err := m.GetBackingStore().Get("alignment")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ContentAlignment)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AddContentFooterAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.InformationProtectionAction.GetFieldDeserializers()
    res["alignment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseContentAlignment)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlignment(val.(*ContentAlignment))
        }
        return nil
    }
    res["fontColor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFontColor(val)
        }
        return nil
    }
    res["fontName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFontName(val)
        }
        return nil
    }
    res["fontSize"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFontSize(val)
        }
        return nil
    }
    res["margin"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMargin(val)
        }
        return nil
    }
    res["text"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetText(val)
        }
        return nil
    }
    res["uiElementName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUiElementName(val)
        }
        return nil
    }
    return res
}
// GetFontColor gets the fontColor property value. Color of the font to use for the footer.
func (m *AddContentFooterAction) GetFontColor()(*string) {
    val, err := m.GetBackingStore().Get("fontColor")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFontName gets the fontName property value. Name of the font to use for the footer.
func (m *AddContentFooterAction) GetFontName()(*string) {
    val, err := m.GetBackingStore().Get("fontName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFontSize gets the fontSize property value. Font size to use for the footer.
func (m *AddContentFooterAction) GetFontSize()(*int32) {
    val, err := m.GetBackingStore().Get("fontSize")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetMargin gets the margin property value. The margin of the header from the bottom of the document.
func (m *AddContentFooterAction) GetMargin()(*int32) {
    val, err := m.GetBackingStore().Get("margin")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetText gets the text property value. The contents of the footer itself.
func (m *AddContentFooterAction) GetText()(*string) {
    val, err := m.GetBackingStore().Get("text")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUiElementName gets the uiElementName property value. The name of the UI element where the footer should be placed.
func (m *AddContentFooterAction) GetUiElementName()(*string) {
    val, err := m.GetBackingStore().Get("uiElementName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AddContentFooterAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.InformationProtectionAction.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAlignment() != nil {
        cast := (*m.GetAlignment()).String()
        err = writer.WriteStringValue("alignment", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("fontColor", m.GetFontColor())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("fontName", m.GetFontName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("fontSize", m.GetFontSize())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("margin", m.GetMargin())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("text", m.GetText())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("uiElementName", m.GetUiElementName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAlignment sets the alignment property value. The alignment property
func (m *AddContentFooterAction) SetAlignment(value *ContentAlignment)() {
    err := m.GetBackingStore().Set("alignment", value)
    if err != nil {
        panic(err)
    }
}
// SetFontColor sets the fontColor property value. Color of the font to use for the footer.
func (m *AddContentFooterAction) SetFontColor(value *string)() {
    err := m.GetBackingStore().Set("fontColor", value)
    if err != nil {
        panic(err)
    }
}
// SetFontName sets the fontName property value. Name of the font to use for the footer.
func (m *AddContentFooterAction) SetFontName(value *string)() {
    err := m.GetBackingStore().Set("fontName", value)
    if err != nil {
        panic(err)
    }
}
// SetFontSize sets the fontSize property value. Font size to use for the footer.
func (m *AddContentFooterAction) SetFontSize(value *int32)() {
    err := m.GetBackingStore().Set("fontSize", value)
    if err != nil {
        panic(err)
    }
}
// SetMargin sets the margin property value. The margin of the header from the bottom of the document.
func (m *AddContentFooterAction) SetMargin(value *int32)() {
    err := m.GetBackingStore().Set("margin", value)
    if err != nil {
        panic(err)
    }
}
// SetText sets the text property value. The contents of the footer itself.
func (m *AddContentFooterAction) SetText(value *string)() {
    err := m.GetBackingStore().Set("text", value)
    if err != nil {
        panic(err)
    }
}
// SetUiElementName sets the uiElementName property value. The name of the UI element where the footer should be placed.
func (m *AddContentFooterAction) SetUiElementName(value *string)() {
    err := m.GetBackingStore().Set("uiElementName", value)
    if err != nil {
        panic(err)
    }
}
// AddContentFooterActionable 
type AddContentFooterActionable interface {
    InformationProtectionActionable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAlignment()(*ContentAlignment)
    GetFontColor()(*string)
    GetFontName()(*string)
    GetFontSize()(*int32)
    GetMargin()(*int32)
    GetText()(*string)
    GetUiElementName()(*string)
    SetAlignment(value *ContentAlignment)()
    SetFontColor(value *string)()
    SetFontName(value *string)()
    SetFontSize(value *int32)()
    SetMargin(value *int32)()
    SetText(value *string)()
    SetUiElementName(value *string)()
}
