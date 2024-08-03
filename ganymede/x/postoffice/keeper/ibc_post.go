package keeper

import (
	"log"
	"errors"
	"strings"
	"strconv"
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v7/modules/core/24-host"
	"github.com/jack139/ganymede/ganymede/x/postoffice/types"
)

// TransmitIbcPostPacket transmits the packet over IBC with the specified source port and source channel
func (k Keeper) TransmitIbcPostPacket(
	ctx sdk.Context,
	packetData types.IbcPostPacketData,
	sourcePort,
	sourceChannel string,
	timeoutHeight clienttypes.Height,
	timeoutTimestamp uint64,
) (uint64, error) {
	channelCap, ok := k.scopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(sourcePort, sourceChannel))
	if !ok {
		return 0, sdkerrors.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
	}

	packetBytes, err := packetData.GetBytes()
	if err != nil {
		return 0, sdkerrors.Wrapf(sdkerrors.ErrJSONMarshal, "cannot marshal the packet: %w", err)
	}

	return k.channelKeeper.SendPacket(ctx, channelCap, sourcePort, sourceChannel, timeoutHeight, timeoutTimestamp, packetBytes)
}

// OnRecvIbcPostPacket processes packet reception
func (k Keeper) OnRecvIbcPostPacket(ctx sdk.Context, packet channeltypes.Packet, data types.IbcPostPacketData) (packetAck types.IbcPostPacketAck, err error) {
	// validate packet data upon receiving
	if err := data.ValidateBasic(); err != nil {
		return packetAck, err
	}

	log.Println("OnRecvIbcPostPacket()", data)

	var sender, receiver, senderInfo, payload string

	if strings.HasPrefix(data.Content, "{") {
		// 解析 content
		// content 的 json 格式： { "sender" : "", "receiver" : "", "payload" : "", "sender_info" : "" }
		var contentData map[string]interface{}
		if err := json.Unmarshal([]byte(data.Content), &contentData); err != nil {
			return packetAck, err
		}

		sender = contentData["sender"].(string)
		receiver = contentData["receiver"].(string)
		senderInfo = contentData["sender_info"].(string)
		payload = contentData["payload"].(string)
	} else {
		payload = data.Content
	}

	// 添加到收到的 post
	id := k.AppendPost(
		ctx,
		types.Post{
			Receiver:   receiver,
			Sender:     sender,
			SenderInfo: senderInfo,
			FromChain:  packet.SourcePort + "-" + packet.SourceChannel, // + "-" + data.Creator,
			Title:      data.Title,
			Payload:    payload,
			SentDate:   data.SentDate,
		},
	)

	packetAck.PostID = strconv.FormatUint(id, 10)

	return packetAck, nil
}

// OnAcknowledgementIbcPostPacket responds to the the success or failure of a packet
// acknowledgement written on the receiving chain.
func (k Keeper) OnAcknowledgementIbcPostPacket(ctx sdk.Context, packet channeltypes.Packet, data types.IbcPostPacketData, ack channeltypes.Acknowledgement) error {
	switch dispatchedAck := ack.Response.(type) {
	case *channeltypes.Acknowledgement_Error:

		// TODO: failed acknowledgement logic
		_ = dispatchedAck.Error

		return nil
	case *channeltypes.Acknowledgement_Result:
		// Decode the packet acknowledgment
		var packetAck types.IbcPostPacketAck

		if err := types.ModuleCdc.UnmarshalJSON(dispatchedAck.Result, &packetAck); err != nil {
			// The counter-party module doesn't implement the correct acknowledgment format
			return errors.New("cannot unmarshal acknowledgment")
		}

		log.Println("OnAcknowledgementIbcPostPacket()", data)

		var sender, receiver, payload string

		if strings.HasPrefix(data.Content, "{") {
			// 解析 content
			// content 的 json 格式： { "sender" : "", "receiver" : "", "payload" : "", "sender_info" : "" }
			var contentData map[string]interface{}
			if err := json.Unmarshal([]byte(data.Content), &contentData); err != nil {
				return errors.New("OnAcknowledgementIbcPostPacket() unmarshal fail: " + err.Error())
			}

			sender = contentData["sender"].(string)
			receiver = contentData["receiver"].(string)
			payload = contentData["payload"].(string)
		} else {
			payload = data.Content
		}

		// 添加已发送
		k.AppendSentPost(
			ctx,
			types.SentPost{
				PostID:   packetAck.PostID,
				Receiver: receiver,
				Sender:   sender,
				ToChain:  packet.DestinationPort + "-" + packet.DestinationChannel, // + "-" + data.Creator,
				Title:    data.Title,
				Payload:  payload,
				SentDate: data.SentDate,
			},
		)

		return nil
	default:
		// The counter-party module doesn't implement the correct acknowledgment format
		return errors.New("invalid acknowledgment format")
	}
}

// OnTimeoutIbcPostPacket responds to the case where a packet has not been transmitted because of a timeout
func (k Keeper) OnTimeoutIbcPostPacket(ctx sdk.Context, packet channeltypes.Packet, data types.IbcPostPacketData) error {

	log.Println("OnTimeoutIbcPostPacket()", data)
	
	var sender, receiver string

	if strings.HasPrefix(data.Content, "{") {
		// 解析 content
		// content 的 json 格式： { "sender" : "", "receiver" : "", "payload" : "", "sender_info" : "" }

		var contentData map[string]interface{}
		if err := json.Unmarshal([]byte(data.Content), &contentData); err != nil {
			return errors.New("OnTimeoutIbcPostPacket() unmarshal fail: " + err.Error())
		}

		sender = contentData["sender"].(string)
		receiver = contentData["receiver"].(string)
	}

	// 添加 超时 post
	k.AppendTimedoutPost(
		ctx,
		types.TimedoutPost{
			Receiver: receiver,
			Sender:   sender,
			Title:    data.Title,
			ToChain:  packet.DestinationPort + "-" + packet.DestinationChannel, // + "-" + data.Creator,
			SentDate: data.SentDate,
		},
	)

	return nil
}
