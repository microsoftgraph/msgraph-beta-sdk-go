package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AppCallsItemRejectPostRequestBody provides operations to call the reject method.
type AppCallsItemRejectPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The callbackUri property
    callbackUri *string
    // The reason property
    reason *RejectReason
}
// NewAppCallsItemRejectPostRequestBody instantiates a new AppCallsItemRejectPostRequestBody and sets the default values.
func NewAppCallsItemRejectPostRequestBody()(*AppCallsItemRejectPostRequestBody) {
    m := &AppCallsItemRejectPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAppCallsItemRejectPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAppCallsItemRejectPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAppCallsItemRejectPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AppCallsItemRejectPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCallbackUri gets the callbackUri property value. The callbackUri property
func (m *AppCallsItemRejectPostRequestBody) GetCallbackUri()(*string) {
    return m.callbackUri
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AppCallsItemRejectPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["callbackUri"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCallbackUri(val)
        }
        return nil
    }
    res["reason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRejectReason)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReason(val.(*RejectReason))
        }
        return nil
    }
    return res
}
// GetReason gets the reason property value. The reason property
func (m *AppCallsItemRejectPostRequestBody) GetReason()(*RejectReason) {
    return m.reason
}
// Serialize serializes information the current object
func (m *AppCallsItemRejectPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("callbackUri", m.GetCallbackUri())
        if err != nil {
            return err
        }
    }
    if m.GetReason() != nil {
        cast := (*m.GetReason()).String()
        err := writer.WriteStringValue("reason", &cast)
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
func (m *AppCallsItemRejectPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCallbackUri sets the callbackUri property value. The callbackUri property
func (m *AppCallsItemRejectPostRequestBody) SetCallbackUri(value *string)() {
    m.callbackUri = value
}
// SetReason sets the reason property value. The reason property
func (m *AppCallsItemRejectPostRequestBody) SetReason(value *RejectReason)() {
    m.reason = value
}
