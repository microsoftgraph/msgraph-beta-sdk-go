package getportalnotifications

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i2edb12705e6a63a8a0fb3f8c7a11f4ab12f4be764e61fa1094f401595fb171bf "github.com/microsoftgraph/msgraph-beta-sdk-go/models/devicemanagement"
)

// GetPortalNotificationsResponse provides operations to call the getPortalNotifications method.
type GetPortalNotificationsResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The value property
    value []i2edb12705e6a63a8a0fb3f8c7a11f4ab12f4be764e61fa1094f401595fb171bf.PortalNotificationable
}
// NewGetPortalNotificationsResponse instantiates a new getPortalNotificationsResponse and sets the default values.
func NewGetPortalNotificationsResponse()(*GetPortalNotificationsResponse) {
    m := &GetPortalNotificationsResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateGetPortalNotificationsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGetPortalNotificationsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGetPortalNotificationsResponse(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GetPortalNotificationsResponse) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GetPortalNotificationsResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i2edb12705e6a63a8a0fb3f8c7a11f4ab12f4be764e61fa1094f401595fb171bf.CreatePortalNotificationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i2edb12705e6a63a8a0fb3f8c7a11f4ab12f4be764e61fa1094f401595fb171bf.PortalNotificationable, len(val))
            for i, v := range val {
                res[i] = v.(i2edb12705e6a63a8a0fb3f8c7a11f4ab12f4be764e61fa1094f401595fb171bf.PortalNotificationable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *GetPortalNotificationsResponse) GetValue()([]i2edb12705e6a63a8a0fb3f8c7a11f4ab12f4be764e61fa1094f401595fb171bf.PortalNotificationable) {
    return m.value
}
// Serialize serializes information the current object
func (m *GetPortalNotificationsResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("value", cast)
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
func (m *GetPortalNotificationsResponse) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetValue sets the value property value. The value property
func (m *GetPortalNotificationsResponse) SetValue(value []i2edb12705e6a63a8a0fb3f8c7a11f4ab12f4be764e61fa1094f401595fb171bf.PortalNotificationable)() {
    m.value = value
}
