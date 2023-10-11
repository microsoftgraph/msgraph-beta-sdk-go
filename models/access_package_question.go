package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// AccessPackageQuestion 
type AccessPackageQuestion struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewAccessPackageQuestion instantiates a new accessPackageQuestion and sets the default values.
func NewAccessPackageQuestion()(*AccessPackageQuestion) {
    m := &AccessPackageQuestion{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAccessPackageQuestionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessPackageQuestionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.accessPackageMultipleChoiceQuestion":
                        return NewAccessPackageMultipleChoiceQuestion(), nil
                    case "#microsoft.graph.accessPackageTextInputQuestion":
                        return NewAccessPackageTextInputQuestion(), nil
                }
            }
        }
    }
    return NewAccessPackageQuestion(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AccessPackageQuestion) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the BackingStore property value. Stores model information.
func (m *AccessPackageQuestion) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessPackageQuestion) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["isAnswerEditable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAnswerEditable(val)
        }
        return nil
    }
    res["isRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsRequired(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["sequence"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSequence(val)
        }
        return nil
    }
    res["text"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessPackageLocalizedContentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetText(val.(AccessPackageLocalizedContentable))
        }
        return nil
    }
    return res
}
// GetId gets the id property value. ID of the question.
func (m *AccessPackageQuestion) GetId()(*string) {
    val, err := m.GetBackingStore().Get("id")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetIsAnswerEditable gets the isAnswerEditable property value. Specifies whether the requestor is allowed to edit answers to questions.
func (m *AccessPackageQuestion) GetIsAnswerEditable()(*bool) {
    val, err := m.GetBackingStore().Get("isAnswerEditable")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsRequired gets the isRequired property value. Whether the requestor is required to supply an answer or not.
func (m *AccessPackageQuestion) GetIsRequired()(*bool) {
    val, err := m.GetBackingStore().Get("isRequired")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AccessPackageQuestion) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSequence gets the sequence property value. Relative position of this question when displaying a list of questions to the requestor.
func (m *AccessPackageQuestion) GetSequence()(*int32) {
    val, err := m.GetBackingStore().Get("sequence")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetText gets the text property value. The text of the question to show to the requestor.
func (m *AccessPackageQuestion) GetText()(AccessPackageLocalizedContentable) {
    val, err := m.GetBackingStore().Get("text")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AccessPackageLocalizedContentable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AccessPackageQuestion) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isAnswerEditable", m.GetIsAnswerEditable())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isRequired", m.GetIsRequired())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("sequence", m.GetSequence())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("text", m.GetText())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AccessPackageQuestion) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *AccessPackageQuestion) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetId sets the id property value. ID of the question.
func (m *AccessPackageQuestion) SetId(value *string)() {
    err := m.GetBackingStore().Set("id", value)
    if err != nil {
        panic(err)
    }
}
// SetIsAnswerEditable sets the isAnswerEditable property value. Specifies whether the requestor is allowed to edit answers to questions.
func (m *AccessPackageQuestion) SetIsAnswerEditable(value *bool)() {
    err := m.GetBackingStore().Set("isAnswerEditable", value)
    if err != nil {
        panic(err)
    }
}
// SetIsRequired sets the isRequired property value. Whether the requestor is required to supply an answer or not.
func (m *AccessPackageQuestion) SetIsRequired(value *bool)() {
    err := m.GetBackingStore().Set("isRequired", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AccessPackageQuestion) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetSequence sets the sequence property value. Relative position of this question when displaying a list of questions to the requestor.
func (m *AccessPackageQuestion) SetSequence(value *int32)() {
    err := m.GetBackingStore().Set("sequence", value)
    if err != nil {
        panic(err)
    }
}
// SetText sets the text property value. The text of the question to show to the requestor.
func (m *AccessPackageQuestion) SetText(value AccessPackageLocalizedContentable)() {
    err := m.GetBackingStore().Set("text", value)
    if err != nil {
        panic(err)
    }
}
// AccessPackageQuestionable 
type AccessPackageQuestionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetId()(*string)
    GetIsAnswerEditable()(*bool)
    GetIsRequired()(*bool)
    GetOdataType()(*string)
    GetSequence()(*int32)
    GetText()(AccessPackageLocalizedContentable)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetId(value *string)()
    SetIsAnswerEditable(value *bool)()
    SetIsRequired(value *bool)()
    SetOdataType(value *string)()
    SetSequence(value *int32)()
    SetText(value AccessPackageLocalizedContentable)()
}
