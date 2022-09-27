package resume

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/identitygovernance"
)

// ResumePostRequestBody provides operations to call the resume method.
type ResumePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The data property
    data i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.CustomTaskExtensionCallbackDataable
    // The source property
    source *string
    // The type property
    type_escaped *string
}
// NewResumePostRequestBody instantiates a new resumePostRequestBody and sets the default values.
func NewResumePostRequestBody()(*ResumePostRequestBody) {
    m := &ResumePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateResumePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateResumePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewResumePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ResumePostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetData gets the data property value. The data property
func (m *ResumePostRequestBody) GetData()(i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.CustomTaskExtensionCallbackDataable) {
    return m.data
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ResumePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["data"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.CreateCustomTaskExtensionCallbackDataFromDiscriminatorValue , m.SetData)
    res["source"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSource)
    res["type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetType)
    return res
}
// GetSource gets the source property value. The source property
func (m *ResumePostRequestBody) GetSource()(*string) {
    return m.source
}
// GetType gets the type property value. The type property
func (m *ResumePostRequestBody) GetType()(*string) {
    return m.type_escaped
}
// Serialize serializes information the current object
func (m *ResumePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("data", m.GetData())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("source", m.GetSource())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type", m.GetType())
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
func (m *ResumePostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetData sets the data property value. The data property
func (m *ResumePostRequestBody) SetData(value i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.CustomTaskExtensionCallbackDataable)() {
    m.data = value
}
// SetSource sets the source property value. The source property
func (m *ResumePostRequestBody) SetSource(value *string)() {
    m.source = value
}
// SetType sets the type property value. The type property
func (m *ResumePostRequestBody) SetType(value *string)() {
    m.type_escaped = value
}
