package managedtenants

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type DeviceHealthStatus struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// DeviceHealthStatus_DeviceHealthStatus_bootTotalDurationInSeconds composed type wrapper for classes float64, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReferenceNumeric, string
type DeviceHealthStatus_DeviceHealthStatus_bootTotalDurationInSeconds struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewDeviceHealthStatus_DeviceHealthStatus_bootTotalDurationInSeconds instantiates a new DeviceHealthStatus_DeviceHealthStatus_bootTotalDurationInSeconds and sets the default values.
func NewDeviceHealthStatus_DeviceHealthStatus_bootTotalDurationInSeconds()(*DeviceHealthStatus_DeviceHealthStatus_bootTotalDurationInSeconds) {
    m := &DeviceHealthStatus_DeviceHealthStatus_bootTotalDurationInSeconds{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateDeviceHealthStatus_DeviceHealthStatus_bootTotalDurationInSecondsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeviceHealthStatus_DeviceHealthStatus_bootTotalDurationInSecondsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewDeviceHealthStatus_DeviceHealthStatus_bootTotalDurationInSeconds()
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
            }
        }
    }
    if val, err := parseNode.GetEnumValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ParseReferenceNumeric); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetReferenceNumeric(val)
    } else if val, err := parseNode.GetFloat64Value(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetDouble(val)
    } else if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetString(val)
    }
    return result, nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *DeviceHealthStatus_DeviceHealthStatus_bootTotalDurationInSeconds) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *DeviceHealthStatus_DeviceHealthStatus_bootTotalDurationInSeconds) GetDouble()(*float64) {
    val, err := m.GetBackingStore().Get("double")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DeviceHealthStatus_DeviceHealthStatus_bootTotalDurationInSeconds) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *DeviceHealthStatus_DeviceHealthStatus_bootTotalDurationInSeconds) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *DeviceHealthStatus_DeviceHealthStatus_bootTotalDurationInSeconds) GetReferenceNumeric()(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReferenceNumeric) {
    val, err := m.GetBackingStore().Get("referenceNumeric")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReferenceNumeric)
    }
    return nil
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *DeviceHealthStatus_DeviceHealthStatus_bootTotalDurationInSeconds) GetString()(*string) {
    val, err := m.GetBackingStore().Get("string")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeviceHealthStatus_DeviceHealthStatus_bootTotalDurationInSeconds) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetReferenceNumeric() != nil {
        cast := (*m.GetReferenceNumeric()).String()
        err := writer.WriteStringValue("", &cast)
        if err != nil {
            return err
        }
    } else if m.GetDouble() != nil {
        err := writer.WriteFloat64Value("", m.GetDouble())
        if err != nil {
            return err
        }
    } else if m.GetString() != nil {
        err := writer.WriteStringValue("", m.GetString())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *DeviceHealthStatus_DeviceHealthStatus_bootTotalDurationInSeconds) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *DeviceHealthStatus_DeviceHealthStatus_bootTotalDurationInSeconds) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReferenceNumeric
func (m *DeviceHealthStatus_DeviceHealthStatus_bootTotalDurationInSeconds) SetReferenceNumeric(value *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *DeviceHealthStatus_DeviceHealthStatus_bootTotalDurationInSeconds) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// DeviceHealthStatus_DeviceHealthStatus_startupPerformanceScore composed type wrapper for classes float64, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReferenceNumeric, string
type DeviceHealthStatus_DeviceHealthStatus_startupPerformanceScore struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewDeviceHealthStatus_DeviceHealthStatus_startupPerformanceScore instantiates a new DeviceHealthStatus_DeviceHealthStatus_startupPerformanceScore and sets the default values.
func NewDeviceHealthStatus_DeviceHealthStatus_startupPerformanceScore()(*DeviceHealthStatus_DeviceHealthStatus_startupPerformanceScore) {
    m := &DeviceHealthStatus_DeviceHealthStatus_startupPerformanceScore{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateDeviceHealthStatus_DeviceHealthStatus_startupPerformanceScoreFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeviceHealthStatus_DeviceHealthStatus_startupPerformanceScoreFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewDeviceHealthStatus_DeviceHealthStatus_startupPerformanceScore()
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
            }
        }
    }
    if val, err := parseNode.GetEnumValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ParseReferenceNumeric); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetReferenceNumeric(val)
    } else if val, err := parseNode.GetFloat64Value(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetDouble(val)
    } else if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetString(val)
    }
    return result, nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *DeviceHealthStatus_DeviceHealthStatus_startupPerformanceScore) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *DeviceHealthStatus_DeviceHealthStatus_startupPerformanceScore) GetDouble()(*float64) {
    val, err := m.GetBackingStore().Get("double")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DeviceHealthStatus_DeviceHealthStatus_startupPerformanceScore) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *DeviceHealthStatus_DeviceHealthStatus_startupPerformanceScore) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *DeviceHealthStatus_DeviceHealthStatus_startupPerformanceScore) GetReferenceNumeric()(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReferenceNumeric) {
    val, err := m.GetBackingStore().Get("referenceNumeric")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReferenceNumeric)
    }
    return nil
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *DeviceHealthStatus_DeviceHealthStatus_startupPerformanceScore) GetString()(*string) {
    val, err := m.GetBackingStore().Get("string")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeviceHealthStatus_DeviceHealthStatus_startupPerformanceScore) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetReferenceNumeric() != nil {
        cast := (*m.GetReferenceNumeric()).String()
        err := writer.WriteStringValue("", &cast)
        if err != nil {
            return err
        }
    } else if m.GetDouble() != nil {
        err := writer.WriteFloat64Value("", m.GetDouble())
        if err != nil {
            return err
        }
    } else if m.GetString() != nil {
        err := writer.WriteStringValue("", m.GetString())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *DeviceHealthStatus_DeviceHealthStatus_startupPerformanceScore) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *DeviceHealthStatus_DeviceHealthStatus_startupPerformanceScore) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReferenceNumeric
