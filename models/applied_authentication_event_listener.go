package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type AppliedAuthenticationEventListener struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewAppliedAuthenticationEventListener instantiates a new AppliedAuthenticationEventListener and sets the default values.
func NewAppliedAuthenticationEventListener()(*AppliedAuthenticationEventListener) {
    m := &AppliedAuthenticationEventListener{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAppliedAuthenticationEventListenerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAppliedAuthenticationEventListenerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAppliedAuthenticationEventListener(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AppliedAuthenticationEventListener) GetAdditionalData()(map[string]any) {
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
func (m *AppliedAuthenticationEventListener) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetEventType gets the eventType property value. The type of authentication event that triggered the custom authentication extension request. The possible values are: tokenIssuanceStart, pageRenderStart, unknownFutureValue.
// returns a *AuthenticationEventType when successful
func (m *AppliedAuthenticationEventListener) GetEventType()(*AuthenticationEventType) {
    val, err := m.GetBackingStore().Get("eventType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AuthenticationEventType)
    }
    return nil
}
// GetExecutedListenerId gets the executedListenerId property value. ID of the authentication event listener that was executed.
// returns a *string when successful
func (m *AppliedAuthenticationEventListener) GetExecutedListenerId()(*string) {
    val, err := m.GetBackingStore().Get("executedListenerId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AppliedAuthenticationEventListener) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["eventType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAuthenticationEventType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventType(val.(*AuthenticationEventType))
        }
        return nil
    }
    res["executedListenerId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExecutedListenerId(val)
        }
        return nil
    }
    res["handlerResult"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthenticationEventHandlerResultFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHandlerResult(val.(AuthenticationEventHandlerResultable))
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    return res
}
// GetHandlerResult gets the handlerResult property value. The result from the listening client, such as an Azure Logic App and Azure Functions, of this authentication event.
// returns a AuthenticationEventHandlerResultable when successful
func (m *AppliedAuthenticationEventListener) GetHandlerResult()(AuthenticationEventHandlerResultable) {
    val, err := m.GetBackingStore().Get("handlerResult")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AuthenticationEventHandlerResultable)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *AppliedAuthenticationEventListener) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AppliedAuthenticationEventListener) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetEventType() != nil {
        cast := (*m.GetEventType()).String()
        err := writer.WriteStringValue("eventType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("executedListenerId", m.GetExecutedListenerId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("handlerResult", m.GetHandlerResult())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
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
func (m *AppliedAuthenticationEventListener) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *AppliedAuthenticationEventListener) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetEventType sets the eventType property value. The type of authentication event that triggered the custom authentication extension request. The possible values are: tokenIssuanceStart, pageRenderStart, unknownFutureValue.
func (m *AppliedAuthenticationEventListener) SetEventType(value *AuthenticationEventType)() {
    err := m.GetBackingStore().Set("eventType", value)
    if err != nil {
        panic(err)
    }
}
// SetExecutedListenerId sets the executedListenerId property value. ID of the authentication event listener that was executed.
func (m *AppliedAuthenticationEventListener) SetExecutedListenerId(value *string)() {
    err := m.GetBackingStore().Set("executedListenerId", value)
    if err != nil {
        panic(err)
    }
}
// SetHandlerResult sets the handlerResult property value. The result from the listening client, such as an Azure Logic App and Azure Functions, of this authentication event.
func (m *AppliedAuthenticationEventListener) SetHandlerResult(value AuthenticationEventHandlerResultable)() {
    err := m.GetBackingStore().Set("handlerResult", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AppliedAuthenticationEventListener) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
type AppliedAuthenticationEventListenerable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetEventType()(*AuthenticationEventType)
    GetExecutedListenerId()(*string)
    GetHandlerResult()(AuthenticationEventHandlerResultable)
    GetOdataType()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetEventType(value *AuthenticationEventType)()
    SetExecutedListenerId(value *string)()
    SetHandlerResult(value AuthenticationEventHandlerResultable)()
    SetOdataType(value *string)()
}
