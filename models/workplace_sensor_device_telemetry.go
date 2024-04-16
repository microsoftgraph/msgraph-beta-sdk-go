package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type WorkplaceSensorDeviceTelemetry struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewWorkplaceSensorDeviceTelemetry instantiates a new WorkplaceSensorDeviceTelemetry and sets the default values.
func NewWorkplaceSensorDeviceTelemetry()(*WorkplaceSensorDeviceTelemetry) {
    m := &WorkplaceSensorDeviceTelemetry{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWorkplaceSensorDeviceTelemetryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWorkplaceSensorDeviceTelemetryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkplaceSensorDeviceTelemetry(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *WorkplaceSensorDeviceTelemetry) GetAdditionalData()(map[string]any) {
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
func (m *WorkplaceSensorDeviceTelemetry) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetBoolValue gets the boolValue property value. The value of the sensor can be true or false. Use it for sensors that report binary values, such as occupancy or heartbeat.
// returns a *bool when successful
func (m *WorkplaceSensorDeviceTelemetry) GetBoolValue()(*bool) {
    val, err := m.GetBackingStore().Get("boolValue")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDeviceId gets the deviceId property value. The user-defined unique identifier of the device provided at the time of creation. Don't use the system generated identifier of the device.
// returns a *string when successful
func (m *WorkplaceSensorDeviceTelemetry) GetDeviceId()(*string) {
    val, err := m.GetBackingStore().Get("deviceId")
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
func (m *WorkplaceSensorDeviceTelemetry) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["boolValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBoolValue(val)
        }
        return nil
    }
    res["deviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceId(val)
        }
        return nil
    }
    res["intValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIntValue(val)
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
    res["timestamp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimestamp(val)
        }
        return nil
    }
    return res
}
// GetIntValue gets the intValue property value. The value of the sensor as an integer. Use it for sensors that report numerical values, such as people count.
// returns a *int32 when successful
func (m *WorkplaceSensorDeviceTelemetry) GetIntValue()(*int32) {
    val, err := m.GetBackingStore().Get("intValue")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *WorkplaceSensorDeviceTelemetry) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSensorId gets the sensorId property value. The user-defined unique identifier of the sensor on the device. Optional. If the device has multiple sensors of the same type, the property must be provided to identify each sensor. If the device has unique sensor types, the property can be omitted. The default value is the sensor type.
// returns a *string when successful
func (m *WorkplaceSensorDeviceTelemetry) GetSensorId()(*string) {
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
func (m *WorkplaceSensorDeviceTelemetry) GetSensorType()(*WorkplaceSensorType) {
    val, err := m.GetBackingStore().Get("sensorType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*WorkplaceSensorType)
    }
    return nil
}
// GetTimestamp gets the timestamp property value. The date and time when the sensor measured and reported its value. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *WorkplaceSensorDeviceTelemetry) GetTimestamp()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("timestamp")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// Serialize serializes information the current object
func (m *WorkplaceSensorDeviceTelemetry) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("boolValue", m.GetBoolValue())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("deviceId", m.GetDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("intValue", m.GetIntValue())
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
        err := writer.WriteTimeValue("timestamp", m.GetTimestamp())
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
func (m *WorkplaceSensorDeviceTelemetry) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *WorkplaceSensorDeviceTelemetry) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetBoolValue sets the boolValue property value. The value of the sensor can be true or false. Use it for sensors that report binary values, such as occupancy or heartbeat.
func (m *WorkplaceSensorDeviceTelemetry) SetBoolValue(value *bool)() {
    err := m.GetBackingStore().Set("boolValue", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceId sets the deviceId property value. The user-defined unique identifier of the device provided at the time of creation. Don't use the system generated identifier of the device.
func (m *WorkplaceSensorDeviceTelemetry) SetDeviceId(value *string)() {
    err := m.GetBackingStore().Set("deviceId", value)
    if err != nil {
        panic(err)
    }
}
// SetIntValue sets the intValue property value. The value of the sensor as an integer. Use it for sensors that report numerical values, such as people count.
func (m *WorkplaceSensorDeviceTelemetry) SetIntValue(value *int32)() {
    err := m.GetBackingStore().Set("intValue", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *WorkplaceSensorDeviceTelemetry) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetSensorId sets the sensorId property value. The user-defined unique identifier of the sensor on the device. Optional. If the device has multiple sensors of the same type, the property must be provided to identify each sensor. If the device has unique sensor types, the property can be omitted. The default value is the sensor type.
func (m *WorkplaceSensorDeviceTelemetry) SetSensorId(value *string)() {
    err := m.GetBackingStore().Set("sensorId", value)
    if err != nil {
        panic(err)
    }
}
// SetSensorType sets the sensorType property value. The sensorType property
func (m *WorkplaceSensorDeviceTelemetry) SetSensorType(value *WorkplaceSensorType)() {
    err := m.GetBackingStore().Set("sensorType", value)
    if err != nil {
        panic(err)
    }
}
// SetTimestamp sets the timestamp property value. The date and time when the sensor measured and reported its value. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *WorkplaceSensorDeviceTelemetry) SetTimestamp(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("timestamp", value)
    if err != nil {
        panic(err)
    }
}
type WorkplaceSensorDeviceTelemetryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetBoolValue()(*bool)
    GetDeviceId()(*string)
    GetIntValue()(*int32)
    GetOdataType()(*string)
    GetSensorId()(*string)
    GetSensorType()(*WorkplaceSensorType)
    GetTimestamp()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetBoolValue(value *bool)()
    SetDeviceId(value *string)()
    SetIntValue(value *int32)()
    SetOdataType(value *string)()
    SetSensorId(value *string)()
    SetSensorType(value *WorkplaceSensorType)()
    SetTimestamp(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
