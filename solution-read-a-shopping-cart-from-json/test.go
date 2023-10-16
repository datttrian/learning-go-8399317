// This is how your code will be called.
// Your answer should be a slice containing cartItem objects.
// You can edit this code to try different testing cases.
jsonString :=
	`[{"name":"apple","price":2.99,"quantity":3},
  	 {"name":"orange","price":1.50,"quantity":8},
  	 {"name":"banana","price":0.49,"quantity":12}]`

result := getCartFromJson(jsonString)
