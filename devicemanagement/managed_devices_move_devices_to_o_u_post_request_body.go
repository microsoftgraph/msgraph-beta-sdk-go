package devicemanagement

import (
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// ManagedDevicesMoveDevicesToOUPostRequestBody 
type ManagedDevicesMoveDevicesToOUPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewManagedDevicesMoveDevicesToOUPostRequestBody instantiates a new ManagedDevicesMoveDevicesToOUPostRequestBody and sets the default values.
func NewManagedDevicesMoveDevicesToOUPostRequestBody()(*ManagedDevicesMoveDevicesToOUPostRequestBody) {
    m := &ManagedDevicesMoveDevicesToOUPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateManagedDevicesMoveDevicesToOUPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedDevicesMoveDevicesToOUPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedDevicesMoveDevicesToOUPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ManagedDevicesMoveDevicesToOUPostRequestBody) GetAdditionalData()(map[string]any) {
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
func (m *ManagedDevicesMoveDevicesToOUPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDeviceIds gets the deviceIds property value. The deviceIds property
func (m *ManagedDevicesMoveDevicesToOUPostRequestBody) GetDeviceIds()([]i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    val, err := m.GetBackingStore().Get("deviceIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedDevicesMoveDevicesToOUPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["deviceIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("uuid")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID, len(val))
            for i, v := range val {
                res[i] = *(v.(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID))
            }
            m.SetDeviceIds(res)
        }
        return nil
    }
    res["organizationalUnitPath"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationalUnitPath(val)
        }
        return nil
    }
    return res
}
// GetOrganizationalUnitPath gets the organizationalUnitPath property value. The organizationalUnitPath property
func (m *ManagedDevicesMoveDevicesToOUPostRequestBody) GetOrganizationalUnitPath()(*string) {
    val, err := m.GetBackingStore().Get("organizationalUnitPath")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ManagedDevicesMoveDevicesToOUPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDeviceIds() != nil {
        err := writer.WriteCollectionOfUUIDValues("deviceIds", m.GetDeviceIds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("organizationalUnitPath", m.GetOrganizationalUnitPath())
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
func (m *ManagedDevicesMoveDevicesToOUPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *ManagedDevicesMoveDevicesToOUPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDeviceIds sets the deviceIds property value. The deviceIds property
func (m *ManagedDevicesMoveDevicesToOUPostRequestBody) SetDeviceIds(value []i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    err := m.GetBackingStore().Set("deviceIds", value)
    if err != nil {
        panic(err)
    }
}
// SetOrganizationalUnitPath sets the organizationalUnitPath property value. The organizationalUnitPath property
func (m *ManagedDevicesMoveDevicesToOUPostRequestBody) SetOrganizationalUnitPath(value *string)() {
    err := m.GetBackingStore().Set("organizationalUnitPath", value)
    if err != nil {
        panic(err)
    }
}
// ManagedDevicesMoveDevicesToOUPostRequestBodyable 
type ManagedDevicesMoveDevicesToOUPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDeviceIds()([]i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetOrganizationalUnitPath()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDeviceIds(value []i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetOrganizationalUnitPath(value *string)()
}
