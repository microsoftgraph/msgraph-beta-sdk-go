package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type UserRegistrationFeatureSummary struct {
    additionalData map[string]interface{};
    totalUserCount *int64;
    userRegistrationFeatureCounts []UserRegistrationFeatureCount;
    userRoles *IncludedUserRoles;
    userTypes *IncludedUserTypes;
}
func NewUserRegistrationFeatureSummary()(*UserRegistrationFeatureSummary) {
    m := &UserRegistrationFeatureSummary{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *UserRegistrationFeatureSummary) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *UserRegistrationFeatureSummary) GetTotalUserCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.totalUserCount
    }
}
func (m *UserRegistrationFeatureSummary) GetUserRegistrationFeatureCounts()([]UserRegistrationFeatureCount) {
    if m == nil {
        return nil
    } else {
        return m.userRegistrationFeatureCounts
    }
}
func (m *UserRegistrationFeatureSummary) GetUserRoles()(*IncludedUserRoles) {
    if m == nil {
        return nil
    } else {
        return m.userRoles
    }
}
func (m *UserRegistrationFeatureSummary) GetUserTypes()(*IncludedUserTypes) {
    if m == nil {
        return nil
    } else {
        return m.userTypes
    }
}
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
func (m *UserRegistrationFeatureSummary) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *UserRegistrationFeatureSummary) SetTotalUserCount(value *int64)() {
    m.totalUserCount = value
}
func (m *UserRegistrationFeatureSummary) SetUserRegistrationFeatureCounts(value []UserRegistrationFeatureCount)() {
    m.userRegistrationFeatureCounts = value
}
func (m *UserRegistrationFeatureSummary) SetUserRoles(value *IncludedUserRoles)() {
    m.userRoles = value
}
func (m *UserRegistrationFeatureSummary) SetUserTypes(value *IncludedUserTypes)() {
    m.userTypes = value
}
