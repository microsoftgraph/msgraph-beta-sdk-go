package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessPackageQuestion 
type AccessPackageQuestion struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // ID of the question.
    id *string
    // Specifies whether the requestor is allowed to edit answers to questions.
    isAnswerEditable *bool
    // Whether the requestor is required to supply an answer or not.
    isRequired *bool
    // The OdataType property
    odataType *string
    // Relative position of this question when displaying a list of questions to the requestor.
    sequence *int32
    // The text of the question to show to the requestor.
    text AccessPackageLocalizedContentable
}
// NewAccessPackageQuestion instantiates a new accessPackageQuestion and sets the default values.
func NewAccessPackageQuestion()(*AccessPackageQuestion) {
    m := &AccessPackageQuestion{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.accessPackageQuestion";
    m.SetOdataType(&odataTypeValue);
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
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AccessPackageQuestion) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessPackageQuestion) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["id"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetId)
    res["isAnswerEditable"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsAnswerEditable)
    res["isRequired"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsRequired)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["sequence"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetSequence)
    res["text"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateAccessPackageLocalizedContentFromDiscriminatorValue , m.SetText)
    return res
}
// GetId gets the id property value. ID of the question.
func (m *AccessPackageQuestion) GetId()(*string) {
    return m.id
}
// GetIsAnswerEditable gets the isAnswerEditable property value. Specifies whether the requestor is allowed to edit answers to questions.
func (m *AccessPackageQuestion) GetIsAnswerEditable()(*bool) {
    return m.isAnswerEditable
}
// GetIsRequired gets the isRequired property value. Whether the requestor is required to supply an answer or not.
func (m *AccessPackageQuestion) GetIsRequired()(*bool) {
    return m.isRequired
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AccessPackageQuestion) GetOdataType()(*string) {
    return m.odataType
}
// GetSequence gets the sequence property value. Relative position of this question when displaying a list of questions to the requestor.
func (m *AccessPackageQuestion) GetSequence()(*int32) {
    return m.sequence
}
// GetText gets the text property value. The text of the question to show to the requestor.
func (m *AccessPackageQuestion) GetText()(AccessPackageLocalizedContentable) {
    return m.text
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AccessPackageQuestion) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetId sets the id property value. ID of the question.
func (m *AccessPackageQuestion) SetId(value *string)() {
    m.id = value
}
// SetIsAnswerEditable sets the isAnswerEditable property value. Specifies whether the requestor is allowed to edit answers to questions.
func (m *AccessPackageQuestion) SetIsAnswerEditable(value *bool)() {
    m.isAnswerEditable = value
}
// SetIsRequired sets the isRequired property value. Whether the requestor is required to supply an answer or not.
func (m *AccessPackageQuestion) SetIsRequired(value *bool)() {
    m.isRequired = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AccessPackageQuestion) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSequence sets the sequence property value. Relative position of this question when displaying a list of questions to the requestor.
func (m *AccessPackageQuestion) SetSequence(value *int32)() {
    m.sequence = value
}
// SetText sets the text property value. The text of the question to show to the requestor.
func (m *AccessPackageQuestion) SetText(value AccessPackageLocalizedContentable)() {
    m.text = value
}
