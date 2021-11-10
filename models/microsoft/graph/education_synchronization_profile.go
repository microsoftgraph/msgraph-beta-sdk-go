package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type EducationSynchronizationProfile struct {
    Entity
    // 
    dataProvider *EducationSynchronizationDataProvider;
    // Name of the configuration profile for syncing identities.
    displayName *string;
    // All errors associated with this synchronization profile.
    errors []EducationSynchronizationError;
    // The date the profile should be considered expired and cease syncing. When null. the profile will never expire. (optional)
    expirationDate *string;
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
// Instantiates a new educationSynchronizationProfile and sets the default values.
func NewEducationSynchronizationProfile()(*EducationSynchronizationProfile) {
    m := &EducationSynchronizationProfile{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the dataProvider property value. 
func (m *EducationSynchronizationProfile) GetDataProvider()(*EducationSynchronizationDataProvider) {
    if m == nil {
        return nil
    } else {
        return m.dataProvider
    }
}
// Gets the displayName property value. Name of the configuration profile for syncing identities.
func (m *EducationSynchronizationProfile) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the errors property value. All errors associated with this synchronization profile.
func (m *EducationSynchronizationProfile) GetErrors()([]EducationSynchronizationError) {
    if m == nil {
        return nil
    } else {
        return m.errors
    }
}
// Gets the expirationDate property value. The date the profile should be considered expired and cease syncing. When null. the profile will never expire. (optional)
func (m *EducationSynchronizationProfile) GetExpirationDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.expirationDate
    }
}
// Gets the handleSpecialCharacterConstraint property value. Determines if School Data Sync should automatically replace unsupported special characters while syncing from source.
func (m *EducationSynchronizationProfile) GetHandleSpecialCharacterConstraint()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.handleSpecialCharacterConstraint
    }
}
// Gets the identitySynchronizationConfiguration property value. 
func (m *EducationSynchronizationProfile) GetIdentitySynchronizationConfiguration()(*EducationIdentitySynchronizationConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.identitySynchronizationConfiguration
    }
}
// Gets the licensesToAssign property value. License setup configuration.
func (m *EducationSynchronizationProfile) GetLicensesToAssign()([]EducationSynchronizationLicenseAssignment) {
    if m == nil {
        return nil
    } else {
        return m.licensesToAssign
    }
}
// Gets the profileStatus property value. The synchronization status.
func (m *EducationSynchronizationProfile) GetProfileStatus()(*EducationSynchronizationProfileStatus) {
    if m == nil {
        return nil
    } else {
        return m.profileStatus
    }
}
// Gets the state property value. The state of the profile. Possible values are: provisioning, provisioned, provisioningFailed, deleting, deletionFailed.
func (m *EducationSynchronizationProfile) GetState()(*EducationSynchronizationProfileState) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// The deserialization information for the current model
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
        val, err := n.GetStringValue()
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
            cast := val.(EducationSynchronizationProfileState)
            m.SetState(&cast)
        }
        return nil
    }
    return res
}
func (m *EducationSynchronizationProfile) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the dataProvider property value. 
// Parameters:
//  - value : Value to set for the dataProvider property.
func (m *EducationSynchronizationProfile) SetDataProvider(value *EducationSynchronizationDataProvider)() {
    m.dataProvider = value
}
// Sets the displayName property value. Name of the configuration profile for syncing identities.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *EducationSynchronizationProfile) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the errors property value. All errors associated with this synchronization profile.
// Parameters:
//  - value : Value to set for the errors property.
func (m *EducationSynchronizationProfile) SetErrors(value []EducationSynchronizationError)() {
    m.errors = value
}
// Sets the expirationDate property value. The date the profile should be considered expired and cease syncing. When null. the profile will never expire. (optional)
// Parameters:
//  - value : Value to set for the expirationDate property.
func (m *EducationSynchronizationProfile) SetExpirationDate(value *string)() {
    m.expirationDate = value
}
// Sets the handleSpecialCharacterConstraint property value. Determines if School Data Sync should automatically replace unsupported special characters while syncing from source.
// Parameters:
//  - value : Value to set for the handleSpecialCharacterConstraint property.
func (m *EducationSynchronizationProfile) SetHandleSpecialCharacterConstraint(value *bool)() {
    m.handleSpecialCharacterConstraint = value
}
// Sets the identitySynchronizationConfiguration property value. 
// Parameters:
//  - value : Value to set for the identitySynchronizationConfiguration property.
func (m *EducationSynchronizationProfile) SetIdentitySynchronizationConfiguration(value *EducationIdentitySynchronizationConfiguration)() {
    m.identitySynchronizationConfiguration = value
}
// Sets the licensesToAssign property value. License setup configuration.
// Parameters:
//  - value : Value to set for the licensesToAssign property.
func (m *EducationSynchronizationProfile) SetLicensesToAssign(value []EducationSynchronizationLicenseAssignment)() {
    m.licensesToAssign = value
}
// Sets the profileStatus property value. The synchronization status.
// Parameters:
//  - value : Value to set for the profileStatus property.
func (m *EducationSynchronizationProfile) SetProfileStatus(value *EducationSynchronizationProfileStatus)() {
    m.profileStatus = value
}
// Sets the state property value. The state of the profile. Possible values are: provisioning, provisioned, provisioningFailed, deleting, deletionFailed.
// Parameters:
//  - value : Value to set for the state property.
func (m *EducationSynchronizationProfile) SetState(value *EducationSynchronizationProfileState)() {
    m.state = value
}
