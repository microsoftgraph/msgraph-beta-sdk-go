package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessPackageLocalizedText 
type AccessPackageLocalizedText struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The ISO code for the intended language. Required.
    languageCode *string
    // The OdataType property
    odataType *string
    // The text in the specific language. Required.
    text *string
}
// NewAccessPackageLocalizedText instantiates a new accessPackageLocalizedText and sets the default values.
func NewAccessPackageLocalizedText()(*AccessPackageLocalizedText) {
    m := &AccessPackageLocalizedText{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.accessPackageLocalizedText";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAccessPackageLocalizedTextFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessPackageLocalizedTextFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessPackageLocalizedText(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AccessPackageLocalizedText) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessPackageLocalizedText) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["languageCode"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetLanguageCode)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["text"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetText)
    return res
}
// GetLanguageCode gets the languageCode property value. The ISO code for the intended language. Required.
func (m *AccessPackageLocalizedText) GetLanguageCode()(*string) {
    return m.languageCode
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AccessPackageLocalizedText) GetOdataType()(*string) {
    return m.odataType
}
// GetText gets the text property value. The text in the specific language. Required.
func (m *AccessPackageLocalizedText) GetText()(*string) {
    return m.text
}
// Serialize serializes information the current object
func (m *AccessPackageLocalizedText) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("languageCode", m.GetLanguageCode())
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
        err := writer.WriteStringValue("text", m.GetText())
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
func (m *AccessPackageLocalizedText) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetLanguageCode sets the languageCode property value. The ISO code for the intended language. Required.
func (m *AccessPackageLocalizedText) SetLanguageCode(value *string)() {
    m.languageCode = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AccessPackageLocalizedText) SetOdataType(value *string)() {
    m.odataType = value
}
// SetText sets the text property value. The text in the specific language. Required.
func (m *AccessPackageLocalizedText) SetText(value *string)() {
    m.text = value
}
