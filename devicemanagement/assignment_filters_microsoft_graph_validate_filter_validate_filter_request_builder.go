package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AssignmentFiltersMicrosoftGraphValidateFilterValidateFilterRequestBuilder provides operations to call the validateFilter method.
type AssignmentFiltersMicrosoftGraphValidateFilterValidateFilterRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AssignmentFiltersMicrosoftGraphValidateFilterValidateFilterRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AssignmentFiltersMicrosoftGraphValidateFilterValidateFilterRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAssignmentFiltersMicrosoftGraphValidateFilterValidateFilterRequestBuilderInternal instantiates a new ValidateFilterRequestBuilder and sets the default values.
func NewAssignmentFiltersMicrosoftGraphValidateFilterValidateFilterRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AssignmentFiltersMicrosoftGraphValidateFilterValidateFilterRequestBuilder) {
    m := &AssignmentFiltersMicrosoftGraphValidateFilterValidateFilterRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/assignmentFilters/microsoft.graph.validateFilter";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewAssignmentFiltersMicrosoftGraphValidateFilterValidateFilterRequestBuilder instantiates a new ValidateFilterRequestBuilder and sets the default values.
func NewAssignmentFiltersMicrosoftGraphValidateFilterValidateFilterRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AssignmentFiltersMicrosoftGraphValidateFilterValidateFilterRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAssignmentFiltersMicrosoftGraphValidateFilterValidateFilterRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action validateFilter
func (m *AssignmentFiltersMicrosoftGraphValidateFilterValidateFilterRequestBuilder) Post(ctx context.Context, body AssignmentFiltersMicrosoftGraphValidateFilterValidateFilterPostRequestBodyable, requestConfiguration *AssignmentFiltersMicrosoftGraphValidateFilterValidateFilterRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AssignmentFilterValidationResultable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAssignmentFilterValidationResultFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AssignmentFilterValidationResultable), nil
}
// ToPostRequestInformation invoke action validateFilter
func (m *AssignmentFiltersMicrosoftGraphValidateFilterValidateFilterRequestBuilder) ToPostRequestInformation(ctx context.Context, body AssignmentFiltersMicrosoftGraphValidateFilterValidateFilterPostRequestBodyable, requestConfiguration *AssignmentFiltersMicrosoftGraphValidateFilterValidateFilterRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
