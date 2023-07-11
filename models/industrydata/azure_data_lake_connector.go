package industrydata

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AzureDataLakeConnector 
type AzureDataLakeConnector struct {
    FileDataConnector
}
// NewAzureDataLakeConnector instantiates a new azureDataLakeConnector and sets the default values.
func NewAzureDataLakeConnector()(*AzureDataLakeConnector) {
    m := &AzureDataLakeConnector{
        FileDataConnector: *NewFileDataConnector(),
    }
    odataTypeValue := "#microsoft.graph.industryData.azureDataLakeConnector"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAzureDataLakeConnectorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAzureDataLakeConnectorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAzureDataLakeConnector(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AzureDataLakeConnector) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.FileDataConnector.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *AzureDataLakeConnector) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.FileDataConnector.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// AzureDataLakeConnectorable 
type AzureDataLakeConnectorable interface {
    FileDataConnectorable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
