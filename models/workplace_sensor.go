package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type WorkplaceSensor struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewWorkplaceSensor instantiates a new WorkplaceSensor and sets the default values.
func NewWorkplaceSensor()(*WorkplaceSensor) {
    m := &WorkplaceSensor{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWorkplaceSensorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWorkplaceSensorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkplaceSensor(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *WorkplaceSensor) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *WorkplaceSensor) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDisplayName gets the displayName property value. The display name of the sensor.
// returns a *string when successful
func (m *WorkplaceSensor) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WorkplaceSensor) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
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
    res["placeId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlaceId(val)
        }
        return nil
    }
    res["sensorId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSensorId(val)
        }
        return nil
    }
    res["sensorType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWorkplaceSensorType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSensorType(val.(*WorkplaceSensorType))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *WorkplaceSensor) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPlaceId gets the placeId property value. The unique identifier of the place served by the sensor. Only needs to be provided if the sensor serves a different place than the device's location.
// returns a *string when successful
func (m *WorkplaceSensor) GetPlaceId()(*string) {
    val, err := m.GetBackingStore().Get("placeId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSensorId gets the sensorId property value. The unique identifier of a sensor within the device. If the sensor Id is not provided, the sensorType will be used as sensorId.
// returns a *string when successful
func (m *WorkplaceSensor) GetSensorId()(*string) {
    val, err := m.GetBackingStore().Get("sensorId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSensorType gets the sensorType property value. The sensorType property
// returns a *WorkplaceSensorType when successful
func (m *WorkplaceSensor) GetSensorType()(*WorkplaceSensorType) {
    val, err := m.GetBackingStore().Get("sensorType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*WorkplaceSensorType)
    }
    return nil
}
// Serialize serializes information the current object
func (m *WorkplaceSensor) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
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
        err := writer.WriteStringValue("placeId", m.GetPlaceId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sensorId", m.GetSensorId())
        if err != nil {
            return err
        }
    }
    if m.GetSensorType() != nil {
        cast := (*m.GetSensorType()).String()
        err := writer.WriteStringValue("sensorType", &cast)
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WorkplaceSensor) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *WorkplaceSensor) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDisplayName sets the displayName property value. The display name of the sensor.
func (m *WorkplaceSensor) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *WorkplaceSensor) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetPlaceId sets the placeId property value. The unique identifier of the place served by the sensor. Only needs to be provided if the sensor serves a different place than the device's location.
func (m *WorkplaceSensor) SetPlaceId(value *string)() {
    err := m.GetBackingStore().Set("placeId", value)
    if err != nil {
        panic(err)
    }
}
// SetSensorId sets the sensorId property value. The unique identifier of a sensor within the device. If the sensor Id is not provided, the sensorType will be used as sensorId.
func (m *WorkplaceSensor) SetSensorId(value *string)() {
    err := m.GetBackingStore().Set("sensorId", value)
    if err != nil {
        panic(err)
    }
}
// SetSensorType sets the sensorType property value. The sensorType property
func (m *WorkplaceSensor) SetSensorType(value *WorkplaceSensorType)() {
    err := m.GetBackingStore().Set("sensorType", value)
    if err != nil {
        panic(err)
    }
}
type WorkplaceSensorable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDisplayName()(*string)
    GetOdataType()(*string)
    GetPlaceId()(*string)
    GetSensorId()(*string)
    GetSensorType()(*WorkplaceSensorType)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDisplayName(value *string)()
    SetOdataType(value *string)()
    SetPlaceId(value *string)()
    SetSensorId(value *string)()
    SetSensorType(value *WorkplaceSensorType)()
}
