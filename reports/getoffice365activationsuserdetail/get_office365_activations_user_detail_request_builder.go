package getoffice365activationsuserdetail

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetOffice365ActivationsUserDetailRequestBuilder provides operations to call the getOffice365ActivationsUserDetail method.
type GetOffice365ActivationsUserDetailRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetOffice365ActivationsUserDetailRequestBuilderGetOptions options for Get
type GetOffice365ActivationsUserDetailRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// NewGetOffice365ActivationsUserDetailRequestBuilderInternal instantiates a new GetOffice365ActivationsUserDetailRequestBuilder and sets the default values.
func NewGetOffice365ActivationsUserDetailRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetOffice365ActivationsUserDetailRequestBuilder) {
    m := &GetOffice365ActivationsUserDetailRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/reports/microsoft.graph.getOffice365ActivationsUserDetail()";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetOffice365ActivationsUserDetailRequestBuilder instantiates a new GetOffice365ActivationsUserDetailRequestBuilder and sets the default values.
func NewGetOffice365ActivationsUserDetailRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetOffice365ActivationsUserDetailRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetOffice365ActivationsUserDetailRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation invoke function getOffice365ActivationsUserDetail
func (m *GetOffice365ActivationsUserDetailRequestBuilder) CreateGetRequestInformation(options *GetOffice365ActivationsUserDetailRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get invoke function getOffice365ActivationsUserDetail
func (m *GetOffice365ActivationsUserDetailRequestBuilder) Get(options *GetOffice365ActivationsUserDetailRequestBuilderGetOptions)(GetOffice365ActivationsUserDetailResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetOffice365ActivationsUserDetailResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetOffice365ActivationsUserDetailResponseable), nil
}
