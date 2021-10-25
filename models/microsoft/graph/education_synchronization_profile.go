package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type EducationSynchronizationProfile struct {
    Entity
    dataProvider *EducationSynchronizationDataProvider;
    displayName *string;
    errors []EducationSynchronizationError;
    expirationDate *string;
    handleSpecialCharacterConstraint *bool;
    identitySynchronizationConfiguration *EducationIdentitySynchronizationConfiguration;
    licensesToAssign []EducationSynchronizationLicenseAssignment;
    profileStatus *EducationSynchronizationProfileStatus;
    state *EducationSynchronizationProfileState;
}
func NewEducationSynchronizationProfile()(*EducationSynchronizationProfile) {
    m := &EducationSynchronizationProfile{
        Entity: *NewEntity(),
    }
    return m
}
func (m *EducationSynchronizationProfile) GetDataProvider()(*EducationSynchronizationDataProvider) {
    if m == nil {
        return nil
    } else {
        return m.dataProvider
    }
}
func (m *EducationSynchronizationProfile) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *EducationSynchronizationProfile) GetErrors()([]EducationSynchronizationError) {
    if m == nil {
        return nil
    } else {
        return m.errors
    }
}
func (m *EducationSynchronizationProfile) GetExpirationDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.expirationDate
    }
}
func (m *EducationSynchronizationProfile) GetHandleSpecialCharacterConstraint()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.handleSpecialCharacterConstraint
    }
}
func (m *EducationSynchronizationProfile) GetIdentitySynchronizationConfiguration()(*EducationIdentitySynchronizationConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.identitySynchronizationConfiguration
    }
}
func (m *EducationSynchronizationProfile) GetLicensesToAssign()([]EducationSynchronizationLicenseAssignment) {
    if m == nil {
        return nil
    } else {
        return m.licensesToAssign
    }
}
func (m *EducationSynchronizationProfile) GetProfileStatus()(*EducationSynchronizationProfileStatus) {
    if m == nil {
        return nil
    } else {
        return m.profileStatus
    }
}
func (m *EducationSynchronizationProfile) GetState()(*EducationSynchronizationProfileState) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
func (m *EducationSynchronizationProfile) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["dataProvider"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationSynchronizationDataProvider() })
        if err != nil {
            return err
        }
        m.SetDataProvider(val.(*EducationSynchronizationDataProvider))
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["errors"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationSynchronizationError() })
        if err != nil {
            return err
        }
        res := make([]EducationSynchronizationError, len(val))
        for i, v := range val {
            res[i] = *(v.(*EducationSynchronizationError))
        }
        m.SetErrors(res)
        return nil
    }
    res["expirationDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetExpirationDate(val)
        return nil
    }
    res["handleSpecialCharacterConstraint"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetHandleSpecialCharacterConstraint(val)
        return nil
    }
    res["identitySynchronizationConfiguration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationIdentitySynchronizationConfiguration() })
        if err != nil {
            return err
        }
        m.SetIdentitySynchronizationConfiguration(val.(*EducationIdentitySynchronizationConfiguration))
        return nil
    }
    res["licensesToAssign"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationSynchronizationLicenseAssignment() })
        if err != nil {
            return err
        }
        res := make([]EducationSynchronizationLicenseAssignment, len(val))
        for i, v := range val {
            res[i] = *(v.(*EducationSynchronizationLicenseAssignment))
        }
        m.SetLicensesToAssign(res)
        return nil
    }
    res["profileStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationSynchronizationProfileStatus() })
        if err != nil {
            return err
        }
        m.SetProfileStatus(val.(*EducationSynchronizationProfileStatus))
        return nil
    }
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseEducationSynchronizationProfileState)
        if err != nil {
            return err
        }
        cast := val.(EducationSynchronizationProfileState)
        m.SetState(&cast)
        return nil
    }
    return res
}
func (m *EducationSynchronizationProfile) IsNil()(bool) {
    return m == nil
}
func (m *EducationSynchronizationProfile) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("dataProvider", m.GetDataProvider())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetErrors()))
        for i, v := range m.GetErrors() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("errors", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("expirationDate", m.GetExpirationDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("handleSpecialCharacterConstraint", m.GetHandleSpecialCharacterConstraint())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("identitySynchronizationConfiguration", m.GetIdentitySynchronizationConfiguration())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetLicensesToAssign()))
        for i, v := range m.GetLicensesToAssign() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("licensesToAssign", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("profileStatus", m.GetProfileStatus())
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := m.GetState().String()
        err = writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *EducationSynchronizationProfile) SetDataProvider(value *EducationSynchronizationDataProvider)() {
    m.dataProvider = value
}
func (m *EducationSynchronizationProfile) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *EducationSynchronizationProfile) SetErrors(value []EducationSynchronizationError)() {
    m.errors = value
}
func (m *EducationSynchronizationProfile) SetExpirationDate(value *string)() {
    m.expirationDate = value
}
func (m *EducationSynchronizationProfile) SetHandleSpecialCharacterConstraint(value *bool)() {
    m.handleSpecialCharacterConstraint = value
}
func (m *EducationSynchronizationProfile) SetIdentitySynchronizationConfiguration(value *EducationIdentitySynchronizationConfiguration)() {
    m.identitySynchronizationConfiguration = value
}
func (m *EducationSynchronizationProfile) SetLicensesToAssign(value []EducationSynchronizationLicenseAssignment)() {
    m.licensesToAssign = value
}
func (m *EducationSynchronizationProfile) SetProfileStatus(value *EducationSynchronizationProfileStatus)() {
    m.profileStatus = value
}
func (m *EducationSynchronizationProfile) SetState(value *EducationSynchronizationProfileState)() {
    m.state = value
}
