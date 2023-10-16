// This is how your code will be called.
// Your answer should be the largest value in the numbers array.
// You can edit this code to try different testing cases.
var cart []cartItem
var apples = cartItem{"apple", 2.99, 3}
var oranges = cartItem{"orange", 1.50, 8}
var bananas = cartItem{"banana", .49, 12}
cart = append(cart, apples, oranges, bananas)
result := calculateTotal(cart)
