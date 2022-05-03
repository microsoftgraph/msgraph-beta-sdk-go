package getpolicysummarywithpolicyid

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// GetPolicySummaryWithPolicyIdRequestBuilder provides operations to call the getPolicySummary method.
type GetPolicySummaryWithPolicyIdRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetPolicySummaryWithPolicyIdRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetPolicySummaryWithPolicyIdRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetPolicySummaryWithPolicyIdRequestBuilderInternal instantiates a new GetPolicySummaryWithPolicyIdRequestBuilder and sets the default values.
func NewGetPolicySummaryWithPolicyIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, policyId *string)(*GetPolicySummaryWithPolicyIdRequestBuilder) {
    m := &GetPolicySummaryWithPolicyIdRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/configManagerCollections/microsoft.graph.getPolicySummary(policyId='{policyId}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if policyId != nil {
        urlTplParams[""] = *policyId
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetPolicySummaryWithPolicyIdRequestBuilder instantiates a new GetPolicySummaryWithPolicyIdRequestBuilder and sets the default values.
func NewGetPolicySummaryWithPolicyIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetPolicySummaryWithPolicyIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetPolicySummaryWithPolicyIdRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CreateGetRequestInformationWithRequestConfiguration invoke function getPolicySummary
func (m *GetPolicySummaryWithPolicyIdRequestBuilder) CreateGetRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration invoke function getPolicySummary
func (m *GetPolicySummaryWithPolicyIdRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *GetPolicySummaryWithPolicyIdRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// GetWithResponseHandler invoke function getPolicySummary
func (m *GetPolicySummaryWithPolicyIdRequestBuilder) GetWithResponseHandler(requestConfiguration *GetPolicySummaryWithPolicyIdRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConfigManagerPolicySummaryable, error) {
    return m.GetWithResponseHandler(requestConfiguration, nil);
}
// GetWithResponseHandler invoke function getPolicySummary
func (m *GetPolicySummaryWithPolicyIdRequestBuilder) GetWithResponseHandler(requestConfiguration *GetPolicySummaryWithPolicyIdRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConfigManagerPolicySummaryable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateConfigManagerPolicySummaryFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConfigManagerPolicySummaryable), nil
}
