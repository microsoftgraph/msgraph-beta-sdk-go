package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// LifecycleWorkflowsDeletedItemsWorkflowsItemIdentityGovernanceActivateActivatePostRequestBody 
type LifecycleWorkflowsDeletedItemsWorkflowsItemIdentityGovernanceActivateActivatePostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewLifecycleWorkflowsDeletedItemsWorkflowsItemIdentityGovernanceActivateActivatePostRequestBody instantiates a new LifecycleWorkflowsDeletedItemsWorkflowsItemIdentityGovernanceActivateActivatePostRequestBody and sets the default values.
func NewLifecycleWorkflowsDeletedItemsWorkflowsItemIdentityGovernanceActivateActivatePostRequestBody()(*LifecycleWorkflowsDeletedItemsWorkflowsItemIdentityGovernanceActivateActivatePostRequestBody) {
    m := &LifecycleWorkflowsDeletedItemsWorkflowsItemIdentityGovernanceActivateActivatePostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateLifecycleWorkflowsDeletedItemsWorkflowsItemIdentityGovernanceActivateActivatePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLifecycleWorkflowsDeletedItemsWorkflowsItemIdentityGovernanceActivateActivatePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLifecycleWorkflowsDeletedItemsWorkflowsItemIdentityGovernanceActivateActivatePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LifecycleWorkflowsDeletedItemsWorkflowsItemIdentityGovernanceActivateActivatePostRequestBody) GetAdditionalData()(map[string]any) {
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
func (m *LifecycleWorkflowsDeletedItemsWorkflowsItemIdentityGovernanceActivateActivatePostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LifecycleWorkflowsDeletedItemsWorkflowsItemIdentityGovernanceActivateActivatePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["subjects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, len(val))
            for i, v := range val {
                res[i] = v.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable)
            }
            m.SetSubjects(res)
        }
        return nil
    }
    return res
}
// GetSubjects gets the subjects property value. The subjects property
func (m *LifecycleWorkflowsDeletedItemsWorkflowsItemIdentityGovernanceActivateActivatePostRequestBody) GetSubjects()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable) {
    val, err := m.GetBackingStore().Get("subjects")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *LifecycleWorkflowsDeletedItemsWorkflowsItemIdentityGovernanceActivateActivatePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetSubjects() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSubjects()))
        for i, v := range m.GetSubjects() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("subjects", cast)
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
func (m *LifecycleWorkflowsDeletedItemsWorkflowsItemIdentityGovernanceActivateActivatePostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *LifecycleWorkflowsDeletedItemsWorkflowsItemIdentityGovernanceActivateActivatePostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetSubjects sets the subjects property value. The subjects property
func (m *LifecycleWorkflowsDeletedItemsWorkflowsItemIdentityGovernanceActivateActivatePostRequestBody) SetSubjects(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable)() {
    err := m.GetBackingStore().Set("subjects", value)
    if err != nil {
        panic(err)
    }
}
// LifecycleWorkflowsDeletedItemsWorkflowsItemIdentityGovernanceActivateActivatePostRequestBodyable 
type LifecycleWorkflowsDeletedItemsWorkflowsItemIdentityGovernanceActivateActivatePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetSubjects()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetSubjects(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable)()
}
