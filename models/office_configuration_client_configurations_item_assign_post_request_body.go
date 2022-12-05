package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OfficeConfigurationClientConfigurationsItemAssignPostRequestBody provides operations to call the assign method.
type OfficeConfigurationClientConfigurationsItemAssignPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The officeConfigurationAssignments property
    officeConfigurationAssignments []OfficeClientConfigurationAssignmentable
}
// NewOfficeConfigurationClientConfigurationsItemAssignPostRequestBody instantiates a new OfficeConfigurationClientConfigurationsItemAssignPostRequestBody and sets the default values.
func NewOfficeConfigurationClientConfigurationsItemAssignPostRequestBody()(*OfficeConfigurationClientConfigurationsItemAssignPostRequestBody) {
    m := &OfficeConfigurationClientConfigurationsItemAssignPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateOfficeConfigurationClientConfigurationsItemAssignPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOfficeConfigurationClientConfigurationsItemAssignPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOfficeConfigurationClientConfigurationsItemAssignPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OfficeConfigurationClientConfigurationsItemAssignPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OfficeConfigurationClientConfigurationsItemAssignPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["officeConfigurationAssignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOfficeClientConfigurationAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OfficeClientConfigurationAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(OfficeClientConfigurationAssignmentable)
            }
            m.SetOfficeConfigurationAssignments(res)
        }
        return nil
    }
    return res
}
// GetOfficeConfigurationAssignments gets the officeConfigurationAssignments property value. The officeConfigurationAssignments property
func (m *OfficeConfigurationClientConfigurationsItemAssignPostRequestBody) GetOfficeConfigurationAssignments()([]OfficeClientConfigurationAssignmentable) {
    return m.officeConfigurationAssignments
}
// Serialize serializes information the current object
func (m *OfficeConfigurationClientConfigurationsItemAssignPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetOfficeConfigurationAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOfficeConfigurationAssignments()))
        for i, v := range m.GetOfficeConfigurationAssignments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("officeConfigurationAssignments", cast)
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
func (m *OfficeConfigurationClientConfigurationsItemAssignPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetOfficeConfigurationAssignments sets the officeConfigurationAssignments property value. The officeConfigurationAssignments property
func (m *OfficeConfigurationClientConfigurationsItemAssignPostRequestBody) SetOfficeConfigurationAssignments(value []OfficeClientConfigurationAssignmentable)() {
    m.officeConfigurationAssignments = value
}
