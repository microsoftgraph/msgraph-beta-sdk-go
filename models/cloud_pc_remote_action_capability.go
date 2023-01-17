package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPcRemoteActionCapability 
type CloudPcRemoteActionCapability struct {
    // The actionCapability property
    actionCapability *ActionCapability
    // The actionName property
    actionName *CloudPcRemoteActionName
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The OdataType property
    odataType *string
}
// NewCloudPcRemoteActionCapability instantiates a new cloudPcRemoteActionCapability and sets the default values.
func NewCloudPcRemoteActionCapability()(*CloudPcRemoteActionCapability) {
    m := &CloudPcRemoteActionCapability{
    }
    m.SetAdditionalData(make(map[string]any));
    return m
}
// CreateCloudPcRemoteActionCapabilityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPcRemoteActionCapabilityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcRemoteActionCapability(), nil
}
// GetActionCapability gets the actionCapability property value. The actionCapability property
func (m *CloudPcRemoteActionCapability) GetActionCapability()(*ActionCapability) {
    return m.actionCapability
}
// GetActionName gets the actionName property value. The actionName property
func (m *CloudPcRemoteActionCapability) GetActionName()(*CloudPcRemoteActionName) {
    return m.actionName
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CloudPcRemoteActionCapability) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcRemoteActionCapability) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["actionCapability"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseActionCapability)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionCapability(val.(*ActionCapability))
        }
        return nil
    }
    res["actionName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcRemoteActionName)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionName(val.(*CloudPcRemoteActionName))
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
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *CloudPcRemoteActionCapability) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *CloudPcRemoteActionCapability) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetActionCapability() != nil {
        cast := (*m.GetActionCapability()).String()
        err := writer.WriteStringValue("actionCapability", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetActionName() != nil {
        cast := (*m.GetActionName()).String()
        err := writer.WriteStringValue("actionName", &cast)
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActionCapability sets the actionCapability property value. The actionCapability property
func (m *CloudPcRemoteActionCapability) SetActionCapability(value *ActionCapability)() {
    m.actionCapability = value
}
// SetActionName sets the actionName property value. The actionName property
func (m *CloudPcRemoteActionCapability) SetActionName(value *CloudPcRemoteActionName)() {
    m.actionName = value
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CloudPcRemoteActionCapability) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CloudPcRemoteActionCapability) SetOdataType(value *string)() {
    m.odataType = value
}
