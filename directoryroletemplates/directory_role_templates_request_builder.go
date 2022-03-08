package directoryroletemplates

import (
    i3f8034c8f6f2b58468d7d887066b5af2d36cab6daa3877702f6e2742abf25509 "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryroletemplates/count"
    i511bfb2a85de68e20bc06bcac26d677c94a462d0cde3590a23a44f1b87fb080e "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryroletemplates/validateproperties"
    i88743b6fef8c46b4b91459762f36ac5865d91e9f7d82bba7253d32e67ee18285 "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryroletemplates/getuserownedobjects"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    ife18911f94e8d446d0a26f38b8cd124d5a68379523a1dc16ae8198cdf21a7152 "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryroletemplates/getbyids"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// DirectoryRoleTemplatesRequestBuilder provides operations to manage the collection of directoryRoleTemplate entities.
type DirectoryRoleTemplatesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DirectoryRoleTemplatesRequestBuilderGetOptions options for Get
type DirectoryRoleTemplatesRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DirectoryRoleTemplatesRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DirectoryRoleTemplatesRequestBuilderGetQueryParameters get entities from directoryRoleTemplates
type DirectoryRoleTemplatesRequestBuilderGetQueryParameters struct {
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
}
// DirectoryRoleTemplatesRequestBuilderPostOptions options for Post
type DirectoryRoleTemplatesRequestBuilderPostOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DirectoryRoleTemplateable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewDirectoryRoleTemplatesRequestBuilderInternal instantiates a new DirectoryRoleTemplatesRequestBuilder and sets the default values.
func NewDirectoryRoleTemplatesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DirectoryRoleTemplatesRequestBuilder) {
    m := &DirectoryRoleTemplatesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/directoryRoleTemplates{?skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDirectoryRoleTemplatesRequestBuilder instantiates a new DirectoryRoleTemplatesRequestBuilder and sets the default values.
func NewDirectoryRoleTemplatesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DirectoryRoleTemplatesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryRoleTemplatesRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *DirectoryRoleTemplatesRequestBuilder) Count()(*i3f8034c8f6f2b58468d7d887066b5af2d36cab6daa3877702f6e2742abf25509.CountRequestBuilder) {
    return i3f8034c8f6f2b58468d7d887066b5af2d36cab6daa3877702f6e2742abf25509.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation get entities from directoryRoleTemplates
func (m *DirectoryRoleTemplatesRequestBuilder) CreateGetRequestInformation(options *DirectoryRoleTemplatesRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePostRequestInformation add new entity to directoryRoleTemplates
func (m *DirectoryRoleTemplatesRequestBuilder) CreatePostRequestInformation(options *DirectoryRoleTemplatesRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get get entities from directoryRoleTemplates
func (m *DirectoryRoleTemplatesRequestBuilder) Get(options *DirectoryRoleTemplatesRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DirectoryRoleTemplateCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateDirectoryRoleTemplateCollectionResponseFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DirectoryRoleTemplateCollectionResponseable), nil
}
func (m *DirectoryRoleTemplatesRequestBuilder) GetByIds()(*ife18911f94e8d446d0a26f38b8cd124d5a68379523a1dc16ae8198cdf21a7152.GetByIdsRequestBuilder) {
    return ife18911f94e8d446d0a26f38b8cd124d5a68379523a1dc16ae8198cdf21a7152.NewGetByIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DirectoryRoleTemplatesRequestBuilder) GetUserOwnedObjects()(*i88743b6fef8c46b4b91459762f36ac5865d91e9f7d82bba7253d32e67ee18285.GetUserOwnedObjectsRequestBuilder) {
    return i88743b6fef8c46b4b91459762f36ac5865d91e9f7d82bba7253d32e67ee18285.NewGetUserOwnedObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Post add new entity to directoryRoleTemplates
func (m *DirectoryRoleTemplatesRequestBuilder) Post(options *DirectoryRoleTemplatesRequestBuilderPostOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DirectoryRoleTemplateable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateDirectoryRoleTemplateFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DirectoryRoleTemplateable), nil
}
func (m *DirectoryRoleTemplatesRequestBuilder) ValidateProperties()(*i511bfb2a85de68e20bc06bcac26d677c94a462d0cde3590a23a44f1b87fb080e.ValidatePropertiesRequestBuilder) {
    return i511bfb2a85de68e20bc06bcac26d677c94a462d0cde3590a23a44f1b87fb080e.NewValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
