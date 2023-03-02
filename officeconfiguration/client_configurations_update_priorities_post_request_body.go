package officeconfiguration

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// ClientConfigurationsUpdatePrioritiesPostRequestBody 
type ClientConfigurationsUpdatePrioritiesPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewClientConfigurationsUpdatePrioritiesPostRequestBody instantiates a new ClientConfigurationsUpdatePrioritiesPostRequestBody and sets the default values.
func NewClientConfigurationsUpdatePrioritiesPostRequestBody()(*ClientConfigurationsUpdatePrioritiesPostRequestBody) {
    m := &ClientConfigurationsUpdatePrioritiesPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateClientConfigurationsUpdatePrioritiesPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateClientConfigurationsUpdatePrioritiesPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewClientConfigurationsUpdatePrioritiesPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ClientConfigurationsUpdatePrioritiesPostRequestBody) GetAdditionalData()(map[string]any) {
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
func (m *ClientConfigurationsUpdatePrioritiesPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ClientConfigurationsUpdatePrioritiesPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["officeConfigurationPolicyIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetOfficeConfigurationPolicyIds(res)
        }
        return nil
    }
    res["officeConfigurationPriorities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("int32")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]int32, len(val))
            for i, v := range val {
                res[i] = *(v.(*int32))
            }
            m.SetOfficeConfigurationPriorities(res)
        }
        return nil
    }
    return res
}
// GetOfficeConfigurationPolicyIds gets the officeConfigurationPolicyIds property value. The officeConfigurationPolicyIds property
func (m *ClientConfigurationsUpdatePrioritiesPostRequestBody) GetOfficeConfigurationPolicyIds()([]string) {
    val, err := m.GetBackingStore().Get("officeConfigurationPolicyIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetOfficeConfigurationPriorities gets the officeConfigurationPriorities property value. The officeConfigurationPriorities property
func (m *ClientConfigurationsUpdatePrioritiesPostRequestBody) GetOfficeConfigurationPriorities()([]int32) {
    val, err := m.GetBackingStore().Get("officeConfigurationPriorities")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ClientConfigurationsUpdatePrioritiesPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetOfficeConfigurationPolicyIds() != nil {
        err := writer.WriteCollectionOfStringValues("officeConfigurationPolicyIds", m.GetOfficeConfigurationPolicyIds())
        if err != nil {
            return err
        }
    }
    if m.GetOfficeConfigurationPriorities() != nil {
        err := writer.WriteCollectionOfInt32Values("officeConfigurationPriorities", m.GetOfficeConfigurationPriorities())
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
func (m *ClientConfigurationsUpdatePrioritiesPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *ClientConfigurationsUpdatePrioritiesPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetOfficeConfigurationPolicyIds sets the officeConfigurationPolicyIds property value. The officeConfigurationPolicyIds property
func (m *ClientConfigurationsUpdatePrioritiesPostRequestBody) SetOfficeConfigurationPolicyIds(value []string)() {
    err := m.GetBackingStore().Set("officeConfigurationPolicyIds", value)
    if err != nil {
        panic(err)
    }
}
// SetOfficeConfigurationPriorities sets the officeConfigurationPriorities property value. The officeConfigurationPriorities property
func (m *ClientConfigurationsUpdatePrioritiesPostRequestBody) SetOfficeConfigurationPriorities(value []int32)() {
    err := m.GetBackingStore().Set("officeConfigurationPriorities", value)
    if err != nil {
        panic(err)
    }
}
// ClientConfigurationsUpdatePrioritiesPostRequestBodyable 
type ClientConfigurationsUpdatePrioritiesPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetOfficeConfigurationPolicyIds()([]string)
    GetOfficeConfigurationPriorities()([]int32)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetOfficeConfigurationPolicyIds(value []string)()
    SetOfficeConfigurationPriorities(value []int32)()
}
