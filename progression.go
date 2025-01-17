// Copyright 2023 Heroic Labs & Contributors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hiro

import (
	"context"

	"github.com/heroiclabs/nakama-common/runtime"
)

var (
	ErrProgressionNotFound             = runtime.NewError("progression not found", 3)                 // INVALID_ARGUMENT
	ErrProgressionNotAvailablePurchase = runtime.NewError("progression not available to purchase", 3) // INVALID_ARGUMENT
	ErrProgressionNotAvailableUpdate   = runtime.NewError("progression not available to update", 3)   // INVALID_ARGUMENT
	ErrProgressionNoCost               = runtime.NewError("progression no cost associated", 3)        // INVALID_ARGUMENT
	ErrProgressionNoCount              = runtime.NewError("progression no count associated", 3)       // INVALID_ARGUMENT
	ErrProgressionAlreadyUnlocked      = runtime.NewError("progression already unlocked", 3)          // INVALID_ARGUMENT
)

// ProgressionConfig is the data definition for a ProgressionSystem type.
type ProgressionConfig struct {
	Progressions map[string]*ProgressionConfigProgression `json:"progressions"`
}

type ProgressionConfigProgression struct {
	Name                 string                         `json:"name"`
	Description          string                         `json:"description"`
	Category             string                         `json:"category"`
	AdditionalProperties map[string]string              `json:"additional_properties"`
	Preconditions        *ProgressionPreconditionsBlock `json:"preconditions"`
}

// A ProgressionSystem is a gameplay system which represents a sequence of progression steps.
type ProgressionSystem interface {
	System

	// Get returns all or an optionally-filtered set of progressions for the given user.
	Get(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID string, lastKnownProgressions map[string]*Progression) (map[string]*Progression, map[string]*ProgressionDelta, error)

	// Purchase permanently unlocks a specified progression, if that progression supports this operation.
	Purchase(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID, progressionID string) (map[string]*Progression, error)

	// Update a specified progression, if that progression supports this operation.
	Update(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID, progressionID string, counts map[string]int64) (map[string]*Progression, error)
}
