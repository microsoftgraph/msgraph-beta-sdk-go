package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DomainDnsTxtRecord 
type DomainDnsTxtRecord struct {
    DomainDnsRecord
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Value used when configuring the text property at the DNS host.
    text *string
}
// NewDomainDnsTxtRecord instantiates a new DomainDnsTxtRecord and sets the default values.
func NewDomainDnsTxtRecord()(*DomainDnsTxtRecord) {
    m := &DomainDnsTxtRecord{
        DomainDnsRecord: *NewDomainDnsRecord(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDomainDnsTxtRecordFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDomainDnsTxtRecordFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDomainDnsTxtRecord(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DomainDnsTxtRecord) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DomainDnsTxtRecord) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DomainDnsRecord.GetFieldDeserializers()
    res["text"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetText(val)
        }
        return nil
    }
    return res
}
// GetText gets the text property value. Value used when configuring the text property at the DNS host.
func (m *DomainDnsTxtRecord) GetText()(*string) {
    if m == nil {
        return nil
    } else {
        return m.text
    }
}
// Serialize serializes information the current object
func (m *DomainDnsTxtRecord) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DomainDnsRecord.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("text", m.GetText())
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
func (m *DomainDnsTxtRecord) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetText sets the text property value. Value used when configuring the text property at the DNS host.
func (m *DomainDnsTxtRecord) SetText(value *string)() {
    if m != nil {
        m.text = value
    }
}
