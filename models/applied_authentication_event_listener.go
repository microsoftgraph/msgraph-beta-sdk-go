package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AppliedAuthenticationEventListener 
type AppliedAuthenticationEventListener struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The eventType property
    eventType *AuthenticationEventType
    // The executedListenerId property
    executedListenerId *string
    // The handlerResult property
    handlerResult AuthenticationEventHandlerResultable
    // The OdataType property
    odataType *string
}
// NewAppliedAuthenticationEventListener instantiates a new appliedAuthenticationEventListener and sets the default values.
func NewAppliedAuthenticationEventListener()(*AppliedAuthenticationEventListener) {
    m := &AppliedAuthenticationEventListener{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.appliedAuthenticationEventListener";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAppliedAuthenticationEventListenerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAppliedAuthenticationEventListenerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAppliedAuthenticationEventListener(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AppliedAuthenticationEventListener) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetEventType gets the eventType property value. The eventType property
func (m *AppliedAuthenticationEventListener) GetEventType()(*AuthenticationEventType) {
    return m.eventType
}
// GetExecutedListenerId gets the executedListenerId property value. The executedListenerId property
func (m *AppliedAuthenticationEventListener) GetExecutedListenerId()(*string) {
    return m.executedListenerId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AppliedAuthenticationEventListener) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["eventType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseAuthenticationEventType , m.SetEventType)
    res["executedListenerId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetExecutedListenerId)
    res["handlerResult"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateAuthenticationEventHandlerResultFromDiscriminatorValue , m.SetHandlerResult)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    return res
}
// GetHandlerResult gets the handlerResult property value. The handlerResult property
func (m *AppliedAuthenticationEventListener) GetHandlerResult()(AuthenticationEventHandlerResultable) {
    return m.handlerResult
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AppliedAuthenticationEventListener) GetOdataType()(*string) {
    return m.odataType
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AppliedAuthenticationEventListener) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetEventType sets the eventType property value. The eventType property
func (m *AppliedAuthenticationEventListener) SetEventType(value *AuthenticationEventType)() {
    m.eventType = value
}
// SetExecutedListenerId sets the executedListenerId property value. The executedListenerId property
func (m *AppliedAuthenticationEventListener) SetExecutedListenerId(value *string)() {
    m.executedListenerId = value
}
// SetHandlerResult sets the handlerResult property value. The handlerResult property
func (m *AppliedAuthenticationEventListener) SetHandlerResult(value AuthenticationEventHandlerResultable)() {
    m.handlerResult = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AppliedAuthenticationEventListener) SetOdataType(value *string)() {
    m.odataType = value
}
