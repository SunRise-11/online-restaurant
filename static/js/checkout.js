let orders= document.getElementsByClassName("ordering")
let totalCost= document.getElementsByClassName("cost")
let food= document.getElementById("food")
let total = document.getElementById("total") 
let footerprice = document.getElementById("footerprice") 
let headerprice = document.getElementById("totalfoot") 

// all meals and number of plates
const order=[]; //convert html collection to array
for (let i = 0; i < orders.length; i++) {
        order[i] = orders[i].value
      }
ordering= order.toString();
// Attach value to hidden input
if (ordering !=0){
      food.value = ordering
}

// get total cost
const price=[]; //convert html collection to array
for (let j = 0; j < totalCost.length; j++) {
        price[j] = totalCost[j].value
      }
var cost=0
var costy=0
for(var k=price.length; k--;){
        costy=parseInt(price[k])
        cost+=costy
}
// Attach value to hidden input
if (cost !=0){
        total.value = cost
        footerprice.innerHTML = '$'+ cost
        headerprice.innerHTML = cost
}
