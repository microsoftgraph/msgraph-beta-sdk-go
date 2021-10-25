package caseexportoperation

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i853045a1c17d869baaeeca6d724b0704bd59922311548198cc7c32c73ab5d870 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/operations/item/caseexportoperation/getdownloadurl"
)

type CaseExportOperationRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
func NewCaseExportOperationRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CaseExportOperationRequestBuilder) {
    m := &CaseExportOperationRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/compliance/ediscovery/cases/{case_id}/operations/{caseOperation_id}/microsoft.graph.ediscovery.caseExportOperation";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewCaseExportOperationRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CaseExportOperationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCaseExportOperationRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *CaseExportOperationRequestBuilder) GetDownloadUrl()(*i853045a1c17d869baaeeca6d724b0704bd59922311548198cc7c32c73ab5d870.GetDownloadUrlRequestBuilder) {
    return i853045a1c17d869baaeeca6d724b0704bd59922311548198cc7c32c73ab5d870.NewGetDownloadUrlRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
