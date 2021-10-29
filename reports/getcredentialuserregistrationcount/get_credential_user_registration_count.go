package getcredentialuserregistrationcount

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// 
type GetCredentialUserRegistrationCount struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    // Provides the total user count in the tenant.
    totalUserCount *int64;
    // A collection of registration count and status information for users in your tenant.
    userRegistrationCounts []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserRegistrationCount;
}
// Instantiates a new getCredentialUserRegistrationCount and sets the default values.
func NewGetCredentialUserRegistrationCount()(*GetCredentialUserRegistrationCount) {
    m := &GetCredentialUserRegistrationCount{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
// Gets the totalUserCount property value. Provides the total user count in the tenant.
func (m *GetCredentialUserRegistrationCount) GetTotalUserCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.totalUserCount
    }
}
// Gets the userRegistrationCounts property value. A collection of registration count and status information for users in your tenant.
func (m *GetCredentialUserRegistrationCount) GetUserRegistrationCounts()([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserRegistrationCount) {
    if m == nil {
        return nil
    } else {
        return m.userRegistrationCounts
    }
}
// The deserialization information for the current model
func (m *GetCredentialUserRegistrationCount) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["totalUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetTotalUserCount(val)
        return nil
    }
    res["userRegistrationCounts"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewUserRegistrationCount() })
        if err != nil {
            return err
        }
        res := make([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserRegistrationCount, len(val))
        for i, v := range val {
            res[i] = *(v.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserRegistrationCount))
        }
        m.SetUserRegistrationCounts(res)
        return nil
    }
    return res
}
func (m *GetCredentialUserRegistrationCount) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *GetCredentialUserRegistrationCount) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
    {
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
// Sets the totalUserCount property value. Provides the total user count in the tenant.
// Parameters:
//  - value : Value to set for the totalUserCount property.
func (m *GetCredentialUserRegistrationCount) SetTotalUserCount(value *int64)() {
    m.totalUserCount = value
}
// Sets the userRegistrationCounts property value. A collection of registration count and status information for users in your tenant.
// Parameters:
//  - value : Value to set for the userRegistrationCounts property.
func (m *GetCredentialUserRegistrationCount) SetUserRegistrationCounts(value []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserRegistrationCount)() {
    m.userRegistrationCounts = value
}
