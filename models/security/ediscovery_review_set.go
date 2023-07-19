package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EdiscoveryReviewSet 
type EdiscoveryReviewSet struct {
    DataSet
}
// NewEdiscoveryReviewSet instantiates a new ediscoveryReviewSet and sets the default values.
func NewEdiscoveryReviewSet()(*EdiscoveryReviewSet) {
    m := &EdiscoveryReviewSet{
        DataSet: *NewDataSet(),
    }
    odataTypeValue := "#microsoft.graph.security.ediscoveryReviewSet"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateEdiscoveryReviewSetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEdiscoveryReviewSetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEdiscoveryReviewSet(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EdiscoveryReviewSet) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DataSet.GetFieldDeserializers()
    res["files"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEdiscoveryFileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EdiscoveryFileable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(EdiscoveryFileable)
                }
            }
            m.SetFiles(res)
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
    res["queries"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEdiscoveryReviewSetQueryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EdiscoveryReviewSetQueryable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(EdiscoveryReviewSetQueryable)
                }
            }
            m.SetQueries(res)
        }
        return nil
    }
    return res
}
// GetFiles gets the files property value. Represents files within the review set.
func (m *EdiscoveryReviewSet) GetFiles()([]EdiscoveryFileable) {
    val, err := m.GetBackingStore().Get("files")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]EdiscoveryFileable)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *EdiscoveryReviewSet) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetQueries gets the queries property value. Represents queries within the review set.
func (m *EdiscoveryReviewSet) GetQueries()([]EdiscoveryReviewSetQueryable) {
    val, err := m.GetBackingStore().Get("queries")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]EdiscoveryReviewSetQueryable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *EdiscoveryReviewSet) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DataSet.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetFiles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetFiles()))
        for i, v := range m.GetFiles() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("files", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetQueries() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetQueries()))
        for i, v := range m.GetQueries() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("queries", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFiles sets the files property value. Represents files within the review set.
func (m *EdiscoveryReviewSet) SetFiles(value []EdiscoveryFileable)() {
    err := m.GetBackingStore().Set("files", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *EdiscoveryReviewSet) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetQueries sets the queries property value. Represents queries within the review set.
func (m *EdiscoveryReviewSet) SetQueries(value []EdiscoveryReviewSetQueryable)() {
    err := m.GetBackingStore().Set("queries", value)
    if err != nil {
        panic(err)
    }
}
// EdiscoveryReviewSetable 
type EdiscoveryReviewSetable interface {
    DataSetable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFiles()([]EdiscoveryFileable)
    GetOdataType()(*string)
    GetQueries()([]EdiscoveryReviewSetQueryable)
    SetFiles(value []EdiscoveryFileable)()
    SetOdataType(value *string)()
    SetQueries(value []EdiscoveryReviewSetQueryable)()
}
