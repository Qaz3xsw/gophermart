package core

import "github.com/Qaz3xsw/gophermart/internal/sharedkernel"

type RewardMechanic struct {
	id           string
	match        string
	rewardType   string
	rewardPoints int
}

func NewRewardMechanic(match string, rewardPoints int, rewardType string) RewardMechanic {
	return RewardMechanic{
		id:           sharedkernel.NewUUID(),
		match:        match,
		rewardPoints: rewardPoints,
		rewardType:   rewardType,
	}
}
