package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OrganizationalMessagePlacementDetail contains the different variants of text that can be displayed for a given placement within a surface
type OrganizationalMessagePlacementDetail struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The OdataType property
    odataType *string
    // Indicates the name of the placement. Possible values are: default, card0, card1, card2, card3, unknownFutureValue.
    placement *OrganizationalMessagePlacement
    // The list of different variants that can be displayed for this placement
    variants []OrganizationalMessageVariantable
}
// NewOrganizationalMessagePlacementDetail instantiates a new organizationalMessagePlacementDetail and sets the default values.
func NewOrganizationalMessagePlacementDetail()(*OrganizationalMessagePlacementDetail) {
    m := &OrganizationalMessagePlacementDetail{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.organizationalMessagePlacementDetail";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateOrganizationalMessagePlacementDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOrganizationalMessagePlacementDetailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOrganizationalMessagePlacementDetail(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OrganizationalMessagePlacementDetail) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OrganizationalMessagePlacementDetail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["placement"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOrganizationalMessagePlacement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlacement(val.(*OrganizationalMessagePlacement))
        }
        return nil
    }
    res["variants"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOrganizationalMessageVariantFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OrganizationalMessageVariantable, len(val))
            for i, v := range val {
                res[i] = v.(OrganizationalMessageVariantable)
            }
            m.SetVariants(res)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *OrganizationalMessagePlacementDetail) GetOdataType()(*string) {
    return m.odataType
}
// GetPlacement gets the placement property value. Indicates the name of the placement. Possible values are: default, card0, card1, card2, card3, unknownFutureValue.
func (m *OrganizationalMessagePlacementDetail) GetPlacement()(*OrganizationalMessagePlacement) {
    return m.placement
}
// GetVariants gets the variants property value. The list of different variants that can be displayed for this placement
func (m *OrganizationalMessagePlacementDetail) GetVariants()([]OrganizationalMessageVariantable) {
    return m.variants
}
// Serialize serializes information the current object
func (m *OrganizationalMessagePlacementDetail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetPlacement() != nil {
        cast := (*m.GetPlacement()).String()
        err := writer.WriteStringValue("placement", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetVariants() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetVariants()))
        for i, v := range m.GetVariants() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("variants", cast)
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
func (m *OrganizationalMessagePlacementDetail) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *OrganizationalMessagePlacementDetail) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPlacement sets the placement property value. Indicates the name of the placement. Possible values are: default, card0, card1, card2, card3, unknownFutureValue.
func (m *OrganizationalMessagePlacementDetail) SetPlacement(value *OrganizationalMessagePlacement)() {
    m.placement = value
}
// SetVariants sets the variants property value. The list of different variants that can be displayed for this placement
func (m *OrganizationalMessagePlacementDetail) SetVariants(value []OrganizationalMessageVariantable)() {
    m.variants = value
}
