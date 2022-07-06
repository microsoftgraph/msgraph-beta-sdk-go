package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DlpEvaluationInput 
type DlpEvaluationInput struct {
    // The accessScope property
    accessScope *AccessScope
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The currentLabel property
    currentLabel CurrentLabelable
    // The discoveredSensitiveTypes property
    discoveredSensitiveTypes []DiscoveredSensitiveTypeable
    // The type property
    type_escaped *string
}
// NewDlpEvaluationInput instantiates a new dlpEvaluationInput and sets the default values.
func NewDlpEvaluationInput()(*DlpEvaluationInput) {
    m := &DlpEvaluationInput{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDlpEvaluationInputFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDlpEvaluationInputFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.dlpEvaluationWindowsDevicesInput":
                        return NewDlpEvaluationWindowsDevicesInput(), nil
                }
            }
        }
    }
    return NewDlpEvaluationInput(), nil
}
// GetAccessScope gets the accessScope property value. The accessScope property
func (m *DlpEvaluationInput) GetAccessScope()(*AccessScope) {
    if m == nil {
        return nil
    } else {
        return m.accessScope
    }
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DlpEvaluationInput) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCurrentLabel gets the currentLabel property value. The currentLabel property
func (m *DlpEvaluationInput) GetCurrentLabel()(CurrentLabelable) {
    if m == nil {
        return nil
    } else {
        return m.currentLabel
    }
}
// GetDiscoveredSensitiveTypes gets the discoveredSensitiveTypes property value. The discoveredSensitiveTypes property
func (m *DlpEvaluationInput) GetDiscoveredSensitiveTypes()([]DiscoveredSensitiveTypeable) {
    if m == nil {
        return nil
    } else {
        return m.discoveredSensitiveTypes
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DlpEvaluationInput) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["accessScope"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAccessScope)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessScope(val.(*AccessScope))
        }
        return nil
    }
    res["currentLabel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCurrentLabelFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentLabel(val.(CurrentLabelable))
        }
        return nil
    }
    res["discoveredSensitiveTypes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDiscoveredSensitiveTypeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DiscoveredSensitiveTypeable, len(val))
            for i, v := range val {
                res[i] = v.(DiscoveredSensitiveTypeable)
            }
            m.SetDiscoveredSensitiveTypes(res)
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
// GetType gets the type property value. The type property
func (m *DlpEvaluationInput) GetType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// Serialize serializes information the current object
func (m *DlpEvaluationInput) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAccessScope() != nil {
        cast := (*m.GetAccessScope()).String()
        err := writer.WriteStringValue("accessScope", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("currentLabel", m.GetCurrentLabel())
        if err != nil {
            return err
        }
    }
    if m.GetDiscoveredSensitiveTypes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDiscoveredSensitiveTypes()))
        for i, v := range m.GetDiscoveredSensitiveTypes() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("discoveredSensitiveTypes", cast)
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
// SetAccessScope sets the accessScope property value. The accessScope property
func (m *DlpEvaluationInput) SetAccessScope(value *AccessScope)() {
    if m != nil {
        m.accessScope = value
    }
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DlpEvaluationInput) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCurrentLabel sets the currentLabel property value. The currentLabel property
func (m *DlpEvaluationInput) SetCurrentLabel(value CurrentLabelable)() {
    if m != nil {
        m.currentLabel = value
    }
}
// SetDiscoveredSensitiveTypes sets the discoveredSensitiveTypes property value. The discoveredSensitiveTypes property
func (m *DlpEvaluationInput) SetDiscoveredSensitiveTypes(value []DiscoveredSensitiveTypeable)() {
    if m != nil {
        m.discoveredSensitiveTypes = value
    }
}
// SetType sets the type property value. The type property
func (m *DlpEvaluationInput) SetType(value *string)() {
    if m != nil {
        m.type_escaped = value
    }
}
