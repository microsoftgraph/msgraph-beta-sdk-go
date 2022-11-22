package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CustomQuestionAnswer 
type CustomQuestionAnswer struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Display name of the custom registration question. Read-only.
    displayName *string
    // The OdataType property
    odataType *string
    // ID the custom registration question. Read-only.
    questionId *string
    // Answer to the custom registration question.
    value *string
}
// NewCustomQuestionAnswer instantiates a new customQuestionAnswer and sets the default values.
func NewCustomQuestionAnswer()(*CustomQuestionAnswer) {
    m := &CustomQuestionAnswer{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCustomQuestionAnswerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCustomQuestionAnswerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCustomQuestionAnswer(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CustomQuestionAnswer) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDisplayName gets the displayName property value. Display name of the custom registration question. Read-only.
func (m *CustomQuestionAnswer) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CustomQuestionAnswer) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["questionId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetQuestionId)
    res["value"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetValue)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *CustomQuestionAnswer) GetOdataType()(*string) {
    return m.odataType
}
// GetQuestionId gets the questionId property value. ID the custom registration question. Read-only.
func (m *CustomQuestionAnswer) GetQuestionId()(*string) {
    return m.questionId
}
// GetValue gets the value property value. Answer to the custom registration question.
func (m *CustomQuestionAnswer) GetValue()(*string) {
    return m.value
}
// Serialize serializes information the current object
func (m *CustomQuestionAnswer) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("questionId", m.GetQuestionId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("value", m.GetValue())
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
func (m *CustomQuestionAnswer) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDisplayName sets the displayName property value. Display name of the custom registration question. Read-only.
func (m *CustomQuestionAnswer) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CustomQuestionAnswer) SetOdataType(value *string)() {
    m.odataType = value
}
// SetQuestionId sets the questionId property value. ID the custom registration question. Read-only.
func (m *CustomQuestionAnswer) SetQuestionId(value *string)() {
    m.questionId = value
}
// SetValue sets the value property value. Answer to the custom registration question.
func (m *CustomQuestionAnswer) SetValue(value *string)() {
    m.value = value
}
