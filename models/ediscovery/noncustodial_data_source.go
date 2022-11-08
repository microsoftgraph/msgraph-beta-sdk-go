package ediscovery

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// NoncustodialDataSource provides operations to manage the compliance singleton.
type NoncustodialDataSource struct {
    DataSourceContainer
    // Indicates if hold is applied to non-custodial data source (such as mailbox or site).
    applyHoldToSource *bool
    // User source or SharePoint site data source as non-custodial data source.
    dataSource DataSourceable
}
// NewNoncustodialDataSource instantiates a new noncustodialDataSource and sets the default values.
func NewNoncustodialDataSource()(*NoncustodialDataSource) {
    m := &NoncustodialDataSource{
        DataSourceContainer: *NewDataSourceContainer(),
    }
    odataTypeValue := "#microsoft.graph.ediscovery.noncustodialDataSource";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateNoncustodialDataSourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateNoncustodialDataSourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNoncustodialDataSource(), nil
}
// GetApplyHoldToSource gets the applyHoldToSource property value. Indicates if hold is applied to non-custodial data source (such as mailbox or site).
func (m *NoncustodialDataSource) GetApplyHoldToSource()(*bool) {
    return m.applyHoldToSource
}
// GetDataSource gets the dataSource property value. User source or SharePoint site data source as non-custodial data source.
func (m *NoncustodialDataSource) GetDataSource()(DataSourceable) {
    return m.dataSource
}
// GetFieldDeserializers the deserialization information for the current model
func (m *NoncustodialDataSource) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DataSourceContainer.GetFieldDeserializers()
    res["applyHoldToSource"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetApplyHoldToSource)
    res["dataSource"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateDataSourceFromDiscriminatorValue , m.SetDataSource)
    return res
}
// Serialize serializes information the current object
func (m *NoncustodialDataSource) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DataSourceContainer.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("applyHoldToSource", m.GetApplyHoldToSource())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("dataSource", m.GetDataSource())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApplyHoldToSource sets the applyHoldToSource property value. Indicates if hold is applied to non-custodial data source (such as mailbox or site).
func (m *NoncustodialDataSource) SetApplyHoldToSource(value *bool)() {
    m.applyHoldToSource = value
}
// SetDataSource sets the dataSource property value. User source or SharePoint site data source as non-custodial data source.
func (m *NoncustodialDataSource) SetDataSource(value DataSourceable)() {
    m.dataSource = value
}
