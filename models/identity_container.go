package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// IdentityContainer 
type IdentityContainer struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewIdentityContainer instantiates a new identityContainer and sets the default values.
func NewIdentityContainer()(*IdentityContainer) {
    m := &IdentityContainer{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateIdentityContainerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIdentityContainerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIdentityContainer(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IdentityContainer) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetApiConnectors gets the apiConnectors property value. Represents entry point for API connectors.
func (m *IdentityContainer) GetApiConnectors()([]IdentityApiConnectorable) {
    val, err := m.GetBackingStore().Get("apiConnectors")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]IdentityApiConnectorable)
    }
    return nil
}
// GetAuthenticationEventListeners gets the authenticationEventListeners property value. The authenticationEventListeners property
func (m *IdentityContainer) GetAuthenticationEventListeners()([]AuthenticationEventListenerable) {
    val, err := m.GetBackingStore().Get("authenticationEventListeners")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AuthenticationEventListenerable)
    }
    return nil
}
// GetAuthenticationEventsFlows gets the authenticationEventsFlows property value. Represents the entry point for self-service sign up and sign in user flows in both Azure AD workforce and customer tenants.
func (m *IdentityContainer) GetAuthenticationEventsFlows()([]AuthenticationEventsFlowable) {
    val, err := m.GetBackingStore().Get("authenticationEventsFlows")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AuthenticationEventsFlowable)
    }
    return nil
}
// GetB2cUserFlows gets the b2cUserFlows property value. Represents entry point for B2C identity userflows.
func (m *IdentityContainer) GetB2cUserFlows()([]B2cIdentityUserFlowable) {
    val, err := m.GetBackingStore().Get("b2cUserFlows")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]B2cIdentityUserFlowable)
    }
    return nil
}
// GetB2xUserFlows gets the b2xUserFlows property value. Represents entry point for B2X and self-service sign-up identity userflows.
func (m *IdentityContainer) GetB2xUserFlows()([]B2xIdentityUserFlowable) {
    val, err := m.GetBackingStore().Get("b2xUserFlows")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]B2xIdentityUserFlowable)
    }
    return nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
func (m *IdentityContainer) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetConditionalAccess gets the conditionalAccess property value. the entry point for the Conditional Access (CA) object model.
func (m *IdentityContainer) GetConditionalAccess()(ConditionalAccessRootable) {
    val, err := m.GetBackingStore().Get("conditionalAccess")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ConditionalAccessRootable)
    }
    return nil
}
// GetContinuousAccessEvaluationPolicy gets the continuousAccessEvaluationPolicy property value. Represents entry point for continuous access evaluation policy.
func (m *IdentityContainer) GetContinuousAccessEvaluationPolicy()(ContinuousAccessEvaluationPolicyable) {
    val, err := m.GetBackingStore().Get("continuousAccessEvaluationPolicy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ContinuousAccessEvaluationPolicyable)
    }
    return nil
}
// GetCustomAuthenticationExtensions gets the customAuthenticationExtensions property value. The customAuthenticationExtensions property
func (m *IdentityContainer) GetCustomAuthenticationExtensions()([]CustomAuthenticationExtensionable) {
    val, err := m.GetBackingStore().Get("customAuthenticationExtensions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CustomAuthenticationExtensionable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IdentityContainer) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["apiConnectors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIdentityApiConnectorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IdentityApiConnectorable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(IdentityApiConnectorable)
                }
            }
            m.SetApiConnectors(res)
        }
        return nil
    }
    res["authenticationEventListeners"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAuthenticationEventListenerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuthenticationEventListenerable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AuthenticationEventListenerable)
                }
            }
            m.SetAuthenticationEventListeners(res)
        }
        return nil
    }
    res["authenticationEventsFlows"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAuthenticationEventsFlowFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuthenticationEventsFlowable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AuthenticationEventsFlowable)
                }
            }
            m.SetAuthenticationEventsFlows(res)
        }
        return nil
    }
    res["b2cUserFlows"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateB2cIdentityUserFlowFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]B2cIdentityUserFlowable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(B2cIdentityUserFlowable)
                }
            }
            m.SetB2cUserFlows(res)
        }
        return nil
    }
    res["b2xUserFlows"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateB2xIdentityUserFlowFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]B2xIdentityUserFlowable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(B2xIdentityUserFlowable)
                }
            }
            m.SetB2xUserFlows(res)
        }
        return nil
    }
    res["conditionalAccess"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateConditionalAccessRootFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConditionalAccess(val.(ConditionalAccessRootable))
        }
        return nil
    }
    res["continuousAccessEvaluationPolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateContinuousAccessEvaluationPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContinuousAccessEvaluationPolicy(val.(ContinuousAccessEvaluationPolicyable))
        }
        return nil
    }
    res["customAuthenticationExtensions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCustomAuthenticationExtensionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CustomAuthenticationExtensionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CustomAuthenticationExtensionable)
                }
            }
            m.SetCustomAuthenticationExtensions(res)
        }
        return nil
    }
    res["identityProviders"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetIdentityProviders(res)
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
    res["userFlowAttributes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIdentityUserFlowAttributeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IdentityUserFlowAttributeable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(IdentityUserFlowAttributeable)
                }
            }
            m.SetUserFlowAttributes(res)
        }
        return nil
    }
    res["userFlows"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIdentityUserFlowFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IdentityUserFlowable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(IdentityUserFlowable)
                }
            }
            m.SetUserFlows(res)
        }
        return nil
    }
    return res
}
// GetIdentityProviders gets the identityProviders property value. Represents entry point for identity provider base.
func (m *IdentityContainer) GetIdentityProviders()([]IdentityProviderBaseable) {
    val, err := m.GetBackingStore().Get("identityProviders")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]IdentityProviderBaseable)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *IdentityContainer) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserFlowAttributes gets the userFlowAttributes property value. Represents entry point for identity userflow attributes.
