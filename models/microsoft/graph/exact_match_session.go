package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ExactMatchSession 
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
// NewExactMatchSession instantiates a new exactMatchSession and sets the default values.
func NewExactMatchSession()(*ExactMatchSession) {
    m := &ExactMatchSession{
        ExactMatchSessionBase: *NewExactMatchSessionBase(),
    }
    return m
}
// GetChecksum gets the checksum property value. 
func (m *ExactMatchSession) GetChecksum()(*string) {
    if m == nil {
        return nil
    } else {
        return m.checksum
    }
}
// GetDataUploadURI gets the dataUploadURI property value. 
func (m *ExactMatchSession) GetDataUploadURI()(*string) {
    if m == nil {
        return nil
    } else {
        return m.dataUploadURI
    }
}
// GetFields gets the fields property value. 
func (m *ExactMatchSession) GetFields()([]string) {
    if m == nil {
        return nil
    } else {
        return m.fields
    }
}
// GetFileName gets the fileName property value. 
func (m *ExactMatchSession) GetFileName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.fileName
    }
}
// GetRowsPerBlock gets the rowsPerBlock property value. 
func (m *ExactMatchSession) GetRowsPerBlock()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.rowsPerBlock
    }
}
// GetSalt gets the salt property value. 
func (m *ExactMatchSession) GetSalt()(*string) {
    if m == nil {
        return nil
    } else {
        return m.salt
    }
}
// GetUploadAgent gets the uploadAgent property value. 
func (m *ExactMatchSession) GetUploadAgent()(*ExactMatchUploadAgent) {
    if m == nil {
        return nil
    } else {
        return m.uploadAgent
    }
}
// GetUploadAgentId gets the uploadAgentId property value. 
func (m *ExactMatchSession) GetUploadAgentId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.uploadAgentId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ExactMatchSession) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ExactMatchSessionBase.GetFieldDeserializers()
    res["checksum"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChecksum(val)
        }
        return nil
    }
    res["dataUploadURI"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataUploadURI(val)
        }
        return nil
    }
    res["fields"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetFields(res)
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
    res["rowsPerBlock"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRowsPerBlock(val)
        }
        return nil
    }
    res["salt"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSalt(val)
        }
        return nil
    }
    res["uploadAgent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewExactMatchUploadAgent() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUploadAgent(val.(*ExactMatchUploadAgent))
        }
        return nil
    }
    res["uploadAgentId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUploadAgentId(val)
        }
        return nil
    }
    return res
}
func (m *ExactMatchSession) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetChecksum sets the checksum property value. 
func (m *ExactMatchSession) SetChecksum(value *string)() {
    if m != nil {
        m.checksum = value
    }
}
// SetDataUploadURI sets the dataUploadURI property value. 
func (m *ExactMatchSession) SetDataUploadURI(value *string)() {
    if m != nil {
        m.dataUploadURI = value
    }
}
// SetFields sets the fields property value. 
func (m *ExactMatchSession) SetFields(value []string)() {
    if m != nil {
        m.fields = value
    }
}
// SetFileName sets the fileName property value. 
func (m *ExactMatchSession) SetFileName(value *string)() {
    if m != nil {
        m.fileName = value
    }
}
// SetRowsPerBlock sets the rowsPerBlock property value. 
func (m *ExactMatchSession) SetRowsPerBlock(value *int32)() {
    if m != nil {
        m.rowsPerBlock = value
    }
}
// SetSalt sets the salt property value. 
func (m *ExactMatchSession) SetSalt(value *string)() {
    if m != nil {
        m.salt = value
    }
}
// SetUploadAgent sets the uploadAgent property value. 
func (m *ExactMatchSession) SetUploadAgent(value *ExactMatchUploadAgent)() {
    if m != nil {
        m.uploadAgent = value
    }
}
// SetUploadAgentId sets the uploadAgentId property value. 
func (m *ExactMatchSession) SetUploadAgentId(value *string)() {
    if m != nil {
        m.uploadAgentId = value
    }
}
