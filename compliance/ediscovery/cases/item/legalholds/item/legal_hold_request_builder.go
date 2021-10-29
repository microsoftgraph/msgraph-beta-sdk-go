package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i06b275b20ff5a31845bd4dd9892fec2b1653544d3b2bd2a68090be24b2058315 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/legalholds/item/usersources"
    i64f0e5258633a3c0c8d5bfc297b4e0736de72c5e3150a38d2fcd17dfb59e046b "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/legalholds/item/sitesources"
    iac7d03a0b596e8feddf3b195648c1f8d07225fc5a621d2181a8d3e15d43673a8 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/legalholds/item/unifiedgroupsources"
    i6b88ba94cf868f1ba6788ff7bfac0991479b8eebedef91a91c1b16be7ca6d060 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/legalholds/item/usersources/item"
    ib34df5322945e42aea8a01751883e193d59745ac0c447872f1671eb484a06936 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/legalholds/item/sitesources/item"
    ifaf28029d98f7bf2244cb8a85e111bec3c46d219a6d9069d26a26ee3797ad760 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/legalholds/item/unifiedgroupsources/item"
)

// Builds and executes requests for operations under \compliance\ediscovery\cases\{case-id}\legalHolds\{legalHold-id}
type LegalHoldRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type LegalHoldRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type LegalHoldRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *LegalHoldRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Returns a list of case legalHold objects for this case.  Nullable.
type LegalHoldRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type LegalHoldRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.LegalHold;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Instantiates a new LegalHoldRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewLegalHoldRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*LegalHoldRequestBuilder) {
    m := &LegalHoldRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/compliance/ediscovery/cases/{case_id}/legalHolds/{legalHold_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new LegalHoldRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewLegalHoldRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*LegalHoldRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLegalHoldRequestBuilderInternal(urlParams, requestAdapter)
}
// Returns a list of case legalHold objects for this case.  Nullable.
// Parameters:
//  - options : Options for the request
func (m *LegalHoldRequestBuilder) CreateDeleteRequestInformation(options *LegalHoldRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Returns a list of case legalHold objects for this case.  Nullable.
// Parameters:
//  - options : Options for the request
func (m *LegalHoldRequestBuilder) CreateGetRequestInformation(options *LegalHoldRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        err := options.Q.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
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
// Returns a list of case legalHold objects for this case.  Nullable.
// Parameters:
//  - options : Options for the request
func (m *LegalHoldRequestBuilder) CreatePatchRequestInformation(options *LegalHoldRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Returns a list of case legalHold objects for this case.  Nullable.
// Parameters:
//  - options : Options for the request
func (m *LegalHoldRequestBuilder) Delete(options *LegalHoldRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// Returns a list of case legalHold objects for this case.  Nullable.
// Parameters:
//  - options : Options for the request
func (m *LegalHoldRequestBuilder) Get(options *LegalHoldRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.LegalHold, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewLegalHold() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.LegalHold), nil
}
// Returns a list of case legalHold objects for this case.  Nullable.
// Parameters:
//  - options : Options for the request
func (m *LegalHoldRequestBuilder) Patch(options *LegalHoldRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *LegalHoldRequestBuilder) SiteSources()(*i64f0e5258633a3c0c8d5bfc297b4e0736de72c5e3150a38d2fcd17dfb59e046b.SiteSourcesRequestBuilder) {
    return i64f0e5258633a3c0c8d5bfc297b4e0736de72c5e3150a38d2fcd17dfb59e046b.NewSiteSourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go.compliance.ediscovery.cases.item.legalHolds.item.siteSources.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *LegalHoldRequestBuilder) SiteSourcesById(id string)(*ib34df5322945e42aea8a01751883e193d59745ac0c447872f1671eb484a06936.SiteSourceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["siteSource_id"] = id
    }
    return ib34df5322945e42aea8a01751883e193d59745ac0c447872f1671eb484a06936.NewSiteSourceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *LegalHoldRequestBuilder) UnifiedGroupSources()(*iac7d03a0b596e8feddf3b195648c1f8d07225fc5a621d2181a8d3e15d43673a8.UnifiedGroupSourcesRequestBuilder) {
    return iac7d03a0b596e8feddf3b195648c1f8d07225fc5a621d2181a8d3e15d43673a8.NewUnifiedGroupSourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go.compliance.ediscovery.cases.item.legalHolds.item.unifiedGroupSources.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *LegalHoldRequestBuilder) UnifiedGroupSourcesById(id string)(*ifaf28029d98f7bf2244cb8a85e111bec3c46d219a6d9069d26a26ee3797ad760.UnifiedGroupSourceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedGroupSource_id"] = id
    }
    return ifaf28029d98f7bf2244cb8a85e111bec3c46d219a6d9069d26a26ee3797ad760.NewUnifiedGroupSourceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *LegalHoldRequestBuilder) UserSources()(*i06b275b20ff5a31845bd4dd9892fec2b1653544d3b2bd2a68090be24b2058315.UserSourcesRequestBuilder) {
    return i06b275b20ff5a31845bd4dd9892fec2b1653544d3b2bd2a68090be24b2058315.NewUserSourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go.compliance.ediscovery.cases.item.legalHolds.item.userSources.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *LegalHoldRequestBuilder) UserSourcesById(id string)(*i6b88ba94cf868f1ba6788ff7bfac0991479b8eebedef91a91c1b16be7ca6d060.UserSourceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userSource_id"] = id
    }
    return i6b88ba94cf868f1ba6788ff7bfac0991479b8eebedef91a91c1b16be7ca6d060.NewUserSourceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
