package config

import (
	"github.com/shippinAI/wardenclyffe/management/server/types"
	"github.com/shippinAI/wardenclyffe/shared/management/proto"
)

func ExtendNetBirdConfig(peerID string, peerGroups []string, config *proto.NetbirdConfig, extraSettings *types.ExtraSettings) *proto.NetbirdConfig {
	return config
}
