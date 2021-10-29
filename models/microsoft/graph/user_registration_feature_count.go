package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type UserRegistrationFeatureCount struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Number of users registered or capable for Multi-Factor Authentication, Self-Service Password Reset and Passwordless Authentication. Possible values are: ssprRegistered, ssprEnabled, ssprCapable, passwordlessCapable, mfaCapable.
    feature *AuthenticationMethodFeature;
    // Number of users.
    userCount *int64;
}
// Instantiates a new userRegistrationFeatureCount and sets the default values.
func NewUserRegistrationFeatureCount()(*UserRegistrationFeatureCount) {
    m := &UserRegistrationFeatureCount{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserRegistrationFeatureCount) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the feature property value. Number of users registered or capable for Multi-Factor Authentication, Self-Service Password Reset and Passwordless Authentication. Possible values are: ssprRegistered, ssprEnabled, ssprCapable, passwordlessCapable, mfaCapable.
func (m *UserRegistrationFeatureCount) GetFeature()(*AuthenticationMethodFeature) {
    if m == nil {
        return nil
    } else {
        return m.feature
    }
}
// Gets the userCount property value. Number of users.
func (m *UserRegistrationFeatureCount) GetUserCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.userCount
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *UserRegistrationFeatureCount) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the feature property value. Number of users registered or capable for Multi-Factor Authentication, Self-Service Password Reset and Passwordless Authentication. Possible values are: ssprRegistered, ssprEnabled, ssprCapable, passwordlessCapable, mfaCapable.
// Parameters:
//  - value : Value to set for the feature property.
func (m *UserRegistrationFeatureCount) SetFeature(value *AuthenticationMethodFeature)() {
    m.feature = value
}
// Sets the userCount property value. Number of users.
// Parameters:
//  - value : Value to set for the userCount property.
func (m *UserRegistrationFeatureCount) SetUserCount(value *int64)() {
    m.userCount = value
}
