package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type UserRegistrationCount struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Provides the registration count for your tenant.
    registrationCount *int64;
    // Represents the status of user registration. Possible values are: registered, enabled, capable, and mfaRegistered.
    registrationStatus *RegistrationStatusType;
}
// Instantiates a new userRegistrationCount and sets the default values.
func NewUserRegistrationCount()(*UserRegistrationCount) {
    m := &UserRegistrationCount{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserRegistrationCount) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the registrationCount property value. Provides the registration count for your tenant.
func (m *UserRegistrationCount) GetRegistrationCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.registrationCount
    }
}
// Gets the registrationStatus property value. Represents the status of user registration. Possible values are: registered, enabled, capable, and mfaRegistered.
func (m *UserRegistrationCount) GetRegistrationStatus()(*RegistrationStatusType) {
    if m == nil {
        return nil
    } else {
        return m.registrationStatus
    }
}
// The deserialization information for the current model
func (m *UserRegistrationCount) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["registrationCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetRegistrationCount(val)
        return nil
    }
    res["registrationStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRegistrationStatusType)
        if err != nil {
            return err
        }
        cast := val.(RegistrationStatusType)
        m.SetRegistrationStatus(&cast)
        return nil
    }
    return res
}
func (m *UserRegistrationCount) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *UserRegistrationCount) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("registrationCount", m.GetRegistrationCount())
        if err != nil {
            return err
        }
    }
    if m.GetRegistrationStatus() != nil {
        cast := m.GetRegistrationStatus().String()
        err := writer.WriteStringValue("registrationStatus", &cast)
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
func (m *UserRegistrationCount) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the registrationCount property value. Provides the registration count for your tenant.
// Parameters:
//  - value : Value to set for the registrationCount property.
func (m *UserRegistrationCount) SetRegistrationCount(value *int64)() {
    m.registrationCount = value
}
// Sets the registrationStatus property value. Represents the status of user registration. Possible values are: registered, enabled, capable, and mfaRegistered.
// Parameters:
//  - value : Value to set for the registrationStatus property.
func (m *UserRegistrationCount) SetRegistrationStatus(value *RegistrationStatusType)() {
    m.registrationStatus = value
}
