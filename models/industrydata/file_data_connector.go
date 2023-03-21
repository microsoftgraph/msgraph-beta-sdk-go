package industrydata

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// FileDataConnector 
type FileDataConnector struct {
    IndustryDataConnector
}
// NewFileDataConnector instantiates a new FileDataConnector and sets the default values.
func NewFileDataConnector()(*FileDataConnector) {
    m := &FileDataConnector{
        IndustryDataConnector: *NewIndustryDataConnector(),
    }
    odataTypeValue := "#microsoft.graph.industryData.fileDataConnector"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateFileDataConnectorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFileDataConnectorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.industryData.azureDataLakeConnector":
                        return NewAzureDataLakeConnector(), nil
                }
            }
        }
    }
    return NewFileDataConnector(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *FileDataConnector) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.IndustryDataConnector.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *FileDataConnector) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.IndustryDataConnector.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// FileDataConnectorable 
type FileDataConnectorable interface {
    IndustryDataConnectorable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
