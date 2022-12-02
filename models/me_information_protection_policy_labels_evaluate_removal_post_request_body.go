package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeInformationProtectionPolicyLabelsEvaluateRemovalPostRequestBody provides operations to call the evaluateRemoval method.
type MeInformationProtectionPolicyLabelsEvaluateRemovalPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The contentInfo property
    contentInfo ContentInfoable
    // The downgradeJustification property
    downgradeJustification DowngradeJustificationable
}
// NewMeInformationProtectionPolicyLabelsEvaluateRemovalPostRequestBody instantiates a new MeInformationProtectionPolicyLabelsEvaluateRemovalPostRequestBody and sets the default values.
func NewMeInformationProtectionPolicyLabelsEvaluateRemovalPostRequestBody()(*MeInformationProtectionPolicyLabelsEvaluateRemovalPostRequestBody) {
    m := &MeInformationProtectionPolicyLabelsEvaluateRemovalPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMeInformationProtectionPolicyLabelsEvaluateRemovalPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMeInformationProtectionPolicyLabelsEvaluateRemovalPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMeInformationProtectionPolicyLabelsEvaluateRemovalPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MeInformationProtectionPolicyLabelsEvaluateRemovalPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetContentInfo gets the contentInfo property value. The contentInfo property
func (m *MeInformationProtectionPolicyLabelsEvaluateRemovalPostRequestBody) GetContentInfo()(ContentInfoable) {
    return m.contentInfo
}
// GetDowngradeJustification gets the downgradeJustification property value. The downgradeJustification property
func (m *MeInformationProtectionPolicyLabelsEvaluateRemovalPostRequestBody) GetDowngradeJustification()(DowngradeJustificationable) {
    return m.downgradeJustification
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeInformationProtectionPolicyLabelsEvaluateRemovalPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["contentInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateContentInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentInfo(val.(ContentInfoable))
        }
        return nil
    }
    res["downgradeJustification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDowngradeJustificationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDowngradeJustification(val.(DowngradeJustificationable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *MeInformationProtectionPolicyLabelsEvaluateRemovalPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("contentInfo", m.GetContentInfo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("downgradeJustification", m.GetDowngradeJustification())
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
func (m *MeInformationProtectionPolicyLabelsEvaluateRemovalPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetContentInfo sets the contentInfo property value. The contentInfo property
func (m *MeInformationProtectionPolicyLabelsEvaluateRemovalPostRequestBody) SetContentInfo(value ContentInfoable)() {
    m.contentInfo = value
}
// SetDowngradeJustification sets the downgradeJustification property value. The downgradeJustification property
func (m *MeInformationProtectionPolicyLabelsEvaluateRemovalPostRequestBody) SetDowngradeJustification(value DowngradeJustificationable)() {
    m.downgradeJustification = value
}
