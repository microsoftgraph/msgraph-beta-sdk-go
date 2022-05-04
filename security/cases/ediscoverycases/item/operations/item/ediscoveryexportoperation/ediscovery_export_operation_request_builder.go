package ediscoveryexportoperation

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i90c4f0b54027c27ebb3d1a8544bedf0f49c1064e12092f75fd81fc8ed649bb14 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/cases/ediscoverycases/item/operations/item/ediscoveryexportoperation/getdownloadurl"
)

// EdiscoveryExportOperationRequestBuilder builds and executes requests for operations under \security\cases\ediscoveryCases\{ediscoveryCase-id}\operations\{caseOperation-id}\microsoft.graph.security.ediscoveryExportOperation
type EdiscoveryExportOperationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewEdiscoveryExportOperationRequestBuilderInternal instantiates a new EdiscoveryExportOperationRequestBuilder and sets the default values.
func NewEdiscoveryExportOperationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryExportOperationRequestBuilder) {
    m := &EdiscoveryExportOperationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/operations/{caseOperation%2Did}/microsoft.graph.security.ediscoveryExportOperation";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEdiscoveryExportOperationRequestBuilder instantiates a new EdiscoveryExportOperationRequestBuilder and sets the default values.
func NewEdiscoveryExportOperationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryExportOperationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEdiscoveryExportOperationRequestBuilderInternal(urlParams, requestAdapter)
}
// GetDownloadUrl provides operations to call the getDownloadUrl method.
func (m *EdiscoveryExportOperationRequestBuilder) GetDownloadUrl()(*i90c4f0b54027c27ebb3d1a8544bedf0f49c1064e12092f75fd81fc8ed649bb14.GetDownloadUrlRequestBuilder) {
    return i90c4f0b54027c27ebb3d1a8544bedf0f49c1064e12092f75fd81fc8ed649bb14.NewGetDownloadUrlRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
