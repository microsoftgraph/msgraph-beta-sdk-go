package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserRegistrationMethodCount 
type UserRegistrationMethodCount struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Name of authentication method.
    authenticationMethod *string;
    // Number of users registered.
    userCount *int64;
}
// NewUserRegistrationMethodCount instantiates a new userRegistrationMethodCount and sets the default values.
func NewUserRegistrationMethodCount()(*UserRegistrationMethodCount) {
    m := &UserRegistrationMethodCount{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserRegistrationMethodCount) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAuthenticationMethod gets the authenticationMethod property value. Name of authentication method.
func (m *UserRegistrationMethodCount) GetAuthenticationMethod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.authenticationMethod
    }
}
// GetUserCount gets the userCount property value. Number of users registered.
func (m *UserRegistrationMethodCount) GetUserCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.userCount
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserRegistrationMethodCount) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["authenticationMethod"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationMethod(val)
        }
        return nil
    }
    res["userCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserCount(val)
        }
        return nil
    }
    return res
}
func (m *UserRegistrationMethodCount) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *UserRegistrationMethodCount) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("authenticationMethod", m.GetAuthenticationMethod())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserRegistrationMethodCount) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAuthenticationMethod sets the authenticationMethod property value. Name of authentication method.
func (m *UserRegistrationMethodCount) SetAuthenticationMethod(value *string)() {
    if m != nil {
        m.authenticationMethod = value
    }
}
// SetUserCount sets the userCount property value. Number of users registered.
func (m *UserRegistrationMethodCount) SetUserCount(value *int64)() {
    if m != nil {
        m.userCount = value
    }
}
