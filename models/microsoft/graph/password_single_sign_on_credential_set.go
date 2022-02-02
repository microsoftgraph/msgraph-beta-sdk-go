package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PasswordSingleSignOnCredentialSet 
type PasswordSingleSignOnCredentialSet struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // A list of credential objects that define the complete sign in flow.
    credentials []Credential;
    // The ID of the user or group this credential set belongs to.
    id *string;
}
// NewPasswordSingleSignOnCredentialSet instantiates a new passwordSingleSignOnCredentialSet and sets the default values.
func NewPasswordSingleSignOnCredentialSet()(*PasswordSingleSignOnCredentialSet) {
    m := &PasswordSingleSignOnCredentialSet{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PasswordSingleSignOnCredentialSet) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCredentials gets the credentials property value. A list of credential objects that define the complete sign in flow.
func (m *PasswordSingleSignOnCredentialSet) GetCredentials()([]Credential) {
    if m == nil {
        return nil
    } else {
        return m.credentials
    }
}
// GetId gets the id property value. The ID of the user or group this credential set belongs to.
func (m *PasswordSingleSignOnCredentialSet) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PasswordSingleSignOnCredentialSet) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["credentials"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCredential() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Credential, len(val))
            for i, v := range val {
                res[i] = *(v.(*Credential))
            }
            m.SetCredentials(res)
        }
        return nil
    }
    res["id"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    return res
}
func (m *PasswordSingleSignOnCredentialSet) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *PasswordSingleSignOnCredentialSet) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err := writer.WriteStringValue("id", m.GetId())
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
func (m *PasswordSingleSignOnCredentialSet) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCredentials sets the credentials property value. A list of credential objects that define the complete sign in flow.
func (m *PasswordSingleSignOnCredentialSet) SetCredentials(value []Credential)() {
    if m != nil {
        m.credentials = value
    }
}
// SetId sets the id property value. The ID of the user or group this credential set belongs to.
func (m *PasswordSingleSignOnCredentialSet) SetId(value *string)() {
    if m != nil {
        m.id = value
    }
}
