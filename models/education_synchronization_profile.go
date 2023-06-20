package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationSynchronizationProfile 
type EducationSynchronizationProfile struct {
    Entity
}
// NewEducationSynchronizationProfile instantiates a new EducationSynchronizationProfile and sets the default values.
func NewEducationSynchronizationProfile()(*EducationSynchronizationProfile) {
    m := &EducationSynchronizationProfile{
        Entity: *NewEntity(),
    }
    return m
}
// CreateEducationSynchronizationProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationSynchronizationProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationSynchronizationProfile(), nil
}
// GetDataProvider gets the dataProvider property value. The dataProvider property
func (m *EducationSynchronizationProfile) GetDataProvider()(EducationSynchronizationDataProviderable) {
    val, err := m.GetBackingStore().Get("dataProvider")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(EducationSynchronizationDataProviderable)
    }
    return nil
}
// GetDisplayName gets the displayName property value. Name of the configuration profile for syncing identities.
func (m *EducationSynchronizationProfile) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetErrors gets the errors property value. All errors associated with this synchronization profile.
func (m *EducationSynchronizationProfile) GetErrors()([]EducationSynchronizationErrorable) {
    val, err := m.GetBackingStore().Get("errors")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]EducationSynchronizationErrorable)
    }
    return nil
}
// GetExpirationDate gets the expirationDate property value. The date the profile should be considered expired and cease syncing. Provide the date in YYYY-MM-DD format, following ISO 8601. Maximum value is 18 months from profile creation.  (optional)
func (m *EducationSynchronizationProfile) GetExpirationDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    val, err := m.GetBackingStore().Get("expirationDate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationSynchronizationProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["dataProvider"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEducationSynchronizationDataProviderFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataProvider(val.(EducationSynchronizationDataProviderable))
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["errors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEducationSynchronizationErrorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EducationSynchronizationErrorable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(EducationSynchronizationErrorable)
                }
            }
            m.SetErrors(res)
        }
        return nil
    }
    res["expirationDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationDate(val)
        }
        return nil
    }
    res["handleSpecialCharacterConstraint"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHandleSpecialCharacterConstraint(val)
        }
        return nil
    }
    res["identitySynchronizationConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEducationIdentitySynchronizationConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentitySynchronizationConfiguration(val.(EducationIdentitySynchronizationConfigurationable))
        }
        return nil
    }
    res["licensesToAssign"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEducationSynchronizationLicenseAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EducationSynchronizationLicenseAssignmentable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(EducationSynchronizationLicenseAssignmentable)
                }
            }
            m.SetLicensesToAssign(res)
        }
        return nil
    }
    res["profileStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEducationSynchronizationProfileStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProfileStatus(val.(EducationSynchronizationProfileStatusable))
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
    val, err := m.GetBackingStore().Get("handleSpecialCharacterConstraint")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIdentitySynchronizationConfiguration gets the identitySynchronizationConfiguration property value. The identitySynchronizationConfiguration property
func (m *EducationSynchronizationProfile) GetIdentitySynchronizationConfiguration()(EducationIdentitySynchronizationConfigurationable) {
    val, err := m.GetBackingStore().Get("identitySynchronizationConfiguration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(EducationIdentitySynchronizationConfigurationable)
    }
    return nil
}
// GetLicensesToAssign gets the licensesToAssign property value. License setup configuration.
func (m *EducationSynchronizationProfile) GetLicensesToAssign()([]EducationSynchronizationLicenseAssignmentable) {
    val, err := m.GetBackingStore().Get("licensesToAssign")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]EducationSynchronizationLicenseAssignmentable)
    }
    return nil
}
// GetProfileStatus gets the profileStatus property value. The synchronization status.
func (m *EducationSynchronizationProfile) GetProfileStatus()(EducationSynchronizationProfileStatusable) {
    val, err := m.GetBackingStore().Get("profileStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(EducationSynchronizationProfileStatusable)
    }
    return nil
}
// GetState gets the state property value. The state of the profile. Possible values are: provisioning, provisioned, provisioningFailed, deleting, deletionFailed.
func (m *EducationSynchronizationProfile) GetState()(*EducationSynchronizationProfileState) {
    val, err := m.GetBackingStore().Get("state")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*EducationSynchronizationProfileState)
    }
    return nil
}
// Serialize serializes information the current object
func (m *EducationSynchronizationProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetErrors()))
        for i, v := range m.GetErrors() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLicensesToAssign()))
        for i, v := range m.GetLicensesToAssign() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
