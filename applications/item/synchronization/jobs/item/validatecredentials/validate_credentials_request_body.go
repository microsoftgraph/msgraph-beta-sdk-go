package validatecredentials

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// ValidateCredentialsRequestBody 
type ValidateCredentialsRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    applicationIdentifier *string;
    // 
    credentials []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SynchronizationSecretKeyStringValuePair;
    // 
    templateId *string;
    // 
    useSavedCredentials *bool;
}
// NewValidateCredentialsRequestBody instantiates a new validateCredentialsRequestBody and sets the default values.
func NewValidateCredentialsRequestBody()(*ValidateCredentialsRequestBody) {
    m := &ValidateCredentialsRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ValidateCredentialsRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetApplicationIdentifier gets the applicationIdentifier property value. 
func (m *ValidateCredentialsRequestBody) GetApplicationIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.applicationIdentifier
    }
}
// GetCredentials gets the credentials property value. 
func (m *ValidateCredentialsRequestBody) GetCredentials()([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SynchronizationSecretKeyStringValuePair) {
    if m == nil {
        return nil
    } else {
        return m.credentials
    }
}
// GetTemplateId gets the templateId property value. 
func (m *ValidateCredentialsRequestBody) GetTemplateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.templateId
    }
}
// GetUseSavedCredentials gets the useSavedCredentials property value. 
func (m *ValidateCredentialsRequestBody) GetUseSavedCredentials()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.useSavedCredentials
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ValidateCredentialsRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["applicationIdentifier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationIdentifier(val)
        }
        return nil
    }
    res["credentials"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewSynchronizationSecretKeyStringValuePair() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SynchronizationSecretKeyStringValuePair, len(val))
            for i, v := range val {
                res[i] = *(v.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SynchronizationSecretKeyStringValuePair))
            }
            m.SetCredentials(res)
        }
        return nil
    }
    res["templateId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTemplateId(val)
        }
        return nil
    }
    res["useSavedCredentials"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUseSavedCredentials(val)
        }
        return nil
    }
    return res
}
func (m *ValidateCredentialsRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ValidateCredentialsRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("applicationIdentifier", m.GetApplicationIdentifier())
        if err != nil {
            return err
        }
    }
    if m.GetCredentials() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCredentials()))
        for i, v := range m.GetCredentials() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("credentials", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("templateId", m.GetTemplateId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("useSavedCredentials", m.GetUseSavedCredentials())
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
func (m *ValidateCredentialsRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetApplicationIdentifier sets the applicationIdentifier property value. 
func (m *ValidateCredentialsRequestBody) SetApplicationIdentifier(value *string)() {
    if m != nil {
        m.applicationIdentifier = value
    }
}
// SetCredentials sets the credentials property value. 
func (m *ValidateCredentialsRequestBody) SetCredentials(value []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SynchronizationSecretKeyStringValuePair)() {
    if m != nil {
        m.credentials = value
    }
}
// SetTemplateId sets the templateId property value. 
func (m *ValidateCredentialsRequestBody) SetTemplateId(value *string)() {
    if m != nil {
        m.templateId = value
    }
}
// SetUseSavedCredentials sets the useSavedCredentials property value. 
func (m *ValidateCredentialsRequestBody) SetUseSavedCredentials(value *bool)() {
    if m != nil {
        m.useSavedCredentials = value
    }
}
