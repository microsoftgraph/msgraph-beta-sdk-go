package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MicrosoftGraphGetAttackSimulationSimulationUserCoverageRequestBuilder provides operations to call the getAttackSimulationSimulationUserCoverage method.
type MicrosoftGraphGetAttackSimulationSimulationUserCoverageRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MicrosoftGraphGetAttackSimulationSimulationUserCoverageRequestBuilderGetQueryParameters invoke function getAttackSimulationSimulationUserCoverage
type MicrosoftGraphGetAttackSimulationSimulationUserCoverageRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// MicrosoftGraphGetAttackSimulationSimulationUserCoverageRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MicrosoftGraphGetAttackSimulationSimulationUserCoverageRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MicrosoftGraphGetAttackSimulationSimulationUserCoverageRequestBuilderGetQueryParameters
}
// NewMicrosoftGraphGetAttackSimulationSimulationUserCoverageRequestBuilderInternal instantiates a new MicrosoftGraphGetAttackSimulationSimulationUserCoverageRequestBuilder and sets the default values.
func NewMicrosoftGraphGetAttackSimulationSimulationUserCoverageRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosoftGraphGetAttackSimulationSimulationUserCoverageRequestBuilder) {
    m := &MicrosoftGraphGetAttackSimulationSimulationUserCoverageRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/reports/microsoft.graph.getAttackSimulationSimulationUserCoverage(){?%24top,%24skip,%24search,%24filter,%24count}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewMicrosoftGraphGetAttackSimulationSimulationUserCoverageRequestBuilder instantiates a new MicrosoftGraphGetAttackSimulationSimulationUserCoverageRequestBuilder and sets the default values.
func NewMicrosoftGraphGetAttackSimulationSimulationUserCoverageRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosoftGraphGetAttackSimulationSimulationUserCoverageRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMicrosoftGraphGetAttackSimulationSimulationUserCoverageRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function getAttackSimulationSimulationUserCoverage
func (m *MicrosoftGraphGetAttackSimulationSimulationUserCoverageRequestBuilder) Get(ctx context.Context, requestConfiguration *MicrosoftGraphGetAttackSimulationSimulationUserCoverageRequestBuilderGetRequestConfiguration)(MicrosoftGraphGetAttackSimulationSimulationUserCoverageGetAttackSimulationSimulationUserCoverageResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateMicrosoftGraphGetAttackSimulationSimulationUserCoverageGetAttackSimulationSimulationUserCoverageResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(MicrosoftGraphGetAttackSimulationSimulationUserCoverageGetAttackSimulationSimulationUserCoverageResponseable), nil
}
// ToGetRequestInformation invoke function getAttackSimulationSimulationUserCoverage
func (m *MicrosoftGraphGetAttackSimulationSimulationUserCoverageRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MicrosoftGraphGetAttackSimulationSimulationUserCoverageRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
