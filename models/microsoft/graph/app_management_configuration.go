package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AppManagementConfiguration 
type AppManagementConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Collection of keyCredential restrictions settings to be applied to an application or service principal.
    keyCredentials []KeyCredentialConfigurationable;
    // Collection of password restrictions settings to be applied to an application or service principal.
    passwordCredentials []PasswordCredentialConfigurationable;
}
// NewAppManagementConfiguration instantiates a new appManagementConfiguration and sets the default values.
func NewAppManagementConfiguration()(*AppManagementConfiguration) {
    m := &AppManagementConfiguration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAppManagementConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAppManagementConfigurationFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewAppManagementConfiguration(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AppManagementConfiguration) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AppManagementConfiguration) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["keyCredentials"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateKeyCredentialConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]KeyCredentialConfigurationable, len(val))
            for i, v := range val {
                res[i] = v.(KeyCredentialConfigurationable)
            }
            m.SetKeyCredentials(res)
        }
        return nil
    }
    res["passwordCredentials"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePasswordCredentialConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PasswordCredentialConfigurationable, len(val))
            for i, v := range val {
                res[i] = v.(PasswordCredentialConfigurationable)
            }
            m.SetPasswordCredentials(res)
        }
        return nil
    }
    return res
}
// GetKeyCredentials gets the keyCredentials property value. Collection of keyCredential restrictions settings to be applied to an application or service principal.
func (m *AppManagementConfiguration) GetKeyCredentials()([]KeyCredentialConfigurationable) {
    if m == nil {
        return nil
    } else {
        return m.keyCredentials
    }
}
// GetPasswordCredentials gets the passwordCredentials property value. Collection of password restrictions settings to be applied to an application or service principal.
func (m *AppManagementConfiguration) GetPasswordCredentials()([]PasswordCredentialConfigurationable) {
    if m == nil {
        return nil
    } else {
        return m.passwordCredentials
    }
}
// Serialize serializes information the current object
func (m *AppManagementConfiguration) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetKeyCredentials() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetKeyCredentials()))
        for i, v := range m.GetKeyCredentials() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("keyCredentials", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPasswordCredentials() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPasswordCredentials()))
        for i, v := range m.GetPasswordCredentials() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("passwordCredentials", cast)
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
func (m *AppManagementConfiguration) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetKeyCredentials sets the keyCredentials property value. Collection of keyCredential restrictions settings to be applied to an application or service principal.
func (m *AppManagementConfiguration) SetKeyCredentials(value []KeyCredentialConfigurationable)() {
    if m != nil {
        m.keyCredentials = value
    }
}
// SetPasswordCredentials sets the passwordCredentials property value. Collection of password restrictions settings to be applied to an application or service principal.
func (m *AppManagementConfiguration) SetPasswordCredentials(value []PasswordCredentialConfigurationable)() {
    if m != nil {
        m.passwordCredentials = value
    }
}
