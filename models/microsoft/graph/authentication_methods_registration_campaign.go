package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AuthenticationMethodsRegistrationCampaign struct {
    additionalData map[string]interface{};
    excludeTargets []ExcludeTarget;
    includeTargets []AuthenticationMethodsRegistrationCampaignIncludeTarget;
    snoozeDurationInDays *int32;
    state *AdvancedConfigState;
}
func NewAuthenticationMethodsRegistrationCampaign()(*AuthenticationMethodsRegistrationCampaign) {
    m := &AuthenticationMethodsRegistrationCampaign{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *AuthenticationMethodsRegistrationCampaign) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *AuthenticationMethodsRegistrationCampaign) GetExcludeTargets()([]ExcludeTarget) {
    if m == nil {
        return nil
    } else {
        return m.excludeTargets
    }
}
func (m *AuthenticationMethodsRegistrationCampaign) GetIncludeTargets()([]AuthenticationMethodsRegistrationCampaignIncludeTarget) {
    if m == nil {
        return nil
    } else {
        return m.includeTargets
    }
}
func (m *AuthenticationMethodsRegistrationCampaign) GetSnoozeDurationInDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.snoozeDurationInDays
    }
}
func (m *AuthenticationMethodsRegistrationCampaign) GetState()(*AdvancedConfigState) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
func (m *AuthenticationMethodsRegistrationCampaign) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["excludeTargets"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewExcludeTarget() })
        if err != nil {
            return err
        }
        res := make([]ExcludeTarget, len(val))
        for i, v := range val {
            res[i] = *(v.(*ExcludeTarget))
        }
        m.SetExcludeTargets(res)
        return nil
    }
    res["includeTargets"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAuthenticationMethodsRegistrationCampaignIncludeTarget() })
        if err != nil {
            return err
        }
        res := make([]AuthenticationMethodsRegistrationCampaignIncludeTarget, len(val))
        for i, v := range val {
            res[i] = *(v.(*AuthenticationMethodsRegistrationCampaignIncludeTarget))
        }
        m.SetIncludeTargets(res)
        return nil
    }
    res["snoozeDurationInDays"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetSnoozeDurationInDays(val)
        return nil
    }
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAdvancedConfigState)
        if err != nil {
            return err
        }
        cast := val.(AdvancedConfigState)
        m.SetState(&cast)
        return nil
    }
    return res
}
func (m *AuthenticationMethodsRegistrationCampaign) IsNil()(bool) {
    return m == nil
}
func (m *AuthenticationMethodsRegistrationCampaign) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetExcludeTargets()))
        for i, v := range m.GetExcludeTargets() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("excludeTargets", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetIncludeTargets()))
        for i, v := range m.GetIncludeTargets() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("includeTargets", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("snoozeDurationInDays", m.GetSnoozeDurationInDays())
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := m.GetState().String()
        err := writer.WriteStringValue("state", &cast)
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
func (m *AuthenticationMethodsRegistrationCampaign) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *AuthenticationMethodsRegistrationCampaign) SetExcludeTargets(value []ExcludeTarget)() {
    m.excludeTargets = value
}
func (m *AuthenticationMethodsRegistrationCampaign) SetIncludeTargets(value []AuthenticationMethodsRegistrationCampaignIncludeTarget)() {
    m.includeTargets = value
}
func (m *AuthenticationMethodsRegistrationCampaign) SetSnoozeDurationInDays(value *int32)() {
    m.snoozeDurationInDays = value
}
func (m *AuthenticationMethodsRegistrationCampaign) SetState(value *AdvancedConfigState)() {
    m.state = value
}
