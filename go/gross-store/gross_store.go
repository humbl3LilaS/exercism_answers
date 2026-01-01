package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	result := map[string]int{}
	result["quarter_of_a_dozen"] = 3
	result["half_of_a_dozen"] = 6
	result["dozen"] = 12
	result["small_gross"] = 120
	result["gross"] = 144
	result["great_gross"] = 1728
	return result
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	cUnit, isValidUnit := units[unit]
	if !isValidUnit {
		return false
	}
	bill[item] += cUnit
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	cUnit, isValidUnit := units[unit]

	// return false if unit is not in units map
	if !isValidUnit {
		return false
	}

	qunatitiy, isInBill := bill[item]
	//  return false if item is not in bill map
	if !isInBill {
		return false
	}

	// return false remove item quantity is greater than the quantity already exist in bill map
	if cUnit > qunatitiy {
		return false
	}

	// remove the entire item from if the final quantity is zero
	if cUnit == qunatitiy {
		delete(bill, item)
		return true
	}
	bill[item] -= cUnit
	return true

}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	quantity, ok := bill[item]
	if !ok {
		return 0, false
	}
	return quantity, true
}
