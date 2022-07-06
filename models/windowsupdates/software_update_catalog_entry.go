package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SoftwareUpdateCatalogEntry 
type SoftwareUpdateCatalogEntry struct {
    CatalogEntry
    // The type property
    type_escaped *string
}
// NewSoftwareUpdateCatalogEntry instantiates a new SoftwareUpdateCatalogEntry and sets the default values.
func NewSoftwareUpdateCatalogEntry()(*SoftwareUpdateCatalogEntry) {
    m := &SoftwareUpdateCatalogEntry{
        CatalogEntry: *NewCatalogEntry(),
    }
    return m
}
// CreateSoftwareUpdateCatalogEntryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSoftwareUpdateCatalogEntryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                mappingStr := *mappingValue
                switch mappingStr {
                    case "#microsoft.graph.windowsUpdates.featureUpdateCatalogEntry":
                        return NewFeatureUpdateCatalogEntry(), nil
                    case "#microsoft.graph.windowsUpdates.qualityUpdateCatalogEntry":
                        return NewQualityUpdateCatalogEntry(), nil
                }
            }
        }
    }
    return NewSoftwareUpdateCatalogEntry(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SoftwareUpdateCatalogEntry) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CatalogEntry.GetFieldDeserializers()
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val)
        }
        return nil
    }
    return res
}
// GetType gets the type property value. The type property
func (m *SoftwareUpdateCatalogEntry) GetType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// Serialize serializes information the current object
func (m *SoftwareUpdateCatalogEntry) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CatalogEntry.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("type", m.GetType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetType sets the type property value. The type property
func (m *SoftwareUpdateCatalogEntry) SetType(value *string)() {
    if m != nil {
        m.type_escaped = value
    }
}
