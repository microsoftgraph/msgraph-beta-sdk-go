package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamworkHardwareDetail 
type TeamworkHardwareDetail struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // MAC address.
    macAddresses []string
    // Device manufacturer.
    manufacturer *string
    // Devie model.
    model *string
    // The OdataType property
    odataType *string
    // Device serial number.
    serialNumber *string
    // The unique identifier for the device.
    uniqueId *string
}
// NewTeamworkHardwareDetail instantiates a new teamworkHardwareDetail and sets the default values.
func NewTeamworkHardwareDetail()(*TeamworkHardwareDetail) {
    m := &TeamworkHardwareDetail{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.teamworkHardwareDetail";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateTeamworkHardwareDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamworkHardwareDetailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamworkHardwareDetail(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkHardwareDetail) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkHardwareDetail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["macAddresses"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetMacAddresses)
    res["manufacturer"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetManufacturer)
    res["model"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetModel)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["serialNumber"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSerialNumber)
    res["uniqueId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUniqueId)
    return res
}
// GetMacAddresses gets the macAddresses property value. MAC address.
func (m *TeamworkHardwareDetail) GetMacAddresses()([]string) {
    return m.macAddresses
}
// GetManufacturer gets the manufacturer property value. Device manufacturer.
func (m *TeamworkHardwareDetail) GetManufacturer()(*string) {
    return m.manufacturer
}
// GetModel gets the model property value. Devie model.
func (m *TeamworkHardwareDetail) GetModel()(*string) {
    return m.model
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TeamworkHardwareDetail) GetOdataType()(*string) {
    return m.odataType
}
// GetSerialNumber gets the serialNumber property value. Device serial number.
func (m *TeamworkHardwareDetail) GetSerialNumber()(*string) {
    return m.serialNumber
}
// GetUniqueId gets the uniqueId property value. The unique identifier for the device.
func (m *TeamworkHardwareDetail) GetUniqueId()(*string) {
    return m.uniqueId
}
// Serialize serializes information the current object
func (m *TeamworkHardwareDetail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetMacAddresses() != nil {
        err := writer.WriteCollectionOfStringValues("macAddresses", m.GetMacAddresses())
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
    {
        err := writer.WriteStringValue("model", m.GetModel())
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
        err := writer.WriteStringValue("uniqueId", m.GetUniqueId())
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
func (m *TeamworkHardwareDetail) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetMacAddresses sets the macAddresses property value. MAC address.
func (m *TeamworkHardwareDetail) SetMacAddresses(value []string)() {
    m.macAddresses = value
}
// SetManufacturer sets the manufacturer property value. Device manufacturer.
func (m *TeamworkHardwareDetail) SetManufacturer(value *string)() {
    m.manufacturer = value
}
// SetModel sets the model property value. Devie model.
func (m *TeamworkHardwareDetail) SetModel(value *string)() {
    m.model = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TeamworkHardwareDetail) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSerialNumber sets the serialNumber property value. Device serial number.
func (m *TeamworkHardwareDetail) SetSerialNumber(value *string)() {
    m.serialNumber = value
}
// SetUniqueId sets the uniqueId property value. The unique identifier for the device.
func (m *TeamworkHardwareDetail) SetUniqueId(value *string)() {
    m.uniqueId = value
}
