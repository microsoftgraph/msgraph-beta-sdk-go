package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type MailFolderOperation struct {
    Entity
}
// NewMailFolderOperation instantiates a new MailFolderOperation and sets the default values.
func NewMailFolderOperation()(*MailFolderOperation) {
    m := &MailFolderOperation{
        Entity: *NewEntity(),
    }
    return m
}
// CreateMailFolderOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMailFolderOperationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.updateAllMessagesReadStateOperation":
                        return NewUpdateAllMessagesReadStateOperation(), nil
                }
            }
        }
    }
    return NewMailFolderOperation(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MailFolderOperation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["resourceLocation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceLocation(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMailFolderOperationStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*MailFolderOperationStatus))
        }
        return nil
    }
    return res
}
// GetResourceLocation gets the resourceLocation property value. The location of the long-running operation.
// returns a *string when successful
func (m *MailFolderOperation) GetResourceLocation()(*string) {
    val, err := m.GetBackingStore().Get("resourceLocation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetStatus gets the status property value. The status of the long-running operation. The possible values are: notStarted, running, succeeded, failed, unknownFutureValue.
// returns a *MailFolderOperationStatus when successful
func (m *MailFolderOperation) GetStatus()(*MailFolderOperationStatus) {
    val, err := m.GetBackingStore().Get("status")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*MailFolderOperationStatus)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MailFolderOperation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("resourceLocation", m.GetResourceLocation())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetResourceLocation sets the resourceLocation property value. The location of the long-running operation.
func (m *MailFolderOperation) SetResourceLocation(value *string)() {
    err := m.GetBackingStore().Set("resourceLocation", value)
    if err != nil {
        panic(err)
    }
}
// SetStatus sets the status property value. The status of the long-running operation. The possible values are: notStarted, running, succeeded, failed, unknownFutureValue.
func (m *MailFolderOperation) SetStatus(value *MailFolderOperationStatus)() {
    err := m.GetBackingStore().Set("status", value)
    if err != nil {
        panic(err)
    }
}
type MailFolderOperationable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetResourceLocation()(*string)
    GetStatus()(*MailFolderOperationStatus)
    SetResourceLocation(value *string)()
    SetStatus(value *MailFolderOperationStatus)()
}
