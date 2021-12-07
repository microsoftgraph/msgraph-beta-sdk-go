package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Compliance 
type Compliance struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    ediscovery *Ediscoveryroot;
}
// NewCompliance instantiates a new Compliance and sets the default values.
func NewCompliance()(*Compliance) {
    m := &Compliance{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Compliance) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetEdiscovery gets the ediscovery property value. 
func (m *Compliance) GetEdiscovery()(*Ediscoveryroot) {
    if m == nil {
        return nil
    } else {
        return m.ediscovery
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Compliance) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["ediscovery"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEdiscoveryroot() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEdiscovery(val.(*Ediscoveryroot))
        }
        return nil
    }
    return res
}
func (m *Compliance) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Compliance) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("ediscovery", m.GetEdiscovery())
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
func (m *Compliance) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetEdiscovery sets the ediscovery property value. 
func (m *Compliance) SetEdiscovery(value *Ediscoveryroot)() {
    if m != nil {
        m.ediscovery = value
    }
}
