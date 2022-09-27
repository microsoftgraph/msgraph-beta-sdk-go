package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ServicePrincipalLockConfiguration 
type ServicePrincipalLockConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The allProperties property
    allProperties *bool
    // The credentialsWithUsageSign property
    credentialsWithUsageSign *bool
    // The credentialsWithUsageVerify property
    credentialsWithUsageVerify *bool
    // The isEnabled property
    isEnabled *bool
    // The OdataType property
    odataType *string
    // The tokenEncryptionKeyId property
    tokenEncryptionKeyId *bool
}
// NewServicePrincipalLockConfiguration instantiates a new servicePrincipalLockConfiguration and sets the default values.
func NewServicePrincipalLockConfiguration()(*ServicePrincipalLockConfiguration) {
    m := &ServicePrincipalLockConfiguration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.servicePrincipalLockConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateServicePrincipalLockConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateServicePrincipalLockConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServicePrincipalLockConfiguration(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ServicePrincipalLockConfiguration) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAllProperties gets the allProperties property value. The allProperties property
func (m *ServicePrincipalLockConfiguration) GetAllProperties()(*bool) {
    return m.allProperties
}
// GetCredentialsWithUsageSign gets the credentialsWithUsageSign property value. The credentialsWithUsageSign property
func (m *ServicePrincipalLockConfiguration) GetCredentialsWithUsageSign()(*bool) {
    return m.credentialsWithUsageSign
}
// GetCredentialsWithUsageVerify gets the credentialsWithUsageVerify property value. The credentialsWithUsageVerify property
func (m *ServicePrincipalLockConfiguration) GetCredentialsWithUsageVerify()(*bool) {
    return m.credentialsWithUsageVerify
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ServicePrincipalLockConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allProperties"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetAllProperties)
    res["credentialsWithUsageSign"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetCredentialsWithUsageSign)
    res["credentialsWithUsageVerify"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetCredentialsWithUsageVerify)
    res["isEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsEnabled)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["tokenEncryptionKeyId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetTokenEncryptionKeyId)
    return res
}
// GetIsEnabled gets the isEnabled property value. The isEnabled property
func (m *ServicePrincipalLockConfiguration) GetIsEnabled()(*bool) {
    return m.isEnabled
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ServicePrincipalLockConfiguration) GetOdataType()(*string) {
    return m.odataType
}
// GetTokenEncryptionKeyId gets the tokenEncryptionKeyId property value. The tokenEncryptionKeyId property
func (m *ServicePrincipalLockConfiguration) GetTokenEncryptionKeyId()(*bool) {
    return m.tokenEncryptionKeyId
}
// Serialize serializes information the current object
func (m *ServicePrincipalLockConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allProperties", m.GetAllProperties())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("credentialsWithUsageSign", m.GetCredentialsWithUsageSign())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("credentialsWithUsageVerify", m.GetCredentialsWithUsageVerify())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isEnabled", m.GetIsEnabled())
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
        err := writer.WriteBoolValue("tokenEncryptionKeyId", m.GetTokenEncryptionKeyId())
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
func (m *ServicePrincipalLockConfiguration) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAllProperties sets the allProperties property value. The allProperties property
func (m *ServicePrincipalLockConfiguration) SetAllProperties(value *bool)() {
    m.allProperties = value
}
// SetCredentialsWithUsageSign sets the credentialsWithUsageSign property value. The credentialsWithUsageSign property
func (m *ServicePrincipalLockConfiguration) SetCredentialsWithUsageSign(value *bool)() {
    m.credentialsWithUsageSign = value
}
// SetCredentialsWithUsageVerify sets the credentialsWithUsageVerify property value. The credentialsWithUsageVerify property
func (m *ServicePrincipalLockConfiguration) SetCredentialsWithUsageVerify(value *bool)() {
    m.credentialsWithUsageVerify = value
}
// SetIsEnabled sets the isEnabled property value. The isEnabled property
func (m *ServicePrincipalLockConfiguration) SetIsEnabled(value *bool)() {
    m.isEnabled = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ServicePrincipalLockConfiguration) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTokenEncryptionKeyId sets the tokenEncryptionKeyId property value. The tokenEncryptionKeyId property
func (m *ServicePrincipalLockConfiguration) SetTokenEncryptionKeyId(value *bool)() {
    m.tokenEncryptionKeyId = value
}
