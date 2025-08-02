function allocateOrder(order, warehouses) {
    const result = [];
  
    for (let { name, inventory } of warehouses) {
      const warehouseAllocation = {};
      for (let [item, needed] of Object.entries(order)) {
        const available = inventory[item] || 0;

        const allocated = Math.min(needed, available);
        order[item] -= warehouseAllocation[name];
        if (allocated > 0)
          warehouseAllocation[name] = {
            ...warehouseAllocation[name],
            ...{ [item]: allocated },
          };
      }
      result.push(warehouseAllocation);
    }
  
    return result;
  }
  
  const order = { apple: 5, banana: 5, orange: 5 };
  const warehouses = [
    { name: "wh1", inventory: { apple: 10 } },
    { name: "wh2", inventory: { banana: 5, orange: 5, apple: 3 } },
  ];
  // expected
  // [
  //     { wh1: { apple: 2 } },
  //     { wh2: { banana: 5, orange: 5, apple: 3 } }
  //   ]
  
  console.log(JSON.stringify(allocateOrder(order, warehouses)));