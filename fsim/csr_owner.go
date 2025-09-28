// SPDX-FileCopyrightText: (C) 2024 Intel Corporation
// SPDX-License-Identifier: Apache 2.0

package fsim

import (
	"context"
	"io"

	"github.com/fido-device-onboard/go-fdo/serviceinfo"
)

type SimpleEnrollResponse struct {
}

func (s SimpleEnrollResponse) HandleInfo(ctx context.Context, messageName string, messageBody io.Reader) error {
	//TODO implement me
	panic("implement me")
}

func (s SimpleEnrollResponse) ProduceInfo(ctx context.Context, producer *serviceinfo.Producer) (blockPeer, moduleDone bool, _ error) {
	//TODO implement me
	panic("implement me")
}

var _ serviceinfo.OwnerModule = (*SimpleEnrollResponse)(nil)
