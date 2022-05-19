package classifytext

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ClassifyTextPostRequestBody provides operations to call the classifyText method.
type ClassifyTextPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The fileExtension property
    fileExtension *string
    // The matchTolerancesToInclude property
    matchTolerancesToInclude *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MlClassificationMatchTolerance
    // The scopesToRun property
    scopesToRun *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SensitiveTypeScope
    // The sensitiveTypeIds property
    sensitiveTypeIds []string
    // The text property
    text *string
}
// NewClassifyTextPostRequestBody instantiates a new classifyTextPostRequestBody and sets the default values.
func NewClassifyTextPostRequestBody()(*ClassifyTextPostRequestBody) {
    m := &ClassifyTextPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateClassifyTextPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateClassifyTextPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewClassifyTextPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ClassifyTextPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ClassifyTextPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
        val, err := n.GetEnumValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ParseMlClassificationMatchTolerance)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMatchTolerancesToInclude(val.(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MlClassificationMatchTolerance))
        }
        return nil
    }
    res["scopesToRun"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ParseSensitiveTypeScope)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScopesToRun(val.(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SensitiveTypeScope))
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
func (m *ClassifyTextPostRequestBody) GetFileExtension()(*string) {
    if m == nil {
        return nil
    } else {
        return m.fileExtension
    }
}
// GetMatchTolerancesToInclude gets the matchTolerancesToInclude property value. The matchTolerancesToInclude property
func (m *ClassifyTextPostRequestBody) GetMatchTolerancesToInclude()(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MlClassificationMatchTolerance) {
    if m == nil {
        return nil
    } else {
        return m.matchTolerancesToInclude
    }
}
// GetScopesToRun gets the scopesToRun property value. The scopesToRun property
func (m *ClassifyTextPostRequestBody) GetScopesToRun()(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SensitiveTypeScope) {
    if m == nil {
        return nil
    } else {
        return m.scopesToRun
    }
}
// GetSensitiveTypeIds gets the sensitiveTypeIds property value. The sensitiveTypeIds property
func (m *ClassifyTextPostRequestBody) GetSensitiveTypeIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.sensitiveTypeIds
    }
}
// GetText gets the text property value. The text property
func (m *ClassifyTextPostRequestBody) GetText()(*string) {
    if m == nil {
        return nil
    } else {
        return m.text
    }
}
// Serialize serializes information the current object
func (m *ClassifyTextPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("fileExtension", m.GetFileExtension())
        if err != nil {
            return err
        }
    }
    if m.GetMatchTolerancesToInclude() != nil {
        cast := (*m.GetMatchTolerancesToInclude()).String()
        err := writer.WriteStringValue("matchTolerancesToInclude", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetScopesToRun() != nil {
        cast := (*m.GetScopesToRun()).String()
        err := writer.WriteStringValue("scopesToRun", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSensitiveTypeIds() != nil {
        err := writer.WriteCollectionOfStringValues("sensitiveTypeIds", m.GetSensitiveTypeIds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("text", m.GetText())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ClassifyTextPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetFileExtension sets the fileExtension property value. The fileExtension property
func (m *ClassifyTextPostRequestBody) SetFileExtension(value *string)() {
    if m != nil {
        m.fileExtension = value
    }
}
// SetMatchTolerancesToInclude sets the matchTolerancesToInclude property value. The matchTolerancesToInclude property
func (m *ClassifyTextPostRequestBody) SetMatchTolerancesToInclude(value *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MlClassificationMatchTolerance)() {
    if m != nil {
        m.matchTolerancesToInclude = value
    }
}
// SetScopesToRun sets the scopesToRun property value. The scopesToRun property
func (m *ClassifyTextPostRequestBody) SetScopesToRun(value *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SensitiveTypeScope)() {
    if m != nil {
        m.scopesToRun = value
    }
}
// SetSensitiveTypeIds sets the sensitiveTypeIds property value. The sensitiveTypeIds property
func (m *ClassifyTextPostRequestBody) SetSensitiveTypeIds(value []string)() {
    if m != nil {
        m.sensitiveTypeIds = value
    }
}
// SetText sets the text property value. The text property
func (m *ClassifyTextPostRequestBody) SetText(value *string)() {
    if m != nil {
        m.text = value
    }
}
