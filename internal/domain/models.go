package domain

type User struct {
    ID    string
    Name  string
    Email string
}

type Product struct {
    ID    string
    Name  string
    Price float64
}

type Order struct {
    ID       string
    UserID   string
    ProductID string
    Quantity int
}