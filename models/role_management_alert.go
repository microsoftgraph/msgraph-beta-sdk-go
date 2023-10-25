package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RoleManagementAlert 
type RoleManagementAlert struct {
    Entity
}
// NewRoleManagementAlert instantiates a new roleManagementAlert and sets the default values.
func NewRoleManagementAlert()(*RoleManagementAlert) {
    m := &RoleManagementAlert{
        Entity: *NewEntity(),
    }
    return m
}
// CreateRoleManagementAlertFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRoleManagementAlertFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRoleManagementAlert(), nil
}
// GetAlertConfigurations gets the alertConfigurations property value. The various configurations of an alert for Microsoft Entra roles. The configurations are predefined and can't be created or deleted, but some of the configurations can be modified.
func (m *RoleManagementAlert) GetAlertConfigurations()([]UnifiedRoleManagementAlertConfigurationable) {
    val, err := m.GetBackingStore().Get("alertConfigurations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UnifiedRoleManagementAlertConfigurationable)
    }
    return nil
}
// GetAlertDefinitions gets the alertDefinitions property value. Defines an alert, its impact, and measures to mitigate or prevent it.
func (m *RoleManagementAlert) GetAlertDefinitions()([]UnifiedRoleManagementAlertDefinitionable) {
    val, err := m.GetBackingStore().Get("alertDefinitions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UnifiedRoleManagementAlertDefinitionable)
    }
    return nil
}
// GetAlerts gets the alerts property value. Represents the alert entity.
func (m *RoleManagementAlert) GetAlerts()([]UnifiedRoleManagementAlertable) {
    val, err := m.GetBackingStore().Get("alerts")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UnifiedRoleManagementAlertable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RoleManagementAlert) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["alertConfigurations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUnifiedRoleManagementAlertConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnifiedRoleManagementAlertConfigurationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UnifiedRoleManagementAlertConfigurationable)
                }
            }
            m.SetAlertConfigurations(res)
        }
        return nil
    }
    res["alertDefinitions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUnifiedRoleManagementAlertDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnifiedRoleManagementAlertDefinitionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UnifiedRoleManagementAlertDefinitionable)
                }
            }
            m.SetAlertDefinitions(res)
        }
        return nil
    }
    res["alerts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUnifiedRoleManagementAlertFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnifiedRoleManagementAlertable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UnifiedRoleManagementAlertable)
                }
            }
            m.SetAlerts(res)
        }
        return nil
    }
    res["operations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateLongRunningOperationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]LongRunningOperationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(LongRunningOperationable)
                }
            }
            m.SetOperations(res)
        }
        return nil
    }
    return res
}
// GetOperations gets the operations property value. Represents operations on resources that take a long time to complete and can run in the background until completion.
func (m *RoleManagementAlert) GetOperations()([]LongRunningOperationable) {
    val, err := m.GetBackingStore().Get("operations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]LongRunningOperationable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *RoleManagementAlert) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetAlertDefinitions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAlertDefinitions()))
        for i, v := range m.GetAlertDefinitions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("alertDefinitions", cast)
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
    if m.GetOperations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOperations()))
        for i, v := range m.GetOperations() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("operations", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAlertConfigurations sets the alertConfigurations property value. The various configurations of an alert for Microsoft Entra roles. The configurations are predefined and can't be created or deleted, but some of the configurations can be modified.
func (m *RoleManagementAlert) SetAlertConfigurations(value []UnifiedRoleManagementAlertConfigurationable)() {
    err := m.GetBackingStore().Set("alertConfigurations", value)
    if err != nil {
        panic(err)
    }
}
// SetAlertDefinitions sets the alertDefinitions property value. Defines an alert, its impact, and measures to mitigate or prevent it.
func (m *RoleManagementAlert) SetAlertDefinitions(value []UnifiedRoleManagementAlertDefinitionable)() {
    err := m.GetBackingStore().Set("alertDefinitions", value)
    if err != nil {
        panic(err)
    }
}
// SetAlerts sets the alerts property value. Represents the alert entity.
func (m *RoleManagementAlert) SetAlerts(value []UnifiedRoleManagementAlertable)() {
    err := m.GetBackingStore().Set("alerts", value)
    if err != nil {
        panic(err)
    }
}
// SetOperations sets the operations property value. Represents operations on resources that take a long time to complete and can run in the background until completion.
func (m *RoleManagementAlert) SetOperations(value []LongRunningOperationable)() {
    err := m.GetBackingStore().Set("operations", value)
    if err != nil {
        panic(err)
    }
}
// RoleManagementAlertable 
type RoleManagementAlertable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAlertConfigurations()([]UnifiedRoleManagementAlertConfigurationable)
    GetAlertDefinitions()([]UnifiedRoleManagementAlertDefinitionable)
    GetAlerts()([]UnifiedRoleManagementAlertable)
    GetOperations()([]LongRunningOperationable)
    SetAlertConfigurations(value []UnifiedRoleManagementAlertConfigurationable)()
    SetAlertDefinitions(value []UnifiedRoleManagementAlertDefinitionable)()
    SetAlerts(value []UnifiedRoleManagementAlertable)()
    SetOperations(value []LongRunningOperationable)()
}
