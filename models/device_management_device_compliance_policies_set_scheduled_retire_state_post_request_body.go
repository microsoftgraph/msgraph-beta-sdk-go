package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementDeviceCompliancePoliciesSetScheduledRetireStatePostRequestBody provides operations to call the setScheduledRetireState method.
type DeviceManagementDeviceCompliancePoliciesSetScheduledRetireStatePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The managedDeviceIds property
    managedDeviceIds []string
    // The scopedToAllDevices property
    scopedToAllDevices *bool
    // Cancel or confirm scheduled retire 
    state *ScheduledRetireState
}
// NewDeviceManagementDeviceCompliancePoliciesSetScheduledRetireStatePostRequestBody instantiates a new DeviceManagementDeviceCompliancePoliciesSetScheduledRetireStatePostRequestBody and sets the default values.
func NewDeviceManagementDeviceCompliancePoliciesSetScheduledRetireStatePostRequestBody()(*DeviceManagementDeviceCompliancePoliciesSetScheduledRetireStatePostRequestBody) {
    m := &DeviceManagementDeviceCompliancePoliciesSetScheduledRetireStatePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementDeviceCompliancePoliciesSetScheduledRetireStatePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementDeviceCompliancePoliciesSetScheduledRetireStatePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementDeviceCompliancePoliciesSetScheduledRetireStatePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementDeviceCompliancePoliciesSetScheduledRetireStatePostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementDeviceCompliancePoliciesSetScheduledRetireStatePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["managedDeviceIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetManagedDeviceIds(res)
        }
        return nil
    }
    res["scopedToAllDevices"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScopedToAllDevices(val)
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseScheduledRetireState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*ScheduledRetireState))
        }
        return nil
    }
    return res
}
// GetManagedDeviceIds gets the managedDeviceIds property value. The managedDeviceIds property
func (m *DeviceManagementDeviceCompliancePoliciesSetScheduledRetireStatePostRequestBody) GetManagedDeviceIds()([]string) {
    return m.managedDeviceIds
}
// GetScopedToAllDevices gets the scopedToAllDevices property value. The scopedToAllDevices property
func (m *DeviceManagementDeviceCompliancePoliciesSetScheduledRetireStatePostRequestBody) GetScopedToAllDevices()(*bool) {
    return m.scopedToAllDevices
}
// GetState gets the state property value. Cancel or confirm scheduled retire 
func (m *DeviceManagementDeviceCompliancePoliciesSetScheduledRetireStatePostRequestBody) GetState()(*ScheduledRetireState) {
    return m.state
}
// Serialize serializes information the current object
func (m *DeviceManagementDeviceCompliancePoliciesSetScheduledRetireStatePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetManagedDeviceIds() != nil {
        err := writer.WriteCollectionOfStringValues("managedDeviceIds", m.GetManagedDeviceIds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("scopedToAllDevices", m.GetScopedToAllDevices())
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err := writer.WriteStringValue("state", &cast)
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
func (m *DeviceManagementDeviceCompliancePoliciesSetScheduledRetireStatePostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetManagedDeviceIds sets the managedDeviceIds property value. The managedDeviceIds property
func (m *DeviceManagementDeviceCompliancePoliciesSetScheduledRetireStatePostRequestBody) SetManagedDeviceIds(value []string)() {
    m.managedDeviceIds = value
}
// SetScopedToAllDevices sets the scopedToAllDevices property value. The scopedToAllDevices property
func (m *DeviceManagementDeviceCompliancePoliciesSetScheduledRetireStatePostRequestBody) SetScopedToAllDevices(value *bool)() {
    m.scopedToAllDevices = value
}
// SetState sets the state property value. Cancel or confirm scheduled retire 
func (m *DeviceManagementDeviceCompliancePoliciesSetScheduledRetireStatePostRequestBody) SetState(value *ScheduledRetireState)() {
    m.state = value
}
