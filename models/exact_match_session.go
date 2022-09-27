package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ExactMatchSession 
type ExactMatchSession struct {
    ExactMatchSessionBase
    // The checksum property
    checksum *string
    // The dataUploadURI property
    dataUploadURI *string
    // The fields property
    fields []string
    // The fileName property
    fileName *string
    // The rowsPerBlock property
    rowsPerBlock *int32
    // The salt property
    salt *string
    // The uploadAgent property
    uploadAgent ExactMatchUploadAgentable
    // The uploadAgentId property
    uploadAgentId *string
}
// NewExactMatchSession instantiates a new ExactMatchSession and sets the default values.
func NewExactMatchSession()(*ExactMatchSession) {
    m := &ExactMatchSession{
        ExactMatchSessionBase: *NewExactMatchSessionBase(),
    }
    odataTypeValue := "#microsoft.graph.exactMatchSession";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateExactMatchSessionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateExactMatchSessionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExactMatchSession(), nil
}
// GetChecksum gets the checksum property value. The checksum property
func (m *ExactMatchSession) GetChecksum()(*string) {
    return m.checksum
}
// GetDataUploadURI gets the dataUploadURI property value. The dataUploadURI property
func (m *ExactMatchSession) GetDataUploadURI()(*string) {
    return m.dataUploadURI
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ExactMatchSession) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ExactMatchSessionBase.GetFieldDeserializers()
    res["checksum"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetChecksum)
    res["dataUploadURI"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDataUploadURI)
    res["fields"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetFields)
    res["fileName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetFileName)
    res["rowsPerBlock"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetRowsPerBlock)
    res["salt"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSalt)
    res["uploadAgent"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateExactMatchUploadAgentFromDiscriminatorValue , m.SetUploadAgent)
    res["uploadAgentId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUploadAgentId)
    return res
}
// GetFields gets the fields property value. The fields property
func (m *ExactMatchSession) GetFields()([]string) {
    return m.fields
}
// GetFileName gets the fileName property value. The fileName property
func (m *ExactMatchSession) GetFileName()(*string) {
    return m.fileName
}
// GetRowsPerBlock gets the rowsPerBlock property value. The rowsPerBlock property
func (m *ExactMatchSession) GetRowsPerBlock()(*int32) {
    return m.rowsPerBlock
}
// GetSalt gets the salt property value. The salt property
func (m *ExactMatchSession) GetSalt()(*string) {
    return m.salt
}
// GetUploadAgent gets the uploadAgent property value. The uploadAgent property
func (m *ExactMatchSession) GetUploadAgent()(ExactMatchUploadAgentable) {
    return m.uploadAgent
}
// GetUploadAgentId gets the uploadAgentId property value. The uploadAgentId property
func (m *ExactMatchSession) GetUploadAgentId()(*string) {
    return m.uploadAgentId
}
// Serialize serializes information the current object
func (m *ExactMatchSession) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ExactMatchSessionBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("checksum", m.GetChecksum())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("dataUploadURI", m.GetDataUploadURI())
        if err != nil {
            return err
        }
    }
    if m.GetFields() != nil {
        err = writer.WriteCollectionOfStringValues("fields", m.GetFields())
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
        err = writer.WriteInt32Value("rowsPerBlock", m.GetRowsPerBlock())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("salt", m.GetSalt())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("uploadAgent", m.GetUploadAgent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("uploadAgentId", m.GetUploadAgentId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetChecksum sets the checksum property value. The checksum property
func (m *ExactMatchSession) SetChecksum(value *string)() {
    m.checksum = value
}
// SetDataUploadURI sets the dataUploadURI property value. The dataUploadURI property
func (m *ExactMatchSession) SetDataUploadURI(value *string)() {
    m.dataUploadURI = value
}
// SetFields sets the fields property value. The fields property
func (m *ExactMatchSession) SetFields(value []string)() {
    m.fields = value
}
// SetFileName sets the fileName property value. The fileName property
func (m *ExactMatchSession) SetFileName(value *string)() {
    m.fileName = value
}
// SetRowsPerBlock sets the rowsPerBlock property value. The rowsPerBlock property
func (m *ExactMatchSession) SetRowsPerBlock(value *int32)() {
    m.rowsPerBlock = value
}
// SetSalt sets the salt property value. The salt property
func (m *ExactMatchSession) SetSalt(value *string)() {
    m.salt = value
}
// SetUploadAgent sets the uploadAgent property value. The uploadAgent property
func (m *ExactMatchSession) SetUploadAgent(value ExactMatchUploadAgentable)() {
    m.uploadAgent = value
}
// SetUploadAgentId sets the uploadAgentId property value. The uploadAgentId property
func (m *ExactMatchSession) SetUploadAgentId(value *string)() {
    m.uploadAgentId = value
}
