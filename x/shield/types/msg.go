package types

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// MsgCreatePool defines the attributes of a create-pool transaction.
type MsgCreatePool struct {
	From           sdk.AccAddress `json:"from" yaml:"from"`
	Shield         sdk.Coins      `json:"shield" yaml:"shield"`
	Deposit        MixedCoins     `json:"deposit" yaml:"deposit"`
	Sponsor        string         `json:"sponsor" yaml:"sponsor"`
	TimeOfCoverage int64          `json:"time_of_coverage" yaml:"time_of_coverage"`
}

// NewMsgCreatePool creates a new NewMsgCreatePool instance.
func NewMsgCreatePool(accAddr sdk.AccAddress, shield sdk.Coins, deposit MixedCoins, sponsor string, time int64) MsgCreatePool {
	return MsgCreatePool{
		From:           accAddr,
		Shield:         shield,
		Deposit:        deposit,
		Sponsor:        sponsor,
		TimeOfCoverage: time,
	}
}

// Route implements the sdk.Msg interface.
func (msg MsgCreatePool) Route() string { return RouterKey }

// Type implements the sdk.Msg interface
func (msg MsgCreatePool) Type() string { return EventTypeCreatePool }

// GetSigners implements the sdk.Msg interface
func (msg MsgCreatePool) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.From}
}

// GetSignBytes implements the sdk.Msg interface.
func (msg MsgCreatePool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgCreatePool) ValidateBasic() error {
	if msg.From.Empty() {
		return ErrEmptySender
	}
	if strings.TrimSpace(msg.Sponsor) == "" {
		return ErrEmptySponsor
	}
	if msg.Deposit.Native.IsZero() || !msg.Deposit.Native.IsValid() {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "native amount: %s", msg.Deposit.Native)
	}
	if msg.Deposit.Foreign.IsZero() || !msg.Deposit.Foreign.IsValid() {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "foreign amount %s", msg.Deposit.Foreign)
	}
	if !msg.Shield.IsValid() || msg.Shield.IsZero() {
		return ErrNoShield
	}
	if msg.TimeOfCoverage <= 0 {
		return ErrInvalidDuration
	}
	return nil
}

// MsgUpdatePool defines the attributes of a shield pool update transaction.
type MsgUpdatePool struct {
	From           sdk.AccAddress `json:"from" yaml:"from"`
	Shield         sdk.Coins      `json:"Shield" yaml:"Shield"`
	Deposit        MixedCoins     `json:"deposit" yaml:"deposit"`
	PoolID         uint64         `json:"pool_id" yaml:"pool_id"`
	AdditionalTime int64          `json:"additional_period" yaml:"additional_period"`
}

// NewMsgUpdatePool creates a new MsgUpdatePool instance.
func NewMsgUpdatePool(accAddr sdk.AccAddress, shield sdk.Coins, deposit MixedCoins, id uint64, time int64) MsgUpdatePool {
	return MsgUpdatePool{
		From:           accAddr,
		Shield:         shield,
		Deposit:        deposit,
		PoolID:         id,
		AdditionalTime: time,
	}
}

// Route implements the sdk.Msg interface.
func (msg MsgUpdatePool) Route() string { return RouterKey }

// Type implements the sdk.Msg interface
func (msg MsgUpdatePool) Type() string { return EventTypeUpdatePool }

// GetSigners implements the sdk.Msg interface
func (msg MsgUpdatePool) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.From}
}

// GetSignBytes implements the sdk.Msg interface.
func (msg MsgUpdatePool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgUpdatePool) ValidateBasic() error {
	if msg.From.Empty() {
		return ErrEmptySender
	}
	if msg.PoolID == 0 {
		return ErrInvalidPoolID
	}
	if !(msg.Deposit.Native.IsValid() && msg.Deposit.Foreign.IsValid()) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "invalid deposit")
	}
	if !msg.Shield.IsValid() {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "invalid shield")
	}
	if msg.AdditionalTime <= 0 {
		return ErrInvalidDuration
	}
	return nil
}

// MsgPausePool defines the attributes of a pausing a shield pool.
type MsgPausePool struct {
	From   sdk.AccAddress `json:"from" yaml:"from"`
	PoolID uint64         `json:"pool_id" yaml:"pool_id"`
}

// NewMsgPausePool creates a new NewMsgPausePool instance.
func NewMsgPausePool(accAddr sdk.AccAddress, id uint64) MsgPausePool {
	return MsgPausePool{
		From:   accAddr,
		PoolID: id,
	}
}

// Route implements the sdk.Msg interface.
func (msg MsgPausePool) Route() string { return RouterKey }

// Type implements the sdk.Msg interface
func (msg MsgPausePool) Type() string { return EventTypePausePool }

// GetSigners implements the sdk.Msg interface
func (msg MsgPausePool) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.From}
}

// GetSignBytes implements the sdk.Msg interface.
func (msg MsgPausePool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgPausePool) ValidateBasic() error {
	if msg.From.Empty() {
		return ErrEmptySender
	}
	if msg.PoolID == 0 {
		return ErrInvalidPoolID
	}
	return nil
}

