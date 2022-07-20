package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPcSubscription 
type CloudPcSubscription struct {
    Entity
    // The ID of the subscription.
    subscriptionId *string
    // The name of the subscription.
    subscriptionName *string
}
// NewCloudPcSubscription instantiates a new cloudPcSubscription and sets the default values.
func NewCloudPcSubscription()(*CloudPcSubscription) {
    m := &CloudPcSubscription{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.cloudPcSubscription";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateCloudPcSubscriptionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPcSubscriptionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcSubscription(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcSubscription) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["subscriptionName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubscriptionName(val)
        }
        return nil
    }
    return res
}
// GetSubscriptionId gets the subscriptionId property value. The ID of the subscription.
func (m *CloudPcSubscription) GetSubscriptionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subscriptionId
    }
}
// GetSubscriptionName gets the subscriptionName property value. The name of the subscription.
func (m *CloudPcSubscription) GetSubscriptionName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subscriptionName
    }
}
// Serialize serializes information the current object
func (m *CloudPcSubscription) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("subscriptionId", m.GetSubscriptionId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("subscriptionName", m.GetSubscriptionName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSubscriptionId sets the subscriptionId property value. The ID of the subscription.
func (m *CloudPcSubscription) SetSubscriptionId(value *string)() {
    if m != nil {
        m.subscriptionId = value
    }
}
// SetSubscriptionName sets the subscriptionName property value. The name of the subscription.
func (m *CloudPcSubscription) SetSubscriptionName(value *string)() {
    if m != nil {
        m.subscriptionName = value
    }
}
