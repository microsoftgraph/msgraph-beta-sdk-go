package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IdentityContainer 
type IdentityContainer struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Represents entry point for API connectors.
    apiConnectors []IdentityApiConnectorable
    // Represents entry point for B2C identity userflows.
    b2cUserFlows []B2cIdentityUserFlowable
    // Represents entry point for B2X and self-service sign-up identity userflows.
    b2xUserFlows []B2xIdentityUserFlowable
    // the entry point for the Conditional Access (CA) object model.
    conditionalAccess ConditionalAccessRootable
    // Represents entry point for continuous access evaluation policy.
    continuousAccessEvaluationPolicy ContinuousAccessEvaluationPolicyable
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
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetApiConnectors gets the apiConnectors property value. Represents entry point for API connectors.
func (m *IdentityContainer) GetApiConnectors()([]IdentityApiConnectorable) {
    if m == nil {
        return nil
    } else {
        return m.apiConnectors
    }
}
// GetB2cUserFlows gets the b2cUserFlows property value. Represents entry point for B2C identity userflows.
func (m *IdentityContainer) GetB2cUserFlows()([]B2cIdentityUserFlowable) {
    if m == nil {
        return nil
    } else {
        return m.b2cUserFlows
    }
}
// GetB2xUserFlows gets the b2xUserFlows property value. Represents entry point for B2X and self-service sign-up identity userflows.
func (m *IdentityContainer) GetB2xUserFlows()([]B2xIdentityUserFlowable) {
    if m == nil {
        return nil
    } else {
        return m.b2xUserFlows
    }
}
// GetConditionalAccess gets the conditionalAccess property value. the entry point for the Conditional Access (CA) object model.
func (m *IdentityContainer) GetConditionalAccess()(ConditionalAccessRootable) {
    if m == nil {
        return nil
    } else {
        return m.conditionalAccess
    }
}
// GetContinuousAccessEvaluationPolicy gets the continuousAccessEvaluationPolicy property value. Represents entry point for continuous access evaluation policy.
func (m *IdentityContainer) GetContinuousAccessEvaluationPolicy()(ContinuousAccessEvaluationPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.continuousAccessEvaluationPolicy
    }
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
                res[i] = v.(IdentityApiConnectorable)
            }
            m.SetApiConnectors(res)
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
                res[i] = v.(B2cIdentityUserFlowable)
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
                res[i] = v.(B2xIdentityUserFlowable)
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
    res["identityProviders"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIdentityProviderBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IdentityProviderBaseable, len(val))
            for i, v := range val {
                res[i] = v.(IdentityProviderBaseable)
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
                res[i] = v.(IdentityUserFlowAttributeable)
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
                res[i] = v.(IdentityUserFlowable)
            }
            m.SetUserFlows(res)
        }
        return nil
    }
    return res
}
// GetIdentityProviders gets the identityProviders property value. Represents entry point for identity provider base.
func (m *IdentityContainer) GetIdentityProviders()([]IdentityProviderBaseable) {
    if m == nil {
        return nil
    } else {
        return m.identityProviders
    }
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *IdentityContainer) GetOdataType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.odataType
    }
}
// GetUserFlowAttributes gets the userFlowAttributes property value. Represents entry point for identity userflow attributes.
func (m *IdentityContainer) GetUserFlowAttributes()([]IdentityUserFlowAttributeable) {
    if m == nil {
        return nil
    } else {
        return m.userFlowAttributes
    }
}
// GetUserFlows gets the userFlows property value. The userFlows property
func (m *IdentityContainer) GetUserFlows()([]IdentityUserFlowable) {
    if m == nil {
        return nil
    } else {
        return m.userFlows
    }
}
// Serialize serializes information the current object
func (m *IdentityContainer) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetApiConnectors() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetApiConnectors()))
        for i, v := range m.GetApiConnectors() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("apiConnectors", cast)
        if err != nil {
            return err
        }
    }
    if m.GetB2cUserFlows() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetB2cUserFlows()))
        for i, v := range m.GetB2cUserFlows() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("b2cUserFlows", cast)
        if err != nil {
            return err
        }
    }
    if m.GetB2xUserFlows() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetB2xUserFlows()))
        for i, v := range m.GetB2xUserFlows() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
    if m.GetIdentityProviders() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetIdentityProviders()))
        for i, v := range m.GetIdentityProviders() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("userFlowAttributes", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserFlows() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserFlows()))
        for i, v := range m.GetUserFlows() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IdentityContainer) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetApiConnectors sets the apiConnectors property value. Represents entry point for API connectors.
func (m *IdentityContainer) SetApiConnectors(value []IdentityApiConnectorable)() {
    if m != nil {
        m.apiConnectors = value
    }
}
// SetB2cUserFlows sets the b2cUserFlows property value. Represents entry point for B2C identity userflows.
func (m *IdentityContainer) SetB2cUserFlows(value []B2cIdentityUserFlowable)() {
    if m != nil {
        m.b2cUserFlows = value
    }
}
// SetB2xUserFlows sets the b2xUserFlows property value. Represents entry point for B2X and self-service sign-up identity userflows.
func (m *IdentityContainer) SetB2xUserFlows(value []B2xIdentityUserFlowable)() {
    if m != nil {
        m.b2xUserFlows = value
    }
}
// SetConditionalAccess sets the conditionalAccess property value. the entry point for the Conditional Access (CA) object model.
func (m *IdentityContainer) SetConditionalAccess(value ConditionalAccessRootable)() {
    if m != nil {
        m.conditionalAccess = value
    }
}
// SetContinuousAccessEvaluationPolicy sets the continuousAccessEvaluationPolicy property value. Represents entry point for continuous access evaluation policy.
func (m *IdentityContainer) SetContinuousAccessEvaluationPolicy(value ContinuousAccessEvaluationPolicyable)() {
    if m != nil {
        m.continuousAccessEvaluationPolicy = value
    }
}
// SetIdentityProviders sets the identityProviders property value. Represents entry point for identity provider base.
func (m *IdentityContainer) SetIdentityProviders(value []IdentityProviderBaseable)() {
    if m != nil {
        m.identityProviders = value
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *IdentityContainer) SetOdataType(value *string)() {
    if m != nil {
        m.odataType = value
    }
}
// SetUserFlowAttributes sets the userFlowAttributes property value. Represents entry point for identity userflow attributes.
func (m *IdentityContainer) SetUserFlowAttributes(value []IdentityUserFlowAttributeable)() {
    if m != nil {
        m.userFlowAttributes = value
    }
}
// SetUserFlows sets the userFlows property value. The userFlows property
func (m *IdentityContainer) SetUserFlows(value []IdentityUserFlowable)() {
    if m != nil {
        m.userFlows = value
    }
}
