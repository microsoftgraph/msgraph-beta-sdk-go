package identity

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// B2cUserFlowsItemUserAttributeAssignmentsSetOrderPostRequestBody 
type B2cUserFlowsItemUserAttributeAssignmentsSetOrderPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewB2cUserFlowsItemUserAttributeAssignmentsSetOrderPostRequestBody instantiates a new B2cUserFlowsItemUserAttributeAssignmentsSetOrderPostRequestBody and sets the default values.
func NewB2cUserFlowsItemUserAttributeAssignmentsSetOrderPostRequestBody()(*B2cUserFlowsItemUserAttributeAssignmentsSetOrderPostRequestBody) {
    m := &B2cUserFlowsItemUserAttributeAssignmentsSetOrderPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateB2cUserFlowsItemUserAttributeAssignmentsSetOrderPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateB2cUserFlowsItemUserAttributeAssignmentsSetOrderPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewB2cUserFlowsItemUserAttributeAssignmentsSetOrderPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *B2cUserFlowsItemUserAttributeAssignmentsSetOrderPostRequestBody) GetAdditionalData()(map[string]any) {
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
func (m *B2cUserFlowsItemUserAttributeAssignmentsSetOrderPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *B2cUserFlowsItemUserAttributeAssignmentsSetOrderPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["newAssignmentOrder"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAssignmentOrderFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNewAssignmentOrder(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AssignmentOrderable))
        }
        return nil
    }
    return res
}
// GetNewAssignmentOrder gets the newAssignmentOrder property value. The newAssignmentOrder property
func (m *B2cUserFlowsItemUserAttributeAssignmentsSetOrderPostRequestBody) GetNewAssignmentOrder()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AssignmentOrderable) {
    val, err := m.GetBackingStore().Get("newAssignmentOrder")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AssignmentOrderable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *B2cUserFlowsItemUserAttributeAssignmentsSetOrderPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("newAssignmentOrder", m.GetNewAssignmentOrder())
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
func (m *B2cUserFlowsItemUserAttributeAssignmentsSetOrderPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *B2cUserFlowsItemUserAttributeAssignmentsSetOrderPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetNewAssignmentOrder sets the newAssignmentOrder property value. The newAssignmentOrder property
func (m *B2cUserFlowsItemUserAttributeAssignmentsSetOrderPostRequestBody) SetNewAssignmentOrder(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AssignmentOrderable)() {
    err := m.GetBackingStore().Set("newAssignmentOrder", value)
    if err != nil {
        panic(err)
    }
}
// B2cUserFlowsItemUserAttributeAssignmentsSetOrderPostRequestBodyable 
type B2cUserFlowsItemUserAttributeAssignmentsSetOrderPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetNewAssignmentOrder()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AssignmentOrderable)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetNewAssignmentOrder(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AssignmentOrderable)()
}