func (m *IdentityContainer) GetUserFlowAttributes()([]IdentityUserFlowAttributeable) {
    val, err := m.GetBackingStore().Get("userFlowAttributes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]IdentityUserFlowAttributeable)
    }
    return nil
}
// GetUserFlows gets the userFlows property value. The userFlows property
func (m *IdentityContainer) GetUserFlows()([]IdentityUserFlowable) {
    val, err := m.GetBackingStore().Get("userFlows")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]IdentityUserFlowable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *IdentityContainer) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetApiConnectors() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetApiConnectors()))
        for i, v := range m.GetApiConnectors() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("apiConnectors", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAuthenticationEventListeners() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAuthenticationEventListeners()))
        for i, v := range m.GetAuthenticationEventListeners() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("authenticationEventListeners", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAuthenticationEventsFlows() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAuthenticationEventsFlows()))
        for i, v := range m.GetAuthenticationEventsFlows() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("authenticationEventsFlows", cast)
        if err != nil {
            return err
        }
    }
    if m.GetB2cUserFlows() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetB2cUserFlows()))
        for i, v := range m.GetB2cUserFlows() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("b2cUserFlows", cast)
        if err != nil {
            return err
        }
    }
    if m.GetB2xUserFlows() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetB2xUserFlows()))
        for i, v := range m.GetB2xUserFlows() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("b2xUserFlows", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("conditionalAccess", m.GetConditionalAccess())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("continuousAccessEvaluationPolicy", m.GetContinuousAccessEvaluationPolicy())
        if err != nil {
            return err
        }
    }
    if m.GetCustomAuthenticationExtensions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCustomAuthenticationExtensions()))
        for i, v := range m.GetCustomAuthenticationExtensions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("customAuthenticationExtensions", cast)
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
        err := writer.WriteCollectionOfObjectValues("identityProviders", cast)
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
    if m.GetUserFlowAttributes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserFlowAttributes()))
        for i, v := range m.GetUserFlowAttributes() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("userFlowAttributes", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserFlows() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserFlows()))
        for i, v := range m.GetUserFlows() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("userFlows", cast)
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IdentityContainer) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetApiConnectors sets the apiConnectors property value. Represents entry point for API connectors.
func (m *IdentityContainer) SetApiConnectors(value []IdentityApiConnectorable)() {
    err := m.GetBackingStore().Set("apiConnectors", value)
    if err != nil {
        panic(err)
    }
}
// SetAuthenticationEventListeners sets the authenticationEventListeners property value. The authenticationEventListeners property
func (m *IdentityContainer) SetAuthenticationEventListeners(value []AuthenticationEventListenerable)() {
    err := m.GetBackingStore().Set("authenticationEventListeners", value)
    if err != nil {
        panic(err)
    }
}
// SetAuthenticationEventsFlows sets the authenticationEventsFlows property value. Represents the entry point for self-service sign up and sign in user flows in both Azure AD workforce and customer tenants.
func (m *IdentityContainer) SetAuthenticationEventsFlows(value []AuthenticationEventsFlowable)() {
    err := m.GetBackingStore().Set("authenticationEventsFlows", value)
    if err != nil {
        panic(err)
    }
}
// SetB2cUserFlows sets the b2cUserFlows property value. Represents entry point for B2C identity userflows.
func (m *IdentityContainer) SetB2cUserFlows(value []B2cIdentityUserFlowable)() {
    err := m.GetBackingStore().Set("b2cUserFlows", value)
    if err != nil {
        panic(err)
    }
}
// SetB2xUserFlows sets the b2xUserFlows property value. Represents entry point for B2X and self-service sign-up identity userflows.
func (m *IdentityContainer) SetB2xUserFlows(value []B2xIdentityUserFlowable)() {
    err := m.GetBackingStore().Set("b2xUserFlows", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *IdentityContainer) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetConditionalAccess sets the conditionalAccess property value. the entry point for the Conditional Access (CA) object model.
func (m *IdentityContainer) SetConditionalAccess(value ConditionalAccessRootable)() {
    err := m.GetBackingStore().Set("conditionalAccess", value)
    if err != nil {
        panic(err)
    }
}
// SetContinuousAccessEvaluationPolicy sets the continuousAccessEvaluationPolicy property value. Represents entry point for continuous access evaluation policy.
func (m *IdentityContainer) SetContinuousAccessEvaluationPolicy(value ContinuousAccessEvaluationPolicyable)() {
    err := m.GetBackingStore().Set("continuousAccessEvaluationPolicy", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomAuthenticationExtensions sets the customAuthenticationExtensions property value. The customAuthenticationExtensions property
func (m *IdentityContainer) SetCustomAuthenticationExtensions(value []CustomAuthenticationExtensionable)() {
    err := m.GetBackingStore().Set("customAuthenticationExtensions", value)
    if err != nil {
        panic(err)
    }
}
// SetIdentityProviders sets the identityProviders property value. Represents entry point for identity provider base.
func (m *IdentityContainer) SetIdentityProviders(value []IdentityProviderBaseable)() {
    err := m.GetBackingStore().Set("identityProviders", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *IdentityContainer) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetUserFlowAttributes sets the userFlowAttributes property value. Represents entry point for identity userflow attributes.
func (m *IdentityContainer) SetUserFlowAttributes(value []IdentityUserFlowAttributeable)() {
    err := m.GetBackingStore().Set("userFlowAttributes", value)
    if err != nil {
        panic(err)
    }
}
// SetUserFlows sets the userFlows property value. The userFlows property
func (m *IdentityContainer) SetUserFlows(value []IdentityUserFlowable)() {
    err := m.GetBackingStore().Set("userFlows", value)
    if err != nil {
        panic(err)
    }
}
// IdentityContainerable 
type IdentityContainerable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApiConnectors()([]IdentityApiConnectorable)
    GetAuthenticationEventListeners()([]AuthenticationEventListenerable)
    GetAuthenticationEventsFlows()([]AuthenticationEventsFlowable)
    GetB2cUserFlows()([]B2cIdentityUserFlowable)
    GetB2xUserFlows()([]B2xIdentityUserFlowable)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetConditionalAccess()(ConditionalAccessRootable)
    GetContinuousAccessEvaluationPolicy()(ContinuousAccessEvaluationPolicyable)
    GetCustomAuthenticationExtensions()([]CustomAuthenticationExtensionable)
    GetIdentityProviders()([]IdentityProviderBaseable)
    GetOdataType()(*string)
    GetUserFlowAttributes()([]IdentityUserFlowAttributeable)
    GetUserFlows()([]IdentityUserFlowable)
    SetApiConnectors(value []IdentityApiConnectorable)()
    SetAuthenticationEventListeners(value []AuthenticationEventListenerable)()
    SetAuthenticationEventsFlows(value []AuthenticationEventsFlowable)()
    SetB2cUserFlows(value []B2cIdentityUserFlowable)()
    SetB2xUserFlows(value []B2xIdentityUserFlowable)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetConditionalAccess(value ConditionalAccessRootable)()
    SetContinuousAccessEvaluationPolicy(value ContinuousAccessEvaluationPolicyable)()
    SetCustomAuthenticationExtensions(value []CustomAuthenticationExtensionable)()
    SetIdentityProviders(value []IdentityProviderBaseable)()
    SetOdataType(value *string)()
    SetUserFlowAttributes(value []IdentityUserFlowAttributeable)()
    SetUserFlows(value []IdentityUserFlowable)()
}
