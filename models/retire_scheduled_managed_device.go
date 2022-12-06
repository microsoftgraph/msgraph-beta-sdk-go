package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RetireScheduledManagedDevice managedDevices that are scheduled for retire
type RetireScheduledManagedDevice struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The complianceState property
    complianceState *ComplianceStatus
    // Device Compliance PolicyId
    deviceCompliancePolicyId *string
    // Device Compliance Policy Name
    deviceCompliancePolicyName *string
    // Device type.
    deviceType *DeviceType
    // Key of the entity.
    id *string
    // Managed DeviceId
    managedDeviceId *string
    // Managed Device Name
    managedDeviceName *string
    // Management agent type.
    managementAgent *ManagementAgentType
    // The OdataType property
    odataType *string
    // Owner type of device.
    ownerType *ManagedDeviceOwnerType
    // Managed Device Retire After DateTime
    retireAfterDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // List of Scope Tags for this Entity instance.
    roleScopeTagIds []string
}
// NewRetireScheduledManagedDevice instantiates a new retireScheduledManagedDevice and sets the default values.
func NewRetireScheduledManagedDevice()(*RetireScheduledManagedDevice) {
    m := &RetireScheduledManagedDevice{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateRetireScheduledManagedDeviceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRetireScheduledManagedDeviceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRetireScheduledManagedDevice(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RetireScheduledManagedDevice) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetComplianceState gets the complianceState property value. The complianceState property
func (m *RetireScheduledManagedDevice) GetComplianceState()(*ComplianceStatus) {
    return m.complianceState
}
// GetDeviceCompliancePolicyId gets the deviceCompliancePolicyId property value. Device Compliance PolicyId
func (m *RetireScheduledManagedDevice) GetDeviceCompliancePolicyId()(*string) {
    return m.deviceCompliancePolicyId
}
// GetDeviceCompliancePolicyName gets the deviceCompliancePolicyName property value. Device Compliance Policy Name
func (m *RetireScheduledManagedDevice) GetDeviceCompliancePolicyName()(*string) {
    return m.deviceCompliancePolicyName
}
// GetDeviceType gets the deviceType property value. Device type.
func (m *RetireScheduledManagedDevice) GetDeviceType()(*DeviceType) {
    return m.deviceType
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RetireScheduledManagedDevice) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["complianceState"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseComplianceStatus , m.SetComplianceState)
    res["deviceCompliancePolicyId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDeviceCompliancePolicyId)
    res["deviceCompliancePolicyName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDeviceCompliancePolicyName)
    res["deviceType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseDeviceType , m.SetDeviceType)
    res["id"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetId)
    res["managedDeviceId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetManagedDeviceId)
    res["managedDeviceName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetManagedDeviceName)
    res["managementAgent"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseManagementAgentType , m.SetManagementAgent)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["ownerType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseManagedDeviceOwnerType , m.SetOwnerType)
    res["retireAfterDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetRetireAfterDateTime)
    res["roleScopeTagIds"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetRoleScopeTagIds)
    return res
}
// GetId gets the id property value. Key of the entity.
func (m *RetireScheduledManagedDevice) GetId()(*string) {
    return m.id
}
// GetManagedDeviceId gets the managedDeviceId property value. Managed DeviceId
func (m *RetireScheduledManagedDevice) GetManagedDeviceId()(*string) {
    return m.managedDeviceId
}
// GetManagedDeviceName gets the managedDeviceName property value. Managed Device Name
func (m *RetireScheduledManagedDevice) GetManagedDeviceName()(*string) {
    return m.managedDeviceName
}
// GetManagementAgent gets the managementAgent property value. Management agent type.
func (m *RetireScheduledManagedDevice) GetManagementAgent()(*ManagementAgentType) {
    return m.managementAgent
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *RetireScheduledManagedDevice) GetOdataType()(*string) {
    return m.odataType
}
// GetOwnerType gets the ownerType property value. Owner type of device.
func (m *RetireScheduledManagedDevice) GetOwnerType()(*ManagedDeviceOwnerType) {
    return m.ownerType
}
// GetRetireAfterDateTime gets the retireAfterDateTime property value. Managed Device Retire After DateTime
func (m *RetireScheduledManagedDevice) GetRetireAfterDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.retireAfterDateTime
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. List of Scope Tags for this Entity instance.
func (m *RetireScheduledManagedDevice) GetRoleScopeTagIds()([]string) {
    return m.roleScopeTagIds
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
func (m *RetireScheduledManagedDevice) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetComplianceState sets the complianceState property value. The complianceState property
func (m *RetireScheduledManagedDevice) SetComplianceState(value *ComplianceStatus)() {
    m.complianceState = value
}
// SetDeviceCompliancePolicyId sets the deviceCompliancePolicyId property value. Device Compliance PolicyId
func (m *RetireScheduledManagedDevice) SetDeviceCompliancePolicyId(value *string)() {
    m.deviceCompliancePolicyId = value
}
// SetDeviceCompliancePolicyName sets the deviceCompliancePolicyName property value. Device Compliance Policy Name
func (m *RetireScheduledManagedDevice) SetDeviceCompliancePolicyName(value *string)() {
    m.deviceCompliancePolicyName = value
}
// SetDeviceType sets the deviceType property value. Device type.
func (m *RetireScheduledManagedDevice) SetDeviceType(value *DeviceType)() {
    m.deviceType = value
}
// SetId sets the id property value. Key of the entity.
func (m *RetireScheduledManagedDevice) SetId(value *string)() {
    m.id = value
}
// SetManagedDeviceId sets the managedDeviceId property value. Managed DeviceId
func (m *RetireScheduledManagedDevice) SetManagedDeviceId(value *string)() {
    m.managedDeviceId = value
}
// SetManagedDeviceName sets the managedDeviceName property value. Managed Device Name
func (m *RetireScheduledManagedDevice) SetManagedDeviceName(value *string)() {
    m.managedDeviceName = value
}
// SetManagementAgent sets the managementAgent property value. Management agent type.
func (m *RetireScheduledManagedDevice) SetManagementAgent(value *ManagementAgentType)() {
    m.managementAgent = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *RetireScheduledManagedDevice) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOwnerType sets the ownerType property value. Owner type of device.
func (m *RetireScheduledManagedDevice) SetOwnerType(value *ManagedDeviceOwnerType)() {
    m.ownerType = value
}
// SetRetireAfterDateTime sets the retireAfterDateTime property value. Managed Device Retire After DateTime
func (m *RetireScheduledManagedDevice) SetRetireAfterDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.retireAfterDateTime = value
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. List of Scope Tags for this Entity instance.
func (m *RetireScheduledManagedDevice) SetRoleScopeTagIds(value []string)() {
    m.roleScopeTagIds = value
}
