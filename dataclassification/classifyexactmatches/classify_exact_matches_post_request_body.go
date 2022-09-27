package classifyexactmatches

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ClassifyExactMatchesPostRequestBody provides operations to call the classifyExactMatches method.
type ClassifyExactMatchesPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The contentClassifications property
    contentClassifications []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentClassificationable
    // The sensitiveTypeIds property
    sensitiveTypeIds []string
    // The text property
    text *string
    // The timeoutInMs property
    timeoutInMs *string
}
// NewClassifyExactMatchesPostRequestBody instantiates a new classifyExactMatchesPostRequestBody and sets the default values.
func NewClassifyExactMatchesPostRequestBody()(*ClassifyExactMatchesPostRequestBody) {
    m := &ClassifyExactMatchesPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateClassifyExactMatchesPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateClassifyExactMatchesPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewClassifyExactMatchesPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ClassifyExactMatchesPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetContentClassifications gets the contentClassifications property value. The contentClassifications property
func (m *ClassifyExactMatchesPostRequestBody) GetContentClassifications()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentClassificationable) {
    return m.contentClassifications
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ClassifyExactMatchesPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["contentClassifications"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateContentClassificationFromDiscriminatorValue , m.SetContentClassifications)
    res["sensitiveTypeIds"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetSensitiveTypeIds)
    res["text"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetText)
    res["timeoutInMs"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetTimeoutInMs)
    return res
}
// GetSensitiveTypeIds gets the sensitiveTypeIds property value. The sensitiveTypeIds property
func (m *ClassifyExactMatchesPostRequestBody) GetSensitiveTypeIds()([]string) {
    return m.sensitiveTypeIds
}
// GetText gets the text property value. The text property
func (m *ClassifyExactMatchesPostRequestBody) GetText()(*string) {
    return m.text
}
// GetTimeoutInMs gets the timeoutInMs property value. The timeoutInMs property
func (m *ClassifyExactMatchesPostRequestBody) GetTimeoutInMs()(*string) {
    return m.timeoutInMs
}
// Serialize serializes information the current object
func (m *ClassifyExactMatchesPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetContentClassifications() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetContentClassifications())
        err := writer.WriteCollectionOfObjectValues("contentClassifications", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSensitiveTypeIds() != nil {
        err := writer.WriteCollectionOfStringValues("sensitiveTypeIds", m.GetSensitiveTypeIds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("text", m.GetText())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("timeoutInMs", m.GetTimeoutInMs())
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
func (m *ClassifyExactMatchesPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetContentClassifications sets the contentClassifications property value. The contentClassifications property
func (m *ClassifyExactMatchesPostRequestBody) SetContentClassifications(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentClassificationable)() {
    m.contentClassifications = value
}
// SetSensitiveTypeIds sets the sensitiveTypeIds property value. The sensitiveTypeIds property
func (m *ClassifyExactMatchesPostRequestBody) SetSensitiveTypeIds(value []string)() {
    m.sensitiveTypeIds = value
}
// SetText sets the text property value. The text property
func (m *ClassifyExactMatchesPostRequestBody) SetText(value *string)() {
    m.text = value
}
// SetTimeoutInMs sets the timeoutInMs property value. The timeoutInMs property
func (m *ClassifyExactMatchesPostRequestBody) SetTimeoutInMs(value *string)() {
    m.timeoutInMs = value
}
