package managedtenants

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type ManagementActionDeploymentStatus struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewManagementActionDeploymentStatus instantiates a new ManagementActionDeploymentStatus and sets the default values.
func NewManagementActionDeploymentStatus()(*ManagementActionDeploymentStatus) {
    m := &ManagementActionDeploymentStatus{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateManagementActionDeploymentStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateManagementActionDeploymentStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagementActionDeploymentStatus(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ManagementActionDeploymentStatus) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *ManagementActionDeploymentStatus) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ManagementActionDeploymentStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["managementActionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementActionId(val)
        }
        return nil
    }
    res["managementTemplateId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementTemplateId(val)
        }
        return nil
    }
    res["managementTemplateVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementTemplateVersion(val)
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
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagementActionStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*ManagementActionStatus))
        }
        return nil
    }
    res["workloadActionDeploymentStatuses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWorkloadActionDeploymentStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WorkloadActionDeploymentStatusable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(WorkloadActionDeploymentStatusable)
                }
            }
            m.SetWorkloadActionDeploymentStatuses(res)
        }
        return nil
    }
    return res
}
// GetManagementActionId gets the managementActionId property value. The identifier for the management action. Required. Read-only.
// returns a *string when successful
func (m *ManagementActionDeploymentStatus) GetManagementActionId()(*string) {
    val, err := m.GetBackingStore().Get("managementActionId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetManagementTemplateId gets the managementTemplateId property value. The management template identifier that was used to generate the management action. Required. Read-only.
// returns a *string when successful
func (m *ManagementActionDeploymentStatus) GetManagementTemplateId()(*string) {
    val, err := m.GetBackingStore().Get("managementTemplateId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetManagementTemplateVersion gets the managementTemplateVersion property value. The managementTemplateVersion property
// returns a *int32 when successful
func (m *ManagementActionDeploymentStatus) GetManagementTemplateVersion()(*int32) {
    val, err := m.GetBackingStore().Get("managementTemplateVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *ManagementActionDeploymentStatus) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetStatus gets the status property value. The status property
// returns a *ManagementActionStatus when successful
func (m *ManagementActionDeploymentStatus) GetStatus()(*ManagementActionStatus) {
    val, err := m.GetBackingStore().Get("status")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ManagementActionStatus)
    }
    return nil
}
// GetWorkloadActionDeploymentStatuses gets the workloadActionDeploymentStatuses property value. The collection of workload action deployment statues for the given management action. Optional.
// returns a []WorkloadActionDeploymentStatusable when successful
func (m *ManagementActionDeploymentStatus) GetWorkloadActionDeploymentStatuses()([]WorkloadActionDeploymentStatusable) {
    val, err := m.GetBackingStore().Get("workloadActionDeploymentStatuses")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]WorkloadActionDeploymentStatusable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ManagementActionDeploymentStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("managementActionId", m.GetManagementActionId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("managementTemplateId", m.GetManagementTemplateId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("managementTemplateVersion", m.GetManagementTemplateVersion())
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
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err := writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetWorkloadActionDeploymentStatuses() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWorkloadActionDeploymentStatuses()))
        for i, v := range m.GetWorkloadActionDeploymentStatuses() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("workloadActionDeploymentStatuses", cast)
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ManagementActionDeploymentStatus) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *ManagementActionDeploymentStatus) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetManagementActionId sets the managementActionId property value. The identifier for the management action. Required. Read-only.
func (m *ManagementActionDeploymentStatus) SetManagementActionId(value *string)() {
    err := m.GetBackingStore().Set("managementActionId", value)
    if err != nil {
        panic(err)
    }
}
// SetManagementTemplateId sets the managementTemplateId property value. The management template identifier that was used to generate the management action. Required. Read-only.
func (m *ManagementActionDeploymentStatus) SetManagementTemplateId(value *string)() {
    err := m.GetBackingStore().Set("managementTemplateId", value)
    if err != nil {
        panic(err)
    }
}
// SetManagementTemplateVersion sets the managementTemplateVersion property value. The managementTemplateVersion property
func (m *ManagementActionDeploymentStatus) SetManagementTemplateVersion(value *int32)() {
    err := m.GetBackingStore().Set("managementTemplateVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ManagementActionDeploymentStatus) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetStatus sets the status property value. The status property
func (m *ManagementActionDeploymentStatus) SetStatus(value *ManagementActionStatus)() {
    err := m.GetBackingStore().Set("status", value)
    if err != nil {
        panic(err)
    }
}
// SetWorkloadActionDeploymentStatuses sets the workloadActionDeploymentStatuses property value. The collection of workload action deployment statues for the given management action. Optional.
func (m *ManagementActionDeploymentStatus) SetWorkloadActionDeploymentStatuses(value []WorkloadActionDeploymentStatusable)() {
    err := m.GetBackingStore().Set("workloadActionDeploymentStatuses", value)
    if err != nil {
        panic(err)
    }
}
type ManagementActionDeploymentStatusable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetManagementActionId()(*string)
    GetManagementTemplateId()(*string)
    GetManagementTemplateVersion()(*int32)
    GetOdataType()(*string)
    GetStatus()(*ManagementActionStatus)
    GetWorkloadActionDeploymentStatuses()([]WorkloadActionDeploymentStatusable)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetManagementActionId(value *string)()
    SetManagementTemplateId(value *string)()
    SetManagementTemplateVersion(value *int32)()
    SetOdataType(value *string)()
    SetStatus(value *ManagementActionStatus)()
    SetWorkloadActionDeploymentStatuses(value []WorkloadActionDeploymentStatusable)()
}
