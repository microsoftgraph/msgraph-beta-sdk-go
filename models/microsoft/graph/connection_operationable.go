package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i3af76fce9a0d8c03f22ff90ccd64c93d01bbef0102a1c4e80376e26d2e22a367 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/externalconnectors"
)

// ConnectionOperationable 
type ConnectionOperationable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetError()(PublicErrorable)
    GetStatus()(*i3af76fce9a0d8c03f22ff90ccd64c93d01bbef0102a1c4e80376e26d2e22a367.ConnectionOperationStatus)
    SetError(value PublicErrorable)()
    SetStatus(value *i3af76fce9a0d8c03f22ff90ccd64c93d01bbef0102a1c4e80376e26d2e22a367.ConnectionOperationStatus)()
}
