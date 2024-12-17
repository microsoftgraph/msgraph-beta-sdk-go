package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

type DataDiscoveryReport struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewDataDiscoveryReport instantiates a new DataDiscoveryReport and sets the default values.
func NewDataDiscoveryReport()(*DataDiscoveryReport) {
    m := &DataDiscoveryReport{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateDataDiscoveryReportFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDataDiscoveryReportFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDataDiscoveryReport(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DataDiscoveryReport) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["uploadedStreams"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCloudAppDiscoveryReportFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudAppDiscoveryReportable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CloudAppDiscoveryReportable)
                }
            }
            m.SetUploadedStreams(res)
        }
        return nil
    }
    return res
}
// GetUploadedStreams gets the uploadedStreams property value. A collection of streams available for generating cloud discovery report.
// returns a []CloudAppDiscoveryReportable when successful
func (m *DataDiscoveryReport) GetUploadedStreams()([]CloudAppDiscoveryReportable) {
    val, err := m.GetBackingStore().Get("uploadedStreams")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CloudAppDiscoveryReportable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DataDiscoveryReport) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetUploadedStreams() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUploadedStreams()))
        for i, v := range m.GetUploadedStreams() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("uploadedStreams", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetUploadedStreams sets the uploadedStreams property value. A collection of streams available for generating cloud discovery report.
func (m *DataDiscoveryReport) SetUploadedStreams(value []CloudAppDiscoveryReportable)() {
    err := m.GetBackingStore().Set("uploadedStreams", value)
    if err != nil {
        panic(err)
    }
}
type DataDiscoveryReportable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetUploadedStreams()([]CloudAppDiscoveryReportable)
    SetUploadedStreams(value []CloudAppDiscoveryReportable)()
}
