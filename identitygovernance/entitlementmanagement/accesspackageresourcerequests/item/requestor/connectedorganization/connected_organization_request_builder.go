package connectedorganization

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    id96d92ea1dc9f03386a6fbed0689baaabc95177a07552b3762839a96c2aff0a0 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageresourcerequests/item/requestor/connectedorganization/ref"
)

// ConnectedOrganizationRequestBuilder builds and executes requests for operations under \identityGovernance\entitlementManagement\accessPackageResourceRequests\{accessPackageResourceRequest-id}\requestor\connectedOrganization
type ConnectedOrganizationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ConnectedOrganizationRequestBuilderGetOptions options for Get
type ConnectedOrganizationRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ConnectedOrganizationRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ConnectedOrganizationRequestBuilderGetQueryParameters the connected organization of the subject. Read-only. Nullable.
type ConnectedOrganizationRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// NewConnectedOrganizationRequestBuilderInternal instantiates a new ConnectedOrganizationRequestBuilder and sets the default values.
func NewConnectedOrganizationRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ConnectedOrganizationRequestBuilder) {
    m := &ConnectedOrganizationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/accessPackageResourceRequests/{accessPackageResourceRequest_id}/requestor/connectedOrganization{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewConnectedOrganizationRequestBuilder instantiates a new ConnectedOrganizationRequestBuilder and sets the default values.
func NewConnectedOrganizationRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ConnectedOrganizationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConnectedOrganizationRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation the connected organization of the subject. Read-only. Nullable.
func (m *ConnectedOrganizationRequestBuilder) CreateGetRequestInformation(options *ConnectedOrganizationRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get the connected organization of the subject. Read-only. Nullable.
func (m *ConnectedOrganizationRequestBuilder) Get(options *ConnectedOrganizationRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ConnectedOrganization, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewConnectedOrganization() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ConnectedOrganization), nil
}
func (m *ConnectedOrganizationRequestBuilder) Ref()(*id96d92ea1dc9f03386a6fbed0689baaabc95177a07552b3762839a96c2aff0a0.RefRequestBuilder) {
    return id96d92ea1dc9f03386a6fbed0689baaabc95177a07552b3762839a96c2aff0a0.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
