package healthmonitoring

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

type HealthMonitoringRoot struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewHealthMonitoringRoot instantiates a new HealthMonitoringRoot and sets the default values.
func NewHealthMonitoringRoot()(*HealthMonitoringRoot) {
    m := &HealthMonitoringRoot{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateHealthMonitoringRootFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateHealthMonitoringRootFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewHealthMonitoringRoot(), nil
}
// GetAlertConfigurations gets the alertConfigurations property value. The configuration of an alert type, which defines behavior that occurs when an alert is created.
// returns a []AlertConfigurationable when successful
func (m *HealthMonitoringRoot) GetAlertConfigurations()([]AlertConfigurationable) {
    val, err := m.GetBackingStore().Get("alertConfigurations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AlertConfigurationable)
    }
    return nil
}
// GetAlerts gets the alerts property value. The collection of health monitoring system detected alerts for anomalous usage patterns found in a Microsoft Entra tenant.
// returns a []Alertable when successful
func (m *HealthMonitoringRoot) GetAlerts()([]Alertable) {
    val, err := m.GetBackingStore().Get("alerts")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]Alertable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *HealthMonitoringRoot) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["alertConfigurations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAlertConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AlertConfigurationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AlertConfigurationable)
                }
            }
            m.SetAlertConfigurations(res)
        }
        return nil
    }
    res["alerts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAlertFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Alertable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Alertable)
                }
            }
            m.SetAlerts(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *HealthMonitoringRoot) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAlertConfigurations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAlertConfigurations()))
        for i, v := range m.GetAlertConfigurations() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("alertConfigurations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAlerts() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAlerts()))
        for i, v := range m.GetAlerts() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("alerts", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAlertConfigurations sets the alertConfigurations property value. The configuration of an alert type, which defines behavior that occurs when an alert is created.
func (m *HealthMonitoringRoot) SetAlertConfigurations(value []AlertConfigurationable)() {
    err := m.GetBackingStore().Set("alertConfigurations", value)
    if err != nil {
        panic(err)
    }
}
// SetAlerts sets the alerts property value. The collection of health monitoring system detected alerts for anomalous usage patterns found in a Microsoft Entra tenant.
func (m *HealthMonitoringRoot) SetAlerts(value []Alertable)() {
    err := m.GetBackingStore().Set("alerts", value)
    if err != nil {
        panic(err)
    }
}
type HealthMonitoringRootable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAlertConfigurations()([]AlertConfigurationable)
    GetAlerts()([]Alertable)
    SetAlertConfigurations(value []AlertConfigurationable)()
    SetAlerts(value []Alertable)()
}
