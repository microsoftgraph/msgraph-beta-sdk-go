package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CrossTenantUserSyncInbound 
type CrossTenantUserSyncInbound struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Defines whether user objects should be synchronized from the partner tenant. If set to false, any current user synchronization from the source tenant to the target tenant will stop. There is no impact on existing users that have already been synchronized.
    isSyncAllowed *bool
    // The OdataType property
    odataType *string
}
// NewCrossTenantUserSyncInbound instantiates a new crossTenantUserSyncInbound and sets the default values.
func NewCrossTenantUserSyncInbound()(*CrossTenantUserSyncInbound) {
    m := &CrossTenantUserSyncInbound{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCrossTenantUserSyncInboundFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCrossTenantUserSyncInboundFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCrossTenantUserSyncInbound(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CrossTenantUserSyncInbound) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CrossTenantUserSyncInbound) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["isSyncAllowed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSyncAllowed(val)
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
// GetIsSyncAllowed gets the isSyncAllowed property value. Defines whether user objects should be synchronized from the partner tenant. If set to false, any current user synchronization from the source tenant to the target tenant will stop. There is no impact on existing users that have already been synchronized.
func (m *CrossTenantUserSyncInbound) GetIsSyncAllowed()(*bool) {
    return m.isSyncAllowed
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *CrossTenantUserSyncInbound) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *CrossTenantUserSyncInbound) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isSyncAllowed", m.GetIsSyncAllowed())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CrossTenantUserSyncInbound) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIsSyncAllowed sets the isSyncAllowed property value. Defines whether user objects should be synchronized from the partner tenant. If set to false, any current user synchronization from the source tenant to the target tenant will stop. There is no impact on existing users that have already been synchronized.
func (m *CrossTenantUserSyncInbound) SetIsSyncAllowed(value *bool)() {
    m.isSyncAllowed = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CrossTenantUserSyncInbound) SetOdataType(value *string)() {
    m.odataType = value
}
