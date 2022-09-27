package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OathTokenMetadata 
type OathTokenMetadata struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The enabled property
    enabled *bool
    // The manufacturer property
    manufacturer *string
    // The manufacturerProperties property
    manufacturerProperties []KeyValueable
    // The OdataType property
    odataType *string
    // The serialNumber property
    serialNumber *string
    // The tokenType property
    tokenType *string
}
// NewOathTokenMetadata instantiates a new oathTokenMetadata and sets the default values.
func NewOathTokenMetadata()(*OathTokenMetadata) {
    m := &OathTokenMetadata{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.oathTokenMetadata";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateOathTokenMetadataFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOathTokenMetadataFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOathTokenMetadata(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OathTokenMetadata) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetEnabled gets the enabled property value. The enabled property
func (m *OathTokenMetadata) GetEnabled()(*bool) {
    return m.enabled
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OathTokenMetadata) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["enabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetEnabled)
    res["manufacturer"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetManufacturer)
    res["manufacturerProperties"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateKeyValueFromDiscriminatorValue , m.SetManufacturerProperties)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["serialNumber"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSerialNumber)
    res["tokenType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetTokenType)
    return res
}
// GetManufacturer gets the manufacturer property value. The manufacturer property
func (m *OathTokenMetadata) GetManufacturer()(*string) {
    return m.manufacturer
}
// GetManufacturerProperties gets the manufacturerProperties property value. The manufacturerProperties property
func (m *OathTokenMetadata) GetManufacturerProperties()([]KeyValueable) {
    return m.manufacturerProperties
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *OathTokenMetadata) GetOdataType()(*string) {
    return m.odataType
}
// GetSerialNumber gets the serialNumber property value. The serialNumber property
func (m *OathTokenMetadata) GetSerialNumber()(*string) {
    return m.serialNumber
}
// GetTokenType gets the tokenType property value. The tokenType property
func (m *OathTokenMetadata) GetTokenType()(*string) {
    return m.tokenType
}
// Serialize serializes information the current object
func (m *OathTokenMetadata) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("enabled", m.GetEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("manufacturer", m.GetManufacturer())
        if err != nil {
            return err
        }
    }
    if m.GetManufacturerProperties() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetManufacturerProperties())
        err := writer.WriteCollectionOfObjectValues("manufacturerProperties", cast)
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
        err := writer.WriteStringValue("serialNumber", m.GetSerialNumber())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("tokenType", m.GetTokenType())
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
func (m *OathTokenMetadata) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetEnabled sets the enabled property value. The enabled property
func (m *OathTokenMetadata) SetEnabled(value *bool)() {
    m.enabled = value
}
// SetManufacturer sets the manufacturer property value. The manufacturer property
func (m *OathTokenMetadata) SetManufacturer(value *string)() {
    m.manufacturer = value
}
// SetManufacturerProperties sets the manufacturerProperties property value. The manufacturerProperties property
func (m *OathTokenMetadata) SetManufacturerProperties(value []KeyValueable)() {
    m.manufacturerProperties = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *OathTokenMetadata) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSerialNumber sets the serialNumber property value. The serialNumber property
func (m *OathTokenMetadata) SetSerialNumber(value *string)() {
    m.serialNumber = value
}
// SetTokenType sets the tokenType property value. The tokenType property
func (m *OathTokenMetadata) SetTokenType(value *string)() {
    m.tokenType = value
}
