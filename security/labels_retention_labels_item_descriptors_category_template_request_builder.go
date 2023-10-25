package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// LabelsRetentionLabelsItemDescriptorsCategoryTemplateRequestBuilder provides operations to manage the categoryTemplate property of the microsoft.graph.security.filePlanDescriptor entity.
type LabelsRetentionLabelsItemDescriptorsCategoryTemplateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LabelsRetentionLabelsItemDescriptorsCategoryTemplateRequestBuilderGetQueryParameters specifies a group of similar types of content in a particular department.
type LabelsRetentionLabelsItemDescriptorsCategoryTemplateRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// LabelsRetentionLabelsItemDescriptorsCategoryTemplateRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LabelsRetentionLabelsItemDescriptorsCategoryTemplateRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *LabelsRetentionLabelsItemDescriptorsCategoryTemplateRequestBuilderGetQueryParameters
}
// NewLabelsRetentionLabelsItemDescriptorsCategoryTemplateRequestBuilderInternal instantiates a new CategoryTemplateRequestBuilder and sets the default values.
func NewLabelsRetentionLabelsItemDescriptorsCategoryTemplateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LabelsRetentionLabelsItemDescriptorsCategoryTemplateRequestBuilder) {
    m := &LabelsRetentionLabelsItemDescriptorsCategoryTemplateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/labels/retentionLabels/{retentionLabel%2Did}/descriptors/categoryTemplate{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewLabelsRetentionLabelsItemDescriptorsCategoryTemplateRequestBuilder instantiates a new CategoryTemplateRequestBuilder and sets the default values.
func NewLabelsRetentionLabelsItemDescriptorsCategoryTemplateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LabelsRetentionLabelsItemDescriptorsCategoryTemplateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLabelsRetentionLabelsItemDescriptorsCategoryTemplateRequestBuilderInternal(urlParams, requestAdapter)
}
// Get specifies a group of similar types of content in a particular department.
func (m *LabelsRetentionLabelsItemDescriptorsCategoryTemplateRequestBuilder) Get(ctx context.Context, requestConfiguration *LabelsRetentionLabelsItemDescriptorsCategoryTemplateRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CategoryTemplateable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateCategoryTemplateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CategoryTemplateable), nil
}
// ToGetRequestInformation specifies a group of similar types of content in a particular department.
func (m *LabelsRetentionLabelsItemDescriptorsCategoryTemplateRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LabelsRetentionLabelsItemDescriptorsCategoryTemplateRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *LabelsRetentionLabelsItemDescriptorsCategoryTemplateRequestBuilder) WithUrl(rawUrl string)(*LabelsRetentionLabelsItemDescriptorsCategoryTemplateRequestBuilder) {
    return NewLabelsRetentionLabelsItemDescriptorsCategoryTemplateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
