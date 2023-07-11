package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// B2cIdentityUserFlow 
type B2cIdentityUserFlow struct {
    IdentityUserFlow
    // The OdataType property
    OdataType *string
}
// NewB2cIdentityUserFlow instantiates a new b2cIdentityUserFlow and sets the default values.
func NewB2cIdentityUserFlow()(*B2cIdentityUserFlow) {
    m := &B2cIdentityUserFlow{
        IdentityUserFlow: *NewIdentityUserFlow(),
    }
    return m
}
// CreateB2cIdentityUserFlowFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateB2cIdentityUserFlowFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewB2cIdentityUserFlow(), nil
}
// GetApiConnectorConfiguration gets the apiConnectorConfiguration property value. Configuration for enabling an API connector for use as part of the user flow. You can only obtain the value of this object using Get userFlowApiConnectorConfiguration.
func (m *B2cIdentityUserFlow) GetApiConnectorConfiguration()(UserFlowApiConnectorConfigurationable) {
    val, err := m.GetBackingStore().Get("apiConnectorConfiguration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserFlowApiConnectorConfigurationable)
    }
    return nil
}
// GetDefaultLanguageTag gets the defaultLanguageTag property value. Indicates the default language of the b2cIdentityUserFlow that is used when no ui_locale tag is specified in the request. This field is RFC 5646 compliant.
func (m *B2cIdentityUserFlow) GetDefaultLanguageTag()(*string) {
    val, err := m.GetBackingStore().Get("defaultLanguageTag")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *B2cIdentityUserFlow) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.IdentityUserFlow.GetFieldDeserializers()
    res["apiConnectorConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserFlowApiConnectorConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApiConnectorConfiguration(val.(UserFlowApiConnectorConfigurationable))
        }
        return nil
    }
    res["defaultLanguageTag"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultLanguageTag(val)
        }
        return nil
    }
    res["identityProviders"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIdentityProviderFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IdentityProviderable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(IdentityProviderable)
                }
            }
            m.SetIdentityProviders(res)
        }
        return nil
    }
    res["isLanguageCustomizationEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsLanguageCustomizationEnabled(val)
        }
        return nil
    }
    res["languages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserFlowLanguageConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserFlowLanguageConfigurationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserFlowLanguageConfigurationable)
                }
            }
            m.SetLanguages(res)
        }
        return nil
    }
    res["userAttributeAssignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIdentityUserFlowAttributeAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IdentityUserFlowAttributeAssignmentable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(IdentityUserFlowAttributeAssignmentable)
                }
            }
            m.SetUserAttributeAssignments(res)
        }
        return nil
    }
    res["userFlowIdentityProviders"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIdentityProviderBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IdentityProviderBaseable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(IdentityProviderBaseable)
                }
            }
            m.SetUserFlowIdentityProviders(res)
        }
        return nil
    }
    return res
}
// GetIdentityProviders gets the identityProviders property value. The identityProviders property
func (m *B2cIdentityUserFlow) GetIdentityProviders()([]IdentityProviderable) {
    val, err := m.GetBackingStore().Get("identityProviders")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]IdentityProviderable)
    }
    return nil
}
// GetIsLanguageCustomizationEnabled gets the isLanguageCustomizationEnabled property value. The property that determines whether language customization is enabled within the B2C user flow. Language customization is not enabled by default for B2C user flows.
func (m *B2cIdentityUserFlow) GetIsLanguageCustomizationEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isLanguageCustomizationEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLanguages gets the languages property value. The languages supported for customization within the user flow. Language customization is not enabled by default in B2C user flows.
func (m *B2cIdentityUserFlow) GetLanguages()([]UserFlowLanguageConfigurationable) {
    val, err := m.GetBackingStore().Get("languages")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UserFlowLanguageConfigurationable)
    }
    return nil
}
// GetUserAttributeAssignments gets the userAttributeAssignments property value. The user attribute assignments included in the user flow.
func (m *B2cIdentityUserFlow) GetUserAttributeAssignments()([]IdentityUserFlowAttributeAssignmentable) {
    val, err := m.GetBackingStore().Get("userAttributeAssignments")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]IdentityUserFlowAttributeAssignmentable)
    }
    return nil
}
// GetUserFlowIdentityProviders gets the userFlowIdentityProviders property value. The userFlowIdentityProviders property
func (m *B2cIdentityUserFlow) GetUserFlowIdentityProviders()([]IdentityProviderBaseable) {
    val, err := m.GetBackingStore().Get("userFlowIdentityProviders")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]IdentityProviderBaseable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *B2cIdentityUserFlow) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetIdentityProviders() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetIdentityProviders()))
        for i, v := range m.GetIdentityProviders() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
    if m.GetLanguages() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLanguages()))
        for i, v := range m.GetLanguages() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("languages", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserAttributeAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserAttributeAssignments()))
        for i, v := range m.GetUserAttributeAssignments() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("userAttributeAssignments", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserFlowIdentityProviders() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserFlowIdentityProviders()))
        for i, v := range m.GetUserFlowIdentityProviders() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("userFlowIdentityProviders", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApiConnectorConfiguration sets the apiConnectorConfiguration property value. Configuration for enabling an API connector for use as part of the user flow. You can only obtain the value of this object using Get userFlowApiConnectorConfiguration.
func (m *B2cIdentityUserFlow) SetApiConnectorConfiguration(value UserFlowApiConnectorConfigurationable)() {
    err := m.GetBackingStore().Set("apiConnectorConfiguration", value)
    if err != nil {
        panic(err)
    }
}
// SetDefaultLanguageTag sets the defaultLanguageTag property value. Indicates the default language of the b2cIdentityUserFlow that is used when no ui_locale tag is specified in the request. This field is RFC 5646 compliant.
func (m *B2cIdentityUserFlow) SetDefaultLanguageTag(value *string)() {
    err := m.GetBackingStore().Set("defaultLanguageTag", value)
    if err != nil {
        panic(err)
    }
}
// SetIdentityProviders sets the identityProviders property value. The identityProviders property
func (m *B2cIdentityUserFlow) SetIdentityProviders(value []IdentityProviderable)() {
    err := m.GetBackingStore().Set("identityProviders", value)
    if err != nil {
        panic(err)
    }
}
// SetIsLanguageCustomizationEnabled sets the isLanguageCustomizationEnabled property value. The property that determines whether language customization is enabled within the B2C user flow. Language customization is not enabled by default for B2C user flows.
func (m *B2cIdentityUserFlow) SetIsLanguageCustomizationEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isLanguageCustomizationEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetLanguages sets the languages property value. The languages supported for customization within the user flow. Language customization is not enabled by default in B2C user flows.
func (m *B2cIdentityUserFlow) SetLanguages(value []UserFlowLanguageConfigurationable)() {
    err := m.GetBackingStore().Set("languages", value)
    if err != nil {
        panic(err)
    }
}
// SetUserAttributeAssignments sets the userAttributeAssignments property value. The user attribute assignments included in the user flow.
func (m *B2cIdentityUserFlow) SetUserAttributeAssignments(value []IdentityUserFlowAttributeAssignmentable)() {
    err := m.GetBackingStore().Set("userAttributeAssignments", value)
    if err != nil {
        panic(err)
    }
}
// SetUserFlowIdentityProviders sets the userFlowIdentityProviders property value. The userFlowIdentityProviders property
func (m *B2cIdentityUserFlow) SetUserFlowIdentityProviders(value []IdentityProviderBaseable)() {
    err := m.GetBackingStore().Set("userFlowIdentityProviders", value)
    if err != nil {
        panic(err)
    }
}
// B2cIdentityUserFlowable 
type B2cIdentityUserFlowable interface {
    IdentityUserFlowable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApiConnectorConfiguration()(UserFlowApiConnectorConfigurationable)
    GetDefaultLanguageTag()(*string)
    GetIdentityProviders()([]IdentityProviderable)
    GetIsLanguageCustomizationEnabled()(*bool)
    GetLanguages()([]UserFlowLanguageConfigurationable)
    GetUserAttributeAssignments()([]IdentityUserFlowAttributeAssignmentable)
    GetUserFlowIdentityProviders()([]IdentityProviderBaseable)
    SetApiConnectorConfiguration(value UserFlowApiConnectorConfigurationable)()
    SetDefaultLanguageTag(value *string)()
    SetIdentityProviders(value []IdentityProviderable)()
    SetIsLanguageCustomizationEnabled(value *bool)()
    SetLanguages(value []UserFlowLanguageConfigurationable)()
    SetUserAttributeAssignments(value []IdentityUserFlowAttributeAssignmentable)()
    SetUserFlowIdentityProviders(value []IdentityProviderBaseable)()
}
