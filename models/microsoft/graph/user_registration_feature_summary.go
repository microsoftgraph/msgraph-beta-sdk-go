package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserRegistrationFeatureSummary 
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
// NewUserRegistrationFeatureSummary instantiates a new UserRegistrationFeatureSummary and sets the default values.
func NewUserRegistrationFeatureSummary()(*UserRegistrationFeatureSummary) {
    m := &UserRegistrationFeatureSummary{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserRegistrationFeatureSummary) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetTotalUserCount gets the totalUserCount property value. Total number of users accounts, excluding those that are blocked
func (m *UserRegistrationFeatureSummary) GetTotalUserCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.totalUserCount
    }
}
// GetUserRegistrationFeatureCounts gets the userRegistrationFeatureCounts property value. Number of users registered or capable for Multi-Factor Authentication, Self-Service Password Reset and Passwordless Authentication.
func (m *UserRegistrationFeatureSummary) GetUserRegistrationFeatureCounts()([]UserRegistrationFeatureCount) {
    if m == nil {
        return nil
    } else {
        return m.userRegistrationFeatureCounts
    }
}
// GetUserRoles gets the userRoles property value. User role type. Possible values are: all, privilegedAdmin, admin, user.
func (m *UserRegistrationFeatureSummary) GetUserRoles()(*IncludedUserRoles) {
    if m == nil {
        return nil
    } else {
        return m.userRoles
    }
}
// GetUserTypes gets the userTypes property value. User type. Possible values are: all, member, guest.
func (m *UserRegistrationFeatureSummary) GetUserTypes()(*IncludedUserTypes) {
    if m == nil {
        return nil
    } else {
        return m.userTypes
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserRegistrationFeatureSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["totalUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalUserCount(val)
        }
        return nil
    }
    res["userRegistrationFeatureCounts"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserRegistrationFeatureCount() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserRegistrationFeatureCount, len(val))
            for i, v := range val {
                res[i] = *(v.(*UserRegistrationFeatureCount))
            }
            m.SetUserRegistrationFeatureCounts(res)
        }
        return nil
    }
    res["userRoles"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseIncludedUserRoles)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRoles(val.(*IncludedUserRoles))
        }
        return nil
    }
    res["userTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseIncludedUserTypes)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserTypes(val.(*IncludedUserTypes))
        }
        return nil
    }
    return res
}
func (m *UserRegistrationFeatureSummary) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *UserRegistrationFeatureSummary) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("totalUserCount", m.GetTotalUserCount())
        if err != nil {
            return err
        }
    }
    if m.GetUserRegistrationFeatureCounts() != nil {
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
        cast := (*m.GetUserRoles()).String()
        err := writer.WriteStringValue("userRoles", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserTypes() != nil {
        cast := (*m.GetUserTypes()).String()
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserRegistrationFeatureSummary) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetTotalUserCount sets the totalUserCount property value. Total number of users accounts, excluding those that are blocked
func (m *UserRegistrationFeatureSummary) SetTotalUserCount(value *int64)() {
    if m != nil {
        m.totalUserCount = value
    }
}
// SetUserRegistrationFeatureCounts sets the userRegistrationFeatureCounts property value. Number of users registered or capable for Multi-Factor Authentication, Self-Service Password Reset and Passwordless Authentication.
func (m *UserRegistrationFeatureSummary) SetUserRegistrationFeatureCounts(value []UserRegistrationFeatureCount)() {
    if m != nil {
        m.userRegistrationFeatureCounts = value
    }
}
// SetUserRoles sets the userRoles property value. User role type. Possible values are: all, privilegedAdmin, admin, user.
func (m *UserRegistrationFeatureSummary) SetUserRoles(value *IncludedUserRoles)() {
    if m != nil {
        m.userRoles = value
    }
}
// SetUserTypes sets the userTypes property value. User type. Possible values are: all, member, guest.
func (m *UserRegistrationFeatureSummary) SetUserTypes(value *IncludedUserTypes)() {
    if m != nil {
        m.userTypes = value
    }
}