// SetDataProvider sets the dataProvider property value. The dataProvider property
func (m *EducationSynchronizationProfile) SetDataProvider(value EducationSynchronizationDataProviderable)() {
    err := m.GetBackingStore().Set("dataProvider", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. Name of the configuration profile for syncing identities.
func (m *EducationSynchronizationProfile) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetErrors sets the errors property value. All errors associated with this synchronization profile.
func (m *EducationSynchronizationProfile) SetErrors(value []EducationSynchronizationErrorable)() {
    err := m.GetBackingStore().Set("errors", value)
    if err != nil {
        panic(err)
    }
}
// SetExpirationDate sets the expirationDate property value. The date the profile should be considered expired and cease syncing. Provide the date in YYYY-MM-DD format, following ISO 8601. Maximum value is 18 months from profile creation.  (optional)
func (m *EducationSynchronizationProfile) SetExpirationDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    err := m.GetBackingStore().Set("expirationDate", value)
    if err != nil {
        panic(err)
    }
}
// SetHandleSpecialCharacterConstraint sets the handleSpecialCharacterConstraint property value. Determines if School Data Sync should automatically replace unsupported special characters while syncing from source.
func (m *EducationSynchronizationProfile) SetHandleSpecialCharacterConstraint(value *bool)() {
    err := m.GetBackingStore().Set("handleSpecialCharacterConstraint", value)
    if err != nil {
        panic(err)
    }
}
// SetIdentitySynchronizationConfiguration sets the identitySynchronizationConfiguration property value. The identitySynchronizationConfiguration property
func (m *EducationSynchronizationProfile) SetIdentitySynchronizationConfiguration(value EducationIdentitySynchronizationConfigurationable)() {
    err := m.GetBackingStore().Set("identitySynchronizationConfiguration", value)
    if err != nil {
        panic(err)
    }
}
// SetLicensesToAssign sets the licensesToAssign property value. License setup configuration.
func (m *EducationSynchronizationProfile) SetLicensesToAssign(value []EducationSynchronizationLicenseAssignmentable)() {
    err := m.GetBackingStore().Set("licensesToAssign", value)
    if err != nil {
        panic(err)
    }
}
// SetProfileStatus sets the profileStatus property value. The synchronization status.
func (m *EducationSynchronizationProfile) SetProfileStatus(value EducationSynchronizationProfileStatusable)() {
    err := m.GetBackingStore().Set("profileStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetState sets the state property value. The state of the profile. Possible values are: provisioning, provisioned, provisioningFailed, deleting, deletionFailed.
func (m *EducationSynchronizationProfile) SetState(value *EducationSynchronizationProfileState)() {
    err := m.GetBackingStore().Set("state", value)
    if err != nil {
        panic(err)
    }
}
// EducationSynchronizationProfileable 
type EducationSynchronizationProfileable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDataProvider()(EducationSynchronizationDataProviderable)
    GetDisplayName()(*string)
    GetErrors()([]EducationSynchronizationErrorable)
    GetExpirationDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetHandleSpecialCharacterConstraint()(*bool)
    GetIdentitySynchronizationConfiguration()(EducationIdentitySynchronizationConfigurationable)
    GetLicensesToAssign()([]EducationSynchronizationLicenseAssignmentable)
    GetProfileStatus()(EducationSynchronizationProfileStatusable)
    GetState()(*EducationSynchronizationProfileState)
    SetDataProvider(value EducationSynchronizationDataProviderable)()
    SetDisplayName(value *string)()
    SetErrors(value []EducationSynchronizationErrorable)()
    SetExpirationDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetHandleSpecialCharacterConstraint(value *bool)()
    SetIdentitySynchronizationConfiguration(value EducationIdentitySynchronizationConfigurationable)()
    SetLicensesToAssign(value []EducationSynchronizationLicenseAssignmentable)()
    SetProfileStatus(value EducationSynchronizationProfileStatusable)()
    SetState(value *EducationSynchronizationProfileState)()
}
