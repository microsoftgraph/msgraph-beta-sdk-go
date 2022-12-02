package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementGroupPolicyUploadedDefinitionFilesItemAddLanguageFilesPostRequestBody provides operations to call the addLanguageFiles method.
type DeviceManagementGroupPolicyUploadedDefinitionFilesItemAddLanguageFilesPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The groupPolicyUploadedLanguageFiles property
    groupPolicyUploadedLanguageFiles []GroupPolicyUploadedLanguageFileable
}
// NewDeviceManagementGroupPolicyUploadedDefinitionFilesItemAddLanguageFilesPostRequestBody instantiates a new DeviceManagementGroupPolicyUploadedDefinitionFilesItemAddLanguageFilesPostRequestBody and sets the default values.
func NewDeviceManagementGroupPolicyUploadedDefinitionFilesItemAddLanguageFilesPostRequestBody()(*DeviceManagementGroupPolicyUploadedDefinitionFilesItemAddLanguageFilesPostRequestBody) {
    m := &DeviceManagementGroupPolicyUploadedDefinitionFilesItemAddLanguageFilesPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementGroupPolicyUploadedDefinitionFilesItemAddLanguageFilesPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementGroupPolicyUploadedDefinitionFilesItemAddLanguageFilesPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementGroupPolicyUploadedDefinitionFilesItemAddLanguageFilesPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementGroupPolicyUploadedDefinitionFilesItemAddLanguageFilesPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementGroupPolicyUploadedDefinitionFilesItemAddLanguageFilesPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["groupPolicyUploadedLanguageFiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGroupPolicyUploadedLanguageFileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GroupPolicyUploadedLanguageFileable, len(val))
            for i, v := range val {
                res[i] = v.(GroupPolicyUploadedLanguageFileable)
            }
            m.SetGroupPolicyUploadedLanguageFiles(res)
        }
        return nil
    }
    return res
}
// GetGroupPolicyUploadedLanguageFiles gets the groupPolicyUploadedLanguageFiles property value. The groupPolicyUploadedLanguageFiles property
func (m *DeviceManagementGroupPolicyUploadedDefinitionFilesItemAddLanguageFilesPostRequestBody) GetGroupPolicyUploadedLanguageFiles()([]GroupPolicyUploadedLanguageFileable) {
    return m.groupPolicyUploadedLanguageFiles
}
// Serialize serializes information the current object
func (m *DeviceManagementGroupPolicyUploadedDefinitionFilesItemAddLanguageFilesPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *DeviceManagementGroupPolicyUploadedDefinitionFilesItemAddLanguageFilesPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetGroupPolicyUploadedLanguageFiles sets the groupPolicyUploadedLanguageFiles property value. The groupPolicyUploadedLanguageFiles property
func (m *DeviceManagementGroupPolicyUploadedDefinitionFilesItemAddLanguageFilesPostRequestBody) SetGroupPolicyUploadedLanguageFiles(value []GroupPolicyUploadedLanguageFileable)() {
    m.groupPolicyUploadedLanguageFiles = value
}
