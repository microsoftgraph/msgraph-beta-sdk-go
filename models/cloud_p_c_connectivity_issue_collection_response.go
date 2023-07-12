package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPCConnectivityIssueCollectionResponse 
type CloudPCConnectivityIssueCollectionResponse struct {
    BaseCollectionPaginationCountResponse
}
// NewCloudPCConnectivityIssueCollectionResponse instantiates a new cloudPCConnectivityIssueCollectionResponse and sets the default values.
func NewCloudPCConnectivityIssueCollectionResponse()(*CloudPCConnectivityIssueCollectionResponse) {
    m := &CloudPCConnectivityIssueCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateCloudPCConnectivityIssueCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPCConnectivityIssueCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPCConnectivityIssueCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPCConnectivityIssueCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCloudPCConnectivityIssueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPCConnectivityIssueable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CloudPCConnectivityIssueable)
                }
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *CloudPCConnectivityIssueCollectionResponse) GetValue()([]CloudPCConnectivityIssueable) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CloudPCConnectivityIssueable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CloudPCConnectivityIssueCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *CloudPCConnectivityIssueCollectionResponse) SetValue(value []CloudPCConnectivityIssueable)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
// CloudPCConnectivityIssueCollectionResponseable 
type CloudPCConnectivityIssueCollectionResponseable interface {
    BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]CloudPCConnectivityIssueable)
    SetValue(value []CloudPCConnectivityIssueable)()
}
