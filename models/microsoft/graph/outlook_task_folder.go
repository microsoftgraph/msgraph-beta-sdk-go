package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// OutlookTaskFolder provides operations to manage the compliance singleton.
type OutlookTaskFolder struct {
    Entity
    // The version of the task folder.
    changeKey *string;
    // True if the folder is the default task folder.
    isDefaultFolder *bool;
    // The collection of multi-value extended properties defined for the task folder. Read-only. Nullable.
    multiValueExtendedProperties []MultiValueLegacyExtendedPropertyable;
    // The name of the task folder.
    name *string;
    // The unique GUID identifier for the task folder's parent group.
    parentGroupKey *string;
    // The collection of single-value extended properties defined for the task folder. Read-only. Nullable.
    singleValueExtendedProperties []SingleValueLegacyExtendedPropertyable;
    // The tasks in this task folder. Read-only. Nullable.
    tasks []OutlookTaskable;
}
// NewOutlookTaskFolder instantiates a new outlookTaskFolder and sets the default values.
func NewOutlookTaskFolder()(*OutlookTaskFolder) {
    m := &OutlookTaskFolder{
        Entity: *NewEntity(),
    }
    return m
}
// CreateOutlookTaskFolderFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOutlookTaskFolderFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewOutlookTaskFolder(), nil
}
// GetChangeKey gets the changeKey property value. The version of the task folder.
func (m *OutlookTaskFolder) GetChangeKey()(*string) {
    if m == nil {
        return nil
    } else {
        return m.changeKey
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OutlookTaskFolder) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["changeKey"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChangeKey(val)
        }
        return nil
    }
    res["isDefaultFolder"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDefaultFolder(val)
        }
        return nil
    }
    res["multiValueExtendedProperties"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMultiValueLegacyExtendedPropertyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MultiValueLegacyExtendedPropertyable, len(val))
            for i, v := range val {
                res[i] = v.(MultiValueLegacyExtendedPropertyable)
            }
            m.SetMultiValueExtendedProperties(res)
        }
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["parentGroupKey"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentGroupKey(val)
        }
        return nil
    }
    res["singleValueExtendedProperties"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSingleValueLegacyExtendedPropertyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SingleValueLegacyExtendedPropertyable, len(val))
            for i, v := range val {
                res[i] = v.(SingleValueLegacyExtendedPropertyable)
            }
            m.SetSingleValueExtendedProperties(res)
        }
        return nil
    }
    res["tasks"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOutlookTaskFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OutlookTaskable, len(val))
            for i, v := range val {
                res[i] = v.(OutlookTaskable)
            }
            m.SetTasks(res)
        }
        return nil
    }
    return res
}
// GetIsDefaultFolder gets the isDefaultFolder property value. True if the folder is the default task folder.
func (m *OutlookTaskFolder) GetIsDefaultFolder()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDefaultFolder
    }
}
// GetMultiValueExtendedProperties gets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the task folder. Read-only. Nullable.
func (m *OutlookTaskFolder) GetMultiValueExtendedProperties()([]MultiValueLegacyExtendedPropertyable) {
    if m == nil {
        return nil
    } else {
        return m.multiValueExtendedProperties
    }
}
// GetName gets the name property value. The name of the task folder.
func (m *OutlookTaskFolder) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetParentGroupKey gets the parentGroupKey property value. The unique GUID identifier for the task folder's parent group.
func (m *OutlookTaskFolder) GetParentGroupKey()(*string) {
    if m == nil {
        return nil
    } else {
        return m.parentGroupKey
    }
}
// GetSingleValueExtendedProperties gets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the task folder. Read-only. Nullable.
func (m *OutlookTaskFolder) GetSingleValueExtendedProperties()([]SingleValueLegacyExtendedPropertyable) {
    if m == nil {
        return nil
    } else {
        return m.singleValueExtendedProperties
    }
}
// GetTasks gets the tasks property value. The tasks in this task folder. Read-only. Nullable.
func (m *OutlookTaskFolder) GetTasks()([]OutlookTaskable) {
    if m == nil {
        return nil
    } else {
        return m.tasks
    }
}
func (m *OutlookTaskFolder) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *OutlookTaskFolder) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("changeKey", m.GetChangeKey())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isDefaultFolder", m.GetIsDefaultFolder())
        if err != nil {
            return err
        }
    }
    if m.GetMultiValueExtendedProperties() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMultiValueExtendedProperties()))
        for i, v := range m.GetMultiValueExtendedProperties() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("multiValueExtendedProperties", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("parentGroupKey", m.GetParentGroupKey())
        if err != nil {
            return err
        }
    }
    if m.GetSingleValueExtendedProperties() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSingleValueExtendedProperties()))
        for i, v := range m.GetSingleValueExtendedProperties() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("singleValueExtendedProperties", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTasks() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTasks()))
        for i, v := range m.GetTasks() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("tasks", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetChangeKey sets the changeKey property value. The version of the task folder.
func (m *OutlookTaskFolder) SetChangeKey(value *string)() {
    if m != nil {
        m.changeKey = value
    }
}
// SetIsDefaultFolder sets the isDefaultFolder property value. True if the folder is the default task folder.
func (m *OutlookTaskFolder) SetIsDefaultFolder(value *bool)() {
    if m != nil {
        m.isDefaultFolder = value
    }
}
// SetMultiValueExtendedProperties sets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the task folder. Read-only. Nullable.
func (m *OutlookTaskFolder) SetMultiValueExtendedProperties(value []MultiValueLegacyExtendedPropertyable)() {
    if m != nil {
        m.multiValueExtendedProperties = value
    }
}
// SetName sets the name property value. The name of the task folder.
func (m *OutlookTaskFolder) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetParentGroupKey sets the parentGroupKey property value. The unique GUID identifier for the task folder's parent group.
func (m *OutlookTaskFolder) SetParentGroupKey(value *string)() {
    if m != nil {
        m.parentGroupKey = value
    }
}
// SetSingleValueExtendedProperties sets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the task folder. Read-only. Nullable.
func (m *OutlookTaskFolder) SetSingleValueExtendedProperties(value []SingleValueLegacyExtendedPropertyable)() {
    if m != nil {
        m.singleValueExtendedProperties = value
    }
}
// SetTasks sets the tasks property value. The tasks in this task folder. Read-only. Nullable.
func (m *OutlookTaskFolder) SetTasks(value []OutlookTaskable)() {
    if m != nil {
        m.tasks = value
    }
}
