package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ConditionalAccessRuleSatisfied 
type ConditionalAccessRuleSatisfied struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Refers to the conditional access policy conditions that are satisfied. The possible values are: none, application, users, devicePlatform, location, clientType, signInRisk, userRisk, time, deviceState, client, ipAddressSeenByAzureAD, ipAddressSeenByResourceProvider, unknownFutureValue, servicePrincipals, servicePrincipalRisk. Note that you must use the Prefer: include-unknown-enum-members request header to get the following values in this evolvable enum: servicePrincipals, servicePrincipalRisk.
    conditionalAccessCondition *ConditionalAccessConditions;
    // Refers to the conditional access policy conditions that were satisfied. The possible values are: allApps, firstPartyApps, office365, appId, acr, appFilter, allUsers, guest, groupId, roleId, userId, allDevicePlatforms, devicePlatform, allLocations, insideCorpnet, allTrustedLocations, locationId, allDevices, deviceFilter, deviceState, unknownFutureValue, deviceFilterIncludeRuleNotMatched, allDeviceStates. Note that you must use the Prefer: include-unknown-enum-members request header to get the following values in this evolvable enum: deviceFilterIncludeRuleNotMatched, allDeviceStates.
    ruleSatisfied *ConditionalAccessRule;
}
// NewConditionalAccessRuleSatisfied instantiates a new conditionalAccessRuleSatisfied and sets the default values.
func NewConditionalAccessRuleSatisfied()(*ConditionalAccessRuleSatisfied) {
    m := &ConditionalAccessRuleSatisfied{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConditionalAccessRuleSatisfied) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetConditionalAccessCondition gets the conditionalAccessCondition property value. Refers to the conditional access policy conditions that are satisfied. The possible values are: none, application, users, devicePlatform, location, clientType, signInRisk, userRisk, time, deviceState, client, ipAddressSeenByAzureAD, ipAddressSeenByResourceProvider, unknownFutureValue, servicePrincipals, servicePrincipalRisk. Note that you must use the Prefer: include-unknown-enum-members request header to get the following values in this evolvable enum: servicePrincipals, servicePrincipalRisk.
func (m *ConditionalAccessRuleSatisfied) GetConditionalAccessCondition()(*ConditionalAccessConditions) {
    if m == nil {
        return nil
    } else {
        return m.conditionalAccessCondition
    }
}
// GetRuleSatisfied gets the ruleSatisfied property value. Refers to the conditional access policy conditions that were satisfied. The possible values are: allApps, firstPartyApps, office365, appId, acr, appFilter, allUsers, guest, groupId, roleId, userId, allDevicePlatforms, devicePlatform, allLocations, insideCorpnet, allTrustedLocations, locationId, allDevices, deviceFilter, deviceState, unknownFutureValue, deviceFilterIncludeRuleNotMatched, allDeviceStates. Note that you must use the Prefer: include-unknown-enum-members request header to get the following values in this evolvable enum: deviceFilterIncludeRuleNotMatched, allDeviceStates.
func (m *ConditionalAccessRuleSatisfied) GetRuleSatisfied()(*ConditionalAccessRule) {
    if m == nil {
        return nil
    } else {
        return m.ruleSatisfied
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ConditionalAccessRuleSatisfied) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["conditionalAccessCondition"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseConditionalAccessConditions)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConditionalAccessCondition(val.(*ConditionalAccessConditions))
        }
        return nil
    }
    res["ruleSatisfied"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseConditionalAccessRule)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRuleSatisfied(val.(*ConditionalAccessRule))
        }
        return nil
    }
    return res
}
func (m *ConditionalAccessRuleSatisfied) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ConditionalAccessRuleSatisfied) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetConditionalAccessCondition() != nil {
        cast := (*m.GetConditionalAccessCondition()).String()
        err := writer.WriteStringValue("conditionalAccessCondition", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetRuleSatisfied() != nil {
        cast := (*m.GetRuleSatisfied()).String()
        err := writer.WriteStringValue("ruleSatisfied", &cast)
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
func (m *ConditionalAccessRuleSatisfied) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetConditionalAccessCondition sets the conditionalAccessCondition property value. Refers to the conditional access policy conditions that are satisfied. The possible values are: none, application, users, devicePlatform, location, clientType, signInRisk, userRisk, time, deviceState, client, ipAddressSeenByAzureAD, ipAddressSeenByResourceProvider, unknownFutureValue, servicePrincipals, servicePrincipalRisk. Note that you must use the Prefer: include-unknown-enum-members request header to get the following values in this evolvable enum: servicePrincipals, servicePrincipalRisk.
func (m *ConditionalAccessRuleSatisfied) SetConditionalAccessCondition(value *ConditionalAccessConditions)() {
    if m != nil {
        m.conditionalAccessCondition = value
    }
}
// SetRuleSatisfied sets the ruleSatisfied property value. Refers to the conditional access policy conditions that were satisfied. The possible values are: allApps, firstPartyApps, office365, appId, acr, appFilter, allUsers, guest, groupId, roleId, userId, allDevicePlatforms, devicePlatform, allLocations, insideCorpnet, allTrustedLocations, locationId, allDevices, deviceFilter, deviceState, unknownFutureValue, deviceFilterIncludeRuleNotMatched, allDeviceStates. Note that you must use the Prefer: include-unknown-enum-members request header to get the following values in this evolvable enum: deviceFilterIncludeRuleNotMatched, allDeviceStates.
func (m *ConditionalAccessRuleSatisfied) SetRuleSatisfied(value *ConditionalAccessRule)() {
    if m != nil {
        m.ruleSatisfied = value
    }
}
