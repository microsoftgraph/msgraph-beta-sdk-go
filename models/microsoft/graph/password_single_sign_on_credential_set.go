package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type PasswordSingleSignOnCredentialSet struct {
    additionalData map[string]interface{};
    credentials []Credential;
    id *string;
}
func NewPasswordSingleSignOnCredentialSet()(*PasswordSingleSignOnCredentialSet) {
    m := &PasswordSingleSignOnCredentialSet{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *PasswordSingleSignOnCredentialSet) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *PasswordSingleSignOnCredentialSet) GetCredentials()([]Credential) {
    if m == nil {
        return nil
    } else {
        return m.credentials
    }
}
func (m *PasswordSingleSignOnCredentialSet) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
func (m *PasswordSingleSignOnCredentialSet) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["credentials"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCredential() })
        if err != nil {
            return err
        }
        res := make([]Credential, len(val))
        for i, v := range val {
            res[i] = *(v.(*Credential))
        }
        m.SetCredentials(res)
        return nil
    }
    res["id"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetId(val)
        return nil
    }
    return res
}
func (m *PasswordSingleSignOnCredentialSet) IsNil()(bool) {
    return m == nil
}
func (m *PasswordSingleSignOnCredentialSet) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
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
func (m *PasswordSingleSignOnCredentialSet) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *PasswordSingleSignOnCredentialSet) SetCredentials(value []Credential)() {
    m.credentials = value
}
func (m *PasswordSingleSignOnCredentialSet) SetId(value *string)() {
    m.id = value
}
