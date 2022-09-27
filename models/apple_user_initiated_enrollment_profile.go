package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AppleUserInitiatedEnrollmentProfile the enrollmentProfile resource represents a collection of configurations which must be provided pre-enrollment to enable enrolling certain devices whose identities have been pre-staged. Pre-staged device identities are assigned to this type of profile to apply the profile's configurations at enrollment of the corresponding device.
type AppleUserInitiatedEnrollmentProfile struct {
    Entity
    // The list of assignments for this profile.
    assignments []AppleEnrollmentProfileAssignmentable
    // List of available enrollment type options
    availableEnrollmentTypeOptions []AppleOwnerTypeEnrollmentTypeable
    // Profile creation time
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The defaultEnrollmentType property
    defaultEnrollmentType *AppleUserInitiatedEnrollmentType
    // Description of the profile
    description *string
    // Name of the profile
    displayName *string
    // Profile last modified time
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Supported platform types.
    platform *DevicePlatformType
    // Priority, 0 is highest
    priority *int32
}
// NewAppleUserInitiatedEnrollmentProfile instantiates a new appleUserInitiatedEnrollmentProfile and sets the default values.
func NewAppleUserInitiatedEnrollmentProfile()(*AppleUserInitiatedEnrollmentProfile) {
    m := &AppleUserInitiatedEnrollmentProfile{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.appleUserInitiatedEnrollmentProfile";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAppleUserInitiatedEnrollmentProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAppleUserInitiatedEnrollmentProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAppleUserInitiatedEnrollmentProfile(), nil
}
// GetAssignments gets the assignments property value. The list of assignments for this profile.
func (m *AppleUserInitiatedEnrollmentProfile) GetAssignments()([]AppleEnrollmentProfileAssignmentable) {
    return m.assignments
}
// GetAvailableEnrollmentTypeOptions gets the availableEnrollmentTypeOptions property value. List of available enrollment type options
func (m *AppleUserInitiatedEnrollmentProfile) GetAvailableEnrollmentTypeOptions()([]AppleOwnerTypeEnrollmentTypeable) {
    return m.availableEnrollmentTypeOptions
}
// GetCreatedDateTime gets the createdDateTime property value. Profile creation time
func (m *AppleUserInitiatedEnrollmentProfile) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDefaultEnrollmentType gets the defaultEnrollmentType property value. The defaultEnrollmentType property
func (m *AppleUserInitiatedEnrollmentProfile) GetDefaultEnrollmentType()(*AppleUserInitiatedEnrollmentType) {
    return m.defaultEnrollmentType
}
// GetDescription gets the description property value. Description of the profile
func (m *AppleUserInitiatedEnrollmentProfile) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. Name of the profile
func (m *AppleUserInitiatedEnrollmentProfile) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AppleUserInitiatedEnrollmentProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAppleEnrollmentProfileAssignmentFromDiscriminatorValue , m.SetAssignments)
    res["availableEnrollmentTypeOptions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAppleOwnerTypeEnrollmentTypeFromDiscriminatorValue , m.SetAvailableEnrollmentTypeOptions)
    res["createdDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetCreatedDateTime)
    res["defaultEnrollmentType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseAppleUserInitiatedEnrollmentType , m.SetDefaultEnrollmentType)
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["lastModifiedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastModifiedDateTime)
    res["platform"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseDevicePlatformType , m.SetPlatform)
    res["priority"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetPriority)
    return res
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Profile last modified time
func (m *AppleUserInitiatedEnrollmentProfile) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetPlatform gets the platform property value. Supported platform types.
func (m *AppleUserInitiatedEnrollmentProfile) GetPlatform()(*DevicePlatformType) {
    return m.platform
}
// GetPriority gets the priority property value. Priority, 0 is highest
func (m *AppleUserInitiatedEnrollmentProfile) GetPriority()(*int32) {
    return m.priority
}
// Serialize serializes information the current object
func (m *AppleUserInitiatedEnrollmentProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssignments() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAssignments())
        err = writer.WriteCollectionOfObjectValues("assignments", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAvailableEnrollmentTypeOptions() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAvailableEnrollmentTypeOptions())
        err = writer.WriteCollectionOfObjectValues("availableEnrollmentTypeOptions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetDefaultEnrollmentType() != nil {
        cast := (*m.GetDefaultEnrollmentType()).String()
        err = writer.WriteStringValue("defaultEnrollmentType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
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
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetPlatform() != nil {
        cast := (*m.GetPlatform()).String()
        err = writer.WriteStringValue("platform", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("priority", m.GetPriority())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignments sets the assignments property value. The list of assignments for this profile.
func (m *AppleUserInitiatedEnrollmentProfile) SetAssignments(value []AppleEnrollmentProfileAssignmentable)() {
    m.assignments = value
}
// SetAvailableEnrollmentTypeOptions sets the availableEnrollmentTypeOptions property value. List of available enrollment type options
func (m *AppleUserInitiatedEnrollmentProfile) SetAvailableEnrollmentTypeOptions(value []AppleOwnerTypeEnrollmentTypeable)() {
    m.availableEnrollmentTypeOptions = value
}
// SetCreatedDateTime sets the createdDateTime property value. Profile creation time
func (m *AppleUserInitiatedEnrollmentProfile) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDefaultEnrollmentType sets the defaultEnrollmentType property value. The defaultEnrollmentType property
func (m *AppleUserInitiatedEnrollmentProfile) SetDefaultEnrollmentType(value *AppleUserInitiatedEnrollmentType)() {
    m.defaultEnrollmentType = value
}
// SetDescription sets the description property value. Description of the profile
func (m *AppleUserInitiatedEnrollmentProfile) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. Name of the profile
func (m *AppleUserInitiatedEnrollmentProfile) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Profile last modified time
func (m *AppleUserInitiatedEnrollmentProfile) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetPlatform sets the platform property value. Supported platform types.
func (m *AppleUserInitiatedEnrollmentProfile) SetPlatform(value *DevicePlatformType)() {
    m.platform = value
}
// SetPriority sets the priority property value. Priority, 0 is highest
func (m *AppleUserInitiatedEnrollmentProfile) SetPriority(value *int32)() {
    m.priority = value
}
