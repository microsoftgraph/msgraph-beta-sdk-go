package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type SynchronizationRule struct {
    additionalData map[string]interface{};
    editable *bool;
    id *string;
    metadata []StringKeyStringValuePair;
    name *string;
    objectMappings []ObjectMapping;
    priority *int32;
    sourceDirectoryName *string;
    targetDirectoryName *string;
}
func NewSynchronizationRule()(*SynchronizationRule) {
    m := &SynchronizationRule{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *SynchronizationRule) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *SynchronizationRule) GetEditable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.editable
    }
}
func (m *SynchronizationRule) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
func (m *SynchronizationRule) GetMetadata()([]StringKeyStringValuePair) {
    if m == nil {
        return nil
    } else {
        return m.metadata
    }
}
func (m *SynchronizationRule) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
func (m *SynchronizationRule) GetObjectMappings()([]ObjectMapping) {
    if m == nil {
        return nil
    } else {
        return m.objectMappings
    }
}
func (m *SynchronizationRule) GetPriority()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.priority
    }
}
func (m *SynchronizationRule) GetSourceDirectoryName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sourceDirectoryName
    }
}
func (m *SynchronizationRule) GetTargetDirectoryName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetDirectoryName
    }
}
func (m *SynchronizationRule) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["editable"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetEditable(val)
        return nil
    }
    res["id"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetId(val)
        return nil
    }
    res["metadata"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewStringKeyStringValuePair() })
        if err != nil {
            return err
        }
        res := make([]StringKeyStringValuePair, len(val))
        for i, v := range val {
            res[i] = *(v.(*StringKeyStringValuePair))
        }
        m.SetMetadata(res)
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetName(val)
        return nil
    }
    res["objectMappings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewObjectMapping() })
        if err != nil {
            return err
        }
        res := make([]ObjectMapping, len(val))
        for i, v := range val {
            res[i] = *(v.(*ObjectMapping))
        }
        m.SetObjectMappings(res)
        return nil
    }
    res["priority"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetPriority(val)
        return nil
    }
    res["sourceDirectoryName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSourceDirectoryName(val)
        return nil
    }
    res["targetDirectoryName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTargetDirectoryName(val)
        return nil
    }
    return res
}
func (m *SynchronizationRule) IsNil()(bool) {
    return m == nil
}
func (m *SynchronizationRule) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("editable", m.GetEditable())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMetadata()))
        for i, v := range m.GetMetadata() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("metadata", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetObjectMappings()))
        for i, v := range m.GetObjectMappings() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("objectMappings", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("priority", m.GetPriority())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sourceDirectoryName", m.GetSourceDirectoryName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("targetDirectoryName", m.GetTargetDirectoryName())
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
func (m *SynchronizationRule) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *SynchronizationRule) SetEditable(value *bool)() {
    m.editable = value
}
func (m *SynchronizationRule) SetId(value *string)() {
    m.id = value
}
func (m *SynchronizationRule) SetMetadata(value []StringKeyStringValuePair)() {
    m.metadata = value
}
func (m *SynchronizationRule) SetName(value *string)() {
    m.name = value
}
func (m *SynchronizationRule) SetObjectMappings(value []ObjectMapping)() {
    m.objectMappings = value
}
func (m *SynchronizationRule) SetPriority(value *int32)() {
    m.priority = value
}
func (m *SynchronizationRule) SetSourceDirectoryName(value *string)() {
    m.sourceDirectoryName = value
}
func (m *SynchronizationRule) SetTargetDirectoryName(value *string)() {
    m.targetDirectoryName = value
}
