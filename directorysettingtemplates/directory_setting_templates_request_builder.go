package directorysettingtemplates

import (
    i03094d9d1021b32e492d609c092d84ed4cfb558c32f0c8ade3ad08ceb1eead66 "github.com/microsoftgraph/msgraph-beta-sdk-go/directorysettingtemplates/validateproperties"
    i8b3159fe7053462b6a5b7687099a8e653ffaaa3e49ac1c52748ad001ac05d4d9 "github.com/microsoftgraph/msgraph-beta-sdk-go/directorysettingtemplates/getbyids"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    if2eec762f4a60cfa480122020666674cae3dd0c0dd67417b7da42d6d7d9efc9f "github.com/microsoftgraph/msgraph-beta-sdk-go/directorysettingtemplates/getuserownedobjects"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// DirectorySettingTemplatesRequestBuilder builds and executes requests for operations under \directorySettingTemplates
type DirectorySettingTemplatesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DirectorySettingTemplatesRequestBuilderGetOptions options for Get
type DirectorySettingTemplatesRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DirectorySettingTemplatesRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DirectorySettingTemplatesRequestBuilderGetQueryParameters get entities from directorySettingTemplates
type DirectorySettingTemplatesRequestBuilderGetQueryParameters struct {
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
// DirectorySettingTemplatesRequestBuilderPostOptions options for Post
type DirectorySettingTemplatesRequestBuilderPostOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DirectorySettingTemplate;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewDirectorySettingTemplatesRequestBuilderInternal instantiates a new DirectorySettingTemplatesRequestBuilder and sets the default values.
func NewDirectorySettingTemplatesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DirectorySettingTemplatesRequestBuilder) {
    m := &DirectorySettingTemplatesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/directorySettingTemplates{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDirectorySettingTemplatesRequestBuilder instantiates a new DirectorySettingTemplatesRequestBuilder and sets the default values.
func NewDirectorySettingTemplatesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DirectorySettingTemplatesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectorySettingTemplatesRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get entities from directorySettingTemplates
func (m *DirectorySettingTemplatesRequestBuilder) CreateGetRequestInformation(options *DirectorySettingTemplatesRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePostRequestInformation add new entity to directorySettingTemplates
func (m *DirectorySettingTemplatesRequestBuilder) CreatePostRequestInformation(options *DirectorySettingTemplatesRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get get entities from directorySettingTemplates
func (m *DirectorySettingTemplatesRequestBuilder) Get(options *DirectorySettingTemplatesRequestBuilderGetOptions)(*DirectorySettingTemplatesResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectorySettingTemplatesResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*DirectorySettingTemplatesResponse), nil
}
func (m *DirectorySettingTemplatesRequestBuilder) GetByIds()(*i8b3159fe7053462b6a5b7687099a8e653ffaaa3e49ac1c52748ad001ac05d4d9.GetByIdsRequestBuilder) {
    return i8b3159fe7053462b6a5b7687099a8e653ffaaa3e49ac1c52748ad001ac05d4d9.NewGetByIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DirectorySettingTemplatesRequestBuilder) GetUserOwnedObjects()(*if2eec762f4a60cfa480122020666674cae3dd0c0dd67417b7da42d6d7d9efc9f.GetUserOwnedObjectsRequestBuilder) {
    return if2eec762f4a60cfa480122020666674cae3dd0c0dd67417b7da42d6d7d9efc9f.NewGetUserOwnedObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Post add new entity to directorySettingTemplates
func (m *DirectorySettingTemplatesRequestBuilder) Post(options *DirectorySettingTemplatesRequestBuilderPostOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DirectorySettingTemplate, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDirectorySettingTemplate() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DirectorySettingTemplate), nil
}
func (m *DirectorySettingTemplatesRequestBuilder) ValidateProperties()(*i03094d9d1021b32e492d609c092d84ed4cfb558c32f0c8ade3ad08ceb1eead66.ValidatePropertiesRequestBuilder) {
    return i03094d9d1021b32e492d609c092d84ed4cfb558c32f0c8ade3ad08ceb1eead66.NewValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
