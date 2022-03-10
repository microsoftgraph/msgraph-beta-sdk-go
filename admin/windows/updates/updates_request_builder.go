package updates

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i0475d54fef399e8dd6955c2d001134c125fb7788ce83c57eedfc935bfa992080 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/catalog"
    i387197d5cc2dce8dbb3b5340ab9e1213cb3d3d425caddd84e7d1c7ddb338d669 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/deployments"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    id8360e8943a0ee251dee4de20317b335abdd17663b8ed227f06c1426f54dc863 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/updatableassets"
    ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/windowsupdates"
    i6ba3b57fd9faa032676b28e7c1ba8fbcadca736ef31ae441fd9cb31c54634f6b "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/updatableassets/item"
    ibf1ad675b59509c5f3216e9d430e98ab3b887534afb4bca984ab70b32d39075b "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/deployments/item"
)

// UpdatesRequestBuilder provides operations to manage the updates property of the microsoft.graph.windowsUpdates.windows entity.
type UpdatesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// UpdatesRequestBuilderDeleteOptions options for Delete
type UpdatesRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UpdatesRequestBuilderGetOptions options for Get
type UpdatesRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *UpdatesRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UpdatesRequestBuilderGetQueryParameters entity that acts as a container for the functionality of the Windows Update for Business deployment service. Read-only.
type UpdatesRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// UpdatesRequestBuilderPatchOptions options for Patch
type UpdatesRequestBuilderPatchOptions struct {
    // 
    Body ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.Updatesable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *UpdatesRequestBuilder) Catalog()(*i0475d54fef399e8dd6955c2d001134c125fb7788ce83c57eedfc935bfa992080.CatalogRequestBuilder) {
    return i0475d54fef399e8dd6955c2d001134c125fb7788ce83c57eedfc935bfa992080.NewCatalogRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUpdatesRequestBuilderInternal instantiates a new UpdatesRequestBuilder and sets the default values.
func NewUpdatesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UpdatesRequestBuilder) {
    m := &UpdatesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/admin/windows/updates{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUpdatesRequestBuilder instantiates a new UpdatesRequestBuilder and sets the default values.
func NewUpdatesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UpdatesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUpdatesRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property updates for admin
func (m *UpdatesRequestBuilder) CreateDeleteRequestInformation(options *UpdatesRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation entity that acts as a container for the functionality of the Windows Update for Business deployment service. Read-only.
func (m *UpdatesRequestBuilder) CreateGetRequestInformation(options *UpdatesRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property updates in admin
func (m *UpdatesRequestBuilder) CreatePatchRequestInformation(options *UpdatesRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property updates for admin
func (m *UpdatesRequestBuilder) Delete(options *UpdatesRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *UpdatesRequestBuilder) Deployments()(*i387197d5cc2dce8dbb3b5340ab9e1213cb3d3d425caddd84e7d1c7ddb338d669.DeploymentsRequestBuilder) {
    return i387197d5cc2dce8dbb3b5340ab9e1213cb3d3d425caddd84e7d1c7ddb338d669.NewDeploymentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeploymentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.admin.windows.updates.deployments.item collection
func (m *UpdatesRequestBuilder) DeploymentsById(id string)(*ibf1ad675b59509c5f3216e9d430e98ab3b887534afb4bca984ab70b32d39075b.DeploymentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deployment_id"] = id
    }
    return ibf1ad675b59509c5f3216e9d430e98ab3b887534afb4bca984ab70b32d39075b.NewDeploymentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get entity that acts as a container for the functionality of the Windows Update for Business deployment service. Read-only.
func (m *UpdatesRequestBuilder) Get(options *UpdatesRequestBuilderGetOptions)(ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.Updatesable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.CreateUpdatesFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ifded49a845bbaa9057da6e2cf565863ac34eb797e99b129c3e0659166af6b7e2.Updatesable), nil
}
// Patch update the navigation property updates in admin
func (m *UpdatesRequestBuilder) Patch(options *UpdatesRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *UpdatesRequestBuilder) UpdatableAssets()(*id8360e8943a0ee251dee4de20317b335abdd17663b8ed227f06c1426f54dc863.UpdatableAssetsRequestBuilder) {
    return id8360e8943a0ee251dee4de20317b335abdd17663b8ed227f06c1426f54dc863.NewUpdatableAssetsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UpdatableAssetsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.admin.windows.updates.updatableAssets.item collection
func (m *UpdatesRequestBuilder) UpdatableAssetsById(id string)(*i6ba3b57fd9faa032676b28e7c1ba8fbcadca736ef31ae441fd9cb31c54634f6b.UpdatableAssetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["updatableAsset_id"] = id
    }
    return i6ba3b57fd9faa032676b28e7c1ba8fbcadca736ef31ae441fd9cb31c54634f6b.NewUpdatableAssetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
