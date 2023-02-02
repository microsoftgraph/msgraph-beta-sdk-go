package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// GroupPolicyUploadedDefinitionFilesItemMicrosoftGraphAddLanguageFilesAddLanguageFilesPostRequestBody 
type GroupPolicyUploadedDefinitionFilesItemMicrosoftGraphAddLanguageFilesAddLanguageFilesPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The groupPolicyUploadedLanguageFiles property
    groupPolicyUploadedLanguageFiles []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyUploadedLanguageFileable
}
// NewGroupPolicyUploadedDefinitionFilesItemMicrosoftGraphAddLanguageFilesAddLanguageFilesPostRequestBody instantiates a new GroupPolicyUploadedDefinitionFilesItemMicrosoftGraphAddLanguageFilesAddLanguageFilesPostRequestBody and sets the default values.
func NewGroupPolicyUploadedDefinitionFilesItemMicrosoftGraphAddLanguageFilesAddLanguageFilesPostRequestBody()(*GroupPolicyUploadedDefinitionFilesItemMicrosoftGraphAddLanguageFilesAddLanguageFilesPostRequestBody) {
    m := &GroupPolicyUploadedDefinitionFilesItemMicrosoftGraphAddLanguageFilesAddLanguageFilesPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGroupPolicyUploadedDefinitionFilesItemMicrosoftGraphAddLanguageFilesAddLanguageFilesPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupPolicyUploadedDefinitionFilesItemMicrosoftGraphAddLanguageFilesAddLanguageFilesPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupPolicyUploadedDefinitionFilesItemMicrosoftGraphAddLanguageFilesAddLanguageFilesPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GroupPolicyUploadedDefinitionFilesItemMicrosoftGraphAddLanguageFilesAddLanguageFilesPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupPolicyUploadedDefinitionFilesItemMicrosoftGraphAddLanguageFilesAddLanguageFilesPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["groupPolicyUploadedLanguageFiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupPolicyUploadedLanguageFileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyUploadedLanguageFileable, len(val))
            for i, v := range val {
                res[i] = v.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyUploadedLanguageFileable)
            }
            m.SetGroupPolicyUploadedLanguageFiles(res)
        }
        return nil
    }
    return res
}
// GetGroupPolicyUploadedLanguageFiles gets the groupPolicyUploadedLanguageFiles property value. The groupPolicyUploadedLanguageFiles property
func (m *GroupPolicyUploadedDefinitionFilesItemMicrosoftGraphAddLanguageFilesAddLanguageFilesPostRequestBody) GetGroupPolicyUploadedLanguageFiles()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyUploadedLanguageFileable) {
    return m.groupPolicyUploadedLanguageFiles
}
// Serialize serializes information the current object
func (m *GroupPolicyUploadedDefinitionFilesItemMicrosoftGraphAddLanguageFilesAddLanguageFilesPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetGroupPolicyUploadedLanguageFiles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetGroupPolicyUploadedLanguageFiles()))
        for i, v := range m.GetGroupPolicyUploadedLanguageFiles() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("groupPolicyUploadedLanguageFiles", cast)
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
func (m *GroupPolicyUploadedDefinitionFilesItemMicrosoftGraphAddLanguageFilesAddLanguageFilesPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetGroupPolicyUploadedLanguageFiles sets the groupPolicyUploadedLanguageFiles property value. The groupPolicyUploadedLanguageFiles property
func (m *GroupPolicyUploadedDefinitionFilesItemMicrosoftGraphAddLanguageFilesAddLanguageFilesPostRequestBody) SetGroupPolicyUploadedLanguageFiles(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyUploadedLanguageFileable)() {
    m.groupPolicyUploadedLanguageFiles = value
}
