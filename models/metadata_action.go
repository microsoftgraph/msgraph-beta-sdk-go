package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MetadataAction 
type MetadataAction struct {
    InformationProtectionAction
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // A collection of key value pairs that should be added to the file.
    metadataToAdd []KeyValuePairable
    // A collection of strings that indicate which keys to remove from the file metadata.
    metadataToRemove []string
}
// NewMetadataAction instantiates a new MetadataAction and sets the default values.
func NewMetadataAction()(*MetadataAction) {
    m := &MetadataAction{
        InformationProtectionAction: *NewInformationProtectionAction(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMetadataActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMetadataActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMetadataAction(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MetadataAction) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MetadataAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.InformationProtectionAction.GetFieldDeserializers()
    res["metadataToAdd"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateKeyValuePairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]KeyValuePairable, len(val))
            for i, v := range val {
                res[i] = v.(KeyValuePairable)
            }
            m.SetMetadataToAdd(res)
        }
        return nil
    }
    res["metadataToRemove"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetMetadataToRemove(res)
        }
        return nil
    }
    return res
}
// GetMetadataToAdd gets the metadataToAdd property value. A collection of key value pairs that should be added to the file.
func (m *MetadataAction) GetMetadataToAdd()([]KeyValuePairable) {
    if m == nil {
        return nil
    } else {
        return m.metadataToAdd
    }
}
// GetMetadataToRemove gets the metadataToRemove property value. A collection of strings that indicate which keys to remove from the file metadata.
func (m *MetadataAction) GetMetadataToRemove()([]string) {
    if m == nil {
        return nil
    } else {
        return m.metadataToRemove
    }
}
// Serialize serializes information the current object
func (m *MetadataAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.InformationProtectionAction.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetMetadataToAdd() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMetadataToAdd()))
        for i, v := range m.GetMetadataToAdd() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("metadataToAdd", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMetadataToRemove() != nil {
        err = writer.WriteCollectionOfStringValues("metadataToRemove", m.GetMetadataToRemove())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MetadataAction) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetMetadataToAdd sets the metadataToAdd property value. A collection of key value pairs that should be added to the file.
func (m *MetadataAction) SetMetadataToAdd(value []KeyValuePairable)() {
    if m != nil {
        m.metadataToAdd = value
    }
}
// SetMetadataToRemove sets the metadataToRemove property value. A collection of strings that indicate which keys to remove from the file metadata.
func (m *MetadataAction) SetMetadataToRemove(value []string)() {
    if m != nil {
        m.metadataToRemove = value
    }
}
