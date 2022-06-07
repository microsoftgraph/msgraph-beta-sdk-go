package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i09cf923f0e84a9700a33e7913ad84e00557e1206775bfb490bfdee5e1eed8ea8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onenote/pages/item/preview"
    i1caba1f27f11980269f83254b7206a72eb91093aeecd2c44476b6ed0621ce573 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onenote/pages/item/parentnotebook"
    i5c1cc1fc7c4561b84aa05066752d58b6db53857db9065a8da941728505e4c732 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onenote/pages/item/onenotepatchcontent"
    i6f0db3c9a10b7adff30a3ede5912a026ed0ef56988193ac07c98a1789af6b242 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onenote/pages/item/copytosection"
    ic1dde0489220b401ddce60f4d7712c2c704faa712e4055c5f74eb9fc5196f035 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onenote/pages/item/content"
    id128a1fdf762fd57ad499e421990f4d8a1cdd77672e8977ee6bbc8ab31c0f0e8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/onenote/pages/item/parentsection"
)

// OnenotePageItemRequestBuilder provides operations to manage the pages property of the microsoft.graph.onenote entity.
type OnenotePageItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// OnenotePageItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnenotePageItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// OnenotePageItemRequestBuilderGetQueryParameters the pages in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
type OnenotePageItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// OnenotePageItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnenotePageItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *OnenotePageItemRequestBuilderGetQueryParameters
}
// OnenotePageItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnenotePageItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewOnenotePageItemRequestBuilderInternal instantiates a new OnenotePageItemRequestBuilder and sets the default values.
func NewOnenotePageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnenotePageItemRequestBuilder) {
    m := &OnenotePageItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/onenote/pages/{onenotePage%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOnenotePageItemRequestBuilder instantiates a new OnenotePageItemRequestBuilder and sets the default values.
func NewOnenotePageItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnenotePageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnenotePageItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Content the content property
func (m *OnenotePageItemRequestBuilder) Content()(*ic1dde0489220b401ddce60f4d7712c2c704faa712e4055c5f74eb9fc5196f035.ContentRequestBuilder) {
    return ic1dde0489220b401ddce60f4d7712c2c704faa712e4055c5f74eb9fc5196f035.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CopyToSection the copyToSection property
func (m *OnenotePageItemRequestBuilder) CopyToSection()(*i6f0db3c9a10b7adff30a3ede5912a026ed0ef56988193ac07c98a1789af6b242.CopyToSectionRequestBuilder) {
    return i6f0db3c9a10b7adff30a3ede5912a026ed0ef56988193ac07c98a1789af6b242.NewCopyToSectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property pages for users
func (m *OnenotePageItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property pages for users
func (m *OnenotePageItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *OnenotePageItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the pages in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
func (m *OnenotePageItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the pages in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
func (m *OnenotePageItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *OnenotePageItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property pages in users
func (m *OnenotePageItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnenotePageable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property pages in users
func (m *OnenotePageItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnenotePageable, requestConfiguration *OnenotePageItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property pages for users
func (m *OnenotePageItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property pages for users
func (m *OnenotePageItemRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *OnenotePageItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the pages in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
func (m *OnenotePageItemRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnenotePageable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler the pages in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
func (m *OnenotePageItemRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *OnenotePageItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnenotePageable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOnenotePageFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnenotePageable), nil
}
// OnenotePatchContent the onenotePatchContent property
func (m *OnenotePageItemRequestBuilder) OnenotePatchContent()(*i5c1cc1fc7c4561b84aa05066752d58b6db53857db9065a8da941728505e4c732.OnenotePatchContentRequestBuilder) {
    return i5c1cc1fc7c4561b84aa05066752d58b6db53857db9065a8da941728505e4c732.NewOnenotePatchContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ParentNotebook the parentNotebook property
func (m *OnenotePageItemRequestBuilder) ParentNotebook()(*i1caba1f27f11980269f83254b7206a72eb91093aeecd2c44476b6ed0621ce573.ParentNotebookRequestBuilder) {
    return i1caba1f27f11980269f83254b7206a72eb91093aeecd2c44476b6ed0621ce573.NewParentNotebookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ParentSection the parentSection property
func (m *OnenotePageItemRequestBuilder) ParentSection()(*id128a1fdf762fd57ad499e421990f4d8a1cdd77672e8977ee6bbc8ab31c0f0e8.ParentSectionRequestBuilder) {
    return id128a1fdf762fd57ad499e421990f4d8a1cdd77672e8977ee6bbc8ab31c0f0e8.NewParentSectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property pages in users
func (m *OnenotePageItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnenotePageable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property pages in users
func (m *OnenotePageItemRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnenotePageable, requestConfiguration *OnenotePageItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Preview provides operations to call the preview method.
func (m *OnenotePageItemRequestBuilder) Preview()(*i09cf923f0e84a9700a33e7913ad84e00557e1206775bfb490bfdee5e1eed8ea8.PreviewRequestBuilder) {
    return i09cf923f0e84a9700a33e7913ad84e00557e1206775bfb490bfdee5e1eed8ea8.NewPreviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
