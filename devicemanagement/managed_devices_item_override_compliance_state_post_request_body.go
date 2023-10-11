package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// ManagedDevicesItemOverrideComplianceStatePostRequestBody 
type ManagedDevicesItemOverrideComplianceStatePostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewManagedDevicesItemOverrideComplianceStatePostRequestBody instantiates a new ManagedDevicesItemOverrideComplianceStatePostRequestBody and sets the default values.
func NewManagedDevicesItemOverrideComplianceStatePostRequestBody()(*ManagedDevicesItemOverrideComplianceStatePostRequestBody) {
    m := &ManagedDevicesItemOverrideComplianceStatePostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateManagedDevicesItemOverrideComplianceStatePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedDevicesItemOverrideComplianceStatePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedDevicesItemOverrideComplianceStatePostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ManagedDevicesItemOverrideComplianceStatePostRequestBody) GetAdditionalData()(map[string]any) {
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
func (m *ManagedDevicesItemOverrideComplianceStatePostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetComplianceState gets the complianceState property value. Administrator configured device compliance state Enum
func (m *ManagedDevicesItemOverrideComplianceStatePostRequestBody) GetComplianceState()(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdministratorConfiguredDeviceComplianceState) {
    val, err := m.GetBackingStore().Get("complianceState")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdministratorConfiguredDeviceComplianceState)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedDevicesItemOverrideComplianceStatePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["complianceState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ParseAdministratorConfiguredDeviceComplianceState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComplianceState(val.(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdministratorConfiguredDeviceComplianceState))
        }
        return nil
    }
    res["remediationUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemediationUrl(val)
        }
        return nil
    }
    return res
}
// GetRemediationUrl gets the remediationUrl property value. The remediationUrl property
func (m *ManagedDevicesItemOverrideComplianceStatePostRequestBody) GetRemediationUrl()(*string) {
    val, err := m.GetBackingStore().Get("remediationUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ManagedDevicesItemOverrideComplianceStatePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetComplianceState() != nil {
        cast := (*m.GetComplianceState()).String()
        err := writer.WriteStringValue("complianceState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("remediationUrl", m.GetRemediationUrl())
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
func (m *ManagedDevicesItemOverrideComplianceStatePostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *ManagedDevicesItemOverrideComplianceStatePostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetComplianceState sets the complianceState property value. Administrator configured device compliance state Enum
func (m *ManagedDevicesItemOverrideComplianceStatePostRequestBody) SetComplianceState(value *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdministratorConfiguredDeviceComplianceState)() {
    err := m.GetBackingStore().Set("complianceState", value)
    if err != nil {
        panic(err)
    }
}
// SetRemediationUrl sets the remediationUrl property value. The remediationUrl property
func (m *ManagedDevicesItemOverrideComplianceStatePostRequestBody) SetRemediationUrl(value *string)() {
    err := m.GetBackingStore().Set("remediationUrl", value)
    if err != nil {
        panic(err)
    }
}
// ManagedDevicesItemOverrideComplianceStatePostRequestBodyable 
type ManagedDevicesItemOverrideComplianceStatePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetComplianceState()(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdministratorConfiguredDeviceComplianceState)
    GetRemediationUrl()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetComplianceState(value *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdministratorConfiguredDeviceComplianceState)()
    SetRemediationUrl(value *string)()
}
