package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TextClassificationRequest 
type TextClassificationRequest struct {
    Entity
}
// NewTextClassificationRequest instantiates a new TextClassificationRequest and sets the default values.
func NewTextClassificationRequest()(*TextClassificationRequest) {
    m := &TextClassificationRequest{
        Entity: *NewEntity(),
    }
    return m
}
// CreateTextClassificationRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTextClassificationRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTextClassificationRequest(), nil
}
// GetContentMetaData gets the contentMetaData property value. The contentMetaData property
func (m *TextClassificationRequest) GetContentMetaData()(ClassificationRequestContentMetaDataable) {
    val, err := m.GetBackingStore().Get("contentMetaData")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ClassificationRequestContentMetaDataable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TextClassificationRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["contentMetaData"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateClassificationRequestContentMetaDataFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentMetaData(val.(ClassificationRequestContentMetaDataable))
        }
        return nil
    }
    res["fileExtension"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileExtension(val)
        }
        return nil
    }
    res["matchTolerancesToInclude"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMlClassificationMatchTolerance)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMatchTolerancesToInclude(val.(*MlClassificationMatchTolerance))
        }
        return nil
    }
    res["scopesToRun"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSensitiveTypeScope)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScopesToRun(val.(*SensitiveTypeScope))
        }
        return nil
    }
    res["sensitiveTypeIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSensitiveTypeIds(res)
        }
        return nil
    }
    res["text"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetText(val)
        }
        return nil
    }
    return res
}
// GetFileExtension gets the fileExtension property value. The fileExtension property
func (m *TextClassificationRequest) GetFileExtension()(*string) {
    val, err := m.GetBackingStore().Get("fileExtension")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMatchTolerancesToInclude gets the matchTolerancesToInclude property value. The matchTolerancesToInclude property
func (m *TextClassificationRequest) GetMatchTolerancesToInclude()(*MlClassificationMatchTolerance) {
    val, err := m.GetBackingStore().Get("matchTolerancesToInclude")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*MlClassificationMatchTolerance)
    }
    return nil
}
// GetScopesToRun gets the scopesToRun property value. The scopesToRun property
func (m *TextClassificationRequest) GetScopesToRun()(*SensitiveTypeScope) {
    val, err := m.GetBackingStore().Get("scopesToRun")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*SensitiveTypeScope)
    }
    return nil
}
// GetSensitiveTypeIds gets the sensitiveTypeIds property value. The sensitiveTypeIds property
func (m *TextClassificationRequest) GetSensitiveTypeIds()([]string) {
    val, err := m.GetBackingStore().Get("sensitiveTypeIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetText gets the text property value. The text property
func (m *TextClassificationRequest) GetText()(*string) {
    val, err := m.GetBackingStore().Get("text")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *TextClassificationRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("contentMetaData", m.GetContentMetaData())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("fileExtension", m.GetFileExtension())
        if err != nil {
            return err
        }
    }
    if m.GetMatchTolerancesToInclude() != nil {
        cast := (*m.GetMatchTolerancesToInclude()).String()
        err = writer.WriteStringValue("matchTolerancesToInclude", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetScopesToRun() != nil {
        cast := (*m.GetScopesToRun()).String()
        err = writer.WriteStringValue("scopesToRun", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSensitiveTypeIds() != nil {
        err = writer.WriteCollectionOfStringValues("sensitiveTypeIds", m.GetSensitiveTypeIds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("text", m.GetText())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContentMetaData sets the contentMetaData property value. The contentMetaData property
func (m *TextClassificationRequest) SetContentMetaData(value ClassificationRequestContentMetaDataable)() {
    err := m.GetBackingStore().Set("contentMetaData", value)
    if err != nil {
        panic(err)
    }
}
// SetFileExtension sets the fileExtension property value. The fileExtension property
func (m *TextClassificationRequest) SetFileExtension(value *string)() {
    err := m.GetBackingStore().Set("fileExtension", value)
    if err != nil {
        panic(err)
    }
}
// SetMatchTolerancesToInclude sets the matchTolerancesToInclude property value. The matchTolerancesToInclude property
func (m *TextClassificationRequest) SetMatchTolerancesToInclude(value *MlClassificationMatchTolerance)() {
    err := m.GetBackingStore().Set("matchTolerancesToInclude", value)
    if err != nil {
        panic(err)
    }
}
// SetScopesToRun sets the scopesToRun property value. The scopesToRun property
func (m *TextClassificationRequest) SetScopesToRun(value *SensitiveTypeScope)() {
    err := m.GetBackingStore().Set("scopesToRun", value)
    if err != nil {
        panic(err)
    }
}
// SetSensitiveTypeIds sets the sensitiveTypeIds property value. The sensitiveTypeIds property
func (m *TextClassificationRequest) SetSensitiveTypeIds(value []string)() {
    err := m.GetBackingStore().Set("sensitiveTypeIds", value)
    if err != nil {
        panic(err)
    }
}
// SetText sets the text property value. The text property
func (m *TextClassificationRequest) SetText(value *string)() {
    err := m.GetBackingStore().Set("text", value)
    if err != nil {
        panic(err)
    }
}
// TextClassificationRequestable 
type TextClassificationRequestable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContentMetaData()(ClassificationRequestContentMetaDataable)
    GetFileExtension()(*string)
    GetMatchTolerancesToInclude()(*MlClassificationMatchTolerance)
    GetScopesToRun()(*SensitiveTypeScope)
    GetSensitiveTypeIds()([]string)
    GetText()(*string)
    SetContentMetaData(value ClassificationRequestContentMetaDataable)()
    SetFileExtension(value *string)()
    SetMatchTolerancesToInclude(value *MlClassificationMatchTolerance)()
    SetScopesToRun(value *SensitiveTypeScope)()
    SetSensitiveTypeIds(value []string)()
    SetText(value *string)()
}
