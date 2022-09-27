package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuthenticationConfigurationValidation 
type AuthenticationConfigurationValidation struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The errors property
    errors []GenericErrorable
    // The OdataType property
    odataType *string
    // The warnings property
    warnings []GenericErrorable
}
// NewAuthenticationConfigurationValidation instantiates a new authenticationConfigurationValidation and sets the default values.
func NewAuthenticationConfigurationValidation()(*AuthenticationConfigurationValidation) {
    m := &AuthenticationConfigurationValidation{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.authenticationConfigurationValidation";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAuthenticationConfigurationValidationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuthenticationConfigurationValidationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthenticationConfigurationValidation(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AuthenticationConfigurationValidation) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetErrors gets the errors property value. The errors property
func (m *AuthenticationConfigurationValidation) GetErrors()([]GenericErrorable) {
    return m.errors
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuthenticationConfigurationValidation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["errors"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateGenericErrorFromDiscriminatorValue , m.SetErrors)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["warnings"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateGenericErrorFromDiscriminatorValue , m.SetWarnings)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AuthenticationConfigurationValidation) GetOdataType()(*string) {
    return m.odataType
}
// GetWarnings gets the warnings property value. The warnings property
func (m *AuthenticationConfigurationValidation) GetWarnings()([]GenericErrorable) {
    return m.warnings
}
// Serialize serializes information the current object
func (m *AuthenticationConfigurationValidation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetErrors() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetErrors())
        err := writer.WriteCollectionOfObjectValues("errors", cast)
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
    if m.GetWarnings() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetWarnings())
        err := writer.WriteCollectionOfObjectValues("warnings", cast)
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
func (m *AuthenticationConfigurationValidation) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetErrors sets the errors property value. The errors property
func (m *AuthenticationConfigurationValidation) SetErrors(value []GenericErrorable)() {
    m.errors = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AuthenticationConfigurationValidation) SetOdataType(value *string)() {
    m.odataType = value
}
// SetWarnings sets the warnings property value. The warnings property
func (m *AuthenticationConfigurationValidation) SetWarnings(value []GenericErrorable)() {
    m.warnings = value
}
