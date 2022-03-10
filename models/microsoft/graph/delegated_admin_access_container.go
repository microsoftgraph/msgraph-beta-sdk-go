package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DelegatedAdminAccessContainer provides operations to manage the tenantRelationship singleton.
type DelegatedAdminAccessContainer struct {
    // 
    accessContainerId *string;
    // 
    accessContainerType *DelegatedAdminAccessContainerType;
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
}
// NewDelegatedAdminAccessContainer instantiates a new delegatedAdminAccessContainer and sets the default values.
func NewDelegatedAdminAccessContainer()(*DelegatedAdminAccessContainer) {
    m := &DelegatedAdminAccessContainer{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDelegatedAdminAccessContainerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDelegatedAdminAccessContainerFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewDelegatedAdminAccessContainer(), nil
}
// GetAccessContainerId gets the accessContainerId property value. 
func (m *DelegatedAdminAccessContainer) GetAccessContainerId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.accessContainerId
    }
}
// GetAccessContainerType gets the accessContainerType property value. 
func (m *DelegatedAdminAccessContainer) GetAccessContainerType()(*DelegatedAdminAccessContainerType) {
    if m == nil {
        return nil
    } else {
        return m.accessContainerType
    }
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DelegatedAdminAccessContainer) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DelegatedAdminAccessContainer) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["accessContainerId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessContainerId(val)
        }
        return nil
    }
    res["accessContainerType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDelegatedAdminAccessContainerType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessContainerType(val.(*DelegatedAdminAccessContainerType))
        }
        return nil
    }
    return res
}
func (m *DelegatedAdminAccessContainer) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DelegatedAdminAccessContainer) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("accessContainerId", m.GetAccessContainerId())
        if err != nil {
            return err
        }
    }
    if m.GetAccessContainerType() != nil {
        cast := (*m.GetAccessContainerType()).String()
        err := writer.WriteStringValue("accessContainerType", &cast)
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
// SetAccessContainerId sets the accessContainerId property value. 
func (m *DelegatedAdminAccessContainer) SetAccessContainerId(value *string)() {
    if m != nil {
        m.accessContainerId = value
    }
}
// SetAccessContainerType sets the accessContainerType property value. 
func (m *DelegatedAdminAccessContainer) SetAccessContainerType(value *DelegatedAdminAccessContainerType)() {
    if m != nil {
        m.accessContainerType = value
    }
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DelegatedAdminAccessContainer) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
