package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type ReusablePolicySettingsItemReferencingConfigurationPoliciesItemSetEnrollmentTimeDeviceMembershipTargetPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewReusablePolicySettingsItemReferencingConfigurationPoliciesItemSetEnrollmentTimeDeviceMembershipTargetPostRequestBody instantiates a new ReusablePolicySettingsItemReferencingConfigurationPoliciesItemSetEnrollmentTimeDeviceMembershipTargetPostRequestBody and sets the default values.
func NewReusablePolicySettingsItemReferencingConfigurationPoliciesItemSetEnrollmentTimeDeviceMembershipTargetPostRequestBody()(*ReusablePolicySettingsItemReferencingConfigurationPoliciesItemSetEnrollmentTimeDeviceMembershipTargetPostRequestBody) {
    m := &ReusablePolicySettingsItemReferencingConfigurationPoliciesItemSetEnrollmentTimeDeviceMembershipTargetPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateReusablePolicySettingsItemReferencingConfigurationPoliciesItemSetEnrollmentTimeDeviceMembershipTargetPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateReusablePolicySettingsItemReferencingConfigurationPoliciesItemSetEnrollmentTimeDeviceMembershipTargetPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewReusablePolicySettingsItemReferencingConfigurationPoliciesItemSetEnrollmentTimeDeviceMembershipTargetPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ReusablePolicySettingsItemReferencingConfigurationPoliciesItemSetEnrollmentTimeDeviceMembershipTargetPostRequestBody) GetAdditionalData()(map[string]any) {
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
func (m *ReusablePolicySettingsItemReferencingConfigurationPoliciesItemSetEnrollmentTimeDeviceMembershipTargetPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetEnrollmentTimeDeviceMembershipTargets gets the enrollmentTimeDeviceMembershipTargets property value. The enrollmentTimeDeviceMembershipTargets property
// returns a []EnrollmentTimeDeviceMembershipTargetable when successful
func (m *ReusablePolicySettingsItemReferencingConfigurationPoliciesItemSetEnrollmentTimeDeviceMembershipTargetPostRequestBody) GetEnrollmentTimeDeviceMembershipTargets()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EnrollmentTimeDeviceMembershipTargetable) {
    val, err := m.GetBackingStore().Get("enrollmentTimeDeviceMembershipTargets")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EnrollmentTimeDeviceMembershipTargetable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ReusablePolicySettingsItemReferencingConfigurationPoliciesItemSetEnrollmentTimeDeviceMembershipTargetPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["enrollmentTimeDeviceMembershipTargets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEnrollmentTimeDeviceMembershipTargetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EnrollmentTimeDeviceMembershipTargetable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EnrollmentTimeDeviceMembershipTargetable)
                }
            }
            m.SetEnrollmentTimeDeviceMembershipTargets(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ReusablePolicySettingsItemReferencingConfigurationPoliciesItemSetEnrollmentTimeDeviceMembershipTargetPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetEnrollmentTimeDeviceMembershipTargets() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEnrollmentTimeDeviceMembershipTargets()))
        for i, v := range m.GetEnrollmentTimeDeviceMembershipTargets() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("enrollmentTimeDeviceMembershipTargets", cast)
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
func (m *ReusablePolicySettingsItemReferencingConfigurationPoliciesItemSetEnrollmentTimeDeviceMembershipTargetPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *ReusablePolicySettingsItemReferencingConfigurationPoliciesItemSetEnrollmentTimeDeviceMembershipTargetPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetEnrollmentTimeDeviceMembershipTargets sets the enrollmentTimeDeviceMembershipTargets property value. The enrollmentTimeDeviceMembershipTargets property
func (m *ReusablePolicySettingsItemReferencingConfigurationPoliciesItemSetEnrollmentTimeDeviceMembershipTargetPostRequestBody) SetEnrollmentTimeDeviceMembershipTargets(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EnrollmentTimeDeviceMembershipTargetable)() {
    err := m.GetBackingStore().Set("enrollmentTimeDeviceMembershipTargets", value)
    if err != nil {
        panic(err)
    }
}
type ReusablePolicySettingsItemReferencingConfigurationPoliciesItemSetEnrollmentTimeDeviceMembershipTargetPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetEnrollmentTimeDeviceMembershipTargets()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EnrollmentTimeDeviceMembershipTargetable)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetEnrollmentTimeDeviceMembershipTargets(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EnrollmentTimeDeviceMembershipTargetable)()
}
