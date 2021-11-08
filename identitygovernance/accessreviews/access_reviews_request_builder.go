package accessreviews

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i65800ed62fdf64d2a6b90ada174f2addc07989dabfa5cf284d4e84f60fc2943e "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/accessreviews/policy"
    ia26090068384411e4f1397570769f217531481e6b5d7843821ed1ba88d0bf3f0 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/accessreviews/historydefinitions"
    icdcfab65b3cad0ae0ef323407563924f12617f5594e6a9fb131fba25fa3c0e64 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/accessreviews/definitions"
    i560e03bc7fd4ff380362907aa5dfe4f5629fc0c3ceb16dcb98159089dfc550b7 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/accessreviews/historydefinitions/item"
    i6da4aa0125ca0757188f525aa6935aefda281d0627445ec063c253a0bd2fccb2 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/accessreviews/definitions/item"
)

// Builds and executes requests for operations under \identityGovernance\accessReviews
type AccessReviewsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type AccessReviewsRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type AccessReviewsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *AccessReviewsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Get accessReviews from identityGovernance
type AccessReviewsRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type AccessReviewsRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessReviewSet;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Instantiates a new AccessReviewsRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewAccessReviewsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AccessReviewsRequestBuilder) {
    m := &AccessReviewsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/accessReviews{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new AccessReviewsRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewAccessReviewsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AccessReviewsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessReviewsRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete navigation property accessReviews for identityGovernance
// Parameters:
//  - options : Options for the request
func (m *AccessReviewsRequestBuilder) CreateDeleteRequestInformation(options *AccessReviewsRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get accessReviews from identityGovernance
// Parameters:
//  - options : Options for the request
func (m *AccessReviewsRequestBuilder) CreateGetRequestInformation(options *AccessReviewsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Update the navigation property accessReviews in identityGovernance
// Parameters:
//  - options : Options for the request
func (m *AccessReviewsRequestBuilder) CreatePatchRequestInformation(options *AccessReviewsRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *AccessReviewsRequestBuilder) Definitions()(*icdcfab65b3cad0ae0ef323407563924f12617f5594e6a9fb131fba25fa3c0e64.DefinitionsRequestBuilder) {
    return icdcfab65b3cad0ae0ef323407563924f12617f5594e6a9fb131fba25fa3c0e64.NewDefinitionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go.identityGovernance.accessReviews.definitions.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *AccessReviewsRequestBuilder) DefinitionsById(id string)(*i6da4aa0125ca0757188f525aa6935aefda281d0627445ec063c253a0bd2fccb2.AccessReviewScheduleDefinitionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewScheduleDefinition_id"] = id
    }
    return i6da4aa0125ca0757188f525aa6935aefda281d0627445ec063c253a0bd2fccb2.NewAccessReviewScheduleDefinitionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Delete navigation property accessReviews for identityGovernance
// Parameters:
//  - options : Options for the request
func (m *AccessReviewsRequestBuilder) Delete(options *AccessReviewsRequestBuilderDeleteOptions)(error) {
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
// Get accessReviews from identityGovernance
// Parameters:
//  - options : Options for the request
func (m *AccessReviewsRequestBuilder) Get(options *AccessReviewsRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessReviewSet, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewAccessReviewSet() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessReviewSet), nil
}
func (m *AccessReviewsRequestBuilder) HistoryDefinitions()(*ia26090068384411e4f1397570769f217531481e6b5d7843821ed1ba88d0bf3f0.HistoryDefinitionsRequestBuilder) {
    return ia26090068384411e4f1397570769f217531481e6b5d7843821ed1ba88d0bf3f0.NewHistoryDefinitionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go.identityGovernance.accessReviews.historyDefinitions.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *AccessReviewsRequestBuilder) HistoryDefinitionsById(id string)(*i560e03bc7fd4ff380362907aa5dfe4f5629fc0c3ceb16dcb98159089dfc550b7.AccessReviewHistoryDefinitionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewHistoryDefinition_id"] = id
    }
    return i560e03bc7fd4ff380362907aa5dfe4f5629fc0c3ceb16dcb98159089dfc550b7.NewAccessReviewHistoryDefinitionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Update the navigation property accessReviews in identityGovernance
// Parameters:
//  - options : Options for the request
func (m *AccessReviewsRequestBuilder) Patch(options *AccessReviewsRequestBuilderPatchOptions)(error) {
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
func (m *AccessReviewsRequestBuilder) Policy()(*i65800ed62fdf64d2a6b90ada174f2addc07989dabfa5cf284d4e84f60fc2943e.PolicyRequestBuilder) {
    return i65800ed62fdf64d2a6b90ada174f2addc07989dabfa5cf284d4e84f60fc2943e.NewPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
