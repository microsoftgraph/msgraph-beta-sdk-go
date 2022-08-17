package models

import (
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
    res["localizedTexts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOrganizationalMessageLocalizedTextFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OrganizationalMessageLocalizedTextable, len(val))
            for i, v := range val {
                res[i] = v.(OrganizationalMessageLocalizedTextable)
            }
            m.SetLocalizedTexts(res)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["variantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVariantId(val)
        }
        return nil
    }
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLocalizedTexts()))
        for i, v := range m.GetLocalizedTexts() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
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
