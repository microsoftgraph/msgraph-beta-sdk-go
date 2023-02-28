package tenantrelationships

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// ManagedTenantsManagedTenantAlertsItemManagedTenantsAddUserInputLogAddUserInputLogPostRequestBody 
type ManagedTenantsManagedTenantAlertsItemManagedTenantsAddUserInputLogAddUserInputLogPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewManagedTenantsManagedTenantAlertsItemManagedTenantsAddUserInputLogAddUserInputLogPostRequestBody instantiates a new ManagedTenantsManagedTenantAlertsItemManagedTenantsAddUserInputLogAddUserInputLogPostRequestBody and sets the default values.
func NewManagedTenantsManagedTenantAlertsItemManagedTenantsAddUserInputLogAddUserInputLogPostRequestBody()(*ManagedTenantsManagedTenantAlertsItemManagedTenantsAddUserInputLogAddUserInputLogPostRequestBody) {
    m := &ManagedTenantsManagedTenantAlertsItemManagedTenantsAddUserInputLogAddUserInputLogPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateManagedTenantsManagedTenantAlertsItemManagedTenantsAddUserInputLogAddUserInputLogPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedTenantsManagedTenantAlertsItemManagedTenantsAddUserInputLogAddUserInputLogPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedTenantsManagedTenantAlertsItemManagedTenantsAddUserInputLogAddUserInputLogPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ManagedTenantsManagedTenantAlertsItemManagedTenantsAddUserInputLogAddUserInputLogPostRequestBody) GetAdditionalData()(map[string]any) {
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
func (m *ManagedTenantsManagedTenantAlertsItemManagedTenantsAddUserInputLogAddUserInputLogPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedTenantsManagedTenantAlertsItemManagedTenantsAddUserInputLogAddUserInputLogPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["logInformation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLogInformation(val)
        }
        return nil
    }
    return res
}
// GetLogInformation gets the logInformation property value. The logInformation property
func (m *ManagedTenantsManagedTenantAlertsItemManagedTenantsAddUserInputLogAddUserInputLogPostRequestBody) GetLogInformation()(*string) {
    val, err := m.GetBackingStore().Get("logInformation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ManagedTenantsManagedTenantAlertsItemManagedTenantsAddUserInputLogAddUserInputLogPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("logInformation", m.GetLogInformation())
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
func (m *ManagedTenantsManagedTenantAlertsItemManagedTenantsAddUserInputLogAddUserInputLogPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *ManagedTenantsManagedTenantAlertsItemManagedTenantsAddUserInputLogAddUserInputLogPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetLogInformation sets the logInformation property value. The logInformation property
func (m *ManagedTenantsManagedTenantAlertsItemManagedTenantsAddUserInputLogAddUserInputLogPostRequestBody) SetLogInformation(value *string)() {
    err := m.GetBackingStore().Set("logInformation", value)
    if err != nil {
        panic(err)
    }
}
// ManagedTenantsManagedTenantAlertsItemManagedTenantsAddUserInputLogAddUserInputLogPostRequestBodyable 
type ManagedTenantsManagedTenantAlertsItemManagedTenantsAddUserInputLogAddUserInputLogPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetLogInformation()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetLogInformation(value *string)()
}
