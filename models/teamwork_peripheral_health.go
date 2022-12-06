package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamworkPeripheralHealth 
type TeamworkPeripheralHealth struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The connected state and time since the peripheral device was connected.
    connection TeamworkConnectionable
    // True if the peripheral is optional. Used for health computation.
    isOptional *bool
    // The OdataType property
    odataType *string
    // The peripheral property
    peripheral TeamworkPeripheralable
}
// NewTeamworkPeripheralHealth instantiates a new teamworkPeripheralHealth and sets the default values.
func NewTeamworkPeripheralHealth()(*TeamworkPeripheralHealth) {
    m := &TeamworkPeripheralHealth{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTeamworkPeripheralHealthFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamworkPeripheralHealthFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamworkPeripheralHealth(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkPeripheralHealth) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetConnection gets the connection property value. The connected state and time since the peripheral device was connected.
func (m *TeamworkPeripheralHealth) GetConnection()(TeamworkConnectionable) {
    return m.connection
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkPeripheralHealth) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["connection"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateTeamworkConnectionFromDiscriminatorValue , m.SetConnection)
    res["isOptional"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsOptional)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["peripheral"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateTeamworkPeripheralFromDiscriminatorValue , m.SetPeripheral)
    return res
}
// GetIsOptional gets the isOptional property value. True if the peripheral is optional. Used for health computation.
func (m *TeamworkPeripheralHealth) GetIsOptional()(*bool) {
    return m.isOptional
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TeamworkPeripheralHealth) GetOdataType()(*string) {
    return m.odataType
}
// GetPeripheral gets the peripheral property value. The peripheral property
func (m *TeamworkPeripheralHealth) GetPeripheral()(TeamworkPeripheralable) {
    return m.peripheral
}
// Serialize serializes information the current object
func (m *TeamworkPeripheralHealth) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("connection", m.GetConnection())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isOptional", m.GetIsOptional())
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
        err := writer.WriteObjectValue("peripheral", m.GetPeripheral())
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
func (m *TeamworkPeripheralHealth) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetConnection sets the connection property value. The connected state and time since the peripheral device was connected.
func (m *TeamworkPeripheralHealth) SetConnection(value TeamworkConnectionable)() {
    m.connection = value
}
// SetIsOptional sets the isOptional property value. True if the peripheral is optional. Used for health computation.
func (m *TeamworkPeripheralHealth) SetIsOptional(value *bool)() {
    m.isOptional = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TeamworkPeripheralHealth) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPeripheral sets the peripheral property value. The peripheral property
func (m *TeamworkPeripheralHealth) SetPeripheral(value TeamworkPeripheralable)() {
    m.peripheral = value
}
