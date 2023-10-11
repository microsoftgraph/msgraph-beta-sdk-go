package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// AndroidManagedStoreAccountEnterpriseSettingsSetAndroidDeviceOwnerFullyManagedEnrollmentStatePostRequestBody 
type AndroidManagedStoreAccountEnterpriseSettingsSetAndroidDeviceOwnerFullyManagedEnrollmentStatePostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewAndroidManagedStoreAccountEnterpriseSettingsSetAndroidDeviceOwnerFullyManagedEnrollmentStatePostRequestBody instantiates a new AndroidManagedStoreAccountEnterpriseSettingsSetAndroidDeviceOwnerFullyManagedEnrollmentStatePostRequestBody and sets the default values.
func NewAndroidManagedStoreAccountEnterpriseSettingsSetAndroidDeviceOwnerFullyManagedEnrollmentStatePostRequestBody()(*AndroidManagedStoreAccountEnterpriseSettingsSetAndroidDeviceOwnerFullyManagedEnrollmentStatePostRequestBody) {
    m := &AndroidManagedStoreAccountEnterpriseSettingsSetAndroidDeviceOwnerFullyManagedEnrollmentStatePostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAndroidManagedStoreAccountEnterpriseSettingsSetAndroidDeviceOwnerFullyManagedEnrollmentStatePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidManagedStoreAccountEnterpriseSettingsSetAndroidDeviceOwnerFullyManagedEnrollmentStatePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidManagedStoreAccountEnterpriseSettingsSetAndroidDeviceOwnerFullyManagedEnrollmentStatePostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AndroidManagedStoreAccountEnterpriseSettingsSetAndroidDeviceOwnerFullyManagedEnrollmentStatePostRequestBody) GetAdditionalData()(map[string]any) {
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
func (m *AndroidManagedStoreAccountEnterpriseSettingsSetAndroidDeviceOwnerFullyManagedEnrollmentStatePostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetEnabled gets the enabled property value. The enabled property
func (m *AndroidManagedStoreAccountEnterpriseSettingsSetAndroidDeviceOwnerFullyManagedEnrollmentStatePostRequestBody) GetEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("enabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidManagedStoreAccountEnterpriseSettingsSetAndroidDeviceOwnerFullyManagedEnrollmentStatePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["enabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnabled(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *AndroidManagedStoreAccountEnterpriseSettingsSetAndroidDeviceOwnerFullyManagedEnrollmentStatePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("enabled", m.GetEnabled())
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
func (m *AndroidManagedStoreAccountEnterpriseSettingsSetAndroidDeviceOwnerFullyManagedEnrollmentStatePostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *AndroidManagedStoreAccountEnterpriseSettingsSetAndroidDeviceOwnerFullyManagedEnrollmentStatePostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetEnabled sets the enabled property value. The enabled property
func (m *AndroidManagedStoreAccountEnterpriseSettingsSetAndroidDeviceOwnerFullyManagedEnrollmentStatePostRequestBody) SetEnabled(value *bool)() {
    err := m.GetBackingStore().Set("enabled", value)
    if err != nil {
        panic(err)
    }
}
// AndroidManagedStoreAccountEnterpriseSettingsSetAndroidDeviceOwnerFullyManagedEnrollmentStatePostRequestBodyable 
type AndroidManagedStoreAccountEnterpriseSettingsSetAndroidDeviceOwnerFullyManagedEnrollmentStatePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetEnabled()(*bool)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetEnabled(value *bool)()
}
