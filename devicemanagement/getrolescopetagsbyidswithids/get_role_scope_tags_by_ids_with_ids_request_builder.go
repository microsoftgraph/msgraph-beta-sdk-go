package getrolescopetagsbyidswithids

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// GetRoleScopeTagsByIdsWithIdsRequestBuilder provides operations to call the getRoleScopeTagsByIds method.
type GetRoleScopeTagsByIdsWithIdsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GetRoleScopeTagsByIdsWithIdsRequestBuilderGetOptions options for Get
type GetRoleScopeTagsByIdsWithIdsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *GetRoleScopeTagsByIdsWithIdsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// GetRoleScopeTagsByIdsWithIdsRequestBuilderGetQueryParameters invoke function getRoleScopeTagsByIds
type GetRoleScopeTagsByIdsWithIdsRequestBuilderGetQueryParameters struct {
    // Usage: ids={ids}
    Ids *string;
}
// NewGetRoleScopeTagsByIdsWithIdsRequestBuilderInternal instantiates a new GetRoleScopeTagsByIdsWithIdsRequestBuilder and sets the default values.
func NewGetRoleScopeTagsByIdsWithIdsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetRoleScopeTagsByIdsWithIdsRequestBuilder) {
    m := &GetRoleScopeTagsByIdsWithIdsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/microsoft.graph.getRoleScopeTagsByIds(ids=@ids){?ids}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetRoleScopeTagsByIdsWithIdsRequestBuilder instantiates a new GetRoleScopeTagsByIdsWithIdsRequestBuilder and sets the default values.
func NewGetRoleScopeTagsByIdsWithIdsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetRoleScopeTagsByIdsWithIdsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetRoleScopeTagsByIdsWithIdsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation invoke function getRoleScopeTagsByIds
func (m *GetRoleScopeTagsByIdsWithIdsRequestBuilder) CreateGetRequestInformation(options *GetRoleScopeTagsByIdsWithIdsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get invoke function getRoleScopeTagsByIds
func (m *GetRoleScopeTagsByIdsWithIdsRequestBuilder) Get(options *GetRoleScopeTagsByIdsWithIdsRequestBuilderGetOptions)(GetRoleScopeTagsByIdsWithIdsResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetRoleScopeTagsByIdsWithIdsResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetRoleScopeTagsByIdsWithIdsResponseable), nil
}
