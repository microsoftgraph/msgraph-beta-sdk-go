package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type MachineLearningDetectedSensitiveContent struct {
    DetectedSensitiveContent
}
// NewMachineLearningDetectedSensitiveContent instantiates a new MachineLearningDetectedSensitiveContent and sets the default values.
func NewMachineLearningDetectedSensitiveContent()(*MachineLearningDetectedSensitiveContent) {
    m := &MachineLearningDetectedSensitiveContent{
        DetectedSensitiveContent: *NewDetectedSensitiveContent(),
    }
    return m
}
// CreateMachineLearningDetectedSensitiveContentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMachineLearningDetectedSensitiveContentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMachineLearningDetectedSensitiveContent(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MachineLearningDetectedSensitiveContent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DetectedSensitiveContent.GetFieldDeserializers()
    res["matchTolerance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMlClassificationMatchTolerance)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMatchTolerance(val.(*MlClassificationMatchTolerance))
        }
        return nil
    }
    res["modelVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModelVersion(val)
        }
        return nil
    }
    return res
}
// GetMatchTolerance gets the matchTolerance property value. The matchTolerance property
// returns a *MlClassificationMatchTolerance when successful
func (m *MachineLearningDetectedSensitiveContent) GetMatchTolerance()(*MlClassificationMatchTolerance) {
    val, err := m.GetBackingStore().Get("matchTolerance")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*MlClassificationMatchTolerance)
    }
    return nil
}
// GetModelVersion gets the modelVersion property value. The modelVersion property
// returns a *string when successful
func (m *MachineLearningDetectedSensitiveContent) GetModelVersion()(*string) {
    val, err := m.GetBackingStore().Get("modelVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
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
    err := m.GetBackingStore().Set("matchTolerance", value)
    if err != nil {
        panic(err)
    }
}
// SetModelVersion sets the modelVersion property value. The modelVersion property
func (m *MachineLearningDetectedSensitiveContent) SetModelVersion(value *string)() {
    err := m.GetBackingStore().Set("modelVersion", value)
    if err != nil {
        panic(err)
    }
}
type MachineLearningDetectedSensitiveContentable interface {
    DetectedSensitiveContentable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMatchTolerance()(*MlClassificationMatchTolerance)
    GetModelVersion()(*string)
    SetMatchTolerance(value *MlClassificationMatchTolerance)()
    SetModelVersion(value *string)()
}
