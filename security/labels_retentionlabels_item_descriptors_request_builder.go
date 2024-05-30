package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// LabelsRetentionlabelsItemDescriptorsRequestBuilder provides operations to manage the descriptors property of the microsoft.graph.security.retentionLabel entity.
type LabelsRetentionlabelsItemDescriptorsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LabelsRetentionlabelsItemDescriptorsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LabelsRetentionlabelsItemDescriptorsRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// LabelsRetentionlabelsItemDescriptorsRequestBuilderGetQueryParameters represents out-of-the-box values that provide more options to improve the manageability and organization of the content you need to label.
type LabelsRetentionlabelsItemDescriptorsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// LabelsRetentionlabelsItemDescriptorsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LabelsRetentionlabelsItemDescriptorsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *LabelsRetentionlabelsItemDescriptorsRequestBuilderGetQueryParameters
}
// LabelsRetentionlabelsItemDescriptorsRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LabelsRetentionlabelsItemDescriptorsRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AuthorityTemplate provides operations to manage the authorityTemplate property of the microsoft.graph.security.filePlanDescriptor entity.
// returns a *LabelsRetentionlabelsItemDescriptorsAuthoritytemplateAuthorityTemplateRequestBuilder when successful
func (m *LabelsRetentionlabelsItemDescriptorsRequestBuilder) AuthorityTemplate()(*LabelsRetentionlabelsItemDescriptorsAuthoritytemplateAuthorityTemplateRequestBuilder) {
    return NewLabelsRetentionlabelsItemDescriptorsAuthoritytemplateAuthorityTemplateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CategoryTemplate provides operations to manage the categoryTemplate property of the microsoft.graph.security.filePlanDescriptor entity.
// returns a *LabelsRetentionlabelsItemDescriptorsCategorytemplateCategoryTemplateRequestBuilder when successful
func (m *LabelsRetentionlabelsItemDescriptorsRequestBuilder) CategoryTemplate()(*LabelsRetentionlabelsItemDescriptorsCategorytemplateCategoryTemplateRequestBuilder) {
    return NewLabelsRetentionlabelsItemDescriptorsCategorytemplateCategoryTemplateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CitationTemplate provides operations to manage the citationTemplate property of the microsoft.graph.security.filePlanDescriptor entity.
// returns a *LabelsRetentionlabelsItemDescriptorsCitationtemplateCitationTemplateRequestBuilder when successful
func (m *LabelsRetentionlabelsItemDescriptorsRequestBuilder) CitationTemplate()(*LabelsRetentionlabelsItemDescriptorsCitationtemplateCitationTemplateRequestBuilder) {
    return NewLabelsRetentionlabelsItemDescriptorsCitationtemplateCitationTemplateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewLabelsRetentionlabelsItemDescriptorsRequestBuilderInternal instantiates a new LabelsRetentionlabelsItemDescriptorsRequestBuilder and sets the default values.
func NewLabelsRetentionlabelsItemDescriptorsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LabelsRetentionlabelsItemDescriptorsRequestBuilder) {
    m := &LabelsRetentionlabelsItemDescriptorsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/labels/retentionLabels/{retentionLabel%2Did}/descriptors{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewLabelsRetentionlabelsItemDescriptorsRequestBuilder instantiates a new LabelsRetentionlabelsItemDescriptorsRequestBuilder and sets the default values.
func NewLabelsRetentionlabelsItemDescriptorsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LabelsRetentionlabelsItemDescriptorsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLabelsRetentionlabelsItemDescriptorsRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property descriptors for security
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *LabelsRetentionlabelsItemDescriptorsRequestBuilder) Delete(ctx context.Context, requestConfiguration *LabelsRetentionlabelsItemDescriptorsRequestBuilderDeleteRequestConfiguration)(error) {
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
// DepartmentTemplate provides operations to manage the departmentTemplate property of the microsoft.graph.security.filePlanDescriptor entity.
// returns a *LabelsRetentionlabelsItemDescriptorsDepartmenttemplateDepartmentTemplateRequestBuilder when successful
func (m *LabelsRetentionlabelsItemDescriptorsRequestBuilder) DepartmentTemplate()(*LabelsRetentionlabelsItemDescriptorsDepartmenttemplateDepartmentTemplateRequestBuilder) {
    return NewLabelsRetentionlabelsItemDescriptorsDepartmenttemplateDepartmentTemplateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// FilePlanReferenceTemplate provides operations to manage the filePlanReferenceTemplate property of the microsoft.graph.security.filePlanDescriptor entity.
// returns a *LabelsRetentionlabelsItemDescriptorsFileplanreferencetemplateFilePlanReferenceTemplateRequestBuilder when successful
func (m *LabelsRetentionlabelsItemDescriptorsRequestBuilder) FilePlanReferenceTemplate()(*LabelsRetentionlabelsItemDescriptorsFileplanreferencetemplateFilePlanReferenceTemplateRequestBuilder) {
    return NewLabelsRetentionlabelsItemDescriptorsFileplanreferencetemplateFilePlanReferenceTemplateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get represents out-of-the-box values that provide more options to improve the manageability and organization of the content you need to label.
// returns a FilePlanDescriptorable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *LabelsRetentionlabelsItemDescriptorsRequestBuilder) Get(ctx context.Context, requestConfiguration *LabelsRetentionlabelsItemDescriptorsRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.FilePlanDescriptorable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateFilePlanDescriptorFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.FilePlanDescriptorable), nil
}
// Patch update the navigation property descriptors in security
// returns a FilePlanDescriptorable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *LabelsRetentionlabelsItemDescriptorsRequestBuilder) Patch(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.FilePlanDescriptorable, requestConfiguration *LabelsRetentionlabelsItemDescriptorsRequestBuilderPatchRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.FilePlanDescriptorable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateFilePlanDescriptorFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.FilePlanDescriptorable), nil
}
// ToDeleteRequestInformation delete navigation property descriptors for security
// returns a *RequestInformation when successful
func (m *LabelsRetentionlabelsItemDescriptorsRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *LabelsRetentionlabelsItemDescriptorsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation represents out-of-the-box values that provide more options to improve the manageability and organization of the content you need to label.
// returns a *RequestInformation when successful
func (m *LabelsRetentionlabelsItemDescriptorsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LabelsRetentionlabelsItemDescriptorsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property descriptors in security
// returns a *RequestInformation when successful
func (m *LabelsRetentionlabelsItemDescriptorsRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.FilePlanDescriptorable, requestConfiguration *LabelsRetentionlabelsItemDescriptorsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *LabelsRetentionlabelsItemDescriptorsRequestBuilder when successful
func (m *LabelsRetentionlabelsItemDescriptorsRequestBuilder) WithUrl(rawUrl string)(*LabelsRetentionlabelsItemDescriptorsRequestBuilder) {
    return NewLabelsRetentionlabelsItemDescriptorsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
