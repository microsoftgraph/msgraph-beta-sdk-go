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
// GetPolicySummaryWithPolicyIdRequestBuilderGetOptions options for Get
type GetPolicySummaryWithPolicyIdRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
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
// CreateGetRequestInformation invoke function getPolicySummary
func (m *GetPolicySummaryWithPolicyIdRequestBuilder) CreateGetRequestInformation(options *GetPolicySummaryWithPolicyIdRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get invoke function getPolicySummary
func (m *GetPolicySummaryWithPolicyIdRequestBuilder) Get(options *GetPolicySummaryWithPolicyIdRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConfigManagerPolicySummaryable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateConfigManagerPolicySummaryFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConfigManagerPolicySummaryable), nil
}
