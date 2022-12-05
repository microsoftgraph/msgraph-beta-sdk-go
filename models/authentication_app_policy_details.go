package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuthenticationAppPolicyDetails 
type AuthenticationAppPolicyDetails struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The adminConfiguration property
    adminConfiguration *AuthenticationAppAdminConfiguration
    // The authenticationEvaluation property
    authenticationEvaluation *AuthenticationAppEvaluation
    // The OdataType property
    odataType *string
    // The policyName property
    policyName *string
    // The status property
    status *AuthenticationAppPolicyStatus
}
// NewAuthenticationAppPolicyDetails instantiates a new authenticationAppPolicyDetails and sets the default values.
func NewAuthenticationAppPolicyDetails()(*AuthenticationAppPolicyDetails) {
    m := &AuthenticationAppPolicyDetails{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAuthenticationAppPolicyDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuthenticationAppPolicyDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthenticationAppPolicyDetails(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AuthenticationAppPolicyDetails) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAdminConfiguration gets the adminConfiguration property value. The adminConfiguration property
func (m *AuthenticationAppPolicyDetails) GetAdminConfiguration()(*AuthenticationAppAdminConfiguration) {
    return m.adminConfiguration
}
// GetAuthenticationEvaluation gets the authenticationEvaluation property value. The authenticationEvaluation property
func (m *AuthenticationAppPolicyDetails) GetAuthenticationEvaluation()(*AuthenticationAppEvaluation) {
    return m.authenticationEvaluation
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuthenticationAppPolicyDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["adminConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAuthenticationAppAdminConfiguration)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdminConfiguration(val.(*AuthenticationAppAdminConfiguration))
        }
        return nil
    }
    res["authenticationEvaluation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAuthenticationAppEvaluation)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationEvaluation(val.(*AuthenticationAppEvaluation))
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["policyName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicyName(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAuthenticationAppPolicyStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*AuthenticationAppPolicyStatus))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AuthenticationAppPolicyDetails) GetOdataType()(*string) {
    return m.odataType
}
// GetPolicyName gets the policyName property value. The policyName property
func (m *AuthenticationAppPolicyDetails) GetPolicyName()(*string) {
    return m.policyName
}
// GetStatus gets the status property value. The status property
func (m *AuthenticationAppPolicyDetails) GetStatus()(*AuthenticationAppPolicyStatus) {
    return m.status
}
// Serialize serializes information the current object
func (m *AuthenticationAppPolicyDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAdminConfiguration() != nil {
        cast := (*m.GetAdminConfiguration()).String()
        err := writer.WriteStringValue("adminConfiguration", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAuthenticationEvaluation() != nil {
        cast := (*m.GetAuthenticationEvaluation()).String()
        err := writer.WriteStringValue("authenticationEvaluation", &cast)
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
        err := writer.WriteStringValue("policyName", m.GetPolicyName())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err := writer.WriteStringValue("status", &cast)
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
func (m *AuthenticationAppPolicyDetails) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAdminConfiguration sets the adminConfiguration property value. The adminConfiguration property
func (m *AuthenticationAppPolicyDetails) SetAdminConfiguration(value *AuthenticationAppAdminConfiguration)() {
    m.adminConfiguration = value
}
// SetAuthenticationEvaluation sets the authenticationEvaluation property value. The authenticationEvaluation property
func (m *AuthenticationAppPolicyDetails) SetAuthenticationEvaluation(value *AuthenticationAppEvaluation)() {
    m.authenticationEvaluation = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AuthenticationAppPolicyDetails) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPolicyName sets the policyName property value. The policyName property
func (m *AuthenticationAppPolicyDetails) SetPolicyName(value *string)() {
    m.policyName = value
}
// SetStatus sets the status property value. The status property
func (m *AuthenticationAppPolicyDetails) SetStatus(value *AuthenticationAppPolicyStatus)() {
    m.status = value
}
