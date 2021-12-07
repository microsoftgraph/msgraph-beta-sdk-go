package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DirectorySizeQuota 
type DirectorySizeQuota struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Total amount of the directory quota.
    total *int32;
    // Used amount of the directory quota.
    used *int32;
}
// NewDirectorySizeQuota instantiates a new directorySizeQuota and sets the default values.
func NewDirectorySizeQuota()(*DirectorySizeQuota) {
    m := &DirectorySizeQuota{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DirectorySizeQuota) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetTotal gets the total property value. Total amount of the directory quota.
func (m *DirectorySizeQuota) GetTotal()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.total
    }
}
// GetUsed gets the used property value. Used amount of the directory quota.
func (m *DirectorySizeQuota) GetUsed()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.used
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DirectorySizeQuota) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["total"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotal(val)
        }
        return nil
    }
    res["used"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsed(val)
        }
        return nil
    }
    return res
}
func (m *DirectorySizeQuota) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DirectorySizeQuota) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("total", m.GetTotal())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("used", m.GetUsed())
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
func (m *DirectorySizeQuota) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetTotal sets the total property value. Total amount of the directory quota.
func (m *DirectorySizeQuota) SetTotal(value *int32)() {
    if m != nil {
        m.total = value
    }
}
// SetUsed sets the used property value. Used amount of the directory quota.
func (m *DirectorySizeQuota) SetUsed(value *int32)() {
    if m != nil {
        m.used = value
    }
}
