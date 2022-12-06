package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EmployeeExperience 
type EmployeeExperience struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // A collection of learning providers.
    learningProviders []LearningProviderable
    // The OdataType property
    odataType *string
}
// NewEmployeeExperience instantiates a new EmployeeExperience and sets the default values.
func NewEmployeeExperience()(*EmployeeExperience) {
    m := &EmployeeExperience{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateEmployeeExperienceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEmployeeExperienceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEmployeeExperience(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EmployeeExperience) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EmployeeExperience) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["learningProviders"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateLearningProviderFromDiscriminatorValue , m.SetLearningProviders)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    return res
}
// GetLearningProviders gets the learningProviders property value. A collection of learning providers.
func (m *EmployeeExperience) GetLearningProviders()([]LearningProviderable) {
    return m.learningProviders
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *EmployeeExperience) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *EmployeeExperience) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetLearningProviders() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetLearningProviders())
        err := writer.WriteCollectionOfObjectValues("learningProviders", cast)
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EmployeeExperience) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetLearningProviders sets the learningProviders property value. A collection of learning providers.
func (m *EmployeeExperience) SetLearningProviders(value []LearningProviderable)() {
    m.learningProviders = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *EmployeeExperience) SetOdataType(value *string)() {
    m.odataType = value
}
