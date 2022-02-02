package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamworkCameraConfiguration 
type TeamworkCameraConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    cameras []TeamworkPeripheral;
    // The configuration for the content camera.
    contentCameraConfiguration *TeamworkContentCameraConfiguration;
    // 
    defaultContentCamera *TeamworkPeripheral;
}
// NewTeamworkCameraConfiguration instantiates a new teamworkCameraConfiguration and sets the default values.
func NewTeamworkCameraConfiguration()(*TeamworkCameraConfiguration) {
    m := &TeamworkCameraConfiguration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkCameraConfiguration) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCameras gets the cameras property value. 
func (m *TeamworkCameraConfiguration) GetCameras()([]TeamworkPeripheral) {
    if m == nil {
        return nil
    } else {
        return m.cameras
    }
}
// GetContentCameraConfiguration gets the contentCameraConfiguration property value. The configuration for the content camera.
func (m *TeamworkCameraConfiguration) GetContentCameraConfiguration()(*TeamworkContentCameraConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.contentCameraConfiguration
    }
}
// GetDefaultContentCamera gets the defaultContentCamera property value. 
func (m *TeamworkCameraConfiguration) GetDefaultContentCamera()(*TeamworkPeripheral) {
    if m == nil {
        return nil
    } else {
        return m.defaultContentCamera
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkCameraConfiguration) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["cameras"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkPeripheral() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TeamworkPeripheral, len(val))
            for i, v := range val {
                res[i] = *(v.(*TeamworkPeripheral))
            }
            m.SetCameras(res)
        }
        return nil
    }
    res["contentCameraConfiguration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkContentCameraConfiguration() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentCameraConfiguration(val.(*TeamworkContentCameraConfiguration))
        }
        return nil
    }
    res["defaultContentCamera"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkPeripheral() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultContentCamera(val.(*TeamworkPeripheral))
        }
        return nil
    }
    return res
}
func (m *TeamworkCameraConfiguration) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TeamworkCameraConfiguration) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetCameras() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCameras()))
        for i, v := range m.GetCameras() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("cameras", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("contentCameraConfiguration", m.GetContentCameraConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("defaultContentCamera", m.GetDefaultContentCamera())
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
func (m *TeamworkCameraConfiguration) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCameras sets the cameras property value. 
func (m *TeamworkCameraConfiguration) SetCameras(value []TeamworkPeripheral)() {
    if m != nil {
        m.cameras = value
    }
}
// SetContentCameraConfiguration sets the contentCameraConfiguration property value. The configuration for the content camera.
func (m *TeamworkCameraConfiguration) SetContentCameraConfiguration(value *TeamworkContentCameraConfiguration)() {
    if m != nil {
        m.contentCameraConfiguration = value
    }
}
// SetDefaultContentCamera sets the defaultContentCamera property value. 
func (m *TeamworkCameraConfiguration) SetDefaultContentCamera(value *TeamworkPeripheral)() {
    if m != nil {
        m.defaultContentCamera = value
    }
}
