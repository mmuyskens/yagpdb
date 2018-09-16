// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("PremiumCodes", testPremiumCodes)
	t.Run("PremiumSlots", testPremiumSlots)
}

func TestDelete(t *testing.T) {
	t.Run("PremiumCodes", testPremiumCodesDelete)
	t.Run("PremiumSlots", testPremiumSlotsDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("PremiumCodes", testPremiumCodesQueryDeleteAll)
	t.Run("PremiumSlots", testPremiumSlotsQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("PremiumCodes", testPremiumCodesSliceDeleteAll)
	t.Run("PremiumSlots", testPremiumSlotsSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("PremiumCodes", testPremiumCodesExists)
	t.Run("PremiumSlots", testPremiumSlotsExists)
}

func TestFind(t *testing.T) {
	t.Run("PremiumCodes", testPremiumCodesFind)
	t.Run("PremiumSlots", testPremiumSlotsFind)
}

func TestBind(t *testing.T) {
	t.Run("PremiumCodes", testPremiumCodesBind)
	t.Run("PremiumSlots", testPremiumSlotsBind)
}

func TestOne(t *testing.T) {
	t.Run("PremiumCodes", testPremiumCodesOne)
	t.Run("PremiumSlots", testPremiumSlotsOne)
}

func TestAll(t *testing.T) {
	t.Run("PremiumCodes", testPremiumCodesAll)
	t.Run("PremiumSlots", testPremiumSlotsAll)
}

func TestCount(t *testing.T) {
	t.Run("PremiumCodes", testPremiumCodesCount)
	t.Run("PremiumSlots", testPremiumSlotsCount)
}

func TestInsert(t *testing.T) {
	t.Run("PremiumCodes", testPremiumCodesInsert)
	t.Run("PremiumCodes", testPremiumCodesInsertWhitelist)
	t.Run("PremiumSlots", testPremiumSlotsInsert)
	t.Run("PremiumSlots", testPremiumSlotsInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("PremiumCodeToPremiumSlotUsingSlot", testPremiumCodeToOnePremiumSlotUsingSlot)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("PremiumSlotToSlotPremiumCodes", testPremiumSlotToManySlotPremiumCodes)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("PremiumCodeToPremiumSlotUsingSlotPremiumCodes", testPremiumCodeToOneSetOpPremiumSlotUsingSlot)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {
	t.Run("PremiumCodeToPremiumSlotUsingSlotPremiumCodes", testPremiumCodeToOneRemoveOpPremiumSlotUsingSlot)
}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("PremiumSlotToSlotPremiumCodes", testPremiumSlotToManyAddOpSlotPremiumCodes)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {
	t.Run("PremiumSlotToSlotPremiumCodes", testPremiumSlotToManySetOpSlotPremiumCodes)
}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {
	t.Run("PremiumSlotToSlotPremiumCodes", testPremiumSlotToManyRemoveOpSlotPremiumCodes)
}

func TestReload(t *testing.T) {
	t.Run("PremiumCodes", testPremiumCodesReload)
	t.Run("PremiumSlots", testPremiumSlotsReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("PremiumCodes", testPremiumCodesReloadAll)
	t.Run("PremiumSlots", testPremiumSlotsReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("PremiumCodes", testPremiumCodesSelect)
	t.Run("PremiumSlots", testPremiumSlotsSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("PremiumCodes", testPremiumCodesUpdate)
	t.Run("PremiumSlots", testPremiumSlotsUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("PremiumCodes", testPremiumCodesSliceUpdateAll)
	t.Run("PremiumSlots", testPremiumSlotsSliceUpdateAll)
}