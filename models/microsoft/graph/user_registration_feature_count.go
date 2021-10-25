package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type UserRegistrationFeatureCount struct {
    additionalData map[string]interface{};
    feature *AuthenticationMethodFeature;
    userCount *int64;
}
func NewUserRegistrationFeatureCount()(*UserRegistrationFeatureCount) {
    m := &UserRegistrationFeatureCount{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *UserRegistrationFeatureCount) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *UserRegistrationFeatureCount) GetFeature()(*AuthenticationMethodFeature) {
    if m == nil {
        return nil
    } else {
        return m.feature
    }
}
func (m *UserRegistrationFeatureCount) GetUserCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.userCount
    }
}
func (m *UserRegistrationFeatureCount) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["feature"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAuthenticationMethodFeature)
        if err != nil {
            return err
        }
        cast := val.(AuthenticationMethodFeature)
        m.SetFeature(&cast)
        return nil
    }
    res["userCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetUserCount(val)
        return nil
    }
    return res
}
func (m *UserRegistrationFeatureCount) IsNil()(bool) {
    return m == nil
}
func (m *UserRegistrationFeatureCount) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetFeature() != nil {
        cast := m.GetFeature().String()
        err := writer.WriteStringValue("feature", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("userCount", m.GetUserCount())
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
func (m *UserRegistrationFeatureCount) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *UserRegistrationFeatureCount) SetFeature(value *AuthenticationMethodFeature)() {
    m.feature = value
}
func (m *UserRegistrationFeatureCount) SetUserCount(value *int64)() {
    m.userCount = value
}
