package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSitesItemListsItemSubscriptionsItemReauthorizeRequestBuilder provides operations to call the reauthorize method.
type ItemSitesItemListsItemSubscriptionsItemReauthorizeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemListsItemSubscriptionsItemReauthorizeRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemListsItemSubscriptionsItemReauthorizeRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSitesItemListsItemSubscriptionsItemReauthorizeRequestBuilderInternal instantiates a new ReauthorizeRequestBuilder and sets the default values.
func NewItemSitesItemListsItemSubscriptionsItemReauthorizeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemListsItemSubscriptionsItemReauthorizeRequestBuilder) {
    m := &ItemSitesItemListsItemSubscriptionsItemReauthorizeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/lists/{list%2Did}/subscriptions/{subscription%2Did}/reauthorize", pathParameters),
    }
    return m
}
// NewItemSitesItemListsItemSubscriptionsItemReauthorizeRequestBuilder instantiates a new ReauthorizeRequestBuilder and sets the default values.
func NewItemSitesItemListsItemSubscriptionsItemReauthorizeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemListsItemSubscriptionsItemReauthorizeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemListsItemSubscriptionsItemReauthorizeRequestBuilderInternal(urlParams, requestAdapter)
}
// Post reauthorize a subscription when you receive a reauthorizationRequired challenge. This API is available in the following national cloud deployments.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/subscription-reauthorize?view=graph-rest-1.0
func (m *ItemSitesItemListsItemSubscriptionsItemReauthorizeRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemSitesItemListsItemSubscriptionsItemReauthorizeRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation reauthorize a subscription when you receive a reauthorizationRequired challenge. This API is available in the following national cloud deployments.
func (m *ItemSitesItemListsItemSubscriptionsItemReauthorizeRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemListsItemSubscriptionsItemReauthorizeRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.TryAdd("Accept", "application/json, application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemSitesItemListsItemSubscriptionsItemReauthorizeRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemListsItemSubscriptionsItemReauthorizeRequestBuilder) {
    return NewItemSitesItemListsItemSubscriptionsItemReauthorizeRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
