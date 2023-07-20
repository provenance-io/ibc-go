package types

import errorsmod "cosmossdk.io/errors"

var (
	ErrInvalid         = errorsmod.Register(ModuleName, 1, "invalid")
	ErrInvalidData     = errorsmod.Register(ModuleName, 2, "invalid data")
	ErrInvalidCodeHash = errorsmod.Register(ModuleName, 3, "invalid code hash")
	// Wasm specific
	ErrWasmEmptyCode        = errorsmod.Register(ModuleName, 4, "empty wasm code")
	ErrWasmCodeTooLarge     = errorsmod.Register(ModuleName, 5, "wasm code too large")
	ErrWasmCodeExists       = errorsmod.Register(ModuleName, 6, "wasm code already exists")
	ErrWasmCodeHashNotFound = errorsmod.Register(ModuleName, 7, "wasm code hash not found")
)
