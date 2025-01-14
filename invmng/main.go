package main

import (
	"fmt"
)

type Product struct {
    ID       int
    Name     string
    Quantity int
    Price    float64
}
// AddProduct adds a new product to the inventory
func AddProduct(inventory []Product, p Product) []Product {
    return append(inventory, p)
}

// ListProducts lists all products in the inventory
func ListProducts(inventory []Product) {
    fmt.Println("Inventory List:")
    for _, product := range inventory {
        fmt.Printf("ID: %d, Name: %s, Quantity: %d, Price: %.2f\n", product.ID, product.Name, product.Quantity, product.Price)
    }
}

// RemoveProduct removes a product by ID from the inventory
func RemoveProduct(inventory []Product, id int) []Product {
    for i, product := range inventory {
        if product.ID == id {
            // Remove product from slice by slicing the slice before and after the product
            return append(inventory[:i], inventory[i+1:]...)
        }
    }
    fmt.Println("Product not found!")
    return inventory
}



func main() {
    // Initialize empty inventory
    var inventory []Product

    // Adding products to the inventory
    inventory = AddProduct(inventory, Product{ID: 1, Name: "Laptop", Quantity: 10, Price: 999.99})
    inventory = AddProduct(inventory, Product{ID: 2, Name: "Smartphone", Quantity: 25, Price: 499.99})
    inventory = AddProduct(inventory, Product{ID: 3, Name: "Headphones", Quantity: 50, Price: 89.99})

    // List all products in the inventory
    ListProducts(inventory)

    // Remove a product from inventory (ID 2)
    inventory = RemoveProduct(inventory, 2)

    // List all products after removal
    ListProducts(inventory)
}

