package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IdentityContainer 
type IdentityContainer struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Represents entry point for API connectors.
    apiConnectors []IdentityApiConnectorable
    // The authenticationEventListeners property
    authenticationEventListeners []AuthenticationEventListenerable
    // Represents entry point for B2C identity userflows.
    b2cUserFlows []B2cIdentityUserFlowable
    // Represents entry point for B2X and self-service sign-up identity userflows.
    b2xUserFlows []B2xIdentityUserFlowable
    // the entry point for the Conditional Access (CA) object model.
    conditionalAccess ConditionalAccessRootable
    // Represents entry point for continuous access evaluation policy.
    continuousAccessEvaluationPolicy ContinuousAccessEvaluationPolicyable
    // The customAuthenticationExtensions property
    customAuthenticationExtensions []CustomAuthenticationExtensionable
    // Represents entry point for identity provider base.
    identityProviders []IdentityProviderBaseable
    // The OdataType property
    odataType *string
    // Represents entry point for identity userflow attributes.
    userFlowAttributes []IdentityUserFlowAttributeable
    // The userFlows property
    userFlows []IdentityUserFlowable
}
// NewIdentityContainer instantiates a new IdentityContainer and sets the default values.
func NewIdentityContainer()(*IdentityContainer) {
    m := &IdentityContainer{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.identityContainer";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateIdentityContainerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIdentityContainerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIdentityContainer(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IdentityContainer) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetApiConnectors gets the apiConnectors property value. Represents entry point for API connectors.
func (m *IdentityContainer) GetApiConnectors()([]IdentityApiConnectorable) {
    return m.apiConnectors
}
// GetAuthenticationEventListeners gets the authenticationEventListeners property value. The authenticationEventListeners property
func (m *IdentityContainer) GetAuthenticationEventListeners()([]AuthenticationEventListenerable) {
    return m.authenticationEventListeners
}
// GetB2cUserFlows gets the b2cUserFlows property value. Represents entry point for B2C identity userflows.
func (m *IdentityContainer) GetB2cUserFlows()([]B2cIdentityUserFlowable) {
    return m.b2cUserFlows
}
// GetB2xUserFlows gets the b2xUserFlows property value. Represents entry point for B2X and self-service sign-up identity userflows.
func (m *IdentityContainer) GetB2xUserFlows()([]B2xIdentityUserFlowable) {
    return m.b2xUserFlows
}
// GetConditionalAccess gets the conditionalAccess property value. the entry point for the Conditional Access (CA) object model.
func (m *IdentityContainer) GetConditionalAccess()(ConditionalAccessRootable) {
    return m.conditionalAccess
}
// GetContinuousAccessEvaluationPolicy gets the continuousAccessEvaluationPolicy property value. Represents entry point for continuous access evaluation policy.
func (m *IdentityContainer) GetContinuousAccessEvaluationPolicy()(ContinuousAccessEvaluationPolicyable) {
    return m.continuousAccessEvaluationPolicy
}
// GetCustomAuthenticationExtensions gets the customAuthenticationExtensions property value. The customAuthenticationExtensions property
func (m *IdentityContainer) GetCustomAuthenticationExtensions()([]CustomAuthenticationExtensionable) {
    return m.customAuthenticationExtensions
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IdentityContainer) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["apiConnectors"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateIdentityApiConnectorFromDiscriminatorValue , m.SetApiConnectors)
    res["authenticationEventListeners"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAuthenticationEventListenerFromDiscriminatorValue , m.SetAuthenticationEventListeners)
    res["b2cUserFlows"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateB2cIdentityUserFlowFromDiscriminatorValue , m.SetB2cUserFlows)
    res["b2xUserFlows"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateB2xIdentityUserFlowFromDiscriminatorValue , m.SetB2xUserFlows)
    res["conditionalAccess"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateConditionalAccessRootFromDiscriminatorValue , m.SetConditionalAccess)
    res["continuousAccessEvaluationPolicy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateContinuousAccessEvaluationPolicyFromDiscriminatorValue , m.SetContinuousAccessEvaluationPolicy)
    res["customAuthenticationExtensions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateCustomAuthenticationExtensionFromDiscriminatorValue , m.SetCustomAuthenticationExtensions)
    res["identityProviders"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateIdentityProviderBaseFromDiscriminatorValue , m.SetIdentityProviders)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["userFlowAttributes"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateIdentityUserFlowAttributeFromDiscriminatorValue , m.SetUserFlowAttributes)
    res["userFlows"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateIdentityUserFlowFromDiscriminatorValue , m.SetUserFlows)
    return res
}
// GetIdentityProviders gets the identityProviders property value. Represents entry point for identity provider base.
func (m *IdentityContainer) GetIdentityProviders()([]IdentityProviderBaseable) {
    return m.identityProviders
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *IdentityContainer) GetOdataType()(*string) {
    return m.odataType
}
// GetUserFlowAttributes gets the userFlowAttributes property value. Represents entry point for identity userflow attributes.
func (m *IdentityContainer) GetUserFlowAttributes()([]IdentityUserFlowAttributeable) {
    return m.userFlowAttributes
}
// GetUserFlows gets the userFlows property value. The userFlows property
func (m *IdentityContainer) GetUserFlows()([]IdentityUserFlowable) {
    return m.userFlows
}
// Serialize serializes information the current object
func (m *IdentityContainer) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetApiConnectors() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetApiConnectors())
        err := writer.WriteCollectionOfObjectValues("apiConnectors", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAuthenticationEventListeners() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAuthenticationEventListeners())
        err := writer.WriteCollectionOfObjectValues("authenticationEventListeners", cast)
        if err != nil {
            return err
        }
    }
    if m.GetB2cUserFlows() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetB2cUserFlows())
        err := writer.WriteCollectionOfObjectValues("b2cUserFlows", cast)
        if err != nil {
            return err
        }
    }
    if m.GetB2xUserFlows() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetB2xUserFlows())
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
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetCustomAuthenticationExtensions())
        err := writer.WriteCollectionOfObjectValues("customAuthenticationExtensions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetIdentityProviders() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetIdentityProviders())
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
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetUserFlowAttributes())
        err := writer.WriteCollectionOfObjectValues("userFlowAttributes", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserFlows() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetUserFlows())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IdentityContainer) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetApiConnectors sets the apiConnectors property value. Represents entry point for API connectors.
