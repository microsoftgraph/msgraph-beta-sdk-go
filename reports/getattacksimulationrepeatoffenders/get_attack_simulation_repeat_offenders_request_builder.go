package getattacksimulationrepeatoffenders

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetAttackSimulationRepeatOffendersRequestBuilder provides operations to call the getAttackSimulationRepeatOffenders method.
type GetAttackSimulationRepeatOffendersRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetAttackSimulationRepeatOffendersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetAttackSimulationRepeatOffendersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetAttackSimulationRepeatOffendersRequestBuilderInternal instantiates a new GetAttackSimulationRepeatOffendersRequestBuilder and sets the default values.
func NewGetAttackSimulationRepeatOffendersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetAttackSimulationRepeatOffendersRequestBuilder) {
    m := &GetAttackSimulationRepeatOffendersRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/reports/microsoft.graph.getAttackSimulationRepeatOffenders()";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetAttackSimulationRepeatOffendersRequestBuilder instantiates a new GetAttackSimulationRepeatOffendersRequestBuilder and sets the default values.
func NewGetAttackSimulationRepeatOffendersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetAttackSimulationRepeatOffendersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetAttackSimulationRepeatOffendersRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformationWithRequestConfiguration invoke function getAttackSimulationRepeatOffenders
func (m *GetAttackSimulationRepeatOffendersRequestBuilder) CreateGetRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration invoke function getAttackSimulationRepeatOffenders
func (m *GetAttackSimulationRepeatOffendersRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *GetAttackSimulationRepeatOffendersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// GetWithResponseHandler invoke function getAttackSimulationRepeatOffenders
func (m *GetAttackSimulationRepeatOffendersRequestBuilder) GetWithResponseHandler(requestConfiguration *GetAttackSimulationRepeatOffendersRequestBuilderGetRequestConfiguration)(GetAttackSimulationRepeatOffendersResponseable, error) {
    return m.GetWithResponseHandler(requestConfiguration, nil);
}
// GetWithResponseHandler invoke function getAttackSimulationRepeatOffenders
func (m *GetAttackSimulationRepeatOffendersRequestBuilder) GetWithResponseHandler(requestConfiguration *GetAttackSimulationRepeatOffendersRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(GetAttackSimulationRepeatOffendersResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetAttackSimulationRepeatOffendersResponseFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetAttackSimulationRepeatOffendersResponseable), nil
}
