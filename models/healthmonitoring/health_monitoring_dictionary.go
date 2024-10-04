package healthmonitoring

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type HealthMonitoringDictionary struct {
    Dictionary
}
// NewHealthMonitoringDictionary instantiates a new HealthMonitoringDictionary and sets the default values.
func NewHealthMonitoringDictionary()(*HealthMonitoringDictionary) {
    m := &HealthMonitoringDictionary{
        Dictionary: *NewDictionary(),
    }
    return m
}
// CreateHealthMonitoringDictionaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateHealthMonitoringDictionaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.healthMonitoring.documentation":
                        return NewDocumentation(), nil
                    case "#microsoft.graph.healthMonitoring.signals":
                        return NewSignals(), nil
                    case "#microsoft.graph.healthMonitoring.supportingData":
                        return NewSupportingData(), nil
                }
            }
        }
    }
    return NewHealthMonitoringDictionary(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *HealthMonitoringDictionary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Dictionary.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *HealthMonitoringDictionary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Dictionary.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type HealthMonitoringDictionaryable interface {
    Dictionaryable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
