package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupPolicyUploadedDefinitionFile 
type GroupPolicyUploadedDefinitionFile struct {
    GroupPolicyDefinitionFile
    // The contents of the uploaded ADMX file.
    content []byte
    // The default language of the uploaded ADMX file.
    defaultLanguageCode *string
    // The list of operations on the uploaded ADMX file.
    groupPolicyOperations []GroupPolicyOperationable
    // The list of ADML files associated with the uploaded ADMX file.
    groupPolicyUploadedLanguageFiles []GroupPolicyUploadedLanguageFileable
    // Type of Group Policy uploaded definition file status.
    status *GroupPolicyUploadedDefinitionFileStatus
    // The uploaded time of the uploaded ADMX file.
    uploadDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewGroupPolicyUploadedDefinitionFile instantiates a new GroupPolicyUploadedDefinitionFile and sets the default values.
func NewGroupPolicyUploadedDefinitionFile()(*GroupPolicyUploadedDefinitionFile) {
    m := &GroupPolicyUploadedDefinitionFile{
        GroupPolicyDefinitionFile: *NewGroupPolicyDefinitionFile(),
    }
    odataTypeValue := "#microsoft.graph.groupPolicyUploadedDefinitionFile";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateGroupPolicyUploadedDefinitionFileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupPolicyUploadedDefinitionFileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupPolicyUploadedDefinitionFile(), nil
}
// GetContent gets the content property value. The contents of the uploaded ADMX file.
func (m *GroupPolicyUploadedDefinitionFile) GetContent()([]byte) {
    return m.content
}
// GetDefaultLanguageCode gets the defaultLanguageCode property value. The default language of the uploaded ADMX file.
func (m *GroupPolicyUploadedDefinitionFile) GetDefaultLanguageCode()(*string) {
    return m.defaultLanguageCode
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupPolicyUploadedDefinitionFile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.GroupPolicyDefinitionFile.GetFieldDeserializers()
    res["content"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetByteArrayValue(m.SetContent)
    res["defaultLanguageCode"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDefaultLanguageCode)
    res["groupPolicyOperations"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateGroupPolicyOperationFromDiscriminatorValue , m.SetGroupPolicyOperations)
    res["groupPolicyUploadedLanguageFiles"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateGroupPolicyUploadedLanguageFileFromDiscriminatorValue , m.SetGroupPolicyUploadedLanguageFiles)
    res["status"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseGroupPolicyUploadedDefinitionFileStatus , m.SetStatus)
    res["uploadDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetUploadDateTime)
    return res
}
// GetGroupPolicyOperations gets the groupPolicyOperations property value. The list of operations on the uploaded ADMX file.
func (m *GroupPolicyUploadedDefinitionFile) GetGroupPolicyOperations()([]GroupPolicyOperationable) {
    return m.groupPolicyOperations
}
// GetGroupPolicyUploadedLanguageFiles gets the groupPolicyUploadedLanguageFiles property value. The list of ADML files associated with the uploaded ADMX file.
func (m *GroupPolicyUploadedDefinitionFile) GetGroupPolicyUploadedLanguageFiles()([]GroupPolicyUploadedLanguageFileable) {
    return m.groupPolicyUploadedLanguageFiles
}
// GetStatus gets the status property value. Type of Group Policy uploaded definition file status.
func (m *GroupPolicyUploadedDefinitionFile) GetStatus()(*GroupPolicyUploadedDefinitionFileStatus) {
    return m.status
}
// GetUploadDateTime gets the uploadDateTime property value. The uploaded time of the uploaded ADMX file.
func (m *GroupPolicyUploadedDefinitionFile) GetUploadDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.uploadDateTime
}
// Serialize serializes information the current object
func (m *GroupPolicyUploadedDefinitionFile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetGroupPolicyOperations() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetGroupPolicyOperations())
        err = writer.WriteCollectionOfObjectValues("groupPolicyOperations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetGroupPolicyUploadedLanguageFiles() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetGroupPolicyUploadedLanguageFiles())
        err = writer.WriteCollectionOfObjectValues("groupPolicyUploadedLanguageFiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
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
// SetContent sets the content property value. The contents of the uploaded ADMX file.
func (m *GroupPolicyUploadedDefinitionFile) SetContent(value []byte)() {
    m.content = value
}
// SetDefaultLanguageCode sets the defaultLanguageCode property value. The default language of the uploaded ADMX file.
func (m *GroupPolicyUploadedDefinitionFile) SetDefaultLanguageCode(value *string)() {
    m.defaultLanguageCode = value
}
// SetGroupPolicyOperations sets the groupPolicyOperations property value. The list of operations on the uploaded ADMX file.
func (m *GroupPolicyUploadedDefinitionFile) SetGroupPolicyOperations(value []GroupPolicyOperationable)() {
    m.groupPolicyOperations = value
}
// SetGroupPolicyUploadedLanguageFiles sets the groupPolicyUploadedLanguageFiles property value. The list of ADML files associated with the uploaded ADMX file.
func (m *GroupPolicyUploadedDefinitionFile) SetGroupPolicyUploadedLanguageFiles(value []GroupPolicyUploadedLanguageFileable)() {
    m.groupPolicyUploadedLanguageFiles = value
}
// SetStatus sets the status property value. Type of Group Policy uploaded definition file status.
func (m *GroupPolicyUploadedDefinitionFile) SetStatus(value *GroupPolicyUploadedDefinitionFileStatus)() {
    m.status = value
}
// SetUploadDateTime sets the uploadDateTime property value. The uploaded time of the uploaded ADMX file.
func (m *GroupPolicyUploadedDefinitionFile) SetUploadDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.uploadDateTime = value
}
