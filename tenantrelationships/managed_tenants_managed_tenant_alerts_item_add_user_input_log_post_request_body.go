package tenantrelationships

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedTenantsManagedTenantAlertsItemAddUserInputLogPostRequestBody 
type ManagedTenantsManagedTenantAlertsItemAddUserInputLogPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The logInformation property
    logInformation *string
}
// NewManagedTenantsManagedTenantAlertsItemAddUserInputLogPostRequestBody instantiates a new ManagedTenantsManagedTenantAlertsItemAddUserInputLogPostRequestBody and sets the default values.
func NewManagedTenantsManagedTenantAlertsItemAddUserInputLogPostRequestBody()(*ManagedTenantsManagedTenantAlertsItemAddUserInputLogPostRequestBody) {
    m := &ManagedTenantsManagedTenantAlertsItemAddUserInputLogPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateManagedTenantsManagedTenantAlertsItemAddUserInputLogPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedTenantsManagedTenantAlertsItemAddUserInputLogPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedTenantsManagedTenantAlertsItemAddUserInputLogPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ManagedTenantsManagedTenantAlertsItemAddUserInputLogPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedTenantsManagedTenantAlertsItemAddUserInputLogPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["logInformation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLogInformation(val)
        }
        return nil
    }
    return res
}
// GetLogInformation gets the logInformation property value. The logInformation property
func (m *ManagedTenantsManagedTenantAlertsItemAddUserInputLogPostRequestBody) GetLogInformation()(*string) {
    return m.logInformation
}
// Serialize serializes information the current object
func (m *ManagedTenantsManagedTenantAlertsItemAddUserInputLogPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("logInformation", m.GetLogInformation())
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
func (m *ManagedTenantsManagedTenantAlertsItemAddUserInputLogPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetLogInformation sets the logInformation property value. The logInformation property
func (m *ManagedTenantsManagedTenantAlertsItemAddUserInputLogPostRequestBody) SetLogInformation(value *string)() {
    m.logInformation = value
}
