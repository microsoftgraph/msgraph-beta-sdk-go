package identity

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type ConditionalAccessEvaluatePostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewConditionalAccessEvaluatePostRequestBody instantiates a new ConditionalAccessEvaluatePostRequestBody and sets the default values.
func NewConditionalAccessEvaluatePostRequestBody()(*ConditionalAccessEvaluatePostRequestBody) {
    m := &ConditionalAccessEvaluatePostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateConditionalAccessEvaluatePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateConditionalAccessEvaluatePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConditionalAccessEvaluatePostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ConditionalAccessEvaluatePostRequestBody) GetAdditionalData()(map[string]any) {
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
// GetAppliedPoliciesOnly gets the appliedPoliciesOnly property value. The appliedPoliciesOnly property
// returns a *bool when successful
func (m *ConditionalAccessEvaluatePostRequestBody) GetAppliedPoliciesOnly()(*bool) {
    val, err := m.GetBackingStore().Get("appliedPoliciesOnly")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *ConditionalAccessEvaluatePostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetConditionalAccessContext gets the conditionalAccessContext property value. The conditionalAccessContext property
// returns a ConditionalAccessContextable when successful
func (m *ConditionalAccessEvaluatePostRequestBody) GetConditionalAccessContext()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConditionalAccessContextable) {
    val, err := m.GetBackingStore().Get("conditionalAccessContext")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConditionalAccessContextable)
    }
    return nil
}
// GetConditionalAccessWhatIfConditions gets the conditionalAccessWhatIfConditions property value. The conditionalAccessWhatIfConditions property
// returns a ConditionalAccessWhatIfConditionsable when successful
func (m *ConditionalAccessEvaluatePostRequestBody) GetConditionalAccessWhatIfConditions()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConditionalAccessWhatIfConditionsable) {
    val, err := m.GetBackingStore().Get("conditionalAccessWhatIfConditions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConditionalAccessWhatIfConditionsable)
    }
    return nil
}
// GetConditionalAccessWhatIfSubject gets the conditionalAccessWhatIfSubject property value. The conditionalAccessWhatIfSubject property
// returns a ConditionalAccessWhatIfSubjectable when successful
func (m *ConditionalAccessEvaluatePostRequestBody) GetConditionalAccessWhatIfSubject()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConditionalAccessWhatIfSubjectable) {
    val, err := m.GetBackingStore().Get("conditionalAccessWhatIfSubject")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConditionalAccessWhatIfSubjectable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ConditionalAccessEvaluatePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["appliedPoliciesOnly"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppliedPoliciesOnly(val)
        }
        return nil
    }
    res["conditionalAccessContext"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateConditionalAccessContextFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConditionalAccessContext(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConditionalAccessContextable))
        }
        return nil
    }
    res["conditionalAccessWhatIfConditions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateConditionalAccessWhatIfConditionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConditionalAccessWhatIfConditions(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConditionalAccessWhatIfConditionsable))
        }
        return nil
    }
    res["conditionalAccessWhatIfSubject"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateConditionalAccessWhatIfSubjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConditionalAccessWhatIfSubject(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConditionalAccessWhatIfSubjectable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ConditionalAccessEvaluatePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("appliedPoliciesOnly", m.GetAppliedPoliciesOnly())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("conditionalAccessContext", m.GetConditionalAccessContext())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("conditionalAccessWhatIfConditions", m.GetConditionalAccessWhatIfConditions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("conditionalAccessWhatIfSubject", m.GetConditionalAccessWhatIfSubject())
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
func (m *ConditionalAccessEvaluatePostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAppliedPoliciesOnly sets the appliedPoliciesOnly property value. The appliedPoliciesOnly property
func (m *ConditionalAccessEvaluatePostRequestBody) SetAppliedPoliciesOnly(value *bool)() {
    err := m.GetBackingStore().Set("appliedPoliciesOnly", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *ConditionalAccessEvaluatePostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetConditionalAccessContext sets the conditionalAccessContext property value. The conditionalAccessContext property
func (m *ConditionalAccessEvaluatePostRequestBody) SetConditionalAccessContext(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConditionalAccessContextable)() {
    err := m.GetBackingStore().Set("conditionalAccessContext", value)
    if err != nil {
        panic(err)
    }
}
// SetConditionalAccessWhatIfConditions sets the conditionalAccessWhatIfConditions property value. The conditionalAccessWhatIfConditions property
func (m *ConditionalAccessEvaluatePostRequestBody) SetConditionalAccessWhatIfConditions(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConditionalAccessWhatIfConditionsable)() {
    err := m.GetBackingStore().Set("conditionalAccessWhatIfConditions", value)
    if err != nil {
        panic(err)
    }
}
// SetConditionalAccessWhatIfSubject sets the conditionalAccessWhatIfSubject property value. The conditionalAccessWhatIfSubject property
func (m *ConditionalAccessEvaluatePostRequestBody) SetConditionalAccessWhatIfSubject(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConditionalAccessWhatIfSubjectable)() {
    err := m.GetBackingStore().Set("conditionalAccessWhatIfSubject", value)
    if err != nil {
        panic(err)
    }
}
type ConditionalAccessEvaluatePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppliedPoliciesOnly()(*bool)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetConditionalAccessContext()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConditionalAccessContextable)
    GetConditionalAccessWhatIfConditions()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConditionalAccessWhatIfConditionsable)
    GetConditionalAccessWhatIfSubject()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConditionalAccessWhatIfSubjectable)
    SetAppliedPoliciesOnly(value *bool)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetConditionalAccessContext(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConditionalAccessContextable)()
    SetConditionalAccessWhatIfConditions(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConditionalAccessWhatIfConditionsable)()
    SetConditionalAccessWhatIfSubject(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConditionalAccessWhatIfSubjectable)()
}
