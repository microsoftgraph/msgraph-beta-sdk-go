package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ContentInfo 
type ContentInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The contentFormat property
    contentFormat *string
    // The identifier property
    identifier *string
    // The metadata property
    metadata []KeyValuePairable
    // The state property
    state *ContentState
}
// NewContentInfo instantiates a new contentInfo and sets the default values.
func NewContentInfo()(*ContentInfo) {
    m := &ContentInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateContentInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateContentInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewContentInfo(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ContentInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetContentFormat gets the contentFormat property value. The contentFormat property
func (m *ContentInfo) GetContentFormat()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contentFormat
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ContentInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["contentFormat"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentFormat(val)
        }
        return nil
    }
    res["identifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentifier(val)
        }
        return nil
    }
    res["metadata"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateKeyValuePairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]KeyValuePairable, len(val))
            for i, v := range val {
                res[i] = v.(KeyValuePairable)
            }
            m.SetMetadata(res)
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseContentState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*ContentState))
        }
        return nil
    }
    return res
}
// GetIdentifier gets the identifier property value. The identifier property
func (m *ContentInfo) GetIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.identifier
    }
}
// GetMetadata gets the metadata property value. The metadata property
func (m *ContentInfo) GetMetadata()([]KeyValuePairable) {
    if m == nil {
        return nil
    } else {
        return m.metadata
    }
}
// GetState gets the state property value. The state property
func (m *ContentInfo) GetState()(*ContentState) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// Serialize serializes information the current object
func (m *ContentInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("contentFormat", m.GetContentFormat())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("identifier", m.GetIdentifier())
        if err != nil {
            return err
        }
    }
    if m.GetMetadata() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMetadata()))
        for i, v := range m.GetMetadata() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("metadata", cast)
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err := writer.WriteStringValue("state", &cast)
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
func (m *ContentInfo) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetContentFormat sets the contentFormat property value. The contentFormat property
func (m *ContentInfo) SetContentFormat(value *string)() {
    if m != nil {
        m.contentFormat = value
    }
}
// SetIdentifier sets the identifier property value. The identifier property
func (m *ContentInfo) SetIdentifier(value *string)() {
    if m != nil {
        m.identifier = value
    }
}
// SetMetadata sets the metadata property value. The metadata property
func (m *ContentInfo) SetMetadata(value []KeyValuePairable)() {
    if m != nil {
        m.metadata = value
    }
}
// SetState sets the state property value. The state property
func (m *ContentInfo) SetState(value *ContentState)() {
    if m != nil {
        m.state = value
    }
}
