package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AddWatermarkAction 
type AddWatermarkAction struct {
    InformationProtectionAction
    // The OdataType property
    OdataType *string
}
// NewAddWatermarkAction instantiates a new addWatermarkAction and sets the default values.
func NewAddWatermarkAction()(*AddWatermarkAction) {
    m := &AddWatermarkAction{
        InformationProtectionAction: *NewInformationProtectionAction(),
    }
    odataTypeValue := "#microsoft.graph.addWatermarkAction"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAddWatermarkActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAddWatermarkActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAddWatermarkAction(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AddWatermarkAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.InformationProtectionAction.GetFieldDeserializers()
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
    res["layout"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWatermarkLayout)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLayout(val.(*WatermarkLayout))
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
// GetFontColor gets the fontColor property value. Color of the font to use for the watermark.
func (m *AddWatermarkAction) GetFontColor()(*string) {
    val, err := m.GetBackingStore().Get("fontColor")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFontName gets the fontName property value. Name of the font to use for the watermark.
func (m *AddWatermarkAction) GetFontName()(*string) {
    val, err := m.GetBackingStore().Get("fontName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFontSize gets the fontSize property value. Font size to use for the watermark.
func (m *AddWatermarkAction) GetFontSize()(*int32) {
    val, err := m.GetBackingStore().Get("fontSize")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetLayout gets the layout property value. The layout property
func (m *AddWatermarkAction) GetLayout()(*WatermarkLayout) {
    val, err := m.GetBackingStore().Get("layout")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*WatermarkLayout)
    }
    return nil
}
// GetText gets the text property value. The contents of the watermark itself.
func (m *AddWatermarkAction) GetText()(*string) {
    val, err := m.GetBackingStore().Get("text")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUiElementName gets the uiElementName property value. The name of the UI element where the watermark should be placed.
func (m *AddWatermarkAction) GetUiElementName()(*string) {
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
func (m *AddWatermarkAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.InformationProtectionAction.Serialize(writer)
    if err != nil {
        return err
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
    if m.GetLayout() != nil {
        cast := (*m.GetLayout()).String()
        err = writer.WriteStringValue("layout", &cast)
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
// SetFontColor sets the fontColor property value. Color of the font to use for the watermark.
func (m *AddWatermarkAction) SetFontColor(value *string)() {
    err := m.GetBackingStore().Set("fontColor", value)
    if err != nil {
        panic(err)
    }
}
// SetFontName sets the fontName property value. Name of the font to use for the watermark.
func (m *AddWatermarkAction) SetFontName(value *string)() {
    err := m.GetBackingStore().Set("fontName", value)
    if err != nil {
        panic(err)
    }
}
// SetFontSize sets the fontSize property value. Font size to use for the watermark.
func (m *AddWatermarkAction) SetFontSize(value *int32)() {
    err := m.GetBackingStore().Set("fontSize", value)
    if err != nil {
        panic(err)
    }
}
// SetLayout sets the layout property value. The layout property
func (m *AddWatermarkAction) SetLayout(value *WatermarkLayout)() {
    err := m.GetBackingStore().Set("layout", value)
    if err != nil {
        panic(err)
    }
}
// SetText sets the text property value. The contents of the watermark itself.
func (m *AddWatermarkAction) SetText(value *string)() {
    err := m.GetBackingStore().Set("text", value)
    if err != nil {
        panic(err)
    }
}
// SetUiElementName sets the uiElementName property value. The name of the UI element where the watermark should be placed.
func (m *AddWatermarkAction) SetUiElementName(value *string)() {
    err := m.GetBackingStore().Set("uiElementName", value)
    if err != nil {
        panic(err)
    }
}
// AddWatermarkActionable 
type AddWatermarkActionable interface {
    InformationProtectionActionable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFontColor()(*string)
    GetFontName()(*string)
    GetFontSize()(*int32)
    GetLayout()(*WatermarkLayout)
    GetText()(*string)
    GetUiElementName()(*string)
    SetFontColor(value *string)()
    SetFontName(value *string)()
    SetFontSize(value *int32)()
    SetLayout(value *WatermarkLayout)()
    SetText(value *string)()
    SetUiElementName(value *string)()
}
