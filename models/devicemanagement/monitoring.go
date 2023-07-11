package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// Monitoring 
type Monitoring struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewMonitoring instantiates a new monitoring and sets the default values.
func NewMonitoring()(*Monitoring) {
    m := &Monitoring{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateMonitoringFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMonitoringFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMonitoring(), nil
}
// GetAlertRecords gets the alertRecords property value. The collection of records of alert events.
func (m *Monitoring) GetAlertRecords()([]AlertRecordable) {
    val, err := m.GetBackingStore().Get("alertRecords")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AlertRecordable)
    }
    return nil
}
// GetAlertRules gets the alertRules property value. The collection of alert rules.
func (m *Monitoring) GetAlertRules()([]AlertRuleable) {
    val, err := m.GetBackingStore().Get("alertRules")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AlertRuleable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Monitoring) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["alertRecords"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAlertRecordFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AlertRecordable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AlertRecordable)
                }
            }
            m.SetAlertRecords(res)
        }
        return nil
    }
    res["alertRules"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAlertRuleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AlertRuleable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AlertRuleable)
                }
            }
            m.SetAlertRules(res)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *Monitoring) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *Monitoring) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAlertRecords() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAlertRecords()))
        for i, v := range m.GetAlertRecords() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("alertRecords", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAlertRules() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAlertRules()))
        for i, v := range m.GetAlertRules() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("alertRules", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAlertRecords sets the alertRecords property value. The collection of records of alert events.
func (m *Monitoring) SetAlertRecords(value []AlertRecordable)() {
    err := m.GetBackingStore().Set("alertRecords", value)
    if err != nil {
        panic(err)
    }
}
// SetAlertRules sets the alertRules property value. The collection of alert rules.
func (m *Monitoring) SetAlertRules(value []AlertRuleable)() {
    err := m.GetBackingStore().Set("alertRules", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *Monitoring) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// Monitoringable 
type Monitoringable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAlertRecords()([]AlertRecordable)
    GetAlertRules()([]AlertRuleable)
    GetOdataType()(*string)
    SetAlertRecords(value []AlertRecordable)()
    SetAlertRules(value []AlertRuleable)()
    SetOdataType(value *string)()
}
