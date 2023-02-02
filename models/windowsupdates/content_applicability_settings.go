package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ContentApplicabilitySettings 
type ContentApplicabilitySettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The OdataType property
    odataType *string
    // The offerWhileRecommendedBy property
    offerWhileRecommendedBy []string
    // The safeguard property
    safeguard SafeguardSettingsable
}
// NewContentApplicabilitySettings instantiates a new contentApplicabilitySettings and sets the default values.
func NewContentApplicabilitySettings()(*ContentApplicabilitySettings) {
    m := &ContentApplicabilitySettings{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateContentApplicabilitySettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateContentApplicabilitySettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewContentApplicabilitySettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ContentApplicabilitySettings) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ContentApplicabilitySettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["offerWhileRecommendedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetOfferWhileRecommendedBy(res)
        }
        return nil
    }
    res["safeguard"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSafeguardSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSafeguard(val.(SafeguardSettingsable))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ContentApplicabilitySettings) GetOdataType()(*string) {
    return m.odataType
}
// GetOfferWhileRecommendedBy gets the offerWhileRecommendedBy property value. The offerWhileRecommendedBy property
func (m *ContentApplicabilitySettings) GetOfferWhileRecommendedBy()([]string) {
    return m.offerWhileRecommendedBy
}
// GetSafeguard gets the safeguard property value. The safeguard property
func (m *ContentApplicabilitySettings) GetSafeguard()(SafeguardSettingsable) {
    return m.safeguard
}
// Serialize serializes information the current object
func (m *ContentApplicabilitySettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetOfferWhileRecommendedBy() != nil {
        err := writer.WriteCollectionOfStringValues("offerWhileRecommendedBy", m.GetOfferWhileRecommendedBy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("safeguard", m.GetSafeguard())
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
func (m *ContentApplicabilitySettings) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ContentApplicabilitySettings) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOfferWhileRecommendedBy sets the offerWhileRecommendedBy property value. The offerWhileRecommendedBy property
func (m *ContentApplicabilitySettings) SetOfferWhileRecommendedBy(value []string)() {
    m.offerWhileRecommendedBy = value
}
// SetSafeguard sets the safeguard property value. The safeguard property
func (m *ContentApplicabilitySettings) SetSafeguard(value SafeguardSettingsable)() {
    m.safeguard = value
}
