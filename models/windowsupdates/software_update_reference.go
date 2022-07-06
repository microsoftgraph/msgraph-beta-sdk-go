package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SoftwareUpdateReference 
type SoftwareUpdateReference struct {
    DeployableContent
    // The type property
    type_escaped *string
}
// NewSoftwareUpdateReference instantiates a new SoftwareUpdateReference and sets the default values.
func NewSoftwareUpdateReference()(*SoftwareUpdateReference) {
    m := &SoftwareUpdateReference{
        DeployableContent: *NewDeployableContent(),
    }
    return m
}
// CreateSoftwareUpdateReferenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSoftwareUpdateReferenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.windowsUpdates.windowsUpdateReference":
                        return NewWindowsUpdateReference(), nil
                }
            }
        }
    }
    return NewSoftwareUpdateReference(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SoftwareUpdateReference) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeployableContent.GetFieldDeserializers()
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
func (m *SoftwareUpdateReference) GetType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// Serialize serializes information the current object
func (m *SoftwareUpdateReference) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeployableContent.Serialize(writer)
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
func (m *SoftwareUpdateReference) SetType(value *string)() {
    if m != nil {
        m.type_escaped = value
    }
}
