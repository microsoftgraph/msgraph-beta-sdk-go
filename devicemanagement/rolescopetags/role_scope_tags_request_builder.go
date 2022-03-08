package rolescopetags

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i33808f02abca01e1267c66ae1677f591fa5a4d6e6c99557066e7ad9ab5bdeb72 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/rolescopetags/getrolescopetagsbyid"
    i52487895386ac5b9b67d47fdade3472c4d5e919e1c2bc876ca92efd882afe9ae "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/rolescopetags/hascustomrolescopetag"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i87e24dbe25af599aa1e532c37a185b0fc3836b1f20d55a440891f8edc915a488 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/rolescopetags/count"
)

// RoleScopeTagsRequestBuilder provides operations to manage the roleScopeTags property of the microsoft.graph.deviceManagement entity.
type RoleScopeTagsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// RoleScopeTagsRequestBuilderGetOptions options for Get
type RoleScopeTagsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *RoleScopeTagsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// RoleScopeTagsRequestBuilderGetQueryParameters the Role Scope Tags.
type RoleScopeTagsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool;
    // Expand related entities
    Expand []string;
    // Filter items by property values
    Filter *string;
    // Order items by property values
    Orderby []string;
    // Search items by search phrases
    Search *string;
    // Select properties to be returned
    Select []string;
    // Skip the first n items
    Skip *int32;
    // Show only the first n items
    Top *int32;
}
// RoleScopeTagsRequestBuilderPostOptions options for Post
type RoleScopeTagsRequestBuilderPostOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.RoleScopeTagable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewRoleScopeTagsRequestBuilderInternal instantiates a new RoleScopeTagsRequestBuilder and sets the default values.
func NewRoleScopeTagsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RoleScopeTagsRequestBuilder) {
    m := &RoleScopeTagsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/roleScopeTags{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRoleScopeTagsRequestBuilder instantiates a new RoleScopeTagsRequestBuilder and sets the default values.
func NewRoleScopeTagsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RoleScopeTagsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRoleScopeTagsRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *RoleScopeTagsRequestBuilder) Count()(*i87e24dbe25af599aa1e532c37a185b0fc3836b1f20d55a440891f8edc915a488.CountRequestBuilder) {
    return i87e24dbe25af599aa1e532c37a185b0fc3836b1f20d55a440891f8edc915a488.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the Role Scope Tags.
func (m *RoleScopeTagsRequestBuilder) CreateGetRequestInformation(options *RoleScopeTagsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to roleScopeTags for deviceManagement
func (m *RoleScopeTagsRequestBuilder) CreatePostRequestInformation(options *RoleScopeTagsRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.POST
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
// Get the Role Scope Tags.
func (m *RoleScopeTagsRequestBuilder) Get(options *RoleScopeTagsRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.RoleScopeTagCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateRoleScopeTagCollectionResponseFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.RoleScopeTagCollectionResponseable), nil
}
func (m *RoleScopeTagsRequestBuilder) GetRoleScopeTagsById()(*i33808f02abca01e1267c66ae1677f591fa5a4d6e6c99557066e7ad9ab5bdeb72.GetRoleScopeTagsByIdRequestBuilder) {
    return i33808f02abca01e1267c66ae1677f591fa5a4d6e6c99557066e7ad9ab5bdeb72.NewGetRoleScopeTagsByIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// HasCustomRoleScopeTag provides operations to call the hasCustomRoleScopeTag method.
func (m *RoleScopeTagsRequestBuilder) HasCustomRoleScopeTag()(*i52487895386ac5b9b67d47fdade3472c4d5e919e1c2bc876ca92efd882afe9ae.HasCustomRoleScopeTagRequestBuilder) {
    return i52487895386ac5b9b67d47fdade3472c4d5e919e1c2bc876ca92efd882afe9ae.NewHasCustomRoleScopeTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Post create new navigation property to roleScopeTags for deviceManagement
func (m *RoleScopeTagsRequestBuilder) Post(options *RoleScopeTagsRequestBuilderPostOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.RoleScopeTagable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateRoleScopeTagFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.RoleScopeTagable), nil
}