// MsgResumePool defines the attributes of a resuming a shield pool.
type MsgResumePool struct {
	From   sdk.AccAddress `json:"from" yaml:"from"`
	PoolID uint64         `json:"pool_id" yaml:"pool_id"`
}

// NewMsgResumePool creates a new NewMsgResumePool instance.
func NewMsgResumePool(accAddr sdk.AccAddress, id uint64) MsgResumePool {
	return MsgResumePool{
		From:   accAddr,
		PoolID: id,
	}
}

// Route implements the sdk.Msg interface.
func (msg MsgResumePool) Route() string { return RouterKey }

// Type implements the sdk.Msg interface
func (msg MsgResumePool) Type() string { return EventTypeResumePool }

// GetSigners implements the sdk.Msg interface
func (msg MsgResumePool) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.From}
}

// GetSignBytes implements the sdk.Msg interface.
func (msg MsgResumePool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgResumePool) ValidateBasic() error {
	if msg.From.Empty() {
		return ErrEmptySender
	}
	if msg.PoolID == 0 {
		return ErrInvalidPoolID
	}
	return nil
}

// MsgDepositCollateral defines the attributes of a depositing collaterals.
type MsgDepositCollateral struct {
	From       sdk.AccAddress `json:"sender" yaml:"sender"`
	PoolID     uint64         `json:"pool_id" yaml:"pool_id"`
	Collateral sdk.Coin       `json:"collateral" yaml:"collateral"`
}

// NewMsgDepositCollateral creates a new MsgDepositCollateral instance.
func NewMsgDepositCollateral(sender sdk.AccAddress, id uint64, collateral sdk.Coin) MsgDepositCollateral {
	return MsgDepositCollateral{
		From:       sender,
		PoolID:     id,
		Collateral: collateral,
	}
}

// Route implements the sdk.Msg interface.
func (msg MsgDepositCollateral) Route() string { return RouterKey }

// Type implements the sdk.Msg interface
func (msg MsgDepositCollateral) Type() string { return "deposit_collateral" }

// GetSigners implements the sdk.Msg interface
func (msg MsgDepositCollateral) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.From}
}

// GetSignBytes implements the sdk.Msg interface.
func (msg MsgDepositCollateral) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgDepositCollateral) ValidateBasic() error {
	if msg.From.Empty() {
		return ErrEmptySender
	}
	if msg.PoolID == 0 {
		return ErrInvalidPoolID
	}
	if !msg.Collateral.IsValid() || msg.Collateral.IsZero() {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "Collateral amount: %s", msg.Collateral)
	}
	return nil
}

// NewMsgWithdrawCollateral defines the attributes of a withdrawing collaterals.
type MsgWithdrawCollateral struct {
	From       sdk.AccAddress `json:"sender" yaml:"sender"`
	PoolID     uint64         `json:"pool_id" yaml:"pool_id"`
	Collateral sdk.Coin       `json:"collateral" yaml:"collateral"`
}

// NewMsgDepositCollateral creates a new MsgDepositCollateral instance.
func NewMsgWithdrawCollateral(sender sdk.AccAddress, id uint64, collateral sdk.Coin) MsgWithdrawCollateral {
	return MsgWithdrawCollateral{
		From:       sender,
		PoolID:     id,
		Collateral: collateral,
	}
}

// Route implements the sdk.Msg interface.
func (msg MsgWithdrawCollateral) Route() string { return RouterKey }

// Type implements the sdk.Msg interface
func (msg MsgWithdrawCollateral) Type() string { return "withdraw_collateral" }

// GetSigners implements the sdk.Msg interface
func (msg MsgWithdrawCollateral) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.From}
}

// GetSignBytes implements the sdk.Msg interface.
func (msg MsgWithdrawCollateral) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgWithdrawCollateral) ValidateBasic() error {
	if msg.PoolID == 0 {
		return ErrInvalidPoolID
	}
	return nil
}

type MsgWithdrawRewards struct {
	From sdk.AccAddress `json:"sender" yaml:"sender"`
}

// NewMsgWithdrawRewards creates a new MsgWithdrawRewards instance.
func NewMsgWithdrawRewards(sender sdk.AccAddress) MsgWithdrawRewards {
	return MsgWithdrawRewards{
		From: sender,
	}
}

func (msg MsgWithdrawRewards) Route() string { return RouterKey }

// Type implements the sdk.Msg interface
func (msg MsgWithdrawRewards) Type() string { return EventTypeWithdrawRewards }

// GetSigners implements the sdk.Msg interface
func (msg MsgWithdrawRewards) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.From}
}

// GetSignBytes implements the sdk.Msg interface.
func (msg MsgWithdrawRewards) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgWithdrawRewards) ValidateBasic() error {
	if msg.From.Empty() {
		return ErrEmptySender
	}
	return nil
}

type MsgWithdrawForeignRewards struct {
	From   sdk.AccAddress `json:"sender" yaml:"sender"`
	Denom  string         `json:"denom" yaml:"denom"`
	ToAddr string         `json:"to_addr" yaml:"to_addr"`
}

