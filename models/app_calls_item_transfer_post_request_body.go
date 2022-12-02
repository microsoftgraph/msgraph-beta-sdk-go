package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AppCallsItemTransferPostRequestBody provides operations to call the transfer method.
type AppCallsItemTransferPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The transferee property
    transferee ParticipantInfoable
    // The transferTarget property
    transferTarget InvitationParticipantInfoable
}
// NewAppCallsItemTransferPostRequestBody instantiates a new AppCallsItemTransferPostRequestBody and sets the default values.
func NewAppCallsItemTransferPostRequestBody()(*AppCallsItemTransferPostRequestBody) {
    m := &AppCallsItemTransferPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAppCallsItemTransferPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAppCallsItemTransferPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAppCallsItemTransferPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AppCallsItemTransferPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AppCallsItemTransferPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["transferee"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateParticipantInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTransferee(val.(ParticipantInfoable))
        }
        return nil
    }
    res["transferTarget"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateInvitationParticipantInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTransferTarget(val.(InvitationParticipantInfoable))
        }
        return nil
    }
    return res
}
// GetTransferee gets the transferee property value. The transferee property
func (m *AppCallsItemTransferPostRequestBody) GetTransferee()(ParticipantInfoable) {
    return m.transferee
}
// GetTransferTarget gets the transferTarget property value. The transferTarget property
func (m *AppCallsItemTransferPostRequestBody) GetTransferTarget()(InvitationParticipantInfoable) {
    return m.transferTarget
}
// Serialize serializes information the current object
func (m *AppCallsItemTransferPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("transferee", m.GetTransferee())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("transferTarget", m.GetTransferTarget())
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
func (m *AppCallsItemTransferPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetTransferee sets the transferee property value. The transferee property
func (m *AppCallsItemTransferPostRequestBody) SetTransferee(value ParticipantInfoable)() {
    m.transferee = value
}
// SetTransferTarget sets the transferTarget property value. The transferTarget property
func (m *AppCallsItemTransferPostRequestBody) SetTransferTarget(value InvitationParticipantInfoable)() {
    m.transferTarget = value
}
