package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type IdentityContainer struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Represents entry point for API connectors.
    apiConnectors []IdentityApiConnector;
    // Represents entry point for B2C identity userflows.
    b2cUserFlows []B2cIdentityUserFlow;
    // Represents entry point for B2X/self-service sign-up identity userflows.
    b2xUserFlows []B2xIdentityUserFlow;
    // the entry point for the Conditional Access (CA) object model.
    conditionalAccess *ConditionalAccessRoot;
    // Represents entry point for continuous access evaluation policy.
    continuousAccessEvaluationPolicy *ContinuousAccessEvaluationPolicy;
    // Represents entry point for identity provider base.
    identityProviders []IdentityProviderBase;
    // Represents entry point for identity userflow attributes.
    userFlowAttributes []IdentityUserFlowAttribute;
    // 
    userFlows []IdentityUserFlow;
}
// Instantiates a new IdentityContainer and sets the default values.
func NewIdentityContainer()(*IdentityContainer) {
    m := &IdentityContainer{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IdentityContainer) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the apiConnectors property value. Represents entry point for API connectors.
func (m *IdentityContainer) GetApiConnectors()([]IdentityApiConnector) {
    if m == nil {
        return nil
    } else {
        return m.apiConnectors
    }
}
// Gets the b2cUserFlows property value. Represents entry point for B2C identity userflows.
func (m *IdentityContainer) GetB2cUserFlows()([]B2cIdentityUserFlow) {
    if m == nil {
        return nil
    } else {
        return m.b2cUserFlows
    }
}
// Gets the b2xUserFlows property value. Represents entry point for B2X/self-service sign-up identity userflows.
func (m *IdentityContainer) GetB2xUserFlows()([]B2xIdentityUserFlow) {
    if m == nil {
        return nil
    } else {
        return m.b2xUserFlows
    }
}
// Gets the conditionalAccess property value. the entry point for the Conditional Access (CA) object model.
func (m *IdentityContainer) GetConditionalAccess()(*ConditionalAccessRoot) {
    if m == nil {
        return nil
    } else {
        return m.conditionalAccess
    }
}
// Gets the continuousAccessEvaluationPolicy property value. Represents entry point for continuous access evaluation policy.
func (m *IdentityContainer) GetContinuousAccessEvaluationPolicy()(*ContinuousAccessEvaluationPolicy) {
    if m == nil {
        return nil
    } else {
        return m.continuousAccessEvaluationPolicy
    }
}
// Gets the identityProviders property value. Represents entry point for identity provider base.
func (m *IdentityContainer) GetIdentityProviders()([]IdentityProviderBase) {
    if m == nil {
        return nil
    } else {
        return m.identityProviders
    }
}
// Gets the userFlowAttributes property value. Represents entry point for identity userflow attributes.
func (m *IdentityContainer) GetUserFlowAttributes()([]IdentityUserFlowAttribute) {
    if m == nil {
        return nil
    } else {
        return m.userFlowAttributes
    }
}
// Gets the userFlows property value. 
func (m *IdentityContainer) GetUserFlows()([]IdentityUserFlow) {
    if m == nil {
        return nil
    } else {
        return m.userFlows
    }
}
// The deserialization information for the current model
func (m *IdentityContainer) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["apiConnectors"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentityApiConnector() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IdentityApiConnector, len(val))
            for i, v := range val {
                res[i] = *(v.(*IdentityApiConnector))
            }
            m.SetApiConnectors(res)
        }
        return nil
    }
    res["b2cUserFlows"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewB2cIdentityUserFlow() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]B2cIdentityUserFlow, len(val))
            for i, v := range val {
                res[i] = *(v.(*B2cIdentityUserFlow))
            }
            m.SetB2cUserFlows(res)
        }
        return nil
    }
    res["b2xUserFlows"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewB2xIdentityUserFlow() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]B2xIdentityUserFlow, len(val))
            for i, v := range val {
                res[i] = *(v.(*B2xIdentityUserFlow))
            }
            m.SetB2xUserFlows(res)
        }
        return nil
    }
    res["conditionalAccess"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewConditionalAccessRoot() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConditionalAccess(val.(*ConditionalAccessRoot))
        }
        return nil
    }
    res["continuousAccessEvaluationPolicy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewContinuousAccessEvaluationPolicy() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContinuousAccessEvaluationPolicy(val.(*ContinuousAccessEvaluationPolicy))
        }
        return nil
    }
    res["identityProviders"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentityProviderBase() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IdentityProviderBase, len(val))
            for i, v := range val {
                res[i] = *(v.(*IdentityProviderBase))
            }
            m.SetIdentityProviders(res)
        }
        return nil
    }
    res["userFlowAttributes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentityUserFlowAttribute() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IdentityUserFlowAttribute, len(val))
            for i, v := range val {
                res[i] = *(v.(*IdentityUserFlowAttribute))
            }
            m.SetUserFlowAttributes(res)
        }
        return nil
    }
    res["userFlows"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentityUserFlow() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IdentityUserFlow, len(val))
            for i, v := range val {
                res[i] = *(v.(*IdentityUserFlow))
            }
            m.SetUserFlows(res)
        }
        return nil
    }
    return res
}
func (m *IdentityContainer) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *IdentityContainer) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetApiConnectors()))
        for i, v := range m.GetApiConnectors() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("apiConnectors", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetB2cUserFlows()))
        for i, v := range m.GetB2cUserFlows() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("b2cUserFlows", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetB2xUserFlows()))
        for i, v := range m.GetB2xUserFlows() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetIdentityProviders()))
        for i, v := range m.GetIdentityProviders() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("identityProviders", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserFlowAttributes()))
        for i, v := range m.GetUserFlowAttributes() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("userFlowAttributes", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserFlows()))
        for i, v := range m.GetUserFlows() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *IdentityContainer) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the apiConnectors property value. Represents entry point for API connectors.
