package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CompliancePoliciesItemSetScheduledActionsRequestBuilder provides operations to call the setScheduledActions method.
type CompliancePoliciesItemSetScheduledActionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CompliancePoliciesItemSetScheduledActionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompliancePoliciesItemSetScheduledActionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCompliancePoliciesItemSetScheduledActionsRequestBuilderInternal instantiates a new SetScheduledActionsRequestBuilder and sets the default values.
func NewCompliancePoliciesItemSetScheduledActionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompliancePoliciesItemSetScheduledActionsRequestBuilder) {
    m := &CompliancePoliciesItemSetScheduledActionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/compliancePolicies/{deviceManagementCompliancePolicy%2Did}/setScheduledActions", pathParameters),
    }
    return m
}
// NewCompliancePoliciesItemSetScheduledActionsRequestBuilder instantiates a new SetScheduledActionsRequestBuilder and sets the default values.
func NewCompliancePoliciesItemSetScheduledActionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompliancePoliciesItemSetScheduledActionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCompliancePoliciesItemSetScheduledActionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action setScheduledActions
func (m *CompliancePoliciesItemSetScheduledActionsRequestBuilder) Post(ctx context.Context, body CompliancePoliciesItemSetScheduledActionsPostRequestBodyable, requestConfiguration *CompliancePoliciesItemSetScheduledActionsRequestBuilderPostRequestConfiguration)(CompliancePoliciesItemSetScheduledActionsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCompliancePoliciesItemSetScheduledActionsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(CompliancePoliciesItemSetScheduledActionsResponseable), nil
}
// ToPostRequestInformation invoke action setScheduledActions
func (m *CompliancePoliciesItemSetScheduledActionsRequestBuilder) ToPostRequestInformation(ctx context.Context, body CompliancePoliciesItemSetScheduledActionsPostRequestBodyable, requestConfiguration *CompliancePoliciesItemSetScheduledActionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *CompliancePoliciesItemSetScheduledActionsRequestBuilder) WithUrl(rawUrl string)(*CompliancePoliciesItemSetScheduledActionsRequestBuilder) {
    return NewCompliancePoliciesItemSetScheduledActionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
