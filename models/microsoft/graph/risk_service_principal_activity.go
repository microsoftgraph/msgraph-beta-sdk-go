package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RiskServicePrincipalActivity 
type RiskServicePrincipalActivity struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Details of the detected risk. Note: Details for this property are only available for Azure AD Premium P2 customers. P1 customers will be returned hidden. The possible values are: none, adminGeneratedTemporaryPassword, userPerformedSecuredPasswordChange, userPerformedSecuredPasswordReset, adminConfirmedSigninSafe, aiConfirmedSigninSafe, userPassedMFADrivenByRiskBasedPolicy, adminDismissedAllRiskForUser, adminConfirmedSigninCompromised, hidden,  adminConfirmedUserCompromised, unknownFutureValue, adminConfirmedServicePrincipalCompromised, adminDismissedAllRiskForServicePrincipal. Note that you must use the Prefer: include-unknown-enum-members request header to get the following value(s) in this evolvable enum: adminConfirmedServicePrincipalCompromised , adminDismissedAllRiskForServicePrincipal.
    detail *RiskDetail;
    // 
    riskEventTypes []string;
}
// NewRiskServicePrincipalActivity instantiates a new riskServicePrincipalActivity and sets the default values.
func NewRiskServicePrincipalActivity()(*RiskServicePrincipalActivity) {
    m := &RiskServicePrincipalActivity{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RiskServicePrincipalActivity) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDetail gets the detail property value. Details of the detected risk. Note: Details for this property are only available for Azure AD Premium P2 customers. P1 customers will be returned hidden. The possible values are: none, adminGeneratedTemporaryPassword, userPerformedSecuredPasswordChange, userPerformedSecuredPasswordReset, adminConfirmedSigninSafe, aiConfirmedSigninSafe, userPassedMFADrivenByRiskBasedPolicy, adminDismissedAllRiskForUser, adminConfirmedSigninCompromised, hidden,  adminConfirmedUserCompromised, unknownFutureValue, adminConfirmedServicePrincipalCompromised, adminDismissedAllRiskForServicePrincipal. Note that you must use the Prefer: include-unknown-enum-members request header to get the following value(s) in this evolvable enum: adminConfirmedServicePrincipalCompromised , adminDismissedAllRiskForServicePrincipal.
func (m *RiskServicePrincipalActivity) GetDetail()(*RiskDetail) {
    if m == nil {
        return nil
    } else {
        return m.detail
    }
}
// GetRiskEventTypes gets the riskEventTypes property value. 
func (m *RiskServicePrincipalActivity) GetRiskEventTypes()([]string) {
    if m == nil {
        return nil
    } else {
        return m.riskEventTypes
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RiskServicePrincipalActivity) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["detail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRiskDetail)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetail(val.(*RiskDetail))
        }
        return nil
    }
    res["riskEventTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetRiskEventTypes(res)
        }
        return nil
    }
    return res
}
func (m *RiskServicePrincipalActivity) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *RiskServicePrincipalActivity) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetDetail() != nil {
        cast := (*m.GetDetail()).String()
        err := writer.WriteStringValue("detail", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetRiskEventTypes() != nil {
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RiskServicePrincipalActivity) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDetail sets the detail property value. Details of the detected risk. Note: Details for this property are only available for Azure AD Premium P2 customers. P1 customers will be returned hidden. The possible values are: none, adminGeneratedTemporaryPassword, userPerformedSecuredPasswordChange, userPerformedSecuredPasswordReset, adminConfirmedSigninSafe, aiConfirmedSigninSafe, userPassedMFADrivenByRiskBasedPolicy, adminDismissedAllRiskForUser, adminConfirmedSigninCompromised, hidden,  adminConfirmedUserCompromised, unknownFutureValue, adminConfirmedServicePrincipalCompromised, adminDismissedAllRiskForServicePrincipal. Note that you must use the Prefer: include-unknown-enum-members request header to get the following value(s) in this evolvable enum: adminConfirmedServicePrincipalCompromised , adminDismissedAllRiskForServicePrincipal.
func (m *RiskServicePrincipalActivity) SetDetail(value *RiskDetail)() {
    if m != nil {
        m.detail = value
    }
}
// SetRiskEventTypes sets the riskEventTypes property value. 
func (m *RiskServicePrincipalActivity) SetRiskEventTypes(value []string)() {
    if m != nil {
        m.riskEventTypes = value
    }
}
