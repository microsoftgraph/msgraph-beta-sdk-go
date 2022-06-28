package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SwapShiftsChangeRequest 
type SwapShiftsChangeRequest struct {
    OfferShiftRequest
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Shift ID for the recipient user with whom the request is to swap.
    recipientShiftId *string
}
// NewSwapShiftsChangeRequest instantiates a new SwapShiftsChangeRequest and sets the default values.
func NewSwapShiftsChangeRequest()(*SwapShiftsChangeRequest) {
    m := &SwapShiftsChangeRequest{
        OfferShiftRequest: *NewOfferShiftRequest(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSwapShiftsChangeRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSwapShiftsChangeRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSwapShiftsChangeRequest(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SwapShiftsChangeRequest) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SwapShiftsChangeRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.OfferShiftRequest.GetFieldDeserializers()
    res["recipientShiftId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecipientShiftId(val)
        }
        return nil
    }
    return res
}
// GetRecipientShiftId gets the recipientShiftId property value. Shift ID for the recipient user with whom the request is to swap.
func (m *SwapShiftsChangeRequest) GetRecipientShiftId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.recipientShiftId
    }
}
// Serialize serializes information the current object
func (m *SwapShiftsChangeRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.OfferShiftRequest.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("recipientShiftId", m.GetRecipientShiftId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SwapShiftsChangeRequest) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetRecipientShiftId sets the recipientShiftId property value. Shift ID for the recipient user with whom the request is to swap.
func (m *SwapShiftsChangeRequest) SetRecipientShiftId(value *string)() {
    if m != nil {
        m.recipientShiftId = value
    }
}
