package keeper

import (
	"context"
	"errors"
	"fmt"

	"projectBit/x/projectbit/types"

	"cosmossdk.io/collections"
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"strconv" // for convert id to string in events
)

func (k msgServer) CreatePost(ctx context.Context, msg *types.MsgCreatePost) (*types.MsgCreatePostResponse, error) {
	// SDK Context
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	if len(msg.Title) < 3 {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "the title is short")
	}

	if len(msg.Body) == 0 {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "the body is empty")
	}

	// 1. Address Conversion
	creatorAddr, err := k.addressCodec.StringToBytes(msg.Creator)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid address: %s", err))
	}

	accAddress := sdk.AccAddress(creatorAddr)

	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid address: %s", err))
	}

	// 3. BankKeeper
    if msg.Amount.IsPositive() {
        err = k.BankKeeper.SendCoinsFromAccountToModule(ctx, accAddress, types.ModuleName, sdk.NewCoins(msg.Amount))
        if err != nil {
            return nil, errorsmod.Wrap(sdkerrors.ErrInsufficientFunds, "failed to pay for post creation")
        }
    }

	nextId, err := k.PostSeq.Next(ctx)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "failed to get next id")
	}

	var post = types.Post{
		Id:      nextId,
		Creator: msg.Creator,
		Title:   msg.Title,
		Body:    msg.Body,
	}

	if err = k.Post.Set(
		ctx,
		nextId,
		post,
	); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to set post")
	}

	// Emit events
	sdkCtx.EventManager().EmitEvent(
        sdk.NewEvent(types.PostCreatedEventType,
            sdk.NewAttribute(types.PostCreatorAttribute, msg.Creator),
            sdk.NewAttribute(types.PostIdAttribute, strconv.FormatUint(nextId, 10)),
            sdk.NewAttribute(types.PostTitleAttribute, msg.Title),
        ),
    )

	return &types.MsgCreatePostResponse{
		Id: nextId,
	}, nil
}

func (k msgServer) UpdatePost(ctx context.Context, msg *types.MsgUpdatePost) (*types.MsgUpdatePostResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid address: %s", err))
	}

	var post = types.Post{
		Creator: msg.Creator,
		Id:      msg.Id,
		Title:   msg.Title,
		Body:    msg.Body,
	}

	// Checks that the element exists
	val, err := k.Post.Get(ctx, msg.Id)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
		}

		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to get post")
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	if err := k.Post.Set(ctx, msg.Id, post); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to update post")
	}

	// Emit events
	sdkCtx.EventManager().EmitEvent(
        sdk.NewEvent(types.PostUpdatedEventType,
            sdk.NewAttribute(types.PostCreatorAttribute, msg.Creator),
            sdk.NewAttribute(types.PostIdAttribute, strconv.FormatUint(msg.Id, 10)),
        ),
    )

	return &types.MsgUpdatePostResponse{}, nil
}

func (k msgServer) DeletePost(ctx context.Context, msg *types.MsgDeletePost) (*types.MsgDeletePostResponse, error) {

	sdkCtx := sdk.UnwrapSDKContext(ctx)

	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid address: %s", err))
	}

	// Checks that the element exists
	val, err := k.Post.Get(ctx, msg.Id)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
		}

		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to get post")
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	if err := k.Post.Remove(ctx, msg.Id); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to delete post")
	}

	sdkCtx.EventManager().EmitEvent(
        sdk.NewEvent(types.PostDeletedEventType,
            sdk.NewAttribute(types.PostCreatorAttribute, msg.Creator),
            sdk.NewAttribute(types.PostIdAttribute, strconv.FormatUint(msg.Id, 10)),
        ),
    )

	return &types.MsgDeletePostResponse{}, nil
}
