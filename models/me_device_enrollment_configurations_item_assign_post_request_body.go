package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeDeviceEnrollmentConfigurationsItemAssignPostRequestBody provides operations to call the assign method.
type MeDeviceEnrollmentConfigurationsItemAssignPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The enrollmentConfigurationAssignments property
    enrollmentConfigurationAssignments []EnrollmentConfigurationAssignmentable
}
// NewMeDeviceEnrollmentConfigurationsItemAssignPostRequestBody instantiates a new MeDeviceEnrollmentConfigurationsItemAssignPostRequestBody and sets the default values.
func NewMeDeviceEnrollmentConfigurationsItemAssignPostRequestBody()(*MeDeviceEnrollmentConfigurationsItemAssignPostRequestBody) {
    m := &MeDeviceEnrollmentConfigurationsItemAssignPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMeDeviceEnrollmentConfigurationsItemAssignPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMeDeviceEnrollmentConfigurationsItemAssignPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMeDeviceEnrollmentConfigurationsItemAssignPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MeDeviceEnrollmentConfigurationsItemAssignPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetEnrollmentConfigurationAssignments gets the enrollmentConfigurationAssignments property value. The enrollmentConfigurationAssignments property
func (m *MeDeviceEnrollmentConfigurationsItemAssignPostRequestBody) GetEnrollmentConfigurationAssignments()([]EnrollmentConfigurationAssignmentable) {
    return m.enrollmentConfigurationAssignments
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeDeviceEnrollmentConfigurationsItemAssignPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["enrollmentConfigurationAssignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEnrollmentConfigurationAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EnrollmentConfigurationAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(EnrollmentConfigurationAssignmentable)
            }
            m.SetEnrollmentConfigurationAssignments(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *MeDeviceEnrollmentConfigurationsItemAssignPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetEnrollmentConfigurationAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEnrollmentConfigurationAssignments()))
        for i, v := range m.GetEnrollmentConfigurationAssignments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("enrollmentConfigurationAssignments", cast)
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
func (m *MeDeviceEnrollmentConfigurationsItemAssignPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetEnrollmentConfigurationAssignments sets the enrollmentConfigurationAssignments property value. The enrollmentConfigurationAssignments property
func (m *MeDeviceEnrollmentConfigurationsItemAssignPostRequestBody) SetEnrollmentConfigurationAssignments(value []EnrollmentConfigurationAssignmentable)() {
    m.enrollmentConfigurationAssignments = value
}
