package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type GroupPolicyUploadedDefinitionFile struct {
    GroupPolicyDefinitionFile
    // The contents of the uploaded ADMX file.
    content []byte;
    // The default language of the uploaded ADMX file.
    defaultLanguageCode *string;
    // The file name of the uploaded ADML file.
    fileName *string;
    // The list of operations on the uploaded ADMX file.
    groupPolicyOperations []GroupPolicyOperation;
    // The list of ADML files associated with the uploaded ADMX file.
    groupPolicyUploadedLanguageFiles []GroupPolicyUploadedLanguageFile;
    // The upload status of the uploaded ADMX file. Possible values are: none, uploadInProgress, available, assigned, removalInProgress, uploadFailed, removalFailed.
    status *GroupPolicyUploadedDefinitionFileStatus;
    // The uploaded time of the uploaded ADMX file.
    uploadDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
}
// Instantiates a new groupPolicyUploadedDefinitionFile and sets the default values.
func NewGroupPolicyUploadedDefinitionFile()(*GroupPolicyUploadedDefinitionFile) {
    m := &GroupPolicyUploadedDefinitionFile{
        GroupPolicyDefinitionFile: *NewGroupPolicyDefinitionFile(),
    }
    return m
}
// Gets the content property value. The contents of the uploaded ADMX file.
func (m *GroupPolicyUploadedDefinitionFile) GetContent()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.content
    }
}
// Gets the defaultLanguageCode property value. The default language of the uploaded ADMX file.
func (m *GroupPolicyUploadedDefinitionFile) GetDefaultLanguageCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.defaultLanguageCode
    }
}
// Gets the fileName property value. The file name of the uploaded ADML file.
func (m *GroupPolicyUploadedDefinitionFile) GetFileName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.fileName
    }
}
// Gets the groupPolicyOperations property value. The list of operations on the uploaded ADMX file.
func (m *GroupPolicyUploadedDefinitionFile) GetGroupPolicyOperations()([]GroupPolicyOperation) {
    if m == nil {
        return nil
    } else {
        return m.groupPolicyOperations
    }
}
// Gets the groupPolicyUploadedLanguageFiles property value. The list of ADML files associated with the uploaded ADMX file.
func (m *GroupPolicyUploadedDefinitionFile) GetGroupPolicyUploadedLanguageFiles()([]GroupPolicyUploadedLanguageFile) {
    if m == nil {
        return nil
    } else {
        return m.groupPolicyUploadedLanguageFiles
    }
}
// Gets the status property value. The upload status of the uploaded ADMX file. Possible values are: none, uploadInProgress, available, assigned, removalInProgress, uploadFailed, removalFailed.
func (m *GroupPolicyUploadedDefinitionFile) GetStatus()(*GroupPolicyUploadedDefinitionFileStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// Gets the uploadDateTime property value. The uploaded time of the uploaded ADMX file.
func (m *GroupPolicyUploadedDefinitionFile) GetUploadDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.uploadDateTime
    }
}
// The deserialization information for the current model
func (m *GroupPolicyUploadedDefinitionFile) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.GroupPolicyDefinitionFile.GetFieldDeserializers()
    res["content"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContent(val)
        }
        return nil
    }
    res["defaultLanguageCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultLanguageCode(val)
        }
        return nil
    }
    res["fileName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileName(val)
        }
        return nil
    }
    res["groupPolicyOperations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGroupPolicyOperation() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GroupPolicyOperation, len(val))
            for i, v := range val {
                res[i] = *(v.(*GroupPolicyOperation))
            }
            m.SetGroupPolicyOperations(res)
        }
        return nil
    }
    res["groupPolicyUploadedLanguageFiles"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGroupPolicyUploadedLanguageFile() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GroupPolicyUploadedLanguageFile, len(val))
            for i, v := range val {
                res[i] = *(v.(*GroupPolicyUploadedLanguageFile))
            }
            m.SetGroupPolicyUploadedLanguageFiles(res)
        }
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseGroupPolicyUploadedDefinitionFileStatus)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(GroupPolicyUploadedDefinitionFileStatus)
            m.SetStatus(&cast)
        }
        return nil
    }
    res["uploadDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUploadDateTime(val)
        }
        return nil
    }
    return res
}
func (m *GroupPolicyUploadedDefinitionFile) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the content property value. The contents of the uploaded ADMX file.
// Parameters:
//  - value : Value to set for the content property.
func (m *GroupPolicyUploadedDefinitionFile) SetContent(value []byte)() {
    m.content = value
}
// Sets the defaultLanguageCode property value. The default language of the uploaded ADMX file.
// Parameters:
//  - value : Value to set for the defaultLanguageCode property.
func (m *GroupPolicyUploadedDefinitionFile) SetDefaultLanguageCode(value *string)() {
    m.defaultLanguageCode = value
}
// Sets the fileName property value. The file name of the uploaded ADML file.
// Parameters:
//  - value : Value to set for the fileName property.
func (m *GroupPolicyUploadedDefinitionFile) SetFileName(value *string)() {
    m.fileName = value
}
// Sets the groupPolicyOperations property value. The list of operations on the uploaded ADMX file.
// Parameters:
//  - value : Value to set for the groupPolicyOperations property.
func (m *GroupPolicyUploadedDefinitionFile) SetGroupPolicyOperations(value []GroupPolicyOperation)() {
    m.groupPolicyOperations = value
}
// Sets the groupPolicyUploadedLanguageFiles property value. The list of ADML files associated with the uploaded ADMX file.
// Parameters:
//  - value : Value to set for the groupPolicyUploadedLanguageFiles property.
func (m *GroupPolicyUploadedDefinitionFile) SetGroupPolicyUploadedLanguageFiles(value []GroupPolicyUploadedLanguageFile)() {
    m.groupPolicyUploadedLanguageFiles = value
}
// Sets the status property value. The upload status of the uploaded ADMX file. Possible values are: none, uploadInProgress, available, assigned, removalInProgress, uploadFailed, removalFailed.
// Parameters:
//  - value : Value to set for the status property.
func (m *GroupPolicyUploadedDefinitionFile) SetStatus(value *GroupPolicyUploadedDefinitionFileStatus)() {
    m.status = value
}
// Sets the uploadDateTime property value. The uploaded time of the uploaded ADMX file.
// Parameters:
//  - value : Value to set for the uploadDateTime property.
func (m *GroupPolicyUploadedDefinitionFile) SetUploadDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.uploadDateTime = value
}
