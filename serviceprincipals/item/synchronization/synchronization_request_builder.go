package synchronization

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i441115fbd107244dd6504672d846ac2adcc602e6da088866d93d998f95107c40 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/synchronization/ping"
    i8d4af922a9d1b5130f294d250ea800e845868ae335b95eccdbe95097a1393573 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/synchronization/templates"
    i9f2b960ccd9a3439ea358b94aa154106d5794fdd4bdd4786b63fb9896b804094 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/synchronization/jobs"
    ie4e7228c8413ba04b04a7f1fac840439946274cff81aadbd7cf5252263fe1cc4 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/synchronization/acquireaccesstoken"
    ibe6466d39a6e0f7d78fef7ec330777772ce9e4942297f090c6e91a7b4f7da1a8 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/synchronization/jobs/item"
    idbe71426e8a7cf17f46df86ea54d1a8d22e20242b9a4e0a5c7d1e8c8405dd581 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/synchronization/templates/item"
)

// SynchronizationRequestBuilder builds and executes requests for operations under \servicePrincipals\{servicePrincipal-id}\synchronization
type SynchronizationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// SynchronizationRequestBuilderDeleteOptions options for Delete
type SynchronizationRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// SynchronizationRequestBuilderGetOptions options for Get
type SynchronizationRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *SynchronizationRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// SynchronizationRequestBuilderGetQueryParameters get synchronization from servicePrincipals
type SynchronizationRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// SynchronizationRequestBuilderPatchOptions options for Patch
type SynchronizationRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Synchronization;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *SynchronizationRequestBuilder) AcquireAccessToken()(*ie4e7228c8413ba04b04a7f1fac840439946274cff81aadbd7cf5252263fe1cc4.AcquireAccessTokenRequestBuilder) {
    return ie4e7228c8413ba04b04a7f1fac840439946274cff81aadbd7cf5252263fe1cc4.NewAcquireAccessTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewSynchronizationRequestBuilderInternal instantiates a new SynchronizationRequestBuilder and sets the default values.
func NewSynchronizationRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SynchronizationRequestBuilder) {
    m := &SynchronizationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/servicePrincipals/{servicePrincipal_id}/synchronization{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSynchronizationRequestBuilder instantiates a new SynchronizationRequestBuilder and sets the default values.
func NewSynchronizationRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SynchronizationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSynchronizationRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property synchronization for servicePrincipals
func (m *SynchronizationRequestBuilder) CreateDeleteRequestInformation(options *SynchronizationRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get synchronization from servicePrincipals
func (m *SynchronizationRequestBuilder) CreateGetRequestInformation(options *SynchronizationRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property synchronization in servicePrincipals
func (m *SynchronizationRequestBuilder) CreatePatchRequestInformation(options *SynchronizationRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property synchronization for servicePrincipals
func (m *SynchronizationRequestBuilder) Delete(options *SynchronizationRequestBuilderDeleteOptions)(error) {
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
// Get get synchronization from servicePrincipals
func (m *SynchronizationRequestBuilder) Get(options *SynchronizationRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Synchronization, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewSynchronization() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Synchronization), nil
}
func (m *SynchronizationRequestBuilder) Jobs()(*i9f2b960ccd9a3439ea358b94aa154106d5794fdd4bdd4786b63fb9896b804094.JobsRequestBuilder) {
    return i9f2b960ccd9a3439ea358b94aa154106d5794fdd4bdd4786b63fb9896b804094.NewJobsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// JobsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.servicePrincipals.item.synchronization.jobs.item collection
func (m *SynchronizationRequestBuilder) JobsById(id string)(*ibe6466d39a6e0f7d78fef7ec330777772ce9e4942297f090c6e91a7b4f7da1a8.SynchronizationJobRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["synchronizationJob_id"] = id
    }
    return ibe6466d39a6e0f7d78fef7ec330777772ce9e4942297f090c6e91a7b4f7da1a8.NewSynchronizationJobRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property synchronization in servicePrincipals
func (m *SynchronizationRequestBuilder) Patch(options *SynchronizationRequestBuilderPatchOptions)(error) {
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
// Ping builds and executes requests for operations under \servicePrincipals\{servicePrincipal-id}\synchronization\microsoft.graph.Ping()
func (m *SynchronizationRequestBuilder) Ping()(*i441115fbd107244dd6504672d846ac2adcc602e6da088866d93d998f95107c40.PingRequestBuilder) {
    return i441115fbd107244dd6504672d846ac2adcc602e6da088866d93d998f95107c40.NewPingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SynchronizationRequestBuilder) Templates()(*i8d4af922a9d1b5130f294d250ea800e845868ae335b95eccdbe95097a1393573.TemplatesRequestBuilder) {
    return i8d4af922a9d1b5130f294d250ea800e845868ae335b95eccdbe95097a1393573.NewTemplatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TemplatesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.servicePrincipals.item.synchronization.templates.item collection
func (m *SynchronizationRequestBuilder) TemplatesById(id string)(*idbe71426e8a7cf17f46df86ea54d1a8d22e20242b9a4e0a5c7d1e8c8405dd581.SynchronizationTemplateRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["synchronizationTemplate_id"] = id
    }
    return idbe71426e8a7cf17f46df86ea54d1a8d22e20242b9a4e0a5c7d1e8c8405dd581.NewSynchronizationTemplateRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
