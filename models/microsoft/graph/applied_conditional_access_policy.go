package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AppliedConditionalAccessPolicy struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    authenticationStrength *AuthenticationStrength;
    // Refers to the conditional access policy conditions that are not satisfied. Possible values are: none, application, users, devicePlatform, location, clientType, signInRisk, userRisk, time, deviceState, client.
    conditionsNotSatisfied *ConditionalAccessConditions;
    // Refers to the conditional access policy conditions that are satisfied. Possible values are: none, application, users, devicePlatform, location, clientType, signInRisk, userRisk, time, deviceState, client.
    conditionsSatisfied *ConditionalAccessConditions;
    // Refers to the Name of the conditional access policy (example: 'Require MFA for Salesforce').
    displayName *string;
    // Refers to the grant controls enforced by the conditional access policy (example: 'Require multi-factor authentication').
    enforcedGrantControls []string;
    // Refers to the session controls enforced by the conditional access policy (example: 'Require app enforced controls').
    enforcedSessionControls []string;
    // 
    excludeRulesSatisfied []ConditionalAccessRuleSatisfied;
    // An identifier of the conditional access policy.
    id *string;
    // 
    includeRulesSatisfied []ConditionalAccessRuleSatisfied;
    // Indicates the result of the CA policy that was triggered. Possible values are: success, failure, notApplied (Policy isn't applied because policy conditions were not met),notEnabled (This is due to the policy in disabled state), unknown, unknownFutureValue.
    result *AppliedConditionalAccessPolicyResult;
}
// Instantiates a new appliedConditionalAccessPolicy and sets the default values.
func NewAppliedConditionalAccessPolicy()(*AppliedConditionalAccessPolicy) {
    m := &AppliedConditionalAccessPolicy{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AppliedConditionalAccessPolicy) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the authenticationStrength property value. 
func (m *AppliedConditionalAccessPolicy) GetAuthenticationStrength()(*AuthenticationStrength) {
    if m == nil {
        return nil
    } else {
        return m.authenticationStrength
    }
}
// Gets the conditionsNotSatisfied property value. Refers to the conditional access policy conditions that are not satisfied. Possible values are: none, application, users, devicePlatform, location, clientType, signInRisk, userRisk, time, deviceState, client.
func (m *AppliedConditionalAccessPolicy) GetConditionsNotSatisfied()(*ConditionalAccessConditions) {
    if m == nil {
        return nil
    } else {
        return m.conditionsNotSatisfied
    }
}
// Gets the conditionsSatisfied property value. Refers to the conditional access policy conditions that are satisfied. Possible values are: none, application, users, devicePlatform, location, clientType, signInRisk, userRisk, time, deviceState, client.
func (m *AppliedConditionalAccessPolicy) GetConditionsSatisfied()(*ConditionalAccessConditions) {
    if m == nil {
        return nil
    } else {
        return m.conditionsSatisfied
    }
}
// Gets the displayName property value. Refers to the Name of the conditional access policy (example: 'Require MFA for Salesforce').
func (m *AppliedConditionalAccessPolicy) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the enforcedGrantControls property value. Refers to the grant controls enforced by the conditional access policy (example: 'Require multi-factor authentication').
func (m *AppliedConditionalAccessPolicy) GetEnforcedGrantControls()([]string) {
    if m == nil {
        return nil
    } else {
        return m.enforcedGrantControls
    }
}
// Gets the enforcedSessionControls property value. Refers to the session controls enforced by the conditional access policy (example: 'Require app enforced controls').
func (m *AppliedConditionalAccessPolicy) GetEnforcedSessionControls()([]string) {
    if m == nil {
        return nil
    } else {
        return m.enforcedSessionControls
    }
}
// Gets the excludeRulesSatisfied property value. 
func (m *AppliedConditionalAccessPolicy) GetExcludeRulesSatisfied()([]ConditionalAccessRuleSatisfied) {
    if m == nil {
        return nil
    } else {
        return m.excludeRulesSatisfied
    }
}
// Gets the id property value. An identifier of the conditional access policy.
func (m *AppliedConditionalAccessPolicy) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
// Gets the includeRulesSatisfied property value. 
func (m *AppliedConditionalAccessPolicy) GetIncludeRulesSatisfied()([]ConditionalAccessRuleSatisfied) {
    if m == nil {
        return nil
    } else {
        return m.includeRulesSatisfied
    }
}
// Gets the result property value. Indicates the result of the CA policy that was triggered. Possible values are: success, failure, notApplied (Policy isn't applied because policy conditions were not met),notEnabled (This is due to the policy in disabled state), unknown, unknownFutureValue.
func (m *AppliedConditionalAccessPolicy) GetResult()(*AppliedConditionalAccessPolicyResult) {
    if m == nil {
        return nil
    } else {
        return m.result
    }
}
// The deserialization information for the current model
func (m *AppliedConditionalAccessPolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["authenticationStrength"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAuthenticationStrength() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationStrength(val.(*AuthenticationStrength))
        }
        return nil
    }
    res["conditionsNotSatisfied"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseConditionalAccessConditions)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(ConditionalAccessConditions)
            m.SetConditionsNotSatisfied(&cast)
        }
        return nil
    }
    res["conditionsSatisfied"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseConditionalAccessConditions)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(ConditionalAccessConditions)
            m.SetConditionsSatisfied(&cast)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["enforcedGrantControls"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetEnforcedGrantControls(res)
        }
        return nil
    }
    res["enforcedSessionControls"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetEnforcedSessionControls(res)
        }
        return nil
    }
    res["excludeRulesSatisfied"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewConditionalAccessRuleSatisfied() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ConditionalAccessRuleSatisfied, len(val))
            for i, v := range val {
                res[i] = *(v.(*ConditionalAccessRuleSatisfied))
            }
            m.SetExcludeRulesSatisfied(res)
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
    res["includeRulesSatisfied"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewConditionalAccessRuleSatisfied() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ConditionalAccessRuleSatisfied, len(val))
            for i, v := range val {
                res[i] = *(v.(*ConditionalAccessRuleSatisfied))
            }
            m.SetIncludeRulesSatisfied(res)
        }
        return nil
    }
    res["result"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppliedConditionalAccessPolicyResult)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(AppliedConditionalAccessPolicyResult)
            m.SetResult(&cast)
        }
        return nil
    }
    return res
}
func (m *AppliedConditionalAccessPolicy) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AppliedConditionalAccessPolicy) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("authenticationStrength", m.GetAuthenticationStrength())
        if err != nil {
            return err
        }
    }
    if m.GetConditionsNotSatisfied() != nil {
        cast := m.GetConditionsNotSatisfied().String()
        err := writer.WriteStringValue("conditionsNotSatisfied", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetConditionsSatisfied() != nil {
        cast := m.GetConditionsSatisfied().String()
        err := writer.WriteStringValue("conditionsSatisfied", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("enforcedGrantControls", m.GetEnforcedGrantControls())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("enforcedSessionControls", m.GetEnforcedSessionControls())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetExcludeRulesSatisfied()))
        for i, v := range m.GetExcludeRulesSatisfied() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("excludeRulesSatisfied", cast)
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetIncludeRulesSatisfied()))
        for i, v := range m.GetIncludeRulesSatisfied() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("includeRulesSatisfied", cast)
        if err != nil {
            return err
        }
    }
    if m.GetResult() != nil {
        cast := m.GetResult().String()
        err := writer.WriteStringValue("result", &cast)
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
func (m *AppliedConditionalAccessPolicy) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the authenticationStrength property value. 
// Parameters:
//  - value : Value to set for the authenticationStrength property.
func (m *AppliedConditionalAccessPolicy) SetAuthenticationStrength(value *AuthenticationStrength)() {
    m.authenticationStrength = value
}
// Sets the conditionsNotSatisfied property value. Refers to the conditional access policy conditions that are not satisfied. Possible values are: none, application, users, devicePlatform, location, clientType, signInRisk, userRisk, time, deviceState, client.
// Parameters:
//  - value : Value to set for the conditionsNotSatisfied property.
func (m *AppliedConditionalAccessPolicy) SetConditionsNotSatisfied(value *ConditionalAccessConditions)() {
    m.conditionsNotSatisfied = value
}
// Sets the conditionsSatisfied property value. Refers to the conditional access policy conditions that are satisfied. Possible values are: none, application, users, devicePlatform, location, clientType, signInRisk, userRisk, time, deviceState, client.
// Parameters:
//  - value : Value to set for the conditionsSatisfied property.
func (m *AppliedConditionalAccessPolicy) SetConditionsSatisfied(value *ConditionalAccessConditions)() {
    m.conditionsSatisfied = value
}
// Sets the displayName property value. Refers to the Name of the conditional access policy (example: 'Require MFA for Salesforce').
// Parameters:
//  - value : Value to set for the displayName property.
func (m *AppliedConditionalAccessPolicy) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the enforcedGrantControls property value. Refers to the grant controls enforced by the conditional access policy (example: 'Require multi-factor authentication').
// Parameters:
//  - value : Value to set for the enforcedGrantControls property.
func (m *AppliedConditionalAccessPolicy) SetEnforcedGrantControls(value []string)() {
    m.enforcedGrantControls = value
}
// Sets the enforcedSessionControls property value. Refers to the session controls enforced by the conditional access policy (example: 'Require app enforced controls').
// Parameters:
//  - value : Value to set for the enforcedSessionControls property.
func (m *AppliedConditionalAccessPolicy) SetEnforcedSessionControls(value []string)() {
    m.enforcedSessionControls = value
}
// Sets the excludeRulesSatisfied property value. 
// Parameters:
//  - value : Value to set for the excludeRulesSatisfied property.
func (m *AppliedConditionalAccessPolicy) SetExcludeRulesSatisfied(value []ConditionalAccessRuleSatisfied)() {
    m.excludeRulesSatisfied = value
}
// Sets the id property value. An identifier of the conditional access policy.
// Parameters:
//  - value : Value to set for the id property.
func (m *AppliedConditionalAccessPolicy) SetId(value *string)() {
    m.id = value
}
// Sets the includeRulesSatisfied property value. 
// Parameters:
//  - value : Value to set for the includeRulesSatisfied property.
func (m *AppliedConditionalAccessPolicy) SetIncludeRulesSatisfied(value []ConditionalAccessRuleSatisfied)() {
    m.includeRulesSatisfied = value
}
// Sets the result property value. Indicates the result of the CA policy that was triggered. Possible values are: success, failure, notApplied (Policy isn't applied because policy conditions were not met),notEnabled (This is due to the policy in disabled state), unknown, unknownFutureValue.
// Parameters:
//  - value : Value to set for the result property.
func (m *AppliedConditionalAccessPolicy) SetResult(value *AppliedConditionalAccessPolicyResult)() {
    m.result = value
}
