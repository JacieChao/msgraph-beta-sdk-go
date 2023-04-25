package compliance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EdiscoveryCasesItemReviewSetsItemQueriesItemEdiscoveryApplyTagsRequestBuilder provides operations to call the applyTags method.
type EdiscoveryCasesItemReviewSetsItemQueriesItemEdiscoveryApplyTagsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EdiscoveryCasesItemReviewSetsItemQueriesItemEdiscoveryApplyTagsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryCasesItemReviewSetsItemQueriesItemEdiscoveryApplyTagsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEdiscoveryCasesItemReviewSetsItemQueriesItemEdiscoveryApplyTagsRequestBuilderInternal instantiates a new EdiscoveryApplyTagsRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemReviewSetsItemQueriesItemEdiscoveryApplyTagsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemReviewSetsItemQueriesItemEdiscoveryApplyTagsRequestBuilder) {
    m := &EdiscoveryCasesItemReviewSetsItemQueriesItemEdiscoveryApplyTagsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/compliance/ediscovery/cases/{case%2Did}/reviewSets/{reviewSet%2Did}/queries/{reviewSetQuery%2Did}/ediscovery.applyTags", pathParameters),
    }
    return m
}
// NewEdiscoveryCasesItemReviewSetsItemQueriesItemEdiscoveryApplyTagsRequestBuilder instantiates a new EdiscoveryApplyTagsRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemReviewSetsItemQueriesItemEdiscoveryApplyTagsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemReviewSetsItemQueriesItemEdiscoveryApplyTagsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEdiscoveryCasesItemReviewSetsItemQueriesItemEdiscoveryApplyTagsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action applyTags
func (m *EdiscoveryCasesItemReviewSetsItemQueriesItemEdiscoveryApplyTagsRequestBuilder) Post(ctx context.Context, body EdiscoveryCasesItemReviewSetsItemQueriesItemEdiscoveryApplyTagsApplyTagsPostRequestBodyable, requestConfiguration *EdiscoveryCasesItemReviewSetsItemQueriesItemEdiscoveryApplyTagsRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation invoke action applyTags
func (m *EdiscoveryCasesItemReviewSetsItemQueriesItemEdiscoveryApplyTagsRequestBuilder) ToPostRequestInformation(ctx context.Context, body EdiscoveryCasesItemReviewSetsItemQueriesItemEdiscoveryApplyTagsApplyTagsPostRequestBodyable, requestConfiguration *EdiscoveryCasesItemReviewSetsItemQueriesItemEdiscoveryApplyTagsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
