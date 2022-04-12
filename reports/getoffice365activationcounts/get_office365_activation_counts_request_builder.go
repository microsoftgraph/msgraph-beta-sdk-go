package getoffice365activationcounts

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetOffice365ActivationCountsRequestBuilder provides operations to call the getOffice365ActivationCounts method.
type GetOffice365ActivationCountsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetOffice365ActivationCountsRequestBuilderGetOptions options for Get
type GetOffice365ActivationCountsRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// NewGetOffice365ActivationCountsRequestBuilderInternal instantiates a new GetOffice365ActivationCountsRequestBuilder and sets the default values.
func NewGetOffice365ActivationCountsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetOffice365ActivationCountsRequestBuilder) {
    m := &GetOffice365ActivationCountsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/reports/microsoft.graph.getOffice365ActivationCounts()";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetOffice365ActivationCountsRequestBuilder instantiates a new GetOffice365ActivationCountsRequestBuilder and sets the default values.
func NewGetOffice365ActivationCountsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetOffice365ActivationCountsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetOffice365ActivationCountsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation invoke function getOffice365ActivationCounts
func (m *GetOffice365ActivationCountsRequestBuilder) CreateGetRequestInformation(options *GetOffice365ActivationCountsRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Get invoke function getOffice365ActivationCounts
func (m *GetOffice365ActivationCountsRequestBuilder) Get(options *GetOffice365ActivationCountsRequestBuilderGetOptions)(GetOffice365ActivationCountsResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetOffice365ActivationCountsResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetOffice365ActivationCountsResponseable), nil
}
