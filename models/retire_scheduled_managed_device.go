package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// RetireScheduledManagedDevice managedDevices that are scheduled for retire
type RetireScheduledManagedDevice struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewRetireScheduledManagedDevice instantiates a new retireScheduledManagedDevice and sets the default values.
func NewRetireScheduledManagedDevice()(*RetireScheduledManagedDevice) {
    m := &RetireScheduledManagedDevice{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRetireScheduledManagedDeviceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRetireScheduledManagedDeviceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRetireScheduledManagedDevice(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RetireScheduledManagedDevice) GetAdditionalData()(map[string]any) {
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
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *RetireScheduledManagedDevice) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetComplianceState gets the complianceState property value. The complianceState property
func (m *RetireScheduledManagedDevice) GetComplianceState()(*ComplianceStatus) {
    val, err := m.GetBackingStore().Get("complianceState")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ComplianceStatus)
    }
    return nil
}
// GetDeviceCompliancePolicyId gets the deviceCompliancePolicyId property value. Device Compliance PolicyId
func (m *RetireScheduledManagedDevice) GetDeviceCompliancePolicyId()(*string) {
    val, err := m.GetBackingStore().Get("deviceCompliancePolicyId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceCompliancePolicyName gets the deviceCompliancePolicyName property value. Device Compliance Policy Name
func (m *RetireScheduledManagedDevice) GetDeviceCompliancePolicyName()(*string) {
    val, err := m.GetBackingStore().Get("deviceCompliancePolicyName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceType gets the deviceType property value. Device type.
func (m *RetireScheduledManagedDevice) GetDeviceType()(*DeviceType) {
    val, err := m.GetBackingStore().Get("deviceType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeviceType)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RetireScheduledManagedDevice) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["complianceState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseComplianceStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComplianceState(val.(*ComplianceStatus))
        }
        return nil
    }
    res["deviceCompliancePolicyId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceCompliancePolicyId(val)
        }
        return nil
    }
    res["deviceCompliancePolicyName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceCompliancePolicyName(val)
        }
        return nil
    }
    res["deviceType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceType(val.(*DeviceType))
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["managedDeviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedDeviceId(val)
        }
        return nil
    }
    res["managedDeviceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedDeviceName(val)
        }
        return nil
    }
    res["managementAgent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagementAgentType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementAgent(val.(*ManagementAgentType))
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
    res["ownerType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedDeviceOwnerType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwnerType(val.(*ManagedDeviceOwnerType))
        }
        return nil
    }
    res["retireAfterDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRetireAfterDateTime(val)
        }
        return nil
    }
    res["roleScopeTagIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetRoleScopeTagIds(res)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. Key of the entity.
func (m *RetireScheduledManagedDevice) GetId()(*string) {
    val, err := m.GetBackingStore().Get("id")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetManagedDeviceId gets the managedDeviceId property value. Managed DeviceId
func (m *RetireScheduledManagedDevice) GetManagedDeviceId()(*string) {
    val, err := m.GetBackingStore().Get("managedDeviceId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetManagedDeviceName gets the managedDeviceName property value. Managed Device Name
func (m *RetireScheduledManagedDevice) GetManagedDeviceName()(*string) {
    val, err := m.GetBackingStore().Get("managedDeviceName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetManagementAgent gets the managementAgent property value. Management agent type.
func (m *RetireScheduledManagedDevice) GetManagementAgent()(*ManagementAgentType) {
    val, err := m.GetBackingStore().Get("managementAgent")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ManagementAgentType)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *RetireScheduledManagedDevice) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOwnerType gets the ownerType property value. Owner type of device.
func (m *RetireScheduledManagedDevice) GetOwnerType()(*ManagedDeviceOwnerType) {
    val, err := m.GetBackingStore().Get("ownerType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ManagedDeviceOwnerType)
    }
    return nil
}
// GetRetireAfterDateTime gets the retireAfterDateTime property value. Managed Device Retire After DateTime
func (m *RetireScheduledManagedDevice) GetRetireAfterDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("retireAfterDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. List of Scope Tags for this Entity instance.
func (m *RetireScheduledManagedDevice) GetRoleScopeTagIds()([]string) {
    val, err := m.GetBackingStore().Get("roleScopeTagIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *RetireScheduledManagedDevice) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetComplianceState() != nil {
        cast := (*m.GetComplianceState()).String()
        err := writer.WriteStringValue("complianceState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("deviceCompliancePolicyId", m.GetDeviceCompliancePolicyId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("deviceCompliancePolicyName", m.GetDeviceCompliancePolicyName())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceType() != nil {
        cast := (*m.GetDeviceType()).String()
        err := writer.WriteStringValue("deviceType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("managedDeviceId", m.GetManagedDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("managedDeviceName", m.GetManagedDeviceName())
        if err != nil {
            return err
        }
    }
    if m.GetManagementAgent() != nil {
        cast := (*m.GetManagementAgent()).String()
        err := writer.WriteStringValue("managementAgent", &cast)
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
    if m.GetOwnerType() != nil {
        cast := (*m.GetOwnerType()).String()
        err := writer.WriteStringValue("ownerType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("retireAfterDateTime", m.GetRetireAfterDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetRoleScopeTagIds() != nil {
        err := writer.WriteCollectionOfStringValues("roleScopeTagIds", m.GetRoleScopeTagIds())
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
func (m *RetireScheduledManagedDevice) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *RetireScheduledManagedDevice) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetComplianceState sets the complianceState property value. The complianceState property
func (m *RetireScheduledManagedDevice) SetComplianceState(value *ComplianceStatus)() {
    err := m.GetBackingStore().Set("complianceState", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceCompliancePolicyId sets the deviceCompliancePolicyId property value. Device Compliance PolicyId
func (m *RetireScheduledManagedDevice) SetDeviceCompliancePolicyId(value *string)() {
    err := m.GetBackingStore().Set("deviceCompliancePolicyId", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceCompliancePolicyName sets the deviceCompliancePolicyName property value. Device Compliance Policy Name
func (m *RetireScheduledManagedDevice) SetDeviceCompliancePolicyName(value *string)() {
    err := m.GetBackingStore().Set("deviceCompliancePolicyName", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceType sets the deviceType property value. Device type.
func (m *RetireScheduledManagedDevice) SetDeviceType(value *DeviceType)() {
    err := m.GetBackingStore().Set("deviceType", value)
    if err != nil {
        panic(err)
    }
}
// SetId sets the id property value. Key of the entity.
func (m *RetireScheduledManagedDevice) SetId(value *string)() {
    err := m.GetBackingStore().Set("id", value)
    if err != nil {
        panic(err)
    }
}
// SetManagedDeviceId sets the managedDeviceId property value. Managed DeviceId
func (m *RetireScheduledManagedDevice) SetManagedDeviceId(value *string)() {
    err := m.GetBackingStore().Set("managedDeviceId", value)
    if err != nil {
        panic(err)
    }
}
// SetManagedDeviceName sets the managedDeviceName property value. Managed Device Name
func (m *RetireScheduledManagedDevice) SetManagedDeviceName(value *string)() {
    err := m.GetBackingStore().Set("managedDeviceName", value)
    if err != nil {
        panic(err)
    }
}
// SetManagementAgent sets the managementAgent property value. Management agent type.
func (m *RetireScheduledManagedDevice) SetManagementAgent(value *ManagementAgentType)() {
    err := m.GetBackingStore().Set("managementAgent", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *RetireScheduledManagedDevice) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetOwnerType sets the ownerType property value. Owner type of device.
func (m *RetireScheduledManagedDevice) SetOwnerType(value *ManagedDeviceOwnerType)() {
    err := m.GetBackingStore().Set("ownerType", value)
    if err != nil {
        panic(err)
    }
}
// SetRetireAfterDateTime sets the retireAfterDateTime property value. Managed Device Retire After DateTime
func (m *RetireScheduledManagedDevice) SetRetireAfterDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("retireAfterDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. List of Scope Tags for this Entity instance.
func (m *RetireScheduledManagedDevice) SetRoleScopeTagIds(value []string)() {
    err := m.GetBackingStore().Set("roleScopeTagIds", value)
    if err != nil {
        panic(err)
    }
}
// RetireScheduledManagedDeviceable 
type RetireScheduledManagedDeviceable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetComplianceState()(*ComplianceStatus)
    GetDeviceCompliancePolicyId()(*string)
    GetDeviceCompliancePolicyName()(*string)
    GetDeviceType()(*DeviceType)
    GetId()(*string)
    GetManagedDeviceId()(*string)
    GetManagedDeviceName()(*string)
    GetManagementAgent()(*ManagementAgentType)
    GetOdataType()(*string)
    GetOwnerType()(*ManagedDeviceOwnerType)
    GetRetireAfterDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRoleScopeTagIds()([]string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetComplianceState(value *ComplianceStatus)()
    SetDeviceCompliancePolicyId(value *string)()
    SetDeviceCompliancePolicyName(value *string)()
    SetDeviceType(value *DeviceType)()
    SetId(value *string)()
    SetManagedDeviceId(value *string)()
    SetManagedDeviceName(value *string)()
    SetManagementAgent(value *ManagementAgentType)()
    SetOdataType(value *string)()
    SetOwnerType(value *ManagedDeviceOwnerType)()
    SetRetireAfterDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRoleScopeTagIds(value []string)()
}
