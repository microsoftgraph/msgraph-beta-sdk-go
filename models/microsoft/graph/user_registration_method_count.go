package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type UserRegistrationMethodCount struct {
    additionalData map[string]interface{};
    authenticationMethod *string;
    userCount *int64;
}
func NewUserRegistrationMethodCount()(*UserRegistrationMethodCount) {
    m := &UserRegistrationMethodCount{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *UserRegistrationMethodCount) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *UserRegistrationMethodCount) GetAuthenticationMethod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.authenticationMethod
    }
}
func (m *UserRegistrationMethodCount) GetUserCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.userCount
    }
}
func (m *UserRegistrationMethodCount) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["authenticationMethod"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAuthenticationMethod(val)
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
func (m *UserRegistrationMethodCount) IsNil()(bool) {
    return m == nil
}
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
func (m *UserRegistrationMethodCount) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *UserRegistrationMethodCount) SetAuthenticationMethod(value *string)() {
    m.authenticationMethod = value
}
func (m *UserRegistrationMethodCount) SetUserCount(value *int64)() {
    m.userCount = value
}
