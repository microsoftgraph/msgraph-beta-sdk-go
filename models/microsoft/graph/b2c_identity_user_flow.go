package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type B2cIdentityUserFlow struct {
    IdentityUserFlow
    // Configuration for enabling an API connector for use as part of the user flow. You can only obtain the value of this object using Get userFlowApiConnectorConfiguration.
    apiConnectorConfiguration *UserFlowApiConnectorConfiguration;
    // Indicates the default language of the b2cIdentityUserFlow that is used when no ui_locale tag is specified in the request. This field is RFC 5646 compliant.
    defaultLanguageTag *string;
    // 
    identityProviders []IdentityProvider;
    // The property that determines whether language customization is enabled within the B2C user flow. Language customization is not enabled by default for B2C user flows.
    isLanguageCustomizationEnabled *bool;
    // The languages supported for customization within the user flow. Language customization is not enabled by default in B2C user flows.
    languages []UserFlowLanguageConfiguration;
    // The user attribute assignments included in the user flow.
    userAttributeAssignments []IdentityUserFlowAttributeAssignment;
    // 
    userFlowIdentityProviders []IdentityProviderBase;
}
// Instantiates a new b2cIdentityUserFlow and sets the default values.
func NewB2cIdentityUserFlow()(*B2cIdentityUserFlow) {
    m := &B2cIdentityUserFlow{
        IdentityUserFlow: *NewIdentityUserFlow(),
    }
    return m
}
// Gets the apiConnectorConfiguration property value. Configuration for enabling an API connector for use as part of the user flow. You can only obtain the value of this object using Get userFlowApiConnectorConfiguration.
func (m *B2cIdentityUserFlow) GetApiConnectorConfiguration()(*UserFlowApiConnectorConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.apiConnectorConfiguration
    }
}
// Gets the defaultLanguageTag property value. Indicates the default language of the b2cIdentityUserFlow that is used when no ui_locale tag is specified in the request. This field is RFC 5646 compliant.
func (m *B2cIdentityUserFlow) GetDefaultLanguageTag()(*string) {
    if m == nil {
        return nil
    } else {
        return m.defaultLanguageTag
    }
}
// Gets the identityProviders property value. 
func (m *B2cIdentityUserFlow) GetIdentityProviders()([]IdentityProvider) {
    if m == nil {
        return nil
    } else {
        return m.identityProviders
    }
}
// Gets the isLanguageCustomizationEnabled property value. The property that determines whether language customization is enabled within the B2C user flow. Language customization is not enabled by default for B2C user flows.
func (m *B2cIdentityUserFlow) GetIsLanguageCustomizationEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isLanguageCustomizationEnabled
    }
}
// Gets the languages property value. The languages supported for customization within the user flow. Language customization is not enabled by default in B2C user flows.
func (m *B2cIdentityUserFlow) GetLanguages()([]UserFlowLanguageConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.languages
    }
}
// Gets the userAttributeAssignments property value. The user attribute assignments included in the user flow.
func (m *B2cIdentityUserFlow) GetUserAttributeAssignments()([]IdentityUserFlowAttributeAssignment) {
    if m == nil {
        return nil
    } else {
        return m.userAttributeAssignments
    }
}
// Gets the userFlowIdentityProviders property value. 
func (m *B2cIdentityUserFlow) GetUserFlowIdentityProviders()([]IdentityProviderBase) {
    if m == nil {
        return nil
    } else {
        return m.userFlowIdentityProviders
    }
}
// The deserialization information for the current model
func (m *B2cIdentityUserFlow) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.IdentityUserFlow.GetFieldDeserializers()
    res["apiConnectorConfiguration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserFlowApiConnectorConfiguration() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApiConnectorConfiguration(val.(*UserFlowApiConnectorConfiguration))
        }
        return nil
    }
    res["defaultLanguageTag"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultLanguageTag(val)
        }
        return nil
    }
    res["identityProviders"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentityProvider() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IdentityProvider, len(val))
            for i, v := range val {
                res[i] = *(v.(*IdentityProvider))
            }
            m.SetIdentityProviders(res)
        }
        return nil
    }
    res["isLanguageCustomizationEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsLanguageCustomizationEnabled(val)
        }
        return nil
    }
    res["languages"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserFlowLanguageConfiguration() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserFlowLanguageConfiguration, len(val))
            for i, v := range val {
                res[i] = *(v.(*UserFlowLanguageConfiguration))
            }
            m.SetLanguages(res)
        }
        return nil
    }
    res["userAttributeAssignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentityUserFlowAttributeAssignment() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IdentityUserFlowAttributeAssignment, len(val))
            for i, v := range val {
                res[i] = *(v.(*IdentityUserFlowAttributeAssignment))
            }
            m.SetUserAttributeAssignments(res)
        }
        return nil
    }
    res["userFlowIdentityProviders"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentityProviderBase() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IdentityProviderBase, len(val))
            for i, v := range val {
                res[i] = *(v.(*IdentityProviderBase))
            }
            m.SetUserFlowIdentityProviders(res)
        }
        return nil
    }
    return res
}
func (m *B2cIdentityUserFlow) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the apiConnectorConfiguration property value. Configuration for enabling an API connector for use as part of the user flow. You can only obtain the value of this object using Get userFlowApiConnectorConfiguration.
// Parameters:
//  - value : Value to set for the apiConnectorConfiguration property.
func (m *B2cIdentityUserFlow) SetApiConnectorConfiguration(value *UserFlowApiConnectorConfiguration)() {
    m.apiConnectorConfiguration = value
}
// Sets the defaultLanguageTag property value. Indicates the default language of the b2cIdentityUserFlow that is used when no ui_locale tag is specified in the request. This field is RFC 5646 compliant.
// Parameters:
//  - value : Value to set for the defaultLanguageTag property.
func (m *B2cIdentityUserFlow) SetDefaultLanguageTag(value *string)() {
    m.defaultLanguageTag = value
}
// Sets the identityProviders property value. 
// Parameters:
//  - value : Value to set for the identityProviders property.
func (m *B2cIdentityUserFlow) SetIdentityProviders(value []IdentityProvider)() {
    m.identityProviders = value
}
// Sets the isLanguageCustomizationEnabled property value. The property that determines whether language customization is enabled within the B2C user flow. Language customization is not enabled by default for B2C user flows.
// Parameters:
//  - value : Value to set for the isLanguageCustomizationEnabled property.
func (m *B2cIdentityUserFlow) SetIsLanguageCustomizationEnabled(value *bool)() {
    m.isLanguageCustomizationEnabled = value
}
// Sets the languages property value. The languages supported for customization within the user flow. Language customization is not enabled by default in B2C user flows.
// Parameters:
//  - value : Value to set for the languages property.
func (m *B2cIdentityUserFlow) SetLanguages(value []UserFlowLanguageConfiguration)() {
    m.languages = value
}
// Sets the userAttributeAssignments property value. The user attribute assignments included in the user flow.
// Parameters:
//  - value : Value to set for the userAttributeAssignments property.
func (m *B2cIdentityUserFlow) SetUserAttributeAssignments(value []IdentityUserFlowAttributeAssignment)() {
    m.userAttributeAssignments = value
}
// Sets the userFlowIdentityProviders property value. 
// Parameters:
//  - value : Value to set for the userFlowIdentityProviders property.
func (m *B2cIdentityUserFlow) SetUserFlowIdentityProviders(value []IdentityProviderBase)() {
    m.userFlowIdentityProviders = value
}
