package industrydata

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// FileValidateOperation 
type FileValidateOperation struct {
    ValidateOperation
}
// NewFileValidateOperation instantiates a new FileValidateOperation and sets the default values.
func NewFileValidateOperation()(*FileValidateOperation) {
    m := &FileValidateOperation{
        ValidateOperation: *NewValidateOperation(),
    }
    odataTypeValue := "#microsoft.graph.industryData.fileValidateOperation"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateFileValidateOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFileValidateOperationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFileValidateOperation(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *FileValidateOperation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ValidateOperation.GetFieldDeserializers()
    res["validatedFiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetValidatedFiles(res)
        }
        return nil
    }
    return res
}
// GetValidatedFiles gets the validatedFiles property value. Set of files validated by the validate operation.
func (m *FileValidateOperation) GetValidatedFiles()([]string) {
    val, err := m.GetBackingStore().Get("validatedFiles")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *FileValidateOperation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ValidateOperation.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// SetValidatedFiles sets the validatedFiles property value. Set of files validated by the validate operation.
func (m *FileValidateOperation) SetValidatedFiles(value []string)() {
    err := m.GetBackingStore().Set("validatedFiles", value)
    if err != nil {
        panic(err)
    }
}
// FileValidateOperationable 
type FileValidateOperationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ValidateOperationable
    GetValidatedFiles()([]string)
    SetValidatedFiles(value []string)()
}
