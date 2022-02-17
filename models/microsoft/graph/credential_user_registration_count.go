package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CredentialUserRegistrationCount 
type CredentialUserRegistrationCount struct {
    Entity
    // Provides the total user count in the tenant.
    totalUserCount *int64;
    // A collection of registration count and status information for users in your tenant.
    userRegistrationCounts []UserRegistrationCount;
}
// NewCredentialUserRegistrationCount instantiates a new CredentialUserRegistrationCount and sets the default values.
func NewCredentialUserRegistrationCount()(*CredentialUserRegistrationCount) {
    m := &CredentialUserRegistrationCount{
        Entity: *NewEntity(),
    }
    return m
}
// GetTotalUserCount gets the totalUserCount property value. Provides the total user count in the tenant.
func (m *CredentialUserRegistrationCount) GetTotalUserCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.totalUserCount
    }
}
// GetUserRegistrationCounts gets the userRegistrationCounts property value. A collection of registration count and status information for users in your tenant.
func (m *CredentialUserRegistrationCount) GetUserRegistrationCounts()([]UserRegistrationCount) {
    if m == nil {
        return nil
    } else {
        return m.userRegistrationCounts
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CredentialUserRegistrationCount) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["userRegistrationCounts"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserRegistrationCount() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserRegistrationCount, len(val))
            for i, v := range val {
                res[i] = *(v.(*UserRegistrationCount))
            }
            m.SetUserRegistrationCounts(res)
        }
        return nil
    }
    return res
}
func (m *CredentialUserRegistrationCount) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *CredentialUserRegistrationCount) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("totalUserCount", m.GetTotalUserCount())
        if err != nil {
            return err
        }
    }
    if m.GetUserRegistrationCounts() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserRegistrationCounts()))
        for i, v := range m.GetUserRegistrationCounts() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("userRegistrationCounts", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetTotalUserCount sets the totalUserCount property value. Provides the total user count in the tenant.
func (m *CredentialUserRegistrationCount) SetTotalUserCount(value *int64)() {
    if m != nil {
        m.totalUserCount = value
    }
}
// SetUserRegistrationCounts sets the userRegistrationCounts property value. A collection of registration count and status information for users in your tenant.
func (m *CredentialUserRegistrationCount) SetUserRegistrationCounts(value []UserRegistrationCount)() {
    if m != nil {
        m.userRegistrationCounts = value
    }
}
