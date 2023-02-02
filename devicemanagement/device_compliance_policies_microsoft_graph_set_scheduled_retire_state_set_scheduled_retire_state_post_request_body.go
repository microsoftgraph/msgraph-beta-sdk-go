package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// DeviceCompliancePoliciesMicrosoftGraphSetScheduledRetireStateSetScheduledRetireStatePostRequestBody 
type DeviceCompliancePoliciesMicrosoftGraphSetScheduledRetireStateSetScheduledRetireStatePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The managedDeviceIds property
    managedDeviceIds []string
    // The scopedToAllDevices property
    scopedToAllDevices *bool
    // Cancel or confirm scheduled retire 
    state *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ScheduledRetireState
}
// NewDeviceCompliancePoliciesMicrosoftGraphSetScheduledRetireStateSetScheduledRetireStatePostRequestBody instantiates a new DeviceCompliancePoliciesMicrosoftGraphSetScheduledRetireStateSetScheduledRetireStatePostRequestBody and sets the default values.
func NewDeviceCompliancePoliciesMicrosoftGraphSetScheduledRetireStateSetScheduledRetireStatePostRequestBody()(*DeviceCompliancePoliciesMicrosoftGraphSetScheduledRetireStateSetScheduledRetireStatePostRequestBody) {
    m := &DeviceCompliancePoliciesMicrosoftGraphSetScheduledRetireStateSetScheduledRetireStatePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDeviceCompliancePoliciesMicrosoftGraphSetScheduledRetireStateSetScheduledRetireStatePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceCompliancePoliciesMicrosoftGraphSetScheduledRetireStateSetScheduledRetireStatePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceCompliancePoliciesMicrosoftGraphSetScheduledRetireStateSetScheduledRetireStatePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceCompliancePoliciesMicrosoftGraphSetScheduledRetireStateSetScheduledRetireStatePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceCompliancePoliciesMicrosoftGraphSetScheduledRetireStateSetScheduledRetireStatePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
        val, err := n.GetEnumValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ParseScheduledRetireState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ScheduledRetireState))
        }
        return nil
    }
    return res
}
// GetManagedDeviceIds gets the managedDeviceIds property value. The managedDeviceIds property
func (m *DeviceCompliancePoliciesMicrosoftGraphSetScheduledRetireStateSetScheduledRetireStatePostRequestBody) GetManagedDeviceIds()([]string) {
    return m.managedDeviceIds
}
// GetScopedToAllDevices gets the scopedToAllDevices property value. The scopedToAllDevices property
func (m *DeviceCompliancePoliciesMicrosoftGraphSetScheduledRetireStateSetScheduledRetireStatePostRequestBody) GetScopedToAllDevices()(*bool) {
    return m.scopedToAllDevices
}
// GetState gets the state property value. Cancel or confirm scheduled retire 
func (m *DeviceCompliancePoliciesMicrosoftGraphSetScheduledRetireStateSetScheduledRetireStatePostRequestBody) GetState()(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ScheduledRetireState) {
    return m.state
}
// Serialize serializes information the current object
func (m *DeviceCompliancePoliciesMicrosoftGraphSetScheduledRetireStateSetScheduledRetireStatePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *DeviceCompliancePoliciesMicrosoftGraphSetScheduledRetireStateSetScheduledRetireStatePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetManagedDeviceIds sets the managedDeviceIds property value. The managedDeviceIds property
func (m *DeviceCompliancePoliciesMicrosoftGraphSetScheduledRetireStateSetScheduledRetireStatePostRequestBody) SetManagedDeviceIds(value []string)() {
    m.managedDeviceIds = value
}
// SetScopedToAllDevices sets the scopedToAllDevices property value. The scopedToAllDevices property
func (m *DeviceCompliancePoliciesMicrosoftGraphSetScheduledRetireStateSetScheduledRetireStatePostRequestBody) SetScopedToAllDevices(value *bool)() {
    m.scopedToAllDevices = value
}
// SetState sets the state property value. Cancel or confirm scheduled retire 
func (m *DeviceCompliancePoliciesMicrosoftGraphSetScheduledRetireStateSetScheduledRetireStatePostRequestBody) SetState(value *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ScheduledRetireState)() {
    m.state = value
}
