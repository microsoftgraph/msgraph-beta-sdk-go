package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ManagedAppRegistrationsItemAppliedPoliciesItemMicrosoftGraphTargetAppsRequestBuilder provides operations to call the targetApps method.
type ManagedAppRegistrationsItemAppliedPoliciesItemMicrosoftGraphTargetAppsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ManagedAppRegistrationsItemAppliedPoliciesItemMicrosoftGraphTargetAppsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedAppRegistrationsItemAppliedPoliciesItemMicrosoftGraphTargetAppsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewManagedAppRegistrationsItemAppliedPoliciesItemMicrosoftGraphTargetAppsRequestBuilderInternal instantiates a new MicrosoftGraphTargetAppsRequestBuilder and sets the default values.
func NewManagedAppRegistrationsItemAppliedPoliciesItemMicrosoftGraphTargetAppsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedAppRegistrationsItemAppliedPoliciesItemMicrosoftGraphTargetAppsRequestBuilder) {
    m := &ManagedAppRegistrationsItemAppliedPoliciesItemMicrosoftGraphTargetAppsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/managedAppRegistrations/{managedAppRegistration%2Did}/appliedPolicies/{managedAppPolicy%2Did}/microsoft.graph.targetApps";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewManagedAppRegistrationsItemAppliedPoliciesItemMicrosoftGraphTargetAppsRequestBuilder instantiates a new MicrosoftGraphTargetAppsRequestBuilder and sets the default values.
func NewManagedAppRegistrationsItemAppliedPoliciesItemMicrosoftGraphTargetAppsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedAppRegistrationsItemAppliedPoliciesItemMicrosoftGraphTargetAppsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedAppRegistrationsItemAppliedPoliciesItemMicrosoftGraphTargetAppsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action targetApps
func (m *ManagedAppRegistrationsItemAppliedPoliciesItemMicrosoftGraphTargetAppsRequestBuilder) Post(ctx context.Context, body ManagedAppRegistrationsItemAppliedPoliciesItemMicrosoftGraphTargetAppsTargetAppsPostRequestBodyable, requestConfiguration *ManagedAppRegistrationsItemAppliedPoliciesItemMicrosoftGraphTargetAppsRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation invoke action targetApps
func (m *ManagedAppRegistrationsItemAppliedPoliciesItemMicrosoftGraphTargetAppsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ManagedAppRegistrationsItemAppliedPoliciesItemMicrosoftGraphTargetAppsTargetAppsPostRequestBodyable, requestConfiguration *ManagedAppRegistrationsItemAppliedPoliciesItemMicrosoftGraphTargetAppsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
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
