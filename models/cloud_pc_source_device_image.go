package models

import (
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
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDisplayName gets the displayName property value. The display name for the source image.
func (m *CloudPcSourceDeviceImage) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcSourceDeviceImage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
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
    res["subscriptionDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubscriptionDisplayName(val)
        }
        return nil
    }
    res["subscriptionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubscriptionId(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The ID of the source image.
func (m *CloudPcSourceDeviceImage) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
// GetSubscriptionDisplayName gets the subscriptionDisplayName property value. The display name of subscription that hosts the source image.
func (m *CloudPcSourceDeviceImage) GetSubscriptionDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subscriptionDisplayName
    }
}
// GetSubscriptionId gets the subscriptionId property value. The ID of subscription that hosts the source image.
func (m *CloudPcSourceDeviceImage) GetSubscriptionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subscriptionId
    }
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
    if m != nil {
        m.additionalData = value
    }
}
// SetDisplayName sets the displayName property value. The display name for the source image.
func (m *CloudPcSourceDeviceImage) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetId sets the id property value. The ID of the source image.
func (m *CloudPcSourceDeviceImage) SetId(value *string)() {
    if m != nil {
        m.id = value
    }
}
// SetSubscriptionDisplayName sets the subscriptionDisplayName property value. The display name of subscription that hosts the source image.
func (m *CloudPcSourceDeviceImage) SetSubscriptionDisplayName(value *string)() {
    if m != nil {
        m.subscriptionDisplayName = value
    }
}
// SetSubscriptionId sets the subscriptionId property value. The ID of subscription that hosts the source image.
func (m *CloudPcSourceDeviceImage) SetSubscriptionId(value *string)() {
    if m != nil {
        m.subscriptionId = value
    }
}
