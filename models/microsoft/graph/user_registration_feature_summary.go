package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type UserRegistrationFeatureSummary struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Total number of users accounts, excluding those that are blocked
    totalUserCount *int64;
    // Number of users registered or capable for Multi-Factor Authentication, Self-Service Password Reset and Passwordless Authentication.
    userRegistrationFeatureCounts []UserRegistrationFeatureCount;
    // User role type. Possible values are: all, privilegedAdmin, admin, user.
    userRoles *IncludedUserRoles;
    // User type. Possible values are: all, member, guest.
    userTypes *IncludedUserTypes;
}
// Instantiates a new UserRegistrationFeatureSummary and sets the default values.
func NewUserRegistrationFeatureSummary()(*UserRegistrationFeatureSummary) {
    m := &UserRegistrationFeatureSummary{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserRegistrationFeatureSummary) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the totalUserCount property value. Total number of users accounts, excluding those that are blocked
func (m *UserRegistrationFeatureSummary) GetTotalUserCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.totalUserCount
    }
}
// Gets the userRegistrationFeatureCounts property value. Number of users registered or capable for Multi-Factor Authentication, Self-Service Password Reset and Passwordless Authentication.
func (m *UserRegistrationFeatureSummary) GetUserRegistrationFeatureCounts()([]UserRegistrationFeatureCount) {
    if m == nil {
        return nil
    } else {
        return m.userRegistrationFeatureCounts
    }
}
// Gets the userRoles property value. User role type. Possible values are: all, privilegedAdmin, admin, user.
func (m *UserRegistrationFeatureSummary) GetUserRoles()(*IncludedUserRoles) {
    if m == nil {
        return nil
    } else {
        return m.userRoles
    }
}
// Gets the userTypes property value. User type. Possible values are: all, member, guest.
func (m *UserRegistrationFeatureSummary) GetUserTypes()(*IncludedUserTypes) {
    if m == nil {
        return nil
    } else {
        return m.userTypes
    }
}
// The deserialization information for the current model
func (m *UserRegistrationFeatureSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["totalUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetTotalUserCount(val)
        return nil
    }
    res["userRegistrationFeatureCounts"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserRegistrationFeatureCount() })
        if err != nil {
            return err
        }
        res := make([]UserRegistrationFeatureCount, len(val))
        for i, v := range val {
            res[i] = *(v.(*UserRegistrationFeatureCount))
        }
        m.SetUserRegistrationFeatureCounts(res)
        return nil
    }
    res["userRoles"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseIncludedUserRoles)
        if err != nil {
            return err
        }
        cast := val.(IncludedUserRoles)
        m.SetUserRoles(&cast)
        return nil
    }
    res["userTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseIncludedUserTypes)
        if err != nil {
            return err
        }
        cast := val.(IncludedUserTypes)
        m.SetUserTypes(&cast)
        return nil
    }
    return res
}
func (m *UserRegistrationFeatureSummary) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *UserRegistrationFeatureSummary) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("totalUserCount", m.GetTotalUserCount())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserRegistrationFeatureCounts()))
        for i, v := range m.GetUserRegistrationFeatureCounts() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("userRegistrationFeatureCounts", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserRoles() != nil {
        cast := m.GetUserRoles().String()
        err := writer.WriteStringValue("userRoles", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserTypes() != nil {
        cast := m.GetUserTypes().String()
        err := writer.WriteStringValue("userTypes", &cast)
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
func (m *UserRegistrationFeatureSummary) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the totalUserCount property value. Total number of users accounts, excluding those that are blocked
// Parameters:
//  - value : Value to set for the totalUserCount property.
func (m *UserRegistrationFeatureSummary) SetTotalUserCount(value *int64)() {
    m.totalUserCount = value
}
// Sets the userRegistrationFeatureCounts property value. Number of users registered or capable for Multi-Factor Authentication, Self-Service Password Reset and Passwordless Authentication.
// Parameters:
//  - value : Value to set for the userRegistrationFeatureCounts property.
func (m *UserRegistrationFeatureSummary) SetUserRegistrationFeatureCounts(value []UserRegistrationFeatureCount)() {
    m.userRegistrationFeatureCounts = value
}
// Sets the userRoles property value. User role type. Possible values are: all, privilegedAdmin, admin, user.
// Parameters:
//  - value : Value to set for the userRoles property.
func (m *UserRegistrationFeatureSummary) SetUserRoles(value *IncludedUserRoles)() {
    m.userRoles = value
}
// Sets the userTypes property value. User type. Possible values are: all, member, guest.
// Parameters:
//  - value : Value to set for the userTypes property.
func (m *UserRegistrationFeatureSummary) SetUserTypes(value *IncludedUserTypes)() {
    m.userTypes = value
}
