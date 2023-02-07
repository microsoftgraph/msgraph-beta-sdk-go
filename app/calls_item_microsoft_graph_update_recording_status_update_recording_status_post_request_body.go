package app

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// CallsItemMicrosoftGraphUpdateRecordingStatusUpdateRecordingStatusPostRequestBody 
type CallsItemMicrosoftGraphUpdateRecordingStatusUpdateRecordingStatusPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The clientContext property
    clientContext *string
    // The status property
    status *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RecordingStatus
}
// NewCallsItemMicrosoftGraphUpdateRecordingStatusUpdateRecordingStatusPostRequestBody instantiates a new CallsItemMicrosoftGraphUpdateRecordingStatusUpdateRecordingStatusPostRequestBody and sets the default values.
func NewCallsItemMicrosoftGraphUpdateRecordingStatusUpdateRecordingStatusPostRequestBody()(*CallsItemMicrosoftGraphUpdateRecordingStatusUpdateRecordingStatusPostRequestBody) {
    m := &CallsItemMicrosoftGraphUpdateRecordingStatusUpdateRecordingStatusPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCallsItemMicrosoftGraphUpdateRecordingStatusUpdateRecordingStatusPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCallsItemMicrosoftGraphUpdateRecordingStatusUpdateRecordingStatusPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCallsItemMicrosoftGraphUpdateRecordingStatusUpdateRecordingStatusPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CallsItemMicrosoftGraphUpdateRecordingStatusUpdateRecordingStatusPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetClientContext gets the clientContext property value. The clientContext property
func (m *CallsItemMicrosoftGraphUpdateRecordingStatusUpdateRecordingStatusPostRequestBody) GetClientContext()(*string) {
    return m.clientContext
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CallsItemMicrosoftGraphUpdateRecordingStatusUpdateRecordingStatusPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["clientContext"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientContext(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ParseRecordingStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RecordingStatus))
        }
        return nil
    }
    return res
}
// GetStatus gets the status property value. The status property
func (m *CallsItemMicrosoftGraphUpdateRecordingStatusUpdateRecordingStatusPostRequestBody) GetStatus()(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RecordingStatus) {
    return m.status
}
// Serialize serializes information the current object
func (m *CallsItemMicrosoftGraphUpdateRecordingStatusUpdateRecordingStatusPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("clientContext", m.GetClientContext())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err := writer.WriteStringValue("status", &cast)
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
func (m *CallsItemMicrosoftGraphUpdateRecordingStatusUpdateRecordingStatusPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetClientContext sets the clientContext property value. The clientContext property
func (m *CallsItemMicrosoftGraphUpdateRecordingStatusUpdateRecordingStatusPostRequestBody) SetClientContext(value *string)() {
    m.clientContext = value
}
// SetStatus sets the status property value. The status property
func (m *CallsItemMicrosoftGraphUpdateRecordingStatusUpdateRecordingStatusPostRequestBody) SetStatus(value *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RecordingStatus)() {
    m.status = value
}
