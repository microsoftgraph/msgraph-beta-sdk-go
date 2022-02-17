package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i18a6924091a3c9bbbe468505115ea5a5d8a057f434df414584dbb1569657989a "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/serviceannouncement/healthoverviews/item/issues"
    i92fa557a83a3a5f72f390d4ba2371036b32f13a2df0fcd47d44a7d53c414e7d8 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/serviceannouncement/healthoverviews/item/issues/item"
)

// ServiceHealthRequestBuilder builds and executes requests for operations under \admin\serviceAnnouncement\healthOverviews\{serviceHealth-id}
type ServiceHealthRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ServiceHealthRequestBuilderDeleteOptions options for Delete
type ServiceHealthRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ServiceHealthRequestBuilderGetOptions options for Get
type ServiceHealthRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ServiceHealthRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ServiceHealthRequestBuilderGetQueryParameters a collection of service health information for tenant. This property is a contained navigation property, it is nullable and readonly.
type ServiceHealthRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ServiceHealthRequestBuilderPatchOptions options for Patch
type ServiceHealthRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ServiceHealth;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewServiceHealthRequestBuilderInternal instantiates a new ServiceHealthRequestBuilder and sets the default values.
func NewServiceHealthRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ServiceHealthRequestBuilder) {
    m := &ServiceHealthRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/admin/serviceAnnouncement/healthOverviews/{serviceHealth_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewServiceHealthRequestBuilder instantiates a new ServiceHealthRequestBuilder and sets the default values.
func NewServiceHealthRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ServiceHealthRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewServiceHealthRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation a collection of service health information for tenant. This property is a contained navigation property, it is nullable and readonly.
func (m *ServiceHealthRequestBuilder) CreateDeleteRequestInformation(options *ServiceHealthRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation a collection of service health information for tenant. This property is a contained navigation property, it is nullable and readonly.
func (m *ServiceHealthRequestBuilder) CreateGetRequestInformation(options *ServiceHealthRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation a collection of service health information for tenant. This property is a contained navigation property, it is nullable and readonly.
func (m *ServiceHealthRequestBuilder) CreatePatchRequestInformation(options *ServiceHealthRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete a collection of service health information for tenant. This property is a contained navigation property, it is nullable and readonly.
func (m *ServiceHealthRequestBuilder) Delete(options *ServiceHealthRequestBuilderDeleteOptions)(error) {
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
// Get a collection of service health information for tenant. This property is a contained navigation property, it is nullable and readonly.
func (m *ServiceHealthRequestBuilder) Get(options *ServiceHealthRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ServiceHealth, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewServiceHealth() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ServiceHealth), nil
}
func (m *ServiceHealthRequestBuilder) Issues()(*i18a6924091a3c9bbbe468505115ea5a5d8a057f434df414584dbb1569657989a.IssuesRequestBuilder) {
    return i18a6924091a3c9bbbe468505115ea5a5d8a057f434df414584dbb1569657989a.NewIssuesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IssuesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.admin.serviceAnnouncement.healthOverviews.item.issues.item collection
func (m *ServiceHealthRequestBuilder) IssuesById(id string)(*i92fa557a83a3a5f72f390d4ba2371036b32f13a2df0fcd47d44a7d53c414e7d8.ServiceHealthIssueRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["serviceHealthIssue_id"] = id
    }
    return i92fa557a83a3a5f72f390d4ba2371036b32f13a2df0fcd47d44a7d53c414e7d8.NewServiceHealthIssueRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch a collection of service health information for tenant. This property is a contained navigation property, it is nullable and readonly.
func (m *ServiceHealthRequestBuilder) Patch(options *ServiceHealthRequestBuilderPatchOptions)(error) {
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
