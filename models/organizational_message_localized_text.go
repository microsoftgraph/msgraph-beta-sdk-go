package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OrganizationalMessageLocalizedText contains the text to be displayed for a given locale
type OrganizationalMessageLocalizedText struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The locale for this text
    locale *string
    // The OdataType property
    odataType *string
    // The text that will be displayed to a user from this specific locale
    text OrganizationalMessageTextable
}
// NewOrganizationalMessageLocalizedText instantiates a new organizationalMessageLocalizedText and sets the default values.
func NewOrganizationalMessageLocalizedText()(*OrganizationalMessageLocalizedText) {
    m := &OrganizationalMessageLocalizedText{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.organizationalMessageLocalizedText";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateOrganizationalMessageLocalizedTextFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOrganizationalMessageLocalizedTextFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOrganizationalMessageLocalizedText(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OrganizationalMessageLocalizedText) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OrganizationalMessageLocalizedText) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["locale"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetLocale)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["text"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateOrganizationalMessageTextFromDiscriminatorValue , m.SetText)
    return res
}
// GetLocale gets the locale property value. The locale for this text
func (m *OrganizationalMessageLocalizedText) GetLocale()(*string) {
    return m.locale
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *OrganizationalMessageLocalizedText) GetOdataType()(*string) {
    return m.odataType
}
// GetText gets the text property value. The text that will be displayed to a user from this specific locale
func (m *OrganizationalMessageLocalizedText) GetText()(OrganizationalMessageTextable) {
    return m.text
}
// Serialize serializes information the current object
func (m *OrganizationalMessageLocalizedText) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("locale", m.GetLocale())
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
        err := writer.WriteObjectValue("text", m.GetText())
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
func (m *OrganizationalMessageLocalizedText) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetLocale sets the locale property value. The locale for this text
func (m *OrganizationalMessageLocalizedText) SetLocale(value *string)() {
    m.locale = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *OrganizationalMessageLocalizedText) SetOdataType(value *string)() {
    m.odataType = value
}
// SetText sets the text property value. The text that will be displayed to a user from this specific locale
func (m *OrganizationalMessageLocalizedText) SetText(value OrganizationalMessageTextable)() {
    m.text = value
}
