package security

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AddWatermarkAction 
type AddWatermarkAction struct {
    InformationProtectionAction
    // Color of the font to use for the watermark.
    fontColor *string
    // Name of the font to use for the watermark.
    fontName *string
    // Font size to use for the watermark.
    fontSize *int32
    // The layout property
    layout *WatermarkLayout
    // The contents of the watermark itself.
    text *string
    // The name of the UI element where the watermark should be placed.
    uiElementName *string
}
// NewAddWatermarkAction instantiates a new AddWatermarkAction and sets the default values.
func NewAddWatermarkAction()(*AddWatermarkAction) {
    m := &AddWatermarkAction{
        InformationProtectionAction: *NewInformationProtectionAction(),
    }
    odataTypeValue := "#microsoft.graph.security.addWatermarkAction";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAddWatermarkActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAddWatermarkActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAddWatermarkAction(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AddWatermarkAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.InformationProtectionAction.GetFieldDeserializers()
    res["fontColor"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetFontColor)
    res["fontName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetFontName)
    res["fontSize"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetFontSize)
    res["layout"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseWatermarkLayout , m.SetLayout)
    res["text"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetText)
    res["uiElementName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUiElementName)
    return res
}
// GetFontColor gets the fontColor property value. Color of the font to use for the watermark.
func (m *AddWatermarkAction) GetFontColor()(*string) {
    return m.fontColor
}
// GetFontName gets the fontName property value. Name of the font to use for the watermark.
func (m *AddWatermarkAction) GetFontName()(*string) {
    return m.fontName
}
// GetFontSize gets the fontSize property value. Font size to use for the watermark.
func (m *AddWatermarkAction) GetFontSize()(*int32) {
    return m.fontSize
}
// GetLayout gets the layout property value. The layout property
func (m *AddWatermarkAction) GetLayout()(*WatermarkLayout) {
    return m.layout
}
// GetText gets the text property value. The contents of the watermark itself.
func (m *AddWatermarkAction) GetText()(*string) {
    return m.text
}
// GetUiElementName gets the uiElementName property value. The name of the UI element where the watermark should be placed.
func (m *AddWatermarkAction) GetUiElementName()(*string) {
    return m.uiElementName
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
    m.fontColor = value
}
// SetFontName sets the fontName property value. Name of the font to use for the watermark.
func (m *AddWatermarkAction) SetFontName(value *string)() {
    m.fontName = value
}
// SetFontSize sets the fontSize property value. Font size to use for the watermark.
func (m *AddWatermarkAction) SetFontSize(value *int32)() {
    m.fontSize = value
}
// SetLayout sets the layout property value. The layout property
func (m *AddWatermarkAction) SetLayout(value *WatermarkLayout)() {
    m.layout = value
}
// SetText sets the text property value. The contents of the watermark itself.
func (m *AddWatermarkAction) SetText(value *string)() {
    m.text = value
}
// SetUiElementName sets the uiElementName property value. The name of the UI element where the watermark should be placed.
func (m *AddWatermarkAction) SetUiElementName(value *string)() {
    m.uiElementName = value
}
