package healthmonitoring

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

type AlertConfiguration struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewAlertConfiguration instantiates a new AlertConfiguration and sets the default values.
func NewAlertConfiguration()(*AlertConfiguration) {
    m := &AlertConfiguration{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateAlertConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAlertConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAlertConfiguration(), nil
}
// GetEmailNotificationConfigurations gets the emailNotificationConfigurations property value. The emailNotificationConfigurations property
// returns a []EmailNotificationConfigurationable when successful
func (m *AlertConfiguration) GetEmailNotificationConfigurations()([]EmailNotificationConfigurationable) {
    val, err := m.GetBackingStore().Get("emailNotificationConfigurations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]EmailNotificationConfigurationable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AlertConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["emailNotificationConfigurations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEmailNotificationConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EmailNotificationConfigurationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(EmailNotificationConfigurationable)
                }
            }
            m.SetEmailNotificationConfigurations(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *AlertConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetEmailNotificationConfigurations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEmailNotificationConfigurations()))
        for i, v := range m.GetEmailNotificationConfigurations() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("emailNotificationConfigurations", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEmailNotificationConfigurations sets the emailNotificationConfigurations property value. The emailNotificationConfigurations property
func (m *AlertConfiguration) SetEmailNotificationConfigurations(value []EmailNotificationConfigurationable)() {
    err := m.GetBackingStore().Set("emailNotificationConfigurations", value)
    if err != nil {
        panic(err)
    }
}
type AlertConfigurationable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEmailNotificationConfigurations()([]EmailNotificationConfigurationable)
    SetEmailNotificationConfigurations(value []EmailNotificationConfigurationable)()
}
