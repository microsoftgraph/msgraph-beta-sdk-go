package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserRegistrationMethodSummary provides operations to call the usersRegisteredByMethod method.
type UserRegistrationMethodSummary struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Total number of users in the tenant.
    totalUserCount *int64;
    // Number of users registered for each authentication method.
    userRegistrationMethodCounts []UserRegistrationMethodCountable;
    // User role type. Possible values are: all, privilegedAdmin, admin, user.
    userRoles *IncludedUserRoles;
    // User type. Possible values are: all, member, guest.
    userTypes *IncludedUserTypes;
}
// NewUserRegistrationMethodSummary instantiates a new UserRegistrationMethodSummary and sets the default values.
func NewUserRegistrationMethodSummary()(*UserRegistrationMethodSummary) {
    m := &UserRegistrationMethodSummary{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUserRegistrationMethodSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserRegistrationMethodSummaryFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewUserRegistrationMethodSummary(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserRegistrationMethodSummary) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserRegistrationMethodSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    res["userRegistrationMethodCounts"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserRegistrationMethodCountFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserRegistrationMethodCountable, len(val))
            for i, v := range val {
                res[i] = v.(UserRegistrationMethodCountable)
            }
            m.SetUserRegistrationMethodCounts(res)
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
// GetTotalUserCount gets the totalUserCount property value. Total number of users in the tenant.
func (m *UserRegistrationMethodSummary) GetTotalUserCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.totalUserCount
    }
}
// GetUserRegistrationMethodCounts gets the userRegistrationMethodCounts property value. Number of users registered for each authentication method.
func (m *UserRegistrationMethodSummary) GetUserRegistrationMethodCounts()([]UserRegistrationMethodCountable) {
    if m == nil {
        return nil
    } else {
        return m.userRegistrationMethodCounts
    }
}
// GetUserRoles gets the userRoles property value. User role type. Possible values are: all, privilegedAdmin, admin, user.
func (m *UserRegistrationMethodSummary) GetUserRoles()(*IncludedUserRoles) {
    if m == nil {
        return nil
    } else {
        return m.userRoles
    }
}
// GetUserTypes gets the userTypes property value. User type. Possible values are: all, member, guest.
func (m *UserRegistrationMethodSummary) GetUserTypes()(*IncludedUserTypes) {
    if m == nil {
        return nil
    } else {
        return m.userTypes
    }
}
func (m *UserRegistrationMethodSummary) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *UserRegistrationMethodSummary) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("totalUserCount", m.GetTotalUserCount())
        if err != nil {
            return err
        }
    }
    if m.GetUserRegistrationMethodCounts() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserRegistrationMethodCounts()))
        for i, v := range m.GetUserRegistrationMethodCounts() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("userRegistrationMethodCounts", cast)
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
func (m *UserRegistrationMethodSummary) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetTotalUserCount sets the totalUserCount property value. Total number of users in the tenant.
func (m *UserRegistrationMethodSummary) SetTotalUserCount(value *int64)() {
    if m != nil {
        m.totalUserCount = value
    }
}
// SetUserRegistrationMethodCounts sets the userRegistrationMethodCounts property value. Number of users registered for each authentication method.
func (m *UserRegistrationMethodSummary) SetUserRegistrationMethodCounts(value []UserRegistrationMethodCountable)() {
    if m != nil {
        m.userRegistrationMethodCounts = value
    }
}
// SetUserRoles sets the userRoles property value. User role type. Possible values are: all, privilegedAdmin, admin, user.
func (m *UserRegistrationMethodSummary) SetUserRoles(value *IncludedUserRoles)() {
    if m != nil {
        m.userRoles = value
    }
}
// SetUserTypes sets the userTypes property value. User type. Possible values are: all, member, guest.
func (m *UserRegistrationMethodSummary) SetUserTypes(value *IncludedUserTypes)() {
    if m != nil {
        m.userTypes = value
    }
}
