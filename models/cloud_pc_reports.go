package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPcReports 
type CloudPcReports struct {
    Entity
    // The exportJobs property
    exportJobs []CloudPcExportJobable
}
// NewCloudPcReports instantiates a new CloudPcReports and sets the default values.
func NewCloudPcReports()(*CloudPcReports) {
    m := &CloudPcReports{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.cloudPcReports";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateCloudPcReportsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPcReportsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcReports(), nil
}
// GetExportJobs gets the exportJobs property value. The exportJobs property
func (m *CloudPcReports) GetExportJobs()([]CloudPcExportJobable) {
    return m.exportJobs
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcReports) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["exportJobs"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateCloudPcExportJobFromDiscriminatorValue , m.SetExportJobs)
    return res
}
// Serialize serializes information the current object
func (m *CloudPcReports) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetExportJobs() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetExportJobs())
        err = writer.WriteCollectionOfObjectValues("exportJobs", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetExportJobs sets the exportJobs property value. The exportJobs property
func (m *CloudPcReports) SetExportJobs(value []CloudPcExportJobable)() {
    m.exportJobs = value
}
