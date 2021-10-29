package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type RiskUserActivity struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Details of the detected risk. Possible values are: none, adminGeneratedTemporaryPassword, userPerformedSecuredPasswordChange, userPerformedSecuredPasswordReset, adminConfirmedSigninSafe, aiConfirmedSigninSafe, userPassedMFADrivenByRiskBasedPolicy, adminDismissedAllRiskForUser, adminConfirmedSigninCompromised, hidden, adminConfirmedUserCompromised, unknownFutureValue.
    detail *RiskDetail;
    // List of risk event types. Deprecated. Use riskEventType instead.
    eventTypes []RiskEventType;
    // The type of risk event detected.
    riskEventTypes []string;
}
// Instantiates a new riskUserActivity and sets the default values.
func NewRiskUserActivity()(*RiskUserActivity) {
    m := &RiskUserActivity{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RiskUserActivity) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the detail property value. Details of the detected risk. Possible values are: none, adminGeneratedTemporaryPassword, userPerformedSecuredPasswordChange, userPerformedSecuredPasswordReset, adminConfirmedSigninSafe, aiConfirmedSigninSafe, userPassedMFADrivenByRiskBasedPolicy, adminDismissedAllRiskForUser, adminConfirmedSigninCompromised, hidden, adminConfirmedUserCompromised, unknownFutureValue.
func (m *RiskUserActivity) GetDetail()(*RiskDetail) {
    if m == nil {
        return nil
    } else {
        return m.detail
    }
}
// Gets the eventTypes property value. List of risk event types. Deprecated. Use riskEventType instead.
func (m *RiskUserActivity) GetEventTypes()([]RiskEventType) {
    if m == nil {
        return nil
    } else {
        return m.eventTypes
    }
}
// Gets the riskEventTypes property value. The type of risk event detected.
func (m *RiskUserActivity) GetRiskEventTypes()([]string) {
    if m == nil {
        return nil
    } else {
        return m.riskEventTypes
    }
}
// The deserialization information for the current model
func (m *RiskUserActivity) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["detail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRiskDetail)
        if err != nil {
            return err
        }
        cast := val.(RiskDetail)
        m.SetDetail(&cast)
        return nil
    }
    res["eventTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseRiskEventType)
        if err != nil {
            return err
        }
        res := make([]RiskEventType, len(val))
        for i, v := range val {
            res[i] = *(v.(*RiskEventType))
        }
        m.SetEventTypes(res)
        return nil
    }
    res["riskEventTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetRiskEventTypes(res)
        return nil
    }
    return res
}
func (m *RiskUserActivity) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *RiskUserActivity) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetDetail() != nil {
        cast := m.GetDetail().String()
        err := writer.WriteStringValue("detail", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("eventTypes", SerializeRiskEventType(m.GetEventTypes()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("riskEventTypes", m.GetRiskEventTypes())
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
func (m *RiskUserActivity) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the detail property value. Details of the detected risk. Possible values are: none, adminGeneratedTemporaryPassword, userPerformedSecuredPasswordChange, userPerformedSecuredPasswordReset, adminConfirmedSigninSafe, aiConfirmedSigninSafe, userPassedMFADrivenByRiskBasedPolicy, adminDismissedAllRiskForUser, adminConfirmedSigninCompromised, hidden, adminConfirmedUserCompromised, unknownFutureValue.
// Parameters:
//  - value : Value to set for the detail property.
func (m *RiskUserActivity) SetDetail(value *RiskDetail)() {
    m.detail = value
}
// Sets the eventTypes property value. List of risk event types. Deprecated. Use riskEventType instead.
// Parameters:
//  - value : Value to set for the eventTypes property.
func (m *RiskUserActivity) SetEventTypes(value []RiskEventType)() {
    m.eventTypes = value
}
// Sets the riskEventTypes property value. The type of risk event detected.
// Parameters:
//  - value : Value to set for the riskEventTypes property.
func (m *RiskUserActivity) SetRiskEventTypes(value []string)() {
    m.riskEventTypes = value
}