// Parameters:
//  - value : Value to set for the apiConnectors property.
func (m *IdentityContainer) SetApiConnectors(value []IdentityApiConnector)() {
    m.apiConnectors = value
}
// Sets the b2cUserFlows property value. Represents entry point for B2C identity userflows.
// Parameters:
//  - value : Value to set for the b2cUserFlows property.
func (m *IdentityContainer) SetB2cUserFlows(value []B2cIdentityUserFlow)() {
    m.b2cUserFlows = value
}
// Sets the b2xUserFlows property value. Represents entry point for B2X/self-service sign-up identity userflows.
// Parameters:
//  - value : Value to set for the b2xUserFlows property.
func (m *IdentityContainer) SetB2xUserFlows(value []B2xIdentityUserFlow)() {
    m.b2xUserFlows = value
}
// Sets the conditionalAccess property value. the entry point for the Conditional Access (CA) object model.
// Parameters:
//  - value : Value to set for the conditionalAccess property.
func (m *IdentityContainer) SetConditionalAccess(value *ConditionalAccessRoot)() {
    m.conditionalAccess = value
}
// Sets the continuousAccessEvaluationPolicy property value. Represents entry point for continuous access evaluation policy.
// Parameters:
//  - value : Value to set for the continuousAccessEvaluationPolicy property.
func (m *IdentityContainer) SetContinuousAccessEvaluationPolicy(value *ContinuousAccessEvaluationPolicy)() {
    m.continuousAccessEvaluationPolicy = value
}
// Sets the identityProviders property value. Represents entry point for identity provider base.
// Parameters:
//  - value : Value to set for the identityProviders property.
func (m *IdentityContainer) SetIdentityProviders(value []IdentityProviderBase)() {
    m.identityProviders = value
}
// Sets the userFlowAttributes property value. Represents entry point for identity userflow attributes.
// Parameters:
//  - value : Value to set for the userFlowAttributes property.
func (m *IdentityContainer) SetUserFlowAttributes(value []IdentityUserFlowAttribute)() {
    m.userFlowAttributes = value
}
// Sets the userFlows property value. 
// Parameters:
//  - value : Value to set for the userFlows property.
func (m *IdentityContainer) SetUserFlows(value []IdentityUserFlow)() {
    m.userFlows = value
}
