package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MoveAction 
type MoveAction struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The name of the location the item was moved from.
    from *string
    // The OdataType property
    odataType *string
    // The name of the location the item was moved to.
    to *string
}
// NewMoveAction instantiates a new moveAction and sets the default values.
func NewMoveAction()(*MoveAction) {
    m := &MoveAction{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.moveAction";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateMoveActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMoveActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMoveAction(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MoveAction) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MoveAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["from"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetFrom)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["to"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetTo)
    return res
}
// GetFrom gets the from property value. The name of the location the item was moved from.
func (m *MoveAction) GetFrom()(*string) {
    return m.from
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *MoveAction) GetOdataType()(*string) {
    return m.odataType
}
// GetTo gets the to property value. The name of the location the item was moved to.
func (m *MoveAction) GetTo()(*string) {
    return m.to
}
// Serialize serializes information the current object
func (m *MoveAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("from", m.GetFrom())
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
        err := writer.WriteStringValue("to", m.GetTo())
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
func (m *MoveAction) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetFrom sets the from property value. The name of the location the item was moved from.
func (m *MoveAction) SetFrom(value *string)() {
    m.from = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *MoveAction) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTo sets the to property value. The name of the location the item was moved to.
func (m *MoveAction) SetTo(value *string)() {
    m.to = value
}
