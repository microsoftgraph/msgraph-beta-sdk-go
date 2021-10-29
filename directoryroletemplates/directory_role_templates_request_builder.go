package directoryroletemplates

import (
    i511bfb2a85de68e20bc06bcac26d677c94a462d0cde3590a23a44f1b87fb080e "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryroletemplates/validateproperties"
    i88743b6fef8c46b4b91459762f36ac5865d91e9f7d82bba7253d32e67ee18285 "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryroletemplates/getuserownedobjects"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    ife18911f94e8d446d0a26f38b8cd124d5a68379523a1dc16ae8198cdf21a7152 "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryroletemplates/getbyids"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// Builds and executes requests for operations under \directoryRoleTemplates
type DirectoryRoleTemplatesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Get
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
// Get entities from directoryRoleTemplates
type DirectoryRoleTemplatesRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
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
    Select_escaped []string;
    // Skip the first n items
    Skip *int32;
}
// Options for Post
type DirectoryRoleTemplatesRequestBuilderPostOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DirectoryRoleTemplate;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Instantiates a new DirectoryRoleTemplatesRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewDirectoryRoleTemplatesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DirectoryRoleTemplatesRequestBuilder) {
    m := &DirectoryRoleTemplatesRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/directoryRoleTemplates{?skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new DirectoryRoleTemplatesRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewDirectoryRoleTemplatesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DirectoryRoleTemplatesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryRoleTemplatesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get entities from directoryRoleTemplates
// Parameters:
//  - options : Options for the request
func (m *DirectoryRoleTemplatesRequestBuilder) CreateGetRequestInformation(options *DirectoryRoleTemplatesRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Add new entity to directoryRoleTemplates
// Parameters:
//  - options : Options for the request
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
// Get entities from directoryRoleTemplates
// Parameters:
//  - options : Options for the request
func (m *DirectoryRoleTemplatesRequestBuilder) Get(options *DirectoryRoleTemplatesRequestBuilderGetOptions)(*DirectoryRoleTemplatesResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryRoleTemplatesResponse() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*DirectoryRoleTemplatesResponse), nil
}
func (m *DirectoryRoleTemplatesRequestBuilder) GetByIds()(*ife18911f94e8d446d0a26f38b8cd124d5a68379523a1dc16ae8198cdf21a7152.GetByIdsRequestBuilder) {
    return ife18911f94e8d446d0a26f38b8cd124d5a68379523a1dc16ae8198cdf21a7152.NewGetByIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DirectoryRoleTemplatesRequestBuilder) GetUserOwnedObjects()(*i88743b6fef8c46b4b91459762f36ac5865d91e9f7d82bba7253d32e67ee18285.GetUserOwnedObjectsRequestBuilder) {
    return i88743b6fef8c46b4b91459762f36ac5865d91e9f7d82bba7253d32e67ee18285.NewGetUserOwnedObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Add new entity to directoryRoleTemplates
// Parameters:
//  - options : Options for the request
func (m *DirectoryRoleTemplatesRequestBuilder) Post(options *DirectoryRoleTemplatesRequestBuilderPostOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DirectoryRoleTemplate, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDirectoryRoleTemplate() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DirectoryRoleTemplate), nil
}
func (m *DirectoryRoleTemplatesRequestBuilder) ValidateProperties()(*i511bfb2a85de68e20bc06bcac26d677c94a462d0cde3590a23a44f1b87fb080e.ValidatePropertiesRequestBuilder) {
    return i511bfb2a85de68e20bc06bcac26d677c94a462d0cde3590a23a44f1b87fb080e.NewValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
