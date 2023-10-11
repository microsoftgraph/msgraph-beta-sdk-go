package managedtenants

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// WorkloadActionDeploymentStatus 
type WorkloadActionDeploymentStatus struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewWorkloadActionDeploymentStatus instantiates a new workloadActionDeploymentStatus and sets the default values.
func NewWorkloadActionDeploymentStatus()(*WorkloadActionDeploymentStatus) {
    m := &WorkloadActionDeploymentStatus{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWorkloadActionDeploymentStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkloadActionDeploymentStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkloadActionDeploymentStatus(), nil
}
// GetActionId gets the actionId property value. The unique identifier for the workload action. Required. Read-only.
func (m *WorkloadActionDeploymentStatus) GetActionId()(*string) {
    val, err := m.GetBackingStore().Get("actionId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WorkloadActionDeploymentStatus) GetAdditionalData()(map[string]any) {
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
func (m *WorkloadActionDeploymentStatus) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDeployedPolicyId gets the deployedPolicyId property value. The identifier of any policy that was created by applying the workload action. Optional. Read-only.
func (m *WorkloadActionDeploymentStatus) GetDeployedPolicyId()(*string) {
    val, err := m.GetBackingStore().Get("deployedPolicyId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetError gets the error property value. The detailed information for exceptions that occur when deploying the workload action. Optional. Required.
func (m *WorkloadActionDeploymentStatus) GetError()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GenericErrorable) {
    val, err := m.GetBackingStore().Get("error")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GenericErrorable)
    }
    return nil
}
// GetExcludeGroups gets the excludeGroups property value. The excludeGroups property
func (m *WorkloadActionDeploymentStatus) GetExcludeGroups()([]string) {
    val, err := m.GetBackingStore().Get("excludeGroups")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkloadActionDeploymentStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["actionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionId(val)
        }
        return nil
    }
    res["deployedPolicyId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeployedPolicyId(val)
        }
        return nil
    }
    res["error"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGenericErrorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetError(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GenericErrorable))
        }
        return nil
    }
    res["excludeGroups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetExcludeGroups(res)
        }
        return nil
    }
    res["includeAllUsers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIncludeAllUsers(val)
        }
        return nil
    }
    res["includeGroups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetIncludeGroups(res)
        }
        return nil
    }
    res["lastDeploymentDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastDeploymentDateTime(val)
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
        val, err := n.GetEnumValue(ParseWorkloadActionStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*WorkloadActionStatus))
        }
        return nil
    }
    return res
}
// GetIncludeAllUsers gets the includeAllUsers property value. The includeAllUsers property
func (m *WorkloadActionDeploymentStatus) GetIncludeAllUsers()(*bool) {
    val, err := m.GetBackingStore().Get("includeAllUsers")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIncludeGroups gets the includeGroups property value. The includeGroups property
func (m *WorkloadActionDeploymentStatus) GetIncludeGroups()([]string) {
    val, err := m.GetBackingStore().Get("includeGroups")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetLastDeploymentDateTime gets the lastDeploymentDateTime property value. The date and time the workload action was last deployed. Optional.
func (m *WorkloadActionDeploymentStatus) GetLastDeploymentDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastDeploymentDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *WorkloadActionDeploymentStatus) GetOdataType()(*string) {
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
func (m *WorkloadActionDeploymentStatus) GetStatus()(*WorkloadActionStatus) {
    val, err := m.GetBackingStore().Get("status")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*WorkloadActionStatus)
    }
    return nil
}
// Serialize serializes information the current object
func (m *WorkloadActionDeploymentStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("actionId", m.GetActionId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("deployedPolicyId", m.GetDeployedPolicyId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("error", m.GetError())
        if err != nil {
            return err
        }
    }
    if m.GetExcludeGroups() != nil {
        err := writer.WriteCollectionOfStringValues("excludeGroups", m.GetExcludeGroups())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("includeAllUsers", m.GetIncludeAllUsers())
        if err != nil {
            return err
        }
    }
    if m.GetIncludeGroups() != nil {
        err := writer.WriteCollectionOfStringValues("includeGroups", m.GetIncludeGroups())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("lastDeploymentDateTime", m.GetLastDeploymentDateTime())
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
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActionId sets the actionId property value. The unique identifier for the workload action. Required. Read-only.
func (m *WorkloadActionDeploymentStatus) SetActionId(value *string)() {
    err := m.GetBackingStore().Set("actionId", value)
    if err != nil {
        panic(err)
    }
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WorkloadActionDeploymentStatus) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *WorkloadActionDeploymentStatus) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDeployedPolicyId sets the deployedPolicyId property value. The identifier of any policy that was created by applying the workload action. Optional. Read-only.
func (m *WorkloadActionDeploymentStatus) SetDeployedPolicyId(value *string)() {
    err := m.GetBackingStore().Set("deployedPolicyId", value)
    if err != nil {
        panic(err)
    }
}
// SetError sets the error property value. The detailed information for exceptions that occur when deploying the workload action. Optional. Required.
func (m *WorkloadActionDeploymentStatus) SetError(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GenericErrorable)() {
    err := m.GetBackingStore().Set("error", value)
    if err != nil {
        panic(err)
    }
}
// SetExcludeGroups sets the excludeGroups property value. The excludeGroups property
func (m *WorkloadActionDeploymentStatus) SetExcludeGroups(value []string)() {
    err := m.GetBackingStore().Set("excludeGroups", value)
    if err != nil {
        panic(err)
    }
}
// SetIncludeAllUsers sets the includeAllUsers property value. The includeAllUsers property
func (m *WorkloadActionDeploymentStatus) SetIncludeAllUsers(value *bool)() {
    err := m.GetBackingStore().Set("includeAllUsers", value)
    if err != nil {
        panic(err)
    }
}
// SetIncludeGroups sets the includeGroups property value. The includeGroups property
func (m *WorkloadActionDeploymentStatus) SetIncludeGroups(value []string)() {
    err := m.GetBackingStore().Set("includeGroups", value)
    if err != nil {
        panic(err)
    }
}
// SetLastDeploymentDateTime sets the lastDeploymentDateTime property value. The date and time the workload action was last deployed. Optional.
func (m *WorkloadActionDeploymentStatus) SetLastDeploymentDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastDeploymentDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *WorkloadActionDeploymentStatus) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetStatus sets the status property value. The status property
func (m *WorkloadActionDeploymentStatus) SetStatus(value *WorkloadActionStatus)() {
    err := m.GetBackingStore().Set("status", value)
    if err != nil {
        panic(err)
    }
}
// WorkloadActionDeploymentStatusable 
type WorkloadActionDeploymentStatusable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActionId()(*string)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDeployedPolicyId()(*string)
    GetError()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GenericErrorable)
    GetExcludeGroups()([]string)
    GetIncludeAllUsers()(*bool)
    GetIncludeGroups()([]string)
    GetLastDeploymentDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetOdataType()(*string)
    GetStatus()(*WorkloadActionStatus)
    SetActionId(value *string)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDeployedPolicyId(value *string)()
    SetError(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GenericErrorable)()
    SetExcludeGroups(value []string)()
    SetIncludeAllUsers(value *bool)()
    SetIncludeGroups(value []string)()
    SetLastDeploymentDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetOdataType(value *string)()
    SetStatus(value *WorkloadActionStatus)()
}
