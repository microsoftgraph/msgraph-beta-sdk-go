package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Win32LobAppRequirement base class to detect a Win32 App
type Win32LobAppRequirement struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The detection value
    detectionValue *string
    // Contains properties for detection operator.
    operator *Win32LobAppDetectionOperator
    // The type property
    type_escaped *string
}
// NewWin32LobAppRequirement instantiates a new win32LobAppRequirement and sets the default values.
func NewWin32LobAppRequirement()(*Win32LobAppRequirement) {
    m := &Win32LobAppRequirement{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    typeValue := "#microsoft.graph.win32LobAppRequirement";
    m.SetType(&typeValue);
    return m
}
// CreateWin32LobAppRequirementFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWin32LobAppRequirementFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.win32LobAppFileSystemRequirement":
                        return NewWin32LobAppFileSystemRequirement(), nil
                    case "#microsoft.graph.win32LobAppPowerShellScriptRequirement":
                        return NewWin32LobAppPowerShellScriptRequirement(), nil
                    case "#microsoft.graph.win32LobAppRegistryRequirement":
                        return NewWin32LobAppRegistryRequirement(), nil
                }
            }
        }
    }
    return NewWin32LobAppRequirement(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Win32LobAppRequirement) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDetectionValue gets the detectionValue property value. The detection value
func (m *Win32LobAppRequirement) GetDetectionValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.detectionValue
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Win32LobAppRequirement) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["detectionValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetectionValue(val)
        }
        return nil
    }
    res["operator"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWin32LobAppDetectionOperator)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperator(val.(*Win32LobAppDetectionOperator))
        }
        return nil
    }
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
// GetOperator gets the operator property value. Contains properties for detection operator.
func (m *Win32LobAppRequirement) GetOperator()(*Win32LobAppDetectionOperator) {
    if m == nil {
        return nil
    } else {
        return m.operator
    }
}
// GetType gets the type property value. The type property
func (m *Win32LobAppRequirement) GetType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// Serialize serializes information the current object
func (m *Win32LobAppRequirement) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("detectionValue", m.GetDetectionValue())
        if err != nil {
            return err
        }
    }
    if m.GetOperator() != nil {
        cast := (*m.GetOperator()).String()
        err := writer.WriteStringValue("operator", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type", m.GetType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Win32LobAppRequirement) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDetectionValue sets the detectionValue property value. The detection value
func (m *Win32LobAppRequirement) SetDetectionValue(value *string)() {
    if m != nil {
        m.detectionValue = value
    }
}
// SetOperator sets the operator property value. Contains properties for detection operator.
func (m *Win32LobAppRequirement) SetOperator(value *Win32LobAppDetectionOperator)() {
    if m != nil {
        m.operator = value
    }
}
// SetType sets the type property value. The type property
func (m *Win32LobAppRequirement) SetType(value *string)() {
    if m != nil {
        m.type_escaped = value
    }
}
