package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CloudPcBulkSetReviewStatus struct {
    CloudPcBulkAction
}
// NewCloudPcBulkSetReviewStatus instantiates a new CloudPcBulkSetReviewStatus and sets the default values.
func NewCloudPcBulkSetReviewStatus()(*CloudPcBulkSetReviewStatus) {
    m := &CloudPcBulkSetReviewStatus{
        CloudPcBulkAction: *NewCloudPcBulkAction(),
    }
    odataTypeValue := "#microsoft.graph.cloudPcBulkSetReviewStatus"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateCloudPcBulkSetReviewStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCloudPcBulkSetReviewStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcBulkSetReviewStatus(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CloudPcBulkSetReviewStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CloudPcBulkAction.GetFieldDeserializers()
    res["reviewStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCloudPcReviewStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReviewStatus(val.(CloudPcReviewStatusable))
        }
        return nil
    }
    return res
}
// GetReviewStatus gets the reviewStatus property value. The reviewStatus property
// returns a CloudPcReviewStatusable when successful
func (m *CloudPcBulkSetReviewStatus) GetReviewStatus()(CloudPcReviewStatusable) {
    val, err := m.GetBackingStore().Get("reviewStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(CloudPcReviewStatusable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CloudPcBulkSetReviewStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CloudPcBulkAction.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("reviewStatus", m.GetReviewStatus())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetReviewStatus sets the reviewStatus property value. The reviewStatus property
func (m *CloudPcBulkSetReviewStatus) SetReviewStatus(value CloudPcReviewStatusable)() {
    err := m.GetBackingStore().Set("reviewStatus", value)
    if err != nil {
        panic(err)
    }
}
type CloudPcBulkSetReviewStatusable interface {
    CloudPcBulkActionable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetReviewStatus()(CloudPcReviewStatusable)
    SetReviewStatus(value CloudPcReviewStatusable)()
}
