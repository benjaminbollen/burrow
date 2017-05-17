pragma solidity ^0.4.4;

contract Store {

	struct Item {
		uint id;
		// string name;
		// bytes8 info;
	}

	mapping (uint=>Item) items;
	uint counter;

	function saveItem(uint id, uint64[2] nums) {
		var newItem = Item(id);
		items[id] = newItem;
		counter++;
	}

	function getCount() returns (uint count) {
		return counter;
	}
}

// contract Storage {
//   int[2] storedData;

//   function set(int[2] x) {
//     storedData = x;
//   }

//   function get() constant returns (int[2] retVal) {
//     return storedData;
//   }
// }