func (m *DeviceHealthStatus_DeviceHealthStatus_startupPerformanceScore) SetReferenceNumeric(value *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *DeviceHealthStatus_DeviceHealthStatus_startupPerformanceScore) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
type DeviceHealthStatus_DeviceHealthStatus_bootTotalDurationInSecondsable interface {
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDouble()(*float64)
    GetReferenceNumeric()(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReferenceNumeric)
    GetString()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDouble(value *float64)()
    SetReferenceNumeric(value *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReferenceNumeric)()
    SetString(value *string)()
}
type DeviceHealthStatus_DeviceHealthStatus_startupPerformanceScoreable interface {
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDouble()(*float64)
    GetReferenceNumeric()(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReferenceNumeric)
    GetString()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDouble(value *float64)()
    SetReferenceNumeric(value *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReferenceNumeric)()
    SetString(value *string)()
}
// NewDeviceHealthStatus instantiates a new DeviceHealthStatus and sets the default values.
func NewDeviceHealthStatus()(*DeviceHealthStatus) {
    m := &DeviceHealthStatus{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateDeviceHealthStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeviceHealthStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceHealthStatus(), nil
}
// GetBlueScreenCount gets the blueScreenCount property value. The blueScreenCount property
// returns a *int32 when successful
func (m *DeviceHealthStatus) GetBlueScreenCount()(*int32) {
    val, err := m.GetBackingStore().Get("blueScreenCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetBootTotalDurationInSeconds gets the bootTotalDurationInSeconds property value. The bootTotalDurationInSeconds property
// returns a DeviceHealthStatus_DeviceHealthStatus_bootTotalDurationInSecondsable when successful
func (m *DeviceHealthStatus) GetBootTotalDurationInSeconds()(DeviceHealthStatus_DeviceHealthStatus_bootTotalDurationInSecondsable) {
    val, err := m.GetBackingStore().Get("bootTotalDurationInSeconds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceHealthStatus_DeviceHealthStatus_bootTotalDurationInSecondsable)
    }
    return nil
}
// GetDeviceId gets the deviceId property value. The deviceId property
// returns a *string when successful
func (m *DeviceHealthStatus) GetDeviceId()(*string) {
    val, err := m.GetBackingStore().Get("deviceId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceMake gets the deviceMake property value. The deviceMake property
// returns a *string when successful
func (m *DeviceHealthStatus) GetDeviceMake()(*string) {
    val, err := m.GetBackingStore().Get("deviceMake")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceModel gets the deviceModel property value. The deviceModel property
// returns a *string when successful
func (m *DeviceHealthStatus) GetDeviceModel()(*string) {
    val, err := m.GetBackingStore().Get("deviceModel")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceName gets the deviceName property value. The deviceName property
// returns a *string when successful
func (m *DeviceHealthStatus) GetDeviceName()(*string) {
    val, err := m.GetBackingStore().Get("deviceName")
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
func (m *DeviceHealthStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["blueScreenCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlueScreenCount(val)
        }
        return nil
    }
    res["bootTotalDurationInSeconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceHealthStatus_DeviceHealthStatus_bootTotalDurationInSecondsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBootTotalDurationInSeconds(val.(*DeviceHealthStatus_DeviceHealthStatus_bootTotalDurationInSeconds))
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
    res["deviceMake"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceMake(val)
        }
        return nil
    }
    res["deviceModel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceModel(val)
        }
        return nil
    }
    res["deviceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceName(val)
        }
        return nil
    }
    res["healthStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHealthStatus(val)
        }
        return nil
    }
    res["lastUpdatedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastUpdatedDateTime(val)
        }
        return nil
    }
    res["osVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsVersion(val)
        }
        return nil
    }
    res["primaryDiskType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrimaryDiskType(val)
        }
        return nil
    }
    res["restartCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRestartCount(val)
        }
        return nil
    }
    res["startupPerformanceScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceHealthStatus_DeviceHealthStatus_startupPerformanceScoreFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartupPerformanceScore(val.(*DeviceHealthStatus_DeviceHealthStatus_startupPerformanceScore))
        }
        return nil
    }
    res["tenantDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantDisplayName(val)
        }
        return nil
    }
    res["tenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantId(val)
        }
        return nil
    }
    res["topProcesses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTopProcesses(val)
        }
        return nil
    }
    return res
}
// GetHealthStatus gets the healthStatus property value. The healthStatus property
// returns a *string when successful
func (m *DeviceHealthStatus) GetHealthStatus()(*string) {
    val, err := m.GetBackingStore().Get("healthStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetLastUpdatedDateTime gets the lastUpdatedDateTime property value. The lastUpdatedDateTime property
// returns a *Time when successful
func (m *DeviceHealthStatus) GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastUpdatedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetOsVersion gets the osVersion property value. The osVersion property
// returns a *string when successful
func (m *DeviceHealthStatus) GetOsVersion()(*string) {
    val, err := m.GetBackingStore().Get("osVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPrimaryDiskType gets the primaryDiskType property value. The primaryDiskType property
// returns a *string when successful
func (m *DeviceHealthStatus) GetPrimaryDiskType()(*string) {
    val, err := m.GetBackingStore().Get("primaryDiskType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRestartCount gets the restartCount property value. The restartCount property
// returns a *int32 when successful
func (m *DeviceHealthStatus) GetRestartCount()(*int32) {
    val, err := m.GetBackingStore().Get("restartCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetStartupPerformanceScore gets the startupPerformanceScore property value. The startupPerformanceScore property
// returns a DeviceHealthStatus_DeviceHealthStatus_startupPerformanceScoreable when successful
func (m *DeviceHealthStatus) GetStartupPerformanceScore()(DeviceHealthStatus_DeviceHealthStatus_startupPerformanceScoreable) {
    val, err := m.GetBackingStore().Get("startupPerformanceScore")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceHealthStatus_DeviceHealthStatus_startupPerformanceScoreable)
    }
    return nil
}
// GetTenantDisplayName gets the tenantDisplayName property value. The tenantDisplayName property
// returns a *string when successful
func (m *DeviceHealthStatus) GetTenantDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("tenantDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTenantId gets the tenantId property value. The tenantId property
// returns a *string when successful
func (m *DeviceHealthStatus) GetTenantId()(*string) {
    val, err := m.GetBackingStore().Get("tenantId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTopProcesses gets the topProcesses property value. The topProcesses property
// returns a *string when successful
func (m *DeviceHealthStatus) GetTopProcesses()(*string) {
    val, err := m.GetBackingStore().Get("topProcesses")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeviceHealthStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("blueScreenCount", m.GetBlueScreenCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("bootTotalDurationInSeconds", m.GetBootTotalDurationInSeconds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceId", m.GetDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceMake", m.GetDeviceMake())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceModel", m.GetDeviceModel())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceName", m.GetDeviceName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("healthStatus", m.GetHealthStatus())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastUpdatedDateTime", m.GetLastUpdatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("osVersion", m.GetOsVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("primaryDiskType", m.GetPrimaryDiskType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("restartCount", m.GetRestartCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("startupPerformanceScore", m.GetStartupPerformanceScore())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tenantDisplayName", m.GetTenantDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tenantId", m.GetTenantId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("topProcesses", m.GetTopProcesses())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBlueScreenCount sets the blueScreenCount property value. The blueScreenCount property
func (m *DeviceHealthStatus) SetBlueScreenCount(value *int32)() {
    err := m.GetBackingStore().Set("blueScreenCount", value)
    if err != nil {
        panic(err)
    }
}
// SetBootTotalDurationInSeconds sets the bootTotalDurationInSeconds property value. The bootTotalDurationInSeconds property
func (m *DeviceHealthStatus) SetBootTotalDurationInSeconds(value DeviceHealthStatus_DeviceHealthStatus_bootTotalDurationInSecondsable)() {
    err := m.GetBackingStore().Set("bootTotalDurationInSeconds", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceId sets the deviceId property value. The deviceId property
func (m *DeviceHealthStatus) SetDeviceId(value *string)() {
    err := m.GetBackingStore().Set("deviceId", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceMake sets the deviceMake property value. The deviceMake property
func (m *DeviceHealthStatus) SetDeviceMake(value *string)() {
    err := m.GetBackingStore().Set("deviceMake", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceModel sets the deviceModel property value. The deviceModel property
func (m *DeviceHealthStatus) SetDeviceModel(value *string)() {
    err := m.GetBackingStore().Set("deviceModel", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceName sets the deviceName property value. The deviceName property
func (m *DeviceHealthStatus) SetDeviceName(value *string)() {
    err := m.GetBackingStore().Set("deviceName", value)
    if err != nil {
        panic(err)
    }
}
// SetHealthStatus sets the healthStatus property value. The healthStatus property
func (m *DeviceHealthStatus) SetHealthStatus(value *string)() {
    err := m.GetBackingStore().Set("healthStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetLastUpdatedDateTime sets the lastUpdatedDateTime property value. The lastUpdatedDateTime property
func (m *DeviceHealthStatus) SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastUpdatedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetOsVersion sets the osVersion property value. The osVersion property
func (m *DeviceHealthStatus) SetOsVersion(value *string)() {
    err := m.GetBackingStore().Set("osVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetPrimaryDiskType sets the primaryDiskType property value. The primaryDiskType property
func (m *DeviceHealthStatus) SetPrimaryDiskType(value *string)() {
    err := m.GetBackingStore().Set("primaryDiskType", value)
    if err != nil {
        panic(err)
    }
}
// SetRestartCount sets the restartCount property value. The restartCount property
func (m *DeviceHealthStatus) SetRestartCount(value *int32)() {
    err := m.GetBackingStore().Set("restartCount", value)
    if err != nil {
        panic(err)
    }
}
// SetStartupPerformanceScore sets the startupPerformanceScore property value. The startupPerformanceScore property
func (m *DeviceHealthStatus) SetStartupPerformanceScore(value DeviceHealthStatus_DeviceHealthStatus_startupPerformanceScoreable)() {
    err := m.GetBackingStore().Set("startupPerformanceScore", value)
    if err != nil {
        panic(err)
    }
}
// SetTenantDisplayName sets the tenantDisplayName property value. The tenantDisplayName property
func (m *DeviceHealthStatus) SetTenantDisplayName(value *string)() {
    err := m.GetBackingStore().Set("tenantDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetTenantId sets the tenantId property value. The tenantId property
func (m *DeviceHealthStatus) SetTenantId(value *string)() {
    err := m.GetBackingStore().Set("tenantId", value)
    if err != nil {
        panic(err)
    }
}
// SetTopProcesses sets the topProcesses property value. The topProcesses property
func (m *DeviceHealthStatus) SetTopProcesses(value *string)() {
    err := m.GetBackingStore().Set("topProcesses", value)
    if err != nil {
        panic(err)
    }
}
type DeviceHealthStatusable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBlueScreenCount()(*int32)
    GetBootTotalDurationInSeconds()(DeviceHealthStatus_DeviceHealthStatus_bootTotalDurationInSecondsable)
    GetDeviceId()(*string)
    GetDeviceMake()(*string)
    GetDeviceModel()(*string)
    GetDeviceName()(*string)
    GetHealthStatus()(*string)
    GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetOsVersion()(*string)
    GetPrimaryDiskType()(*string)
    GetRestartCount()(*int32)
    GetStartupPerformanceScore()(DeviceHealthStatus_DeviceHealthStatus_startupPerformanceScoreable)
    GetTenantDisplayName()(*string)
    GetTenantId()(*string)
    GetTopProcesses()(*string)
    SetBlueScreenCount(value *int32)()
    SetBootTotalDurationInSeconds(value DeviceHealthStatus_DeviceHealthStatus_bootTotalDurationInSecondsable)()
    SetDeviceId(value *string)()
    SetDeviceMake(value *string)()
    SetDeviceModel(value *string)()
    SetDeviceName(value *string)()
    SetHealthStatus(value *string)()
    SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetOsVersion(value *string)()
    SetPrimaryDiskType(value *string)()
    SetRestartCount(value *int32)()
    SetStartupPerformanceScore(value DeviceHealthStatus_DeviceHealthStatus_startupPerformanceScoreable)()
    SetTenantDisplayName(value *string)()
    SetTenantId(value *string)()
    SetTopProcesses(value *string)()
}
