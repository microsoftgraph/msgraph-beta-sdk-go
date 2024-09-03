package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSecurityInformationProtectionLabelPolicySettingsRequestBuilder provides operations to manage the labelPolicySettings property of the microsoft.graph.security.informationProtection entity.
type ItemSecurityInformationProtectionLabelPolicySettingsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSecurityInformationProtectionLabelPolicySettingsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSecurityInformationProtectionLabelPolicySettingsRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemSecurityInformationProtectionLabelPolicySettingsRequestBuilderGetQueryParameters read the properties and relationships of an informationProtectionPolicySetting object. The settings exposed by this API should be used in applications to populate the moreInfoUrl property for Microsoft Purview Information Protection help, and indicate whether labeling is mandatory for the user and whether justification must be provided on downgrade.
type ItemSecurityInformationProtectionLabelPolicySettingsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemSecurityInformationProtectionLabelPolicySettingsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSecurityInformationProtectionLabelPolicySettingsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSecurityInformationProtectionLabelPolicySettingsRequestBuilderGetQueryParameters
}
// ItemSecurityInformationProtectionLabelPolicySettingsRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSecurityInformationProtectionLabelPolicySettingsRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSecurityInformationProtectionLabelPolicySettingsRequestBuilderInternal instantiates a new ItemSecurityInformationProtectionLabelPolicySettingsRequestBuilder and sets the default values.
func NewItemSecurityInformationProtectionLabelPolicySettingsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSecurityInformationProtectionLabelPolicySettingsRequestBuilder) {
    m := &ItemSecurityInformationProtectionLabelPolicySettingsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/security/informationProtection/labelPolicySettings{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemSecurityInformationProtectionLabelPolicySettingsRequestBuilder instantiates a new ItemSecurityInformationProtectionLabelPolicySettingsRequestBuilder and sets the default values.
func NewItemSecurityInformationProtectionLabelPolicySettingsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSecurityInformationProtectionLabelPolicySettingsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSecurityInformationProtectionLabelPolicySettingsRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property labelPolicySettings for users
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSecurityInformationProtectionLabelPolicySettingsRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemSecurityInformationProtectionLabelPolicySettingsRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get read the properties and relationships of an informationProtectionPolicySetting object. The settings exposed by this API should be used in applications to populate the moreInfoUrl property for Microsoft Purview Information Protection help, and indicate whether labeling is mandatory for the user and whether justification must be provided on downgrade.
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a InformationProtectionPolicySettingable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-informationprotectionpolicysetting-get?view=graph-rest-beta
func (m *ItemSecurityInformationProtectionLabelPolicySettingsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSecurityInformationProtectionLabelPolicySettingsRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.InformationProtectionPolicySettingable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateInformationProtectionPolicySettingFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.InformationProtectionPolicySettingable), nil
}
// Patch update the navigation property labelPolicySettings in users
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a InformationProtectionPolicySettingable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSecurityInformationProtectionLabelPolicySettingsRequestBuilder) Patch(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.InformationProtectionPolicySettingable, requestConfiguration *ItemSecurityInformationProtectionLabelPolicySettingsRequestBuilderPatchRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.InformationProtectionPolicySettingable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateInformationProtectionPolicySettingFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.InformationProtectionPolicySettingable), nil
}
// ToDeleteRequestInformation delete navigation property labelPolicySettings for users
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a *RequestInformation when successful
func (m *ItemSecurityInformationProtectionLabelPolicySettingsRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemSecurityInformationProtectionLabelPolicySettingsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of an informationProtectionPolicySetting object. The settings exposed by this API should be used in applications to populate the moreInfoUrl property for Microsoft Purview Information Protection help, and indicate whether labeling is mandatory for the user and whether justification must be provided on downgrade.
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a *RequestInformation when successful
func (m *ItemSecurityInformationProtectionLabelPolicySettingsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSecurityInformationProtectionLabelPolicySettingsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property labelPolicySettings in users
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a *RequestInformation when successful
func (m *ItemSecurityInformationProtectionLabelPolicySettingsRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.InformationProtectionPolicySettingable, requestConfiguration *ItemSecurityInformationProtectionLabelPolicySettingsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a *ItemSecurityInformationProtectionLabelPolicySettingsRequestBuilder when successful
func (m *ItemSecurityInformationProtectionLabelPolicySettingsRequestBuilder) WithUrl(rawUrl string)(*ItemSecurityInformationProtectionLabelPolicySettingsRequestBuilder) {
    return NewItemSecurityInformationProtectionLabelPolicySettingsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