// NewMsgWithdrawForeignRewards creates a new MsgWithdrawForeignRewards instance.
func NewMsgWithdrawForeignRewards(sender sdk.AccAddress, denom, toAddr string) MsgWithdrawForeignRewards {
	return MsgWithdrawForeignRewards{
		From:   sender,
		Denom:  denom,
		ToAddr: toAddr,
	}
}

// Route implements the sdk.Msg interface.
func (msg MsgWithdrawForeignRewards) Route() string { return RouterKey }

// Type implements the sdk.Msg interface
func (msg MsgWithdrawForeignRewards) Type() string { return EventTypeWithdrawForeignRewards }

// GetSigners implements the sdk.Msg interface
func (msg MsgWithdrawForeignRewards) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.From}
}

// GetSignBytes implements the sdk.Msg interface.
func (msg MsgWithdrawForeignRewards) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgWithdrawForeignRewards) ValidateBasic() error {
	if msg.From.Empty() {
		return ErrEmptySender
	}
	if strings.TrimSpace(msg.ToAddr) == "" {
		return ErrInvalidToAddr
	}
	return nil
}

type MsgClearPayouts struct {
	From  sdk.AccAddress `json:"sender" yaml:"sender"`
	Denom string         `json:"denom" yaml:"denom"`
}

// NewMsgClearPayouts creates a new MsgClearPayouts instance.
func NewMsgClearPayouts(sender sdk.AccAddress, denom string) MsgClearPayouts {
	return MsgClearPayouts{
		From:  sender,
		Denom: denom,
	}
}

// Route implements the sdk.Msg interface.
func (msg MsgClearPayouts) Route() string { return RouterKey }

// Type implements the sdk.Msg interface
func (msg MsgClearPayouts) Type() string { return EventTypeClearPayouts }

// GetSigners implements the sdk.Msg interface
func (msg MsgClearPayouts) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.From}
}

// GetSignBytes implements the sdk.Msg interface.
func (msg MsgClearPayouts) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgClearPayouts) ValidateBasic() error {
	if msg.From.Empty() {
		return ErrEmptySender
	}
	if err := sdk.ValidateDenom(msg.Denom); err != nil {
		return ErrInvalidDenom
	}
	return nil
}

// MsgPurchaseShield defines the attributes of purchase shield transaction
type MsgPurchaseShield struct {
	PoolID      uint64         `json:"pool_id" yaml:"pool_id"`
	Shield      sdk.Coins      `json:"shield" yaml:"shield"`
	Description string         `json:"description" yaml:"description"`
	From        sdk.AccAddress `json:"from" yaml:"from"`
	Simulate    bool           `json:"simulate" yaml:"simulate"`
	SimTxHash   []byte         `json:"sim_txhash" yaml:"sim_txhash"`
}

// NewMsgPurchaseShield creates a new MsgPurchaseShield instance.
func NewMsgPurchaseShield(poolID uint64, shield sdk.Coins, description string, from sdk.AccAddress) MsgPurchaseShield {
	return MsgPurchaseShield{
		PoolID:      poolID,
		Shield:      shield,
		Description: description,
		From:        from,
	}
}

// Route implements the sdk.Msg interface.
func (msg MsgPurchaseShield) Route() string { return RouterKey }

// Type implements the sdk.Msg interface.
func (msg MsgPurchaseShield) Type() string { return EventTypePurchaseShield }

// GetSigners implements the sdk.Msg interface.
func (msg MsgPurchaseShield) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.From}
}

// GetSignBytes implements the sdk.Msg interface.
func (msg MsgPurchaseShield) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgPurchaseShield) ValidateBasic() error {
	if msg.PoolID == 0 {
		return ErrInvalidPoolID
	}
	if !msg.Shield.IsValid() || msg.Shield.IsZero() {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "shield amount: %s", msg.Shield)
	}
	if strings.TrimSpace(msg.Description) == "" {
		return ErrPurchaseMissingDescription
	}
	if msg.From.Empty() {
		return ErrEmptySender
	}
	return nil
}

// MsgWithdrawReimburse defines the attributes of withdraw reimbursement transaction.
type MsgWithdrawReimbursement struct {
	ProposalID uint64         `json:"proposal_id" yaml:"proposal_id"`
	From       sdk.AccAddress `json:"from" yaml:"from"`
}

// NewMsgWithdrawReimbursement creates a new MsgWithdrawReimbursement instance.
func NewMsgWithdrawReimbursement(proposalID uint64, from sdk.AccAddress) MsgWithdrawReimbursement {
	return MsgWithdrawReimbursement{
		ProposalID: proposalID,
		From:       from,
	}
}

// Route implements the sdk.Msg interface.
func (msg MsgWithdrawReimbursement) Route() string { return RouterKey }

// Type implements the sdk.Msg interface.
func (msg MsgWithdrawReimbursement) Type() string { return EventTypeWithdrawReimbursement }

// GetSigners implements the sdk.Msg interface.
func (msg MsgWithdrawReimbursement) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.From}
}

// GetSignBytes implements the sdk.Msg interface.
func (msg MsgWithdrawReimbursement) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgWithdrawReimbursement) ValidateBasic() error {
	return nil
}