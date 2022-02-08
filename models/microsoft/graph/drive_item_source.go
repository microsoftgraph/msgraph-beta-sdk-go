package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DriveItemSource 
type DriveItemSource struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Enumeration value that indicates the source application where the file was created.
    application *DriveItemSourceApplication;
    // The external identifier for the drive item from the source.
    externalId *string;
}
// NewDriveItemSource instantiates a new driveItemSource and sets the default values.
func NewDriveItemSource()(*DriveItemSource) {
    m := &DriveItemSource{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DriveItemSource) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetApplication gets the application property value. Enumeration value that indicates the source application where the file was created.
func (m *DriveItemSource) GetApplication()(*DriveItemSourceApplication) {
    if m == nil {
        return nil
    } else {
        return m.application
    }
}
// GetExternalId gets the externalId property value. The external identifier for the drive item from the source.
func (m *DriveItemSource) GetExternalId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DriveItemSource) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["application"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDriveItemSourceApplication)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplication(val.(*DriveItemSourceApplication))
        }
        return nil
    }
    res["externalId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalId(val)
        }
        return nil
    }
    return res
}
func (m *DriveItemSource) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DriveItemSource) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetApplication() != nil {
        cast := (*m.GetApplication()).String()
        err := writer.WriteStringValue("application", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("externalId", m.GetExternalId())
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
func (m *DriveItemSource) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetApplication sets the application property value. Enumeration value that indicates the source application where the file was created.
func (m *DriveItemSource) SetApplication(value *DriveItemSourceApplication)() {
    if m != nil {
        m.application = value
    }
}
// SetExternalId sets the externalId property value. The external identifier for the drive item from the source.
func (m *DriveItemSource) SetExternalId(value *string)() {
    if m != nil {
        m.externalId = value
    }
}
