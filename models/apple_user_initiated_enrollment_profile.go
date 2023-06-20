package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AppleUserInitiatedEnrollmentProfile the enrollmentProfile resource represents a collection of configurations which must be provided pre-enrollment to enable enrolling certain devices whose identities have been pre-staged. Pre-staged device identities are assigned to this type of profile to apply the profile's configurations at enrollment of the corresponding device.
type AppleUserInitiatedEnrollmentProfile struct {
    Entity
}
// NewAppleUserInitiatedEnrollmentProfile instantiates a new appleUserInitiatedEnrollmentProfile and sets the default values.
func NewAppleUserInitiatedEnrollmentProfile()(*AppleUserInitiatedEnrollmentProfile) {
    m := &AppleUserInitiatedEnrollmentProfile{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAppleUserInitiatedEnrollmentProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAppleUserInitiatedEnrollmentProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAppleUserInitiatedEnrollmentProfile(), nil
}
// GetAssignments gets the assignments property value. The list of assignments for this profile.
func (m *AppleUserInitiatedEnrollmentProfile) GetAssignments()([]AppleEnrollmentProfileAssignmentable) {
    val, err := m.GetBackingStore().Get("assignments")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AppleEnrollmentProfileAssignmentable)
    }
    return nil
}
// GetAvailableEnrollmentTypeOptions gets the availableEnrollmentTypeOptions property value. List of available enrollment type options
func (m *AppleUserInitiatedEnrollmentProfile) GetAvailableEnrollmentTypeOptions()([]AppleOwnerTypeEnrollmentTypeable) {
    val, err := m.GetBackingStore().Get("availableEnrollmentTypeOptions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AppleOwnerTypeEnrollmentTypeable)
    }
    return nil
}
// GetCreatedDateTime gets the createdDateTime property value. Profile creation time
func (m *AppleUserInitiatedEnrollmentProfile) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetDefaultEnrollmentType gets the defaultEnrollmentType property value. The defaultEnrollmentType property
func (m *AppleUserInitiatedEnrollmentProfile) GetDefaultEnrollmentType()(*AppleUserInitiatedEnrollmentType) {
    val, err := m.GetBackingStore().Get("defaultEnrollmentType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AppleUserInitiatedEnrollmentType)
    }
    return nil
}
// GetDescription gets the description property value. Description of the profile
func (m *AppleUserInitiatedEnrollmentProfile) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. Name of the profile
func (m *AppleUserInitiatedEnrollmentProfile) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AppleUserInitiatedEnrollmentProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAppleEnrollmentProfileAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AppleEnrollmentProfileAssignmentable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AppleEnrollmentProfileAssignmentable)
                }
            }
            m.SetAssignments(res)
        }
        return nil
    }
    res["availableEnrollmentTypeOptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAppleOwnerTypeEnrollmentTypeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AppleOwnerTypeEnrollmentTypeable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AppleOwnerTypeEnrollmentTypeable)
                }
            }
            m.SetAvailableEnrollmentTypeOptions(res)
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["defaultEnrollmentType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppleUserInitiatedEnrollmentType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultEnrollmentType(val.(*AppleUserInitiatedEnrollmentType))
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
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
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["platform"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDevicePlatformType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlatform(val.(*DevicePlatformType))
        }
        return nil
    }
    res["priority"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPriority(val)
        }
        return nil
    }
    return res
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Profile last modified time
func (m *AppleUserInitiatedEnrollmentProfile) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastModifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetPlatform gets the platform property value. Supported platform types.
func (m *AppleUserInitiatedEnrollmentProfile) GetPlatform()(*DevicePlatformType) {
    val, err := m.GetBackingStore().Get("platform")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DevicePlatformType)
    }
    return nil
}
// GetPriority gets the priority property value. Priority, 0 is highest
func (m *AppleUserInitiatedEnrollmentProfile) GetPriority()(*int32) {
    val, err := m.GetBackingStore().Get("priority")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AppleUserInitiatedEnrollmentProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignments()))
        for i, v := range m.GetAssignments() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("assignments", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAvailableEnrollmentTypeOptions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAvailableEnrollmentTypeOptions()))
        for i, v := range m.GetAvailableEnrollmentTypeOptions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
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
    err := m.GetBackingStore().Set("assignments", value)
    if err != nil {
        panic(err)
    }
}
// SetAvailableEnrollmentTypeOptions sets the availableEnrollmentTypeOptions property value. List of available enrollment type options
func (m *AppleUserInitiatedEnrollmentProfile) SetAvailableEnrollmentTypeOptions(value []AppleOwnerTypeEnrollmentTypeable)() {
    err := m.GetBackingStore().Set("availableEnrollmentTypeOptions", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. Profile creation time
func (m *AppleUserInitiatedEnrollmentProfile) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetDefaultEnrollmentType sets the defaultEnrollmentType property value. The defaultEnrollmentType property
func (m *AppleUserInitiatedEnrollmentProfile) SetDefaultEnrollmentType(value *AppleUserInitiatedEnrollmentType)() {
    err := m.GetBackingStore().Set("defaultEnrollmentType", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. Description of the profile
func (m *AppleUserInitiatedEnrollmentProfile) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. Name of the profile
func (m *AppleUserInitiatedEnrollmentProfile) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Profile last modified time
func (m *AppleUserInitiatedEnrollmentProfile) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetPlatform sets the platform property value. Supported platform types.
func (m *AppleUserInitiatedEnrollmentProfile) SetPlatform(value *DevicePlatformType)() {
    err := m.GetBackingStore().Set("platform", value)
    if err != nil {
        panic(err)
    }
}
// SetPriority sets the priority property value. Priority, 0 is highest
func (m *AppleUserInitiatedEnrollmentProfile) SetPriority(value *int32)() {
    err := m.GetBackingStore().Set("priority", value)
    if err != nil {
        panic(err)
    }
}
// AppleUserInitiatedEnrollmentProfileable 
type AppleUserInitiatedEnrollmentProfileable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignments()([]AppleEnrollmentProfileAssignmentable)
    GetAvailableEnrollmentTypeOptions()([]AppleOwnerTypeEnrollmentTypeable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDefaultEnrollmentType()(*AppleUserInitiatedEnrollmentType)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPlatform()(*DevicePlatformType)
    GetPriority()(*int32)
    SetAssignments(value []AppleEnrollmentProfileAssignmentable)()
    SetAvailableEnrollmentTypeOptions(value []AppleOwnerTypeEnrollmentTypeable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDefaultEnrollmentType(value *AppleUserInitiatedEnrollmentType)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPlatform(value *DevicePlatformType)()
    SetPriority(value *int32)()
}
