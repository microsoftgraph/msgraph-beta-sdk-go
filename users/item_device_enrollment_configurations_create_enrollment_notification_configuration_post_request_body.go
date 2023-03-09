package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// ItemDeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationPostRequestBody 
type ItemDeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewItemDeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationPostRequestBody instantiates a new ItemDeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationPostRequestBody and sets the default values.
func NewItemDeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationPostRequestBody()(*ItemDeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationPostRequestBody) {
    m := &ItemDeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemDeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemDeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemDeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemDeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationPostRequestBody) GetAdditionalData()(map[string]any) {
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
func (m *ItemDeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDeviceEnrollmentNotificationConfigurations gets the deviceEnrollmentNotificationConfigurations property value. The deviceEnrollmentNotificationConfigurations property
func (m *ItemDeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationPostRequestBody) GetDeviceEnrollmentNotificationConfigurations()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceEnrollmentConfigurationable) {
    val, err := m.GetBackingStore().Get("deviceEnrollmentNotificationConfigurations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceEnrollmentConfigurationable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemDeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["deviceEnrollmentNotificationConfigurations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceEnrollmentConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceEnrollmentConfigurationable, len(val))
            for i, v := range val {
                res[i] = v.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceEnrollmentConfigurationable)
            }
            m.SetDeviceEnrollmentNotificationConfigurations(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ItemDeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDeviceEnrollmentNotificationConfigurations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceEnrollmentNotificationConfigurations()))
        for i, v := range m.GetDeviceEnrollmentNotificationConfigurations() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("deviceEnrollmentNotificationConfigurations", cast)
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
func (m *ItemDeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *ItemDeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDeviceEnrollmentNotificationConfigurations sets the deviceEnrollmentNotificationConfigurations property value. The deviceEnrollmentNotificationConfigurations property
func (m *ItemDeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationPostRequestBody) SetDeviceEnrollmentNotificationConfigurations(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceEnrollmentConfigurationable)() {
    err := m.GetBackingStore().Set("deviceEnrollmentNotificationConfigurations", value)
    if err != nil {
        panic(err)
    }
}
// ItemDeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationPostRequestBodyable 
type ItemDeviceEnrollmentConfigurationsCreateEnrollmentNotificationConfigurationPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDeviceEnrollmentNotificationConfigurations()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceEnrollmentConfigurationable)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDeviceEnrollmentNotificationConfigurations(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceEnrollmentConfigurationable)()
}
