package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OrganizationalMessageVariant contains the localized text to be displayed for a given variant
type OrganizationalMessageVariant struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The list of localized texts for this variant
    localizedTexts []OrganizationalMessageLocalizedTextable
    // The name of the variant
    name *string
    // The OdataType property
    odataType *string
    // A unique identifier for this variant
    variantId *string
}
// NewOrganizationalMessageVariant instantiates a new organizationalMessageVariant and sets the default values.
func NewOrganizationalMessageVariant()(*OrganizationalMessageVariant) {
    m := &OrganizationalMessageVariant{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.organizationalMessageVariant";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateOrganizationalMessageVariantFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOrganizationalMessageVariantFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOrganizationalMessageVariant(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OrganizationalMessageVariant) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OrganizationalMessageVariant) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["localizedTexts"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateOrganizationalMessageLocalizedTextFromDiscriminatorValue , m.SetLocalizedTexts)
    res["name"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetName)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["variantId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetVariantId)
    return res
}
// GetLocalizedTexts gets the localizedTexts property value. The list of localized texts for this variant
func (m *OrganizationalMessageVariant) GetLocalizedTexts()([]OrganizationalMessageLocalizedTextable) {
    return m.localizedTexts
}
// GetName gets the name property value. The name of the variant
func (m *OrganizationalMessageVariant) GetName()(*string) {
    return m.name
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *OrganizationalMessageVariant) GetOdataType()(*string) {
    return m.odataType
}
// GetVariantId gets the variantId property value. A unique identifier for this variant
func (m *OrganizationalMessageVariant) GetVariantId()(*string) {
    return m.variantId
}
// Serialize serializes information the current object
func (m *OrganizationalMessageVariant) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetLocalizedTexts() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetLocalizedTexts())
        err := writer.WriteCollectionOfObjectValues("localizedTexts", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
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
        err := writer.WriteStringValue("variantId", m.GetVariantId())
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
func (m *OrganizationalMessageVariant) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetLocalizedTexts sets the localizedTexts property value. The list of localized texts for this variant
func (m *OrganizationalMessageVariant) SetLocalizedTexts(value []OrganizationalMessageLocalizedTextable)() {
    m.localizedTexts = value
}
// SetName sets the name property value. The name of the variant
func (m *OrganizationalMessageVariant) SetName(value *string)() {
    m.name = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *OrganizationalMessageVariant) SetOdataType(value *string)() {
    m.odataType = value
}
// SetVariantId sets the variantId property value. A unique identifier for this variant
func (m *OrganizationalMessageVariant) SetVariantId(value *string)() {
    m.variantId = value
}
