package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EmbeddedSIMActivationCodePool a pool represents a group of embedded SIM activation codes.
type EmbeddedSIMActivationCodePool struct {
    Entity
    // The total count of activation codes which belong to this pool.
    activationCodeCount *int32
    // The activation codes which belong to this pool. This navigation property is used to post activation codes to Intune but cannot be used to read activation codes from Intune.
    activationCodes []EmbeddedSIMActivationCodeable
    // Navigational property to a list of targets to which this pool is assigned.
    assignments []EmbeddedSIMActivationCodePoolAssignmentable
    // The time the embedded SIM activation code pool was created. Generated service side.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Navigational property to a list of device states for this pool.
    deviceStates []EmbeddedSIMDeviceStateable
    // The admin defined name of the embedded SIM activation code pool.
    displayName *string
    // The time the embedded SIM activation code pool was last modified. Updated service side.
    modifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewEmbeddedSIMActivationCodePool instantiates a new embeddedSIMActivationCodePool and sets the default values.
func NewEmbeddedSIMActivationCodePool()(*EmbeddedSIMActivationCodePool) {
    m := &EmbeddedSIMActivationCodePool{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.embeddedSIMActivationCodePool";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateEmbeddedSIMActivationCodePoolFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEmbeddedSIMActivationCodePoolFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEmbeddedSIMActivationCodePool(), nil
}
// GetActivationCodeCount gets the activationCodeCount property value. The total count of activation codes which belong to this pool.
func (m *EmbeddedSIMActivationCodePool) GetActivationCodeCount()(*int32) {
    return m.activationCodeCount
}
// GetActivationCodes gets the activationCodes property value. The activation codes which belong to this pool. This navigation property is used to post activation codes to Intune but cannot be used to read activation codes from Intune.
func (m *EmbeddedSIMActivationCodePool) GetActivationCodes()([]EmbeddedSIMActivationCodeable) {
    return m.activationCodes
}
// GetAssignments gets the assignments property value. Navigational property to a list of targets to which this pool is assigned.
func (m *EmbeddedSIMActivationCodePool) GetAssignments()([]EmbeddedSIMActivationCodePoolAssignmentable) {
    return m.assignments
}
// GetCreatedDateTime gets the createdDateTime property value. The time the embedded SIM activation code pool was created. Generated service side.
func (m *EmbeddedSIMActivationCodePool) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDeviceStates gets the deviceStates property value. Navigational property to a list of device states for this pool.
func (m *EmbeddedSIMActivationCodePool) GetDeviceStates()([]EmbeddedSIMDeviceStateable) {
    return m.deviceStates
}
// GetDisplayName gets the displayName property value. The admin defined name of the embedded SIM activation code pool.
func (m *EmbeddedSIMActivationCodePool) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EmbeddedSIMActivationCodePool) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activationCodeCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetActivationCodeCount)
    res["activationCodes"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateEmbeddedSIMActivationCodeFromDiscriminatorValue , m.SetActivationCodes)
    res["assignments"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateEmbeddedSIMActivationCodePoolAssignmentFromDiscriminatorValue , m.SetAssignments)
    res["createdDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetCreatedDateTime)
    res["deviceStates"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateEmbeddedSIMDeviceStateFromDiscriminatorValue , m.SetDeviceStates)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["modifiedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetModifiedDateTime)
    return res
}
// GetModifiedDateTime gets the modifiedDateTime property value. The time the embedded SIM activation code pool was last modified. Updated service side.
func (m *EmbeddedSIMActivationCodePool) GetModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.modifiedDateTime
}
// Serialize serializes information the current object
func (m *EmbeddedSIMActivationCodePool) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("activationCodeCount", m.GetActivationCodeCount())
        if err != nil {
            return err
        }
    }
    if m.GetActivationCodes() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetActivationCodes())
        err = writer.WriteCollectionOfObjectValues("activationCodes", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAssignments() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAssignments())
        err = writer.WriteCollectionOfObjectValues("assignments", cast)
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
    if m.GetDeviceStates() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDeviceStates())
        err = writer.WriteCollectionOfObjectValues("deviceStates", cast)
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
        err = writer.WriteTimeValue("modifiedDateTime", m.GetModifiedDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActivationCodeCount sets the activationCodeCount property value. The total count of activation codes which belong to this pool.
func (m *EmbeddedSIMActivationCodePool) SetActivationCodeCount(value *int32)() {
    m.activationCodeCount = value
}
// SetActivationCodes sets the activationCodes property value. The activation codes which belong to this pool. This navigation property is used to post activation codes to Intune but cannot be used to read activation codes from Intune.
func (m *EmbeddedSIMActivationCodePool) SetActivationCodes(value []EmbeddedSIMActivationCodeable)() {
    m.activationCodes = value
}
// SetAssignments sets the assignments property value. Navigational property to a list of targets to which this pool is assigned.
func (m *EmbeddedSIMActivationCodePool) SetAssignments(value []EmbeddedSIMActivationCodePoolAssignmentable)() {
    m.assignments = value
}
// SetCreatedDateTime sets the createdDateTime property value. The time the embedded SIM activation code pool was created. Generated service side.
func (m *EmbeddedSIMActivationCodePool) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDeviceStates sets the deviceStates property value. Navigational property to a list of device states for this pool.
func (m *EmbeddedSIMActivationCodePool) SetDeviceStates(value []EmbeddedSIMDeviceStateable)() {
    m.deviceStates = value
}
// SetDisplayName sets the displayName property value. The admin defined name of the embedded SIM activation code pool.
func (m *EmbeddedSIMActivationCodePool) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetModifiedDateTime sets the modifiedDateTime property value. The time the embedded SIM activation code pool was last modified. Updated service side.
func (m *EmbeddedSIMActivationCodePool) SetModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.modifiedDateTime = value
}
