/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package vvp_client

import (
	"errors"
	"fmt"
)

const (
	InSyncState        string = "IN SYNC"
	OutOfSyncState     string = "OUT OF SYNC"
	SynchronizingState string = "SYNCHRONIZING"

	CommunityEditionNamespace = "default"
)

type retryableError struct {
	msg string
}

func (err retryableError) Error() string {
	return err.msg
}

func NewRetryableError(err error) error {
	return retryableError{
		msg: err.Error(),
	}
}

var ErrRetryable = NewRetryableError(errors.New("error"))

func FormatOutOfSync(err error) string {
	return fmt.Sprintf("%s: %s", OutOfSyncState, err.Error())
}
