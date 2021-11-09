package getrolescopetagsbyidsids=@ids

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Builds and executes requests for operations under \deviceManagement\microsoft.graph.getRoleScopeTagsByIds(ids=@ids)
type GetRoleScopeTagsByIdsids=@idsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Get
type GetRoleScopeTagsByIdsids=@idsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *GetRoleScopeTagsByIdsids=@idsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Invoke function getRoleScopeTagsByIds
type GetRoleScopeTagsByIdsids=@idsRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Usage: ids={ids}
    Ids *string;
}
// Instantiates a new GetRoleScopeTagsByIdsids=@idsRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewGetRoleScopeTagsByIdsids=@idsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetRoleScopeTagsByIdsids=@idsRequestBuilder) {
    m := &GetRoleScopeTagsByIdsids=@idsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/microsoft.graph.getRoleScopeTagsByIds(ids=@ids){?ids}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new GetRoleScopeTagsByIdsids=@idsRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewGetRoleScopeTagsByIdsids=@idsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetRoleScopeTagsByIdsids=@idsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetRoleScopeTagsByIdsids=@idsRequestBuilderInternal(urlParams, requestAdapter)
}
// Invoke function getRoleScopeTagsByIds
// Parameters:
//  - options : Options for the request
func (m *GetRoleScopeTagsByIdsids=@idsRequestBuilder) CreateGetRequestInformation(options *GetRoleScopeTagsByIdsids=@idsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Invoke function getRoleScopeTagsByIds
// Parameters:
//  - options : Options for the request
func (m *GetRoleScopeTagsByIdsids=@idsRequestBuilder) Get(options *GetRoleScopeTagsByIdsids=@idsRequestBuilderGetOptions)([]GetRoleScopeTagsByIdsids=@ids, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendCollectionAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGetRoleScopeTagsByIdsids=@ids() }, nil)
    if err != nil {
        return nil, err
    }
    val := make([]GetRoleScopeTagsByIdsids=@ids, len(res))
    for i, v := range res {
        val[i] = *(v.(*GetRoleScopeTagsByIdsids=@ids))
    }
    return val, nil
}
