package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EmbeddedSIMActivationCodePool 
type EmbeddedSIMActivationCodePool struct {
    Entity
}
// NewEmbeddedSIMActivationCodePool instantiates a new EmbeddedSIMActivationCodePool and sets the default values.
func NewEmbeddedSIMActivationCodePool()(*EmbeddedSIMActivationCodePool) {
    m := &EmbeddedSIMActivationCodePool{
        Entity: *NewEntity(),
    }
    return m
}
// CreateEmbeddedSIMActivationCodePoolFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEmbeddedSIMActivationCodePoolFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEmbeddedSIMActivationCodePool(), nil
}
// GetActivationCodeCount gets the activationCodeCount property value. The total count of activation codes which belong to this pool.
func (m *EmbeddedSIMActivationCodePool) GetActivationCodeCount()(*int32) {
    val, err := m.GetBackingStore().Get("activationCodeCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetActivationCodes gets the activationCodes property value. The activation codes which belong to this pool. This navigation property is used to post activation codes to Intune but cannot be used to read activation codes from Intune.
func (m *EmbeddedSIMActivationCodePool) GetActivationCodes()([]EmbeddedSIMActivationCodeable) {
    val, err := m.GetBackingStore().Get("activationCodes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]EmbeddedSIMActivationCodeable)
    }
    return nil
}
// GetAssignments gets the assignments property value. Navigational property to a list of targets to which this pool is assigned.
func (m *EmbeddedSIMActivationCodePool) GetAssignments()([]EmbeddedSIMActivationCodePoolAssignmentable) {
    val, err := m.GetBackingStore().Get("assignments")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]EmbeddedSIMActivationCodePoolAssignmentable)
    }
    return nil
}
// GetCreatedDateTime gets the createdDateTime property value. The time the embedded SIM activation code pool was created. Generated service side.
func (m *EmbeddedSIMActivationCodePool) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetDeviceStates gets the deviceStates property value. Navigational property to a list of device states for this pool.
func (m *EmbeddedSIMActivationCodePool) GetDeviceStates()([]EmbeddedSIMDeviceStateable) {
    val, err := m.GetBackingStore().Get("deviceStates")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]EmbeddedSIMDeviceStateable)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The admin defined name of the embedded SIM activation code pool.
func (m *EmbeddedSIMActivationCodePool) GetDisplayName()(*string) {
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
func (m *EmbeddedSIMActivationCodePool) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activationCodeCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivationCodeCount(val)
        }
        return nil
    }
    res["activationCodes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEmbeddedSIMActivationCodeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EmbeddedSIMActivationCodeable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(EmbeddedSIMActivationCodeable)
                }
            }
            m.SetActivationCodes(res)
        }
        return nil
    }
    res["assignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEmbeddedSIMActivationCodePoolAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EmbeddedSIMActivationCodePoolAssignmentable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(EmbeddedSIMActivationCodePoolAssignmentable)
                }
            }
            m.SetAssignments(res)
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
    res["deviceStates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEmbeddedSIMDeviceStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EmbeddedSIMDeviceStateable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(EmbeddedSIMDeviceStateable)
                }
            }
            m.SetDeviceStates(res)
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
    res["modifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModifiedDateTime(val)
        }
        return nil
    }
    return res
}
// GetModifiedDateTime gets the modifiedDateTime property value. The time the embedded SIM activation code pool was last modified. Updated service side.
func (m *EmbeddedSIMActivationCodePool) GetModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("modifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetActivationCodes()))
        for i, v := range m.GetActivationCodes() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("activationCodes", cast)
        if err != nil {
            return err
        }
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
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceStates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceStates()))
        for i, v := range m.GetDeviceStates() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
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
    err := m.GetBackingStore().Set("activationCodeCount", value)
    if err != nil {
        panic(err)
    }
}
// SetActivationCodes sets the activationCodes property value. The activation codes which belong to this pool. This navigation property is used to post activation codes to Intune but cannot be used to read activation codes from Intune.
func (m *EmbeddedSIMActivationCodePool) SetActivationCodes(value []EmbeddedSIMActivationCodeable)() {
    err := m.GetBackingStore().Set("activationCodes", value)
    if err != nil {
        panic(err)
    }
}
// SetAssignments sets the assignments property value. Navigational property to a list of targets to which this pool is assigned.
func (m *EmbeddedSIMActivationCodePool) SetAssignments(value []EmbeddedSIMActivationCodePoolAssignmentable)() {
    err := m.GetBackingStore().Set("assignments", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The time the embedded SIM activation code pool was created. Generated service side.
func (m *EmbeddedSIMActivationCodePool) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceStates sets the deviceStates property value. Navigational property to a list of device states for this pool.
func (m *EmbeddedSIMActivationCodePool) SetDeviceStates(value []EmbeddedSIMDeviceStateable)() {
    err := m.GetBackingStore().Set("deviceStates", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The admin defined name of the embedded SIM activation code pool.
func (m *EmbeddedSIMActivationCodePool) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetModifiedDateTime sets the modifiedDateTime property value. The time the embedded SIM activation code pool was last modified. Updated service side.
func (m *EmbeddedSIMActivationCodePool) SetModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("modifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// EmbeddedSIMActivationCodePoolable 
type EmbeddedSIMActivationCodePoolable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActivationCodeCount()(*int32)
    GetActivationCodes()([]EmbeddedSIMActivationCodeable)
    GetAssignments()([]EmbeddedSIMActivationCodePoolAssignmentable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDeviceStates()([]EmbeddedSIMDeviceStateable)
    GetDisplayName()(*string)
    GetModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetActivationCodeCount(value *int32)()
    SetActivationCodes(value []EmbeddedSIMActivationCodeable)()
    SetAssignments(value []EmbeddedSIMActivationCodePoolAssignmentable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDeviceStates(value []EmbeddedSIMDeviceStateable)()
    SetDisplayName(value *string)()
    SetModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
