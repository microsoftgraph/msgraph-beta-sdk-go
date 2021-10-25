package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type GroupPolicyUploadedDefinitionFile struct {
    GroupPolicyDefinitionFile
    content []byte;
    defaultLanguageCode *string;
    fileName *string;
    groupPolicyOperations []GroupPolicyOperation;
    groupPolicyUploadedLanguageFiles []GroupPolicyUploadedLanguageFile;
    status *GroupPolicyUploadedDefinitionFileStatus;
    uploadDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
}
func NewGroupPolicyUploadedDefinitionFile()(*GroupPolicyUploadedDefinitionFile) {
    m := &GroupPolicyUploadedDefinitionFile{
        GroupPolicyDefinitionFile: *NewGroupPolicyDefinitionFile(),
    }
    return m
}
func (m *GroupPolicyUploadedDefinitionFile) GetContent()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.content
    }
}
func (m *GroupPolicyUploadedDefinitionFile) GetDefaultLanguageCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.defaultLanguageCode
    }
}
func (m *GroupPolicyUploadedDefinitionFile) GetFileName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.fileName
    }
}
func (m *GroupPolicyUploadedDefinitionFile) GetGroupPolicyOperations()([]GroupPolicyOperation) {
    if m == nil {
        return nil
    } else {
        return m.groupPolicyOperations
    }
}
func (m *GroupPolicyUploadedDefinitionFile) GetGroupPolicyUploadedLanguageFiles()([]GroupPolicyUploadedLanguageFile) {
    if m == nil {
        return nil
    } else {
        return m.groupPolicyUploadedLanguageFiles
    }
}
func (m *GroupPolicyUploadedDefinitionFile) GetStatus()(*GroupPolicyUploadedDefinitionFileStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
func (m *GroupPolicyUploadedDefinitionFile) GetUploadDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.uploadDateTime
    }
}
func (m *GroupPolicyUploadedDefinitionFile) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.GroupPolicyDefinitionFile.GetFieldDeserializers()
    res["content"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        m.SetContent(val)
        return nil
    }
    res["defaultLanguageCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDefaultLanguageCode(val)
        return nil
    }
    res["fileName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetFileName(val)
        return nil
    }
    res["groupPolicyOperations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGroupPolicyOperation() })
        if err != nil {
            return err
        }
        res := make([]GroupPolicyOperation, len(val))
        for i, v := range val {
            res[i] = *(v.(*GroupPolicyOperation))
        }
        m.SetGroupPolicyOperations(res)
        return nil
    }
    res["groupPolicyUploadedLanguageFiles"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGroupPolicyUploadedLanguageFile() })
        if err != nil {
            return err
        }
        res := make([]GroupPolicyUploadedLanguageFile, len(val))
        for i, v := range val {
            res[i] = *(v.(*GroupPolicyUploadedLanguageFile))
        }
        m.SetGroupPolicyUploadedLanguageFiles(res)
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseGroupPolicyUploadedDefinitionFileStatus)
        if err != nil {
            return err
        }
        cast := val.(GroupPolicyUploadedDefinitionFileStatus)
        m.SetStatus(&cast)
        return nil
    }
    res["uploadDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetUploadDateTime(val)
        return nil
    }
    return res
}
func (m *GroupPolicyUploadedDefinitionFile) IsNil()(bool) {
    return m == nil
}
func (m *GroupPolicyUploadedDefinitionFile) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.GroupPolicyDefinitionFile.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteByteArrayValue("content", m.GetContent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("defaultLanguageCode", m.GetDefaultLanguageCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("fileName", m.GetFileName())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetGroupPolicyOperations()))
        for i, v := range m.GetGroupPolicyOperations() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("groupPolicyOperations", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetGroupPolicyUploadedLanguageFiles()))
        for i, v := range m.GetGroupPolicyUploadedLanguageFiles() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("groupPolicyUploadedLanguageFiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := m.GetStatus().String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("uploadDateTime", m.GetUploadDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *GroupPolicyUploadedDefinitionFile) SetContent(value []byte)() {
    m.content = value
}
func (m *GroupPolicyUploadedDefinitionFile) SetDefaultLanguageCode(value *string)() {
    m.defaultLanguageCode = value
}
func (m *GroupPolicyUploadedDefinitionFile) SetFileName(value *string)() {
    m.fileName = value
}
func (m *GroupPolicyUploadedDefinitionFile) SetGroupPolicyOperations(value []GroupPolicyOperation)() {
    m.groupPolicyOperations = value
}
func (m *GroupPolicyUploadedDefinitionFile) SetGroupPolicyUploadedLanguageFiles(value []GroupPolicyUploadedLanguageFile)() {
    m.groupPolicyUploadedLanguageFiles = value
}
func (m *GroupPolicyUploadedDefinitionFile) SetStatus(value *GroupPolicyUploadedDefinitionFileStatus)() {
    m.status = value
}
func (m *GroupPolicyUploadedDefinitionFile) SetUploadDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.uploadDateTime = value
}
