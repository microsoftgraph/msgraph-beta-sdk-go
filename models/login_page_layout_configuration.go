package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LoginPageLayoutConfiguration 
type LoginPageLayoutConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The isFooterShown property
    isFooterShown *bool
    // The isHeaderShown property
    isHeaderShown *bool
    // The layoutTemplateType property
    layoutTemplateType *LayoutTemplateType
    // The OdataType property
    odataType *string
}
// NewLoginPageLayoutConfiguration instantiates a new loginPageLayoutConfiguration and sets the default values.
func NewLoginPageLayoutConfiguration()(*LoginPageLayoutConfiguration) {
    m := &LoginPageLayoutConfiguration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateLoginPageLayoutConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLoginPageLayoutConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLoginPageLayoutConfiguration(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LoginPageLayoutConfiguration) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LoginPageLayoutConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["isFooterShown"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsFooterShown)
    res["isHeaderShown"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsHeaderShown)
    res["layoutTemplateType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseLayoutTemplateType , m.SetLayoutTemplateType)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    return res
}
// GetIsFooterShown gets the isFooterShown property value. The isFooterShown property
func (m *LoginPageLayoutConfiguration) GetIsFooterShown()(*bool) {
    return m.isFooterShown
}
// GetIsHeaderShown gets the isHeaderShown property value. The isHeaderShown property
func (m *LoginPageLayoutConfiguration) GetIsHeaderShown()(*bool) {
    return m.isHeaderShown
}
// GetLayoutTemplateType gets the layoutTemplateType property value. The layoutTemplateType property
func (m *LoginPageLayoutConfiguration) GetLayoutTemplateType()(*LayoutTemplateType) {
    return m.layoutTemplateType
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *LoginPageLayoutConfiguration) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *LoginPageLayoutConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isFooterShown", m.GetIsFooterShown())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isHeaderShown", m.GetIsHeaderShown())
        if err != nil {
            return err
        }
    }
    if m.GetLayoutTemplateType() != nil {
        cast := (*m.GetLayoutTemplateType()).String()
        err := writer.WriteStringValue("layoutTemplateType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LoginPageLayoutConfiguration) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetIsFooterShown sets the isFooterShown property value. The isFooterShown property
func (m *LoginPageLayoutConfiguration) SetIsFooterShown(value *bool)() {
    m.isFooterShown = value
}
// SetIsHeaderShown sets the isHeaderShown property value. The isHeaderShown property
func (m *LoginPageLayoutConfiguration) SetIsHeaderShown(value *bool)() {
    m.isHeaderShown = value
}
// SetLayoutTemplateType sets the layoutTemplateType property value. The layoutTemplateType property
func (m *LoginPageLayoutConfiguration) SetLayoutTemplateType(value *LayoutTemplateType)() {
    m.layoutTemplateType = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *LoginPageLayoutConfiguration) SetOdataType(value *string)() {
    m.odataType = value
}
