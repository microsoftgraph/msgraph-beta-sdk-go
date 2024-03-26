package windowsupdates

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

type Edition struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewEdition instantiates a new Edition and sets the default values.
func NewEdition()(*Edition) {
    m := &Edition{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateEditionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEditionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEdition(), nil
}
// GetDeviceFamily gets the deviceFamily property value. The device family targeted by the edition.
// returns a *string when successful
func (m *Edition) GetDeviceFamily()(*string) {
    val, err := m.GetBackingStore().Get("deviceFamily")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetEndOfServiceDateTime gets the endOfServiceDateTime property value. The date and time when the edition reached the end of service. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
// returns a *Time when successful
func (m *Edition) GetEndOfServiceDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("endOfServiceDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Edition) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deviceFamily"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceFamily(val)
        }
        return nil
    }
    res["endOfServiceDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndOfServiceDateTime(val)
        }
        return nil
    }
    res["generalAvailabilityDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGeneralAvailabilityDateTime(val)
        }
        return nil
    }
    res["isInService"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsInService(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["releasedName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReleasedName(val)
        }
        return nil
    }
    res["servicingPeriods"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateServicingPeriodFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ServicingPeriodable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ServicingPeriodable)
                }
            }
            m.SetServicingPeriods(res)
        }
        return nil
    }
    return res
}
// GetGeneralAvailabilityDateTime gets the generalAvailabilityDateTime property value. The date and time when the edition became available to the general customers for the first time. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
// returns a *Time when successful
func (m *Edition) GetGeneralAvailabilityDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("generalAvailabilityDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetIsInService gets the isInService property value. Indicates whether the edition is in service or out of service.
// returns a *bool when successful
func (m *Edition) GetIsInService()(*bool) {
    val, err := m.GetBackingStore().Get("isInService")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetName gets the name property value. The name of the edition. Read-only.
// returns a *string when successful
func (m *Edition) GetName()(*string) {
    val, err := m.GetBackingStore().Get("name")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetReleasedName gets the releasedName property value. The public name of the edition. Read-only.
// returns a *string when successful
func (m *Edition) GetReleasedName()(*string) {
    val, err := m.GetBackingStore().Get("releasedName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetServicingPeriods gets the servicingPeriods property value. The servicingPeriods property
// returns a []ServicingPeriodable when successful
func (m *Edition) GetServicingPeriods()([]ServicingPeriodable) {
    val, err := m.GetBackingStore().Get("servicingPeriods")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ServicingPeriodable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *Edition) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("deviceFamily", m.GetDeviceFamily())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("endOfServiceDateTime", m.GetEndOfServiceDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("generalAvailabilityDateTime", m.GetGeneralAvailabilityDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isInService", m.GetIsInService())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("releasedName", m.GetReleasedName())
        if err != nil {
            return err
        }
    }
    if m.GetServicingPeriods() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetServicingPeriods()))
        for i, v := range m.GetServicingPeriods() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("servicingPeriods", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDeviceFamily sets the deviceFamily property value. The device family targeted by the edition.
func (m *Edition) SetDeviceFamily(value *string)() {
    err := m.GetBackingStore().Set("deviceFamily", value)
    if err != nil {
        panic(err)
    }
}
// SetEndOfServiceDateTime sets the endOfServiceDateTime property value. The date and time when the edition reached the end of service. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *Edition) SetEndOfServiceDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("endOfServiceDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetGeneralAvailabilityDateTime sets the generalAvailabilityDateTime property value. The date and time when the edition became available to the general customers for the first time. The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *Edition) SetGeneralAvailabilityDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("generalAvailabilityDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetIsInService sets the isInService property value. Indicates whether the edition is in service or out of service.
func (m *Edition) SetIsInService(value *bool)() {
    err := m.GetBackingStore().Set("isInService", value)
    if err != nil {
        panic(err)
    }
}
// SetName sets the name property value. The name of the edition. Read-only.
func (m *Edition) SetName(value *string)() {
    err := m.GetBackingStore().Set("name", value)
    if err != nil {
        panic(err)
    }
}
// SetReleasedName sets the releasedName property value. The public name of the edition. Read-only.
func (m *Edition) SetReleasedName(value *string)() {
    err := m.GetBackingStore().Set("releasedName", value)
    if err != nil {
        panic(err)
    }
}
// SetServicingPeriods sets the servicingPeriods property value. The servicingPeriods property
func (m *Edition) SetServicingPeriods(value []ServicingPeriodable)() {
    err := m.GetBackingStore().Set("servicingPeriods", value)
    if err != nil {
        panic(err)
    }
}
type Editionable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeviceFamily()(*string)
    GetEndOfServiceDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetGeneralAvailabilityDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetIsInService()(*bool)
    GetName()(*string)
    GetReleasedName()(*string)
    GetServicingPeriods()([]ServicingPeriodable)
    SetDeviceFamily(value *string)()
    SetEndOfServiceDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetGeneralAvailabilityDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetIsInService(value *bool)()
    SetName(value *string)()
    SetReleasedName(value *string)()
    SetServicingPeriods(value []ServicingPeriodable)()
}
