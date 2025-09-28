// SPDX-FileCopyrightText: (C) 2024 Intel Corporation
// SPDX-License-Identifier: Apache 2.0

package fsim

import (
	"context"
	"io"

	"github.com/fido-device-onboard/go-fdo/serviceinfo"
)

type SimpleEnrollRequest struct {
}

func (s SimpleEnrollRequest) Transition(active bool) error {
	//TODO implement me
	panic("implement me")
}

func (s SimpleEnrollRequest) Receive(ctx context.Context, messageName string, messageBody io.Reader, respond func(message string) io.Writer, yield func()) error {
	//TODO implement me
	panic("implement me")
}

func (s SimpleEnrollRequest) Yield(ctx context.Context, respond func(message string) io.Writer, yield func()) error {
	//TODO implement me
	panic("implement me")
}

var _ serviceinfo.DeviceModule = (*SimpleEnrollRequest)(nil)
