package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// DeviceCompliancePoliciesSetScheduledRetireStatePostRequestBody 
type DeviceCompliancePoliciesSetScheduledRetireStatePostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewDeviceCompliancePoliciesSetScheduledRetireStatePostRequestBody instantiates a new DeviceCompliancePoliciesSetScheduledRetireStatePostRequestBody and sets the default values.
func NewDeviceCompliancePoliciesSetScheduledRetireStatePostRequestBody()(*DeviceCompliancePoliciesSetScheduledRetireStatePostRequestBody) {
    m := &DeviceCompliancePoliciesSetScheduledRetireStatePostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDeviceCompliancePoliciesSetScheduledRetireStatePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceCompliancePoliciesSetScheduledRetireStatePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceCompliancePoliciesSetScheduledRetireStatePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceCompliancePoliciesSetScheduledRetireStatePostRequestBody) GetAdditionalData()(map[string]any) {
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
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *DeviceCompliancePoliciesSetScheduledRetireStatePostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceCompliancePoliciesSetScheduledRetireStatePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["managedDeviceIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetManagedDeviceIds(res)
        }
        return nil
    }
    res["scopedToAllDevices"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScopedToAllDevices(val)
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ParseScheduledRetireState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ScheduledRetireState))
        }
        return nil
    }
    return res
}
// GetManagedDeviceIds gets the managedDeviceIds property value. The managedDeviceIds property
func (m *DeviceCompliancePoliciesSetScheduledRetireStatePostRequestBody) GetManagedDeviceIds()([]string) {
    val, err := m.GetBackingStore().Get("managedDeviceIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetScopedToAllDevices gets the scopedToAllDevices property value. The scopedToAllDevices property
func (m *DeviceCompliancePoliciesSetScheduledRetireStatePostRequestBody) GetScopedToAllDevices()(*bool) {
    val, err := m.GetBackingStore().Get("scopedToAllDevices")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetState gets the state property value. Cancel or confirm scheduled retire 
func (m *DeviceCompliancePoliciesSetScheduledRetireStatePostRequestBody) GetState()(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ScheduledRetireState) {
    val, err := m.GetBackingStore().Get("state")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ScheduledRetireState)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeviceCompliancePoliciesSetScheduledRetireStatePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetManagedDeviceIds() != nil {
        err := writer.WriteCollectionOfStringValues("managedDeviceIds", m.GetManagedDeviceIds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("scopedToAllDevices", m.GetScopedToAllDevices())
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err := writer.WriteStringValue("state", &cast)
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
func (m *DeviceCompliancePoliciesSetScheduledRetireStatePostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *DeviceCompliancePoliciesSetScheduledRetireStatePostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetManagedDeviceIds sets the managedDeviceIds property value. The managedDeviceIds property
func (m *DeviceCompliancePoliciesSetScheduledRetireStatePostRequestBody) SetManagedDeviceIds(value []string)() {
    err := m.GetBackingStore().Set("managedDeviceIds", value)
    if err != nil {
        panic(err)
    }
}
// SetScopedToAllDevices sets the scopedToAllDevices property value. The scopedToAllDevices property
func (m *DeviceCompliancePoliciesSetScheduledRetireStatePostRequestBody) SetScopedToAllDevices(value *bool)() {
    err := m.GetBackingStore().Set("scopedToAllDevices", value)
    if err != nil {
        panic(err)
    }
}
// SetState sets the state property value. Cancel or confirm scheduled retire 
func (m *DeviceCompliancePoliciesSetScheduledRetireStatePostRequestBody) SetState(value *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ScheduledRetireState)() {
    err := m.GetBackingStore().Set("state", value)
    if err != nil {
        panic(err)
    }
}
// DeviceCompliancePoliciesSetScheduledRetireStatePostRequestBodyable 
type DeviceCompliancePoliciesSetScheduledRetireStatePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetManagedDeviceIds()([]string)
    GetScopedToAllDevices()(*bool)
    GetState()(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ScheduledRetireState)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetManagedDeviceIds(value []string)()
    SetScopedToAllDevices(value *bool)()
    SetState(value *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ScheduledRetireState)()
}
