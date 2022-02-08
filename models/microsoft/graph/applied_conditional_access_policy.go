package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AppliedConditionalAccessPolicy 
type AppliedConditionalAccessPolicy struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The custom authentication strength enforced in a Conditional Access policy.
    authenticationStrength *AuthenticationStrength;
    // Refers to the conditional access policy conditions that are not satisfied. The possible values are: none, application, users, devicePlatform, location, clientType, signInRisk, userRisk, time, deviceState, client,ipAddressSeenByAzureAD,ipAddressSeenByResourceProvider,unknownFutureValue,servicePrincipals,servicePrincipalRisk. Note that you must use the Prefer: include-unknown-enum-members request header to get the following values in this evolvable enum: servicePrincipals,servicePrincipalRisk.
    conditionsNotSatisfied *ConditionalAccessConditions;
    // Refers to the conditional access policy conditions that are satisfied. The possible values are: none, application, users, devicePlatform, location, clientType, signInRisk, userRisk, time, deviceState, client,ipAddressSeenByAzureAD,ipAddressSeenByResourceProvider,unknownFutureValue,servicePrincipals,servicePrincipalRisk. Note that you must use the Prefer: include-unknown-enum-members request header to get the following values in this evolvable enum: servicePrincipals,servicePrincipalRisk.
    conditionsSatisfied *ConditionalAccessConditions;
    // Refers to the Name of the conditional access policy (example: 'Require MFA for Salesforce').
    displayName *string;
    // Refers to the grant controls enforced by the conditional access policy (example: 'Require multi-factor authentication').
    enforcedGrantControls []string;
    // Refers to the session controls enforced by the conditional access policy (example: 'Require app enforced controls').
    enforcedSessionControls []string;
    // List of key-value pairs containing each matched exclude condition in the conditional access policy. Example: [{'devicePlatform' : 'DevicePlatform'}] means the policy didn’t apply, because the DevicePlatform condition was a match.
    excludeRulesSatisfied []ConditionalAccessRuleSatisfied;
    // An identifier of the conditional access policy.
    id *string;
    // List of key-value pairs containing each matched include condition in the conditional access policy. Example: [{ 'application' : 'AllApps'}, {'users': 'Group'}], meaning Application condition was a match because AllApps are included and Users condition was a match because the user was part of the included Group rule.
    includeRulesSatisfied []ConditionalAccessRuleSatisfied;
    // Indicates the result of the CA policy that was triggered. Possible values are: success, failure, notApplied (Policy isn't applied because policy conditions were not met),notEnabled (This is due to the policy in disabled state), unknown, unknownFutureValue.
    result *AppliedConditionalAccessPolicyResult;
}
// NewAppliedConditionalAccessPolicy instantiates a new appliedConditionalAccessPolicy and sets the default values.
func NewAppliedConditionalAccessPolicy()(*AppliedConditionalAccessPolicy) {
    m := &AppliedConditionalAccessPolicy{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AppliedConditionalAccessPolicy) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAuthenticationStrength gets the authenticationStrength property value. The custom authentication strength enforced in a Conditional Access policy.
func (m *AppliedConditionalAccessPolicy) GetAuthenticationStrength()(*AuthenticationStrength) {
    if m == nil {
        return nil
    } else {
        return m.authenticationStrength
    }
}
// GetConditionsNotSatisfied gets the conditionsNotSatisfied property value. Refers to the conditional access policy conditions that are not satisfied. The possible values are: none, application, users, devicePlatform, location, clientType, signInRisk, userRisk, time, deviceState, client,ipAddressSeenByAzureAD,ipAddressSeenByResourceProvider,unknownFutureValue,servicePrincipals,servicePrincipalRisk. Note that you must use the Prefer: include-unknown-enum-members request header to get the following values in this evolvable enum: servicePrincipals,servicePrincipalRisk.
func (m *AppliedConditionalAccessPolicy) GetConditionsNotSatisfied()(*ConditionalAccessConditions) {
    if m == nil {
        return nil
    } else {
        return m.conditionsNotSatisfied
    }
}
// GetConditionsSatisfied gets the conditionsSatisfied property value. Refers to the conditional access policy conditions that are satisfied. The possible values are: none, application, users, devicePlatform, location, clientType, signInRisk, userRisk, time, deviceState, client,ipAddressSeenByAzureAD,ipAddressSeenByResourceProvider,unknownFutureValue,servicePrincipals,servicePrincipalRisk. Note that you must use the Prefer: include-unknown-enum-members request header to get the following values in this evolvable enum: servicePrincipals,servicePrincipalRisk.
func (m *AppliedConditionalAccessPolicy) GetConditionsSatisfied()(*ConditionalAccessConditions) {
    if m == nil {
        return nil
    } else {
        return m.conditionsSatisfied
    }
}
// GetDisplayName gets the displayName property value. Refers to the Name of the conditional access policy (example: 'Require MFA for Salesforce').
func (m *AppliedConditionalAccessPolicy) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetEnforcedGrantControls gets the enforcedGrantControls property value. Refers to the grant controls enforced by the conditional access policy (example: 'Require multi-factor authentication').
func (m *AppliedConditionalAccessPolicy) GetEnforcedGrantControls()([]string) {
    if m == nil {
        return nil
    } else {
        return m.enforcedGrantControls
    }
}
// GetEnforcedSessionControls gets the enforcedSessionControls property value. Refers to the session controls enforced by the conditional access policy (example: 'Require app enforced controls').
func (m *AppliedConditionalAccessPolicy) GetEnforcedSessionControls()([]string) {
    if m == nil {
        return nil
    } else {
        return m.enforcedSessionControls
    }
}
// GetExcludeRulesSatisfied gets the excludeRulesSatisfied property value. List of key-value pairs containing each matched exclude condition in the conditional access policy. Example: [{'devicePlatform' : 'DevicePlatform'}] means the policy didn’t apply, because the DevicePlatform condition was a match.
func (m *AppliedConditionalAccessPolicy) GetExcludeRulesSatisfied()([]ConditionalAccessRuleSatisfied) {
    if m == nil {
        return nil
    } else {
        return m.excludeRulesSatisfied
    }
}
// GetId gets the id property value. An identifier of the conditional access policy.
func (m *AppliedConditionalAccessPolicy) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
// GetIncludeRulesSatisfied gets the includeRulesSatisfied property value. List of key-value pairs containing each matched include condition in the conditional access policy. Example: [{ 'application' : 'AllApps'}, {'users': 'Group'}], meaning Application condition was a match because AllApps are included and Users condition was a match because the user was part of the included Group rule.
func (m *AppliedConditionalAccessPolicy) GetIncludeRulesSatisfied()([]ConditionalAccessRuleSatisfied) {
    if m == nil {
        return nil
    } else {
        return m.includeRulesSatisfied
    }
}
// GetResult gets the result property value. Indicates the result of the CA policy that was triggered. Possible values are: success, failure, notApplied (Policy isn't applied because policy conditions were not met),notEnabled (This is due to the policy in disabled state), unknown, unknownFutureValue.
func (m *AppliedConditionalAccessPolicy) GetResult()(*AppliedConditionalAccessPolicyResult) {
    if m == nil {
        return nil
    } else {
        return m.result
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
            m.SetConditionsNotSatisfied(val.(*ConditionalAccessConditions))
        }
        return nil
    }
    res["conditionsSatisfied"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseConditionalAccessConditions)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConditionsSatisfied(val.(*ConditionalAccessConditions))
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
            m.SetResult(val.(*AppliedConditionalAccessPolicyResult))
        }
        return nil
    }
    return res
}
func (m *AppliedConditionalAccessPolicy) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AppliedConditionalAccessPolicy) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("authenticationStrength", m.GetAuthenticationStrength())
        if err != nil {
            return err
        }
    }
    if m.GetConditionsNotSatisfied() != nil {
        cast := (*m.GetConditionsNotSatisfied()).String()
        err := writer.WriteStringValue("conditionsNotSatisfied", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetConditionsSatisfied() != nil {
        cast := (*m.GetConditionsSatisfied()).String()
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
    if m.GetEnforcedGrantControls() != nil {
        err := writer.WriteCollectionOfStringValues("enforcedGrantControls", m.GetEnforcedGrantControls())
        if err != nil {
            return err
        }
    }
    if m.GetEnforcedSessionControls() != nil {
        err := writer.WriteCollectionOfStringValues("enforcedSessionControls", m.GetEnforcedSessionControls())
        if err != nil {
            return err
        }
    }
    if m.GetExcludeRulesSatisfied() != nil {
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
    if m.GetIncludeRulesSatisfied() != nil {
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
        cast := (*m.GetResult()).String()
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AppliedConditionalAccessPolicy) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAuthenticationStrength sets the authenticationStrength property value. The custom authentication strength enforced in a Conditional Access policy.
func (m *AppliedConditionalAccessPolicy) SetAuthenticationStrength(value *AuthenticationStrength)() {
    if m != nil {
        m.authenticationStrength = value
    }
}
// SetConditionsNotSatisfied sets the conditionsNotSatisfied property value. Refers to the conditional access policy conditions that are not satisfied. The possible values are: none, application, users, devicePlatform, location, clientType, signInRisk, userRisk, time, deviceState, client,ipAddressSeenByAzureAD,ipAddressSeenByResourceProvider,unknownFutureValue,servicePrincipals,servicePrincipalRisk. Note that you must use the Prefer: include-unknown-enum-members request header to get the following values in this evolvable enum: servicePrincipals,servicePrincipalRisk.
func (m *AppliedConditionalAccessPolicy) SetConditionsNotSatisfied(value *ConditionalAccessConditions)() {
    if m != nil {
        m.conditionsNotSatisfied = value
    }
}
// SetConditionsSatisfied sets the conditionsSatisfied property value. Refers to the conditional access policy conditions that are satisfied. The possible values are: none, application, users, devicePlatform, location, clientType, signInRisk, userRisk, time, deviceState, client,ipAddressSeenByAzureAD,ipAddressSeenByResourceProvider,unknownFutureValue,servicePrincipals,servicePrincipalRisk. Note that you must use the Prefer: include-unknown-enum-members request header to get the following values in this evolvable enum: servicePrincipals,servicePrincipalRisk.
func (m *AppliedConditionalAccessPolicy) SetConditionsSatisfied(value *ConditionalAccessConditions)() {
    if m != nil {
        m.conditionsSatisfied = value
    }
}
// SetDisplayName sets the displayName property value. Refers to the Name of the conditional access policy (example: 'Require MFA for Salesforce').
func (m *AppliedConditionalAccessPolicy) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetEnforcedGrantControls sets the enforcedGrantControls property value. Refers to the grant controls enforced by the conditional access policy (example: 'Require multi-factor authentication').
func (m *AppliedConditionalAccessPolicy) SetEnforcedGrantControls(value []string)() {
    if m != nil {
        m.enforcedGrantControls = value
    }
}
// SetEnforcedSessionControls sets the enforcedSessionControls property value. Refers to the session controls enforced by the conditional access policy (example: 'Require app enforced controls').
func (m *AppliedConditionalAccessPolicy) SetEnforcedSessionControls(value []string)() {
    if m != nil {
        m.enforcedSessionControls = value
    }
}
// SetExcludeRulesSatisfied sets the excludeRulesSatisfied property value. List of key-value pairs containing each matched exclude condition in the conditional access policy. Example: [{'devicePlatform' : 'DevicePlatform'}] means the policy didn’t apply, because the DevicePlatform condition was a match.
func (m *AppliedConditionalAccessPolicy) SetExcludeRulesSatisfied(value []ConditionalAccessRuleSatisfied)() {
    if m != nil {
        m.excludeRulesSatisfied = value
    }
}
// SetId sets the id property value. An identifier of the conditional access policy.
func (m *AppliedConditionalAccessPolicy) SetId(value *string)() {
    if m != nil {
        m.id = value
    }
}
// SetIncludeRulesSatisfied sets the includeRulesSatisfied property value. List of key-value pairs containing each matched include condition in the conditional access policy. Example: [{ 'application' : 'AllApps'}, {'users': 'Group'}], meaning Application condition was a match because AllApps are included and Users condition was a match because the user was part of the included Group rule.
func (m *AppliedConditionalAccessPolicy) SetIncludeRulesSatisfied(value []ConditionalAccessRuleSatisfied)() {
    if m != nil {
        m.includeRulesSatisfied = value
    }
}
// SetResult sets the result property value. Indicates the result of the CA policy that was triggered. Possible values are: success, failure, notApplied (Policy isn't applied because policy conditions were not met),notEnabled (This is due to the policy in disabled state), unknown, unknownFutureValue.
func (m *AppliedConditionalAccessPolicy) SetResult(value *AppliedConditionalAccessPolicyResult)() {
    if m != nil {
        m.result = value
    }
}
