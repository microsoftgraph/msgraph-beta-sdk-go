package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// DeviceConfigurationsItemAssignedAccessMultiModeProfilesPostRequestBody 
type DeviceConfigurationsItemAssignedAccessMultiModeProfilesPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewDeviceConfigurationsItemAssignedAccessMultiModeProfilesPostRequestBody instantiates a new DeviceConfigurationsItemAssignedAccessMultiModeProfilesPostRequestBody and sets the default values.
func NewDeviceConfigurationsItemAssignedAccessMultiModeProfilesPostRequestBody()(*DeviceConfigurationsItemAssignedAccessMultiModeProfilesPostRequestBody) {
    m := &DeviceConfigurationsItemAssignedAccessMultiModeProfilesPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDeviceConfigurationsItemAssignedAccessMultiModeProfilesPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceConfigurationsItemAssignedAccessMultiModeProfilesPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceConfigurationsItemAssignedAccessMultiModeProfilesPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceConfigurationsItemAssignedAccessMultiModeProfilesPostRequestBody) GetAdditionalData()(map[string]any) {
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
// GetAssignedAccessMultiModeProfiles gets the assignedAccessMultiModeProfiles property value. The assignedAccessMultiModeProfiles property
func (m *DeviceConfigurationsItemAssignedAccessMultiModeProfilesPostRequestBody) GetAssignedAccessMultiModeProfiles()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsAssignedAccessProfileable) {
    val, err := m.GetBackingStore().Get("assignedAccessMultiModeProfiles")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsAssignedAccessProfileable)
    }
    return nil
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *DeviceConfigurationsItemAssignedAccessMultiModeProfilesPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceConfigurationsItemAssignedAccessMultiModeProfilesPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["assignedAccessMultiModeProfiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsAssignedAccessProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsAssignedAccessProfileable, len(val))
            for i, v := range val {
                res[i] = v.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsAssignedAccessProfileable)
            }
            m.SetAssignedAccessMultiModeProfiles(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *DeviceConfigurationsItemAssignedAccessMultiModeProfilesPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAssignedAccessMultiModeProfiles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignedAccessMultiModeProfiles()))
        for i, v := range m.GetAssignedAccessMultiModeProfiles() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("assignedAccessMultiModeProfiles", cast)
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
func (m *DeviceConfigurationsItemAssignedAccessMultiModeProfilesPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAssignedAccessMultiModeProfiles sets the assignedAccessMultiModeProfiles property value. The assignedAccessMultiModeProfiles property
func (m *DeviceConfigurationsItemAssignedAccessMultiModeProfilesPostRequestBody) SetAssignedAccessMultiModeProfiles(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsAssignedAccessProfileable)() {
    err := m.GetBackingStore().Set("assignedAccessMultiModeProfiles", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *DeviceConfigurationsItemAssignedAccessMultiModeProfilesPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// DeviceConfigurationsItemAssignedAccessMultiModeProfilesPostRequestBodyable 
type DeviceConfigurationsItemAssignedAccessMultiModeProfilesPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignedAccessMultiModeProfiles()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsAssignedAccessProfileable)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    SetAssignedAccessMultiModeProfiles(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsAssignedAccessProfileable)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
}
