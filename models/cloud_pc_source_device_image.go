package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPcSourceDeviceImage 
type CloudPcSourceDeviceImage struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The display name for the source image.
    displayName *string
    // The ID of the source image.
    id *string
    // The OdataType property
    odataType *string
    // The display name of subscription that hosts the source image.
    subscriptionDisplayName *string
    // The ID of subscription that hosts the source image.
    subscriptionId *string
}
// NewCloudPcSourceDeviceImage instantiates a new cloudPcSourceDeviceImage and sets the default values.
func NewCloudPcSourceDeviceImage()(*CloudPcSourceDeviceImage) {
    m := &CloudPcSourceDeviceImage{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCloudPcSourceDeviceImageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPcSourceDeviceImageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcSourceDeviceImage(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CloudPcSourceDeviceImage) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDisplayName gets the displayName property value. The display name for the source image.
func (m *CloudPcSourceDeviceImage) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcSourceDeviceImage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["id"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetId)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["subscriptionDisplayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSubscriptionDisplayName)
    res["subscriptionId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSubscriptionId)
    return res
}
// GetId gets the id property value. The ID of the source image.
func (m *CloudPcSourceDeviceImage) GetId()(*string) {
    return m.id
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *CloudPcSourceDeviceImage) GetOdataType()(*string) {
    return m.odataType
}
// GetSubscriptionDisplayName gets the subscriptionDisplayName property value. The display name of subscription that hosts the source image.
func (m *CloudPcSourceDeviceImage) GetSubscriptionDisplayName()(*string) {
    return m.subscriptionDisplayName
}
// GetSubscriptionId gets the subscriptionId property value. The ID of subscription that hosts the source image.
func (m *CloudPcSourceDeviceImage) GetSubscriptionId()(*string) {
    return m.subscriptionId
}
// Serialize serializes information the current object
func (m *CloudPcSourceDeviceImage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
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
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("subscriptionDisplayName", m.GetSubscriptionDisplayName())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CloudPcSourceDeviceImage) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDisplayName sets the displayName property value. The display name for the source image.
func (m *CloudPcSourceDeviceImage) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetId sets the id property value. The ID of the source image.
func (m *CloudPcSourceDeviceImage) SetId(value *string)() {
    m.id = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CloudPcSourceDeviceImage) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSubscriptionDisplayName sets the subscriptionDisplayName property value. The display name of subscription that hosts the source image.
func (m *CloudPcSourceDeviceImage) SetSubscriptionDisplayName(value *string)() {
    m.subscriptionDisplayName = value
}
// SetSubscriptionId sets the subscriptionId property value. The ID of subscription that hosts the source image.
func (m *CloudPcSourceDeviceImage) SetSubscriptionId(value *string)() {
    m.subscriptionId = value
}
