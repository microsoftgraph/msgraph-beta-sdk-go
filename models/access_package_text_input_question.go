package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessPackageTextInputQuestion 
type AccessPackageTextInputQuestion struct {
    AccessPackageQuestion
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Indicates whether the answer will be in single or multiple line format.
    isSingleLineQuestion *bool
}
// NewAccessPackageTextInputQuestion instantiates a new AccessPackageTextInputQuestion and sets the default values.
func NewAccessPackageTextInputQuestion()(*AccessPackageTextInputQuestion) {
    m := &AccessPackageTextInputQuestion{
        AccessPackageQuestion: *NewAccessPackageQuestion(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAccessPackageTextInputQuestionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessPackageTextInputQuestionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessPackageTextInputQuestion(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AccessPackageTextInputQuestion) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessPackageTextInputQuestion) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AccessPackageQuestion.GetFieldDeserializers()
    res["isSingleLineQuestion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSingleLineQuestion(val)
        }
        return nil
    }
    return res
}
// GetIsSingleLineQuestion gets the isSingleLineQuestion property value. Indicates whether the answer will be in single or multiple line format.
func (m *AccessPackageTextInputQuestion) GetIsSingleLineQuestion()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isSingleLineQuestion
    }
}
// Serialize serializes information the current object
func (m *AccessPackageTextInputQuestion) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AccessPackageQuestion.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("isSingleLineQuestion", m.GetIsSingleLineQuestion())
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
func (m *AccessPackageTextInputQuestion) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetIsSingleLineQuestion sets the isSingleLineQuestion property value. Indicates whether the answer will be in single or multiple line format.
func (m *AccessPackageTextInputQuestion) SetIsSingleLineQuestion(value *bool)() {
    if m != nil {
        m.isSingleLineQuestion = value
    }
}
