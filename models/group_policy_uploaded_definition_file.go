package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupPolicyUploadedDefinitionFile 
type GroupPolicyUploadedDefinitionFile struct {
    GroupPolicyDefinitionFile
}
// NewGroupPolicyUploadedDefinitionFile instantiates a new GroupPolicyUploadedDefinitionFile and sets the default values.
func NewGroupPolicyUploadedDefinitionFile()(*GroupPolicyUploadedDefinitionFile) {
    m := &GroupPolicyUploadedDefinitionFile{
        GroupPolicyDefinitionFile: *NewGroupPolicyDefinitionFile(),
    }
    odataTypeValue := "#microsoft.graph.groupPolicyUploadedDefinitionFile"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateGroupPolicyUploadedDefinitionFileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupPolicyUploadedDefinitionFileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupPolicyUploadedDefinitionFile(), nil
}
// GetContent gets the content property value. The contents of the uploaded ADMX file.
func (m *GroupPolicyUploadedDefinitionFile) GetContent()([]byte) {
    val, err := m.GetBackingStore().Get("content")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]byte)
    }
    return nil
}
// GetDefaultLanguageCode gets the defaultLanguageCode property value. The default language of the uploaded ADMX file.
func (m *GroupPolicyUploadedDefinitionFile) GetDefaultLanguageCode()(*string) {
    val, err := m.GetBackingStore().Get("defaultLanguageCode")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupPolicyUploadedDefinitionFile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.GroupPolicyDefinitionFile.GetFieldDeserializers()
    res["content"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContent(val)
        }
        return nil
    }
    res["defaultLanguageCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultLanguageCode(val)
        }
        return nil
    }
    res["groupPolicyOperations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGroupPolicyOperationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GroupPolicyOperationable, len(val))
            for i, v := range val {
                res[i] = v.(GroupPolicyOperationable)
            }
            m.SetGroupPolicyOperations(res)
        }
        return nil
    }
    res["groupPolicyUploadedLanguageFiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGroupPolicyUploadedLanguageFileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GroupPolicyUploadedLanguageFileable, len(val))
            for i, v := range val {
                res[i] = v.(GroupPolicyUploadedLanguageFileable)
            }
            m.SetGroupPolicyUploadedLanguageFiles(res)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseGroupPolicyUploadedDefinitionFileStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*GroupPolicyUploadedDefinitionFileStatus))
        }
        return nil
    }
    res["uploadDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetGroupPolicyOperations gets the groupPolicyOperations property value. The list of operations on the uploaded ADMX file.
func (m *GroupPolicyUploadedDefinitionFile) GetGroupPolicyOperations()([]GroupPolicyOperationable) {
    val, err := m.GetBackingStore().Get("groupPolicyOperations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]GroupPolicyOperationable)
    }
    return nil
}
// GetGroupPolicyUploadedLanguageFiles gets the groupPolicyUploadedLanguageFiles property value. The list of ADML files associated with the uploaded ADMX file.
func (m *GroupPolicyUploadedDefinitionFile) GetGroupPolicyUploadedLanguageFiles()([]GroupPolicyUploadedLanguageFileable) {
    val, err := m.GetBackingStore().Get("groupPolicyUploadedLanguageFiles")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]GroupPolicyUploadedLanguageFileable)
    }
    return nil
}
// GetStatus gets the status property value. Type of Group Policy uploaded definition file status.
func (m *GroupPolicyUploadedDefinitionFile) GetStatus()(*GroupPolicyUploadedDefinitionFileStatus) {
    val, err := m.GetBackingStore().Get("status")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*GroupPolicyUploadedDefinitionFileStatus)
    }
    return nil
}
// GetUploadDateTime gets the uploadDateTime property value. The uploaded time of the uploaded ADMX file.
func (m *GroupPolicyUploadedDefinitionFile) GetUploadDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("uploadDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetGroupPolicyOperations()))
        for i, v := range m.GetGroupPolicyOperations() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("groupPolicyOperations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetGroupPolicyUploadedLanguageFiles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetGroupPolicyUploadedLanguageFiles()))
        for i, v := range m.GetGroupPolicyUploadedLanguageFiles() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
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
    err := m.GetBackingStore().Set("content", value)
    if err != nil {
        panic(err)
    }
}
// SetDefaultLanguageCode sets the defaultLanguageCode property value. The default language of the uploaded ADMX file.
func (m *GroupPolicyUploadedDefinitionFile) SetDefaultLanguageCode(value *string)() {
    err := m.GetBackingStore().Set("defaultLanguageCode", value)
    if err != nil {
        panic(err)
    }
}
// SetGroupPolicyOperations sets the groupPolicyOperations property value. The list of operations on the uploaded ADMX file.
func (m *GroupPolicyUploadedDefinitionFile) SetGroupPolicyOperations(value []GroupPolicyOperationable)() {
    err := m.GetBackingStore().Set("groupPolicyOperations", value)
    if err != nil {
        panic(err)
    }
}
// SetGroupPolicyUploadedLanguageFiles sets the groupPolicyUploadedLanguageFiles property value. The list of ADML files associated with the uploaded ADMX file.
func (m *GroupPolicyUploadedDefinitionFile) SetGroupPolicyUploadedLanguageFiles(value []GroupPolicyUploadedLanguageFileable)() {
    err := m.GetBackingStore().Set("groupPolicyUploadedLanguageFiles", value)
    if err != nil {
        panic(err)
    }
}
// SetStatus sets the status property value. Type of Group Policy uploaded definition file status.
func (m *GroupPolicyUploadedDefinitionFile) SetStatus(value *GroupPolicyUploadedDefinitionFileStatus)() {
    err := m.GetBackingStore().Set("status", value)
    if err != nil {
        panic(err)
    }
}
// SetUploadDateTime sets the uploadDateTime property value. The uploaded time of the uploaded ADMX file.
func (m *GroupPolicyUploadedDefinitionFile) SetUploadDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("uploadDateTime", value)
    if err != nil {
        panic(err)
    }
}
// GroupPolicyUploadedDefinitionFileable 
type GroupPolicyUploadedDefinitionFileable interface {
    GroupPolicyDefinitionFileable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContent()([]byte)
    GetDefaultLanguageCode()(*string)
    GetGroupPolicyOperations()([]GroupPolicyOperationable)
    GetGroupPolicyUploadedLanguageFiles()([]GroupPolicyUploadedLanguageFileable)
    GetStatus()(*GroupPolicyUploadedDefinitionFileStatus)
    GetUploadDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetContent(value []byte)()
    SetDefaultLanguageCode(value *string)()
    SetGroupPolicyOperations(value []GroupPolicyOperationable)()
    SetGroupPolicyUploadedLanguageFiles(value []GroupPolicyUploadedLanguageFileable)()
    SetStatus(value *GroupPolicyUploadedDefinitionFileStatus)()
    SetUploadDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
