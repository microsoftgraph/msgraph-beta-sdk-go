package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementTemplatesImportOffice365DeviceConfigurationPoliciesResponse provides operations to call the importOffice365DeviceConfigurationPolicies method.
type DeviceManagementTemplatesImportOffice365DeviceConfigurationPoliciesResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []DeviceManagementIntentable
}
// NewDeviceManagementTemplatesImportOffice365DeviceConfigurationPoliciesResponse instantiates a new DeviceManagementTemplatesImportOffice365DeviceConfigurationPoliciesResponse and sets the default values.
func NewDeviceManagementTemplatesImportOffice365DeviceConfigurationPoliciesResponse()(*DeviceManagementTemplatesImportOffice365DeviceConfigurationPoliciesResponse) {
    m := &DeviceManagementTemplatesImportOffice365DeviceConfigurationPoliciesResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateDeviceManagementTemplatesImportOffice365DeviceConfigurationPoliciesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementTemplatesImportOffice365DeviceConfigurationPoliciesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementTemplatesImportOffice365DeviceConfigurationPoliciesResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementTemplatesImportOffice365DeviceConfigurationPoliciesResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementIntentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementIntentable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementIntentable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *DeviceManagementTemplatesImportOffice365DeviceConfigurationPoliciesResponse) GetValue()([]DeviceManagementIntentable) {
    return m.value
}
// Serialize serializes information the current object
func (m *DeviceManagementTemplatesImportOffice365DeviceConfigurationPoliciesResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *DeviceManagementTemplatesImportOffice365DeviceConfigurationPoliciesResponse) SetValue(value []DeviceManagementIntentable)() {
    m.value = value
}
