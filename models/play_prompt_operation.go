package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PlayPromptOperation 
type PlayPromptOperation struct {
    CommsOperation
    // Possible values are: unknown, completedSuccessfully, mediaOperationCanceled.
    completionReason *PlayPromptCompletionReason
}
// NewPlayPromptOperation instantiates a new PlayPromptOperation and sets the default values.
func NewPlayPromptOperation()(*PlayPromptOperation) {
    m := &PlayPromptOperation{
        CommsOperation: *NewCommsOperation(),
    }
    return m
}
// CreatePlayPromptOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlayPromptOperationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPlayPromptOperation(), nil
}
// GetCompletionReason gets the completionReason property value. Possible values are: unknown, completedSuccessfully, mediaOperationCanceled.
func (m *PlayPromptOperation) GetCompletionReason()(*PlayPromptCompletionReason) {
    return m.completionReason
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlayPromptOperation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CommsOperation.GetFieldDeserializers()
    res["completionReason"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParsePlayPromptCompletionReason , m.SetCompletionReason)
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
    return nil
}
// SetCompletionReason sets the completionReason property value. Possible values are: unknown, completedSuccessfully, mediaOperationCanceled.
func (m *PlayPromptOperation) SetCompletionReason(value *PlayPromptCompletionReason)() {
    m.completionReason = value
}
