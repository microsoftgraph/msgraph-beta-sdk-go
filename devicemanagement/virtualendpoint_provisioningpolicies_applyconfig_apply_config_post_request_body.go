package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type VirtualendpointProvisioningpoliciesApplyconfigApplyConfigPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewVirtualendpointProvisioningpoliciesApplyconfigApplyConfigPostRequestBody instantiates a new VirtualendpointProvisioningpoliciesApplyconfigApplyConfigPostRequestBody and sets the default values.
func NewVirtualendpointProvisioningpoliciesApplyconfigApplyConfigPostRequestBody()(*VirtualendpointProvisioningpoliciesApplyconfigApplyConfigPostRequestBody) {
    m := &VirtualendpointProvisioningpoliciesApplyconfigApplyConfigPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateVirtualendpointProvisioningpoliciesApplyconfigApplyConfigPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateVirtualendpointProvisioningpoliciesApplyconfigApplyConfigPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVirtualendpointProvisioningpoliciesApplyconfigApplyConfigPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *VirtualendpointProvisioningpoliciesApplyconfigApplyConfigPostRequestBody) GetAdditionalData()(map[string]any) {
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
func (m *VirtualendpointProvisioningpoliciesApplyconfigApplyConfigPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetCloudPcIds gets the cloudPcIds property value. The cloudPcIds property
// returns a []string when successful
func (m *VirtualendpointProvisioningpoliciesApplyconfigApplyConfigPostRequestBody) GetCloudPcIds()([]string) {
    val, err := m.GetBackingStore().Get("cloudPcIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *VirtualendpointProvisioningpoliciesApplyconfigApplyConfigPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["cloudPcIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetCloudPcIds(res)
        }
        return nil
    }
    res["policySettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ParseCloudPcPolicySettingType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicySettings(val.(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcPolicySettingType))
        }
        return nil
    }
    return res
}
// GetPolicySettings gets the policySettings property value. The policySettings property
// returns a *CloudPcPolicySettingType when successful
func (m *VirtualendpointProvisioningpoliciesApplyconfigApplyConfigPostRequestBody) GetPolicySettings()(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcPolicySettingType) {
    val, err := m.GetBackingStore().Get("policySettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcPolicySettingType)
    }
    return nil
}
// Serialize serializes information the current object
func (m *VirtualendpointProvisioningpoliciesApplyconfigApplyConfigPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCloudPcIds() != nil {
        err := writer.WriteCollectionOfStringValues("cloudPcIds", m.GetCloudPcIds())
        if err != nil {
            return err
        }
    }
    if m.GetPolicySettings() != nil {
        cast := (*m.GetPolicySettings()).String()
        err := writer.WriteStringValue("policySettings", &cast)
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
func (m *VirtualendpointProvisioningpoliciesApplyconfigApplyConfigPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *VirtualendpointProvisioningpoliciesApplyconfigApplyConfigPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetCloudPcIds sets the cloudPcIds property value. The cloudPcIds property
func (m *VirtualendpointProvisioningpoliciesApplyconfigApplyConfigPostRequestBody) SetCloudPcIds(value []string)() {
    err := m.GetBackingStore().Set("cloudPcIds", value)
    if err != nil {
        panic(err)
    }
}
// SetPolicySettings sets the policySettings property value. The policySettings property
func (m *VirtualendpointProvisioningpoliciesApplyconfigApplyConfigPostRequestBody) SetPolicySettings(value *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcPolicySettingType)() {
    err := m.GetBackingStore().Set("policySettings", value)
    if err != nil {
        panic(err)
    }
}
type VirtualendpointProvisioningpoliciesApplyconfigApplyConfigPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetCloudPcIds()([]string)
    GetPolicySettings()(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcPolicySettingType)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetCloudPcIds(value []string)()
    SetPolicySettings(value *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcPolicySettingType)()
}
