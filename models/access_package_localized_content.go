package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessPackageLocalizedContent 
type AccessPackageLocalizedContent struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The fallback string, which is used when a requested localization is not available. Required.
    defaultText *string
    // Content represented in a format for a specific locale.
    localizedTexts []AccessPackageLocalizedTextable
    // The OdataType property
    odataType *string
}
// NewAccessPackageLocalizedContent instantiates a new accessPackageLocalizedContent and sets the default values.
func NewAccessPackageLocalizedContent()(*AccessPackageLocalizedContent) {
    m := &AccessPackageLocalizedContent{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.accessPackageLocalizedContent";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAccessPackageLocalizedContentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessPackageLocalizedContentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessPackageLocalizedContent(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AccessPackageLocalizedContent) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDefaultText gets the defaultText property value. The fallback string, which is used when a requested localization is not available. Required.
func (m *AccessPackageLocalizedContent) GetDefaultText()(*string) {
    return m.defaultText
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessPackageLocalizedContent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["defaultText"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDefaultText)
    res["localizedTexts"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAccessPackageLocalizedTextFromDiscriminatorValue , m.SetLocalizedTexts)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    return res
}
// GetLocalizedTexts gets the localizedTexts property value. Content represented in a format for a specific locale.
func (m *AccessPackageLocalizedContent) GetLocalizedTexts()([]AccessPackageLocalizedTextable) {
    return m.localizedTexts
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AccessPackageLocalizedContent) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *AccessPackageLocalizedContent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("defaultText", m.GetDefaultText())
        if err != nil {
            return err
        }
    }
    if m.GetLocalizedTexts() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetLocalizedTexts())
        err := writer.WriteCollectionOfObjectValues("localizedTexts", cast)
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
func (m *AccessPackageLocalizedContent) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDefaultText sets the defaultText property value. The fallback string, which is used when a requested localization is not available. Required.
func (m *AccessPackageLocalizedContent) SetDefaultText(value *string)() {
    m.defaultText = value
}
// SetLocalizedTexts sets the localizedTexts property value. Content represented in a format for a specific locale.
func (m *AccessPackageLocalizedContent) SetLocalizedTexts(value []AccessPackageLocalizedTextable)() {
    m.localizedTexts = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AccessPackageLocalizedContent) SetOdataType(value *string)() {
    m.odataType = value
}
