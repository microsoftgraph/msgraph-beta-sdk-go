package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementGroupPolicyUploadedDefinitionFilesItemUploadNewVersionPostRequestBody provides operations to call the uploadNewVersion method.
type DeviceManagementGroupPolicyUploadedDefinitionFilesItemUploadNewVersionPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The content property
    content []byte
    // The groupPolicyUploadedLanguageFiles property
    groupPolicyUploadedLanguageFiles []GroupPolicyUploadedLanguageFileable
}
// NewDeviceManagementGroupPolicyUploadedDefinitionFilesItemUploadNewVersionPostRequestBody instantiates a new DeviceManagementGroupPolicyUploadedDefinitionFilesItemUploadNewVersionPostRequestBody and sets the default values.
func NewDeviceManagementGroupPolicyUploadedDefinitionFilesItemUploadNewVersionPostRequestBody()(*DeviceManagementGroupPolicyUploadedDefinitionFilesItemUploadNewVersionPostRequestBody) {
    m := &DeviceManagementGroupPolicyUploadedDefinitionFilesItemUploadNewVersionPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementGroupPolicyUploadedDefinitionFilesItemUploadNewVersionPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementGroupPolicyUploadedDefinitionFilesItemUploadNewVersionPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementGroupPolicyUploadedDefinitionFilesItemUploadNewVersionPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementGroupPolicyUploadedDefinitionFilesItemUploadNewVersionPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetContent gets the content property value. The content property
func (m *DeviceManagementGroupPolicyUploadedDefinitionFilesItemUploadNewVersionPostRequestBody) GetContent()([]byte) {
    return m.content
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementGroupPolicyUploadedDefinitionFilesItemUploadNewVersionPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["content"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContent(val)
        }
        return nil
    }
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
func (m *DeviceManagementGroupPolicyUploadedDefinitionFilesItemUploadNewVersionPostRequestBody) GetGroupPolicyUploadedLanguageFiles()([]GroupPolicyUploadedLanguageFileable) {
    return m.groupPolicyUploadedLanguageFiles
}
// Serialize serializes information the current object
func (m *DeviceManagementGroupPolicyUploadedDefinitionFilesItemUploadNewVersionPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteByteArrayValue("content", m.GetContent())
        if err != nil {
            return err
        }
    }
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
func (m *DeviceManagementGroupPolicyUploadedDefinitionFilesItemUploadNewVersionPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetContent sets the content property value. The content property
func (m *DeviceManagementGroupPolicyUploadedDefinitionFilesItemUploadNewVersionPostRequestBody) SetContent(value []byte)() {
    m.content = value
}
// SetGroupPolicyUploadedLanguageFiles sets the groupPolicyUploadedLanguageFiles property value. The groupPolicyUploadedLanguageFiles property
func (m *DeviceManagementGroupPolicyUploadedDefinitionFilesItemUploadNewVersionPostRequestBody) SetGroupPolicyUploadedLanguageFiles(value []GroupPolicyUploadedLanguageFileable)() {
    m.groupPolicyUploadedLanguageFiles = value
}
