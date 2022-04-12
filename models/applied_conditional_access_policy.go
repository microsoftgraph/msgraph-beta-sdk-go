package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AppliedConditionalAccessPolicy 
type AppliedConditionalAccessPolicy struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The custom authentication strength enforced in a Conditional Access policy.
    authenticationStrength AuthenticationStrengthable
    // Refers to the conditional access policy conditions that are not satisfied. The possible values are: none, application, users, devicePlatform, location, clientType, signInRisk, userRisk, time, deviceState, client,ipAddressSeenByAzureAD,ipAddressSeenByResourceProvider,unknownFutureValue,servicePrincipals,servicePrincipalRisk. Note that you must use the Prefer: include-unknown-enum-members request header to get the following values in this evolvable enum: servicePrincipals,servicePrincipalRisk.
    conditionsNotSatisfied *ConditionalAccessConditions
    // Refers to the conditional access policy conditions that are satisfied. The possible values are: none, application, users, devicePlatform, location, clientType, signInRisk, userRisk, time, deviceState, client,ipAddressSeenByAzureAD,ipAddressSeenByResourceProvider,unknownFutureValue,servicePrincipals,servicePrincipalRisk. Note that you must use the Prefer: include-unknown-enum-members request header to get the following values in this evolvable enum: servicePrincipals,servicePrincipalRisk.
    conditionsSatisfied *ConditionalAccessConditions
    // Refers to the Name of the conditional access policy (example: 'Require MFA for Salesforce').
    displayName *string
    // Refers to the grant controls enforced by the conditional access policy (example: 'Require multi-factor authentication').
    enforcedGrantControls []string
    // Refers to the session controls enforced by the conditional access policy (example: 'Require app enforced controls').
    enforcedSessionControls []string
    // List of key-value pairs containing each matched exclude condition in the conditional access policy. Example: [{'devicePlatform' : 'DevicePlatform'}] means the policy didn’t apply, because the DevicePlatform condition was a match.
    excludeRulesSatisfied []ConditionalAccessRuleSatisfiedable
    // An identifier of the conditional access policy.
    id *string
    // List of key-value pairs containing each matched include condition in the conditional access policy. Example: [{ 'application' : 'AllApps'}, {'users': 'Group'}], meaning Application condition was a match because AllApps are included and Users condition was a match because the user was part of the included Group rule.
    includeRulesSatisfied []ConditionalAccessRuleSatisfiedable
    // Indicates the result of the CA policy that was triggered. Possible values are: success, failure, notApplied (Policy isn't applied because policy conditions were not met),notEnabled (This is due to the policy in disabled state), unknown, unknownFutureValue.
    result *AppliedConditionalAccessPolicyResult
}
// NewAppliedConditionalAccessPolicy instantiates a new appliedConditionalAccessPolicy and sets the default values.
func NewAppliedConditionalAccessPolicy()(*AppliedConditionalAccessPolicy) {
    m := &AppliedConditionalAccessPolicy{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAppliedConditionalAccessPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAppliedConditionalAccessPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAppliedConditionalAccessPolicy(), nil
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
func (m *AppliedConditionalAccessPolicy) GetAuthenticationStrength()(AuthenticationStrengthable) {
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
func (m *AppliedConditionalAccessPolicy) GetExcludeRulesSatisfied()([]ConditionalAccessRuleSatisfiedable) {
    if m == nil {
        return nil
    } else {
        return m.excludeRulesSatisfied
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AppliedConditionalAccessPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["authenticationStrength"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthenticationStrengthFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationStrength(val.(AuthenticationStrengthable))
        }
        return nil
    }
    res["conditionsNotSatisfied"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseConditionalAccessConditions)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConditionsNotSatisfied(val.(*ConditionalAccessConditions))
        }
        return nil
    }
    res["conditionsSatisfied"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseConditionalAccessConditions)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConditionsSatisfied(val.(*ConditionalAccessConditions))
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["enforcedGrantControls"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
    res["enforcedSessionControls"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
    res["excludeRulesSatisfied"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateConditionalAccessRuleSatisfiedFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ConditionalAccessRuleSatisfiedable, len(val))
            for i, v := range val {
                res[i] = v.(ConditionalAccessRuleSatisfiedable)
            }
            m.SetExcludeRulesSatisfied(res)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["includeRulesSatisfied"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateConditionalAccessRuleSatisfiedFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ConditionalAccessRuleSatisfiedable, len(val))
            for i, v := range val {
                res[i] = v.(ConditionalAccessRuleSatisfiedable)
            }
            m.SetIncludeRulesSatisfied(res)
        }
        return nil
    }
    res["result"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetId gets the id property value. An identifier of the conditional access policy.
func (m *AppliedConditionalAccessPolicy) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
// GetIncludeRulesSatisfied gets the includeRulesSatisfied property value. List of key-value pairs containing each matched include condition in the conditional access policy. Example: [{ 'application' : 'AllApps'}, {'users': 'Group'}], meaning Application condition was a match because AllApps are included and Users condition was a match because the user was part of the included Group rule.
func (m *AppliedConditionalAccessPolicy) GetIncludeRulesSatisfied()([]ConditionalAccessRuleSatisfiedable) {
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
// Serialize serializes information the current object
func (m *AppliedConditionalAccessPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetExcludeRulesSatisfied()))
        for i, v := range m.GetExcludeRulesSatisfied() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetIncludeRulesSatisfied()))
        for i, v := range m.GetIncludeRulesSatisfied() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
func (m *AppliedConditionalAccessPolicy) SetAuthenticationStrength(value AuthenticationStrengthable)() {
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
func (m *AppliedConditionalAccessPolicy) SetExcludeRulesSatisfied(value []ConditionalAccessRuleSatisfiedable)() {
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
func (m *AppliedConditionalAccessPolicy) SetIncludeRulesSatisfied(value []ConditionalAccessRuleSatisfiedable)() {
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
