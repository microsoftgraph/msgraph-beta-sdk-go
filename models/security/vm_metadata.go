package security

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VmMetadata 
type VmMetadata struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The cloudProvider property
    cloudProvider *VmCloudProvider
    // The OdataType property
    odataType *string
    // Unique identifier of the Azure resource.
    resourceId *string
    // Unique identifier of the Azure subscription the customer tenant belongs to.
    subscriptionId *string
    // Unique identifier of the virtual machine instance.
    vmId *string
}
// NewVmMetadata instantiates a new vmMetadata and sets the default values.
func NewVmMetadata()(*VmMetadata) {
    m := &VmMetadata{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateVmMetadataFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVmMetadataFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVmMetadata(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *VmMetadata) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCloudProvider gets the cloudProvider property value. The cloudProvider property
func (m *VmMetadata) GetCloudProvider()(*VmCloudProvider) {
    return m.cloudProvider
}
// GetFieldDeserializers the deserialization information for the current model
func (m *VmMetadata) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["cloudProvider"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseVmCloudProvider , m.SetCloudProvider)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["resourceId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetResourceId)
    res["subscriptionId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSubscriptionId)
    res["vmId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetVmId)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *VmMetadata) GetOdataType()(*string) {
    return m.odataType
}
// GetResourceId gets the resourceId property value. Unique identifier of the Azure resource.
func (m *VmMetadata) GetResourceId()(*string) {
    return m.resourceId
}
// GetSubscriptionId gets the subscriptionId property value. Unique identifier of the Azure subscription the customer tenant belongs to.
func (m *VmMetadata) GetSubscriptionId()(*string) {
    return m.subscriptionId
}
// GetVmId gets the vmId property value. Unique identifier of the virtual machine instance.
func (m *VmMetadata) GetVmId()(*string) {
    return m.vmId
}
// Serialize serializes information the current object
func (m *VmMetadata) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCloudProvider() != nil {
        cast := (*m.GetCloudProvider()).String()
        err := writer.WriteStringValue("cloudProvider", &cast)
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
        err := writer.WriteStringValue("resourceId", m.GetResourceId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("subscriptionId", m.GetSubscriptionId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("vmId", m.GetVmId())
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
func (m *VmMetadata) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCloudProvider sets the cloudProvider property value. The cloudProvider property
func (m *VmMetadata) SetCloudProvider(value *VmCloudProvider)() {
    m.cloudProvider = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *VmMetadata) SetOdataType(value *string)() {
    m.odataType = value
}
// SetResourceId sets the resourceId property value. Unique identifier of the Azure resource.
func (m *VmMetadata) SetResourceId(value *string)() {
    m.resourceId = value
}
// SetSubscriptionId sets the subscriptionId property value. Unique identifier of the Azure subscription the customer tenant belongs to.
func (m *VmMetadata) SetSubscriptionId(value *string)() {
    m.subscriptionId = value
}
// SetVmId sets the vmId property value. Unique identifier of the virtual machine instance.
func (m *VmMetadata) SetVmId(value *string)() {
    m.vmId = value
}