func (m *IdentityContainer) SetApiConnectors(value []IdentityApiConnectorable)() {
    m.apiConnectors = value
}
// SetAuthenticationEventListeners sets the authenticationEventListeners property value. The authenticationEventListeners property
func (m *IdentityContainer) SetAuthenticationEventListeners(value []AuthenticationEventListenerable)() {
    m.authenticationEventListeners = value
}
// SetB2cUserFlows sets the b2cUserFlows property value. Represents entry point for B2C identity userflows.
func (m *IdentityContainer) SetB2cUserFlows(value []B2cIdentityUserFlowable)() {
    m.b2cUserFlows = value
}
// SetB2xUserFlows sets the b2xUserFlows property value. Represents entry point for B2X and self-service sign-up identity userflows.
func (m *IdentityContainer) SetB2xUserFlows(value []B2xIdentityUserFlowable)() {
    m.b2xUserFlows = value
}
// SetConditionalAccess sets the conditionalAccess property value. the entry point for the Conditional Access (CA) object model.
func (m *IdentityContainer) SetConditionalAccess(value ConditionalAccessRootable)() {
    m.conditionalAccess = value
}
// SetContinuousAccessEvaluationPolicy sets the continuousAccessEvaluationPolicy property value. Represents entry point for continuous access evaluation policy.
func (m *IdentityContainer) SetContinuousAccessEvaluationPolicy(value ContinuousAccessEvaluationPolicyable)() {
    m.continuousAccessEvaluationPolicy = value
}
// SetCustomAuthenticationExtensions sets the customAuthenticationExtensions property value. The customAuthenticationExtensions property
func (m *IdentityContainer) SetCustomAuthenticationExtensions(value []CustomAuthenticationExtensionable)() {
    m.customAuthenticationExtensions = value
}
// SetIdentityProviders sets the identityProviders property value. Represents entry point for identity provider base.
func (m *IdentityContainer) SetIdentityProviders(value []IdentityProviderBaseable)() {
    m.identityProviders = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *IdentityContainer) SetOdataType(value *string)() {
    m.odataType = value
}
// SetUserFlowAttributes sets the userFlowAttributes property value. Represents entry point for identity userflow attributes.
func (m *IdentityContainer) SetUserFlowAttributes(value []IdentityUserFlowAttributeable)() {
    m.userFlowAttributes = value
}
// SetUserFlows sets the userFlows property value. The userFlows property
func (m *IdentityContainer) SetUserFlows(value []IdentityUserFlowable)() {
    m.userFlows = value
}
