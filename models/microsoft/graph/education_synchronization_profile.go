package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// EducationSynchronizationProfile 
type EducationSynchronizationProfile struct {
    Entity
    // 
    dataProvider *EducationSynchronizationDataProvider;
    // Name of the configuration profile for syncing identities.
    displayName *string;
    // All errors associated with this synchronization profile.
    errors []EducationSynchronizationError;
    // The date the profile should be considered expired and cease syncing. When null. the profile will never expire. (optional)
    expirationDate *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly;
    // Determines if School Data Sync should automatically replace unsupported special characters while syncing from source.
    handleSpecialCharacterConstraint *bool;
    // 
    identitySynchronizationConfiguration *EducationIdentitySynchronizationConfiguration;
    // License setup configuration.
    licensesToAssign []EducationSynchronizationLicenseAssignment;
    // The synchronization status.
    profileStatus *EducationSynchronizationProfileStatus;
    // The state of the profile. Possible values are: provisioning, provisioned, provisioningFailed, deleting, deletionFailed.
    state *EducationSynchronizationProfileState;
}
// NewEducationSynchronizationProfile instantiates a new educationSynchronizationProfile and sets the default values.
func NewEducationSynchronizationProfile()(*EducationSynchronizationProfile) {
    m := &EducationSynchronizationProfile{
        Entity: *NewEntity(),
    }
    return m
}
// GetDataProvider gets the dataProvider property value. 
func (m *EducationSynchronizationProfile) GetDataProvider()(*EducationSynchronizationDataProvider) {
    if m == nil {
        return nil
    } else {
        return m.dataProvider
    }
}
// GetDisplayName gets the displayName property value. Name of the configuration profile for syncing identities.
func (m *EducationSynchronizationProfile) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetErrors gets the errors property value. All errors associated with this synchronization profile.
func (m *EducationSynchronizationProfile) GetErrors()([]EducationSynchronizationError) {
    if m == nil {
        return nil
    } else {
        return m.errors
    }
}
// GetExpirationDate gets the expirationDate property value. The date the profile should be considered expired and cease syncing. When null. the profile will never expire. (optional)
func (m *EducationSynchronizationProfile) GetExpirationDate()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly) {
    if m == nil {
        return nil
    } else {
        return m.expirationDate
    }
}
// GetHandleSpecialCharacterConstraint gets the handleSpecialCharacterConstraint property value. Determines if School Data Sync should automatically replace unsupported special characters while syncing from source.
func (m *EducationSynchronizationProfile) GetHandleSpecialCharacterConstraint()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.handleSpecialCharacterConstraint
    }
}
// GetIdentitySynchronizationConfiguration gets the identitySynchronizationConfiguration property value. 
func (m *EducationSynchronizationProfile) GetIdentitySynchronizationConfiguration()(*EducationIdentitySynchronizationConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.identitySynchronizationConfiguration
    }
}
// GetLicensesToAssign gets the licensesToAssign property value. License setup configuration.
func (m *EducationSynchronizationProfile) GetLicensesToAssign()([]EducationSynchronizationLicenseAssignment) {
    if m == nil {
        return nil
    } else {
        return m.licensesToAssign
    }
}
// GetProfileStatus gets the profileStatus property value. The synchronization status.
func (m *EducationSynchronizationProfile) GetProfileStatus()(*EducationSynchronizationProfileStatus) {
    if m == nil {
        return nil
    } else {
        return m.profileStatus
    }
}
// GetState gets the state property value. The state of the profile. Possible values are: provisioning, provisioned, provisioningFailed, deleting, deletionFailed.
func (m *EducationSynchronizationProfile) GetState()(*EducationSynchronizationProfileState) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationSynchronizationProfile) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["dataProvider"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationSynchronizationDataProvider() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataProvider(val.(*EducationSynchronizationDataProvider))
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["errors"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationSynchronizationError() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EducationSynchronizationError, len(val))
            for i, v := range val {
                res[i] = *(v.(*EducationSynchronizationError))
            }
            m.SetErrors(res)
        }
        return nil
    }
    res["expirationDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationDate(val)
        }
        return nil
    }
    res["handleSpecialCharacterConstraint"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHandleSpecialCharacterConstraint(val)
        }
        return nil
    }
    res["identitySynchronizationConfiguration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationIdentitySynchronizationConfiguration() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentitySynchronizationConfiguration(val.(*EducationIdentitySynchronizationConfiguration))
        }
        return nil
    }
    res["licensesToAssign"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationSynchronizationLicenseAssignment() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EducationSynchronizationLicenseAssignment, len(val))
            for i, v := range val {
                res[i] = *(v.(*EducationSynchronizationLicenseAssignment))
            }
            m.SetLicensesToAssign(res)
        }
        return nil
    }
    res["profileStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationSynchronizationProfileStatus() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProfileStatus(val.(*EducationSynchronizationProfileStatus))
        }
        return nil
    }
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseEducationSynchronizationProfileState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*EducationSynchronizationProfileState))
        }
        return nil
    }
    return res
}
func (m *EducationSynchronizationProfile) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
    if m.GetErrors() != nil {
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
        err = writer.WriteDateOnlyValue("expirationDate", m.GetExpirationDate())
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
    if m.GetLicensesToAssign() != nil {
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
        cast := (*m.GetState()).String()
        err = writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDataProvider sets the dataProvider property value. 
func (m *EducationSynchronizationProfile) SetDataProvider(value *EducationSynchronizationDataProvider)() {
    if m != nil {
        m.dataProvider = value
    }
}
// SetDisplayName sets the displayName property value. Name of the configuration profile for syncing identities.
func (m *EducationSynchronizationProfile) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetErrors sets the errors property value. All errors associated with this synchronization profile.
func (m *EducationSynchronizationProfile) SetErrors(value []EducationSynchronizationError)() {
    if m != nil {
        m.errors = value
    }
}
// SetExpirationDate sets the expirationDate property value. The date the profile should be considered expired and cease syncing. When null. the profile will never expire. (optional)
func (m *EducationSynchronizationProfile) SetExpirationDate(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)() {
    if m != nil {
        m.expirationDate = value
    }
}
// SetHandleSpecialCharacterConstraint sets the handleSpecialCharacterConstraint property value. Determines if School Data Sync should automatically replace unsupported special characters while syncing from source.
func (m *EducationSynchronizationProfile) SetHandleSpecialCharacterConstraint(value *bool)() {
    if m != nil {
        m.handleSpecialCharacterConstraint = value
    }
}
// SetIdentitySynchronizationConfiguration sets the identitySynchronizationConfiguration property value. 
func (m *EducationSynchronizationProfile) SetIdentitySynchronizationConfiguration(value *EducationIdentitySynchronizationConfiguration)() {
    if m != nil {
        m.identitySynchronizationConfiguration = value
    }
}
// SetLicensesToAssign sets the licensesToAssign property value. License setup configuration.
func (m *EducationSynchronizationProfile) SetLicensesToAssign(value []EducationSynchronizationLicenseAssignment)() {
    if m != nil {
        m.licensesToAssign = value
    }
}
// SetProfileStatus sets the profileStatus property value. The synchronization status.
func (m *EducationSynchronizationProfile) SetProfileStatus(value *EducationSynchronizationProfileStatus)() {
    if m != nil {
        m.profileStatus = value
    }
}
// SetState sets the state property value. The state of the profile. Possible values are: provisioning, provisioned, provisioningFailed, deleting, deletionFailed.
func (m *EducationSynchronizationProfile) SetState(value *EducationSynchronizationProfileState)() {
    if m != nil {
        m.state = value
    }
}
