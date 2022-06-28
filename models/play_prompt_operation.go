package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PlayPromptOperation 
type PlayPromptOperation struct {
    CommsOperation
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Possible values are: unknown, completedSuccessfully, mediaOperationCanceled.
    completionReason *PlayPromptCompletionReason
}
// NewPlayPromptOperation instantiates a new PlayPromptOperation and sets the default values.
func NewPlayPromptOperation()(*PlayPromptOperation) {
    m := &PlayPromptOperation{
        CommsOperation: *NewCommsOperation(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreatePlayPromptOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlayPromptOperationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPlayPromptOperation(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PlayPromptOperation) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCompletionReason gets the completionReason property value. Possible values are: unknown, completedSuccessfully, mediaOperationCanceled.
func (m *PlayPromptOperation) GetCompletionReason()(*PlayPromptCompletionReason) {
    if m == nil {
        return nil
    } else {
        return m.completionReason
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlayPromptOperation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CommsOperation.GetFieldDeserializers()
    res["completionReason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePlayPromptCompletionReason)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompletionReason(val.(*PlayPromptCompletionReason))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *PlayPromptOperation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CommsOperation.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCompletionReason() != nil {
        cast := (*m.GetCompletionReason()).String()
        err = writer.WriteStringValue("completionReason", &cast)
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
func (m *PlayPromptOperation) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCompletionReason sets the completionReason property value. Possible values are: unknown, completedSuccessfully, mediaOperationCanceled.
func (m *PlayPromptOperation) SetCompletionReason(value *PlayPromptCompletionReason)() {
    if m != nil {
        m.completionReason = value
    }
}
