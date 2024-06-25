package industrydata

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AzureDataLakeConnector struct {
    FileDataConnector
}
// NewAzureDataLakeConnector instantiates a new AzureDataLakeConnector and sets the default values.
func NewAzureDataLakeConnector()(*AzureDataLakeConnector) {
    m := &AzureDataLakeConnector{
        FileDataConnector: *NewFileDataConnector(),
    }
    odataTypeValue := "#microsoft.graph.industryData.azureDataLakeConnector"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAzureDataLakeConnectorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAzureDataLakeConnectorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAzureDataLakeConnector(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AzureDataLakeConnector) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.FileDataConnector.GetFieldDeserializers()
    res["fileFormat"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateFileFormatReferenceValueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileFormat(val.(FileFormatReferenceValueable))
        }
        return nil
    }
    return res
}
// GetFileFormat gets the fileFormat property value. The file format that external systems can upload using this connector.
// returns a FileFormatReferenceValueable when successful
func (m *AzureDataLakeConnector) GetFileFormat()(FileFormatReferenceValueable) {
    val, err := m.GetBackingStore().Get("fileFormat")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(FileFormatReferenceValueable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AzureDataLakeConnector) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.FileDataConnector.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("fileFormat", m.GetFileFormat())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFileFormat sets the fileFormat property value. The file format that external systems can upload using this connector.
func (m *AzureDataLakeConnector) SetFileFormat(value FileFormatReferenceValueable)() {
    err := m.GetBackingStore().Set("fileFormat", value)
    if err != nil {
        panic(err)
    }
}
type AzureDataLakeConnectorable interface {
    FileDataConnectorable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFileFormat()(FileFormatReferenceValueable)
    SetFileFormat(value FileFormatReferenceValueable)()
}
