package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OnPremisesApplicationSegment 
type OnPremisesApplicationSegment struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // If you're configuring a traffic manager in front of multiple App Proxy application segments, contains the user-friendly URL that will point to the traffic manager.
    alternateUrl *string
    // CORS Rule definition for a particular application segment.
    corsConfigurations []CorsConfigurationable
    // The published external URL for the application segment; for example, https://intranet.contoso.com./
    externalUrl *string
    // The internal URL of the application segment; for example, https://intranet/.
    internalUrl *string
    // The OdataType property
    odataType *string
}
// NewOnPremisesApplicationSegment instantiates a new onPremisesApplicationSegment and sets the default values.
func NewOnPremisesApplicationSegment()(*OnPremisesApplicationSegment) {
    m := &OnPremisesApplicationSegment{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateOnPremisesApplicationSegmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOnPremisesApplicationSegmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOnPremisesApplicationSegment(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OnPremisesApplicationSegment) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAlternateUrl gets the alternateUrl property value. If you're configuring a traffic manager in front of multiple App Proxy application segments, contains the user-friendly URL that will point to the traffic manager.
func (m *OnPremisesApplicationSegment) GetAlternateUrl()(*string) {
    return m.alternateUrl
}
// GetCorsConfigurations gets the corsConfigurations property value. CORS Rule definition for a particular application segment.
func (m *OnPremisesApplicationSegment) GetCorsConfigurations()([]CorsConfigurationable) {
    return m.corsConfigurations
}
// GetExternalUrl gets the externalUrl property value. The published external URL for the application segment; for example, https://intranet.contoso.com./
func (m *OnPremisesApplicationSegment) GetExternalUrl()(*string) {
    return m.externalUrl
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OnPremisesApplicationSegment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["alternateUrl"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAlternateUrl)
    res["corsConfigurations"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateCorsConfigurationFromDiscriminatorValue , m.SetCorsConfigurations)
    res["externalUrl"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetExternalUrl)
    res["internalUrl"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetInternalUrl)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    return res
}
// GetInternalUrl gets the internalUrl property value. The internal URL of the application segment; for example, https://intranet/.
func (m *OnPremisesApplicationSegment) GetInternalUrl()(*string) {
    return m.internalUrl
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *OnPremisesApplicationSegment) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *OnPremisesApplicationSegment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("alternateUrl", m.GetAlternateUrl())
        if err != nil {
            return err
        }
    }
    if m.GetCorsConfigurations() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetCorsConfigurations())
        err := writer.WriteCollectionOfObjectValues("corsConfigurations", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("externalUrl", m.GetExternalUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("internalUrl", m.GetInternalUrl())
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
func (m *OnPremisesApplicationSegment) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAlternateUrl sets the alternateUrl property value. If you're configuring a traffic manager in front of multiple App Proxy application segments, contains the user-friendly URL that will point to the traffic manager.
func (m *OnPremisesApplicationSegment) SetAlternateUrl(value *string)() {
    m.alternateUrl = value
}
// SetCorsConfigurations sets the corsConfigurations property value. CORS Rule definition for a particular application segment.
func (m *OnPremisesApplicationSegment) SetCorsConfigurations(value []CorsConfigurationable)() {
    m.corsConfigurations = value
}
// SetExternalUrl sets the externalUrl property value. The published external URL for the application segment; for example, https://intranet.contoso.com./
func (m *OnPremisesApplicationSegment) SetExternalUrl(value *string)() {
    m.externalUrl = value
}
// SetInternalUrl sets the internalUrl property value. The internal URL of the application segment; for example, https://intranet/.
func (m *OnPremisesApplicationSegment) SetInternalUrl(value *string)() {
    m.internalUrl = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *OnPremisesApplicationSegment) SetOdataType(value *string)() {
    m.odataType = value
}
