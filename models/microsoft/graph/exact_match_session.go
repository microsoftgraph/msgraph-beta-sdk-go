package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ExactMatchSession struct {
    ExactMatchSessionBase
    // 
    checksum *string;
    // 
    dataUploadURI *string;
    // 
    fields []string;
    // 
    fileName *string;
    // 
    rowsPerBlock *int32;
    // 
    salt *string;
    // 
    uploadAgent *ExactMatchUploadAgent;
    // 
    uploadAgentId *string;
}
// Instantiates a new exactMatchSession and sets the default values.
func NewExactMatchSession()(*ExactMatchSession) {
    m := &ExactMatchSession{
        ExactMatchSessionBase: *NewExactMatchSessionBase(),
    }
    return m
}
// Gets the checksum property value. 
func (m *ExactMatchSession) GetChecksum()(*string) {
    if m == nil {
        return nil
    } else {
        return m.checksum
    }
}
// Gets the dataUploadURI property value. 
func (m *ExactMatchSession) GetDataUploadURI()(*string) {
    if m == nil {
        return nil
    } else {
        return m.dataUploadURI
    }
}
// Gets the fields property value. 
func (m *ExactMatchSession) GetFields()([]string) {
    if m == nil {
        return nil
    } else {
        return m.fields
    }
}
// Gets the fileName property value. 
func (m *ExactMatchSession) GetFileName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.fileName
    }
}
// Gets the rowsPerBlock property value. 
func (m *ExactMatchSession) GetRowsPerBlock()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.rowsPerBlock
    }
}
// Gets the salt property value. 
func (m *ExactMatchSession) GetSalt()(*string) {
    if m == nil {
        return nil
    } else {
        return m.salt
    }
}
// Gets the uploadAgent property value. 
func (m *ExactMatchSession) GetUploadAgent()(*ExactMatchUploadAgent) {
    if m == nil {
        return nil
    } else {
        return m.uploadAgent
    }
}
// Gets the uploadAgentId property value. 
func (m *ExactMatchSession) GetUploadAgentId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.uploadAgentId
    }
}
// The deserialization information for the current model
func (m *ExactMatchSession) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ExactMatchSessionBase.GetFieldDeserializers()
    res["checksum"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetChecksum(val)
        return nil
    }
    res["dataUploadURI"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDataUploadURI(val)
        return nil
    }
    res["fields"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetFields(res)
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
    res["rowsPerBlock"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetRowsPerBlock(val)
        return nil
    }
    res["salt"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSalt(val)
        return nil
    }
    res["uploadAgent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewExactMatchUploadAgent() })
        if err != nil {
            return err
        }
        m.SetUploadAgent(val.(*ExactMatchUploadAgent))
        return nil
    }
    res["uploadAgentId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUploadAgentId(val)
        return nil
    }
    return res
}
func (m *ExactMatchSession) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ExactMatchSession) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
    {
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
// Sets the checksum property value. 
// Parameters:
//  - value : Value to set for the checksum property.
func (m *ExactMatchSession) SetChecksum(value *string)() {
    m.checksum = value
}
// Sets the dataUploadURI property value. 
// Parameters:
//  - value : Value to set for the dataUploadURI property.
func (m *ExactMatchSession) SetDataUploadURI(value *string)() {
    m.dataUploadURI = value
}
// Sets the fields property value. 
// Parameters:
//  - value : Value to set for the fields property.
func (m *ExactMatchSession) SetFields(value []string)() {
    m.fields = value
}
// Sets the fileName property value. 
// Parameters:
//  - value : Value to set for the fileName property.
func (m *ExactMatchSession) SetFileName(value *string)() {
    m.fileName = value
}
// Sets the rowsPerBlock property value. 
// Parameters:
//  - value : Value to set for the rowsPerBlock property.
func (m *ExactMatchSession) SetRowsPerBlock(value *int32)() {
    m.rowsPerBlock = value
}
// Sets the salt property value. 
// Parameters:
//  - value : Value to set for the salt property.
func (m *ExactMatchSession) SetSalt(value *string)() {
    m.salt = value
}
// Sets the uploadAgent property value. 
// Parameters:
//  - value : Value to set for the uploadAgent property.
func (m *ExactMatchSession) SetUploadAgent(value *ExactMatchUploadAgent)() {
    m.uploadAgent = value
}
// Sets the uploadAgentId property value. 
// Parameters:
//  - value : Value to set for the uploadAgentId property.
func (m *ExactMatchSession) SetUploadAgentId(value *string)() {
    m.uploadAgentId = value
}
