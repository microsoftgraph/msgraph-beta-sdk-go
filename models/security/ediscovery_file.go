package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EdiscoveryFile 
type EdiscoveryFile struct {
    File
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Custodians associated with the file.
    custodian EdiscoveryCustodianable
    // Tags associated with the file.
    tags []EdiscoveryReviewTagable
}
// NewEdiscoveryFile instantiates a new EdiscoveryFile and sets the default values.
func NewEdiscoveryFile()(*EdiscoveryFile) {
    m := &EdiscoveryFile{
        File: *NewFile(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateEdiscoveryFileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEdiscoveryFileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEdiscoveryFile(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EdiscoveryFile) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCustodian gets the custodian property value. Custodians associated with the file.
func (m *EdiscoveryFile) GetCustodian()(EdiscoveryCustodianable) {
    if m == nil {
        return nil
    } else {
        return m.custodian
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EdiscoveryFile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.File.GetFieldDeserializers()
    res["custodian"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEdiscoveryCustodianFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustodian(val.(EdiscoveryCustodianable))
        }
        return nil
    }
    res["tags"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEdiscoveryReviewTagFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EdiscoveryReviewTagable, len(val))
            for i, v := range val {
                res[i] = v.(EdiscoveryReviewTagable)
            }
            m.SetTags(res)
        }
        return nil
    }
    return res
}
// GetTags gets the tags property value. Tags associated with the file.
func (m *EdiscoveryFile) GetTags()([]EdiscoveryReviewTagable) {
    if m == nil {
        return nil
    } else {
        return m.tags
    }
}
// Serialize serializes information the current object
func (m *EdiscoveryFile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.File.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("custodian", m.GetCustodian())
        if err != nil {
            return err
        }
    }
    if m.GetTags() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTags()))
        for i, v := range m.GetTags() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("tags", cast)
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
func (m *EdiscoveryFile) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCustodian sets the custodian property value. Custodians associated with the file.
func (m *EdiscoveryFile) SetCustodian(value EdiscoveryCustodianable)() {
    if m != nil {
        m.custodian = value
    }
}
// SetTags sets the tags property value. Tags associated with the file.
func (m *EdiscoveryFile) SetTags(value []EdiscoveryReviewTagable)() {
    if m != nil {
        m.tags = value
    }
}
