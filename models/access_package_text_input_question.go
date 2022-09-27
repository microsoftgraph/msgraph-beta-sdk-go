package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessPackageTextInputQuestion 
type AccessPackageTextInputQuestion struct {
    AccessPackageQuestion
    // Indicates whether the answer will be in single or multiple line format.
    isSingleLineQuestion *bool
    // This is the regex pattern that the corresponding text answer must follow.
    regexPattern *string
}
// NewAccessPackageTextInputQuestion instantiates a new AccessPackageTextInputQuestion and sets the default values.
func NewAccessPackageTextInputQuestion()(*AccessPackageTextInputQuestion) {
    m := &AccessPackageTextInputQuestion{
        AccessPackageQuestion: *NewAccessPackageQuestion(),
    }
    odataTypeValue := "#microsoft.graph.accessPackageTextInputQuestion";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAccessPackageTextInputQuestionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessPackageTextInputQuestionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessPackageTextInputQuestion(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessPackageTextInputQuestion) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AccessPackageQuestion.GetFieldDeserializers()
    res["isSingleLineQuestion"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsSingleLineQuestion)
    res["regexPattern"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetRegexPattern)
    return res
}
// GetIsSingleLineQuestion gets the isSingleLineQuestion property value. Indicates whether the answer will be in single or multiple line format.
func (m *AccessPackageTextInputQuestion) GetIsSingleLineQuestion()(*bool) {
    return m.isSingleLineQuestion
}
// GetRegexPattern gets the regexPattern property value. This is the regex pattern that the corresponding text answer must follow.
func (m *AccessPackageTextInputQuestion) GetRegexPattern()(*string) {
    return m.regexPattern
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
        err = writer.WriteStringValue("regexPattern", m.GetRegexPattern())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIsSingleLineQuestion sets the isSingleLineQuestion property value. Indicates whether the answer will be in single or multiple line format.
func (m *AccessPackageTextInputQuestion) SetIsSingleLineQuestion(value *bool)() {
    m.isSingleLineQuestion = value
}
// SetRegexPattern sets the regexPattern property value. This is the regex pattern that the corresponding text answer must follow.
func (m *AccessPackageTextInputQuestion) SetRegexPattern(value *string)() {
    m.regexPattern = value
}
