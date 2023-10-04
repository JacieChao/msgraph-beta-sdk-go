package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CasesEdiscoveryCasesItemReviewSetsItemMicrosoftGraphSecurityExportRequestBuilder provides operations to call the export method.
type CasesEdiscoveryCasesItemReviewSetsItemMicrosoftGraphSecurityExportRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CasesEdiscoveryCasesItemReviewSetsItemMicrosoftGraphSecurityExportRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoveryCasesItemReviewSetsItemMicrosoftGraphSecurityExportRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCasesEdiscoveryCasesItemReviewSetsItemMicrosoftGraphSecurityExportRequestBuilderInternal instantiates a new MicrosoftGraphSecurityExportRequestBuilder and sets the default values.
func NewCasesEdiscoveryCasesItemReviewSetsItemMicrosoftGraphSecurityExportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoveryCasesItemReviewSetsItemMicrosoftGraphSecurityExportRequestBuilder) {
    m := &CasesEdiscoveryCasesItemReviewSetsItemMicrosoftGraphSecurityExportRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/reviewSets/{ediscoveryReviewSet%2Did}/microsoft.graph.security.export", pathParameters),
    }
    return m
}
// NewCasesEdiscoveryCasesItemReviewSetsItemMicrosoftGraphSecurityExportRequestBuilder instantiates a new MicrosoftGraphSecurityExportRequestBuilder and sets the default values.
func NewCasesEdiscoveryCasesItemReviewSetsItemMicrosoftGraphSecurityExportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoveryCasesItemReviewSetsItemMicrosoftGraphSecurityExportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoveryCasesItemReviewSetsItemMicrosoftGraphSecurityExportRequestBuilderInternal(urlParams, requestAdapter)
}
// Post initiate an export from a reviewSet.  For details, see Export documents from a review set in eDiscovery (Premium). This API is supported in the following national cloud deployments.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-ediscoveryreviewset-export?view=graph-rest-1.0
func (m *CasesEdiscoveryCasesItemReviewSetsItemMicrosoftGraphSecurityExportRequestBuilder) Post(ctx context.Context, body CasesEdiscoveryCasesItemReviewSetsItemMicrosoftGraphSecurityExportExportPostRequestBodyable, requestConfiguration *CasesEdiscoveryCasesItemReviewSetsItemMicrosoftGraphSecurityExportRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation initiate an export from a reviewSet.  For details, see Export documents from a review set in eDiscovery (Premium). This API is supported in the following national cloud deployments.
func (m *CasesEdiscoveryCasesItemReviewSetsItemMicrosoftGraphSecurityExportRequestBuilder) ToPostRequestInformation(ctx context.Context, body CasesEdiscoveryCasesItemReviewSetsItemMicrosoftGraphSecurityExportExportPostRequestBodyable, requestConfiguration *CasesEdiscoveryCasesItemReviewSetsItemMicrosoftGraphSecurityExportRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
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
func (m *CasesEdiscoveryCasesItemReviewSetsItemMicrosoftGraphSecurityExportRequestBuilder) WithUrl(rawUrl string)(*CasesEdiscoveryCasesItemReviewSetsItemMicrosoftGraphSecurityExportRequestBuilder) {
    return NewCasesEdiscoveryCasesItemReviewSetsItemMicrosoftGraphSecurityExportRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
