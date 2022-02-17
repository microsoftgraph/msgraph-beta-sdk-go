package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i3033e6924fd7bdba5cc472a9ef874d137078ea8e323da669089145de32632ec9 "github.com/microsoftgraph/msgraph-beta-sdk-go/onpremisespublishingprofiles/item/connectorgroups/item/members"
    i65f60af5b24192c2cf36c30b834edfc25784e306f7de673f8c1129eccf89bcdc "github.com/microsoftgraph/msgraph-beta-sdk-go/onpremisespublishingprofiles/item/connectorgroups/item/applications"
    i753cfff90925bea726b5ebc5e71149db3f766477cada75985f673e10929c55ae "github.com/microsoftgraph/msgraph-beta-sdk-go/onpremisespublishingprofiles/item/connectorgroups/item/members/item"
)

// ConnectorGroupRequestBuilder builds and executes requests for operations under \onPremisesPublishingProfiles\{onPremisesPublishingProfile-id}\connectorGroups\{connectorGroup-id}
type ConnectorGroupRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ConnectorGroupRequestBuilderDeleteOptions options for Delete
type ConnectorGroupRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ConnectorGroupRequestBuilderGetOptions options for Get
type ConnectorGroupRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ConnectorGroupRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ConnectorGroupRequestBuilderGetQueryParameters list of existing connectorGroup objects for applications published through Application Proxy. Read-only. Nullable.
type ConnectorGroupRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ConnectorGroupRequestBuilderPatchOptions options for Patch
type ConnectorGroupRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ConnectorGroup;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *ConnectorGroupRequestBuilder) Applications()(*i65f60af5b24192c2cf36c30b834edfc25784e306f7de673f8c1129eccf89bcdc.ApplicationsRequestBuilder) {
    return i65f60af5b24192c2cf36c30b834edfc25784e306f7de673f8c1129eccf89bcdc.NewApplicationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewConnectorGroupRequestBuilderInternal instantiates a new ConnectorGroupRequestBuilder and sets the default values.
func NewConnectorGroupRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ConnectorGroupRequestBuilder) {
    m := &ConnectorGroupRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/onPremisesPublishingProfiles/{onPremisesPublishingProfile_id}/connectorGroups/{connectorGroup_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewConnectorGroupRequestBuilder instantiates a new ConnectorGroupRequestBuilder and sets the default values.
func NewConnectorGroupRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ConnectorGroupRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConnectorGroupRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation list of existing connectorGroup objects for applications published through Application Proxy. Read-only. Nullable.
func (m *ConnectorGroupRequestBuilder) CreateDeleteRequestInformation(options *ConnectorGroupRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation list of existing connectorGroup objects for applications published through Application Proxy. Read-only. Nullable.
func (m *ConnectorGroupRequestBuilder) CreateGetRequestInformation(options *ConnectorGroupRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation list of existing connectorGroup objects for applications published through Application Proxy. Read-only. Nullable.
func (m *ConnectorGroupRequestBuilder) CreatePatchRequestInformation(options *ConnectorGroupRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete list of existing connectorGroup objects for applications published through Application Proxy. Read-only. Nullable.
func (m *ConnectorGroupRequestBuilder) Delete(options *ConnectorGroupRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get list of existing connectorGroup objects for applications published through Application Proxy. Read-only. Nullable.
func (m *ConnectorGroupRequestBuilder) Get(options *ConnectorGroupRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ConnectorGroup, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewConnectorGroup() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ConnectorGroup), nil
}
func (m *ConnectorGroupRequestBuilder) Members()(*i3033e6924fd7bdba5cc472a9ef874d137078ea8e323da669089145de32632ec9.MembersRequestBuilder) {
    return i3033e6924fd7bdba5cc472a9ef874d137078ea8e323da669089145de32632ec9.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.onPremisesPublishingProfiles.item.connectorGroups.item.members.item collection
func (m *ConnectorGroupRequestBuilder) MembersById(id string)(*i753cfff90925bea726b5ebc5e71149db3f766477cada75985f673e10929c55ae.ConnectorRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["connector_id"] = id
    }
    return i753cfff90925bea726b5ebc5e71149db3f766477cada75985f673e10929c55ae.NewConnectorRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch list of existing connectorGroup objects for applications published through Application Proxy. Read-only. Nullable.
func (m *ConnectorGroupRequestBuilder) Patch(options *ConnectorGroupRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
