package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MachineLearningDetectedSensitiveContent 
type MachineLearningDetectedSensitiveContent struct {
    DetectedSensitiveContent
    // The matchTolerance property
    matchTolerance *MlClassificationMatchTolerance
    // The modelVersion property
    modelVersion *string
}
// NewMachineLearningDetectedSensitiveContent instantiates a new MachineLearningDetectedSensitiveContent and sets the default values.
func NewMachineLearningDetectedSensitiveContent()(*MachineLearningDetectedSensitiveContent) {
    m := &MachineLearningDetectedSensitiveContent{
        DetectedSensitiveContent: *NewDetectedSensitiveContent(),
    }
    return m
}
// CreateMachineLearningDetectedSensitiveContentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMachineLearningDetectedSensitiveContentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMachineLearningDetectedSensitiveContent(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MachineLearningDetectedSensitiveContent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DetectedSensitiveContent.GetFieldDeserializers()
    res["matchTolerance"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseMlClassificationMatchTolerance , m.SetMatchTolerance)
    res["modelVersion"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetModelVersion)
    return res
}
// GetMatchTolerance gets the matchTolerance property value. The matchTolerance property
func (m *MachineLearningDetectedSensitiveContent) GetMatchTolerance()(*MlClassificationMatchTolerance) {
    return m.matchTolerance
}
// GetModelVersion gets the modelVersion property value. The modelVersion property
func (m *MachineLearningDetectedSensitiveContent) GetModelVersion()(*string) {
    return m.modelVersion
}
// Serialize serializes information the current object
func (m *MachineLearningDetectedSensitiveContent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DetectedSensitiveContent.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetMatchTolerance() != nil {
        cast := (*m.GetMatchTolerance()).String()
        err = writer.WriteStringValue("matchTolerance", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("modelVersion", m.GetModelVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetMatchTolerance sets the matchTolerance property value. The matchTolerance property
func (m *MachineLearningDetectedSensitiveContent) SetMatchTolerance(value *MlClassificationMatchTolerance)() {
    m.matchTolerance = value
}
// SetModelVersion sets the modelVersion property value. The modelVersion property
func (m *MachineLearningDetectedSensitiveContent) SetModelVersion(value *string)() {
    m.modelVersion = value
}
