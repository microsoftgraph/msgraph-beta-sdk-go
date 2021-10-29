package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DirectorySizeQuota struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Total amount of the directory quota.
    total *int32;
    // Used amount of the directory quota.
    used *int32;
}
// Instantiates a new directorySizeQuota and sets the default values.
func NewDirectorySizeQuota()(*DirectorySizeQuota) {
    m := &DirectorySizeQuota{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DirectorySizeQuota) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the total property value. Total amount of the directory quota.
func (m *DirectorySizeQuota) GetTotal()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.total
    }
}
// Gets the used property value. Used amount of the directory quota.
func (m *DirectorySizeQuota) GetUsed()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.used
    }
}
// The deserialization information for the current model
func (m *DirectorySizeQuota) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["total"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetTotal(val)
        return nil
    }
    res["used"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetUsed(val)
        return nil
    }
    return res
}
func (m *DirectorySizeQuota) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *DirectorySizeQuota) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the total property value. Total amount of the directory quota.
// Parameters:
//  - value : Value to set for the total property.
func (m *DirectorySizeQuota) SetTotal(value *int32)() {
    m.total = value
}
// Sets the used property value. Used amount of the directory quota.
// Parameters:
//  - value : Value to set for the used property.
func (m *DirectorySizeQuota) SetUsed(value *int32)() {
    m.used = value
}
