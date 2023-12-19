package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// AppliedConditionalAccessPolicy 
type AppliedConditionalAccessPolicy struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewAppliedConditionalAccessPolicy instantiates a new appliedConditionalAccessPolicy and sets the default values.
func NewAppliedConditionalAccessPolicy()(*AppliedConditionalAccessPolicy) {
    m := &AppliedConditionalAccessPolicy{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAppliedConditionalAccessPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAppliedConditionalAccessPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAppliedConditionalAccessPolicy(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AppliedConditionalAccessPolicy) GetAdditionalData()(map[string]any) {
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
// GetAuthenticationStrength gets the authenticationStrength property value. The custom authentication strength enforced in a Conditional Access policy.
func (m *AppliedConditionalAccessPolicy) GetAuthenticationStrength()(AuthenticationStrengthable) {
    val, err := m.GetBackingStore().Get("authenticationStrength")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AuthenticationStrengthable)
    }
    return nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
func (m *AppliedConditionalAccessPolicy) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetConditionsNotSatisfied gets the conditionsNotSatisfied property value. Refers to the conditional access policy conditions that aren't satisfied. The possible values are: none, application, users, devicePlatform, location, clientType, signInRisk, userRisk, time, deviceState, client,ipAddressSeenByAzureAD,ipAddressSeenByResourceProvider,unknownFutureValue,servicePrincipals,servicePrincipalRisk, authenticationFlows, insiderRisk . You must use the Prefer: include-unknown-enum-members request header to get the following values in this evolvable enum: servicePrincipals,servicePrincipalRisk, authenticationFlows, insiderRisk. conditionalAccessConditions is a multi-valued enumeration and the property can contain multiple values in a comma-separated list.
func (m *AppliedConditionalAccessPolicy) GetConditionsNotSatisfied()(*AppliedConditionalAccessPolicy_conditionsNotSatisfied) {
    val, err := m.GetBackingStore().Get("conditionsNotSatisfied")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AppliedConditionalAccessPolicy_conditionsNotSatisfied)
    }
    return nil
}
// GetConditionsSatisfied gets the conditionsSatisfied property value. Refers to the conditional access policy conditions that are satisfied. The possible values are: none, application, users, devicePlatform, location, clientType, signInRisk, userRisk, time, deviceState, client,ipAddressSeenByAzureAD,ipAddressSeenByResourceProvider,unknownFutureValue,servicePrincipals,servicePrincipalRisk, authenticationFlows, insiderRisk. You must use the Prefer: include-unknown-enum-members request header to get the following values in this evolvable enum: servicePrincipals,servicePrincipalRisk, authenticationFlows, insiderRisk. conditionalAccessConditions is a multi-valued enumeration and the property can contain multiple values in a comma-separated list.
func (m *AppliedConditionalAccessPolicy) GetConditionsSatisfied()(*AppliedConditionalAccessPolicy_conditionsSatisfied) {
    val, err := m.GetBackingStore().Get("conditionsSatisfied")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AppliedConditionalAccessPolicy_conditionsSatisfied)
    }
    return nil
}
// GetDisplayName gets the displayName property value. Name of the conditional access policy.
func (m *AppliedConditionalAccessPolicy) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetEnforcedGrantControls gets the enforcedGrantControls property value. Refers to the grant controls enforced by the conditional access policy (example: 'Require multi-factor authentication').
func (m *AppliedConditionalAccessPolicy) GetEnforcedGrantControls()([]string) {
    val, err := m.GetBackingStore().Get("enforcedGrantControls")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetEnforcedSessionControls gets the enforcedSessionControls property value. Refers to the session controls enforced by the conditional access policy (example: 'Require app enforced controls').
func (m *AppliedConditionalAccessPolicy) GetEnforcedSessionControls()([]string) {
    val, err := m.GetBackingStore().Get("enforcedSessionControls")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetExcludeRulesSatisfied gets the excludeRulesSatisfied property value. List of key-value pairs containing each matched exclude condition in the conditional access policy. Example: [{'devicePlatform' : 'DevicePlatform'}] means the policy didn't apply, because the DevicePlatform condition was a match.
func (m *AppliedConditionalAccessPolicy) GetExcludeRulesSatisfied()([]ConditionalAccessRuleSatisfiedable) {
    val, err := m.GetBackingStore().Get("excludeRulesSatisfied")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ConditionalAccessRuleSatisfiedable)
    }
    return nil
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
        val, err := n.GetEnumValue(ParseAppliedConditionalAccessPolicy_conditionsNotSatisfied)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConditionsNotSatisfied(val.(*AppliedConditionalAccessPolicy_conditionsNotSatisfied))
        }
        return nil
    }
    res["conditionsSatisfied"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppliedConditionalAccessPolicy_conditionsSatisfied)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConditionsSatisfied(val.(*AppliedConditionalAccessPolicy_conditionsSatisfied))
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
                if v != nil {
                    res[i] = *(v.(*string))
                }
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
                if v != nil {
                    res[i] = *(v.(*string))
                }
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
                if v != nil {
                    res[i] = v.(ConditionalAccessRuleSatisfiedable)
                }
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
                if v != nil {
                    res[i] = v.(ConditionalAccessRuleSatisfiedable)
                }
            }
            m.SetIncludeRulesSatisfied(res)
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
    res["result"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppliedConditionalAccessPolicy_result)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResult(val.(*AppliedConditionalAccessPolicy_result))
        }
        return nil
    }
    res["sessionControlsNotSatisfied"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetSessionControlsNotSatisfied(res)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. Identifier of the conditional access policy.
func (m *AppliedConditionalAccessPolicy) GetId()(*string) {
    val, err := m.GetBackingStore().Get("id")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetIncludeRulesSatisfied gets the includeRulesSatisfied property value. List of key-value pairs containing each matched include condition in the conditional access policy. Example: [{ 'application' : 'AllApps'}, {'users': 'Group'}], meaning Application condition was a match because AllApps are included and Users condition was a match because the user was part of the included Group rule.
func (m *AppliedConditionalAccessPolicy) GetIncludeRulesSatisfied()([]ConditionalAccessRuleSatisfiedable) {
    val, err := m.GetBackingStore().Get("includeRulesSatisfied")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ConditionalAccessRuleSatisfiedable)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AppliedConditionalAccessPolicy) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetResult gets the result property value. Indicates the result of the CA policy that was triggered. Possible values are: success, failure, notApplied (Policy isn't applied because policy conditions weren't met),notEnabled (This is due to the policy in disabled state), unknown, unknownFutureValue, reportOnlySuccess, reportOnlyFailure, reportOnlyNotApplied, reportOnlyInterrupted. You must use the Prefer: include-unknown-enum-members request header to get the following values in this evolvable enum: reportOnlySuccess, reportOnlyFailure, reportOnlyNotApplied, reportOnlyInterrupted.
func (m *AppliedConditionalAccessPolicy) GetResult()(*AppliedConditionalAccessPolicy_result) {
    val, err := m.GetBackingStore().Get("result")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AppliedConditionalAccessPolicy_result)
    }
    return nil
}
// GetSessionControlsNotSatisfied gets the sessionControlsNotSatisfied property value. Refers to the session controls that a sign-in activity didn't satisfy. (Example: Application enforced Restrictions).
func (m *AppliedConditionalAccessPolicy) GetSessionControlsNotSatisfied()([]string) {
    val, err := m.GetBackingStore().Get("sessionControlsNotSatisfied")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
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
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("includeRulesSatisfied", cast)
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
    if m.GetResult() != nil {
        cast := (*m.GetResult()).String()
        err := writer.WriteStringValue("result", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSessionControlsNotSatisfied() != nil {
        err := writer.WriteCollectionOfStringValues("sessionControlsNotSatisfied", m.GetSessionControlsNotSatisfied())
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
func (m *AppliedConditionalAccessPolicy) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAuthenticationStrength sets the authenticationStrength property value. The custom authentication strength enforced in a Conditional Access policy.
func (m *AppliedConditionalAccessPolicy) SetAuthenticationStrength(value AuthenticationStrengthable)() {
    err := m.GetBackingStore().Set("authenticationStrength", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *AppliedConditionalAccessPolicy) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetConditionsNotSatisfied sets the conditionsNotSatisfied property value. Refers to the conditional access policy conditions that aren't satisfied. The possible values are: none, application, users, devicePlatform, location, clientType, signInRisk, userRisk, time, deviceState, client,ipAddressSeenByAzureAD,ipAddressSeenByResourceProvider,unknownFutureValue,servicePrincipals,servicePrincipalRisk, authenticationFlows, insiderRisk . You must use the Prefer: include-unknown-enum-members request header to get the following values in this evolvable enum: servicePrincipals,servicePrincipalRisk, authenticationFlows, insiderRisk. conditionalAccessConditions is a multi-valued enumeration and the property can contain multiple values in a comma-separated list.
func (m *AppliedConditionalAccessPolicy) SetConditionsNotSatisfied(value *AppliedConditionalAccessPolicy_conditionsNotSatisfied)() {
    err := m.GetBackingStore().Set("conditionsNotSatisfied", value)
    if err != nil {
        panic(err)
    }
}
// SetConditionsSatisfied sets the conditionsSatisfied property value. Refers to the conditional access policy conditions that are satisfied. The possible values are: none, application, users, devicePlatform, location, clientType, signInRisk, userRisk, time, deviceState, client,ipAddressSeenByAzureAD,ipAddressSeenByResourceProvider,unknownFutureValue,servicePrincipals,servicePrincipalRisk, authenticationFlows, insiderRisk. You must use the Prefer: include-unknown-enum-members request header to get the following values in this evolvable enum: servicePrincipals,servicePrincipalRisk, authenticationFlows, insiderRisk. conditionalAccessConditions is a multi-valued enumeration and the property can contain multiple values in a comma-separated list.
func (m *AppliedConditionalAccessPolicy) SetConditionsSatisfied(value *AppliedConditionalAccessPolicy_conditionsSatisfied)() {
    err := m.GetBackingStore().Set("conditionsSatisfied", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. Name of the conditional access policy.
func (m *AppliedConditionalAccessPolicy) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetEnforcedGrantControls sets the enforcedGrantControls property value. Refers to the grant controls enforced by the conditional access policy (example: 'Require multi-factor authentication').
func (m *AppliedConditionalAccessPolicy) SetEnforcedGrantControls(value []string)() {
    err := m.GetBackingStore().Set("enforcedGrantControls", value)
    if err != nil {
        panic(err)
    }
}
// SetEnforcedSessionControls sets the enforcedSessionControls property value. Refers to the session controls enforced by the conditional access policy (example: 'Require app enforced controls').
func (m *AppliedConditionalAccessPolicy) SetEnforcedSessionControls(value []string)() {
    err := m.GetBackingStore().Set("enforcedSessionControls", value)
    if err != nil {
        panic(err)
    }
}
// SetExcludeRulesSatisfied sets the excludeRulesSatisfied property value. List of key-value pairs containing each matched exclude condition in the conditional access policy. Example: [{'devicePlatform' : 'DevicePlatform'}] means the policy didn't apply, because the DevicePlatform condition was a match.
func (m *AppliedConditionalAccessPolicy) SetExcludeRulesSatisfied(value []ConditionalAccessRuleSatisfiedable)() {
    err := m.GetBackingStore().Set("excludeRulesSatisfied", value)
    if err != nil {
        panic(err)
    }
}
// SetId sets the id property value. Identifier of the conditional access policy.
func (m *AppliedConditionalAccessPolicy) SetId(value *string)() {
    err := m.GetBackingStore().Set("id", value)
    if err != nil {
        panic(err)
    }
}
// SetIncludeRulesSatisfied sets the includeRulesSatisfied property value. List of key-value pairs containing each matched include condition in the conditional access policy. Example: [{ 'application' : 'AllApps'}, {'users': 'Group'}], meaning Application condition was a match because AllApps are included and Users condition was a match because the user was part of the included Group rule.
func (m *AppliedConditionalAccessPolicy) SetIncludeRulesSatisfied(value []ConditionalAccessRuleSatisfiedable)() {
    err := m.GetBackingStore().Set("includeRulesSatisfied", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AppliedConditionalAccessPolicy) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetResult sets the result property value. Indicates the result of the CA policy that was triggered. Possible values are: success, failure, notApplied (Policy isn't applied because policy conditions weren't met),notEnabled (This is due to the policy in disabled state), unknown, unknownFutureValue, reportOnlySuccess, reportOnlyFailure, reportOnlyNotApplied, reportOnlyInterrupted. You must use the Prefer: include-unknown-enum-members request header to get the following values in this evolvable enum: reportOnlySuccess, reportOnlyFailure, reportOnlyNotApplied, reportOnlyInterrupted.
func (m *AppliedConditionalAccessPolicy) SetResult(value *AppliedConditionalAccessPolicy_result)() {
    err := m.GetBackingStore().Set("result", value)
    if err != nil {
        panic(err)
    }
}
// SetSessionControlsNotSatisfied sets the sessionControlsNotSatisfied property value. Refers to the session controls that a sign-in activity didn't satisfy. (Example: Application enforced Restrictions).
func (m *AppliedConditionalAccessPolicy) SetSessionControlsNotSatisfied(value []string)() {
    err := m.GetBackingStore().Set("sessionControlsNotSatisfied", value)
    if err != nil {
        panic(err)
    }
}
// AppliedConditionalAccessPolicyable 
type AppliedConditionalAccessPolicyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthenticationStrength()(AuthenticationStrengthable)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetConditionsNotSatisfied()(*AppliedConditionalAccessPolicy_conditionsNotSatisfied)
    GetConditionsSatisfied()(*AppliedConditionalAccessPolicy_conditionsSatisfied)
    GetDisplayName()(*string)
    GetEnforcedGrantControls()([]string)
    GetEnforcedSessionControls()([]string)
    GetExcludeRulesSatisfied()([]ConditionalAccessRuleSatisfiedable)
    GetId()(*string)
    GetIncludeRulesSatisfied()([]ConditionalAccessRuleSatisfiedable)
    GetOdataType()(*string)
    GetResult()(*AppliedConditionalAccessPolicy_result)
    GetSessionControlsNotSatisfied()([]string)
    SetAuthenticationStrength(value AuthenticationStrengthable)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetConditionsNotSatisfied(value *AppliedConditionalAccessPolicy_conditionsNotSatisfied)()
    SetConditionsSatisfied(value *AppliedConditionalAccessPolicy_conditionsSatisfied)()
    SetDisplayName(value *string)()
    SetEnforcedGrantControls(value []string)()
    SetEnforcedSessionControls(value []string)()
    SetExcludeRulesSatisfied(value []ConditionalAccessRuleSatisfiedable)()
    SetId(value *string)()
    SetIncludeRulesSatisfied(value []ConditionalAccessRuleSatisfiedable)()
    SetOdataType(value *string)()
    SetResult(value *AppliedConditionalAccessPolicy_result)()
    SetSessionControlsNotSatisfied(value []string)()
}
