package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DriveItemSource struct {
    additionalData map[string]interface{};
    application *DriveItemSourceApplication;
    externalId *string;
}
func NewDriveItemSource()(*DriveItemSource) {
    m := &DriveItemSource{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *DriveItemSource) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *DriveItemSource) GetApplication()(*DriveItemSourceApplication) {
    if m == nil {
        return nil
    } else {
        return m.application
    }
}
func (m *DriveItemSource) GetExternalId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalId
    }
}
func (m *DriveItemSource) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["application"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDriveItemSourceApplication)
        if err != nil {
            return err
        }
        cast := val.(DriveItemSourceApplication)
        m.SetApplication(&cast)
        return nil
    }
    res["externalId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetExternalId(val)
        return nil
    }
    return res
}
func (m *DriveItemSource) IsNil()(bool) {
    return m == nil
}
func (m *DriveItemSource) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetApplication() != nil {
        cast := m.GetApplication().String()
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
func (m *DriveItemSource) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *DriveItemSource) SetApplication(value *DriveItemSourceApplication)() {
    m.application = value
}
func (m *DriveItemSource) SetExternalId(value *string)() {
    m.externalId = value
}
