package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AppliedAuthenticationEventListener 
type AppliedAuthenticationEventListener struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The type of authentication event that triggered the custom extension request. The possible values are: tokenIssuanceStart, pageRenderStart, unknownFutureValue.
    eventType *AuthenticationEventType
    // ID of the Event Listener that was executed.
    executedListenerId *string
    // The result from the listening client, such as an Azure Logic App and Azure Functions, of this authentication event.
    handlerResult AuthenticationEventHandlerResultable
    // The OdataType property
    odataType *string
}
// NewAppliedAuthenticationEventListener instantiates a new appliedAuthenticationEventListener and sets the default values.
func NewAppliedAuthenticationEventListener()(*AppliedAuthenticationEventListener) {
    m := &AppliedAuthenticationEventListener{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAppliedAuthenticationEventListenerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAppliedAuthenticationEventListenerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAppliedAuthenticationEventListener(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AppliedAuthenticationEventListener) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEventType gets the eventType property value. The type of authentication event that triggered the custom extension request. The possible values are: tokenIssuanceStart, pageRenderStart, unknownFutureValue.
func (m *AppliedAuthenticationEventListener) GetEventType()(*AuthenticationEventType) {
    return m.eventType
}
// GetExecutedListenerId gets the executedListenerId property value. ID of the Event Listener that was executed.
func (m *AppliedAuthenticationEventListener) GetExecutedListenerId()(*string) {
    return m.executedListenerId
}
// GetFieldDeserializers the deserialization information for the current model
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
func (m *AppliedAuthenticationEventListener) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEventType sets the eventType property value. The type of authentication event that triggered the custom extension request. The possible values are: tokenIssuanceStart, pageRenderStart, unknownFutureValue.
func (m *AppliedAuthenticationEventListener) SetEventType(value *AuthenticationEventType)() {
    m.eventType = value
}
// SetExecutedListenerId sets the executedListenerId property value. ID of the Event Listener that was executed.
func (m *AppliedAuthenticationEventListener) SetExecutedListenerId(value *string)() {
    m.executedListenerId = value
}
// SetHandlerResult sets the handlerResult property value. The result from the listening client, such as an Azure Logic App and Azure Functions, of this authentication event.
func (m *AppliedAuthenticationEventListener) SetHandlerResult(value AuthenticationEventHandlerResultable)() {
    m.handlerResult = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AppliedAuthenticationEventListener) SetOdataType(value *string)() {
    m.odataType = value
}
