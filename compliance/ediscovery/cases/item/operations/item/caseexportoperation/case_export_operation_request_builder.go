package caseexportoperation

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i853045a1c17d869baaeeca6d724b0704bd59922311548198cc7c32c73ab5d870 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/operations/item/caseexportoperation/getdownloadurl"
)

// CaseExportOperationRequestBuilder builds and executes requests for operations under \compliance\ediscovery\cases\{case-id}\operations\{caseOperation-id}\microsoft.graph.ediscovery.caseExportOperation
type CaseExportOperationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// NewCaseExportOperationRequestBuilderInternal instantiates a new CaseExportOperationRequestBuilder and sets the default values.
func NewCaseExportOperationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CaseExportOperationRequestBuilder) {
    m := &CaseExportOperationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/compliance/ediscovery/cases/{case_id}/operations/{caseOperation_id}/microsoft.graph.ediscovery.caseExportOperation";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCaseExportOperationRequestBuilder instantiates a new CaseExportOperationRequestBuilder and sets the default values.
func NewCaseExportOperationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CaseExportOperationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCaseExportOperationRequestBuilderInternal(urlParams, requestAdapter)
}
// GetDownloadUrl provides operations to call the getDownloadUrl method.
func (m *CaseExportOperationRequestBuilder) GetDownloadUrl()(*i853045a1c17d869baaeeca6d724b0704bd59922311548198cc7c32c73ab5d870.GetDownloadUrlRequestBuilder) {
    return i853045a1c17d869baaeeca6d724b0704bd59922311548198cc7c32c73ab5d870.NewGetDownloadUrlRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
