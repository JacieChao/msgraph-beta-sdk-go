package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemChatsItemPermissionGrantsItemCheckMemberGroupsRequestBuilder provides operations to call the checkMemberGroups method.
type ItemChatsItemPermissionGrantsItemCheckMemberGroupsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemChatsItemPermissionGrantsItemCheckMemberGroupsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemChatsItemPermissionGrantsItemCheckMemberGroupsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemChatsItemPermissionGrantsItemCheckMemberGroupsRequestBuilderInternal instantiates a new CheckMemberGroupsRequestBuilder and sets the default values.
func NewItemChatsItemPermissionGrantsItemCheckMemberGroupsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemChatsItemPermissionGrantsItemCheckMemberGroupsRequestBuilder) {
    m := &ItemChatsItemPermissionGrantsItemCheckMemberGroupsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/chats/{chat%2Did}/permissionGrants/{resourceSpecificPermissionGrant%2Did}/checkMemberGroups", pathParameters),
    }
    return m
}
// NewItemChatsItemPermissionGrantsItemCheckMemberGroupsRequestBuilder instantiates a new CheckMemberGroupsRequestBuilder and sets the default values.
func NewItemChatsItemPermissionGrantsItemCheckMemberGroupsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemChatsItemPermissionGrantsItemCheckMemberGroupsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemChatsItemPermissionGrantsItemCheckMemberGroupsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action checkMemberGroups
func (m *ItemChatsItemPermissionGrantsItemCheckMemberGroupsRequestBuilder) Post(ctx context.Context, body ItemChatsItemPermissionGrantsItemCheckMemberGroupsPostRequestBodyable, requestConfiguration *ItemChatsItemPermissionGrantsItemCheckMemberGroupsRequestBuilderPostRequestConfiguration)(ItemChatsItemPermissionGrantsItemCheckMemberGroupsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemChatsItemPermissionGrantsItemCheckMemberGroupsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemChatsItemPermissionGrantsItemCheckMemberGroupsResponseable), nil
}
// ToPostRequestInformation invoke action checkMemberGroups
func (m *ItemChatsItemPermissionGrantsItemCheckMemberGroupsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemChatsItemPermissionGrantsItemCheckMemberGroupsPostRequestBodyable, requestConfiguration *ItemChatsItemPermissionGrantsItemCheckMemberGroupsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
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
