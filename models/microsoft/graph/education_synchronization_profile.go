package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// EducationSynchronizationProfile provides operations to manage the educationRoot singleton.
type EducationSynchronizationProfile struct {
    Entity
    // 
    dataProvider EducationSynchronizationDataProviderable;
    // Name of the configuration profile for syncing identities.
    displayName *string;
    // All errors associated with this synchronization profile.
    errors []EducationSynchronizationErrorable;
    // The date the profile should be considered expired and cease syncing. Provide the date in YYYY-MM-DD format, following ISO 8601. Maximum value is 18 months from profile creation.  (optional)
    expirationDate *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly;
    // Determines if School Data Sync should automatically replace unsupported special characters while syncing from source.
    handleSpecialCharacterConstraint *bool;
    // 
    identitySynchronizationConfiguration EducationIdentitySynchronizationConfigurationable;
    // License setup configuration.
    licensesToAssign []EducationSynchronizationLicenseAssignmentable;
    // The synchronization status.
    profileStatus EducationSynchronizationProfileStatusable;
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
// CreateEducationSynchronizationProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationSynchronizationProfileFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewEducationSynchronizationProfile(), nil
}
// GetDataProvider gets the dataProvider property value. 
func (m *EducationSynchronizationProfile) GetDataProvider()(EducationSynchronizationDataProviderable) {
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
func (m *EducationSynchronizationProfile) GetErrors()([]EducationSynchronizationErrorable) {
    if m == nil {
        return nil
    } else {
        return m.errors
    }
}
// GetExpirationDate gets the expirationDate property value. The date the profile should be considered expired and cease syncing. Provide the date in YYYY-MM-DD format, following ISO 8601. Maximum value is 18 months from profile creation.  (optional)
func (m *EducationSynchronizationProfile) GetExpirationDate()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly) {
    if m == nil {
        return nil
    } else {
        return m.expirationDate
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationSynchronizationProfile) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["dataProvider"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateEducationSynchronizationDataProviderFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataProvider(val.(EducationSynchronizationDataProviderable))
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
        val, err := n.GetCollectionOfObjectValues(CreateEducationSynchronizationErrorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EducationSynchronizationErrorable, len(val))
            for i, v := range val {
                res[i] = v.(EducationSynchronizationErrorable)
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
        val, err := n.GetObjectValue(CreateEducationIdentitySynchronizationConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentitySynchronizationConfiguration(val.(EducationIdentitySynchronizationConfigurationable))
        }
        return nil
    }
    res["licensesToAssign"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEducationSynchronizationLicenseAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EducationSynchronizationLicenseAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(EducationSynchronizationLicenseAssignmentable)
            }
            m.SetLicensesToAssign(res)
        }
        return nil
    }
    res["profileStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateEducationSynchronizationProfileStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProfileStatus(val.(EducationSynchronizationProfileStatusable))
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
// GetHandleSpecialCharacterConstraint gets the handleSpecialCharacterConstraint property value. Determines if School Data Sync should automatically replace unsupported special characters while syncing from source.
func (m *EducationSynchronizationProfile) GetHandleSpecialCharacterConstraint()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.handleSpecialCharacterConstraint
    }
}
// GetIdentitySynchronizationConfiguration gets the identitySynchronizationConfiguration property value. 
func (m *EducationSynchronizationProfile) GetIdentitySynchronizationConfiguration()(EducationIdentitySynchronizationConfigurationable) {
    if m == nil {
        return nil
    } else {
        return m.identitySynchronizationConfiguration
    }
}
// GetLicensesToAssign gets the licensesToAssign property value. License setup configuration.
func (m *EducationSynchronizationProfile) GetLicensesToAssign()([]EducationSynchronizationLicenseAssignmentable) {
    if m == nil {
        return nil
    } else {
        return m.licensesToAssign
    }
}
// GetProfileStatus gets the profileStatus property value. The synchronization status.
func (m *EducationSynchronizationProfile) GetProfileStatus()(EducationSynchronizationProfileStatusable) {
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
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
func (m *EducationSynchronizationProfile) SetDataProvider(value EducationSynchronizationDataProviderable)() {
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
func (m *EducationSynchronizationProfile) SetErrors(value []EducationSynchronizationErrorable)() {
    if m != nil {
        m.errors = value
    }
}
// SetExpirationDate sets the expirationDate property value. The date the profile should be considered expired and cease syncing. Provide the date in YYYY-MM-DD format, following ISO 8601. Maximum value is 18 months from profile creation.  (optional)
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
func (m *EducationSynchronizationProfile) SetIdentitySynchronizationConfiguration(value EducationIdentitySynchronizationConfigurationable)() {
    if m != nil {
        m.identitySynchronizationConfiguration = value
    }
}
// SetLicensesToAssign sets the licensesToAssign property value. License setup configuration.
func (m *EducationSynchronizationProfile) SetLicensesToAssign(value []EducationSynchronizationLicenseAssignmentable)() {
    if m != nil {
        m.licensesToAssign = value
    }
}
// SetProfileStatus sets the profileStatus property value. The synchronization status.
func (m *EducationSynchronizationProfile) SetProfileStatus(value EducationSynchronizationProfileStatusable)() {
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
