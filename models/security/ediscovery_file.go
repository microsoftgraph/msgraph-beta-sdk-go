package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EdiscoveryFile 
type EdiscoveryFile struct {
    File
}
// NewEdiscoveryFile instantiates a new ediscoveryFile and sets the default values.
func NewEdiscoveryFile()(*EdiscoveryFile) {
    m := &EdiscoveryFile{
        File: *NewFile(),
    }
    return m
}
// CreateEdiscoveryFileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEdiscoveryFileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEdiscoveryFile(), nil
}
// GetCustodian gets the custodian property value. Custodians associated with the file.
func (m *EdiscoveryFile) GetCustodian()(EdiscoveryCustodianable) {
    val, err := m.GetBackingStore().Get("custodian")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(EdiscoveryCustodianable)
    }
    return nil
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
                if v != nil {
                    res[i] = v.(EdiscoveryReviewTagable)
                }
            }
            m.SetTags(res)
        }
        return nil
    }
    return res
}
// GetTags gets the tags property value. Tags associated with the file.
func (m *EdiscoveryFile) GetTags()([]EdiscoveryReviewTagable) {
    val, err := m.GetBackingStore().Get("tags")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]EdiscoveryReviewTagable)
    }
    return nil
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
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("tags", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCustodian sets the custodian property value. Custodians associated with the file.
func (m *EdiscoveryFile) SetCustodian(value EdiscoveryCustodianable)() {
    err := m.GetBackingStore().Set("custodian", value)
    if err != nil {
        panic(err)
    }
}
// SetTags sets the tags property value. Tags associated with the file.
func (m *EdiscoveryFile) SetTags(value []EdiscoveryReviewTagable)() {
    err := m.GetBackingStore().Set("tags", value)
    if err != nil {
        panic(err)
    }
}
// EdiscoveryFileable 
type EdiscoveryFileable interface {
    Fileable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCustodian()(EdiscoveryCustodianable)
    GetTags()([]EdiscoveryReviewTagable)
    SetCustodian(value EdiscoveryCustodianable)()
    SetTags(value []EdiscoveryReviewTagable)()
}
