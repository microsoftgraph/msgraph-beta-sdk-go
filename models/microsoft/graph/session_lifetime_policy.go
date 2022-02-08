package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SessionLifetimePolicy 
type SessionLifetimePolicy struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The human-readable details of the conditional access session management policy applied to the sign-in.
    detail *string;
    // If a conditional access session management policy required the user to authenticate in this sign-in event, this field describes the policy type that required authentication. The possible values are: rememberMultifactorAuthenticationOnTrustedDevices, tenantTokenLifetimePolicy, audienceTokenLifetimePolicy, signInFrequencyPeriodicReauthentication, ngcMfa, signInFrequencyEveryTime, unknownFutureValue.
    expirationRequirement *ExpirationRequirement;
}
// NewSessionLifetimePolicy instantiates a new sessionLifetimePolicy and sets the default values.
func NewSessionLifetimePolicy()(*SessionLifetimePolicy) {
    m := &SessionLifetimePolicy{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SessionLifetimePolicy) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDetail gets the detail property value. The human-readable details of the conditional access session management policy applied to the sign-in.
func (m *SessionLifetimePolicy) GetDetail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.detail
    }
}
// GetExpirationRequirement gets the expirationRequirement property value. If a conditional access session management policy required the user to authenticate in this sign-in event, this field describes the policy type that required authentication. The possible values are: rememberMultifactorAuthenticationOnTrustedDevices, tenantTokenLifetimePolicy, audienceTokenLifetimePolicy, signInFrequencyPeriodicReauthentication, ngcMfa, signInFrequencyEveryTime, unknownFutureValue.
func (m *SessionLifetimePolicy) GetExpirationRequirement()(*ExpirationRequirement) {
    if m == nil {
        return nil
    } else {
        return m.expirationRequirement
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SessionLifetimePolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["detail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetail(val)
        }
        return nil
    }
    res["expirationRequirement"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseExpirationRequirement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationRequirement(val.(*ExpirationRequirement))
        }
        return nil
    }
    return res
}
func (m *SessionLifetimePolicy) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *SessionLifetimePolicy) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("detail", m.GetDetail())
        if err != nil {
            return err
        }
    }
    if m.GetExpirationRequirement() != nil {
        cast := (*m.GetExpirationRequirement()).String()
        err := writer.WriteStringValue("expirationRequirement", &cast)
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
func (m *SessionLifetimePolicy) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDetail sets the detail property value. The human-readable details of the conditional access session management policy applied to the sign-in.
func (m *SessionLifetimePolicy) SetDetail(value *string)() {
    if m != nil {
        m.detail = value
    }
}
// SetExpirationRequirement sets the expirationRequirement property value. If a conditional access session management policy required the user to authenticate in this sign-in event, this field describes the policy type that required authentication. The possible values are: rememberMultifactorAuthenticationOnTrustedDevices, tenantTokenLifetimePolicy, audienceTokenLifetimePolicy, signInFrequencyPeriodicReauthentication, ngcMfa, signInFrequencyEveryTime, unknownFutureValue.
func (m *SessionLifetimePolicy) SetExpirationRequirement(value *ExpirationRequirement)() {
    if m != nil {
        m.expirationRequirement = value
    }
}
