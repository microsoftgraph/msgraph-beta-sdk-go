package me

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// MailFoldersItemChildFoldersMicrosoftGraphDeltaDeltaResponse 
type MailFoldersItemChildFoldersMicrosoftGraphDeltaDeltaResponse struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BaseDeltaFunctionResponse
    // The value property
    value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailFolderable
}
// NewMailFoldersItemChildFoldersMicrosoftGraphDeltaDeltaResponse instantiates a new MailFoldersItemChildFoldersMicrosoftGraphDeltaDeltaResponse and sets the default values.
func NewMailFoldersItemChildFoldersMicrosoftGraphDeltaDeltaResponse()(*MailFoldersItemChildFoldersMicrosoftGraphDeltaDeltaResponse) {
    m := &MailFoldersItemChildFoldersMicrosoftGraphDeltaDeltaResponse{
        BaseDeltaFunctionResponse: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewBaseDeltaFunctionResponse(),
    }
    return m
}
// CreateMailFoldersItemChildFoldersMicrosoftGraphDeltaDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMailFoldersItemChildFoldersMicrosoftGraphDeltaDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMailFoldersItemChildFoldersMicrosoftGraphDeltaDeltaResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MailFoldersItemChildFoldersMicrosoftGraphDeltaDeltaResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseDeltaFunctionResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMailFolderFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailFolderable, len(val))
            for i, v := range val {
                res[i] = v.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailFolderable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *MailFoldersItemChildFoldersMicrosoftGraphDeltaDeltaResponse) GetValue()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailFolderable) {
    return m.value
}
// Serialize serializes information the current object
func (m *MailFoldersItemChildFoldersMicrosoftGraphDeltaDeltaResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseDeltaFunctionResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *MailFoldersItemChildFoldersMicrosoftGraphDeltaDeltaResponse) SetValue(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailFolderable)() {
    m.value = value
}
