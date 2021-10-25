package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type B2cIdentityUserFlow struct {
    IdentityUserFlow
    apiConnectorConfiguration *UserFlowApiConnectorConfiguration;
    defaultLanguageTag *string;
    identityProviders []IdentityProvider;
    isLanguageCustomizationEnabled *bool;
    languages []UserFlowLanguageConfiguration;
    userAttributeAssignments []IdentityUserFlowAttributeAssignment;
    userFlowIdentityProviders []IdentityProviderBase;
}
func NewB2cIdentityUserFlow()(*B2cIdentityUserFlow) {
    m := &B2cIdentityUserFlow{
        IdentityUserFlow: *NewIdentityUserFlow(),
    }
    return m
}
func (m *B2cIdentityUserFlow) GetApiConnectorConfiguration()(*UserFlowApiConnectorConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.apiConnectorConfiguration
    }
}
func (m *B2cIdentityUserFlow) GetDefaultLanguageTag()(*string) {
    if m == nil {
        return nil
    } else {
        return m.defaultLanguageTag
    }
}
func (m *B2cIdentityUserFlow) GetIdentityProviders()([]IdentityProvider) {
    if m == nil {
        return nil
    } else {
        return m.identityProviders
    }
}
func (m *B2cIdentityUserFlow) GetIsLanguageCustomizationEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isLanguageCustomizationEnabled
    }
}
func (m *B2cIdentityUserFlow) GetLanguages()([]UserFlowLanguageConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.languages
    }
}
func (m *B2cIdentityUserFlow) GetUserAttributeAssignments()([]IdentityUserFlowAttributeAssignment) {
    if m == nil {
        return nil
    } else {
        return m.userAttributeAssignments
    }
}
func (m *B2cIdentityUserFlow) GetUserFlowIdentityProviders()([]IdentityProviderBase) {
    if m == nil {
        return nil
    } else {
        return m.userFlowIdentityProviders
    }
}
func (m *B2cIdentityUserFlow) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.IdentityUserFlow.GetFieldDeserializers()
    res["apiConnectorConfiguration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserFlowApiConnectorConfiguration() })
        if err != nil {
            return err
        }
        m.SetApiConnectorConfiguration(val.(*UserFlowApiConnectorConfiguration))
        return nil
    }
    res["defaultLanguageTag"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDefaultLanguageTag(val)
        return nil
    }
    res["identityProviders"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentityProvider() })
        if err != nil {
            return err
        }
        res := make([]IdentityProvider, len(val))
        for i, v := range val {
            res[i] = *(v.(*IdentityProvider))
        }
        m.SetIdentityProviders(res)
        return nil
    }
    res["isLanguageCustomizationEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsLanguageCustomizationEnabled(val)
        return nil
    }
    res["languages"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserFlowLanguageConfiguration() })
        if err != nil {
            return err
        }
        res := make([]UserFlowLanguageConfiguration, len(val))
        for i, v := range val {
            res[i] = *(v.(*UserFlowLanguageConfiguration))
        }
        m.SetLanguages(res)
        return nil
    }
    res["userAttributeAssignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentityUserFlowAttributeAssignment() })
        if err != nil {
            return err
        }
        res := make([]IdentityUserFlowAttributeAssignment, len(val))
        for i, v := range val {
            res[i] = *(v.(*IdentityUserFlowAttributeAssignment))
        }
        m.SetUserAttributeAssignments(res)
        return nil
    }
    res["userFlowIdentityProviders"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentityProviderBase() })
        if err != nil {
            return err
        }
        res := make([]IdentityProviderBase, len(val))
        for i, v := range val {
            res[i] = *(v.(*IdentityProviderBase))
        }
        m.SetUserFlowIdentityProviders(res)
        return nil
    }
    return res
}
func (m *B2cIdentityUserFlow) IsNil()(bool) {
    return m == nil
}
func (m *B2cIdentityUserFlow) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.IdentityUserFlow.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("apiConnectorConfiguration", m.GetApiConnectorConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("defaultLanguageTag", m.GetDefaultLanguageTag())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetIdentityProviders()))
        for i, v := range m.GetIdentityProviders() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("identityProviders", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isLanguageCustomizationEnabled", m.GetIsLanguageCustomizationEnabled())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetLanguages()))
        for i, v := range m.GetLanguages() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("languages", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserAttributeAssignments()))
        for i, v := range m.GetUserAttributeAssignments() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("userAttributeAssignments", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserFlowIdentityProviders()))
        for i, v := range m.GetUserFlowIdentityProviders() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("userFlowIdentityProviders", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *B2cIdentityUserFlow) SetApiConnectorConfiguration(value *UserFlowApiConnectorConfiguration)() {
    m.apiConnectorConfiguration = value
}
func (m *B2cIdentityUserFlow) SetDefaultLanguageTag(value *string)() {
    m.defaultLanguageTag = value
}
func (m *B2cIdentityUserFlow) SetIdentityProviders(value []IdentityProvider)() {
    m.identityProviders = value
}
func (m *B2cIdentityUserFlow) SetIsLanguageCustomizationEnabled(value *bool)() {
    m.isLanguageCustomizationEnabled = value
}
func (m *B2cIdentityUserFlow) SetLanguages(value []UserFlowLanguageConfiguration)() {
    m.languages = value
}
func (m *B2cIdentityUserFlow) SetUserAttributeAssignments(value []IdentityUserFlowAttributeAssignment)() {
    m.userAttributeAssignments = value
}
func (m *B2cIdentityUserFlow) SetUserFlowIdentityProviders(value []IdentityProviderBase)() {
    m.userFlowIdentityProviders = value
}
